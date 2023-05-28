package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"os"
	"time"

	"gorm.io/gorm"
)

func Login() {
	// helpers.CLearScreen()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("==================== Silahkan Login ====================")
	fmt.Print("Masukkan username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan password: ")
	scanner.Scan()
	password := scanner.Text()

	var user entity.User

	err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Username atau Password Anda Salah")
			fmt.Println("Silahkan Login Kembali")
		} else {
			fmt.Println("Terjadi kesalahan:", err)
		}
		Login()
	}

	fmt.Println("Selamat Datang", user.Username, "!")
	fmt.Println("========================================================")
}

func Register() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=================== Silahkan Rerister ==================")
	fmt.Print("Masukkan username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Masukkan password: ")
	scanner.Scan()
	password := scanner.Text()

	user := entity.User{
		Username:   username,
		Password:   password,
		Created_at: time.Now(),
	}

	err := config.DB.Create(&user).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	Login()
}
