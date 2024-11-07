package ministerialmember

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

func EnvLoad() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func MsvcMinisterialMember(c *fiber.Ctx) error {
	EnvLoad()
	MSVC_MINISTERIAL_URL := os.Getenv("MSVC_MINISTERIALs_URL")
	id := c.Params(utils.ID)
	url := MSVC_MINISTERIAL_URL

	if len(id) != 0 && url != "" {
		url = MSVC_MINISTERIAL_URL + "/" + id
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

func MsvcMinisterialUserMinisterial(c *fiber.Ctx) error {
	EnvLoad()
	MSVC_MINISTERIAL_USER_URL := os.Getenv("MSVC_MINISTERIAL_MEMBER_URL")
	id := c.Params(utils.ID)
	url := MSVC_MINISTERIAL_USER_URL

	if len(id) != 0 && url != "" {
		url = MSVC_MINISTERIAL_USER_URL + "/" + id
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
