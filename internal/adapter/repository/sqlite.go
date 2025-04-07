package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/rkperes/blog/db"
	"github.com/rkperes/blog/db/queries"
	"github.com/rkperes/blog/internal/core/domain"
	"github.com/rkperes/blog/internal/core/domain/errs"
)

type SQLiteRepository struct {
	q *queries.Queries
}

func NewSQLiteRepository() (*SQLiteRepository, error) {
	db, err := db.NewSQLLite()
	if err != nil {
		return nil, err
	}

	return &SQLiteRepository{
		q: queries.New(db),
	}, nil
}

func (r *SQLiteRepository) GetSession(ctx context.Context, sessionID string) (domain.Session, error) {
	s, err := r.q.GetSessionById(ctx, sessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Session{}, errs.ErrNotFound
		}
		return domain.Session{}, err
	}

	return toDomainSession(s), nil
}

func (r *SQLiteRepository) CreateSession(ctx context.Context, userID string) (domain.Session, error) {
	s, err := r.q.CreateSession(ctx, queries.CreateSessionParams{
		ID:     uuid.New().String(),
		UserID: userID,
	})
	if err != nil {
		return domain.Session{}, err
	}

	slog.Info("new session", slog.String("session_id", s.ID))
	return toDomainSession(s), nil
}

func toDomainSession(s queries.Session) domain.Session {
	return domain.Session{
		ID:     domain.UUID(s.ID),
		UserID: domain.UUID(s.UserID),
	}
}
