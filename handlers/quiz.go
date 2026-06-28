package handlers

import (
	"net/http"
	"strconv"

	"ego/services"

	"github.com/gin-gonic/gin"
)

// SubmitTest memproses jawaban kuis dan mengembalikan ID user sebagai JSON
func SubmitTest(c *gin.Context) {
	email := c.PostForm("email")
	nama := c.PostForm("nama")

	// Baca 15 jawaban dari form
	answers := make([]int, 15)
	for i := 1; i <= 15; i++ {
		val, err := strconv.Atoi(c.PostForm("q" + strconv.Itoa(i)))
		if err != nil {
			val = 0 // default jika tidak terisi
		}
		answers[i-1] = val
	}

	userID, err := services.ProcessQuizAnswers(email, nama, answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan data tes: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": userID,
	})
}

// ShowPaywall menampilkan halaman pembayaran
func ShowPaywall(c *gin.Context) {
	id := c.Param("id")
	data, err := services.GetPaywallData(id)
	if err != nil {
		c.String(http.StatusNotFound, "User tidak ditemukan")
		return
	}

	c.HTML(http.StatusOK, "paywall.html", gin.H{
		"ID":   data.ID,
		"Nama": data.Nama,
	})
}

// ShowResult menampilkan hasil kuis (hanya jika sudah bayar)
func ShowResult(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetQuizResult(id)
	if err != nil {
		c.String(http.StatusNotFound, "Data tidak ditemukan")
		return
	}
	if result == nil {
		// Belum bayar, redirect ke paywall
		c.Redirect(http.StatusSeeOther, "/paywall/"+id+"?error=belum_bayar")
		return
	}

	c.HTML(http.StatusOK, "hasil.html", gin.H{
		"Nama":                result.Nama,
		"Narsisme":            result.Narsisme,
		"Machiavellian":       result.Machiavellian,
		"Psikopati":           result.Psikopati,
		"ExecutiveSummary":    result.ExecutiveSummary,
		"RelationshipProfile": result.RelationshipProfile,
		"Kekuatan":            result.Kekuatan,
		"AreaPerhatian":       result.AreaPerhatian,
		"RelationshipInsight": result.RelationshipInsight,
		"CompatibilityNotes":  result.CompatibilityNotes,
		"ReflectionQuestions": result.ReflectionQuestions,
	})
}

// KonfirmasiBayar memproses konfirmasi pembayaran dari user
func KonfirmasiBayar(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		NamaPengirim string `json:"nama_pengirim"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak valid",
		})
		return
	}

	err := services.ConfirmPayment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal memproses pembayaran: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      id,
	})
}
