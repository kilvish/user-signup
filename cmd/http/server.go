package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kilvish/user-signup/cmd/middlewares"
	configmanager "github.com/kilvish/user-signup/internal/configmanager"
	mysql "github.com/kilvish/user-signup/internal/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var host = flag.String("host", "localhost", "HOst ip")
var port = flag.String("port", "8080", "HOst ip")

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1024KB"))
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middlewares.RequestID)
	e.Use(middlewares.Method)
	// initialize empty context

	//ctx := types.NewContext(nil, fmt.Sprint(os.Getpid()))

	config, err := configmanager.GetConfig()

	if err != nil {
		log.Println("Error initializing confg. Error = ", err)
		os.Exit(1)
	}
	err = mysql.InitMysqlConnection(config.MySQL)
	if err != nil {
		log.Println("Error initializing mysql. Error = ", err)
		os.Exit(1)
	}

	AddRoutes(e)
	if err := e.Start(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
		log.Println("Failed to start server!", err)
		os.Exit(1)
	}
	select {}
}
