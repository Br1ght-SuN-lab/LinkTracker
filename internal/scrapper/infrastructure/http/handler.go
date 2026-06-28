package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/client"
)

type Handler struct {
	service *application.Service
}


func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type ErrorResponse  struct {
	Description string `yaml:"description"`
}


func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Description: msg,
	})
}

func (h *Handler) RegisterChat(w http.ResponseWriter, r *http.Request) {
	chatID, err := extractChatID(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid chat id")
		return
	}

	if err := h.service.RegisterChat(r.Context(), chatID); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
}


func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	chatID, err := extractChatID(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid chat id")
		return
	}

	if err := h.service.DeleteChat(r.Context(), chatID); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
}


func extractChatID(path string) (int64, error) {
	clearPath := strings.Trim(path, "/")
	tokens := strings.Split(clearPath, "/")

	if len(tokens) != 2 {
		return 0, fmt.Errorf("invalid path")
	} 

	if tokens[0] != "tg-chat" {
		return 0, fmt.Errorf("invalid path")
	}

	chatID, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid chat id: %w", err)
	}

	return chatID, nil 
}

func extractTgChatID(r *http.Request) (int64, error) {
	value := r.Header.Get("Tg-Chat-Id")
	if value == "" {
		return 0, fmt.Errorf("missing Tg-Chat-Id header")
	}

	chatID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid Tg-Chat-Id header")
	}

	return chatID, nil
}

func (h *Handler) AddLink(w http.ResponseWriter, r *http.Request) {
	chatID, err := extractTgChatID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req client.AddLinkRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	link, err := h.service.AddLink(r.Context(), chatID, req.Link, req.Tags, req.Filters)
	if err != nil {
		if err.Error() == "chat not found" {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	resp := client.LinkResponseDTO{
		ID:      link.ID,
		URL:     link.URL,
		Tags:    link.Tags,
		Filters: link.Filters,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListLinks(w http.ResponseWriter, r *http.Request) {
	chatID, err := extractTgChatID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	links, err := h.service.ListLinks(r.Context(), chatID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	respLinks := make([]client.LinkResponseDTO, 0, len(links))
	for _, link := range links {
		respLinks = append(respLinks, client.LinkResponseDTO{
			ID:      link.ID,
			URL:     link.URL,
			Tags:    link.Tags,
			Filters: link.Filters,
		})
	}

	resp := client.ListLinksResponseDTO{
		Links: respLinks,
		Size:  int32(len(respLinks)),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) RemoveLink(w http.ResponseWriter, r *http.Request) {
	chatID, err := extractTgChatID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req client.RemoveLinkRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	link, err := h.service.RemoveLink(r.Context(), chatID, req.Link)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	resp := client.LinkResponseDTO{
		ID:      link.ID,
		URL:     link.URL,
		Tags:    link.Tags,
		Filters: link.Filters,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}