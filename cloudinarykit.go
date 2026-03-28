package cloudinarykit

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Config struct {
	CloudName string
	APIKey    string
	APISecret string
}

type CloudinaryKit struct {
	CLD *cloudinary.Cloudinary
}

func New(cfg Config) (*CloudinaryKit, error) {
	cloudinaryURL := fmt.Sprintf("cloudinary://%s:%s@%s", cfg.APIKey, cfg.APISecret, cfg.CloudName)
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %w", err)
	}
	return &CloudinaryKit{CLD: cld}, nil
}

func (c *CloudinaryKit) UploadFile(ctx context.Context, filePath string, folder string, publicID string, overwrite bool) (string, error) {
	resp, err := c.CLD.Upload.Upload(ctx, filePath, uploader.UploadParams{
		Folder:    folder,
		PublicID:  publicID,
		Overwrite: &overwrite,
	})
	if err != nil {
		return "", fmt.Errorf("upload file failed: %w", err)
	}
	return resp.SecureURL, nil
}

func (c *CloudinaryKit) UploadBytes(ctx context.Context, data []byte, folder string, fileName string) (string, error) {
	mimeType := http.DetectContentType(data[:512])

	base64Str := base64.StdEncoding.EncodeToString(data)
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)
	overwrite := true
	resp, err := c.CLD.Upload.Upload(ctx, dataURI, uploader.UploadParams{
		Folder:    folder,
		PublicID:  fileName,
		Overwrite: &overwrite,
	})
	if err != nil {
		return "", fmt.Errorf("upload from bytes failed: %w", err)
	}
	return resp.SecureURL, nil
}

func (c *CloudinaryKit) UploadVideoFile(ctx context.Context, filePath, folder, publicID string, overwrite bool) (string, error) {
	resp, err := c.CLD.Upload.Upload(ctx, filePath, uploader.UploadParams{
		Folder:       folder,
		PublicID:     publicID,
		Overwrite:    &overwrite,
		ResourceType: "video",
	})
	if err != nil {
		return "", fmt.Errorf("upload video file failed: %w", err)
	}
	return resp.SecureURL, nil
}

// Delete removes an image from Cloudinary using its public ID
func (c *CloudinaryKit) Delete(ctx context.Context, publicID string) error {
	_, err := c.CLD.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}
	return nil
}
