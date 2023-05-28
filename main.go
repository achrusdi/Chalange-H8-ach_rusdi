package main

import (
	"go-trial-class/cli"
	"go-trial-class/config"
	"go-trial-class/helpers"
)

func main() {
	helpers.CLearScreen()
	config.DBConnect()
	cli.MainMenu()
}
