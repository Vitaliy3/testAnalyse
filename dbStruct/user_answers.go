package dbStruct

//go:generate reform

// UserAnswers represents a row in user_answers table.
//reform:user_answers
type UserAnswers struct {
	ID         int32  `reform:"id,pk"`
	Answer     string `reform:"answer"`
	FkQuestion int32  `reform:"fk_question"`
}
