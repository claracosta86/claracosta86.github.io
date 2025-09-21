package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"poc2/back/model"
	"poc2/back/service"
	"poc2/back/lib/logging"

)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler(es service.EventService) *EventHandler {
	return &EventHandler{
		eventService: es,
	}
}

// @ Register a new event handler
// @ Accept json
// /events/register [POST]
func (h *EventHandler) HandleRegisterEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.eventService.RegisterEvent(event)
	if err != nil {
		http.Error(w, "Erro ao salvar evento", http.StatusInternalServerError)
		return
	}

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusNoContent}
	rec.WriteHeader(http.StatusNoContent)
}

// @ Gets all events
// /events/all [GET]
func (h *EventHandler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.eventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

// @ Gets an event by its ID
// /events/:id [GET]
func (h *EventHandler) HandleGetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	event, err := h.eventService.GetEventByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(event)
}

// @ Updates an event
// @ Accepts JSON
// /events/update [PATCH]
func (h *EventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = h.eventService.UpdateEvent(event)
	if err != nil {
		http.Error(w, "Erro ao atualizar evento", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}

// Deletes an event by its ID
// /events/delete/:id [DELETE]
func (h *EventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "id")
	ID, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.eventService.DeleteEventByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}