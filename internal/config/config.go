package config

import "flag"

type Conf struct {
	ENV         *string
	FilePath    *string
	ServiceName *string
	Image       *string
	Port        *string
}

func InitConfig() Conf {

	return getFlags()
}

func getFlags() Conf {

	// Define command line flags
	env := flag.String("env", "dev", "")
	filePath := flag.String("file", "compose.yaml", "Path to docker-compose file")
	serviceName := flag.String("service", "", "Name of the service to modify")
	image := flag.String("image", "", "New image for the service")
	port := flag.String("port", "", "New port for the service")

	tmp := Conf{ENV: env, FilePath: filePath, ServiceName: serviceName, Image: image, Port: port}

	flag.Parse()

	return tmp
}
