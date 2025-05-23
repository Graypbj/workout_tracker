-- +goose Up
CREATE TABLE workouts (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	workout_date TIMESTAMP DEFAULT NOW() NOT NULL,
	workout_type VARCHAR(50) CHECK (workout_type IN ('strength_training', 'cardio')) NOT NULL,
	notes TEXT,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE workouts;

