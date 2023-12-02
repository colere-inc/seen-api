package adapter

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, chatbotController ChatbotController) {
	e.GET("/api/v1/spaces/:spaceID/surveys/:surveyID/chatbot", chatbotController.Get())
}
