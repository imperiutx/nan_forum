package gapi

import (
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		UserName:          user.UserName,
		Email:             user.Email,
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
