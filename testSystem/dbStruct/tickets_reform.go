// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dbStruct

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type ticketsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *ticketsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("tickets").
func (v *ticketsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *ticketsTableType) Columns() []string {
	return []string{"id", "number"}
}

// NewStruct makes a new struct for that view or table.
func (v *ticketsTableType) NewStruct() reform.Struct {
	return new(Tickets)
}

// NewRecord makes a new record for that table.
func (v *ticketsTableType) NewRecord() reform.Record {
	return new(Tickets)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *ticketsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TicketsTable represents tickets view or table in SQL database.
var TicketsTable = &ticketsTableType{
	s: parse.StructInfo{Type: "Tickets", SQLSchema: "", SQLName: "tickets", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Number", Type: "int32", Column: "number"}}, PKFieldIndex: 0},
	z: new(Tickets).Values(),
}

// String returns a string representation of this struct or record.
func (s Tickets) String() string {
	res := make([]string, 2)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Number: " + reform.Inspect(s.Number, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Tickets) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Number,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Tickets) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Number,
	}
}

// View returns View object for that struct.
func (s *Tickets) View() reform.View {
	return TicketsTable
}

// Table returns Table object for that record.
func (s *Tickets) Table() reform.Table {
	return TicketsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Tickets) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Tickets) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Tickets) HasPK() bool {
	return s.ID != TicketsTable.z[TicketsTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Tickets) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = TicketsTable
	_ reform.Struct = (*Tickets)(nil)
	_ reform.Table  = TicketsTable
	_ reform.Record = (*Tickets)(nil)
	_ fmt.Stringer  = (*Tickets)(nil)
)

func init() {
	parse.AssertUpToDate(&TicketsTable.s, new(Tickets))
}
