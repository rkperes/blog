package domain

type User struct {
	ID string
}

type Session struct {
	ID     string
	UserID string
}
