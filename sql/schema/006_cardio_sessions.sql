-- +goose Up
CREATE TABLE cardio_sessions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	workout_id UUID REFERENCES workouts(id) ON DELETE CASCADE,
	exercise_id UUID REFERENCES exercises(id),
	distance DECIMAL(6, 2),
	time INTERVAL,
	notes TEXT
);

-- +goose Down
DROP TABLE cardio_sessions;

