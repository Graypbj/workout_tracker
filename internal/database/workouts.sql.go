// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: workouts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO workouts (id, user_id, workout_date, workout_type, notes, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	NOW(),
	$2,
	$3,
	NOW(),
	NOW()
)
RETURNING id, user_id, workout_date, workout_type, notes, created_at, updated_at
`

type CreateWorkoutParams struct {
	UserID      uuid.UUID
	WorkoutType string
	Notes       sql.NullString
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout, arg.UserID, arg.WorkoutType, arg.Notes)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.WorkoutDate,
		&i.WorkoutType,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateWorkout = `-- name: UpdateWorkout :one
UPDATE workouts
SET workout_date = $3, workout_type = $4, notes = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING workout_date, workout_type, notes, created_at, updated_at
`

type UpdateWorkoutParams struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	WorkoutDate time.Time
	WorkoutType string
	Notes       sql.NullString
}

type UpdateWorkoutRow struct {
	WorkoutDate time.Time
	WorkoutType string
	Notes       sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (UpdateWorkoutRow, error) {
	row := q.db.QueryRowContext(ctx, updateWorkout,
		arg.ID,
		arg.UserID,
		arg.WorkoutDate,
		arg.WorkoutType,
		arg.Notes,
	)
	var i UpdateWorkoutRow
	err := row.Scan(
		&i.WorkoutDate,
		&i.WorkoutType,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
