package config

type ObjectStorageConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
}

func LoadObjectStorageConfig() *ObjectStorageConfig {
	endpoint := GetEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKey := GetEnv("MINIO_ACCESS_KEY", "minioadmin")
	secretKey := GetEnv("MINIO_SECRET_KEY", "minioadmin")

	return &ObjectStorageConfig{
		Endpoint:  endpoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}
