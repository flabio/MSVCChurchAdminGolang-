package handler

import (
	"msvc_function_ministerial/usecase/services"

	"github.com/gofiber/fiber/v2"
)

type functionMinisterialHandler struct {
	functionMinisterial services.IFunctionMinisterialService
}

func NewFunctionMinisterialHandler() IFunctionMinisterialHandler {

	return &functionMinisterialHandler{
		functionMinisterial: services.NewFunctionMinisterialService(),
	}
}

func (r *functionMinisterialHandler) GetFindAll(c *fiber.Ctx) error {
	return r.functionMinisterial.GetFindAll(c)

}

func (r *functionMinisterialHandler) GetFindById(c *fiber.Ctx) error {
	return r.functionMinisterial.GetFindById(c)

}

func (r *functionMinisterialHandler) Create(c *fiber.Ctx) error {
	return r.functionMinisterial.Create(c)
}
func (r *functionMinisterialHandler) Update(c *fiber.Ctx) error {
	return r.functionMinisterial.Update(c)
}

func (r *functionMinisterialHandler) Delete(c *fiber.Ctx) error {
	return r.functionMinisterial.Delete(c)
}
