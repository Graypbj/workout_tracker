-- name: CreateStrengthTrainingSession :one
INSERT INTO strength_training_sessions (id, workout_id, exercise_id, notes, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateStrengthTrainingSession :one
UPDATE strength_training_sessions
SET workout_id = $3, exercise_id = $4, notes = $5, updated_at = NOW()
WHERE workout_id = $1 AND exercise_id = $2
RETURNING id, workout_id, exercise_id, notes, created_at, updated_at;

-- name: DeleteStrengthTrainingSession :exec
DELETE FROM strength_training_sessions
WHERE id = $1;

