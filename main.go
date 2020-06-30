package main

import (
	"github.com/shao1f/user.tag.logic/server/http"
	"github.com/shao1f/user.tag.logic/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init service
	srv := service.New()

	// init http server
	http.Init(srv)
	defer http.ShutDown()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("user.tag.logic server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
