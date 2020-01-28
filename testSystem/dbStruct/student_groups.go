package dbStruct

//go:generate reform

// StudentGroups represents a row in student_groups table.
//reform:student_groups
type StudentGroups struct {
	ID           int32  `reform:"id,pk"`
	GroupName    string `reform:"group_name"`
	FkSpeciality int32  `reform:"fk_speciality"`
}
