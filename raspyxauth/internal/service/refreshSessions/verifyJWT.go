package refreshSessions

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func (s *RefreshSessionsService) VerifyJWT(tokenStr string) (*jwt.MapClaims, error) {
	tokenStr = strings.TrimPrefix(strings.TrimSpace(tokenStr), "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte(s.cfg.JWT.JWTSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	//exp, err := claims.GetExpirationTime()
	//if err != nil {
	//	return fmt.Errorf("error get expiration time: %w", err)
	//}
	//
	//if exp.Time.Unix()+60 < time.Now().Unix() {
	//	return fmt.Errorf("token expired")
	//}

	return &claims, nil
}
