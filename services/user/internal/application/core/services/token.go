package services

import (
	"log"
	"strconv"
	"time"

	"github.com/Yusufdot101/ribble/services/user/config"
	"github.com/Yusufdot101/ribble/services/user/internal/application/core/domain"
	"github.com/Yusufdot101/ribble/services/user/internal/ports"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenService struct {
	repo ports.Repository
}

func NewTokenService(repo ports.Repository) *TokenService {
	return &TokenService{
		repo: repo,
	}
}

// New() (*domain.Token, error)
// Save(token *domain.Token) error

func (tsvc *TokenService) New(tokenType domain.TokenType, tokenUse domain.TokenUse, userID uint) (*domain.Token, error) {
	var ttl time.Duration
	switch tokenUse {
	case domain.REFRESH:
		ttl = config.GetRefreshTokenTTL()
	case domain.ACCESS:
		ttl = config.GetAccessTokenTTL()
	default:
		return nil, domain.ErrInvalidTokenUse
	}

	token := &domain.Token{
		TokenType: tokenType,
		UserID:    userID,
		Expires:   time.Now().Add(ttl),
		Use:       tokenUse,
	}
	var tokenString string
	switch tokenType {
	case domain.RANDOMSTRING:
		tokenString = generateRandomString()
	case domain.JWT:
		tokenString = genenerateJWT(userID)
	default:
		return nil, domain.ErrInvalidTokeType
	}
	token.TokenString = tokenString
	return token, nil
}

func generateRandomString() string {
	return uuid.New().String()
}

func genenerateJWT(userID uint) string {
	claims := jwt.RegisteredClaims{
		Issuer:    config.GetJWTIssuer(),
		Subject:   strconv.Itoa(int(userID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.GetAccessTokenTTL())),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetJWTSecret())
	if err != nil {
		log.Fatalf("error signing jwt: %v\n", err)
	}
	return tokenString
}

func (tsvc *TokenService) Save(token *domain.Token) error {
	return tsvc.repo.InsertToken(token)
}
