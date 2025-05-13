-- +goose Up
CREATE TABLE workouts (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	workout_date TIMESTAMP DEFAULT NOW(),
	workout_type VARCHAR(50) CHECK (workout_type IN ('strength_training', 'cardio')),
	notes TEXT
);

-- +goose Down
DROP TABLE workouts;

