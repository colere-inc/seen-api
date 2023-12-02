package adapter

import (
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChatbotController struct {
	chatbotRepository repository.ChatbotRepository
}

func NewChatbotController(chatbotRepository repository.ChatbotRepository) *ChatbotController {
	return &ChatbotController{
		chatbotRepository: chatbotRepository,
	}
}

func (ctrl *ChatbotController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		spaceID := c.Param("spaceID")
		surveyID := c.Param("surveyID")
		answers, err := ctrl.chatbotRepository.GetChatbotAnswers(spaceID, surveyID)
		if err != nil {
			panic(any(err))
		}
		return c.JSON(http.StatusOK, answers)
	}
}
