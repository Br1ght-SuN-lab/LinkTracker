package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
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