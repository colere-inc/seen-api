package adapter

import (
	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/colere-inc/seen-api/app/domain/repository/mock_repository"
)

func TestChatbotController_Get(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/spaces/:spaceID/surveys/:surveyID/chatbot")
	c.SetParamNames("spaceID", "surveyID")
	c.SetParamValues("testSpaceID", "testSurveyID")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock_repository.NewMockChatbotRepository(ctrl)
	mock.EXPECT().GetChatbotAnswers("testSpaceID", "testSurveyID").Return(&model.ChatbotAnswers{
		Answers: []model.ChatbotAnswer{
			{
				QuestionID:      1,
				Question:        "Test Question 1 Text",
				Answer:          "Test Answer 1 Text",
				NextQuestionIDs: []int32{2, 3, 4},
			},
			{
				QuestionID:      2,
				Question:        "Test Question 2 Text",
				Answer:          "Test Answer 2 Text",
				NextQuestionIDs: []int32{3, 4, 5},
			},
		},
	}, nil)
	chatbotController := NewChatbotController(mock)

	// Act
	h := chatbotController.Get()
	err := h(c)
	if err != nil {
		t.Fatal(err)
	}
}
