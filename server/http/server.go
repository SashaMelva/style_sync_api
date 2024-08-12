package http

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/SashaMelva/style_sync_api/internal/app"
	"github.com/SashaMelva/style_sync_api/internal/config"
	"github.com/SashaMelva/style_sync_api/server/hendler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
	log *zap.SugaredLogger
}

var routes = flag.Bool("routes", false, "Generate router documentation")

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func NewServer(log *zap.SugaredLogger, app *app.App, config *config.ConfigHttpServer) *Server {
	log.Info("URL api " + config.Host + ":" + config.Port)
	log.Debug("URL api running " + config.Host + ":" + config.Port)

	h := hendler.NewHendler(log, app)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Post("/auth", h.Authorization)
	router.Post("/un_auth", h.Authorization)
	router.Post("/reg", h.Registration)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Mount("/user", userRouter(h))
		r.Mount("/weather_report", weatherReportRouter(h))
		r.Mount("/", wardrobeRouter(h))

	})

	if *routes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi/v5",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		return &Server{
			srv: &http.Server{
				Addr:    config.Host + ":" + config.Port,
				Handler: router,
			},
			log: log,
		}
	}
	return nil
}

func (s *Server) Start(ctx context.Context) {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Fatal("listen: ", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(1)
}
