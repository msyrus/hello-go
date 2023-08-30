package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	hellolog "github.com/msyrus/hello-go/log"
	pb "github.com/msyrus/hello-go/proto/v1"
	"github.com/msyrus/hello-go/rpc"
	"github.com/msyrus/hello-go/service"
	"github.com/msyrus/hello-go/version"
	"github.com/msyrus/hello-go/web"
	"github.com/msyrus/hello-go/web/middleware"
)

var (
	gPort, wPort int
	host, name   string
	showVersion  bool
)

var mdls = middleware.Group(
	middleware.Recover,
	middleware.Logger(hellolog.DefaultOutputLogger),
)

var unaryInterceptors = []grpc.UnaryServerInterceptor{}

func main() {
	hName, _ := os.Hostname()
	flag.StringVar(&host, "host", "", "server host")
	flag.IntVar(&wPort, "web", 8080, "web server listening port")
	flag.IntVar(&gPort, "grpc", 8081, "grpc server listening port")
	flag.StringVar(&name, "name", hName, "server name")
	flag.BoolVar(&showVersion, "version", false, "print version")
	flag.Parse()

	if showVersion {
		fmt.Println("Version:", version.Version)
		return
	}

	var wAddr, gAddr string
	if wPort != 0 {
		wAddr = fmt.Sprintf("%s:%d", host, wPort)
	}
	if gPort != 0 {
		gAddr = fmt.Sprintf("%s:%d", host, gPort)
	}

	gSvc, err := service.NewGreeting(name)
	if err != nil {
		log.Fatalln(err)
	}

	gSrvr := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
	)

	pb.RegisterGreetingServer(gSrvr, rpc.NewServer(gSvc))
	lsnr, err := net.Listen("tcp", gAddr)
	if err != nil {
		log.Fatalln(err)
	}

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	go func() {
		log.Println("Starting gRPC server on", gAddr)
		errCh <- gSrvr.Serve(lsnr)
	}()

	srvr := http.Server{
		Addr:    wAddr,
		Handler: mdls(web.NewRouter(gSvc)),
	}

	go func() {
		log.Println("Starting web server on", wAddr)
		errCh <- srvr.ListenAndServe()
	}()

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)

	breakIt := false
	for i := 0; i < 2 && !breakIt; i++ {
		select {
		case err := <-errCh:
			if err != nil {
				log.Fatalln(err)
			}
			breakIt = true

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
					log.Println("Web Server shutteddown gracefully")

					gSrvr.GracefulStop()
					log.Println("gRPC Server shutteddown gracefully")
				}()
				continue
			}

			log.Println("Suttingdown web server forcefully")
			if err := srvr.Close(); err != nil {
				log.Fatalln(err)
			}

			log.Println("Suttingdown gRPC server forcefully")
			gSrvr.Stop()
		}
	}
}
