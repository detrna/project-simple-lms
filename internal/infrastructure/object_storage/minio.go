package infrastructure

import (
	"bytes"
	"context"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/pkg"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	client *minio.Client
}

func SetupMinIO(config *config.ObjectStorageConfig) (pkg.ObjectStorage, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL == "true",
	})

	if err != nil {
		return nil, err
	}

	return &MinIO{client: client}, nil
}

func (m *MinIO) UploadFile(
	ctx context.Context,
	bucketName, objectName, contentType string,
	file []byte,
) (*domain.File, error) {
	reader := bytes.NewReader(file)

	info, err := m.client.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		reader.Size(),
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)

	if err != nil {
		return nil, err
	}

	id := uuid.New()
	return &domain.File{
		ID:          id,
		Name:        objectName,
		FileURL:     "http://" + m.client.EndpointURL().Host + "/" + bucketName + "/" + objectName,
		Size:        info.Size,
		ContentType: contentType,
		Bucket:      bucketName,
	}, nil
}
