package main

import "plDockerKubernetes/server"

func main() {
	server := server.NewServer(8080)
	server.Start()
}
