package models

import (
	"api-login/config"
	"errors"
)

type Akun struct {
	ID            int
	Username      string
	Email         string
	Password      string
	TanggalDibuat string
}

func GetAkunByUsername(username string) (*Akun, error) {
	akun := &Akun{}
	query := "SELECT id, username, email, password, tanggal_dibuat FROM akun WHERE username=$1"

	err := config.DB.QueryRow(query, username).Scan(&akun.ID, &akun.Username, &akun.Email, &akun.Password, &akun.TanggalDibuat)
	if err != nil {
		return nil, errors.New("Akun tidak ditemukan")
	}
	return akun, nil
}
