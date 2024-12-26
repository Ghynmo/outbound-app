package config

import (
	"e-commerce-1/helper"
	"fmt"
)

type FirebaseConfig struct {
	CredentialFile string
	BucketName     string
	ProjectID      string
	StorageURL     string
}

func NewFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		CredentialFile: helper.GetEnv("FIREBASE_CREDENTIAL_FILE", "firebase-credentials.json"),
		BucketName:     helper.GetEnv("FIREBASE_BUCKET_NAME", "your-bucket.appspot.com"),
		ProjectID:      helper.GetEnv("FIREBASE_PROJECT_ID", "your-project-id"),
		StorageURL:     helper.GetEnv("FIREBASE_STORAGE_URL", ""),
	}
}

func (c FirebaseConfig) Validate() error {
	if c.CredentialFile == "" {
		return fmt.Errorf("firebase credential file is required")
	}
	if c.BucketName == "" {
		return fmt.Errorf("firebase bucket name is required")
	}
	return nil
}