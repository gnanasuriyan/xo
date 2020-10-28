// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/gnanasuriyan/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// PgTruetypid calls the stored procedure 'information_schema._pg_truetypid(pg_attribute, pg_type) oid' on db.
func PgTruetypid(db XODB, v0 pgtypes.PgAttribute, v1 pgtypes.PgType) (pgtypes.Oid, error) {
	var err error

	// sql query
	const sqlstr = `SELECT information_schema._pg_truetypid($1, $2)`

	// run query
	var ret pgtypes.Oid
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return pgtypes.Oid{}, err
	}

	return ret, nil
}
