// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/gnanasuriyan/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// AdministrableRoleAuthorization represents a row from 'information_schema.administrable_role_authorizations'.
type AdministrableRoleAuthorization struct {
	Grantee     pgtypes.SQLIdentifier `json:"grantee"`      // grantee
	RoleName    pgtypes.SQLIdentifier `json:"role_name"`    // role_name
	IsGrantable pgtypes.YesOrNo       `json:"is_grantable"` // is_grantable
}
