package server

import (
	"net/http"

	"github.com/QianJiuGe/mysite/backend/internal/service"
)

func NewHTTPHandler(svc *service.Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", svc.Healthz)
	mux.HandleFunc("/v1/auth/register", svc.Register)
	mux.HandleFunc("/v1/auth/login", svc.Login)
	mux.HandleFunc("/v1/admin/users/pending", svc.PendingUsers)
	mux.HandleFunc("/v1/admin/users/", svc.ApproveUser)
	mux.HandleFunc("/v1/home", svc.Home)

	return withCORS(mux)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
