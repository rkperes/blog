package ports

import (
	"context"

	"github.com/rkperes/blog/internal/core/domain"
)

type SessionRepository interface {
	GetSession(ctx context.Context, sessionID string) (domain.Session, error)
	CreateSession(ctx context.Context, userID string) (domain.Session, error)
}

type UserRepository interface{}
