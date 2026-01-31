package middleware

import (
	"fmt"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

const (
	UserUID = "user_uid"
)

func ContextSetMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		client, err := app.Auth(ctx)
		if err != nil {
			panic("error getting Auth client")
		}

		// リクエストヘッダのAuthorizationからFirebaseのIDトークンを取得し、Firebase Admin SDKで検証
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Bearerプレフィックスを除去
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader {
			c.JSON(401, gin.H{"error": "Invalid Authorization header format. Expected 'Bearer <token>'"})
			c.Abort()
			return
		}
		
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			panic("error verifying ID token")
		}
		fmt.Printf("Verified ID token: %v\n", token)
		
		// 正しく認証されていればコンテキストにセット
		c.Set(UserUID, token.UID)
		c.Next()
	}
}