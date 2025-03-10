package delivery

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/internal/domain"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/pkg/log"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc-gen/auth"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuthDelivery)

type AuthDelivery struct {
	usecase domain.Usecase
}

func NewAuthDelivery(usecase domain.Usecase) *AuthDelivery {
	return &AuthDelivery{usecase}
}
func (d *AuthDelivery) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	var (
		token string
	)
	if req.UserId == 0 {
		return nil, errors.New("no userID")
	}
	token, err = d.usecase.DeliverToken(ctx, req.UserId)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	return &auth.DeliveryResp{
		Token: token,
	}, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (d *AuthDelivery) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	var (
		isValid bool
	)
	isValid, err = d.usecase.VerifyToken(ctx, req.Token)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid token")
	}

	return &auth.VerifyResp{
		Res: isValid,
	}, err
}

// RenewTokenByRPC implements the AuthServiceImpl interface.
func (d *AuthDelivery) RenewTokenByRPC(ctx context.Context, req *auth.RenewTokenReq) (resp *auth.DeliveryResp, err error) {
	var (
		newToken string
	)

	newToken, err = d.usecase.RenewToken(ctx, req.OldToken)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	return &auth.DeliveryResp{
		Token: newToken,
	}, err
}
