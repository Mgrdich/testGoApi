package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	".com/configs"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	Router  *chi.Mux
	SqlConn *pgx.Conn
}

func NewServer(sqlConn *pgx.Conn) *Server {
	server := &Server{
		Router:  chi.NewRouter(),
		SqlConn: sqlConn,
	}

	return server
}

func (s *Server) Start(ctx context.Context) {
	addr := fmt.Sprintf(":%s", configs.GetAppConfig().Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      s.Router,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second * 2,
	}

	shutdownCompleted := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed %v\n", err)
		}
	})

	fmt.Printf("Server is starting at %s", addr)

	if err := server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		<-shutdownCompleted
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
