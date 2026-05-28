package main

import (
	"encoding/json"
	"net/http"
)

type App struct{ store *Store }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	stations := a.store.All()
	writeJSON(w, http.StatusOK, stations)
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	station, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "station introuvable")
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
		writeError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	if station.ID == "" {
		writeError(w, http.StatusBadRequest, "id manquant")
		return
	}

	if a.store.Has(station.ID) {
		writeError(w, http.StatusConflict, "id déjà utilisé")
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
		writeError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	if station.ID != "" && station.ID != id {
		writeError(w, http.StatusBadRequest, "id du body différent de l'id de l'URL")
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
