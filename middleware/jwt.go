package middleware

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
	"zeropiece/common"
	"zeropiece/structData"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := Decrypt(c.Request.Header.Get("Authorization"))

		if token == "" {
			ResponseFail(c, 401, "请求未携带token，无权限访问")
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token[len("Bearer "):])
		if err != nil {

			if err == TokenExpired {
				ResponseFail(c, 401, "token 已经过期，请重新登录！")
				c.Abort()
				return
			}
			ResponseFail(c, 401, err.Error())
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	TokenExpireAt    int64
)

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims structData.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*structData.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structData.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*structData.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &structData.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*structData.Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix() // 默认token有效期为1个小时
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// GenerateToken 生成令牌
func GenerateToken(userClaims *structData.Claims) (token string, msg string, ok bool) {
	j := &JWT{[]byte(common.CONFIG.Jwt.SigningKey)}
	res, err := strconv.ParseInt(common.CONFIG.Jwt.Expire, 10, 64)
	if err != nil {
		common.LOG.Error("jwt expire config info err")
	}
	TokenExpireAt = res
	userClaims.StandardClaims = jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,          // 签名生效时间
		ExpiresAt: time.Now().Unix() + TokenExpireAt, // 过期时间 一小时
		Issuer:    common.CONFIG.Jwt.SigningKey,      //签名的发行者
	}
	token, err = j.CreateToken(*userClaims)
	if err != nil {
		common.LOG.Error("创建Token失败", zap.Any("err", err))
		return token, "创建token失败", false
	} else {
		bearer := "Bearer "
		token = fmt.Sprintf("%s%s", bearer, token)
		return Encrypt(token), "登录成功", true
	}
}

func GetClaims(c *gin.Context) *structData.Claims {
	claims, exist := c.Get("claims")
	if !exist {
		return &structData.Claims{}
	}
	return claims.(*structData.Claims)
}

func Encrypt(message string) (STR string) {
	// 处理Base64编码里的等号
	data := base64.StdEncoding.EncodeToString([]byte(message))
	re := regexp.MustCompile(`=+$`)
	return re.ReplaceAllString(data, "")
}

func Decrypt(message string) (STR string) {
	// 因为字符串必须是4的倍数 所以用等号补位
	num := len(message) % 4
	if num != 0 {
		num = 4 - len(message)%4
	}
	message = message + strings.Repeat("=", num)
	strByte, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println(err)
	}
	return string(strByte)
}
