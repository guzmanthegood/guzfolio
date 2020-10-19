package main

// ** WARNING **
// This program deletes all database tables and seeds with default values

import (
	"bufio"
	"fmt"
	"gorm.io/gorm/logger"
	"guzfolio/model"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if !askForConfirmation("Do you really want to reset your database?") {
		return
	}

	// new data base connection
	db, err := gorm.Open(postgres.Open(os.Getenv("PG_CONNECTION_STRING")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// drop tables
	if db.Migrator().HasTable(&model.Currency{}) {
		db.Migrator().DropTable(&model.Currency{})
	}
	if db.Migrator().HasTable(&model.User{}) {
		db.Migrator().DropTable(&model.User{})
	}
	if db.Migrator().HasTable(&model.Portfolio{}) {
		db.Migrator().DropTable(&model.Portfolio{})
	}

	// automigrate model
	db.AutoMigrate(&model.Currency{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Portfolio{})

	// create currencies
	db.Create(&model.Currency{Code: "USD", Name: "United States dollar", Type: model.CurrencyTypeFiat})
	db.Create(&model.Currency{Code: "EUR", Name: "Euro", Type: model.CurrencyTypeFiat})
	db.Create(&model.Currency{Code: "JPY", Name: "Japanese yen", Type: model.CurrencyTypeFiat})
	db.Create(&model.Currency{Code: "GBP", Name: "Pound sterling", Type: model.CurrencyTypeFiat})
	db.Create(&model.Currency{Code: "AUD", Name: "Australian dollar", Type: model.CurrencyTypeFiat})

	db.Create(&model.Currency{Code: "BTC", Name: "Bitcoin", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "ETH", Name: "Ethereum", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "USDT", Name: "Tether", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "XRP", Name: "XRP", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "BCH", Name: "Bitcoin Cash", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "BNB", Name: "Binance Coin", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "LINK", Name: "Chainlink", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "DOT", Name: "Polkadot", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "ADA", Name: "Litecoin", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "BSV", Name: "Bitcoin SV", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "CRO", Name: "Crypto.com Coin", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "EOS", Name: "EOS", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "XMR", Name: "Monero", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "TRON", Name: "TRX", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "Tezos", Name: "XTZ", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "XLM", Name: "Stellar", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "NEO", Name: "Neo", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "WBTC", Name: "Wrapped Bitcoin", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "LEO", Name: "UNUS SED LEO", Type: model.CurrencyTypeCrypto})
	db.Create(&model.Currency{Code: "ATOM", Name: "Cosmos", Type: model.CurrencyTypeCrypto})

	// create users
	db.Create(&model.User{Name: "Ravi Radish", Email: "rradish0@github.com"})
	db.Create(&model.User{Name: "Eadie Duffield", Email: "eduffield1@opensource.org"})
	db.Create(&model.User{Name: "Dory Burchett", Email: "dburchett2@slashdot.org"})
	db.Create(&model.User{Name: "Gaylord Son", Email: "gson3@chicagotribune.com"})
	db.Create(&model.User{Name: "Justin Plastow", Email: "jplastow4@tuttocitta.it"})
	db.Create(&model.User{Name: "Lorene McGarel", Email: "lmcgarel5@scientificamerican.com"})
	db.Create(&model.User{Name: "Faun McLugaish", Email: "fmclugaish6@paginegialle.it"})
	db.Create(&model.User{Name: "Maris Jensen", Email: "mjensen7@si.edu"})
	db.Create(&model.User{Name: "Agneta Pantlin", Email: "apantlin8@macromedia.com"})
	db.Create(&model.User{Name: "Janetta Sagar", Email: "jsagar9@squarespace.com"})
	db.Create(&model.User{Name: "Derby O'Duane", Email: "doduanea@theguardian.com"})
	db.Create(&model.User{Name: "Hettie Glendinning", Email: "hglendinningb@weather.com"})
	db.Create(&model.User{Name: "Findlay Barlass", Email: "fbarlassc@hatena.ne.jp"})
	db.Create(&model.User{Name: "Clerissa Merrywether", Email: "cmerrywetherd@fda.gov"})
	db.Create(&model.User{Name: "Max Sleath", Email: "msleathe@latimes.com"})
	db.Create(&model.User{Name: "Beale Fransewich", Email: "bfransewichf@yandex.ru"})
	db.Create(&model.User{Name: "Morgun Gerin", Email: "mgering@independent.co.uk"})
	db.Create(&model.User{Name: "Devin Antoinet", Email: "dantoineth@typepad.com"})
	db.Create(&model.User{Name: "Cody Ossulton", Email: "cossultoni@cnbc.com"})
	db.Create(&model.User{Name: "Susette Varndall", Email: "varndallj@wsj.com"})

	// create portfolios
	defaultPortfolio := "Default portfolio"
	currencyID := 1
	for userID := 1; userID < 21; userID++ {
		db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: uint(userID), FiatCurrencyID: uint(currencyID)})
		currencyID++
		if currencyID > 5 {
			currencyID = 1
		}
	}

	defaultPortfolio = "Default portfolio #02"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1, FiatCurrencyID: 2})
	defaultPortfolio = "Default portfolio #03"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1, FiatCurrencyID: 3})
	defaultPortfolio = "Default portfolio #04"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1, FiatCurrencyID: 4})

	defaultPortfolio = "Default portfolio #02"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2, FiatCurrencyID: 3})
	defaultPortfolio = "Default portfolio #03"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2, FiatCurrencyID: 4})
	defaultPortfolio = "Default portfolio #04"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2, FiatCurrencyID: 5})
}

func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
