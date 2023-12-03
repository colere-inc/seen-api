package model

import (
	"context"
	"fmt"
	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/colere-inc/seen-api/app/infrastructure"
	"google.golang.org/api/iterator"
)

const maxAnswers = 100

type ChatbotRepository struct {
	db *infrastructure.DB
}

func NewChatbotRepository(db *infrastructure.DB) repository.ChatbotRepository {
	return ChatbotRepository{db: db}
}

func (c ChatbotRepository) GetChatbotAnswers(
	spaceID string,
	surveyID string,
) (*model.ChatbotAnswers, error) {
	ctx := context.Background()
	answers := make([]model.ChatbotAnswer, 0, maxAnswers)
	iter := c.db.Collection(fmt.Sprintf("spaces/%s/surveys/%s/chatbot", spaceID, surveyID)).Limit(maxAnswers).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var answer model.ChatbotAnswer
		if err := doc.DataTo(&answer); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return &model.ChatbotAnswers{
		Answers: answers,
	}, nil
}
