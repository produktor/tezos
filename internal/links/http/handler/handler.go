package handler

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"link_api/internal/links"
	"net/http"
)

type Handler struct {
	s      *links.LinkService
	logger *zap.SugaredLogger
}


func New(s *links.LinkService, logger *zap.SugaredLogger)  {
	r := mux.NewRouter()

	h := &Handler{
		s:      s,
		logger: logger,
	}

	r.HandleFunc("/groups", h.GroupsList).Methods("GET")
	r.HandleFunc("/link", h.Link).Methods("GET")
	r.HandleFunc("/links", h.Links).Methods("POST")

	http.ListenAndServe("localhost:3050", r)
}
