package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
	)
	userDetail, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user detail")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return response, errors.Wrap(err, "failed to compare password")
	}
	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := models.UserSession{
		UserID:               userDetail.ID,
		Token:                token,
		RefreshToken:         refreshToken,
		TokenExpired:         now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired:  now.Add(helpers.MapTypeToken["refresh_token"]),
	}
	if err := s.UserRepo.InsertNewUserSession(ctx, &userSession); err != nil {
		return response, errors.Wrap(err, "failed to insert new user session")
	}

	response.UserID = userDetail.ID
	response.Username = userDetail.Username
	response.FullName = userDetail.FullName
	response.Email = userDetail.Email
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}
