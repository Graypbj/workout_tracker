-- name: CreateStrengthTrainingSet :one
INSERT INTO strength_training_sets (id, user_id, session_id, set_number, reps, weight, created_at, updated_at)
VALUES (
	gen_random_uuid,
	$1,
	$2,
	$3,
	$4,
	$5,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateStrengthTrainingSet :one
UPDATE strength_training_sets
SET set_number = $3, reps = $4, weight = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, session_id, set_number, reps, weight, created_at, updated_at;

-- name: DeleteStrengthTrainingSet :exec
DELETE FROM strength_training_sets
WHERE id = $1 AND user_id = $2;

