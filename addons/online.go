package addons

import (
	"log"
	"net/http"
	"strings"

	"github.com/CarlosNunezMX/WMC_GO/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/khatibomar/kobayashi"
)

func FembedToAPI(id string) string {
	url := "https://www.fembed.com/api/source/" + id
	return url
}

var OnlineTempAnime []*interfaces.TempFile

func GetSource(c *fiber.Ctx) error {
	url := c.Query("video_url")
	media_type := c.Query("type")

	if media_type == "Fembed" {
		x := strings.Split(url, "/v/")[1]
		if x == "" {
			return interfaces.ErrorHandling(c, interfaces.Error{Cause: "Incomplete Fembed URL!"}, *fiber.ErrBadRequest)
		}

		url = FembedToAPI(x)
	}
	if url == "" {
		return interfaces.ErrorHandling(c, interfaces.Error{Cause: "Give a url in query!"}, *fiber.ErrBadRequest)
	}
	d := kobayashi.NewDecoder()
	println(url)
	link, err := d.Decode(url)
	if err != nil {
		log.Printf(err.Error())
		return interfaces.ErrorHandling(c, interfaces.Error{Cause: err.Error()}, *fiber.ErrConflict)
	}
	return c.Status(200).JSON(interfaces.Meida{
		OnlineMedia: true,
		Url:         link,
		MediaType:   interfaces.Video,
		Name:        "Video online",
	})
}

func TempStoreOnlineMedia(c *fiber.Ctx) error {
	f := new(interfaces.TempFile)

	if err := c.BodyParser(f); err != nil {
		return interfaces.ErrorHandling(c, interfaces.Error{
			Cause: "All values is required!",
		}, *fiber.ErrBadRequest)
	}
	if f.Name == "" || f.Url == "" {
		return interfaces.ErrorHandling(c, interfaces.Error{
			Cause: "All fields is required!",
		}, *fiber.ErrBadRequest)
	}

	OnlineTempAnime = append(OnlineTempAnime, f)
	return c.JSON(interfaces.StandarMessage{
		Message: "Complete!",
	})
}

func GetTemps(c *fiber.Ctx) error {
	return c.JSON(OnlineTempAnime)
}

func Proxy(c *fiber.Ctx) error {
	Url := c.Query("url")
	println(Url)
	if Url == "" {
		return interfaces.ErrorHandling(c, interfaces.Error{
			Cause: "Url is required!",
		}, *fiber.ErrBadRequest)
	}

	req, _ := http.NewRequest(http.MethodGet, Url, nil)

	req.Header.Set("Authorization", "Bearer ---------")
	req.Header.Set("Client-ID", "---------")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Set("Content-Type", "application/json; charset=utf-8")

	return c.Status(res.StatusCode).SendStream(res.Body)
}
