package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	db "github.com/IskanderA1/handly/iternal/db/sqlc"
	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/repository"
	"github.com/IskanderA1/handly/pkg/config"
	passwordHash "github.com/IskanderA1/handly/pkg/hash"
	"github.com/IskanderA1/handly/pkg/token"
)

type AdminsService struct {
	adminRepository      repository.Admins
	sessionRepository    repository.Sessions
	tokenManger          token.Maker
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewAdminsService(adminRepository repository.Admins, sessionRepository repository.Sessions, tokenManger token.Maker, config config.Config) *AdminsService {
	return &AdminsService{
		tokenManger:          tokenManger,
		adminRepository:      adminRepository,
		sessionRepository:    sessionRepository,
		accessTokenDuration:  config.AccessTokenDuration,
		refreshTokenDuration: config.RefreshTokenDuration,
	}
}

func (s *AdminsService) SignIn(ctx context.Context, input AdminSingInInput, adminConfig AdminConfig) (domain.Session, error) {

	admin, err := s.adminRepository.GetByUsername(ctx, input.Username)
	if err != nil {
		return domain.Session{}, fmt.Errorf("User not found")
	}

	if err := passwordHash.CheckPassword(input.Password, admin.Password); err != nil {
		return domain.Session{}, fmt.Errorf("Invalid login or password")
	}
	return s.createSession(ctx, admin, adminConfig)
}

func (s *AdminsService) SignUp(ctx context.Context, input AdminSignUpInput, adminConfig AdminConfig) (domain.Admin, error) {
	hashedPassword, err := passwordHash.HashPassword(input.Password)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("Invalid password")
	}
	admin, err := s.adminRepository.Create(ctx, db.CreateAdminParams{
		Username: input.Username,
		FullName: input.FullName,
		Password: hashedPassword,
	})

	if err != nil {
		return domain.Admin{}, err
	}

	return domain.NewAdmin(admin), err
}

func (s *AdminsService) RefreshToken(ctx context.Context, refreshToken string) (domain.Session, error) {

	refreshPayload, err := s.tokenManger.VerifyAdminToken(refreshToken)
	if err != nil {
		return domain.Session{}, fmt.Errorf("invalid refresh token")
	}

	session, err := s.sessionRepository.GetById(ctx, refreshPayload.ID)
	if err != nil {
		return domain.Session{}, err
	}

	if session.IsBlocked {
		return domain.Session{}, fmt.Errorf("blocked session")
	}

	if session.Username != refreshPayload.Username {
		return domain.Session{}, fmt.Errorf("incorrect session user")
	}

	if session.RefreshToken != refreshToken {
		return domain.Session{}, fmt.Errorf("mismatched session token")
	}

	if time.Now().After(session.ExpiresAt) {
		return domain.Session{}, fmt.Errorf("expired session")
	}

	accessToken, accessPayload, err := s.tokenManger.CreateAdminToken(
		refreshPayload.Username,
		s.accessTokenDuration,
	)
	if err != nil {
		return domain.Session{}, fmt.Errorf("Failed to create accessToken")
	}

	return domain.Session{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}, err
}

func (s *AdminsService) GetList(ctx context.Context, input ListInput) ([]domain.Admin, error) {
	admins, err := s.adminRepository.GetList(ctx, db.ListAdminsParams{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return make([]domain.Admin, 0), err
	}
	adminsWithoutPass := make([]domain.Admin, 0)
	for _, admin := range admins {
		adminsWithoutPass = append(adminsWithoutPass, domain.NewAdmin(admin))
	}
	return adminsWithoutPass, err
}

func (s *AdminsService) GetByName(ctx context.Context, username string) (domain.Admin, error) {
	admin, err := s.adminRepository.GetByUsername(ctx, username)
	if err != nil {
		return domain.Admin{}, fmt.Errorf("admin not found")
	}
	return domain.NewAdmin(admin), err
}

func (s *AdminsService) Delete(ctx context.Context, username string) error {
	err := s.adminRepository.Delete(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("admin not found")
		}
		return fmt.Errorf("unable to remove admin")
	}
	return nil
}

func (s *AdminsService) createSession(ctx context.Context, admin db.Admin, adminConfig AdminConfig) (domain.Session, error) {

	accessToken, accessPayload, err := s.tokenManger.CreateAdminToken(admin.Username, s.accessTokenDuration)
	if err != nil {
		return domain.Session{}, err
	}

	refreshToken, refreshPayload, err := s.tokenManger.CreateAdminToken(admin.Username, s.refreshTokenDuration)
	if err != nil {
		return domain.Session{}, err
	}

	_, err = s.sessionRepository.Create(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     admin.Username,
		RefreshToken: refreshToken,
		UserAgent:    adminConfig.UserAgent,
		ClientIp:     adminConfig.ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	return domain.Session{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}, err
}
