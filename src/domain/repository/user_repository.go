package repository

type IUserRepository interface {
	Save(lineUserID string) uint
	FindOne(lineUserID string) uint
}
