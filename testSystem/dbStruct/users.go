package dbStruct

//go:generate reform

// Users represents a row in users table.
//reform:users

type Users struct {
	ID         int32   `reform:"id,pk"`
	UserName   string  `reform:"user_name"`
	Surname    string  `reform:"surname"`
	Patronymic *string `reform:"patronymic"`
	UserPass   string  `reform:"user_pass"`
	FkRole     int32   `reform:"fk_role"`
}
