package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var secret string = "l85XKbW9lDc8LyYhZPGG_lNEJHTDZa3SK2ioD3Phm"

//生成token
func GenerateToken(user string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userId,
		"username": user,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
		//"iat": time.Now().Unix(),//用作和exp对比的时间
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return tokenString, err
	} else {
		return tokenString, nil //对应的字符串请自行生成，最后足够使用加密后的字符串
	}
}

//验证token是否有效
func VerifyToken(m map[string]string, tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		m["id"] = claims["id"].(string)
		m["username"] = claims["username"].(string)
		return true, nil
	} else {
		return false, err
	}
}
