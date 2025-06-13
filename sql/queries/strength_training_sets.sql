-- name: CreateStrengthTrainingSet :one
INSERT INTO strength_training_sets (id, user_id, session_id, set_number, reps, weight, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateStrengthTrainingSetByID :one
UPDATE strength_training_sets
SET set_number = $3, reps = $4, weight = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, session_id, set_number, reps, weight, created_at, updated_at;

-- name: DeleteStrengthTrainingSetByID :exec
DELETE FROM strength_training_sets
WHERE id = $1 AND user_id = $2;

-- name: ListStrengthTrainingSetsBySession :many
SELECT id, session_id, set_number, reps, weight, created_at, updated_at
FROM strength_training_sets
WHERE session_id = $1 AND user_id = $2;
