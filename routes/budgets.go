package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bargenson/go-budget-api/services"
)

const BudgetPath = "/api/budgets"

type budgetsController struct {

}

func (ctlr *budgetsController) getBudgets(w http.ResponseWriter, r *http.Request) {
	var service services.BudgetService
	json.NewEncoder(w).Encode(service.GetBudgets())
}

func AddBudgetsController(r *mux.Router) {
	var controller budgetsController
	r.Methods("GET").HandlerFunc(controller.getBudgets)
}