-- name: CreateExercise :one
INSERT INTO exercises (id, user_id, name, exercise_type, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	NOW(),
	NOW()
)
RETURNING *;

-- name: UpdateExercise :one
UPDATE exercises
SET name = $3, exercise_type = $4
WHERE id = $1 AND user_id = $2
RETURNING id, name, exercise_type, created_at, updated_at;
