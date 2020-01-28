package dbStruct

//go:generate reform

// TestUsersSubject represents a row in test_users_subject table.
//reform:test_users_subject
type TestUsersSubject struct {
	ID          int32 `reform:"id,pk"`
	FkTestUsers int32 `reform:"fk_test_users"`
	FkSubject   int32 `reform:"fk_subject"`
}
