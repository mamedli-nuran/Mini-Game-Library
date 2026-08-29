package service

import "time"

type UserRepo interface {
}

type UserService struct {
	repo          UserRepo
	jwtSecret     string
	accessExpire  time.Duration
	refreshExpire time.Duration
}

func NewUserService(repo UserRepo, jwtSecret string, accessExpire, refreshExpire time.Duration) *UserService {
	return &UserService{
		repo:          repo,
		jwtSecret:     jwtSecret,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}
