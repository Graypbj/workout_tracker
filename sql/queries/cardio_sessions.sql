-- name: CreateCardioSession :one
INSERT INTO cardio_sessions (id, user_id, workout_id, exercise_id, distance, time, notes, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateCardioSessionByID :one
UPDATE cardio_sessions
SET distance = $3, time = $4, notes = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, workout_id, exercise_id, distance, time, notes, created_at, updated_at;

-- name: DeleteCardioSessionByID :exec
DELETE FROM cardio_sessions
WHERE id = $1 AND user_id = $2;

-- name: ListCardioSessionsByWorkout :many
SELECT id, workout_id, exercise_id, distance, time, notes, created_at, updated_at
FROM cardio_sessions
WHERE workout_id = $1 AND user_id = $2;
