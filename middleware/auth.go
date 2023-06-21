package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		data, err := ctx.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				if ctx.GetHeader("Content-Type") == "application/json" {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				} else {
					ctx.Redirect(http.StatusSeeOther, "/user/login")
					ctx.Abort()
				}
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(data, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
