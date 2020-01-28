package dbStruct

//go:generate reform

// Tickets represents a row in tickets table.
//reform:tickets
type Tickets struct {
	ID     int32 `reform:"id,pk"`
	Number int32 `reform:"number"`
}
