package dbStruct

//go:generate reform

// TestUsers represents a row in test_users table.
//reform:test_users

type TestUsers struct {
	ID       int32  `reform:"id,pk"`
	UserName string `reform:"user_name"`
	FkGroup  int32  `reform:"fk_group"`
}
