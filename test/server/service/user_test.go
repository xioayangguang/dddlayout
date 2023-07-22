package service_test

import (
	"context"
	"errors"
	"fmt"
	mock_repository "layout/mocks/repository"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"layout/internal/middleware"
	"layout/internal/model"
	"layout/internal/service"
	"layout/pkg/configParse"

	"layout/pkg/log"
)

var (
	srv *service.Service
)

func TestMain(m *testing.M) {
	fmt.Println("begin")

	os.Setenv("APP_CONF", "../../../config/local.yml")

	conf := configParse.NewConfig()

	logger := log.NewLog(conf)
	jwt := middleware.NewJwt(conf)
	sf := sid.NewSid()
	srv = service.NewService(logger, sf, jwt)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &service.RegisterRequest{
		Username: "testuser",
		Password: "password",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByUsername(ctx, req.Username).Return(nil, nil)
	mockUserRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	err := userService.Register(ctx, req)

	assert.NoError(t, err)
}

func TestUserService_Register_UsernameExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &service.RegisterRequest{
		Username: "testuser",
		Password: "password",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByUsername(ctx, req.Username).Return(&model.User{}, nil)

	err := userService.Register(ctx, req)

	assert.Error(t, err)
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &service.LoginRequest{
		Username: "testuser",
		Password: "password",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		t.Error("failed to hash password")
	}

	mockUserRepo.EXPECT().GetByUsername(ctx, req.Username).Return(&model.User{
		Password: string(hashedPassword),
	}, nil)

	token, err := userService.Login(ctx, req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	req := &service.LoginRequest{
		Username: "testuser",
		Password: "password",
	}

	mockUserRepo.EXPECT().GetByUsername(ctx, req.Username).Return(nil, errors.New("user not found"))

	_, err := userService.Login(ctx, req)

	assert.Error(t, err)
}

func TestUserService_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(&model.User{
		UserId:   userId,
		Username: "testuser",
		Email:    "test@example.com",
	}, nil)

	user, err := userService.GetProfile(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, userId, user.UserId)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"
	req := &service.UpdateProfileRequest{
		Nickname: "testuser",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(&model.User{
		UserId:   userId,
		Username: "testuser",
		Email:    "old@example.com",
	}, nil)
	mockUserRepo.EXPECT().Update(ctx, gomock.Any()).Return(nil)

	err := userService.UpdateProfile(ctx, userId, req)

	assert.NoError(t, err)
}

func TestUserService_UpdateProfile_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := "123"
	req := &service.UpdateProfileRequest{
		Nickname: "testuser",
		Email:    "test@example.com",
	}

	mockUserRepo.EXPECT().GetByID(ctx, userId).Return(nil, errors.New("user not found"))

	err := userService.UpdateProfile(ctx, userId, req)

	assert.Error(t, err)
}

func TestUserService_GenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

	userService := service.NewUserService(srv, mockUserRepo)

	ctx := context.Background()
	userId := 123

	result, err := userService.GenerateToken(ctx, userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
