package main

import (
	"docker-compose-builder/internal/config"
	"docker-compose-builder/internal/parser"
	"fmt"
	"strings"
)

func main() {
	conf := config.InitConfig()

	compose := parser.ComposeParser(conf)

	composer := parser.Edit(&compose)

	//
	parser.MarshalCompose(conf.FilePath, *composer)
	parser.Print()

	confirm := getInput()
	if confirm == "yes" || confirm == "y" {
		parser.WriteToFile(conf.FilePath)
	} else {
		fmt.Println("Операция отменена")
	}

}

func getInput() string {
	var input string
	fmt.Println("Вы ходите записать в файл?")
	fmt.Scan(&input)

	input = strings.ToLower(input)

	return input

}
