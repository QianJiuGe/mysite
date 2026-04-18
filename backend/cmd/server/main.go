package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/QianJiuGe/mysite/backend/internal/biz"
	"github.com/QianJiuGe/mysite/backend/internal/conf"
	"github.com/QianJiuGe/mysite/backend/internal/data"
	"github.com/QianJiuGe/mysite/backend/internal/server"
	"github.com/QianJiuGe/mysite/backend/internal/service"
)

func main() {
	configPath := flag.String("config", "./configs/app.example.yaml", "config file path")
	flag.Parse()

	cfg, err := conf.Load(*configPath)
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	store, err := data.NewStore(cfg.MySQL.DSN, cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Fatalf("init data store failed: %v", err)
	}
	defer func() {
		if err := store.Close(); err != nil {
			log.Printf("close store: %v", err)
		}
	}()

	uc := biz.NewUsecase(store)
	if err := uc.BootstrapAdmin(context.Background(), cfg.Auth.DefaultAdminUsername, cfg.Auth.DefaultAdminPassword); err != nil {
		log.Fatalf("bootstrap default admin failed: %v", err)
	}

	svc := service.New(uc)
	h := server.NewHTTPHandler(svc)

	log.Printf("backend listening on %s", cfg.Server.HTTPAddr)
	if err := http.ListenAndServe(cfg.Server.HTTPAddr, h); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
