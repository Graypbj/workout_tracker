-- name: CreateExercise :one
INSERT INTO exercises (id, user_id, name, exercise_type)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3
)
RETURNING *;
