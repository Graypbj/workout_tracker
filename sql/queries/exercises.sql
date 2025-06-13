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

-- name: UpdateExerciseByID :one
UPDATE exercises
SET name = $3, exercise_type = $4, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, name, exercise_type, created_at, updated_at;

-- name: DeleteExerciseByID :exec
DELETE FROM exercises
WHERE id = $1 AND user_id = $2;

-- name: ListExercisesByUser :many
SELECT id, name, exercise_type, created_at, updated_at
FROM exercises
WHERE user_id = $1
ORDER BY name ASC;

-- name: GetExerciseByID :one
SELECT id, name, exercise_type, created_at, updated_at
FROM exercises
WHERE id = $1 AND user_id = $2;

