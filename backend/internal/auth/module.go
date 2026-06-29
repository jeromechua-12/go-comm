package auth

import (
	"database/sql"
	
	"github.com/jeromechua-12/go-comm/module"
)

type AuthModule struct {
	Handler *handler
}

func New(db *sql.DB) module.Module {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	return &AuthModule{Handler: handler}
}

func (m *AuthModule) Routes() []module.Route {
	return []module.Route{
		{
			Method: "POST",
			Path: "/api/user/signup",
			HandlerFunc: m.Handler.userSignup,
			Protected: false,
		},
		{
			Method: "POST",
			Path: "/api/user/login",
			HandlerFunc: m.Handler.userLogin,
			Protected: false,
		},
	}
}
