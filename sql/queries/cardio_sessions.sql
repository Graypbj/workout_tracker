-- name: CreateCardioSession :one
INSERT INTO cardio_sessions (workout_id, exercise_id, distance, time, notes, created_at, updated_at)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	NOW(),
	NOW()
)
RETURNING *;
