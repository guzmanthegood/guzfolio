package model

type Portfolio struct {
	ID           	uint      	`json:"id"`
	Name         	*string   	`json:"name"`
	User         	User     	`json:"user"`
	UserID         	uint	  	`json:"userID"`
}

type PortfolioReport struct {
	TotalTransactions	int			`json:"totalTransactions"`
	TotalValue			float64		`json:"totalValue"`
	NetCost				float64		`json:"netCost"`
	PercentChange		float64		`json:"percentChange"`
}