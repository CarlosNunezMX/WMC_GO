package controllers

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/CarlosNunezMX/WMC_GO/interfaces"
	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {
	folderPath := c.Query("path")
	fmt.Println(folderPath)
	if folderPath == "" {
		error_http := interfaces.Error{
			Cause: "Not path in query",
		}
		return c.Status(fiber.StatusBadRequest).JSON(error_http)
	}
	place := path.Join("./", folderPath)
	files, err := os.Open(place)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return c.Status(fiber.ErrBadRequest.Code).JSON(interfaces.Error{
			Cause: "Folder isn't exit's",
		})
	}

	filesInFolder, err_readDir := files.ReadDir(0)
	if err_readDir != nil {
		panic(err)
	}
	var res []interfaces.Meida
	for _, file := range filesInFolder {
		info, err := file.Info()
		if err != nil {
			panic(err)
		}
		var tipo string
		if !info.IsDir() {
			tipo = interfaces.Video
			if !strings.Contains(file.Name(), ".mp4") {
				println("The file: ", file.Name(), " isn't a video")
				continue
			}
		} else {
			tipo = interfaces.Folder
		}
		res = append(res, interfaces.Meida{
			OnlineMedia: false,
			MediaType:   tipo,
			Url:         folderPath + "/" + file.Name(),
			Name:        file.Name(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func Video(c *fiber.Ctx) error {
	place := c.Query("path")
	if place == "" {
		error_http := interfaces.Error{
			Cause: "Not path in query",
		}
		return c.Status(fiber.StatusBadRequest).JSON(error_http)
	}
	if !strings.Contains(place, ".mp4") {
		return c.Status(fiber.ErrBadRequest.Code).JSON(interfaces.Error{
			Cause: "File isn't a Video",
		})
	}
	exits := check(place)
	fmt.Printf("exits: %v\n", exits)
	if !exits {
		return c.Status(fiber.ErrBadRequest.Code).JSON(interfaces.Error{
			Cause: "File isn't exit's",
		})
	}

	return c.SendFile(path.Join("./", place), true)
}

func check(place string) bool {
	_, err := os.Stat(path.Join("./", place))
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
