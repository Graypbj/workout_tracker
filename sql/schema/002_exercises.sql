-- +goose Up
CREATE TABLE exercises (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	exercise_type VARCHAR(50) CHECK (exercise_type IN ('strength_training', 'cardio')) NOT NULL
);

-- +goose Down
DROP TABLE exercises;

