package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New implaments APIServer create helper
func New(c *Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIServer) Start() error {
	if err := s.configureLoger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("starting API Server")
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.BindAddr), s.router)
}

func (s *APIServer) configureLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	//
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
