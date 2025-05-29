-- name: CreateWorkout :one
INSERT INTO workouts (id, user_id, workout_date, workout_type, notes, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	NOW(),
	$2,
	$3,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateWorkout :one
UPDATE workouts
SET workout_date = $3, workout_type = $4, notes = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, workout_date, workout_type, notes, created_at, updated_at;

-- name: DeleteWorkout :exec
DELETE FROM workouts
WHERE id = $1 AND user_id = $2;

