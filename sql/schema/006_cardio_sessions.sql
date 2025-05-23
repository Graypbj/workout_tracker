-- +goose Up
CREATE TABLE cardio_sessions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	workout_id UUID REFERENCES workouts(id) ON DELETE CASCADE NOT NULL,
	exercise_id UUID REFERENCES exercises(id) NOT NULL,
	distance DECIMAL(6, 2) NOT NULL,
	time INTERVAL NOT NULL,
	notes TEXT,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE cardio_sessions;

