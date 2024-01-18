package parser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var result string

func MarshalCompose(filepath *string, compose DockerCompose) {

	// Write modified docker-compose file to stdout
	output, err := yaml.Marshal(&compose)
	if err != nil {
		log.Fatal(err)
	}
	result = string(output)
}

func WriteToFile(filepath *string) {
	err := os.WriteFile(*filepath, []byte(result), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func Print() {
	fmt.Println("======================================")
	fmt.Println(result)
}
