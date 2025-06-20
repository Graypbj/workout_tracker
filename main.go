package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Graypbj/workout_tracker/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	jwtSecret      string
}

func main() {
	fmt.Println("Hello world!")

	const filepathRoot = "."
	const port = "8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if platform == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
		jwtSecret:      jwtSecret,
	}

	mux := http.NewServeMux()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerUsersUpdate)
	mux.HandleFunc("DELETE /api/users", apiCfg.handlerUsersDelete)

	mux.HandleFunc("POST /api/exercises", apiCfg.handlerExercisesCreate)
	mux.HandleFunc("PUT /api/exercises", apiCfg.handlerExercisesUpdate)
	mux.HandleFunc("DELETE /api/exercises", apiCfg.handlerExercisesDelete)
	mux.HandleFunc("GET /api/exercises", apiCfg.handlerExercisesRetrieve)

	mux.HandleFunc("POST /api/workouts", apiCfg.handlerWorkoutsCreate)
	mux.HandleFunc("PUT /api/workouts", apiCfg.handlerWorkoutsUpdate)
	mux.HandleFunc("DELETE /api/workouts", apiCfg.handlerWorkoutsDelete)
	mux.HandleFunc("GET /api/workouts", apiCfg.handlerWorkoutsRetrieve)

	mux.HandleFunc("POST /api/strength_training_sessions", apiCfg.handlerStrengthTrainingSessionsCreate)
	mux.HandleFunc("PUT /api/strength_training_sessions", apiCfg.handlerStrengthTrainingSessionsUpdate)
	mux.HandleFunc("DELETE /api/strength_training_sessions", apiCfg.handlerStrengthTrainingSessionsDelete)
	mux.HandleFunc("GET /api/strength_training_sessions", apiCfg.handlerStrengthTrainingSessionsRetrieve)

	mux.HandleFunc("POST /api/strength_training_sets", apiCfg.handlerStrengthTrainingSetsCreate)
	mux.HandleFunc("PUT /api/strength_training_sets", apiCfg.handlerStrengthTrainingSetsUpdate)
	mux.HandleFunc("DELETE /api/strength_training_sets", apiCfg.handlerStrengthTrainingSetsDelete)
	mux.HandleFunc("GET /api/strength_training_sets", apiCfg.handlerStrengthTrainingSetsRetrieve)

	mux.HandleFunc("POST /api/cardio_training_sessions", apiCfg.handlerCardioTrainingSessionsCreate)
	mux.HandleFunc("PUT /api/cardio_training_sessions", apiCfg.handlerCardioTrainingSessionsUpdate)
	mux.HandleFunc("DELETE /api/cardio_training_sessions", apiCfg.handlerCardioTrainingSessionsDelete)
	mux.HandleFunc("GET /api/cardio_training_sessions", apiCfg.handlerCardioTrainingSessionsRetrieve)

	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerCount)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: withCORS(mux),
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
