package model

import (
	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
)

type ChatbotRepository struct{}

func NewChatbotRepository() repository.ChatbotRepository {
	return ChatbotRepository{}
}

func (c ChatbotRepository) GetChatbotAnswers(spaceID string, surveyID string) (*model.ChatbotAnswers, error) {
	return &model.ChatbotAnswers{
		Answers: []model.ChatbotAnswer{
			{
				QuestionID:      1,
				Question:        "question text",
				Answer:          "answer  text",
				NextQuestionIDs: []int32{2, 3, 4},
			},
		},
	}, nil
}
