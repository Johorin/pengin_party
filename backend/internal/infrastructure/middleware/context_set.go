package middleware

import (
	"fmt"

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
		idToken := c.Request.Header.Get("Authorization")
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