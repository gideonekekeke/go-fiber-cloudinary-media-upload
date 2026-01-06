package handlers

import (
	"context"
	"fiber-file-upload/config"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func UploadSingle(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(400, "File is required")
	}

	// Open file stream
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload to Cloudinary using stream
	result, err := config.Cloudinary.Upload.Upload(
		context.Background(),
		file, // ✅ io.Reader (NOT file path)
		uploader.UploadParams{
			Folder:       "fiber_uploads",
			ResourceType: "auto",
		},
	)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(fiber.Map{
		"url":      result.SecureURL,
		"publicId": result.PublicID,
		"type":     result.ResourceType,
	})
}

func UploadMultiple(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	results := []fiber.Map{}

	for _, file := range files {
		tempPath := filepath.Join(os.TempDir(), file.Filename)
		c.SaveFile(file, tempPath)

		res, err := config.Cloudinary.Upload.Upload(
			context.Background(),
			tempPath,
			uploader.UploadParams{
				Folder: "fiber_uploads",
			},
		)

		os.Remove(tempPath)

		if err == nil {
			results = append(results, fiber.Map{
				"url": res.SecureURL,
			})
		}
	}

	return c.JSON(results)
}
