package main

import (
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
)

func (cfg *apiConfig) handlerExercisesRetrieve(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Exercises []Exercise `json:"exercises"`
	}

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

	dbExercises, err := cfg.db.ListExercisesByUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve exercises", err)
		return
	}

	exercises := make([]Exercise, len(dbExercises))
	for i, dbEx := range dbExercises {
		exercises[i] = Exercise{
			ID:           dbEx.ID,
			Name:         dbEx.Name,
			ExerciseType: dbEx.ExerciseType,
			CreatedAt:    dbEx.CreatedAt,
			UpdatedAt:    dbEx.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		Exercises: exercises,
	})
}
