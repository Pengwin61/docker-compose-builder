package parser

func Edit(compose *DockerCompose) *DockerCompose {
	compose.Version = "3.9"
	for k, v := range compose.Services {
		if k == "web" {
			// v.Ports = nil
			v.Ports = append(v.Ports, "443:443")
		}
	}
	return compose
}
