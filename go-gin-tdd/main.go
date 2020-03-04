package main

import "github.com/alochym01/ftp-api/api"

func main() {
	router := api.SetupRouter()
	router.Run()
}
