package http

import (
	nethttp "net/http"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/config"
)

func NewServer(service *application.Service, cfg *config.Config) *nethttp.Server {
	handler := NewHandler(service)
	mux := nethttp.NewServeMux()

	mux.HandleFunc("/tg-chat/", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		switch r.Method{
		case nethttp.MethodPost:
			handler.RegisterChat(w, r)
		case nethttp.MethodDelete:
			handler.DeleteChat(w, r)
		default:
			w.WriteHeader(nethttp.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/links", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		switch r.Method {
		case nethttp.MethodGet:
			handler.ListLinks(w, r)
		case nethttp.MethodPost:
			handler.AddLink(w, r)
		case nethttp.MethodDelete:
			handler.RemoveLink(w, r)
		default:
			w.WriteHeader(nethttp.StatusMethodNotAllowed)
		}
	})
	
	return &nethttp.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}
}