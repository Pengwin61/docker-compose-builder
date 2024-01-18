package parser

import (
	"docker-compose-builder/internal/config"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DockerCompose struct {
	Version  string
	Services map[string]Service
}

type Service struct {
	Image          string
	Ports          []string
	Container_Name string
	Restart        string
	Command        string
	Environment    []string
	Volumes        []string
	Expose         []string
	Depends_ON     []string
	Extra_Hosts    []string
	// HealtCheck     map[string]HealtCheck
}

// type HealtCheck struct {
// 	Test         []string
// 	Interval     string
// 	Timeout      string
// 	Retries      int
// 	Start_period string
// }

var compose DockerCompose

func ComposeParser(conf config.Conf) DockerCompose {
	// Read docker-compose file
	data, err := os.ReadFile(*conf.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse docker-compose file
	err = yaml.Unmarshal(data, &compose)
	if err != nil {
		log.Fatal(err)
	}

	return compose
}
