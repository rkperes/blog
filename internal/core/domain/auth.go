package domain

type UUID string

type User struct {
	ID UUID
}

type Session struct {
	ID     UUID
	UserID UUID
}
