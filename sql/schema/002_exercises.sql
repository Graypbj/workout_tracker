-- +goose Up
CREATE TABLE exercises (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	name VARCHAR(100) NOT NULL,
	exercise_type VARCHAR(50) CHECK (exercise_type IN ('strength_training', 'cardio')) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE exercises;

