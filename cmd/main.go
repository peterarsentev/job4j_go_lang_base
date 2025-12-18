package main

import (
	"context"
	"job4j.ru/go-lang-base/internal/db"
	item "job4j.ru/go-lang-base/internal/repository"
	"job4j.ru/go-lang-base/internal/tracker"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	cfg := db.Config{
		Host:     env("DB_HOST", "localhost"),
		Port:     envInt("DB_PORT", 5432),
		User:     env("DB_USER", "postgres"),
		Password: env("DB_PASSWORD", "password"),
		DBName:   env("DB_NAME", "tracker"),
		SSLMode:  env("DB_SSLMODE", "disable"),
	}

	pool, err := db.NewPool(ctx, cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := item.NewRepoPg(pool)

	it := tracker.Item{
		ID:   uuid.New().String(),
		Name: "first",
	}

	if err := repo.Create(ctx, it); err != nil {
		log.Fatal(err)
	}

	got, err := repo.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, it := range got {
		log.Printf("item: %+v", it)
	}
}

func env(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
