package dbStruct

//go:generate reform

// TestResults represents a row in test_results table.
//reform:test_results
type TestResults struct {
	ID       int32  `reform:"id,pk"`
	TestInfo string `reform:"test_info"`
	FkUser   int32  `reform:"fk_user"`
}
