package model

type ChatbotAnswers struct {
	Answers []ChatbotAnswer `json:"answers"`
}

type ChatbotAnswer struct {
	QuestionID      int32   `json:"question_id"`
	Question        string  `json:"question"`
	Answer          string  `json:"answer"`
	NextQuestionIDs []int32 `json:"next_question_ids"`
}
