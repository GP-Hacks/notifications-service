package grpc

import (
	"context"

	notification_service "github.com/GP-Hacks/notifications/internal/services/notifications_service"
	desc "github.com/GP-Hacks/proto/pkg/api/notifications"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TokensController struct {
	desc.UnimplementedNotificationsServer
	notificationsService *notification_service.NotificationsService
}

func NewTokensController(ns *notification_service.NotificationsService) *TokensController {
	return &TokensController{
		notificationsService: ns,
	}
}

func (tc *TokensController) AddUserToken(ctx context.Context, in *desc.AddUserTokenRequest) (*emptypb.Empty, error) {
	err := tc.notificationsService.AddUserToken(ctx, in.GetUserId(), in.GetToken())

	return &emptypb.Empty{}, err
}
