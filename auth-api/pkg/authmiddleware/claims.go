package authmiddleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

const (
	AccessTokenTTL  = time.Hour * 8
	RefreshTokenTTL = time.Hour * 24 * 7
)

type BaseClaims struct {
	jwt.StandardClaims
}

type AccessClaims struct {
	BaseClaims
	AccessUUID string `json:"access_uuid"`
}

type RefreshClaims struct {
	BaseClaims
	RefreshUUID string `json:"refresh_uuid"`
}

type Tokens struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

func NewClaims(ttl time.Duration) BaseClaims {
	return BaseClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Id:        uuid.NewV4().String(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func GenerateClaims() (*AccessClaims, *RefreshClaims) {
	access := AccessClaims{
		BaseClaims: NewClaims(AccessTokenTTL),
	}

	refresh := RefreshClaims{
		BaseClaims: NewClaims(RefreshTokenTTL),
	}

	access.AccessUUID = refresh.Id
	refresh.RefreshUUID = access.Id

	return &access, &refresh
}
