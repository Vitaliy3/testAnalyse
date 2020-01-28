package dbStruct

//go:generate reform

// Graphs represents a row in graphs table.
//reform:graphs
type Graphs struct {
	ID           int32  `reform:"id,pk"`
	FkGraph      []byte `reform:"fk_graph"`
	FkTestResult *int32 `reform:"fk_test_result"`
}
