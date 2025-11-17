package server

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

// Config хранит настройки сервера
type Config struct {
	Domains    []string
	Email      string
	CacheDir   string
	Production bool
}

type Server struct {
	config      *Config
	httpServer  *http.Server
	certManager *autocert.Manager
}

func New(cfg *Config) *Server {

	certManager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,

		HostPolicy: autocert.HostWhitelist(cfg.Domains...),

		Cache: autocert.DirCache(cfg.CacheDir),

		Email: cfg.Email,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/api/info", APIHandler)

	httpServer := &http.Server{
		Addr:    ":443",
		Handler: mux,

		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,

		TLSConfig: &tls.Config{

			GetCertificate: certManager.GetCertificate,

			MinVersion: tls.VersionTLS12,

			PreferServerCipherSuites: true,

			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			},
		},
	}

	return &Server{
		config:      cfg,
		httpServer:  httpServer,
		certManager: certManager,
	}
}

func (s *Server) Start() error {

	go func() {
		log.Println("Запуск HTTP-сервера на порту 80 для ACME challenge...")
		if err := http.ListenAndServe(":80", s.certManager.HTTPHandler(nil)); err != nil {
			log.Printf("HTTP-сервер остановлен: %v", err)
		}
	}()

	log.Printf("Запуск HTTPS-сервера на порту 443 для доменов: %v", s.config.Domains)
	log.Println("Сертификаты будут получены автоматически от Let's Encrypt")
	log.Printf("Сертификаты сохраняются в: %s", s.config.CacheDir)

	return s.httpServer.ListenAndServeTLS("", "")
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Получен сигнал остановки. Завершаем работу...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Ошибка при остановке сервера: %v", err)
	}

	log.Println("Сервер остановлен корректно")
}
