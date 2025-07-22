package api

import (
	"encoding/json"
	"log"
	"net/http"

	"poc2/back/model"
	"poc2/back/service"

)

type EventHandler struct {
	eventService service.EventService
}

func EventHandler() *EventHandler {
	return &EventHandler{
		eventService: service.NewEventService(),
	}
}

func (h *EventHandler) HandleGetUserEvents(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	events, err := h.eventService.GetUserEvents(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}

func (h *EventHandler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.eventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (h *EventHandler) HandleGetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID := r.Context().Value("eventID").(string)
	event, err := h.eventService.GetEventByID(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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