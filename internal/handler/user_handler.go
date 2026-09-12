package handler

import (
	"context"
	"encoding/json"
	"errors"
	"mini-game-library/internal/apperror"
	"mini-game-library/internal/config"
	"mini-game-library/internal/constant"
	"mini-game-library/internal/dto"
	"mini-game-library/internal/models"
	"mini-game-library/internal/service"
	"net/http"
	"time"
)

type UserService interface {
	RegisterUser(ctx context.Context, request dto.RegisterRequest) (*models.User, error)
	LoginUser(ctx context.Context, request dto.LoginRequest) (*service.TokenPair, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*service.TokenPair, error)
	GetMeInfo(ctx context.Context) (*models.User, error)
}
type UserHandler struct {
	svc UserService
	cfg config.Config
}

func NewUserHandler(svc UserService, cfg config.Config) *UserHandler {
	return &UserHandler{
		svc: svc,
		cfg: cfg,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}

	req.Sanitize()
	if errorDetails := req.Validate(); len(errorDetails) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errorDetails...)
		return
	}

	user, err := h.svc.RegisterUser(r.Context(), req)

	if err != nil {
		if errors.Is(err, apperror.ErrUserDuplicate) {
			WriteError(w, r, http.StatusConflict, constant.ErrUserAlreadyExists)
		} else {
			WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		}
		return
	}
	writeJSON(w, http.StatusCreated, dto.NewRegisterResponse(user))

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, r, http.StatusBadRequest, constant.ErrInvalidBody)
		return
	}

	req.Sanitize()
	if errorDetails := req.Validate(); len(errorDetails) > 0 {
		WriteError(w, r, http.StatusBadRequest, constant.ErrValidationFailed, errorDetails...)
		return
	}

	tokens, err := h.svc.LoginUser(r.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrUserFind) || errors.Is(err, apperror.ErrInvalidCredentials) {
			WriteError(w, r, http.StatusBadRequest, err.Error())
		} else {
			WriteError(w, r, http.StatusInternalServerError, apperror.ErrInternalServerError)
		}
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Expires:  time.Now().Add(h.cfg.RefreshTokenExpireHours),
	}
	http.SetCookie(w, cookie)

	writeJSON(w, http.StatusOK, dto.TokenResponse{
		AccessToken: tokens.AccessToken,
	})
}

func (h *UserHandler) MeInfo(w http.ResponseWriter, r *http.Request) {
	user, err := h.svc.GetMeInfo(r.Context())
	if err != nil {
		if errors.Is(err, apperror.ErrUnauthorized) {
			WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		} else if errors.Is(err, apperror.ErrUserFind) {
			WriteError(w, r, http.StatusNotFound, "user not found")
		} else {
			WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		}
		return
	}

	userResponse := dto.NewUserResponse(user)
	writeJSON(w, http.StatusOK, userResponse)
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			WriteError(w, r, http.StatusUnauthorized, "missing refresh_token cookie")
		} else {
			WriteError(w, r, http.StatusBadRequest, "error reading refresh_token cookie")
		}
		return
	}

	refreshToken := cookie.Value
	if refreshToken == "" {
		WriteError(w, r, http.StatusUnauthorized, "refresh_token is empty")
		return
	}

	tokens, err := h.svc.RefreshTokens(r.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, apperror.ErrUnauthorized) {
			WriteError(w, r, http.StatusUnauthorized, constant.ErrUnauthorized)
		} else {
			WriteError(w, r, http.StatusInternalServerError, constant.ErrInternalServerError)
		}
		return
	}

	newCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Expires:  time.Now().Add(h.cfg.RefreshTokenExpireHours),
	}
	http.SetCookie(w, newCookie)

	writeJSON(w, http.StatusOK, dto.TokenResponse{
		AccessToken: tokens.AccessToken,
	})
}
