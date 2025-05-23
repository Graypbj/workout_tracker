-- name: CreateStrengthTrainingSession :one
INSERT INTO strength_training_sessions (workout_id, exercise_id, notes, created_at, updated_at)
VALUES (
	$1,
	$2,
	$3,
	NOW(),
	NOW()
)
RETURNING *;
