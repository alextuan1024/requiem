package web

import (
	"errors"
	"github.com/alextuan1024/requiem/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type loginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type loginResp struct {
	Status int      `json:"status"`
	Data   jwtToken `json:"data"`
}

func (l *loginReq) GetUser() string {
	return l.Username
}

type Profiler interface {
	GetUser() string
}

type jwtToken struct {
	Token string `json:"token"`
}

const EnvAuthSign = "REQUIEM_AUTH_SIGH"

func LoginHandler(c *gin.Context) {
	var req loginReq
	if err := c.Bind(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token, err := generateToken(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &loginResp{
		Data: jwtToken{
			Token: token,
		},
	})

}

func generateToken(p Profiler) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  p.GetUser(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString([]byte(settings.Current.AuthSign))
}

func AuthMiddleware(c *gin.Context) {
	token, err := getTokenFromReq(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if _, err := parseToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
}

func getTokenFromReq(c *gin.Context) (string, error) {
	var token string
	token = c.GetHeader("Authorization")
	if token != "" {
		if strings.Index(token, "Bearer ") < 0 {
			return "", ErrInvalidToken
		}
		return strings.TrimPrefix(token, "Bearer "), nil
	}
	token = c.Query("auth_token")
	if token != "" {
		return token, nil
	}

	token, _ = c.Cookie("auth_token")
	if token != "" {
		return token, nil
	}
	return "", ErrInvalidToken
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(settings.Current.AuthSign), nil
	})
	if err != nil {
		if !errors.Is(err, ErrInvalidToken) {
			logger().Info(err)
		}
		return nil, ErrInvalidToken
	}
	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
