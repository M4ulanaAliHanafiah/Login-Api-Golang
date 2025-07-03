package controllers

import (
	"api-login/config"
	"api-login/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Input JSON error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Input username:", input.Username)

	db := config.DB
	var akun models.Akun
	row := db.QueryRow("SELECT username, password FROM akun WHERE username = $1", input.Username)
	err := row.Scan(&akun.Username, &akun.Password)
	if err == sql.ErrNoRows {
		fmt.Println("User tidak ditemukan")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username tidak ditemukan"})
		return
	} else if err != nil {
		fmt.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if akun.Password != input.Password {
		fmt.Println("Password salah")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	fmt.Println("Login berhasil!")
	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil"})
}
