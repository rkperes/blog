package authctx

import (
	"context"
	"log/slog"

	"github.com/rkperes/blog/internal/core/domain"
)

type ctxKey string

const (
	authCookie    string = "session_token"
	ctxKeySession ctxKey = "session"
)

func WithSession(ctx context.Context, session domain.Session) context.Context {
	return context.WithValue(ctx, ctxKeySession, session)
}

func Session(ctx context.Context) domain.Session {
	session, ok := ctx.Value(ctxKeySession).(domain.Session)
	if !ok {
		return domain.NoSession
	}
	return session
}

func MustSession(ctx context.Context) domain.Session {
	session, ok := ctx.Value(ctxKeySession).(domain.Session)
	if !ok {
		slog.Error("session not found in context", slog.Any("ctx", ctx))
		panic("session not found in context")
	}

	return session
}
