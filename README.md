# CloudinaryKit

**CloudinaryKit** is a lightweight and reusable Go library for integrating [Cloudinary](https://cloudinary.com/) into your Go applications. It provides an easy way to **upload and manage images and videos**, supporting both **file paths** and **in-memory byte slices**. Designed for **simplicity, reusability, and secure public API usage**.

---

## Features

- Upload **images from file** (`UploadFile`)  
- Upload **images from bytes** (`UploadBytes`) — ideal for APIs or web/mobile uploads  
- Upload **videos from file** (`UploadVideoFile`)  
- Upload **videos from bytes** (`UploadVideoBytes`)  
- Automatic **MIME type detection** for byte uploads  
- Supports **overwrite** for existing resources  
- Delete uploaded resources by **public ID**  
- Fully compatible with **Go contexts**

---

## Installation

```bash
go get github.com/4CHILL3S101/cloudinarykit



Usage Example

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/4CHILL3S101/cloudinarykit"
)

func main() {
	ctx := context.Background()

	// Initialize CloudinaryKit with your credentials
	ck, err := cloudinarykit.New(cloudinarykit.Config{
		CloudName: "your_cloud_name",
		APIKey:    "your_api_key",
		APISecret: "your_api_secret",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload an image file
	url, err := ck.UploadFile(ctx, "test.jpg", "images", "myimage", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded URL:", url)

	// Delete the uploaded image
	err = ck.Delete(ctx, "images/myimage")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted successfully")
}



Use Cases
Backend services handling user-generated content
APIs that accept file or FormData uploads
Web or mobile apps needing secure image/video storage
Projects that want quick Cloudinary integration in Go
License

MIT License © 4CHILL3S101

