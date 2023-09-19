package jwt

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	JWT_SECRET = "JWT_SECRET"
)

type JwtCustomInfo struct {
	Infos map[string]string `json:"info,omitempty"`
}

func (j *JwtCustomInfo) ToString() (string, error) {
	jo, err := json.Marshal(j.Infos)
	if err != nil {
		return "", err
	}
	return string(jo), nil
}

// 生成一个token
func GenerateToken(infos map[string]string, expire_hours int64) (string, error) {
	jwtInfo := JwtCustomInfo{Infos: infos}
	infoString, err := jwtInfo.ToString()

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * time.Duration(expire_hours)).Unix(), // TTL: 3 小时
		"iat":   time.Now().Unix(),
		"infos": infoString,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv(JWT_SECRET)))

	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(infos map[string]string, expire_hours int64) (string, error) {
	jwtInfo := JwtCustomInfo{Infos: infos}

	infoString, err := jwtInfo.ToString()

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * time.Duration(expire_hours)).Unix(),
		"infos": infoString,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv(JWT_SECRET)))

	if err != nil {
		return "", err
	}
	return t, nil
}

func GenrateTokenPair(infos map[string]string, expire_hours int64) (map[string]string, error) {
	accessToken, err := GenerateToken(infos, expire_hours)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateRefreshToken(infos, expire_hours)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func GetTokenFromHeader(header string) (*jwt.Token, error) {
	const BearerSchema string = "Bearer "

	if header == "" {
		return nil, fmt.Errorf("empty authorization header")

	}
	tokenString := header[len(BearerSchema):]
	token, err := ValidateToken(tokenString)
	return token, err
}

// func ParseCustomInfoFromToken(token *jwt.Token) (map[string]string) {

// }

func ParseCustomInfosFromToken(token *jwt.Token) (*JwtCustomInfo, error) {
	var infos map[string]string

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		infoString := claims["infos"].(string)
		err := json.Unmarshal([]byte(infoString), &infos)

		if err != nil {
			return nil, err
		}
	}

	return &JwtCustomInfo{Infos: infos}, nil
}
