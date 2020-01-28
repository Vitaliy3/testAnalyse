package dbStruct

//go:generate reform

// Subjects represents a row in subjects table.
//reform:subjects
type Subjects struct {
	ID          int32  `reform:"id,pk"`
	SubjectName string `reform:"subject_name"`
}
