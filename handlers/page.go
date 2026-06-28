package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowHome menampilkan halaman utama
func ShowHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowQuiz menampilkan halaman kuesioner
func ShowQuiz(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", nil)
}

// ShowTentang menampilkan halaman tentang kami
func ShowTentang(c *gin.Context) {
	c.HTML(http.StatusOK, "tentang.html", nil)
}

// Show404 menampilkan halaman 404
func Show404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"Message": "Halaman yang Anda cari tidak ditemukan.",
	})
}

// Show500 menampilkan halaman error server
func Show500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"Message": "Terjadi kesalahan pada server. Silakan coba lagi.",
	})
}
