package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/User_model"
	services "github.com/suhas-developer07/GuessVibe-Server/internals/service/user"
)

type UserHandler struct {
	UserServices *services.UserService
}

func (h *UserHandler) UserRegisterHandler(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "Invalid request")
	}
	id, err := h.UserServices.RegisterUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status: "error",
			Error:  "Failed to register user: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    map[string]int64{"user_id": id},
	})

}
