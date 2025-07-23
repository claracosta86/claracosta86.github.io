package api

import (
	"encoding/json"
	"net/http"
	"log"

	"poc2/back/model"
	"poc2/back/service"
	"poc2/back/lib/logging"

)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventService: service.NewEventService(),
	}
}

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

	log.Printf("Evento recebido: %+v\n", event)
	// err = h.eventService.RegisterEvent(event)
	// if err != nil {
	// 	http.Error(w, "Erro ao salvar evento", http.StatusInternalServerError)
	// 	return
	// }

	rec := &logging.StatusRecorder{ResponseWriter: w, Status: http.StatusOK}
	rec.WriteHeader(http.StatusOK)
}

// /events/list [GET]
func (h *EventHandler) HandleGetUserEvents(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	events, err := h.eventService.GetUserEvents(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

// /events/all [GET]
func (h *EventHandler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.eventService.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}

// /events/:id [GET]
func (h *EventHandler) HandleGetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID := r.Context().Value("eventID").(string)
	event, err := h.eventService.GetEventByID(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(event)
}

// /events/update [PUT]
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

// /events/delete/:id [DELETE]
func (h *EventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.Context().Value("eventID").(string)
	err := h.eventService.DeleteEventByID(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"sucesso"}`))
}