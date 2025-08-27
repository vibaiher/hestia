package i18n

import (
	"encoding/json"
	"os"
	
	"hestia/static"
)

type Translations struct {
	App struct {
		Name      string `json:"name"`
		Subtitle  string `json:"subtitle"`
		TitleFull string `json:"title_full"`
	} `json:"app"`
	Home struct {
		Welcome     string `json:"welcome"`
		Description string `json:"description"`
		SystemInfo  string `json:"system_info"`
		CurrentDate string `json:"current_date"`
		Status      string `json:"status"`
		StatusOk    string `json:"status_ok"`
	} `json:"home"`
	Features struct {
		Title                 string `json:"title"`
		IncomeTracking        string `json:"income_tracking"`
		ExpenseManagement     string `json:"expense_management"`
		BankMonitoring        string `json:"bank_monitoring"`
		InvestmentTracking    string `json:"investment_tracking"`
		ExpenseCategorization string `json:"expense_categorization"`
	} `json:"features"`
	Footer struct {
		Tagline string `json:"tagline"`
		Motto   string `json:"motto"`
	} `json:"footer"`
}

func Load() (Translations, error) {
	var translations Translations
	
	data, err := os.ReadFile(static.File("locales/es.json"))
	if err != nil {
		return translations, err
	}
	
	err = json.Unmarshal(data, &translations)
	return translations, err
}