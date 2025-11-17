package main

import (
	"log"
	"os"

	"myserver/server"
)

func main() {
	domains := os.Getenv("DOMAINS")
	if domains == "" {
		log.Fatal("Переменная окружения DOMAINS не установлена. Укажите ваш домен, например: export DOMAINS=example.com,www.example.com")
	}

	email := os.Getenv("EMAIL")
	if email == "" {
		log.Fatal("Переменная окружения EMAIL не установлена. Укажите email для уведомлений Let's Encrypt")
	}

	cfg := &server.Config{
		Domains:    []string{domains},
		Email:      email,
		CacheDir:   "./certs",
		Production: true,
	}

	srv := server.New(cfg)

	go srv.GracefulShutdown()

	if err := srv.Start(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
