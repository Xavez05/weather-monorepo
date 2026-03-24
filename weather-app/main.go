package main

import "log"

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("error de configuración: %v", err)
	}

	if err := StartServer(cfg); err != nil {
		log.Fatalf("error en el servidor: %v", err)
	}
}
