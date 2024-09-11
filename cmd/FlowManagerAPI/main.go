package main

import (
	"log"

	"github.com/rickalon/FlowManagerAPI/internal/config"
	"github.com/rickalon/FlowManagerAPI/internal/repositories"
	"github.com/rickalon/FlowManagerAPI/internal/router"
)

func main() {
	log.Println("[1] Setting enviroment variables")
	config := config.NewConfig(".env")
	config.SetConfigFile()
	strPqConfig := config.GetPostgresConfig()
	log.Println("[2] Setting DB Pool of connections")
	pqDB := repositories.NewPqDriver(strPqConfig)
	pqDB.SetUpDatabases()
	log.Println("[3] Configuring the router")
	router := router.NewRouter(":8080", pqDB) //localhost:8080
	log.Println("[4] Configuring the handlers and services[5]")
	router.CreateHandlersForSubrouter()
	log.Println("[6] ON...")
	router.StartListenAndServe()
}
