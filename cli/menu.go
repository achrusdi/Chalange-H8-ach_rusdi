package cli

import (
	"bufio"
	"fmt"
	"os"
)

func MainMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("==================== Selamat datang ====================")
	fmt.Println("Tekan (1) untuk login")
	fmt.Println("Tekan (2) untuk register")
	fmt.Println("Tekan (q) untuk keluar aplikasi")

	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "1":
		Login()
	case "2":
		Register()
	case "q":
		fmt.Println("Terimakasih telah mengguankan applikasi")
		os.Exit(1)
	default:
		MainMenu()
	}
}
