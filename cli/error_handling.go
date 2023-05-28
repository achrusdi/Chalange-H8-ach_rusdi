package cli

import (
	"fmt"
)

func ErrorHandler(msg string) {
	fmt.Println("Terjadi kesalahan pada aplikasi")
	fmt.Println(msg)

	// var input string
	// fmt.Println("Tekan (m) untuk kembali ke menu utama")
	// fmt.Println("Tekan (q) untuk keluar aplikasi")

	// _, err := fmt.Scanln(&input)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// switch input {
	// case "m":
	// 	MainMenu()
	// case "q":
	// 	fmt.Println("Terima kasih telah menggunakan aplikasi ini")
	// 	os.Exit(1)
	// default:
	// 	ErrorHandler(msg)
	// }
}
