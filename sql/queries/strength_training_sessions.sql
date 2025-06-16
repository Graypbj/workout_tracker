-- name: CreateStrengthTrainingSession :one
INSERT INTO strength_training_sessions (id, user_id, workout_id, exercise_id, notes, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateStrengthTrainingSessionByID :one
UPDATE strength_training_sessions
SET workout_id = $3, exercise_id = $4, notes = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, workout_id, exercise_id, notes, created_at, updated_at;

-- name: DeleteStrengthTrainingSessionByID :exec
DELETE FROM strength_training_sessions
WHERE id = $1 AND user_id = $2;

-- name: ListStrengthTrainingSessionsByWorkout :many
SELECT id, workout_id, exercise_id, notes, created_at, updated_at
FROM strength_training_sessions
WHERE workout_id = $1 AND user_id = $2;
