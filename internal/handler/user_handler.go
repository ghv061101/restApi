package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/ghv061101/RestApiAge/internal/models"
	"github.com/ghv061101/RestApiAge/internal/service"
)

type UserHandler struct {
	svc      *service.UserService
	validate *validator.Validate
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc, validate: validator.New()}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := struct {
		Name string `json:"name" validate:"required,min=2"`
		Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "invalid request"})
	}
	if err := h.validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "validation failed", "error": err.Error()})
	}
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)
	u := &models.Users{Name: req.Name, Dob: parsedDob}
	if err := h.svc.Create(u); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not create user"})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"id": u.ID, "name": u.Name, "dob": u.Dob.Format("2006-01-02")})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.svc.List()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not get users"})
	}
	type ur struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Dob  string `json:"dob"`
		Age  int    `json:"age"`
	}
	resp := make([]ur, 0, len(users))
	for _, u := range users {
		resp = append(resp, ur{ID: u.ID, Name: u.Name, Dob: u.Dob.Format("2006-01-02"), Age: u.Age()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": resp})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	u, err := h.svc.GetByID(uint(id64))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": fiber.Map{"id": u.ID, "name": u.Name, "dob": u.Dob.Format("2006-01-02"), "age": u.Age()}})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	existing, err := h.svc.GetByID(uint(id64))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	req := struct {
		Name string `json:"name" validate:"required,min=2"`
		Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "invalid request"})
	}
	if err := h.validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "validation failed", "error": err.Error()})
	}
	parsedDob, _ := time.Parse("2006-01-02", req.Dob)
	existing.Name = req.Name
	existing.Dob = parsedDob
	if err := h.svc.Update(existing); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not update user"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"id": existing.ID, "name": existing.Name, "dob": existing.Dob.Format("2006-01-02")})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid id"})
	}
	rows, err := h.svc.Delete(uint(id64))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not delete user"})
	}
	if rows == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	return c.SendStatus(http.StatusNoContent)
}
