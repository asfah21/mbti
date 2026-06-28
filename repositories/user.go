package repositories

import (
	"ego/database"
	"ego/models"
)

// InsertUser menyimpan data user baru dan mengembalikan ID
func InsertUser(email, nama string, skorNarsisme, skorMachiavellian, skorPsikopati int) (string, error) {
	var userID string
	query := `INSERT INTO users_test (nama, email, skor_narsisme, skor_machiavellian, skor_psikopati) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := database.DB.QueryRow(query, nama, email, skorNarsisme, skorMachiavellian, skorPsikopati).Scan(&userID)
	return userID, err
}

// GetUserName mengambil nama user berdasarkan ID
func GetUserName(id string) (string, error) {
	var nama string
	err := database.DB.QueryRow("SELECT nama FROM users_test WHERE id = $1", id).Scan(&nama)
	return nama, err
}

// GetUserResult mengambil data hasil lengkap user (tanpa email)
func GetUserResult(id string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT nama, email, skor_narsisme, skor_machiavellian, skor_psikopati, status_pembayaran FROM users_test WHERE id = $1"
	err := database.DB.QueryRow(query, id).Scan(
		&user.Nama, &user.Email, &user.SkorNarsisme, &user.SkorMachiavellian, &user.SkorPsikopati, &user.StatusPembayaran,
	)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

// UpdatePaymentStatus memperbarui status pembayaran user menjadi PAID
func UpdatePaymentStatus(id string) error {
	query := "UPDATE users_test SET status_pembayaran = 'PAID' WHERE id = $1"
	_, err := database.DB.Exec(query, id)
	return err
}
