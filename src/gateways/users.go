package gateways

import (
	"go-fiber-template/src/domain/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *HTTPGateway) GetAllUserData(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h *HTTPGateway) CreateNewUserAccount(ctx *fiber.Ctx) error {

	bodyData := new(entities.NewUserBody)
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	err := h.UserService.InsertNewAccount(bodyData)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseMessage{Message: "cannot insert new user data"})
	}
	log.Info("new user created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(entities.ResponseModel{Message: "created successfully"})
}

func (h *HTTPGateway) UpdateUserData(ctx *fiber.Ctx) error {

	bodyData := new(entities.NewUserBody)
	if err := ctx.BodyParser(&bodyData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}
	params := ctx.Queries()
	userID := params["user_id"]

	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid query params"})
	}

	err := h.UserService.UpdateUser(userID, bodyData)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseMessage{Message: "cannot update user data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "updated successfully"})
}

func (h *HTTPGateway) DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("user_id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid query params"})
	}
	err := h.UserService.DeleteUser(userID)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseMessage{Message: "cannot delete user data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "deleted successfully"})
}

func (h *HTTPGateway) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(entities.ResponseMessage{Message: "invalid query params"})
	}
	data, err := h.UserService.GetUser(userID)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseMessage{Message: "cannot get user data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}
