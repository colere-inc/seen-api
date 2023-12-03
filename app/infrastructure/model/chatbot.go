package model

import (
	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/colere-inc/seen-api/app/infrastructure"
)

type ChatbotRepository struct {
	db *infrastructure.DB
}

func NewChatbotRepository(db *infrastructure.DB) repository.ChatbotRepository {
	return ChatbotRepository{db: db}
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
