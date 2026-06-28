# MBTI Test — Technical Documentation
### Versi: 2.0 | Untuk: AI Client & Developer | Basis: Teori Kognitif Carl Jung

---

## DAFTAR ISI

1. [Landasan Teori — Fungsi Kognitif Jung](#1-landasan-teori--fungsi-kognitif-jung)
2. [Arsitektur 4 Dikotomi & Mekanisme Skoring](#2-arsitektur-4-dikotomi--mekanisme-skoring)
3. [Sistem Fungsi Kognitif — Stack & Dinamika](#3-sistem-fungsi-kognitif--stack--dinamika)
4. [Matriks Pembobotan Soal](#4-matriks-pembobotan-soal)
5. [Algoritma Skoring — Pseudocode & Implementasi](#5-algoritma-skoring--pseudocode--implementasi)
6. [Skema Database](#6-skema-database)
7. [Bank Soal — Contoh & Klasifikasi](#7-bank-soal--contoh--klasifikasi)
8. [Profil 16 Tipe MBTI — Analisis Mendalam](#8-profil-16-tipe-mbti--analisis-mendalam)
9. [Validasi & Edge Cases](#9-validasi--edge-cases)
10. [API Response Schema](#10-api-response-schema)

---

## 1. Landasan Teori — Fungsi Kognitif Jung

### 1.1 Delapan Fungsi Kognitif Dasar

Carl Jung mengidentifikasi bahwa pikiran manusia memproses realitas melalui **dua dimensi utama**: cara kita *mengumpulkan informasi* (Perceiving) dan cara kita *membuat keputusan* (Judging). Masing-masing dimensi berjalan dalam dua arah: ke dalam diri (Introverted) atau ke luar dunia (Extraverted).

| Kode | Nama Lengkap | Domain | Arah | Deskripsi Esensial |
|------|-------------|--------|------|--------------------|
| `Se` | Extraverted Sensing | Perceiving | Eksternal | Mengalami realitas fisik secara langsung, saat ini, melalui indera. Respons instan terhadap stimulus konkret. |
| `Si` | Introverted Sensing | Perceiving | Internal | Mencocokkan pengalaman baru dengan memori subjektif yang tersimpan. Membangun referensi dari masa lalu. |
| `Ne` | Extraverted iNtuition | Perceiving | Eksternal | Melihat pola, koneksi, dan kemungkinan di dunia luar. Berpikir lateral dan asosiatif secara luas. |
| `Ni` | Introverted iNtuition | Perceiving | Internal | Mensintesis informasi menjadi insight mendalam tentang makna tersembunyi atau masa depan. Visi holistik. |
| `Te` | Extraverted Thinking | Judging | Eksternal | Mengorganisasi dunia eksternal secara logis dan efisien. Fokus pada sistem, prosedur, dan hasil terukur. |
| `Ti` | Introverted Thinking | Judging | Internal | Membangun kerangka logika internal yang konsisten dan presisi. Analisis mendalam, bukan efisiensi. |
| `Fe` | Extraverted Feeling | Judging | Eksternal | Menyelaraskan dengan nilai dan emosi kelompok. Menjaga harmoni sosial dan kebutuhan orang lain. |
| `Fi` | Introverted Feeling | Judging | Internal | Mengevaluasi berdasarkan nilai moral internal yang kuat. Autentisitas dan integritas personal. |

### 1.2 Hukum Oposisi Fungsi

Setiap fungsi memiliki **fungsi berlawanan** (axis) yang tidak bisa aktif secara simultan pada level sadar yang sama:

```
Se ←→ Ni   (Sensing Eksternal vs Intuisi Internal)
Si ←→ Ne   (Sensing Internal vs Intuisi Eksternal)
Te ←→ Fi   (Thinking Eksternal vs Feeling Internal)
Ti ←→ Fe   (Thinking Internal vs Feeling Eksternal)
```

> **Implikasi untuk Skoring:** Jika seseorang memiliki `Se` kuat sebagai fungsi dominan, maka `Ni` secara statistis akan muncul lemah (inferior). Ini membentuk validasi silang antar dimensi.

### 1.3 Orientasi Sikap (Attitude)

Setiap fungsi kognitif memiliki **orientasi sikap** yang menentukan apakah fungsi tersebut diekspresikan secara Extraverted (E) atau Introverted (I):

- **Extraverted functions** → Energi mengalir keluar, dapat diamati secara eksternal, cepat beradaptasi
- **Introverted functions** → Energi mengalir ke dalam, sulit diamati, kaya makna internal

---

## 2. Arsitektur 4 Dikotomi & Mekanisme Skoring

### 2.1 Empat Sumbu Dikotomi

MBTI mengoperasionalkan teori Jung menjadi 4 pasang dikotomi yang dapat diukur:

| Sumbu | Pole A | Pole B | Apa yang Diukur |
|-------|--------|--------|-----------------|
| **E/I** | Extraversion | Introversion | Arah orientasi energi psikis |
| **S/N** | Sensing | iNtuition | Preferensi fungsi persepsi |
| **T/F** | Thinking | Feeling | Preferensi fungsi penilaian |
| **J/P** | Judging | Perceiving | Gaya hidup & orientasi terhadap dunia luar |

### 2.2 Interpretasi Psikologis per Dikotomi

#### Dikotomi E/I — Orientasi Energi

```
EXTRAVERSION                         INTROVERSION
─────────────────────────────────    ──────────────────────────────────
• Energi dipulihkan oleh interaksi   • Energi dipulihkan oleh solitude
• Lebar > dalam (banyak koneksi)     • Dalam > lebar (koneksi bermakna)
• Think-out-loud (verbalisasi)       • Think-before-speak (refleksi)
• Preferensi aksi → refleksi         • Preferensi refleksi → aksi
• Dunia luar sebagai orientasi       • Dunia dalam sebagai orientasi
```

**PENTING — Kesalahpahaman Umum:** E/I bukan tentang "suka bergaul vs pemalu". Ini tentang **arah aliran energi**. Seorang Extravert bisa introvert secara sosial, dan seorang Introvert bisa sangat ekspresif. Yang diukur adalah: *di mana energi psikis diisi ulang?*

#### Dikotomi S/N — Fungsi Persepsi

```
SENSING                              iNTUITION
─────────────────────────────────    ──────────────────────────────────
• Kepercayaan pada data indrawi      • Kepercayaan pada pola & abstraksi
• Present-focused (kini & nyata)     • Future/conceptually focused
• Tahapan: konkret → abstrak         • Tahapan: abstrak → konkret
• Akurasi detail faktual             • Koneksi antar domain yang jauh
• "Apa yang ada" (eksplisit)         • "Apa yang mungkin" (implisit)
```

#### Dikotomi T/F — Fungsi Penilaian

```
THINKING                             FEELING
─────────────────────────────────    ──────────────────────────────────
• Kriteria: logika & konsistensi     • Kriteria: nilai & dampak manusia
• Objektif, impersonal               • Subjektif, personal
• Analisis sebab-akibat              • Pertimbangan harmoni & empati
• Kebenaran > kesopanan              • Kesopanan bisa = kebenaran
• Detach untuk memutuskan            • Engage untuk memutuskan
```

**PENTING:** T/F bukan tentang "cerdas vs emosional". Keduanya adalah cara rasional membuat keputusan. T menggunakan logika impersonal; F menggunakan nilai personal yang tak kalah sistematis.

#### Dikotomi J/P — Orientasi Gaya Hidup

```
JUDGING                              PERCEIVING
─────────────────────────────────    ──────────────────────────────────
• Preferensi closure & struktur      • Preferensi keterbukaan & fleksibilitas
• Rencana sebelum eksekusi           • Adaptasi saat eksekusi
• Lingkungan terorganisir            • Lingkungan organik & spontan
• Deadline sebagai patokan keras     • Deadline sebagai panduan lunak
• Fungsi Judging (T/F) di depan      • Fungsi Perceiving (S/N) di depan
```

**CATATAN KRITIS untuk Developer:** Dikotomi J/P secara teknis menentukan **fungsi mana yang diekstraversionkan** (ditampilkan ke dunia). Bukan semata tentang "terorganisir vs spontan" — ini memiliki implikasi langsung pada stack fungsi kognitif (lihat Bagian 3).

### 2.3 Mekanisme Skoring Raw per Dikotomi

Setiap jawaban soal berkontribusi pada salah satu pole dikotomi dengan bobot tertentu:

```
Skor Akhir Per Dikotomi:
  score_EI = Σ(bobot_soal × nilai_jawaban_E) - Σ(bobot_soal × nilai_jawaban_I)
  score_SN = Σ(bobot_soal × nilai_jawaban_S) - Σ(bobot_soal × nilai_jawaban_N)
  score_TF = Σ(bobot_soal × nilai_jawaban_T) - Σ(bobot_soal × nilai_jawaban_F)
  score_JP = Σ(bobot_soal × nilai_jawaban_J) - Σ(bobot_soal × nilai_jawaban_P)

Klasifikasi Tipe:
  E/I → score_EI > 0 ? "E" : "I"
  S/N → score_SN > 0 ? "S" : "N"
  T/F → score_TF > 0 ? "T" : "F"
  J/P → score_JP > 0 ? "J" : "P"
```

### 2.4 Skor Kekuatan Preferensi (Preference Clarity Index)

Skor raw dikonversi menjadi **persentase kekuatan** untuk menunjukkan seberapa jelas preferensi:

| Range Persentase | Label Kekuatan | Interpretasi |
|-----------------|----------------|--------------|
| 0% – 25% | Slight (Sedikit) | Preferensi lemah, hampir di tengah; fleksibel di kedua sisi |
| 26% – 50% | Moderate (Sedang) | Preferensi cukup jelas, masih fleksibel |
| 51% – 74% | Clear (Jelas) | Preferensi jelas dan konsisten |
| 75% – 100% | Very Clear (Sangat Jelas) | Preferensi sangat kuat dan dominan |

```
Formula PCI (Preference Clarity Index):
  pci = |score_dikotomi| / max_possible_score × 100
  
  Contoh:
    max_possible_score = jumlah_soal × bobot_max × skala_max
    pci_EI = |score_EI| / max_EI × 100
```

---

## 3. Sistem Fungsi Kognitif — Stack & Dinamika

### 3.1 Aturan Pembentukan Stack

Setiap tipe MBTI memiliki **4 fungsi kognitif** yang tersusun dalam hierarki (stack):

```
Posisi 1: Dominant    → Fungsi paling kuat, paling sadar, identitas utama
Posisi 2: Auxiliary   → Fungsi pendukung, menyeimbangkan dominan
Posisi 3: Tertiary    → Berkembang di usia dewasa, lebih tidak sadar
Posisi 4: Inferior    → Fungsi paling lemah, sumber stres & pertumbuhan
```

**Aturan Stack:**
1. Dominan dan Inferior selalu dari **axis yang sama** (Se↔Ni, Si↔Ne, Te↔Fi, Ti↔Fe)
2. Auxiliary dan Tertiary selalu dari **axis yang sama** (Se↔Ni, Si↔Ne, Te↔Fi, Ti↔Fe)
3. Stack harus berisi **campuran Perceiving dan Judging** (tidak boleh semuanya P atau J)
4. Stack harus berisi **campuran Extraverted dan Introverted** (idealnya 2E+2I atau dengan bias)

### 3.2 Tabel Stack Lengkap 16 Tipe

| Tipe | Dominant | Auxiliary | Tertiary | Inferior |
|------|----------|-----------|----------|---------|
| **INTJ** | Ni | Te | Fi | Se |
| **INTP** | Ti | Ne | Si | Fe |
| **ENTJ** | Te | Ni | Se | Fi |
| **ENTP** | Ne | Ti | Fe | Si |
| **INFJ** | Ni | Fe | Ti | Se |
| **INFP** | Fi | Ne | Si | Te |
| **ENFJ** | Fe | Ni | Se | Ti |
| **ENFP** | Ne | Fi | Te | Si |
| **ISTJ** | Si | Te | Fi | Ne |
| **ISFJ** | Si | Fe | Ti | Ne |
| **ESTJ** | Te | Si | Ne | Fi |
| **ESFJ** | Fe | Si | Ne | Ti |
| **ISTP** | Ti | Se | Ni | Fe |
| **ISFP** | Fi | Se | Ni | Te |
| **ESTP** | Se | Ti | Fe | Ni |
| **ESFP** | Se | Fi | Te | Ni |

### 3.3 Aturan Derivasi Otomatis dari 4 Huruf

Developer dapat menurunkan stack secara algoritmik dari 4 huruf MBTI menggunakan aturan berikut:

```python
def derive_cognitive_stack(mbti: str) -> dict:
    """
    Menurunkan stack fungsi kognitif dari 4 huruf MBTI.
    Aturan berdasarkan teori Jungian.
    """
    e_i = mbti[0]   # 'E' atau 'I'
    s_n = mbti[1]   # 'S' atau 'N'
    t_f = mbti[2]   # 'T' atau 'F'
    j_p = mbti[3]   # 'J' atau 'P'
    
    # Tentukan fungsi persepsi
    perceiving = "Se" if s_n == "S" else "Ne"
    perceiving_i = "Si" if s_n == "S" else "Ni"
    
    # Tentukan fungsi penilaian
    judging = "Te" if t_f == "T" else "Fe"
    judging_i = "Ti" if t_f == "T" else "Fi"
    
    if e_i == "E":
        if j_p == "J":
            # Extravert-Judging: Dominant = fungsi judging ekstraverted
            dominant   = judging       # Te atau Fe (ekstraverted)
            auxiliary  = perceiving_i  # Ni atau Si (introverted)
            tertiary   = axis_opposite(perceiving_i)  # Se atau Ne
            inferior   = axis_opposite(dominant)       # Fi atau Ti
        else:  # j_p == "P"
            # Extravert-Perceiving: Dominant = fungsi perceiving ekstraverted
            dominant   = perceiving    # Se atau Ne (ekstraverted)
            auxiliary  = judging_i     # Ti atau Fi (introverted)
            tertiary   = axis_opposite(judging_i)     # Fe atau Te
            inferior   = axis_opposite(dominant)       # Ni atau Si
    else:  # e_i == "I"
        if j_p == "J":
            # Introvert-Judging: Dominant = fungsi judging introverted
            dominant   = judging_i     # Ti atau Fi (introverted)
            auxiliary  = perceiving    # Se atau Ne (ekstraverted)
            tertiary   = axis_opposite(perceiving)    # Ni atau Si
            inferior   = axis_opposite(dominant)       # Fe atau Te
        else:  # j_p == "P"
            # Introvert-Perceiving: Dominant = fungsi perceiving introverted
            dominant   = perceiving_i  # Si atau Ni (introverted)
            auxiliary  = judging       # Te atau Fe (ekstraverted)
            tertiary   = axis_opposite(judging)       # Fi atau Ti
            inferior   = axis_opposite(dominant)       # Se atau Ne
    
    return {
        "dominant": dominant,
        "auxiliary": auxiliary,
        "tertiary": tertiary,
        "inferior": inferior
    }

AXIS_OPPOSITES = {
    "Se": "Ni", "Ni": "Se",
    "Si": "Ne", "Ne": "Si",
    "Te": "Fi", "Fi": "Te",
    "Ti": "Fe", "Fe": "Ti"
}

def axis_opposite(func: str) -> str:
    return AXIS_OPPOSITES[func]
```

### 3.4 Dinamika Fungsi: Shadow & Trickster

Selain 4 fungsi utama (ego stack), setiap tipe memiliki **4 fungsi shadow** (unconscious stack):

| Posisi | Ego Stack | Shadow Stack | Dinamika |
|--------|-----------|--------------|---------|
| 1 | Dominant | Opposing Role (5th) | Defensif saat terancam |
| 2 | Auxiliary | Critical Parent (6th) | Perfeksionis internal |
| 3 | Tertiary | Trickster (7th) | Miskomunikasi & tipu daya |
| 4 | Inferior | Demon (8th) | Sisi gelap, destruktif |

> **Untuk MVP:** Shadow functions tidak perlu diimplementasikan dalam skoring. Cukup catatan psikologis dalam profil tipe.

---

## 4. Matriks Pembobotan Soal

### 4.1 Kategori Pertanyaan & Bobot

Tidak semua pertanyaan memiliki bobot yang sama. Pertanyaan yang mengukur **fungsi dominan/auxiliary** lebih diskriminatif dibanding yang mengukur tertiary/inferior.

| Kategori | Bobot | Alasan Psikologis |
|----------|-------|------------------|
| **Core Function** | 3.0 | Mengukur fungsi dominan atau auxiliary secara langsung |
| **Preference Behavior** | 2.0 | Mengukur perilaku yang berkorelasi kuat dengan dikotomi |
| **Lifestyle Orientation** | 1.5 | Mengukur gaya hidup J/P dan orientasi E/I |
| **Situational** | 1.0 | Pertanyaan skenario yang lebih kontekstual |
| **Self-Report Baseline** | 0.5 | Pertanyaan self-report yang rawan social desirability bias |

### 4.2 Skala Respons

Gunakan skala **6-poin Likert** (menghindari tengah-tengah/neutral midpoint) untuk memaksa arah preferensi:

```
Nilai Skala:
  1 = Sangat setuju dengan pernyataan A (bobot penuh ke pole A)
  2 = Setuju dengan pernyataan A        (bobot 0.67 ke pole A)
  3 = Sedikit condong ke A              (bobot 0.33 ke pole A)
  4 = Sedikit condong ke B              (bobot 0.33 ke pole B)
  5 = Setuju dengan pernyataan B        (bobot 0.67 ke pole B)
  6 = Sangat setuju dengan pernyataan B (bobot penuh ke pole B)
```

**Alternatif: Forced Choice Format**
Lebih diskriminatif secara psikometri — responden memilih salah satu dari dua pernyataan:

```json
{
  "question_id": "Q_EI_001",
  "format": "forced_choice",
  "option_a": {
    "text": "Saya merasa berenergi setelah menghabiskan waktu bersama orang banyak",
    "pole": "E",
    "weight": 2.0
  },
  "option_b": {
    "text": "Saya merasa berenergi setelah menghabiskan waktu sendiri",
    "pole": "I",
    "weight": 2.0
  }
}
```

### 4.3 Distribusi Soal yang Direkomendasikan

**Total soal: 60–93 soal** (standar instrument psikometri MBTI)

| Dikotomi | Jumlah Soal | Distribusi Bobot |
|----------|-------------|-----------------|
| E/I | 15–21 | 30% Core, 40% Behavior, 30% Situational |
| S/N | 18–24 | 35% Core, 35% Behavior, 30% Situational |
| T/F | 18–24 | 35% Core, 35% Behavior, 30% Situational |
| J/P | 15–21 | 30% Core, 40% Lifestyle, 30% Situational |
| **Total** | **66–90** | |

> **Mengapa S/N dan T/F lebih banyak soal?** Kedua dikotomi ini memiliki overlap dengan bias sosial dan gender (T/F khususnya), sehingga butuh lebih banyak item untuk akurasi.

### 4.4 Matriks Tag Soal

Setiap soal harus di-tag dengan metadata berikut untuk filtering dan analisis:

```json
{
  "question_id": "string",          // Unik, e.g., "Q_SN_012"
  "dikotomi": "EI|SN|TF|JP",
  "pole_primary": "E|I|S|N|T|F|J|P",
  "cognitive_function": "Se|Si|Ne|Ni|Te|Ti|Fe|Fi|null",
  "weight": 0.5 | 1.0 | 1.5 | 2.0 | 3.0,
  "format": "likert_6|forced_choice|scenario",
  "reverse_scored": true | false,
  "bias_risk": "social_desirability|gender|culture|none",
  "validated": true | false,
  "domain": "work|social|cognitive|emotional|lifestyle"
}
```

---

## 5. Algoritma Skoring — Pseudocode & Implementasi

### 5.1 Struktur Data Input

```go
// models.go

package mbti

import "time"

// QuestionResponse merepresentasikan satu jawaban dari user.
// AnswerValue: 1–6 untuk Likert, atau 0 (A) / 1 (B) untuk forced choice.
type QuestionResponse struct {
	QuestionID  string  `json:"question_id"`
	AnswerValue float64 `json:"answer_value"`         // 1–6 Likert | 0=A / 1=B forced choice
	TimeTakenMs *int64  `json:"time_taken_ms,omitempty"` // Opsional, untuk analisis response time
}

// SessionMetadata menyimpan konteks perangkat & lokal pengguna.
type SessionMetadata struct {
	Device string `json:"device"`
	Locale string `json:"locale"`
}

// TestSession merepresentasikan satu sesi pengerjaan tes MBTI.
type TestSession struct {
	SessionID   string              `json:"session_id"`
	UserID      *string             `json:"user_id,omitempty"`
	Responses   []QuestionResponse  `json:"responses"`
	StartedAt   time.Time           `json:"started_at"`
	CompletedAt *time.Time          `json:"completed_at,omitempty"`
	Metadata    *SessionMetadata    `json:"metadata,omitempty"`
}
```

### 5.2 Algoritma Skoring Utama

```go
// scoring.go

package mbti

import (
	"math"
)

// QuestionDefinition mendefinisikan metadata setiap soal dalam question bank.
type QuestionDefinition struct {
	QuestionID   string  `json:"question_id"`
	Dikotomi     string  `json:"dikotomi"`      // "EI" | "SN" | "TF" | "JP"
	PolePrimary  string  `json:"pole_primary"`  // "E"|"I"|"S"|"N"|"T"|"F"|"J"|"P"
	Weight       float64 `json:"weight"`
	Format       string  `json:"format"`        // "likert_6" | "forced_choice"
	ReverseScored bool   `json:"reverse_scored"`
}

// DikotomiScore menyimpan hasil skoring untuk satu sumbu dikotomi.
type DikotomiScore struct {
	RawScore    float64 `json:"raw_score"`    // Positif = pole pertama, Negatif = pole kedua
	PoleAScore  float64 `json:"pole_a_score"` // e.g., untuk EI: total skor E
	PoleBScore  float64 `json:"pole_b_score"` // e.g., untuk EI: total skor I
	MaxPossible float64 `json:"max_possible"`
	Preference  string  `json:"preference"`   // "E" atau "I"
	PCI         float64 `json:"pci"`          // Preference Clarity Index (0–100)
	Strength    string  `json:"strength"`     // "slight"|"moderate"|"clear"|"very_clear"
}

// CognitiveStack menyimpan urutan 4 fungsi kognitif hasil derivasi.
type CognitiveStack struct {
	Dominant  string `json:"dominant"`
	Auxiliary string `json:"auxiliary"`
	Tertiary  string `json:"tertiary"`
	Inferior  string `json:"inferior"`
}

// ReliabilityIndicators menyimpan metrik keandalan sesi tes.
type ReliabilityIndicators struct {
	CompletionRate     float64 `json:"completion_rate"`      // % soal dijawab
	AvgResponseTimeMs  int64   `json:"avg_response_time_ms"`
	InconsistencyScore float64 `json:"inconsistency_score"`  // 0–100, semakin tinggi semakin inkonsisten
	IsReliable         bool    `json:"is_reliable"`
}

// MBTIResult adalah output akhir kalkulasi satu sesi tes.
type MBTIResult struct {
	Type                 string                `json:"type"` // e.g., "INTJ"
	Scores               map[string]DikotomiScore `json:"scores"`
	CognitiveStack       CognitiveStack        `json:"cognitive_stack"`
	ReliabilityIndicators ReliabilityIndicators `json:"reliability_indicators"`
}

// accumulator adalah struktur internal untuk menampung skor per dikotomi.
type accumulator struct {
	poleA float64
	poleB float64
	max   float64
}

// likertContribution memetakan nilai skala Likert 1–6 ke bobot kontribusi (0.0–1.0).
var likertContribution = map[int]float64{
	1: 1.00, // Sangat kuat ke pole_primary
	2: 0.67, // Kuat ke pole_primary
	3: 0.33, // Lemah ke pole_primary
	4: 0.33, // Lemah ke pole_opposite
	5: 0.67, // Kuat ke pole_opposite
	6: 1.00, // Sangat kuat ke pole_opposite
}

// buildDikotomiScore menghitung DikotomiScore dari akumulator satu sumbu.
func buildDikotomiScore(acc accumulator, poleALetter, poleBLetter string) DikotomiScore {
	rawScore := acc.poleA - acc.poleB
	preference := poleALetter
	if rawScore < 0 {
		preference = poleBLetter
	}

	pci := 0.0
	if acc.max > 0 {
		pci = math.Abs(rawScore) / acc.max * 100
	}
	pci = math.Round(pci*10) / 10 // 1 desimal

	strength := "very_clear"
	switch {
	case pci <= 25:
		strength = "slight"
	case pci <= 50:
		strength = "moderate"
	case pci <= 75:
		strength = "clear"
	}

	return DikotomiScore{
		RawScore:    rawScore,
		PoleAScore:  acc.poleA,
		PoleBScore:  acc.poleB,
		MaxPossible: acc.max,
		Preference:  preference,
		PCI:         pci,
		Strength:    strength,
	}
}

// CalculateMBTI menghitung hasil MBTI lengkap dari satu sesi tes.
func CalculateMBTI(session TestSession, questionBank []QuestionDefinition) MBTIResult {

	// Step 1: Buat map question_id → definisi soal
	questionMap := make(map[string]QuestionDefinition, len(questionBank))
	for _, q := range questionBank {
		questionMap[q.QuestionID] = q
	}

	// Step 2: Inisialisasi akumulator skor
	// A=E, A=S, A=T, A=J (pole pertama masing-masing dikotomi)
	accumulators := map[string]*accumulator{
		"EI": {},
		"SN": {},
		"TF": {},
		"JP": {},
	}

	// Step 3: Proses setiap respons
	for _, response := range session.Responses {
		q, ok := questionMap[response.QuestionID]
		if !ok {
			continue
		}

		acc := accumulators[q.Dikotomi]

		switch q.Format {
		case "likert_6":
			// Skala 1–6: nilai 1–3 ke pole_primary, 4–6 ke pole_opposite
			raw := int(response.AnswerValue)
			adjusted := raw
			if q.ReverseScored {
				adjusted = 7 - raw
			}
			contribution := likertContribution[adjusted]
			weighted := contribution * q.Weight
			if adjusted <= 3 {
				acc.poleA += weighted
			} else {
				acc.poleB += weighted
			}

		case "forced_choice":
			// AnswerValue: 0 = pilih A (pole_primary), 1 = pilih B (pole_opposite)
			choseA := response.AnswerValue == 0
			primaryChosen := choseA
			if q.ReverseScored {
				primaryChosen = !choseA
			}
			if primaryChosen {
				acc.poleA += q.Weight
			} else {
				acc.poleB += q.Weight
			}
		}

		// Track max possible
		acc.max += q.Weight
	}

	// Step 4: Hitung DikotomiScore untuk setiap dikotomi
	scores := map[string]DikotomiScore{
		"EI": buildDikotomiScore(*accumulators["EI"], "E", "I"),
		"SN": buildDikotomiScore(*accumulators["SN"], "S", "N"),
		"TF": buildDikotomiScore(*accumulators["TF"], "T", "F"),
		"JP": buildDikotomiScore(*accumulators["JP"], "J", "P"),
	}

	// Step 5: Derive tipe MBTI
	mbtiType := scores["EI"].Preference +
		scores["SN"].Preference +
		scores["TF"].Preference +
		scores["JP"].Preference

	// Step 6: Derive cognitive stack
	cognitiveStack := DeriveCognitiveStack(mbtiType)

	// Step 7: Hitung reliability indicators
	completionRate := float64(len(session.Responses)) / float64(len(questionBank)) * 100
	completionRate = math.Round(completionRate*10) / 10

	var totalMs int64
	var msCount int64
	for _, r := range session.Responses {
		if r.TimeTakenMs != nil {
			totalMs += *r.TimeTakenMs
			msCount++
		}
	}
	avgResponseMs := int64(0)
	if msCount > 0 {
		avgResponseMs = totalMs / msCount
	}

	inconsistencyScore := CalculateInconsistencyScore(session.Responses, questionMap)

	return MBTIResult{
		Type:           mbtiType,
		Scores:         scores,
		CognitiveStack: cognitiveStack,
		ReliabilityIndicators: ReliabilityIndicators{
			CompletionRate:     completionRate,
			AvgResponseTimeMs:  avgResponseMs,
			InconsistencyScore: inconsistencyScore,
			IsReliable:         completionRate >= 90 && inconsistencyScore <= 35,
		},
	}
}
```

### 5.3 Deteksi Inkonsistensi Jawaban

```go
// inconsistency.go

package mbti

import "math"

// AnchorPair mendefinisikan sepasang soal yang seharusnya konsisten satu sama lain.
type AnchorPair struct {
	QuestionAID string // ID soal pertama
	QuestionBID string // ID soal kedua
	Expected    string // "same" | "opposite"
}

// anchorPairs adalah daftar pasangan soal validasi.
// Populate dari question bank sesuai desain instrumen.
var anchorPairs = []AnchorPair{
	// {QuestionAID: "Q_EI_001", QuestionBID: "Q_EI_007", Expected: "same"},
	// {QuestionAID: "Q_SN_003", QuestionBID: "Q_SN_015", Expected: "opposite"},
	// ... populate dari question bank
}

// CalculateInconsistencyScore menghitung skor inkonsistensi jawaban (0–100).
//
// Beberapa soal dirancang sebagai "anchor pairs" —
// pasangan soal yang mengukur konstruk yang sama dari sudut berbeda.
// Jika jawaban keduanya berlawanan secara drastis → inkonsistensi tinggi.
//
// Skor inkonsistensi 0–100. Threshold reliable: <= 35.
func CalculateInconsistencyScore(
	responses []QuestionResponse,
	questionMap map[string]QuestionDefinition,
) float64 {
	// Buat map cepat question_id → AnswerValue
	responseMap := make(map[string]float64, len(responses))
	for _, r := range responses {
		responseMap[r.QuestionID] = r.AnswerValue
	}

	var totalInconsistency float64
	pairCount := 0

	for _, pair := range anchorPairs {
		v1, ok1 := responseMap[pair.QuestionAID]
		v2, ok2 := responseMap[pair.QuestionBID]
		if !ok1 || !ok2 {
			continue
		}

		if pair.Expected == "same" {
			// Keduanya harus arah yang sama (sama-sama < 3.5 atau sama-sama > 3.5)
			inconsistent := (v1 < 3.5) != (v2 < 3.5)
			magnitude := math.Abs(v1 - v2)
			if inconsistent {
				totalInconsistency += magnitude * 10
			}
		} else {
			// Keduanya harus berlawanan arah
			consistent := (v1 < 3.5) != (v2 < 3.5)
			magnitude := math.Abs(v1 - (7 - v2))
			if !consistent {
				totalInconsistency += magnitude * 10
			}
		}

		pairCount++
	}

	if pairCount == 0 {
		return 0
	}
	return math.Min(100, totalInconsistency/float64(pairCount))
}
```

---

## 6. Skema Database

### 6.1 Entity Relationship

```
users (opsional, jika authenticated)
  └──< test_sessions
         └──< session_responses
         └──> mbti_results
              └──> mbti_type_profiles

questions
  └──< question_options (untuk forced choice)
  └──< anchor_pairs

mbti_type_profiles
  └──< cognitive_functions_stack
```

### 6.2 DDL SQL (PostgreSQL)

```sql
-- =============================================
-- TABEL PERTANYAAN
-- =============================================
CREATE TABLE questions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_code   VARCHAR(20) UNIQUE NOT NULL,  -- e.g., "Q_EI_001"
    dikotomi        VARCHAR(2) NOT NULL CHECK (dikotomi IN ('EI','SN','TF','JP')),
    pole_primary    VARCHAR(1) NOT NULL CHECK (pole_primary IN ('E','I','S','N','T','F','J','P')),
    cognitive_fn    VARCHAR(2),                    -- e.g., 'Se', 'Ni', NULL jika tidak spesifik
    weight          DECIMAL(3,1) NOT NULL DEFAULT 1.0,
    format          VARCHAR(20) NOT NULL CHECK (format IN ('likert_6','forced_choice','scenario')),
    reverse_scored  BOOLEAN NOT NULL DEFAULT FALSE,
    bias_risk       VARCHAR(30),
    domain          VARCHAR(20),
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    validated       BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Teks pertanyaan multibahasa
CREATE TABLE question_translations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id     UUID NOT NULL REFERENCES questions(id),
    locale          VARCHAR(10) NOT NULL,           -- e.g., "id", "en", "ja"
    question_text   TEXT NOT NULL,
    option_a_text   TEXT,                           -- Untuk forced choice
    option_b_text   TEXT,
    UNIQUE(question_id, locale)
);

-- Anchor pairs untuk deteksi inkonsistensi
CREATE TABLE anchor_pairs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    question_a_id   UUID NOT NULL REFERENCES questions(id),
    question_b_id   UUID NOT NULL REFERENCES questions(id),
    expected_relation VARCHAR(10) NOT NULL CHECK (expected_relation IN ('same','opposite'))
);

-- =============================================
-- TABEL SESI TES
-- =============================================
CREATE TABLE test_sessions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID,                           -- NULL jika anonymous
    session_token   VARCHAR(64) UNIQUE NOT NULL,
    locale          VARCHAR(10) NOT NULL DEFAULT 'id',
    started_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at    TIMESTAMPTZ,
    device_type     VARCHAR(20),
    is_completed    BOOLEAN NOT NULL DEFAULT FALSE,
    is_invalidated  BOOLEAN NOT NULL DEFAULT FALSE,
    metadata        JSONB
);

-- =============================================
-- TABEL RESPONS INDIVIDUAL
-- =============================================
CREATE TABLE session_responses (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id      UUID NOT NULL REFERENCES test_sessions(id),
    question_id     UUID NOT NULL REFERENCES questions(id),
    answer_value    DECIMAL(4,2) NOT NULL,          -- 1-6 untuk Likert, 0/1 untuk forced choice
    answered_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    time_taken_ms   INTEGER,
    UNIQUE(session_id, question_id)
);

-- =============================================
-- TABEL HASIL MBTI
-- =============================================
CREATE TABLE mbti_results (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id          UUID NOT NULL REFERENCES test_sessions(id) UNIQUE,
    mbti_type           VARCHAR(4) NOT NULL,

    -- Raw scores per dikotomi
    ei_raw_score        DECIMAL(8,2) NOT NULL,      -- Positif = E, Negatif = I
    sn_raw_score        DECIMAL(8,2) NOT NULL,      -- Positif = S, Negatif = N
    tf_raw_score        DECIMAL(8,2) NOT NULL,      -- Positif = T, Negatif = F
    jp_raw_score        DECIMAL(8,2) NOT NULL,      -- Positif = J, Negatif = P

    -- Preference Clarity Index (0-100)
    ei_pci              DECIMAL(5,1) NOT NULL,
    sn_pci              DECIMAL(5,1) NOT NULL,
    tf_pci              DECIMAL(5,1) NOT NULL,
    jp_pci              DECIMAL(5,1) NOT NULL,

    -- Strength labels
    ei_strength         VARCHAR(12) NOT NULL,
    sn_strength         VARCHAR(12) NOT NULL,
    tf_strength         VARCHAR(12) NOT NULL,
    jp_strength         VARCHAR(12) NOT NULL,

    -- Cognitive stack
    cognitive_dominant  VARCHAR(2) NOT NULL,
    cognitive_auxiliary VARCHAR(2) NOT NULL,
    cognitive_tertiary  VARCHAR(2) NOT NULL,
    cognitive_inferior  VARCHAR(2) NOT NULL,

    -- Reliability
    completion_rate     DECIMAL(5,1) NOT NULL,
    avg_response_ms     INTEGER,
    inconsistency_score DECIMAL(5,1) NOT NULL DEFAULT 0,
    is_reliable         BOOLEAN NOT NULL DEFAULT TRUE,

    calculated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- =============================================
-- TABEL PROFIL 16 TIPE (Data Statis)
-- =============================================
CREATE TABLE mbti_type_profiles (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mbti_type           VARCHAR(4) UNIQUE NOT NULL,
    nickname            VARCHAR(50) NOT NULL,       -- e.g., "The Architect"
    nickname_id         VARCHAR(50) NOT NULL,       -- e.g., "Sang Arsitek"
    cognitive_dominant  VARCHAR(2) NOT NULL,
    cognitive_auxiliary VARCHAR(2) NOT NULL,
    cognitive_tertiary  VARCHAR(2) NOT NULL,
    cognitive_inferior  VARCHAR(2) NOT NULL,
    temperament_group   VARCHAR(2) NOT NULL,        -- NT, NF, SJ, SP
    interaction_style   VARCHAR(20) NOT NULL,       -- "Chart-the-Course", "Behind-the-Scenes", dll
    description_short   TEXT NOT NULL,
    description_full    TEXT NOT NULL,
    strengths           JSONB NOT NULL,             -- Array of strings
    challenges          JSONB NOT NULL,
    career_themes       JSONB NOT NULL,
    relationship_style  TEXT NOT NULL,
    growth_areas        TEXT NOT NULL,
    famous_examples     JSONB NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- =============================================
-- INDEKS PERFORMA
-- =============================================
CREATE INDEX idx_session_responses_session ON session_responses(session_id);
CREATE INDEX idx_mbti_results_type ON mbti_results(mbti_type);
CREATE INDEX idx_test_sessions_user ON test_sessions(user_id) WHERE user_id IS NOT NULL;
CREATE INDEX idx_test_sessions_token ON test_sessions(session_token);
```

### 6.3 Skema NoSQL Alternatif (MongoDB)

```javascript
// Collection: questions
{
  _id: ObjectId,
  question_code: "Q_EI_001",
  dikotomi: "EI",
  pole_primary: "E",
  cognitive_fn: "Se",    // nullable
  weight: 2.0,
  format: "forced_choice",
  reverse_scored: false,
  translations: {
    id: {
      text: "...",
      option_a: "...",
      option_b: "..."
    },
    en: { text: "...", option_a: "...", option_b: "..." }
  },
  tags: {
    bias_risk: "none",
    domain: "social",
    validated: true
  }
}

// Collection: test_results
{
  _id: ObjectId,
  session_id: "uuid",
  mbti_type: "INTJ",
  scores: {
    EI: { raw: 12.5, pci: 68.2, preference: "I", strength: "clear" },
    SN: { raw: -8.0, pci: 44.4, preference: "N", strength: "moderate" },
    TF: { raw: 15.0, pci: 75.0, preference: "T", strength: "very_clear" },
    JP: { raw: -6.5, pci: 40.6, preference: "J", strength: "moderate" }
  },
  cognitive_stack: {
    dominant: "Ni",
    auxiliary: "Te",
    tertiary: "Fi",
    inferior: "Se"
  },
  reliability: {
    completion_rate: 100,
    avg_response_ms: 4200,
    inconsistency_score: 12,
    is_reliable: true
  },
  calculated_at: ISODate("2025-01-01T00:00:00Z")
}
```

---

## 7. Bank Soal — Contoh & Klasifikasi

### 7.1 Contoh Soal per Dikotomi

**DIKOTOMI E/I:**

```json
[
  {
    "code": "Q_EI_001",
    "format": "forced_choice",
    "weight": 2.0,
    "cognitive_fn": null,
    "id": {
      "text": "Setelah hari yang panjang bersama banyak orang, saya cenderung...",
      "option_a": "Merasa berenergi dan ingin melanjutkan bersosialisasi",
      "option_b": "Merasa lelah dan butuh waktu sendiri untuk recharge"
    },
    "pole_a": "E", "pole_b": "I",
    "domain": "social"
  },
  {
    "code": "Q_EI_002",
    "format": "likert_6",
    "weight": 1.5,
    "cognitive_fn": null,
    "id": {
      "text": "Saya lebih sering memikirkan sesuatu secara mendalam di dalam kepala sebelum mengungkapkannya.",
      "scale_a": "Sangat tidak setuju (saya lebih sering berbicara untuk berpikir)",
      "scale_b": "Sangat setuju (saya perlu merenung dulu sebelum berbicara)"
    },
    "pole_primary": "I",
    "domain": "cognitive"
  }
]
```

**DIKOTOMI S/N:**

```json
[
  {
    "code": "Q_SN_001",
    "format": "forced_choice",
    "weight": 3.0,
    "cognitive_fn": "Ne",
    "id": {
      "text": "Ketika membaca sebuah fakta menarik, saya cenderung...",
      "option_a": "Mengingat detail faktanya dengan akurat",
      "option_b": "Langsung memikirkan koneksi dan implikasinya ke hal lain"
    },
    "pole_a": "S", "pole_b": "N",
    "domain": "cognitive"
  },
  {
    "code": "Q_SN_002",
    "format": "forced_choice",
    "weight": 2.5,
    "cognitive_fn": "Si",
    "id": {
      "text": "Dalam belajar sesuatu yang baru, saya lebih nyaman...",
      "option_a": "Mengikuti petunjuk langkah demi langkah secara terurut",
      "option_b": "Memahami gambaran besarnya dulu, baru detail-detailnya"
    },
    "pole_a": "S", "pole_b": "N",
    "domain": "cognitive"
  }
]
```

**DIKOTOMI T/F:**

```json
[
  {
    "code": "Q_TF_001",
    "format": "forced_choice",
    "weight": 2.0,
    "cognitive_fn": "Te",
    "id": {
      "text": "Ketika seorang teman meminta pendapat tentang keputusan penting, saya lebih dulu...",
      "option_a": "Menganalisis pro-kontra secara objektif dan logis",
      "option_b": "Mempertimbangkan bagaimana perasaan mereka dan apa yang terbaik untuk mereka"
    },
    "pole_a": "T", "pole_b": "F",
    "domain": "social"
  },
  {
    "code": "Q_TF_002",
    "format": "likert_6",
    "weight": 1.5,
    "cognitive_fn": "Fi",
    "id": {
      "text": "Saya merasa tidak nyaman ketika harus membuat keputusan yang logis namun menyakiti perasaan orang lain.",
      "scale_a": "Sangat tidak setuju (keputusan harus berdasar logika, terlepas dari perasaan)",
      "scale_b": "Sangat setuju (dampak emosional sangat penting bagi saya)"
    },
    "pole_primary": "F",
    "domain": "emotional"
  }
]
```

**DIKOTOMI J/P:**

```json
[
  {
    "code": "Q_JP_001",
    "format": "forced_choice",
    "weight": 2.0,
    "cognitive_fn": null,
    "id": {
      "text": "Saya lebih suka...",
      "option_a": "Memiliki rencana yang jelas sebelum memulai suatu proyek",
      "option_b": "Memulai dengan fleksibel dan menyesuaikan rencana seiring berjalan"
    },
    "pole_a": "J", "pole_b": "P",
    "domain": "lifestyle"
  }
]
```

---

## 8. Profil 16 Tipe MBTI — Analisis Mendalam

> **Format setiap profil:** Kode → Stack Fungsi → Dinamika Psikologis → Kekuatan/Tantangan → Karir → Relasi

---

### 8.1 INTJ — The Architect (Sang Arsitek)

**Stack Kognitif:** `Ni` → `Te` → `Fi` → `Se`

| Posisi | Fungsi | Deskripsi Peran dalam INTJ |
|--------|--------|---------------------------|
| Dominant | **Ni** (Introverted iNtuition) | Mesin penggerak INTJ. Menyerap pola kompleks dari lingkungan dan mensintesisnya menjadi visi jangka panjang yang koheren. INTJ tidak "berpikir" tentang masa depan—mereka *melihatnya* dengan kejelasan yang mengejutkan, meski sulit dijelaskan ke orang lain. |
| Auxiliary | **Te** (Extraverted Thinking) | Eksekutor visi Ni. Te mengorganisir dunia eksternal: membuat sistem, mengoptimalkan proses, mengeksekusi rencana dengan efisiensi. Memberi INTJ kemampuan untuk tidak hanya bermimpi tapi juga membangun. |
| Tertiary | **Fi** (Introverted Feeling) | Kompas moral personal yang sering tidak disadari INTJ. Walau terlihat dingin, INTJ memiliki nilai-nilai yang sangat kuat dan pribadi. Berkembang seiring usia—INTJ yang matang lebih sadar akan perasaan batin mereka. |
| Inferior | **Se** (Extraverted Sensing) | Titik kelemahan bawah sadar. INTJ bisa mengabaikan kebutuhan fisik, detail sensoris, atau kenikmatan saat-ini demi visi masa depan. Saat stres, Se bisa meledak: impulsif, overindulge, atau hyper-aware terhadap sensasi fisik. |

**Dinamika Psikologis:**
INTJ beroperasi dari "masa depan ke masa kini"—Ni menentukan kemana, Te membangun jalan. Ini membuat mereka tampak **determinatif dan sering keras kepala**: bukan karena arogan, tapi karena Ni mereka telah "melihat" hasilnya dan Te merasa tidak perlu mendiskusikan hal yang sudah *jelas*. Konflik umum: ketidakmampuan untuk menjelaskan bagaimana mereka tahu apa yang mereka tahu (insight Ni sering tak dapat diverbalkan dengan mudah).

**Profil Ringkas:**

```json
{
  "mbti_type": "INTJ",
  "nickname": "The Architect",
  "nickname_id": "Sang Arsitek",
  "temperament": "NT (Rational)",
  "population_estimate": "2-4%",
  "cognitive_stack": ["Ni", "Te", "Fi", "Se"],
  "core_drive": "Mengimplementasikan visi jangka panjang dengan sistem yang efisien",
  "strengths": [
    "Kemampuan berpikir strategis jangka panjang yang langka",
    "Determinasi tinggi dalam mengeksekusi rencana",
    "Standar kualitas yang sangat tinggi",
    "Kemampuan melihat kelemahan sistem dengan cepat",
    "Independensi intelektual yang kuat"
  ],
  "challenges": [
    "Kesulitan menjelaskan intuisi kepada orang lain",
    "Impatient terhadap 'inefisiensi' orang lain",
    "Mengabaikan kebutuhan emosional diri sendiri dan orang lain (Fi inferior rendah)",
    "Rentan terhadap perfeksionisme yang menghambat",
    "Terlihat dingin atau arogan meski niatnya baik"
  ],
  "career_themes": ["Strategy", "Engineering", "Research", "Architecture", "Law", "Academia"],
  "relationship_note": "Berkomitmen dalam, loyal ekstrem, namun kesulitan mengekspresikan afeksi secara verbal. Pasangan perlu memahami bahwa tindakan adalah bahasa cinta INTJ."
}
```

---

### 8.2 INTP — The Logician (Sang Logikawan)

**Stack Kognitif:** `Ti` → `Ne` → `Si` → `Fe`

| Posisi | Fungsi | Deskripsi Peran dalam INTP |
|--------|--------|---------------------------|
| Dominant | **Ti** (Introverted Thinking) | INTP adalah pembangun kerangka logika internal. Ti terus-menerus menganalisis konsistensi, mencari presisi, dan membangun model mental yang akurat. Berbeda dengan Te yang berorientasi hasil, Ti puas dengan *pemahaman* itu sendiri. |
| Auxiliary | **Ne** (Extraverted iNtuition) | Mesin asosiasi ide tanpa batas. Ne memberi INTP kegembiraan dalam menemukan koneksi antar domain yang tampak tidak berhubungan. "Bagaimana jika..." adalah pertanyaan favorit INTP. |
| Tertiary | **Si** (Introverted Sensing) | Memberikan referensi pada pengalaman dan pengetahuan yang terakumulasi. INTP bisa sangat terikat pada rutinitas atau cara berpikir lama (Si) sembari mengeksplorasi ide baru (Ne)—sebuah tegangan produktif. |
| Inferior | **Fe** (Extraverted Feeling) | Kebutuhan tersembunyi akan koneksi emosional dan persetujuan sosial. INTP sering tidak menyadari betapa mereka mendambakan kedekatan emosional. Saat stres, Fe bisa meledak: hipersensitif terhadap kritik atau tiba-tiba sangat ingin disukai. |

**Dinamika Psikologis:**
INTP hidup di "alam ide"—sebuah dunia internal yang lebih nyata dari dunia fisik. Mereka sering tampak **tidak hadir** secara sosial bukan karena tidak peduli, tapi karena Ti sedang sibuk memverifikasi konsistensi sistem internal. Pola umum: INTP sangat produktif dalam fase eksplorasi ide (Ti+Ne), namun sering kesulitan menyelesaikan proyek ketika fase implementasi membutuhkan disiplin rutin (yang melawan preferensi mereka).

```json
{
  "mbti_type": "INTP",
  "nickname": "The Logician",
  "nickname_id": "Sang Logikawan",
  "temperament": "NT (Rational)",
  "population_estimate": "3-5%",
  "cognitive_stack": ["Ti", "Ne", "Si", "Fe"],
  "core_drive": "Membangun model logika yang sempurna dan memahami bagaimana segalanya bekerja",
  "strengths": [
    "Kemampuan analisis logika yang mendalam dan presisi",
    "Kreativitas intelektual yang tinggi (koneksi lintas domain)",
    "Objektif dan tidak memihak dalam penilaian",
    "Keterbukaan terhadap revisi ide jika ada bukti baru",
    "Kemampuan berpikir abstrak yang langka"
  ],
  "challenges": [
    "Prokrastinasi akibat terlalu banyak kemungkinan (analysis paralysis)",
    "Kesulitan finalisasi dan eksekusi",
    "Bisa terlihat tidak peduli atau dingin padahal tidak",
    "Overthinking hingga ragu-ragu dalam keputusan",
    "Standar logika yang sangat tinggi membuat INTP sering mengkritik sendiri"
  ],
  "career_themes": ["Philosophy", "Mathematics", "Software Engineering", "Physics", "Writing", "Research"],
  "relationship_note": "Deeply loyal dan thoughtful, namun mengekspresikan kasih sayang melalui berbagi ide dan waktu intelektual, bukan sentuhan fisik atau kata-kata romantis."
}
```

---

### 8.3 ENTJ — The Commander (Sang Komandan)

**Stack Kognitif:** `Te` → `Ni` → `Se` → `Fi`

| Posisi | Fungsi | Deskripsi Peran dalam ENTJ |
|--------|--------|---------------------------|
| Dominant | **Te** (Extraverted Thinking) | ENTJ didorong oleh dorongan untuk mengorganisir dan mengoptimalkan dunia. Te yang dominan membuat mereka natural dalam kepemimpinan—mereka secara insting melihat apa yang perlu dilakukan, siapa yang harus melakukannya, dan bagaimana caranya paling efisien. |
| Auxiliary | **Ni** (Introverted iNtuition) | Memberikan visi strategis jangka panjang. ENTJ bukan sekadar eksekutor—mereka memiliki kemampuan untuk melihat ke depan dan memposisikan sumber daya dengan tepat. Ni memberi kedalaman pada kepemimpinan Te yang visioner. |
| Tertiary | **Se** (Extraverted Sensing) | Kepekaan terhadap realitas saat ini dan lingkungan fisik. ENTJ yang matang bisa "membaca ruangan" dan beradaptasi taktis. Se yang berkembang membuat ENTJ lebih karismatik dan hadir secara fisik. |
| Inferior | **Fi** (Introverted Feeling) | Titik kelemahan: ENTJ sering tidak sadar akan nilai-nilai dan emosi pribadi mereka, serta dampak keputusan mereka pada orang lain secara emosional. Saat stres: ENTJ bisa menjadi hypersensitif terhadap kritik personal atau tiba-tiba emosional secara tidak proporsional. |

```json
{
  "mbti_type": "ENTJ",
  "nickname": "The Commander",
  "nickname_id": "Sang Komandan",
  "temperament": "NT (Rational)",
  "population_estimate": "2-5%",
  "cognitive_stack": ["Te", "Ni", "Se", "Fi"],
  "core_drive": "Memimpin orang dan sistem menuju tujuan yang ambisius dan terstruktur",
  "strengths": [
    "Kepemimpinan natural yang karismatik dan visioner",
    "Kemampuan eksekusi dan pengambilan keputusan yang cepat",
    "Determinasi dan kepercayaan diri tinggi",
    "Kemampuan membangun sistem dan organisasi yang efisien",
    "Visi jangka panjang yang strategis"
  ],
  "challenges": [
    "Bisa tampak intimidating atau dominan",
    "Kurang sadar akan kebutuhan emosional orang lain",
    "Intoleran terhadap inefisiensi atau ketidakmampuan",
    "Workaholic—sulit berhenti dan menikmati proses",
    "Bisa terlalu kritis dan demanding terhadap tim"
  ],
  "career_themes": ["Executive Leadership", "Entrepreneurship", "Law", "Management Consulting", "Politics"],
  "relationship_note": "Berkomitmen dan melindungi, namun perlu belajar bahwa 'menyelesaikan masalah' bukan selalu yang dibutuhkan pasangan—kadang hanya didengarkan."
}
```

---

### 8.4 ENTP — The Debater (Sang Debater)

**Stack Kognitif:** `Ne` → `Ti` → `Fe` → `Si`

| Posisi | Fungsi | Deskripsi Peran dalam ENTP |
|--------|--------|---------------------------|
| Dominant | **Ne** (Extraverted iNtuition) | ENTP adalah *idea machine*. Ne yang dominan terus-menerus menghasilkan koneksi baru, kemungkinan baru, dan perspektif baru. ENTP tidak bisa berhenti melihat "sisi lain" dari setiap argumen—bukan karena tidak memiliki pendirian, tapi karena Ne ingin mengeksplorasi semua kemungkinan. |
| Auxiliary | **Ti** (Introverted Thinking) | Memberikan kerangka logis untuk mengevaluasi ide-ide yang dihasilkan Ne. Ti membuat ENTP menjadi debater yang formidable—mereka tidak hanya menghasilkan ide, tapi juga bisa menganalisis kekuatan dan kelemahannya secara mendalam. |
| Tertiary | **Fe** (Extraverted Feeling) | Kesadaran akan dinamika sosial dan emosi kelompok. ENTP yang matang menggunakan Fe untuk membaca audiens dan menyesuaikan argumen agar lebih persuasif. Fe juga memberi ENTP kemampuan menghibur dan karisma sosial. |
| Inferior | **Si** (Introverted Sensing) | Kelemahan terbesar ENTP: konsistensi, rutinitas, dan perhatian terhadap detail praktis. ENTP sering mengulang kesalahan yang sama (Si lemah berarti kurang belajar dari pengalaman masa lalu). Saat stres: obsesif terhadap detail kecil yang biasanya diabaikan. |

```json
{
  "mbti_type": "ENTP",
  "nickname": "The Debater",
  "nickname_id": "Sang Debater",
  "temperament": "NT (Rational)",
  "population_estimate": "3-5%",
  "cognitive_stack": ["Ne", "Ti", "Fe", "Si"],
  "core_drive": "Mengeksplorasi ide-ide inovatif dan menantang status quo melalui debat intelektual",
  "strengths": [
    "Kreativitas dan inovasi yang luar biasa",
    "Kemampuan berdebat dan berargumentasi dari berbagai sudut",
    "Adaptabilitas dan fleksibilitas intelektual tinggi",
    "Karisma dan kemampuan sosial yang natural",
    "Quick thinker dalam situasi tak terduga"
  ],
  "challenges": [
    "Kesulitan menindaklanjuti ide hingga selesai",
    "Bisa menjengkelkan karena selalu 'devil's advocate'",
    "Bosan dengan rutinitas dan detail operasional",
    "Prokrastinasi menjelang deadline",
    "Terkadang berdebat demi berdebat, bukan demi kebenaran"
  ],
  "career_themes": ["Entrepreneurship", "Innovation", "Law", "Technology", "Media", "Consulting"],
  "relationship_note": "Mengasyikkan dan stimulating sebagai pasangan, namun perlu komitmen ekstra untuk tidak mengabaikan kebutuhan emosional pasangan demi diskusi intelektual."
}
```

---

### 8.5 INFJ — The Advocate (Sang Advokat)

**Stack Kognitif:** `Ni` → `Fe` → `Ti` → `Se`

| Posisi | Fungsi | Deskripsi Peran dalam INFJ |
|--------|--------|---------------------------|
| Dominant | **Ni** (Introverted iNtuition) | Fungsi paling misterius. INFJ "tahu" hal-hal tanpa bisa menjelaskan bagaimana mereka tahu. Ni mensintesis sinyal-sinyal samar menjadi pemahaman holistik tentang pola, makna, dan arah masa depan. Inilah asal muasal reputasi INFJ sebagai "the mystic" atau "psychic". |
| Auxiliary | **Fe** (Extraverted Feeling) | Mengarahkan visi Ni untuk kepentingan orang banyak. Fe membuat INFJ benar-benar peduli pada kesejahteraan orang lain—bukan sekadar performa. INFJ *merasakan* emosi orang lain hampir seperti emosi sendiri (bukan empati biasa, tapi sesuatu mendekati kemampuan menyerap perasaan orang lain). |
| Tertiary | **Ti** (Introverted Thinking) | Analitik internal yang memberi INFJ kemampuan untuk menyusun argumen logis mendukung nilai-nilai Fe mereka. Ti yang berkembang membuat INFJ lebih kritis dan tidak hanya beroperasi pada domain perasaan. |
| Inferior | **Se** (Extraverted Sensing) | Kelemahan bersama INTJ. INFJ bisa terputus dari realitas sensoris—lupa makan, tidak sadar waktu, dll. Saat stres: overindulgence pada sensasi fisik (makanan, belanja, dll.) sebagai pelarian dari beban emosional. |

**Pola Khusus INFJ — "The Door Slam":**
INFJ memiliki pola pemutusan hubungan yang dikenal sebagai "door slam"—ketika seseorang terus-menerus melanggar nilai-nilai Fi tersembunyi atau menguras energi Fe INFJ, mereka akan menutup hubungan secara total dan permanen. Ini bukan drama—ini mekanisme pertahanan diri yang kuat.

```json
{
  "mbti_type": "INFJ",
  "nickname": "The Advocate",
  "nickname_id": "Sang Advokat",
  "temperament": "NF (Idealist)",
  "population_estimate": "1-3%",
  "cognitive_stack": ["Ni", "Fe", "Ti", "Se"],
  "core_drive": "Membantu orang berkembang menuju versi terbaik mereka, didorong oleh visi mendalam tentang potensi manusia",
  "strengths": [
    "Insight tentang orang dan situasi yang sangat dalam",
    "Empati yang mendalam dan autentik",
    "Kemampuan visi jangka panjang tentang pertumbuhan personal",
    "Kreativitas dalam mengkomunikasikan ide kompleks",
    "Determinasi dalam memperjuangkan nilai-nilai yang dipercaya"
  ],
  "challenges": [
    "Kelelahan emosional karena menyerap perasaan orang lain",
    "Perfeksionisme terhadap visi ideal yang tidak realistis",
    "Kesulitan menetapkan batasan sehat",
    "Rentan terhadap burnout karena terlalu banyak memberi",
    "Paradoks: sangat memahami orang namun sering merasa disalahpahami"
  ],
  "career_themes": ["Counseling", "Writing", "Teaching", "Activism", "Healthcare", "Psychology"],
  "relationship_note": "Pasangan yang sangat berkomitmen dan intuitif, namun membutuhkan ruang solitude dan pasangan yang bisa hadir secara emosional dengan kedalaman setara."
}
```

---

### 8.6 INFP — The Mediator (Sang Mediator)

**Stack Kognitif:** `Fi` → `Ne` → `Si` → `Te`

| Posisi | Fungsi | Deskripsi Peran dalam INFP |
|--------|--------|---------------------------|
| Dominant | **Fi** (Introverted Feeling) | Kompas moral internal yang sangat kuat dan kaya. Fi adalah tentang *nilai yang autentik*, bukan aturan eksternal. INFP tahu secara intuitif apakah sesuatu "benar" atau "salah" berdasarkan sistem nilai internal yang sangat kompleks dan personal. |
| Auxiliary | **Ne** (Extraverted iNtuition) | Memungkinkan INFP mengekspresikan nilai Fi-nya melalui kreativitas, imajinasi, dan koneksi ide yang unik. INFP adalah pemimpi kreatif yang melihat potensi di mana-mana. |
| Tertiary | **Si** (Introverted Sensing) | Koneksi dengan kenangan personal dan tradisi bermakna. INFP sering sangat nostalgis dan melekat pada pengalaman masa lalu yang membentuk identitas. |
| Inferior | **Te** (Extraverted Thinking) | Kelemahan terbesar: eksekusi, manajemen waktu, dan efisiensi. INFP penuh ide dan nilai tapi sering kesulitan mentranslasikan menjadi aksi sistematis. Saat stres: tiba-tiba menjadi hiperkritis dan judgemental (Te yang tidak matang). |

```json
{
  "mbti_type": "INFP",
  "nickname": "The Mediator",
  "nickname_id": "Sang Mediator",
  "temperament": "NF (Idealist)",
  "population_estimate": "4-5%",
  "cognitive_stack": ["Fi", "Ne", "Si", "Te"],
  "core_drive": "Hidup sesuai nilai-nilai personal yang autentik dan membantu orang lain menemukan kebenaran mereka",
  "strengths": [
    "Kreativitas dan imajinasi yang kaya",
    "Empati mendalam dan genuinely caring",
    "Komitmen kuat terhadap nilai-nilai personal",
    "Kemampuan melihat potensi dalam orang lain",
    "Keaslian dan autentisitas yang langka"
  ],
  "challenges": [
    "Terlalu idealistis hingga sulit berdamai dengan realita",
    "Menghindari konflik hingga masalah menumpuk",
    "Prokrastinasi dan kesulitan eksekusi",
    "Hipersensitif terhadap kritik yang terasa sebagai serangan identitas",
    "Rentan terhadap isolasi ketika tidak diterima lingkungan"
  ],
  "career_themes": ["Writing", "Art", "Counseling", "Education", "Social Work", "UX Design"],
  "relationship_note": "Passionate dan loyal, namun butuh waktu lama untuk membuka diri. Ketika sudah trust, mereka adalah pasangan yang sangat dalam dan berkomitmen."
}
```

---

### 8.7 ENFJ — The Protagonist (Sang Protagonis)

**Stack Kognitif:** `Fe` → `Ni` → `Se` → `Ti`

| Posisi | Fungsi | Deskripsi Peran dalam ENFJ |
|--------|--------|---------------------------|
| Dominant | **Fe** (Extraverted Feeling) | ENFJ hidup untuk menciptakan harmoni dan memfasilitasi pertumbuhan orang lain. Fe yang dominan berarti mereka secara konstan memindai kebutuhan emosional orang di sekitar dan bereaksi dengan penuh kepedulian. |
| Auxiliary | **Ni** (Introverted iNtuition) | Memberi ENFJ kemampuan untuk melihat potensi masa depan dalam diri orang lain. ENFJ bukan sekadar supportif—mereka memiliki visi tentang *kemana seseorang bisa berkembang* dan bekerja aktif untuk memfasilitasinya. |
| Tertiary | **Se** (Extraverted Sensing) | Kepekaan terhadap lingkungan fisik dan presentasi diri. ENFJ sering karismatik secara fisik, ekspresif, dan hadir secara penuh dalam interaksi. |
| Inferior | **Ti** (Introverted Thinking) | Kelemahan: analisis logis impersonal. ENFJ bisa kesulitan membuat keputusan yang secara logis benar tapi menyakiti orang yang mereka cintai. Saat stres: hiperkritis terhadap diri sendiri atau menjadi overly analytical dengan cara yang tidak produktif. |

```json
{
  "mbti_type": "ENFJ",
  "nickname": "The Protagonist",
  "nickname_id": "Sang Protagonis",
  "temperament": "NF (Idealist)",
  "population_estimate": "2-3%",
  "cognitive_stack": ["Fe", "Ni", "Se", "Ti"],
  "core_drive": "Memimpin dan menginspirasi orang lain untuk mencapai potensi penuh mereka",
  "strengths": [
    "Kepemimpinan yang hangat dan menginspirasi",
    "Kemampuan memahami dan memotivasi orang lain",
    "Karisma dan kemampuan komunikasi yang tinggi",
    "Visi tentang potensi orang lain yang sangat kuat",
    "Kemampuan membangun harmoni dalam kelompok"
  ],
  "challenges": [
    "Terlalu mengutamakan kebutuhan orang lain di atas diri sendiri",
    "Kesulitan menerima kritik personal",
    "Bisa manipulatif tanpa sadar demi menciptakan harmoni",
    "Burnout karena terlalu banyak memberi",
    "Kesulitan membuat keputusan yang menyakiti orang lain meski logis"
  ],
  "career_themes": ["Teaching", "Coaching", "Counseling", "Human Resources", "Public Relations", "Politics"],
  "relationship_note": "Pasangan yang sangat devoted dan supportif, namun perlu diingatkan bahwa kebutuhan mereka sendiri sama pentingnya."
}
```

---

### 8.8 ENFP — The Campaigner (Sang Kampanyer)

**Stack Kognitif:** `Ne` → `Fi` → `Te` → `Si`

| Posisi | Fungsi | Deskripsi Peran dalam ENFP |
|--------|--------|---------------------------|
| Dominant | **Ne** (Extraverted iNtuition) | Mesin kemungkinan ENFP. Ne terus-menerus menghasilkan koneksi baru, kemungkinan baru, dan cara pandang baru. ENFP adalah tentang potensi—potensi ide, potensi orang, potensi masa depan. |
| Auxiliary | **Fi** (Introverted Feeling) | Memberi arah dan makna bagi eksplorasi Ne. Bukan semua kemungkinan layak dikejar—Fi mengevaluasi mana yang *benar* secara personal. ENFP bukan sekadar petualang ide; mereka petualang yang berpegang pada nilai. |
| Tertiary | **Te** (Extraverted Thinking) | Kemampuan untuk mengorganisir dan mengeksekusi ketika termotivasi. ENFP bisa sangat produktif dan efisien—tapi hanya ketika proyeknya selaras dengan Fi mereka. |
| Inferior | **Si** (Introverted Sensing) | Kelemahan: konsistensi, perhatian detail rutin, dan follow-through. ENFP memiliki reputasi "banyak mulai, sedikit selesai"—ini adalah manifestasi Si yang lemah. |

```json
{
  "mbti_type": "ENFP",
  "nickname": "The Campaigner",
  "nickname_id": "Sang Kampanyer",
  "temperament": "NF (Idealist)",
  "population_estimate": "7-8%",
  "cognitive_stack": ["Ne", "Fi", "Te", "Si"],
  "core_drive": "Mengeksplorasi dan mewujudkan kemungkinan-kemungkinan bermakna dalam diri, orang lain, dan dunia",
  "strengths": [
    "Antusiasme dan energi yang menular",
    "Kreativitas dan kemampuan koneksi ide yang luar biasa",
    "Empati genuine dan perhatian pada individu",
    "Kemampuan melihat potensi dalam orang dan situasi",
    "Adaptabilitas dan keterbukaan terhadap pengalaman baru"
  ],
  "challenges": [
    "Kesulitan menyelesaikan proyek jangka panjang",
    "Mudah terdistraksi oleh peluang baru",
    "Overthinking dan kecemasan tersembunyi",
    "Hipersensitif terhadap penolakan",
    "Bisa impulsif dalam keputusan emosional"
  ],
  "career_themes": ["Marketing", "Journalism", "Counseling", "Teaching", "Entertainment", "Social Work"],
  "relationship_note": "Penuh cinta dan perhatian, sangat enthusiastic tentang pasangan. Namun perlu komitmen aktif untuk tidak teralihkan dan tetap present dalam hubungan jangka panjang."
}
```

---

### 8.9 ISTJ — The Logistician (Sang Logistisian)

**Stack Kognitif:** `Si` → `Te` → `Fi` → `Ne`

| Posisi | Fungsi | Deskripsi Peran dalam ISTJ |
|--------|--------|---------------------------|
| Dominant | **Si** (Introverted Sensing) | ISTJ adalah *database hidup* dari pengalaman dan prosedur yang telah terbukti. Si mendorong kepercayaan pada apa yang telah berhasil di masa lalu. Setiap situasi baru dibandingkan dengan database internal ini untuk memastikan konsistensi dan keandalan. |
| Auxiliary | **Te** (Extraverted Thinking) | Menerapkan pengetahuan Si secara efisien dan sistematis. ISTJ adalah master implementasi: mereka mengambil apa yang "terbukti benar" (Si) dan mengeksekusinya dengan presisi dan efisiensi (Te). |
| Tertiary | **Fi** (Introverted Feeling) | Nilai-nilai personal yang kuat namun jarang diekspresikan secara verbal. ISTJ sangat bermoral—mereka hanya tidak membicarakannya, mereka hidup sesuainya. |
| Inferior | **Ne** (Extraverted iNtuition) | Kelemahan: imajinasi spekulatif, toleransi terhadap ambiguitas, dan kemampuan melihat kemungkinan luar biasa (bukan berbasis pengalaman). Saat stres: ISTJ bisa menjadi paranoid tentang kemungkinan buruk yang tidak realistis. |

```json
{
  "mbti_type": "ISTJ",
  "nickname": "The Logistician",
  "nickname_id": "Sang Logistisian",
  "temperament": "SJ (Guardian)",
  "population_estimate": "11-14%",
  "cognitive_stack": ["Si", "Te", "Fi", "Ne"],
  "core_drive": "Memastikan stabilitas, keandalan, dan penyelesaian tanggung jawab melalui metode yang terbukti",
  "strengths": [
    "Keandalan dan konsistensi yang luar biasa",
    "Perhatian detail dan akurasi yang tinggi",
    "Tanggung jawab dan integritas yang kuat",
    "Kemampuan implementasi sistematis yang excellent",
    "Ingatan faktual dan prosedural yang kuat"
  ],
  "challenges": [
    "Resistensi terhadap perubahan dan cara baru",
    "Kaku dalam situasi yang membutuhkan fleksibilitas",
    "Kesulitan mengekspresikan emosi dan kebutuhan personal",
    "Bisa terlihat cold atau kritik berlebihan",
    "Overcommit terhadap aturan meski situasi membutuhkan pengecualian"
  ],
  "career_themes": ["Accounting", "Law Enforcement", "Military", "Administration", "Engineering", "Healthcare (nursing)"],
  "relationship_note": "Extremely loyal dan reliable. Mengekspresikan kasih sayang melalui tindakan nyata (acts of service). Perlu pasangan yang menghargai konsistensi dan tidak membutuhkan ekspresi verbal berlebihan."
}
```

---

### 8.10 ISFJ — The Defender (Sang Pembela)

**Stack Kognitif:** `Si` → `Fe` → `Ti` → `Ne`

| Posisi | Fungsi | Deskripsi Peran dalam ISFJ |
|--------|--------|---------------------------|
| Dominant | **Si** (Introverted Sensing) | Seperti ISTJ, ISFJ menyimpan database pengalaman yang kaya—namun Si ISFJ lebih berwarna emosional. Mereka mengingat *bagaimana perasaan orang* dalam situasi tertentu, bukan hanya fakta prosedural. |
| Auxiliary | **Fe** (Extraverted Feeling) | Mendorong Si ISFJ untuk digunakan dalam pelayanan orang lain. ISFJ adalah caretaker par excellence: mengingat preferensi orang (Si) dan bertindak untuk membuat mereka bahagia (Fe). |
| Tertiary | **Ti** (Introverted Thinking) | Analisis internal yang memberi ISFJ kemampuan untuk mengevaluasi informasi secara kritis. Kurang dominan, namun ISFJ yang matang menggunakannya untuk membuat batasan sehat. |
| Inferior | **Ne** (Extraverted iNtuition) | Kelemahan dan sumber kecemasan: ISFJ sering mengkhawatirkan kemungkinan terburuk yang belum tentu terjadi. Saat stres: spiraling ke skenario negatif yang semakin tidak realistis. |

```json
{
  "mbti_type": "ISFJ",
  "nickname": "The Defender",
  "nickname_id": "Sang Pembela",
  "temperament": "SJ (Guardian)",
  "population_estimate": "9-14%",
  "cognitive_stack": ["Si", "Fe", "Ti", "Ne"],
  "core_drive": "Merawat dan melindungi orang-orang yang dicintai dengan cara yang konkret dan penuh perhatian",
  "strengths": [
    "Kepedulian dan perhatian terhadap detail kebutuhan orang lain",
    "Keandalan dan dedikasi yang sangat tinggi",
    "Ingatan yang kuat tentang preferensi dan kebutuhan orang",
    "Kesabaran dan keuletan dalam membantu",
    "Menciptakan lingkungan yang nyaman dan aman"
  ],
  "challenges": [
    "Kesulitan menetapkan batasan dan berkata 'tidak'",
    "Mengorbankan kebutuhan sendiri secara berlebihan",
    "Kecemasan berlebihan terhadap kemungkinan buruk",
    "Resistensi terhadap perubahan yang mengancam kenyamanan",
    "Memendam perasaan negatif sampai meledak"
  ],
  "career_themes": ["Nursing", "Teaching (elementary)", "Social Work", "Administration", "Childcare", "Religious service"],
  "relationship_note": "Pasangan yang sangat devoted dan perhatif. Sering memberikan lebih dari yang dibutuhkan—pasangan perlu aktif memastikan kebutuhan ISFJ juga terpenuhi."
}
```

---

### 8.11 ESTJ — The Executive (Sang Eksekutif)

**Stack Kognitif:** `Te` → `Si` → `Ne` → `Fi`

| Posisi | Fungsi | Deskripsi Peran dalam ESTJ |
|--------|--------|---------------------------|
| Dominant | **Te** (Extraverted Thinking) | ESTJ mengorganisir dunia. Te yang dominan membuat mereka natural dalam manajemen, pembuatan kebijakan, dan memastikan sistem berjalan efisien. Mereka percaya bahwa "ada cara yang benar untuk melakukan sesuatu"—dan mereka tahu caranya. |
| Auxiliary | **Si** (Introverted Sensing) | Memberikan ESTJ referensi tradisi, prosedur yang terbukti, dan preseden. "Jika ini berhasil di masa lalu, gunakan lagi." Si membuat ESTJ menjadi penjaga tradisi dan standar. |
| Tertiary | **Ne** (Extraverted iNtuition) | Kemampuan brainstorming dan melihat kemungkinan baru—lebih terbatas dibanding NT, namun ESTJ yang matang bisa cukup inovatif dalam problem-solving. |
| Inferior | **Fi** (Introverted Feeling) | Kelemahan: kesadaran akan nilai-nilai personal dan emosi orang lain. ESTJ sering tidak sadar bahwa pendekatan mereka yang "efisien" bisa terasa kasar atau tidak berperasaan kepada orang lain. |

```json
{
  "mbti_type": "ESTJ",
  "nickname": "The Executive",
  "nickname_id": "Sang Eksekutif",
  "temperament": "SJ (Guardian)",
  "population_estimate": "8-12%",
  "cognitive_stack": ["Te", "Si", "Ne", "Fi"],
  "core_drive": "Memimpin dan mengorganisir komunitas sesuai standar dan tradisi yang terbukti efektif",
  "strengths": [
    "Kepemimpinan yang tegas dan terorganisir",
    "Kemampuan implementasi prosedur yang excellent",
    "Keandalan dan konsistensi tinggi",
    "Kemampuan manajemen dan delegasi",
    "Dedikasi terhadap tanggung jawab dan komitmen"
  ],
  "challenges": [
    "Kaku dan resistensi terhadap cara-cara baru",
    "Kurang peka terhadap kebutuhan emosional orang lain",
    "Terlalu fokus pada aturan tanpa melihat konteks",
    "Bisa mendominasi dan mengontrol",
    "Kesulitan mengakui kesalahan atau mengubah pendapat"
  ],
  "career_themes": ["Management", "Military", "Law", "Business Operations", "Finance", "Government"],
  "relationship_note": "Loyal dan protektif. Mengekspresikan cinta melalui stabilitas finansial dan perlindungan fisik. Perlu bekerja aktif untuk menunjukkan empati secara verbal."
}
```

---

### 8.12 ESFJ — The Consul (Sang Konsul)

**Stack Kognitif:** `Fe` → `Si` → `Ne` → `Ti`

| Posisi | Fungsi | Deskripsi Peran dalam ESFJ |
|--------|--------|---------------------------|
| Dominant | **Fe** (Extraverted Feeling) | ESFJ adalah "social glue"—mereka memiliki kemampuan bawaan untuk merasakan dinamika emosional kelompok dan bertindak untuk menjaga harmoni. Fe dominan membuat mereka sangat peduli dengan persetujuan sosial dan kebutuhan orang lain. |
| Auxiliary | **Si** (Introverted Sensing) | Memberikan ESFJ orientasi pada tradisi, prosedur sosial yang telah terbukti, dan cara-cara "yang benar" dalam berinteraksi. ESFJ sangat menghargai tradisi, upacara, dan cara-cara lama yang bermakna. |
| Tertiary | **Ne** (Extraverted iNtuition) | Fleksibilitas dan kemampuan brainstorming untuk solusi kreatif dalam membantu orang lain. ESFJ yang matang bisa sangat inovatif dalam pelayanan mereka. |
| Inferior | **Ti** (Introverted Thinking) | Kelemahan: analisis logis yang impersonal dan independensi intelektual. Saat stres, ESFJ bisa menjadi hiperkritis dan menghakimi dengan cara yang tidak biasanya mereka lakukan. |

```json
{
  "mbti_type": "ESFJ",
  "nickname": "The Consul",
  "nickname_id": "Sang Konsul",
  "temperament": "SJ (Guardian)",
  "population_estimate": "9-13%",
  "cognitive_stack": ["Fe", "Si", "Ne", "Ti"],
  "core_drive": "Menciptakan harmoni sosial dan merawat kebutuhan komunitas melalui tradisi dan kepedulian aktif",
  "strengths": [
    "Kehangatan sosial dan kemampuan membangun koneksi",
    "Sangat perhatif terhadap kebutuhan orang lain",
    "Kemampuan organisasi sosial dan event",
    "Loyal dan supportif terhadap orang yang dicintai",
    "Kepekaan terhadap dinamika sosial"
  ],
  "challenges": [
    "Sangat sensitif terhadap kritik dan penolakan",
    "Terlalu mementingkan persetujuan sosial",
    "Kesulitan menangani konflik secara langsung",
    "Bisa terlalu controlling dalam "membantu"",
    "Kesulitan menerima pandangan yang berbeda dari norma"
  ],
  "career_themes": ["Healthcare", "Education", "Event Planning", "Human Resources", "Social Services", "Hospitality"],
  "relationship_note": "Pasangan yang sangat caring dan ekspresif. Membutuhkan validasi dan ekspresi kasih sayang yang konsisten. Akan sangat terluka jika usaha mereka diabaikan."
}
```

---

### 8.13 ISTP — The Virtuoso (Sang Virtuoso)

**Stack Kognitif:** `Ti` → `Se` → `Ni` → `Fe`

| Posisi | Fungsi | Deskripsi Peran dalam ISTP |
|--------|--------|---------------------------|
| Dominant | **Ti** (Introverted Thinking) | ISTP membangun pemahaman mendalam tentang *bagaimana sesuatu bekerja* secara mekanis dan logis. Berbeda dengan INTP (Ti+Ne), ISTP menerapkan logika pada dunia fisik nyata, bukan abstraksi. |
| Auxiliary | **Se** (Extraverted Sensing) | Memberikan ISTP koneksi yang kuat dengan realitas fisik. ISTP sangat adept secara fisik—mereka "merasakan" mesin, material, dan lingkungan fisik dengan cara yang hampir intuitif. Ini yang membuat mereka excellent dalam pekerjaan teknis dan fisik. |
| Tertiary | **Ni** (Introverted iNtuition) | Memberikan kemampuan untuk melihat pola tersembunyi dan mengantisipasi masalah. ISTP yang matang menggunakan Ni untuk perencanaan taktis. |
| Inferior | **Fe** (Extraverted Feeling) | Kelemahan bawah sadar: ekspresi emosi dan sensitivitas sosial. ISTP sering terlihat dingin atau tidak peduli meski sebenarnya peduli—mereka hanya tidak tahu bagaimana mengekspresikannya. Saat stres: oversharing emosi secara tidak proporsional (Fe yang tidak terkendali). |

```json
{
  "mbti_type": "ISTP",
  "nickname": "The Virtuoso",
  "nickname_id": "Sang Virtuoso",
  "temperament": "SP (Artisan)",
  "population_estimate": "4-6%",
  "cognitive_stack": ["Ti", "Se", "Ni", "Fe"],
  "core_drive": "Memahami secara mendalam bagaimana sesuatu bekerja dan menguasainya secara praktis",
  "strengths": [
    "Kemampuan teknis dan mekanis yang sangat tinggi",
    "Pemecahan masalah dalam krisis dengan tenang",
    "Observasi lingkungan fisik yang tajam",
    "Efisiensi dan pragmatisme tinggi",
    "Ketenangan di bawah tekanan"
  ],
  "challenges": [
    "Kesulitan mengekspresikan emosi dan kebutuhan",
    "Menghindari komitmen jangka panjang",
    "Bisa terlihat tidak peduli atau detached",
    "Impulsif dalam pencarian sensasi baru",
    "Kesulitan mengikuti aturan yang terasa tidak logis"
  ],
  "career_themes": ["Engineering", "Mechanics", "Surgery", "Forensics", "Martial Arts", "Programming"],
  "relationship_note": "Menunjukkan kasih sayang melalui bantuan praktis dan kehadiran fisik. Butuh pasangan yang menghargai kemandirian mereka dan tidak menuntut ekspresi verbal yang berlebihan."
}
```

---

### 8.14 ISFP — The Adventurer (Sang Petualang)

**Stack Kognitif:** `Fi` → `Se` → `Ni` → `Te`

| Posisi | Fungsi | Deskripsi Peran dalam ISFP |
|--------|--------|---------------------------|
| Dominant | **Fi** (Introverted Feeling) | ISFP memiliki kehidupan emosional internal yang sangat kaya dan dalam, namun sangat jarang ditampilkan ke luar. Mereka hidup sesuai nilai-nilai personal yang autentik. Berbeda dengan INFP (Fi+Ne), ISFP lebih *hadir* dalam momen kini daripada di dunia imajinasi. |
| Auxiliary | **Se** (Extraverted Sensing) | Menghubungkan Fi ke dunia fisik. ISFP mengekspresikan nilai-nilai batin mereka melalui pengalaman sensoris: seni, musik, alam, gerakan. Se membuat ISFP sangat peka terhadap keindahan dan detail estetika. |
| Tertiary | **Ni** (Introverted iNtuition) | Kemampuan untuk merasakan pola tersembunyi dan makna di balik pengalaman sensoris. ISFP yang matang bisa sangat insightful tentang makna yang lebih dalam dari hal-hal fisik. |
| Inferior | **Te** (Extraverted Thinking) | Kelemahan: efisiensi sistematis, manajemen, dan pendelegasian. Saat stres, ISFP bisa menjadi hiperkritis dan tiba-tiba menjadi sangat judgemental terhadap orang lain (Te yang tidak matang). |

```json
{
  "mbti_type": "ISFP",
  "nickname": "The Adventurer",
  "nickname_id": "Sang Petualang",
  "temperament": "SP (Artisan)",
  "population_estimate": "5-9%",
  "cognitive_stack": ["Fi", "Se", "Ni", "Te"],
  "core_drive": "Mengalami keindahan dunia secara autentik dan mengekspresikan identitas sejati melalui kreasi dan pengalaman",
  "strengths": [
    "Kepekaan estetika dan artistik yang tinggi",
    "Kehangatan genuine dan perhatian pada individu",
    "Fleksibilitas dan keterbukaan terhadap pengalaman",
    "Keaslian dan keengganan berpura-pura",
    "Kehadiran penuh dalam momen kini"
  ],
  "challenges": [
    "Kesulitan merencanakan jangka panjang",
    "Menghindari konflik hingga berlarut",
    "Mudah terluka meski tidak menampakkannya",
    "Kesulitan menegaskan diri dan kebutuhan",
    "Bisa impulsif dalam keputusan berbasis perasaan"
  ],
  "career_themes": ["Fine Arts", "Music", "Fashion Design", "Veterinary", "Culinary Arts", "Photography"],
  "relationship_note": "Penuh kasih dan penuh perhatian melalui tindakan kecil yang bermakna. Sangat loyal namun butuh ruang dan kebebasan yang dihormati."
}
```

---

### 8.15 ESTP — The Entrepreneur (Sang Entrepreneur)

**Stack Kognitif:** `Se` → `Ti` → `Fe` → `Ni`

| Posisi | Fungsi | Deskripsi Peran dalam ESTP |
|--------|--------|---------------------------|
| Dominant | **Se** (Extraverted Sensing) | ESTP adalah *master of the present moment*. Se yang dominan membuat mereka sangat responsif terhadap lingkungan fisik, peluang yang muncul saat ini, dan realitas yang bisa ditangkap secara langsung. |
| Auxiliary | **Ti** (Introverted Thinking) | Memberikan ESTP kemampuan analitis untuk memahami cara kerja sistem yang mereka amati melalui Se. ESTP tidak hanya bereaksi—mereka juga menganalisis secara cepat dan tajam. |
| Tertiary | **Fe** (Extraverted Feeling) | Kepekaan sosial yang memberikan ESTP kemampuan membaca orang dan situasi sosial. ESTP sering sangat karismatik dan charming. |
| Inferior | **Ni** (Introverted iNtuition) | Kelemahan: visi jangka panjang dan perencanaan. ESTP bisa sangat sukses jangka pendek namun gagal melihat konsekuensi jangka panjang dari keputusan impulsif. |

```json
{
  "mbti_type": "ESTP",
  "nickname": "The Entrepreneur",
  "nickname_id": "Sang Entrepreneur",
  "temperament": "SP (Artisan)",
  "population_estimate": "4-6%",
  "cognitive_stack": ["Se", "Ti", "Fe", "Ni"],
  "core_drive": "Bertindak langsung pada peluang nyata yang ada saat ini dengan kecepatan dan ketajaman maksimal",
  "strengths": [
    "Responsivitas dan adaptabilitas dalam situasi krisis",
    "Kemampuan membaca situasi secara real-time",
    "Karisma dan kemampuan persuasi alami",
    "Keberanian mengambil risiko yang terkalkulasi",
    "Kemampuan problem-solving praktis yang cepat"
  ],
  "challenges": [
    "Impulsivitas dan kurangnya perencanaan jangka panjang",
    "Bisa manipulatif tanpa sadar",
    "Mudah bosan dengan rutinitas",
    "Mengabaikan aturan jika terasa membatasi",
    "Kesulitan berkomitmen pada hubungan/proyek jangka panjang"
  ],
  "career_themes": ["Sales", "Entrepreneurship", "Emergency Services", "Sports", "Trading", "Event Management"],
  "relationship_note": "Spontan, menyenangkan, dan penuh energi sebagai pasangan. Namun komitmen jangka panjang membutuhkan usaha sadar. Butuh pasangan yang bisa mengikuti energi mereka."
}
```

---

### 8.16 ESFP — The Entertainer (Sang Entertainer)

**Stack Kognitif:** `Se` → `Fi` → `Te` → `Ni`

| Posisi | Fungsi | Deskripsi Peran dalam ESFP |
|--------|--------|---------------------------|
| Dominant | **Se** (Extraverted Sensing) | Seperti ESTP, ESFP hadir sepenuhnya dalam momen kini. Namun Se ESFP diwarnai oleh Fi—pengalaman sensoris mereka penuh dengan nuansa emosional dan estetika. |
| Auxiliary | **Fi** (Introverted Feeling) | Memberikan kedalaman dan autentisitas pada energi Se. ESFP sangat kuat dalam nilai personal mereka—mereka tahu siapa mereka dan tidak akan mengkhianatinya demi persetujuan orang lain. |
| Tertiary | **Te** (Extraverted Thinking) | Kemampuan organisasi dan efisiensi yang berkembang seiring usia. ESFP yang matang bisa sangat produktif ketika proyeknya selaras dengan nilai Fi mereka. |
| Inferior | **Ni** (Introverted iNtuition) | Kelemahan: konsekuensi jangka panjang dan visi ke depan. ESFP bisa terjebak dalam siklus kesenangan sesaat tanpa mempertimbangkan dampak masa depan. |

```json
{
  "mbti_type": "ESFP",
  "nickname": "The Entertainer",
  "nickname_id": "Sang Entertainer",
  "temperament": "SP (Artisan)",
  "population_estimate": "7-10%",
  "cognitive_stack": ["Se", "Fi", "Te", "Ni"],
  "core_drive": "Menciptakan kegembiraan dan koneksi autentik melalui pengalaman yang hidup dan penuh warna",
  "strengths": [
    "Energi positif yang menular",
    "Kemampuan membuat orang lain merasa diterima dan dihibur",
    "Spontanitas dan kreativitas dalam momen",
    "Kepekaan estetika dan kemampuan perform",
    "Genuine caring terhadap orang-orang di sekitar"
  ],
  "challenges": [
    "Menghindari konflik dan topik serius",
    "Kesulitan merencanakan dan menabung untuk masa depan",
    "Mudah terdistraksi oleh peluang menyenangkan",
    "Hypersensitif terhadap kritik meski tampak tidak terpengaruh",
    "Bisa mengambil keputusan impulsif berbasis emosi"
  ],
  "career_themes": ["Performing Arts", "Event Planning", "Teaching (arts)", "Sales", "Tourism", "Social Media"],
  "relationship_note": "Pasangan yang penuh cinta, spontan, dan mengasyikkan. Butuh lingkungan relasi yang penuh penerimaan dan sedikit menuntut komitmen jadwal yang ketat."
}
```

---

## 9. Validasi & Edge Cases

### 9.1 Skenario Edge Case

| Skenario | Kondisi | Penanganan |
|----------|---------|-----------|
| **Near-zero score** | `|pci| < 5` pada satu atau lebih dikotomi | Tandai sebagai "ambiguous" untuk dikotomi tersebut; saran retest |
| **Multiple ambiguous** | 2+ dikotomi dengan pci < 5 | Tandai hasil sebagai "low confidence"; saran konsultasi profesional |
| **Speed responses** | `avg_response_ms < 1500` | Flag sebagai potentially invalid; warning kepada user |
| **Inconsistency tinggi** | `inconsistency_score > 50` | Hasil tidak valid; minta retest |
| **Incomplete session** | `completion_rate < 80%` | Jangan hitung hasil; minta penyelesaian |
| **Straight-line responding** | Semua jawaban nilai sama | Deteksi dan invalidasi |

### 9.2 Deteksi Straight-Line Responding

```typescript
function detectStraightLine(responses: QuestionResponse[]): boolean {
  // Cek apakah semua jawaban identik atau variance sangat rendah
  const values = responses
    .filter(r => typeof r.answer_value === 'number')
    .map(r => r.answer_value as number);
  
  if (values.length === 0) return false;
  
  const uniqueValues = new Set(values);
  const maxVariance = 0.5; // Threshold variance minimum
  
  if (uniqueValues.size <= 2) {
    // Semua jawaban hanya 1-2 nilai berbeda
    const mean = values.reduce((a, b) => a + b, 0) / values.length;
    const variance = values.reduce((a, b) => a + Math.pow(b - mean, 2), 0) / values.length;
    return variance <= maxVariance;
  }
  
  return false;
}
```

### 9.3 Skema Confidence Score

```typescript
interface ConfidenceScore {
  overall: number;           // 0-100
  per_dikotomi: {
    EI: number;
    SN: number;
    TF: number;
    JP: number;
  };
  flags: string[];           // Array peringatan
  recommendation: "reliable" | "review_suggested" | "retest_recommended";
}

function calculateConfidence(result: MBTIResult): ConfidenceScore {
  const flags: string[] = [];
  let confidence = 100;
  
  // Penalti per dikotomi jika ambiguous
  const dikotomi_scores = [
    result.scores.EI.pci,
    result.scores.SN.pci,
    result.scores.TF.pci,
    result.scores.JP.pci,
  ];
  
  for (const pci of dikotomi_scores) {
    if (pci < 5) { confidence -= 25; flags.push("near_zero_preference"); }
    else if (pci < 15) { confidence -= 10; flags.push("slight_preference"); }
  }
  
  // Penalti inkonsistensi
  if (result.reliability_indicators.inconsistency_score > 50) {
    confidence -= 30;
    flags.push("high_inconsistency");
  } else if (result.reliability_indicators.inconsistency_score > 30) {
    confidence -= 15;
    flags.push("moderate_inconsistency");
  }
  
  // Penalti completion rate
  if (result.reliability_indicators.completion_rate < 90) {
    confidence -= 20;
    flags.push("incomplete_responses");
  }
  
  // Penalti response time
  if (result.reliability_indicators.avg_response_ms < 1500) {
    confidence -= 20;
    flags.push("responses_too_fast");
  }
  
  confidence = Math.max(0, confidence);
  
  let recommendation: ConfidenceScore["recommendation"];
  if (confidence >= 70) recommendation = "reliable";
  else if (confidence >= 40) recommendation = "review_suggested";
  else recommendation = "retest_recommended";
  
  return {
    overall: confidence,
    per_dikotomi: {
      EI: result.scores.EI.pci,
      SN: result.scores.SN.pci,
      TF: result.scores.TF.pci,
      JP: result.scores.JP.pci,
    },
    flags,
    recommendation,
  };
}
```

---

## 10. API Response Schema

### 10.1 Response Body — Test Result

```json
{
  "status": "success",
  "data": {
    "session_id": "uuid-string",
    "result": {
      "mbti_type": "INTJ",
      "scores": {
        "EI": {
          "raw_score": -14.5,
          "pole_a_score": 8.0,
          "pole_b_score": 22.5,
          "max_possible": 45.0,
          "preference": "I",
          "pci": 32.2,
          "strength": "moderate"
        },
        "SN": {
          "raw_score": -18.0,
          "pole_a_score": 10.0,
          "pole_b_score": 28.0,
          "max_possible": 48.0,
          "preference": "N",
          "pci": 37.5,
          "strength": "moderate"
        },
        "TF": {
          "raw_score": 22.0,
          "pole_a_score": 29.0,
          "pole_b_score": 7.0,
          "max_possible": 48.0,
          "preference": "T",
          "pci": 45.8,
          "strength": "moderate"
        },
        "JP": {
          "raw_score": -10.0,
          "pole_a_score": 11.0,
          "pole_b_score": 21.0,
          "max_possible": 45.0,
          "preference": "J",
          "pci": 22.2,
          "strength": "slight"
        }
      },
      "cognitive_stack": {
        "dominant": "Ni",
        "auxiliary": "Te",
        "tertiary": "Fi",
        "inferior": "Se"
      },
      "confidence": {
        "overall": 75,
        "flags": ["slight_jp_preference"],
        "recommendation": "reliable"
      },
      "reliability": {
        "completion_rate": 100.0,
        "avg_response_ms": 4800,
        "inconsistency_score": 8.5,
        "is_reliable": true
      }
    },
    "profile": {
      "mbti_type": "INTJ",
      "nickname": "The Architect",
      "nickname_id": "Sang Arsitek",
      "temperament": "NT",
      "description_short": "Pemikir strategis dengan visi jangka panjang yang kuat...",
      "strengths": ["Strategic thinking", "Determinasi", "Standar tinggi"],
      "challenges": ["Kesulitan komunikasi emosional", "Impatience"],
      "cognitive_stack_description": {
        "dominant": {
          "function": "Ni",
          "name": "Introverted iNtuition",
          "description": "Mesin visi jangka panjang yang mensintesis pola kompleks..."
        },
        "auxiliary": {
          "function": "Te",
          "name": "Extraverted Thinking",
          "description": "Eksekutor sistematis yang mengorganisir dunia eksternal..."
        }
      }
    }
  },
  "meta": {
    "calculated_at": "2025-01-01T12:00:00Z",
    "algorithm_version": "2.0",
    "total_questions": 60,
    "questions_answered": 60
  }
}
```

### 10.2 HTTP Status Codes

| Status | Kondisi |
|--------|---------|
| `200 OK` | Hasil berhasil dihitung |
| `202 Accepted` | Sesi diterima, belum selesai |
| `400 Bad Request` | Request tidak valid (format jawaban salah) |
| `422 Unprocessable Entity` | Respons tidak cukup untuk kalkulasi (`completion_rate < 80%`) |
| `429 Too Many Requests` | Rate limiting |
| `500 Internal Server Error` | Kesalahan server |

### 10.3 Error Response Schema

```json
{
  "status": "error",
  "error": {
    "code": "INSUFFICIENT_RESPONSES",
    "message": "Minimal 80% pertanyaan harus dijawab untuk kalkulasi hasil.",
    "details": {
      "questions_required": 48,
      "questions_answered": 38,
      "completion_rate": 63.3
    }
  }
}
```

---

## Appendix A — Tabel Referensi Cepat: 16 Tipe

| Tipe | Stack | Temperamen | Populasi | Nickname |
|------|-------|------------|----------|---------|
| INTJ | Ni-Te-Fi-Se | NT | 2-4% | The Architect |
| INTP | Ti-Ne-Si-Fe | NT | 3-5% | The Logician |
| ENTJ | Te-Ni-Se-Fi | NT | 2-5% | The Commander |
| ENTP | Ne-Ti-Fe-Si | NT | 3-5% | The Debater |
| INFJ | Ni-Fe-Ti-Se | NF | 1-3% | The Advocate |
| INFP | Fi-Ne-Si-Te | NF | 4-5% | The Mediator |
| ENFJ | Fe-Ni-Se-Ti | NF | 2-3% | The Protagonist |
| ENFP | Ne-Fi-Te-Si | NF | 7-8% | The Campaigner |
| ISTJ | Si-Te-Fi-Ne | SJ | 11-14% | The Logistician |
| ISFJ | Si-Fe-Ti-Ne | SJ | 9-14% | The Defender |
| ESTJ | Te-Si-Ne-Fi | SJ | 8-12% | The Executive |
| ESFJ | Fe-Si-Ne-Ti | SJ | 9-13% | The Consul |
| ISTP | Ti-Se-Ni-Fe | SP | 4-6% | The Virtuoso |
| ISFP | Fi-Se-Ni-Te | SP | 5-9% | The Adventurer |
| ESTP | Se-Ti-Fe-Ni | SP | 4-6% | The Entrepreneur |
| ESFP | Se-Fi-Te-Ni | SP | 7-10% | The Entertainer |

---

## Appendix B — Tabel Kompatibilitas Fungsi Antar Tipe

Tipe-tipe dengan **fungsi dominan & auxiliary yang sama namun urutan berbeda** sering memiliki dinamika relasi yang kuat (saling melengkapi):

| Pasangan Komplementer | Alasan |
|-----------------------|--------|
| INTJ ↔ ENTP | Ni+Te ↔ Ne+Ti (sharing sumbu N dan T) |
| INFJ ↔ ENFP | Ni+Fe ↔ Ne+Fi (sharing sumbu N dan F) |
| ISTJ ↔ ESTP | Si+Te ↔ Se+Ti (sharing sumbu S dan T) |
| ISFJ ↔ ESFP | Si+Fe ↔ Se+Fi (sharing sumbu S dan F) |

---

## Appendix C — Checklist Implementasi untuk Developer

```
FASE 1: FOUNDATION
□ Setup database dengan skema di Bagian 6
□ Seed tabel mbti_type_profiles dengan 16 profil
□ Implementasi API endpoint: POST /sessions (buat sesi)
□ Implementasi API endpoint: POST /sessions/{id}/responses (kirim jawaban)
□ Implementasi API endpoint: GET /sessions/{id}/result (ambil hasil)

FASE 2: CORE ALGORITHM  
□ Implementasi fungsi calculateMBTI() (Bagian 5.2)
□ Implementasi fungsi deriveCognitiveStack() (Bagian 3.3)
□ Implementasi PCI calculation per dikotomi
□ Implementasi strength classification (slight/moderate/clear/very_clear)

FASE 3: QUESTION BANK
□ Buat minimal 60 soal dengan distribusi sesuai Bagian 4.3
□ Tag setiap soal dengan metadata lengkap (Bagian 4.4)
□ Setup anchor pairs untuk inconsistency detection (min. 10 pasang)
□ Implementasi question randomization dan counterbalancing

FASE 4: VALIDATION LAYER
□ Implementasi deteksi straight-line responding
□ Implementasi response time tracking
□ Implementasi inconsistency score calculation
□ Implementasi confidence score & recommendation

FASE 5: OPTIMASI
□ Index database untuk query performa
□ Caching untuk type profiles (data statis)
□ Rate limiting pada API endpoints
□ Logging & monitoring untuk data analisis
```

---

*Dokumen ini mengikuti teori Jungian original dan format instrumen MBTI standar psikometri. Untuk implementasi klinis, pastikan validasi tambahan oleh psikolog bersertifikat.*

**Versi Dokumen:** 2.0  
**Terakhir diperbarui:** 2025  
**Lisensi penggunaan:** Internal Development Use
