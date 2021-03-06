// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dbStruct

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type specialitiesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *specialitiesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("specialities").
func (v *specialitiesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *specialitiesTableType) Columns() []string {
	return []string{"id", "speciality_name", "fk_institute"}
}

// NewStruct makes a new struct for that view or table.
func (v *specialitiesTableType) NewStruct() reform.Struct {
	return new(Specialities)
}

// NewRecord makes a new record for that table.
func (v *specialitiesTableType) NewRecord() reform.Record {
	return new(Specialities)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *specialitiesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// SpecialitiesTable represents specialities view or table in SQL database.
var SpecialitiesTable = &specialitiesTableType{
	s: parse.StructInfo{Type: "Specialities", SQLSchema: "", SQLName: "specialities", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "SpecialityName", Type: "string", Column: "speciality_name"}, {Name: "FkInstitute", Type: "int32", Column: "fk_institute"}}, PKFieldIndex: 0},
	z: new(Specialities).Values(),
}

// String returns a string representation of this struct or record.
func (s Specialities) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "SpecialityName: " + reform.Inspect(s.SpecialityName, true)
	res[2] = "FkInstitute: " + reform.Inspect(s.FkInstitute, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Specialities) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.SpecialityName,
		s.FkInstitute,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Specialities) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.SpecialityName,
		&s.FkInstitute,
	}
}

// View returns View object for that struct.
func (s *Specialities) View() reform.View {
	return SpecialitiesTable
}

// Table returns Table object for that record.
func (s *Specialities) Table() reform.Table {
	return SpecialitiesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Specialities) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Specialities) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Specialities) HasPK() bool {
	return s.ID != SpecialitiesTable.z[SpecialitiesTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Specialities) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = SpecialitiesTable
	_ reform.Struct = (*Specialities)(nil)
	_ reform.Table  = SpecialitiesTable
	_ reform.Record = (*Specialities)(nil)
	_ fmt.Stringer  = (*Specialities)(nil)
)

func init() {
	parse.AssertUpToDate(&SpecialitiesTable.s, new(Specialities))
}
