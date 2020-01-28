package dbStruct

//go:generate reform

// Questions represents a row in questions table.
//reform:questions
type Questions struct {
	ID         int32  `reform:"id,pk"`
	Question   string `reform:"question"`
	TrueAnswer string `reform:"true_answer"`
	FkTicket   int32  `reform:"fk_ticket"`
}
