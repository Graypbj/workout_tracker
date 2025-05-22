-- name: CreateExercise :one
INSERT INTO exercises (id, name, exercise_type, notes)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3
)
RETURNING *;
