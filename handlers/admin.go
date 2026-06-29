package handlers

import (
	"net/http"

	"ego/services"

	"github.com/gin-gonic/gin"
)

// Admin credentials (hardcoded)
const (
	adminUsername = "admin"
	adminPassword = "admin360"
)

// ShowLogin menampilkan halaman login admin
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginProcess memproses login admin
func LoginProcess(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != adminUsername || password != adminPassword {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Error": "Username atau password salah",
		})
		return
	}

	// Set session cookie (simple token-based)
	c.SetCookie("admin_token", "authenticated", 3600*2, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

// ShowDashboard menampilkan dashboard admin dengan data user
func ShowDashboard(c *gin.Context) {
	// Cek autentikasi
	token, err := c.Cookie("admin_token")
	if err != nil || token != "authenticated" {
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	users, err := services.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"Message": "Gagal mengambil data pengguna.",
		})
		return
	}

	// Hitung statistik
	totalUser := len(users)
	sudahBayar := 0
	belumBayar := 0
	totalPendapatan := 0

	type UserRow struct {
		ID         string
		Nama       string
		Email      string
		SudahBayar bool
		MBTITipe   string
		SkorEI     int
		SkorSN     int
		SkorTF     int
		SkorJP     int
		Dibuat     string
	}

	var rows []UserRow
	for _, u := range users {
		isPaid := u.StatusPembayaran == "paid"
		if isPaid {
			sudahBayar++
			totalPendapatan += 14900
		} else {
			belumBayar++
		}
		rows = append(rows, UserRow{
			ID:         u.ID,
			Nama:       u.Nama,
			Email:      u.Email,
			SudahBayar: isPaid,
			MBTITipe:   u.MBTITipe,
			SkorEI:     u.SkorEI,
			SkorSN:     u.SkorSN,
			SkorTF:     u.SkorTF,
			SkorJP:     u.SkorJP,
			Dibuat:     "-",
		})
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Users":           rows,
		"TotalUser":       totalUser,
		"SudahBayar":      sudahBayar,
		"BelumBayar":      belumBayar,
		"TotalPendapatan": totalPendapatan,
	})
}

// ShowUserDetail menampilkan detail user untuk admin
func ShowUserDetail(c *gin.Context) {
	// Cek autentikasi
	token, err := c.Cookie("admin_token")
	if err != nil || token != "authenticated" {
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"Message": "User tidak ditemukan.",
		})
		return
	}

	c.HTML(http.StatusOK, "user_detail.html", gin.H{
		"User": user,
	})
}

// LogoutProcess menghapus session admin
func LogoutProcess(c *gin.Context) {
	c.SetCookie("admin_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/login")
}
