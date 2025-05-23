-- name: CreateStrengthTrainingSet :exec
INSERT INTO strength_training_sets (session_id, set_number, reps, weight, created_at, updated_at)
VALUES (
	$1,
	$2,
	$3,
	$4,
	NOW(),
	NOW()
);
