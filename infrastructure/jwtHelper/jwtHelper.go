package jwthelper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("oX7XqVVSttlw3oy7tZbv")

// Claims 定义了 JWT 的自定义声明
type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	Address  string `json:"address"`
	Role     string `json:"role"` // 添加角色字段用于权限控制
	jwt.StandardClaims
}

func GenerateToken(userID uint, username string, address string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		Address:  address,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "LearnGin",
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("无效的token")
	}

	return claims, nil
}

// ValidateToken 验证JWT令牌是否有效
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
