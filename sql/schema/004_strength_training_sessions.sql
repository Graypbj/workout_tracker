-- +goose Up
CREATE TABLE strength_training_sessions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	workout_id UUID REFERENCES workouts(id) ON DELETE CASCADE,
	exercise_id UUID REFERENCES exercises(id),
	notes TEXT
);

-- +goose Down
DROP TABLE strength_training_sessions;

