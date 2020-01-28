package dbStruct

import (
	"time"
)

type Tests struct {
	ID             int32     `reform:"id,pk"`
	Duration       int32     `reform:"duration"`
	Mark           int32     `reform:"mark"`
	DateCompletion time.Time `reform:"date_completion"`
	FkSubject      int32     `reform:"fk_subject"`
	FkTicket       int32     `reform:"fk_ticket"`
}

//go:generate reform

// Tests represents a row in tests table.
//reform:tests
