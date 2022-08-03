package wallet

import (
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/samthehai/simpleblockchain/server"
)

//go:embed frontend/dist
var frontend embed.FS

type walletServer struct {
	port    uint64
	gateway string
	server  *http.Server
}

func NewWalletServer(port uint64, gateway string) server.Server {
	return &walletServer{
		port:    port,
		gateway: gateway,
	}
}

func (s *walletServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	stripped, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		return err
	}

	r.Handle("/*", http.FileServer(http.FS(stripped)))

	s.server = &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(int(s.port)),
		Handler: r,
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *walletServer) Close() error {
	if s.server == nil {
		return nil
	}

	return s.server.Close()
}
