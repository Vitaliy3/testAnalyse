package dbStruct

//go:generate reform

// Institutes represents a row in institutes table.
//reform:institutes
type Institutes struct {
	ID            int32  `reform:"id,pk"`
	InstituteName string `reform:"institute_name"`
	ShortName     string `reform:"short_name"`
}
