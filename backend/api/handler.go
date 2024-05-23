package api

import (
	"apica-assignment/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	cacheService *service.CacheService
}

func NewHandler(cacheService *service.CacheService) *Handler {
	return &Handler{cacheService: cacheService}
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, found := h.cacheService.Get(key)
	if !found {
		log.Printf("GET failed: key not found %s\n", key)
		http.NotFound(w, r)
		return
	}
	log.Printf("GET successful: key %s\n", key)
	json.NewEncoder(w).Encode(value)
}

func (h *Handler) SetItem(w http.ResponseWriter, r *http.Request) {
	var item service.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("POST failed: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.cacheService.Set(item.Key, item.Value, time.Duration(item.Expiration)*time.Second)
	log.Printf("POST successful: key %s\n", item.Key)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	h.cacheService.Delete(key)
	log.Printf("DELETE successful: key %s\n", key)
	w.WriteHeader(http.StatusOK)
}

// Add more handler methods as needed
