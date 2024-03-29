// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type orderTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *orderTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("orders").
func (v *orderTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *orderTableType) Columns() []string {
	return []string{
		"id",
		"created_at",
		"user_id",
		"product_id",
		"discount",
		"quantity",
		"subtotal",
		"tax",
		"total",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *orderTableType) NewStruct() reform.Struct {
	return new(Order)
}

// NewRecord makes a new record for that table.
func (v *orderTableType) NewRecord() reform.Record {
	return new(Order)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *orderTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// OrderTable represents orders view or table in SQL database.
var OrderTable = &orderTableType{
	s: parse.StructInfo{
		Type:    "Order",
		SQLName: "orders",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UserID", Type: "int64", Column: "user_id"},
			{Name: "ProductID", Type: "int64", Column: "product_id"},
			{Name: "Discount", Type: "*float64", Column: "discount"},
			{Name: "Quantity", Type: "int32", Column: "quantity"},
			{Name: "Subtotal", Type: "float64", Column: "subtotal"},
			{Name: "Tax", Type: "float64", Column: "tax"},
			{Name: "Total", Type: "float64", Column: "total"},
		},
		PKFieldIndex: 0,
	},
	z: new(Order).Values(),
}

// String returns a string representation of this struct or record.
func (s Order) String() string {
	res := make([]string, 9)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[2] = "UserID: " + reform.Inspect(s.UserID, true)
	res[3] = "ProductID: " + reform.Inspect(s.ProductID, true)
	res[4] = "Discount: " + reform.Inspect(s.Discount, true)
	res[5] = "Quantity: " + reform.Inspect(s.Quantity, true)
	res[6] = "Subtotal: " + reform.Inspect(s.Subtotal, true)
	res[7] = "Tax: " + reform.Inspect(s.Tax, true)
	res[8] = "Total: " + reform.Inspect(s.Total, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Order) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.CreatedAt,
		s.UserID,
		s.ProductID,
		s.Discount,
		s.Quantity,
		s.Subtotal,
		s.Tax,
		s.Total,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Order) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.CreatedAt,
		&s.UserID,
		&s.ProductID,
		&s.Discount,
		&s.Quantity,
		&s.Subtotal,
		&s.Tax,
		&s.Total,
	}
}

// View returns View object for that struct.
func (s *Order) View() reform.View {
	return OrderTable
}

// Table returns Table object for that record.
func (s *Order) Table() reform.Table {
	return OrderTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Order) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Order) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Order) HasPK() bool {
	return s.ID != OrderTable.z[OrderTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Order) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = OrderTable
	_ reform.Struct = (*Order)(nil)
	_ reform.Table  = OrderTable
	_ reform.Record = (*Order)(nil)
	_ fmt.Stringer  = (*Order)(nil)
)

func init() {
	parse.AssertUpToDate(&OrderTable.s, new(Order))
}
