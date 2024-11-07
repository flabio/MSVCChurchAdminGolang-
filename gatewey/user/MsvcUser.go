package user

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

func EnvLoad() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading.env file")
	}
}

// Handler for the  service
func MsvcUser(c *fiber.Ctx) error {
	EnvLoad()
	MSVC_ROL_URL := os.Getenv("MSVC_USER_URL")
	id := c.Params(utils.ID)
	url := MSVC_ROL_URL
	if len(id) != 0 && url != "" {
		url = MSVC_ROL_URL + "/" + id
	}
	req, err := http.NewRequest(c.Method(), url, bytes.NewBuffer(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FAILED_CREATE)
	}
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.SERVICE_NOT_AVAILALE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return c.Send(respBody)
}
