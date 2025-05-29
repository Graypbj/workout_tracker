package main

import (
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
)

func (cfg *apiConfig) handlerUsersDelete(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	err = cfg.db.DeleteUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete user", err)
		return
	}

	var a any
	respondWithJSON(w, http.StatusOK, a)
}
