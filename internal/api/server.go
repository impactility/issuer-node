package api

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/iden3/iden3comm/v2"

	"github.com/polygonid/sh-id-platform/internal/config"
	"github.com/polygonid/sh-id-platform/internal/core/ports"
	"github.com/polygonid/sh-id-platform/internal/health"
	"github.com/polygonid/sh-id-platform/internal/network"
)

// Server implements StrictServerInterface and holds the implementation of all API controllers
// This is the glue to the API autogenerated code
type Server struct {
	cfg                *config.Configuration
	accountService     ports.AccountService
	claimService       ports.ClaimService
	connectionsService ports.ConnectionService
	health             *health.Status
	identityService    ports.IdentityService
	linkService        ports.LinkService
	networkResolver    network.Resolver
	packageManager     *iden3comm.PackageManager
	publisherGateway   ports.Publisher
	qrService          ports.QrStoreService
	schemaService      ports.SchemaService
}

// NewServer is a Server constructor
func NewServer(cfg *config.Configuration, identityService ports.IdentityService, accountService ports.AccountService, connectionsService ports.ConnectionService, claimsService ports.ClaimService, qrService ports.QrStoreService, publisherGateway ports.Publisher, packageManager *iden3comm.PackageManager, networkResolver network.Resolver, health *health.Status, schemaService ports.SchemaService, linkService ports.LinkService) *Server {
	return &Server{
		cfg:                cfg,
		accountService:     accountService,
		claimService:       claimsService,
		connectionsService: connectionsService,
		health:             health,
		identityService:    identityService,
		linkService:        linkService,
		networkResolver:    networkResolver,
		publisherGateway:   publisherGateway,
		packageManager:     packageManager,
		qrService:          qrService,
		schemaService:      schemaService,
	}
}

// Health is a method
func (s *Server) Health(_ context.Context, _ HealthRequestObject) (HealthResponseObject, error) {
	var resp Health200JSONResponse = s.health.Status()

	return resp, nil
}

// RegisterStatic add method to the mux that are not documented in the API.
func RegisterStatic(mux *chi.Mux) {
	mux.Get("/docs", documentation)
	mux.Get("/static/docs/api/api.yaml", swagger)
	mux.Get("/favicon.ico", favicon)
	fs := http.FileServer(http.Dir("./ui/dist/"))
	mux.Get("/schemas", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
	mux.Route("/credentials", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})
	mux.Route("/connections", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})
	mux.Route("/issuer-state", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})
	mux.Route("/identities", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})
	mux.Handle("/*", http.StripPrefix("/", fs))
	mux.Handle("/assets/*", fs)
	mux.Handle("/images/*", fs)
}

func documentation(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/spec.html", "text/html; charset=UTF-8", w)
}

func favicon(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/privadoid.png", "image/png", w)
}

func swagger(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/api.yaml", "text/html; charset=UTF-8", w)
}

func writeFile(path string, mimeType string, w http.ResponseWriter) {
	f, err := os.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(f)
}
