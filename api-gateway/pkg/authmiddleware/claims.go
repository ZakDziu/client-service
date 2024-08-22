package authmiddleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
