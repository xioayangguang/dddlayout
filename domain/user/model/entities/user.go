package entities

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"layout/domain/user/model/entities/pdo"
	"layout/infrastructure/berror"
	"layout/infrastructure/http/response"
	"layout/infrastructure/redis"
	"layout/pkg/contextValue"
	"layout/pkg/idBuilder"
	"layout/pkg/md5"
	"strconv"
	"time"
)

type User struct {
	Id             uint64
	CreatedAt      uint
	UpdatedAt      uint
	DeletedAt      uint
	Uuid           uint
	Serial         uint
	Nickname       string
	Mail           string
	Describe       string
	Code           string
	InvitationCode string
	Avatar         string
	Status         int
}

type UserService interface {
	Login(ctx context.Context, req *pdo.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userId uint64) (*User, error)
	UpdateProfile(ctx context.Context, userId uint64, req *pdo.UpdateProfileRequest) error
	GenerateToken(ctx context.Context, userInfo *User) string
}

type userService struct {
	userRepo facade.UserRepository
	*entities.Service
}

func NewUserService(service *entities.Service, userRepo facade.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

// Login 登录
func (s *userService) Login(ctx context.Context, req *pdo.LoginRequest) (string, error) {
	userModel := &User{}
	var err error
	if s.transaction(ctx, func(ctx context.Context) error {
		userModel, err = s.userRepo.GetByUsername(ctx, req.Nickname)
		if err != nil {
			if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
				userModel.Serial = uint(idBuilder.Generate("user_id_id_builder", func() int {
					return s.userRepo.GetMaxSerial(ctx)
				}))
				userModel.InvitationCode = idBuilder.Id2Code(int(userModel.Serial))
				userModel.Uuid = idBuilder.From32To10(userModel.InvitationCode)
				userModel.Nickname = "SAG_" + strconv.Itoa(int(userModel.Uuid))
				//userModel.Nickname = req.Nickname
				return s.userRepo.Create(ctx, userModel)
			}
			return berror.New(response.LoginError)
		} else {
			return nil
		}
	}) != nil {
		return "", berror.New(response.LoginError)
	}
	return s.GenerateToken(ctx, userModel), nil
}

// GetProfile 获取用户信息
func (s *userService) GetProfile(ctx context.Context, userId uint64) (*User, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, berror.New(response.Error)
	}
	return user, nil
}

// UpdateProfile 修改用户信息
func (s *userService) UpdateProfile(ctx context.Context, userId uint64, req *pdo.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return berror.New(response.Error)
	}
	user.Mail = req.Email
	user.Nickname = req.Nickname
	if err = s.userRepo.Update(ctx, user); err != nil {
		return berror.New(response.Error)
	}
	return nil
}

// GenerateToken 生成用户token
func (s *userService) GenerateToken(ctx context.Context, userInfo *User) string {
	channel := "app"                        //此处演示写死
	var duration time.Duration = 86400 * 30 //此处演示写死
	token := md5.Md5(strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(int(userInfo.Id)))
	jsonStr, _ := json.Marshal(contextValue.LoginUserInfo{
		Id:             userInfo.Id,
		Nickname:       userInfo.Nickname,
		Uuid:           userInfo.Uuid,
		InvitationCode: userInfo.InvitationCode,
		ApiAuth:        token,
		Serial:         userInfo.Serial,
	})
	strUserId := strconv.FormatUint(userInfo.Id, 10)
	if oldToken, _ := redis.Instances.HGet(ctx, channel, strUserId).Result(); oldToken != "" {
		redis.Instances.Del(ctx, oldToken)
	}
	redis.Instances.Set(ctx, token, jsonStr, duration*time.Second)
	redis.Instances.HSet(ctx, channel, strUserId, token)
	return token
}
