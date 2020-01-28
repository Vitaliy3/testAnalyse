package dbStruct

//go:generate reform

// Roles represents a row in roles table.
//reform:roles
type Roles struct {
	ID   int32  `reform:"id,pk"`
	Role string `reform:"role"`
}
