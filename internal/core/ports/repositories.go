package ports

import "github.com/rkperes/blog/internal/core/domain"

type SessionRepository interface {
	GetSession(sessionID string) (domain.Session, error)
}

type UserRepository interface{}
