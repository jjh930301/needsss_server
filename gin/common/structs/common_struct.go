package structs

import "github.com/dgrijalva/jwt-go"

type AuthClaim struct {
	UserID string `json:"id"`    // 유저 ID
	Type   int    `json:"type"`  // 유저 타입
	Email  string `json:"email"` // 유저 메일
	jwt.StandardClaims
}
