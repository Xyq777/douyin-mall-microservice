package domain

import "context"

type Usecase interface {
	DeliverToken(ctx context.Context, userID int32) (tokenStr string, err error)
	RenewToken(ctx context.Context, oldToken string) (tokenStr string, err error)
	VerifyToken(ctx context.Context, token string) (isValid bool, err error)
}
