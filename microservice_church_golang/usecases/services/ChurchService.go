package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
	"microservice_church.com/core/interfaces"
	"microservice_church.com/core/repositories"
	"microservice_church.com/infrastructure/entities"
	"microservice_church.com/infrastructure/utils"

	"microservice_church.com/usecases/dto"
)

type churchService struct {
	IChurch interfaces.IChurch
}

func NewChurchService() IChurchService {
	return &churchService{
		IChurch: repositories.GetChurchInstance(),
	}
}

func (s *churchService) GetChurchFindAll(c *fiber.Ctx) error {
	result, err := s.IChurch.GetChurchFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (s *churchService) GetChurchFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.IChurch.GetChurchFindById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(result)
}
func (s *churchService) CreateChurch(c *fiber.Ctx) error {
	var churchCreate entities.Church
	var churchDto dto.ChurchDTO
	church, msgError := validateChurch(0, churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(church).To(&churchCreate)
	result, err := s.IChurch.CreateChurch(churchCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}

func (s *churchService) UpdateChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var churchUpdate entities.Church
	var churchDto dto.ChurchDTO

	findById, _ := s.IChurch.GetChurchFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	church, msgError := validateChurch(uint(id), churchDto, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(church).To(&churchUpdate)
	result, err := s.IChurch.UpdateChurch(uint(id), churchUpdate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}

func (s *churchService) DeleteChurch(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))

	findById, _ := s.IChurch.GetChurchFindById(uint(id))
	if findById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.IChurch.DeleteChurch(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func validateChurch(id uint, churchDto dto.ChurchDTO, s *churchService, c *fiber.Ctx) (dto.ChurchDTO, string) {
	var msg string = ""
	b := c.Body()
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)
	if errJson != nil {
		msg = errJson.Error()
	}

	msgValid := validateField(dataMap)
	if msgValid != "" {
		return dto.ChurchDTO{}, msgValid
	}

	MapToStruct(&churchDto, dataMap)
	msgRequired := validateRequired(churchDto)
	if msgRequired != "" {
		return dto.ChurchDTO{}, msgRequired
	}
	log.Println(churchDto)
	existName, _ := s.IChurch.GetChurchFindByName(id, churchDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	existEmail, _ := s.IChurch.GetChurchFindByEmail(id, churchDto.Email)
	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return churchDto, msg
}

func MapToStruct(dataDto *dto.ChurchDTO, dataMap map[string]interface{}) {
	rol := dto.ChurchDTO{
		Name:    dataMap["name"].(string),
		Email:   dataMap["email"].(string),
		Address: dataMap["address"].(string),
		Phone:   dataMap["phone"].(string),
		Active:  dataMap["active"].(bool),
	}
	*dataDto = rol
}
func validateField(value map[string]interface{}) string {
	var msg string = ""
	if value["name"] == nil {
		msg = "The field name is required"
	}
	if value["email"] == nil {
		msg = "The field email is required"
	}
	if value["address"] == nil {
		msg = "The field address is required"
	}
	if value["phone"] == nil {
		msg = "The field phone is required"
	}
	// if value["active"] == nil {
	// 	msg = "The field active is required"
	// }
	return msg
}

func validateRequired(field dto.ChurchDTO) string {
	var msg string = ""
	if field.Name == "" {
		msg = "The name is required"
	}
	if field.Email == "" {
		msg = "The email is required"
	}
	if field.Address == "" {
		msg = "The address is required"
	}
	if field.Phone == "" {
		msg = "The phone is required"
	}
	if field.Active == false {
		msg = "The active field is required"
	}

	return msg
}
