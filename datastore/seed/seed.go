package main

// ** WARNING **
// This program deletes all database tables and seeds with default values

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"guzfolio/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db.Migrator().DropTable(&model.Currency{})
	db.Migrator().DropTable(&model.User{})
	db.Migrator().DropTable(&model.Portfolio{})
	db.Migrator().DropTable(&model.Transaction{})

	// automigrate model
	db.AutoMigrate(&model.Currency{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Portfolio{})
	db.AutoMigrate(&model.Transaction{})

	// create currencies
	t := time.Date(2020, time.Month(10), 22, 11, 0, 0, 0, time.UTC)
	db.Create(&model.Currency{Code: "BTC", Name: "Bitcoin", MarketValue: 12909.29, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "ETH", Name: "Ethereum", MarketValue: 411.30, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "USDT", Name: "Tether", MarketValue: 1.01, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "XRP", Name: "XRP", MarketValue: 0.262765, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "BCH", Name: "Bitcoin Cash", MarketValue: 270.01, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "BNB", Name: "Binance Coin", MarketValue: 31.02, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "LINK", Name: "Chainlink", MarketValue: 11.54, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "DOT", Name: "Polkadot", MarketValue: 4.31, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "ADA", Name: "Litecoin", MarketValue: 0.109642, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "BSV", Name: "Bitcoin SV", MarketValue: 170.10, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "CRO", Name: "Crypto.com Coin", MarketValue: 0.100035, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "EOS", Name: "EOS", MarketValue: 2.68, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "XMR", Name: "Monero", MarketValue: 126.12, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "TRON", Name: "TRX", MarketValue: 0.02739823, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "Tezos", Name: "XTZ", MarketValue: 2.22, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "XLM", Name: "Stellar", MarketValue: 0.086645, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "NEO", Name: "Neo", MarketValue: 18.86, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "WBTC", Name: "Wrapped Bitcoin", MarketValue: 13002.70, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "LEO", Name: "UNUS SED LEO", MarketValue: 1.25, CreatedAt: t, UpdatedAt: t})
	db.Create(&model.Currency{Code: "ATOM", Name: "Cosmos", MarketValue: 5.48, CreatedAt: t, UpdatedAt: t})

	// create users
	bytePassword := []byte("guzfolio1234")
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	pass := string(passwordHash)

	db.Create(&model.User{Name: "Default User", Email: "user@guzfolio.dev", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Default Admin", Email: "admin@guzfolio.dev", Password: pass, IsAdmin: true})
	db.Create(&model.User{Name: "Eadie Duffield", Email: "eduffield1@opensource.org", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Dory Burchett", Email: "dburchett2@slashdot.org", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Gaylord Son", Email: "gson3@chicagotribune.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Justin Plastow", Email: "jplastow4@tuttocitta.it", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Lorene McGarel", Email: "lmcgarel5@scientificamerican.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Faun McLugaish", Email: "fmclugaish6@paginegialle.it", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Maris Jensen", Email: "mjensen7@si.edu", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Agneta Pantlin", Email: "apantlin8@macromedia.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Janetta Sagar", Email: "jsagar9@squarespace.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Derby O'Duane", Email: "doduanea@theguardian.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Hettie Glendinning", Email: "hglendinningb@weather.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Findlay Barlass", Email: "fbarlassc@hatena.ne.jp", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Clerissa Merrywether", Email: "cmerrywetherd@fda.gov", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Max Sleath", Email: "msleathe@latimes.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Beale Fransewich", Email: "bfransewichf@yandex.ru", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Morgun Gerin", Email: "mgering@independent.co.uk", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Devin Antoinet", Email: "dantoineth@typepad.com", Password: pass, IsAdmin: false})
	db.Create(&model.User{Name: "Cody Ossulton", Email: "cossultoni@cnbc.com", Password: pass, IsAdmin: false})

	// new random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// get all currencies
	var allCurrencies []model.Currency
	db.Find(&allCurrencies)

	// create portfolios
	defaultPortfolio := "Default portfolio"
	for userID := 1; userID < 21; userID++ {
		portfolio := &model.Portfolio{Name: &defaultPortfolio, UserID: uint(userID)}
		db.Create(portfolio)

		totalTx := r.Intn(20) + 5
		for i := 0; i < totalTx; i++ {
			currency := allCurrencies[r.Intn(len(allCurrencies))]
			q := rand.Float64() * 10

			// create some transactions
			tx := model.Transaction{
				Quantity:     q,
				CurrencyID:   currency.ID,
				PricePerCoin: currency.MarketValue,
				PortfolioID:  portfolio.ID,
			}
			db.Create(&tx)
		}
	}

	defaultPortfolio = "Default portfolio #02"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1})
	defaultPortfolio = "Default portfolio #03"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1})
	defaultPortfolio = "Default portfolio #04"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 1})

	defaultPortfolio = "Default portfolio #02"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2})
	defaultPortfolio = "Default portfolio #03"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2})
	defaultPortfolio = "Default portfolio #04"
	db.Create(&model.Portfolio{Name: &defaultPortfolio, UserID: 2})
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