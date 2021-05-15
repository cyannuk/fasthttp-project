// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type reviewTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *reviewTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("reviews").
func (v *reviewTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *reviewTableType) Columns() []string {
	return []string{
		"id",
		"created_at",
		"reviewer",
		"product_id",
		"rating",
		"body",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *reviewTableType) NewStruct() reform.Struct {
	return new(Review)
}

// NewRecord makes a new record for that table.
func (v *reviewTableType) NewRecord() reform.Record {
	return new(Review)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *reviewTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ReviewTable represents reviews view or table in SQL database.
var ReviewTable = &reviewTableType{
	s: parse.StructInfo{
		Type:    "Review",
		SQLName: "reviews",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "Reviewer", Type: "string", Column: "reviewer"},
			{Name: "ProductID", Type: "int64", Column: "product_id"},
			{Name: "Rating", Type: "int32", Column: "rating"},
			{Name: "Body", Type: "string", Column: "body"},
		},
		PKFieldIndex: 0,
	},
	z: new(Review).Values(),
}

// String returns a string representation of this struct or record.
func (s Review) String() string {
	res := make([]string, 6)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[2] = "Reviewer: " + reform.Inspect(s.Reviewer, true)
	res[3] = "ProductID: " + reform.Inspect(s.ProductID, true)
	res[4] = "Rating: " + reform.Inspect(s.Rating, true)
	res[5] = "Body: " + reform.Inspect(s.Body, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Review) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.CreatedAt,
		s.Reviewer,
		s.ProductID,
		s.Rating,
		s.Body,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Review) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.CreatedAt,
		&s.Reviewer,
		&s.ProductID,
		&s.Rating,
		&s.Body,
	}
}

// View returns View object for that struct.
func (s *Review) View() reform.View {
	return ReviewTable
}

// Table returns Table object for that record.
func (s *Review) Table() reform.Table {
	return ReviewTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Review) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Review) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Review) HasPK() bool {
	return s.ID != ReviewTable.z[ReviewTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Review) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ReviewTable
	_ reform.Struct = (*Review)(nil)
	_ reform.Table  = ReviewTable
	_ reform.Record = (*Review)(nil)
	_ fmt.Stringer  = (*Review)(nil)
)

func init() {
	parse.AssertUpToDate(&ReviewTable.s, new(Review))
}
