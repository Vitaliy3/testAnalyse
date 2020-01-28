package dbStruct

//go:generate reform

// Specialities represents a row in specialities table.
//reform:specialities
type Specialities struct {
	ID             int32  `reform:"id,pk"`
	SpecialityName string `reform:"speciality_name"`
	FkInstitute    int32  `reform:"fk_institute"`
}
