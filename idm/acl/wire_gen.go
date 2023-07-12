// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package acl

import (
	"context"
	"github.com/pydio/cells/v4/common/sql"
	"github.com/pydio/cells/v4/common/utils/configx"
)

// Injectors from wire.go:

func InitializeDAO() MyDAO {
	context := newContext()
	values := newConfigStore()
	dao := newSQLDAO(context, values)
	aclSqlimpl := newSQLImpl(dao)
	myDAO := NewMyDAO(aclSqlimpl)
	return myDAO
}

// wire.go:

type MyDAO struct {
	DAO
}

func NewMyDAO(dao DAO) MyDAO {
	return MyDAO{
		DAO: dao,
	}
}

func newSQLImpl(dao sql.DAO) *sqlimpl {
	return &sqlimpl{dao}
}

func newSQLDAO(ctx context.Context, store configx.Values) sql.DAO {
	dao, _ := sql.NewDAO(ctx, store.Val("driver").String(), store.Val("dsn").String(), store.Val("prefix").String())
	return dao.(sql.DAO)
}

func newConfigStore() configx.Values {
	c := configx.New()
	c.Val("driver").Set("mysql")
	c.Val("dsn").Set("root:P@ssw0rd@tcp(localhost:3306)/cells?parseTime=true")
	c.Val("prefix").Set("idm_acl_")

	return c
}

func newContext() context.Context {
	return context.TODO()
}
