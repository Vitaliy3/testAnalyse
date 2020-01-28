// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dbStruct

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type graphsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *graphsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("graphs").
func (v *graphsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *graphsTableType) Columns() []string {
	return []string{"id", "fk_graph", "fk_test_result"}
}

// NewStruct makes a new struct for that view or table.
func (v *graphsTableType) NewStruct() reform.Struct {
	return new(Graphs)
}

// NewRecord makes a new record for that table.
func (v *graphsTableType) NewRecord() reform.Record {
	return new(Graphs)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *graphsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// GraphsTable represents graphs view or table in SQL database.
var GraphsTable = &graphsTableType{
	s: parse.StructInfo{Type: "Graphs", SQLSchema: "", SQLName: "graphs", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "FkGraph", Type: "[]uint8", Column: "fk_graph"}, {Name: "FkTestResult", Type: "*int32", Column: "fk_test_result"}}, PKFieldIndex: 0},
	z: new(Graphs).Values(),
}

// String returns a string representation of this struct or record.
func (s Graphs) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "FkGraph: " + reform.Inspect(s.FkGraph, true)
	res[2] = "FkTestResult: " + reform.Inspect(s.FkTestResult, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Graphs) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.FkGraph,
		s.FkTestResult,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Graphs) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.FkGraph,
		&s.FkTestResult,
	}
}

// View returns View object for that struct.
func (s *Graphs) View() reform.View {
	return GraphsTable
}

// Table returns Table object for that record.
func (s *Graphs) Table() reform.Table {
	return GraphsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Graphs) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Graphs) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Graphs) HasPK() bool {
	return s.ID != GraphsTable.z[GraphsTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Graphs) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = GraphsTable
	_ reform.Struct = (*Graphs)(nil)
	_ reform.Table  = GraphsTable
	_ reform.Record = (*Graphs)(nil)
	_ fmt.Stringer  = (*Graphs)(nil)
)

func init() {
	parse.AssertUpToDate(&GraphsTable.s, new(Graphs))
}