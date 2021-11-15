package repository

import "context"

type IUserRepository interface {
	Save(ctx context.Context, lineUserID string) uint
	FindOne(ctx context.Context, lineUserID string) uint
}
