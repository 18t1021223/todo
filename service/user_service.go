package service

import (
	"context"
	"database/sql"
	db "todo_list_roadmap/db/genarated"
	"todo_list_roadmap/dto"
	"todo_list_roadmap/handle/response"
	"todo_list_roadmap/util"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	sqlDB *sql.DB
}

func (s *UserService) Create(ctx context.Context, req dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	session := db.New(s.sqlDB)
	ID := uuid.New().String()
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		zap.L().Error("Failed GenerateFromPassword", zap.Error(err))
		return dto.UserRegisterResponse{}, response.ErrInternalServer
	}
	if err := session.CreateUser(
		ctx,
		db.CreateUserParams{
			ID:       ID,
			Name:     req.Name,
			Email:    req.Email,
			Password: pwdHash,
		},
	); err != nil {
		if util.IsDuplicate(err) {
			return dto.UserRegisterResponse{}, response.ErrEmailDuplicate
		}
		zap.L().Error("Failed CreateUser", zap.Error(err))
		return dto.UserRegisterResponse{}, response.ErrInternalServer
	}
	token, err := CreateToken(&db.User{ID: ID, Name: req.Name, Email: req.Email})
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}
	return dto.UserRegisterResponse{Token: token}, nil
}

func (s *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	session := db.New(s.sqlDB)
	user, err := session.GetUserByEmail(ctx, req.Email)
	if err != nil {
		zap.L().Error("Failed GetUserByEmail", zap.Error(err))
		return dto.UserLoginResponse{}, response.ErrUserNotFound
	}
	token, err := CreateToken(&user)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	return dto.UserLoginResponse{Token: token}, nil
}

func NewUserService(sqlDB *sql.DB) *UserService {
	return &UserService{sqlDB: sqlDB}
}
