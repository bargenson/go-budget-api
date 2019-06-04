package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bargenson/go-budget-api/models"
)

const BudgetPath = "/api/budgets"

type budgetsController struct {

}

func (ctlr *budgetsController) getBudgets(w http.ResponseWriter, r *http.Request) {
	var budgets = []models.Budgets{}
	json.NewEncoder(w).Encode(budgets)
}

func AddBudgetsController(r *mux.Router) {
	ctlr := new(budgetsController)
	r.Methods("GET").HandlerFunc(ctlr.getBudgets)
}