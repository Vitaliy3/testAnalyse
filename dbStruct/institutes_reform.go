// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dbStruct

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type institutesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *institutesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("institutes").
func (v *institutesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *institutesTableType) Columns() []string {
	return []string{"id", "institute_name", "short_name"}
}

// NewStruct makes a new struct for that view or table.
func (v *institutesTableType) NewStruct() reform.Struct {
	return new(Institutes)
}

// NewRecord makes a new record for that table.
func (v *institutesTableType) NewRecord() reform.Record {
	return new(Institutes)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *institutesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// InstitutesTable represents institutes view or table in SQL database.
var InstitutesTable = &institutesTableType{
	s: parse.StructInfo{Type: "Institutes", SQLSchema: "", SQLName: "institutes", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "InstituteName", Type: "string", Column: "institute_name"}, {Name: "ShortName", Type: "string", Column: "short_name"}}, PKFieldIndex: 0},
	z: new(Institutes).Values(),
}

// String returns a string representation of this struct or record.
func (s Institutes) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "InstituteName: " + reform.Inspect(s.InstituteName, true)
	res[2] = "ShortName: " + reform.Inspect(s.ShortName, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Institutes) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.InstituteName,
		s.ShortName,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Institutes) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.InstituteName,
		&s.ShortName,
	}
}

// View returns View object for that struct.
func (s *Institutes) View() reform.View {
	return InstitutesTable
}

// Table returns Table object for that record.
func (s *Institutes) Table() reform.Table {
	return InstitutesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Institutes) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Institutes) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Institutes) HasPK() bool {
	return s.ID != InstitutesTable.z[InstitutesTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Institutes) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = InstitutesTable
	_ reform.Struct = (*Institutes)(nil)
	_ reform.Table  = InstitutesTable
	_ reform.Record = (*Institutes)(nil)
	_ fmt.Stringer  = (*Institutes)(nil)
)

func init() {
	parse.AssertUpToDate(&InstitutesTable.s, new(Institutes))
}
