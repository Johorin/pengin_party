package firebase

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FirebaseInterface interface {
	GetFirebase() *firebase.App
}

type FirebaseApp struct {
	app *firebase.App
}

func Init() (FirebaseInterface, error) {
	ctx := context.Background()
	
	// 環境変数からサービスアカウントキーのパスを取得
	credentialPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	
	if credentialPath == "" {
		return nil, fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS environment variable is not set")
	}
	
	// 相対パスの場合、現在の作業ディレクトリからの相対パスとして解決
	// 絶対パスの場合はそのまま使用
	if !filepath.IsAbs(credentialPath) {
		// 作業ディレクトリを取得（通常は backend/ ディレクトリ）
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory: %v", err)
		}
		fmt.Printf("Loading Firebase credentials from: %s\n", wd)
		credentialPath = filepath.Join(wd, credentialPath)
	}
	
	// credentialファイルの存在確認
	if _, err := os.Stat(credentialPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("credentials file not found: %s (working directory: %s)", credentialPath, os.Getenv("PWD"))
	}
	
	fmt.Printf("Loading Firebase credentials from: %s\n", credentialPath)

	opt := option.WithCredentialsFile(credentialPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase admin: %v", err)
	}

	return &FirebaseApp{
		app: app,
	}, nil
}

func (fc *FirebaseApp) GetFirebase() *firebase.App {
	return fc.app
}