package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {
	//load config
	helpers.SetupConfig()
	//load logger
	helpers.SetupLogger()
	//load database
	helpers.SetupMySQL()

	//run grpc
	go cmd.ServerGRPC()
	
	//run http
	cmd.ServerHTTP()
}