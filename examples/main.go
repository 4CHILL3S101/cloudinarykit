package main

import (
	"context"
	"fmt"
	"log"

	"github.com/4CHILL3S101/cloudinarykit"
)

func main() {
	ctx := context.Background()

	ck, err := cloudinarykit.New(cloudinarykit.Config{
		CloudName: "your_cloud_name",
		APIKey:    "your_api_key",
		APISecret: "your_api_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload file
	url, err := ck.UploadFile(ctx, "test.jpg", "myfolder", "test", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded URL:", url)

	// Delete file
	err = ck.Delete(ctx, "myfolder/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted successfully")
}
