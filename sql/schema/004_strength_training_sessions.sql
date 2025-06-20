-- +goose Up
CREATE TABLE strength_training_sessions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
	workout_id UUID REFERENCES workouts(id) ON DELETE CASCADE NOT NULL,
	exercise_id UUID REFERENCES exercises(id) NOT NULL,
	notes TEXT,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE strength_training_sessions;
