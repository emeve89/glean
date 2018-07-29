package main

import (
	"github.com/emeve89/glean/infrastructure"
	"github.com/emeve89/glean/interfaces"
	"github.com/emeve89/glean/usecases"
	"net/http"
	)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.db")

	handlers := make(map[string] interfaces.DbHandler)
	handlers["DbUserRepo"] 		= dbHandler
	handlers["DbCustomerRepo"] 	= dbHandler
	handlers["DbOrderRepo"] 	= dbHandler
	handlers["DbItemRepo"] 		= dbHandler

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

	webserviceHandler := interfaces.WebServiceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}