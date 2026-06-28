package services

import (
	"ego/models"
	"ego/repositories"
)

// Question metadata — mapping index ke trait dan reverse status
// 15 questions distributed: 5 Narcissism, 5 Machiavellian, 5 Psychopathy
// Indices: 0-based, matching form fields q1..q15
var questionMeta = []struct {
	trait   string
	reverse bool
}{
	// Narcissism (direct)
	{trait: "narcissism", reverse: false},
	// Machiavellian (direct)
	{trait: "machiavellian", reverse: false},
	// Psychopathy (direct)
	{trait: "psychopathy", reverse: false},
	// Narcissism (reverse)
	{trait: "narcissism", reverse: true},
	// Machiavellian (reverse)
	{trait: "machiavellian", reverse: true},
	// Psychopathy (reverse)
	{trait: "psychopathy", reverse: true},
	// Narcissism (direct)
	{trait: "narcissism", reverse: false},
	// Machiavellian (direct)
	{trait: "machiavellian", reverse: false},
	// Psychopathy (direct)
	{trait: "psychopathy", reverse: false},
	// Narcissism (direct)
	{trait: "narcissism", reverse: false},
	// Machiavellian (direct)
	{trait: "machiavellian", reverse: false},
	// Psychopathy (direct)
	{trait: "psychopathy", reverse: false},
	// Narcissism (reverse)
	{trait: "narcissism", reverse: true},
	// Machiavellian (reverse)
	{trait: "machiavellian", reverse: true},
	// Psychopathy (reverse)
	{trait: "psychopathy", reverse: true},
}

// ProcessQuizAnswers memproses 15 jawaban kuesioner dengan reverse scoring
// dan normalisasi ke persentase 0-100 per trait.
func ProcessQuizAnswers(email, nama string, answers []int) (string, error) {
	// Skala 4 poin (tanpa netral): 1=Sangat Tidak Sesuai, 4=Sangat Sesuai
	// Untuk direct items: skor mentah = nilai jawaban (1-4)
	// Untuk reverse items: skor mentah = 5 - nilai jawaban (1->4, 2->3, 3->2, 4->1)
	// Setiap trait punya 5 items, max raw score = 5*4 = 20, min = 5*1 = 5

	rawScores := map[string]int{
		"narcissism":    0,
		"machiavellian": 0,
		"psychopathy":   0,
	}

	for i, answer := range answers {
		if i >= len(questionMeta) {
			break
		}
		meta := questionMeta[i]
		score := answer
		if meta.reverse {
			score = 5 - answer // reverse: 1->4, 2->3, 3->2, 4->1
		}
		rawScores[meta.trait] += score
	}

	// Normalisasi ke persentase 0-100
	// Min raw = 5 (all 1s), Max raw = 20 (all 4s)
	// Rumus: ((raw - min) / (max - min)) * 100
	normalize := func(raw int) int {
		min := 5
		max := 20
		percent := ((raw - min) * 100) / (max - min)
		if percent < 0 {
			percent = 0
		}
		if percent > 100 {
			percent = 100
		}
		return percent
	}

	skorNarsisme := normalize(rawScores["narcissism"])
	skorMachiavellian := normalize(rawScores["machiavellian"])
	skorPsikopati := normalize(rawScores["psychopathy"])

	userID, err := repositories.InsertUser(email, nama, skorNarsisme, skorMachiavellian, skorPsikopati)
	if err != nil {
		return "", err
	}
	return userID, nil
}

// GetPaywallData mengambil data untuk halaman paywall
func GetPaywallData(id string) (*models.PaywallData, error) {
	nama, err := repositories.GetUserName(id)
	if err != nil {
		return nil, err
	}
	return &models.PaywallData{ID: id, Nama: nama}, nil
}

// GetQuizResult mengambil data hasil kuis (dengan proteksi paywall)
// dan menghasilkan narasi yang dipersonalisasi.
func GetQuizResult(id string) (*models.QuizResult, error) {
	user, err := repositories.GetUserResult(id)
	if err != nil {
		return nil, err
	}

	// Proteksi: hanya tampilkan hasil jika sudah PAID
	if user.StatusPembayaran != "PAID" {
		return nil, nil // nil menandakan belum bayar
	}

	// Generate all narrative sections
	execSummary, relProfile, kekuatan, areaPerhatian, relInsight, compatNotes, refQuestions := GenerateAllNarratives(
		user.Nama,
		user.SkorNarsisme,
		user.SkorMachiavellian,
		user.SkorPsikopati,
	)

	return &models.QuizResult{
		Nama:                user.Nama,
		Narsisme:            user.SkorNarsisme,
		Machiavellian:       user.SkorMachiavellian,
		Psikopati:           user.SkorPsikopati,
		ExecutiveSummary:    execSummary,
		RelationshipProfile: relProfile,
		Kekuatan:            kekuatan,
		AreaPerhatian:       areaPerhatian,
		RelationshipInsight: relInsight,
		CompatibilityNotes:  compatNotes,
		ReflectionQuestions: refQuestions,
	}, nil
}

// ConfirmPayment mengonfirmasi pembayaran user
func ConfirmPayment(id string) error {
	return repositories.UpdatePaymentStatus(id)
}
