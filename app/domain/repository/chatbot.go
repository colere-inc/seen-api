package repository

import "github.com/colere-inc/seen-api/app/domain/model"

type ChatbotRepository interface {
	GetChatbotAnswers(spaceID string, surveyID string) (*model.ChatbotAnswers, error)
}
