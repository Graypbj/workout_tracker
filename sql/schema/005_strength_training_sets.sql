-- +goose Up
CREATE TABLE strength_training_sets (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	session_id UUID REFERENCES strength_training_sessions(id) ON DELETE CASCADE NOT NULL,
	set_number INT NOT NULL,
	reps INT NOT NULL,
	weight DECIMAL(5, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE strength_training_sets;

