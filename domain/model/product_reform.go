// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type productTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *productTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("products").
func (v *productTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *productTableType) Columns() []string {
	return []string{
		"id",
		"created_at",
		"category",
		"ean",
		"price",
		"quantity",
		"rating",
		"title",
		"vendor",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *productTableType) NewStruct() reform.Struct {
	return new(Product)
}

// NewRecord makes a new record for that table.
func (v *productTableType) NewRecord() reform.Record {
	return new(Product)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *productTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProductTable represents products view or table in SQL database.
var ProductTable = &productTableType{
	s: parse.StructInfo{
		Type:    "Product",
		SQLName: "products",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "Category", Type: "string", Column: "category"},
			{Name: "Ean", Type: "string", Column: "ean"},
			{Name: "Price", Type: "float64", Column: "price"},
			{Name: "Quantity", Type: "int32", Column: "quantity"},
			{Name: "Rating", Type: "float64", Column: "rating"},
			{Name: "Title", Type: "string", Column: "title"},
			{Name: "Vendor", Type: "string", Column: "vendor"},
		},
		PKFieldIndex: 0,
	},
	z: new(Product).Values(),
}

// String returns a string representation of this struct or record.
func (s Product) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[2] = "Category: " + reform.Inspect(s.Category, true)
	res[3] = "Ean: " + reform.Inspect(s.Ean, true)
	res[4] = "Price: " + reform.Inspect(s.Price, true)
	res[5] = "Quantity: " + reform.Inspect(s.Quantity, true)
	res[6] = "Rating: " + reform.Inspect(s.Rating, true)
	res[7] = "Title: " + reform.Inspect(s.Title, true)
	res[8] = "Vendor: " + reform.Inspect(s.Vendor, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Product) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.CreatedAt,
		s.Category,
		s.Ean,
		s.Price,
		s.Quantity,
		s.Rating,
		s.Title,
		s.Vendor,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Product) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.CreatedAt,
		&s.Category,
		&s.Ean,
		&s.Price,
		&s.Quantity,
		&s.Rating,
		&s.Title,
		&s.Vendor,
	}
}

// View returns View object for that struct.
func (s *Product) View() reform.View {
	return ProductTable
}

// Table returns Table object for that record.
func (s *Product) Table() reform.Table {
	return ProductTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Product) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Product) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Product) HasPK() bool {
	return s.ID != ProductTable.z[ProductTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Product) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ProductTable
	_ reform.Struct = (*Product)(nil)
	_ reform.Table  = ProductTable
	_ reform.Record = (*Product)(nil)
	_ fmt.Stringer  = (*Product)(nil)
)

func init() {
	parse.AssertUpToDate(&ProductTable.s, new(Product))
}
