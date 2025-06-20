// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: strength_training_sets.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createStrengthTrainingSet = `-- name: CreateStrengthTrainingSet :one
INSERT INTO strength_training_sets (id, user_id, session_id, set_number, reps, weight, created_at, updated_at)
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5,
	NOW(),
	NOW()
)
RETURNING id, user_id, session_id, set_number, reps, weight, created_at, updated_at
`

type CreateStrengthTrainingSetParams struct {
	UserID    uuid.UUID
	SessionID uuid.UUID
	SetNumber int32
	Reps      int32
	Weight    string
}

func (q *Queries) CreateStrengthTrainingSet(ctx context.Context, arg CreateStrengthTrainingSetParams) (StrengthTrainingSet, error) {
	row := q.db.QueryRowContext(ctx, createStrengthTrainingSet,
		arg.UserID,
		arg.SessionID,
		arg.SetNumber,
		arg.Reps,
		arg.Weight,
	)
	var i StrengthTrainingSet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SessionID,
		&i.SetNumber,
		&i.Reps,
		&i.Weight,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStrengthTrainingSetByID = `-- name: DeleteStrengthTrainingSetByID :exec
DELETE FROM strength_training_sets
WHERE id = $1 AND user_id = $2
`

type DeleteStrengthTrainingSetByIDParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteStrengthTrainingSetByID(ctx context.Context, arg DeleteStrengthTrainingSetByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteStrengthTrainingSetByID, arg.ID, arg.UserID)
	return err
}

const listStrengthTrainingSetsBySession = `-- name: ListStrengthTrainingSetsBySession :many
SELECT id, session_id, set_number, reps, weight, created_at, updated_at
FROM strength_training_sets
WHERE session_id = $1 AND user_id = $2
`

type ListStrengthTrainingSetsBySessionParams struct {
	SessionID uuid.UUID
	UserID    uuid.UUID
}

type ListStrengthTrainingSetsBySessionRow struct {
	ID        uuid.UUID
	SessionID uuid.UUID
	SetNumber int32
	Reps      int32
	Weight    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) ListStrengthTrainingSetsBySession(ctx context.Context, arg ListStrengthTrainingSetsBySessionParams) ([]ListStrengthTrainingSetsBySessionRow, error) {
	rows, err := q.db.QueryContext(ctx, listStrengthTrainingSetsBySession, arg.SessionID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListStrengthTrainingSetsBySessionRow
	for rows.Next() {
		var i ListStrengthTrainingSetsBySessionRow
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.SetNumber,
			&i.Reps,
			&i.Weight,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStrengthTrainingSetByID = `-- name: UpdateStrengthTrainingSetByID :one
UPDATE strength_training_sets
SET set_number = $3, reps = $4, weight = $5, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING id, session_id, set_number, reps, weight, created_at, updated_at
`

type UpdateStrengthTrainingSetByIDParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	SetNumber int32
	Reps      int32
	Weight    string
}

type UpdateStrengthTrainingSetByIDRow struct {
	ID        uuid.UUID
	SessionID uuid.UUID
	SetNumber int32
	Reps      int32
	Weight    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) UpdateStrengthTrainingSetByID(ctx context.Context, arg UpdateStrengthTrainingSetByIDParams) (UpdateStrengthTrainingSetByIDRow, error) {
	row := q.db.QueryRowContext(ctx, updateStrengthTrainingSetByID,
		arg.ID,
		arg.UserID,
		arg.SetNumber,
		arg.Reps,
		arg.Weight,
	)
	var i UpdateStrengthTrainingSetByIDRow
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.SetNumber,
		&i.Reps,
		&i.Weight,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
