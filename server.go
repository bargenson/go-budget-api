package main

import (
	"github.com/gorilla/mux"
	"github.com/bargenson/go-budget-api/routes"
)

func newServer(db *database) *mux.Router {
	router := mux.NewRouter()
  routes.AddBudgetsController(router.PathPrefix("/budgets").Subrouter())
  return router
}
