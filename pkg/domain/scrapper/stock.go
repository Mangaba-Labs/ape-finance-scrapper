package scrapper

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/Mangaba-Labs/ape-finance-scrapper/database"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	stockRepository "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
	"github.com/mxschmitt/playwright-go"
)

// Scrapper struct implementation
// UpdateShares function called in our cron job
func UpdateShares() error {
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	repositoryStock := stockRepository.Repository{
		DB: db,
	}

	stocks, err := repositoryStock.FindAll()
	fmt.Println(stocks)
	if err != nil {
		return err
	}
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return err
	}

	stocksUpdated := []model.Share{}

	var wg sync.WaitGroup
	for i := 0; i < len(stocks); i++ {
		wg.Add(1)
		go worker(&wg, browser, stocks[i], &stocksUpdated)
	}
	wg.Wait()
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v\n", err)
		return err
	}

	err = repositoryStock.Update(stocksUpdated)
	return err
}

// ScrapFullStock to create stock in database
func ScrapFullStock(bvmf string) (share model.Share, err error) {
	pw, err := playwright.Run()
	if err != nil {
		return share, err
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return share, err
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
		return
	}

	searchPage := fmt.Sprintf("https://www.tradingview.com/symbols/BMFBOVESPA-%s/", bvmf)
	if _, err = page.Goto(searchPage); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	companyEntry, err := page.QuerySelectorAll("div.tv-symbol-header__first-line")
	if err != nil {
		log.Fatalln(err)
		return
	}
	company, err := companyEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}

	imageEntry, err := page.QuerySelectorAll("img.tv-circle-logo.tv-circle-logo--large.tv-category-header__icon")
	if err != nil {
		log.Fatalln(err)
		return
	}
	image, err := imageEntry[0].GetAttribute("src")
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err != nil {
		log.Fatalln(err)
		return
	}
	// Variation
	variationValuesEntry, err := page.QuerySelectorAll("div.js-symbol-change-direction.tv-symbol-price-quote__change")
	if err != nil {
		log.Fatalln(err)
		return
	}
	variation, err := variationValuesEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}
	// Stock Value
	valueEntry, err := page.QuerySelectorAll("div.tv-symbol-price-quote__value.js-symbol-last")
	if err != nil {
		log.Fatalf("could not get entries: %v\n", err)
		return
	}
	value, err := valueEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}
	price, _ := strconv.ParseFloat(value, 2)

	pw.Stop()

	share.Bvmf = bvmf
	share.Company = company
	share.Price = float32(price)
	share.Variation = variation
	share.Image = image
	return share, nil
}

func scrapStock(browser playwright.Browser, bvmf string) (scrapped model.VariableData, err error) {
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
		return
	}

	searchPage := fmt.Sprintf("https://www.tradingview.com/symbols/BMFBOVESPA-%s/", bvmf)
	if _, err = page.Goto(searchPage); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// Variation
	variationValuesEntry, err := page.QuerySelectorAll("div.js-symbol-change-direction.tv-symbol-price-quote__change")
	if err != nil {
		log.Fatalln(err)
		return
	}
	variation, err := variationValuesEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}
	// Stock Value
	valueEntry, err := page.QuerySelectorAll("div.tv-symbol-price-quote__value.js-symbol-last")
	if err != nil {
		log.Fatalf("could not get entries: %v\n", err)
		return
	}
	value, err := valueEntry[0].InnerText()
	if err != nil {
		log.Fatalln(err)
		return
	}
	price, _ := strconv.ParseFloat(value, 2)

	scrapped.Price = float32(price)
	scrapped.Variation = variation

	return scrapped, nil
}

func worker(wg *sync.WaitGroup, browser playwright.Browser, stock model.Share, stockResponse *[]model.Share) {
	defer wg.Done()

	scrapped, _ := scrapStock(browser, stock.Bvmf)

	response := stock
	response.Price = scrapped.Price
	response.Variation = scrapped.Variation

	*stockResponse = append(*stockResponse, response)
}
