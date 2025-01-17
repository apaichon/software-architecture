package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vi-cqrs/commands"
	"vi-cqrs/queries"
)

type Handler struct {
	commandHandler *commands.CommandHandler
	queryHandler   *queries.QueryHandler
}

func NewHandler(ch *commands.CommandHandler, qh *queries.QueryHandler) *Handler {
	return &Handler{
		commandHandler: ch,
		queryHandler:   qh,
	}
}

func (h *Handler) PurchaseTicket(w http.ResponseWriter, r *http.Request) {
	var cmd commands.PurchaseTicketCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.commandHandler.HandlePurchaseTicket(cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetConcert(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid concert ID", http.StatusBadRequest)
		return
	}

	concert, err := h.queryHandler.HandleGetConcertByID(queries.GetConcertByIDQuery{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concert)
}

func (h *Handler) GetAvailableConcerts(w http.ResponseWriter, r *http.Request) {
	concerts, err := h.queryHandler.HandleGetAvailableConcerts(queries.GetAvailableConcertsQuery{MinAvailableSeats: 1})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concerts)
}

func (h *Handler) CreateConcert(w http.ResponseWriter, r *http.Request) {
	var cmd commands.CreateConcertCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.commandHandler.HandleCreateConcert(cmd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
