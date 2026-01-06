//for using a simple upload just like using multer in express

// package normal_file

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func normalFile() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("We are live....")
// 	})

// 	app.Post("/upload", func(c *fiber.Ctx) error {
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}

// 		// Individual file size limit
// 		maxFileSize := int64(5 * 1024 * 1024) // 5MB per file
// 		if file.Size > maxFileSize {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "File size exceeds limit (5MB)",
// 			})
// 		}

// 		// Save file
// 		err = c.SaveFile(file, "./uploads/"+file.Filename)
// 		if err != nil {
// 			return err
// 		}

// 		return c.JSON(fiber.Map{
// 			"message":  "uploaded sucessful",
// 			"filename": file.Filename,
// 			"size":     file.Size,
// 		})
// 	})

// 	app.Listen(":8000")

// }
