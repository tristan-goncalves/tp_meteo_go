package main

import (
	"encoding/json"
	"net/http"
)

type App struct{ store *Store }

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code string, msg string) {
	writeJSON(w, status, ErrorResponse{
		Error: msg,
		Code:  code,
	})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	stations := a.store.All()
	writeJSON(w, http.StatusOK, stations)
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	station, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "La station est introuvable")
		return
	}
	writeJSON(w, http.StatusOK, station)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Pour fermer la lecture
	var station Station

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Pour refuser ce qui n'existe pas dans station

	err := decoder.Decode(&station)
	if err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "Le JSON est invalide")
		return
	}
	if station.ID == "" {
		writeError(w, http.StatusBadRequest, "MISSING_ID", "L'id est manquant")
		return
	}
	if a.store.Has(station.ID) {
		writeError(w, http.StatusConflict, "ID_TAKEN", "L'id est déjà utilisé")
		return
	}

	a.store.Put(station)
	writeJSON(w, http.StatusCreated, station)
}

func (a *App) updateStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // Pour fermer la lecture
	id := r.PathValue("id")

	var station Station
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Pour refuser ce qui n'existe pas dans station

	err := decoder.Decode(&station)
	if err != nil {
		writeError(w, http.StatusBadRequest, "BAD_JSON", "Le JSON est invalide")
		return
	}

	if station.ID != "" && station.ID != id {
		writeError(w, http.StatusBadRequest, "ID_MISMATCH", "L'id du body différent de l'id de l'URL")
		return
	}

	created := !a.store.Has(id)
	station.ID = id
	a.store.Put(station)

	if created {
		writeJSON(w, http.StatusCreated, station)
		return
	}
	writeJSON(w, http.StatusOK, station)
}

func (a *App) deleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ok := a.store.Delete(id)

	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "La station est introuvable")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *App) listObservations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	station, ok := a.store.Get(id)

	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "La station est introuvable")
		return
	}
	writeJSON(w, http.StatusOK, station.Observations)
}
