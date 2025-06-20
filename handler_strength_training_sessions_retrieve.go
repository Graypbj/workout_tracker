package main

import (
	"net/http"

	"github.com/Graypbj/workout_tracker/internal/auth"
	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStrengthTrainingSessionsRetrieve(w http.ResponseWriter, r *http.Request) {
	type response struct {
		StrengthTrainingSessions []StrengthTrainingSession `json:"strength_training_sessions"`
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

	worktoutIDStr := r.URL.Query().Get("workout_id")
	if worktoutIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Missing workout_id query parameter", err)
		return
	}

	workoutID, err := uuid.Parse(worktoutIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid workout_id format", err)
		return
	}

	dbStrengthTrainingSessions, err := cfg.db.ListStrengthTrainingSessionsByWorkout(r.Context(), database.ListStrengthTrainingSessionsByWorkoutParams{
		WorkoutID: workoutID,
		UserID:    userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve strength training sessions", err)
		return
	}

	sessions := make([]StrengthTrainingSession, len(dbStrengthTrainingSessions))
	for i, dbSess := range dbStrengthTrainingSessions {
		sessions[i] = StrengthTrainingSession{
			ID:         dbSess.ID,
			WorkoutID:  dbSess.WorkoutID,
			ExerciseID: dbSess.ExerciseID,
			Notes:      dbSess.Notes.String,
			CreatedAt:  dbSess.CreatedAt,
			UpdatedAt:  dbSess.UpdatedAt,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		StrengthTrainingSessions: sessions,
	})
}
