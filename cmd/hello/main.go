package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/msyrus/hello-go/service"
	"github.com/msyrus/hello-go/web"
)

var port int
var host, msg string

func main() {
	hName, _ := os.Hostname()
	flag.StringVar(&host, "host", "", "server host")
	flag.IntVar(&port, "port", 8080, "server listening port")
	flag.StringVar(&msg, "name", hName, "server name")
	flag.Parse()

	if port != 0 {
		host = fmt.Sprintf("%s:%d", host, port)
	}

	gSvc, err := service.NewGreeting(msg)
	if err != nil {
		log.Fatalln(err)
	}

	errCh := make(chan error, 0)
	sigCh := make(chan os.Signal, 0)

	srvr := http.Server{
		Addr:    host,
		Handler: web.NewRouter(gSvc),
	}

	go func() {
		log.Println("Starting server on", host)
		err := srvr.ListenAndServe()
		errCh <- err
	}()

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)

	for i := 0; i < 2; i++ {
		select {
		case err := <-errCh:
			if err != nil {
				log.Fatalln(err)
			}
			break

		case <-sigCh:
			if i == 0 {
				d := 30 * time.Second
				log.Println("Suttingdown server gracefully with in", d)
				log.Println("To shutdown immedietly press again")
				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), d)
					defer cancel()

					if err := srvr.Shutdown(ctx); err != nil {
						log.Fatalln(err)
					}

					log.Println("Server shutteddown gracefully")
				}()
				continue
			}

			log.Println("Suttingdown server forcefully")
			if err := srvr.Close(); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
