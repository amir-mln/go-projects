package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amir-mln/go-projects/micro-cafe/product-management/app"
	"github.com/amir-mln/go-projects/micro-cafe/product-management/app/handlers"
	"github.com/amir-mln/go-projects/micro-cafe/product-management/infrastructure/repositories"
)

func main() {
	l := log.New(os.Stdout, "\tproduct-api\t", log.Flags())
	pr := repositories.NewProductRepository()
	as := app.NewApplicationService(pr)

	productsHandler := handlers.NewProductHandler(l, as)

	serveMux := http.NewServeMux()
	serveMux.Handle(handlers.ProductHandlerRoute, productsHandler)

	server := http.Server{
		Addr:         ":3000",
		Handler:      serveMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	l.Println("Gracefully terminating the server...", <-sigChan)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(tc)
}
