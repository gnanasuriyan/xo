// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/gnanasuriyan/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// UserMapping represents a row from 'information_schema.user_mappings'.
type UserMapping struct {
	AuthorizationIdentifier pgtypes.SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
	ForeignServerCatalog    pgtypes.SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       pgtypes.SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
}
