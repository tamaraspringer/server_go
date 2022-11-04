package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"server_mod/config"
	"server_mod/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	cfg := config.GetConfig()
	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	log.Println("start app")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Type))
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server ist listening port ")
	log.Fatalln(server.Serve(listener))

}
