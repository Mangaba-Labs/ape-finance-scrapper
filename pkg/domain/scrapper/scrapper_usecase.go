package scrapper

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
	"github.com/mxschmitt/playwright-go"
)

// Scrapper struct implementation
type Scrapper struct {
	Repository repository.Repository
}

// GetStocks Scrapper to get all stocks in database
func (s Scrapper) UpdateShares() (error) {
	stockModels, err := s.Repository.FindAll()
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

	responseSlice := []model.Share{}

	var wg sync.WaitGroup
	for i := 0; i < len(stockModels); i++ {
		wg.Add(1)
		go worker(&wg, browser, stockModels[i], &responseSlice)
	}
	wg.Wait()
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v\n", err)
		return err
	}
	return nil
}

// CheckIfExists check if stock exists before register in database
func (s Scrapper) CheckIfExists(bvmf string) (stock *model.Share, err error) {
	// Opening browser
	pw, err := playwright.Run()
	if err != nil {
		return
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return
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

	// Finding company name
	companyEntry, err := page.QuerySelectorAll("div.tv-symbol-header__first-line")

	if err != nil {
		log.Fatalln(err)
		return
	}

	if len(companyEntry) == 0 {
		return stock, errors.New("Company not founded")
	}

	stock.Company, _ = companyEntry[0].InnerText()

	// Closing browser
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v\n", err)
		return
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v\n", err)
		return
	}
	return stock, nil
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
	// imageEntry, err := page.QuerySelectorAll("img.tv-circle-logo.tv-circle-logo--large.tv-category-header__icon")
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// image, err := imageEntry[0].GetAttribute("src")
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

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
	price, _ := strconv.ParseFloat(value, 2)

	scrapped.Price = float32(price)
	scrapped.Variation = variation
	return scrapped, nil
}

// Async method to get scrapped data and parse to stockResponse
func workerUpdate(wg *sync.WaitGroup, browser playwright.Browser, stock model.Share, stockResponse *[]model.Share) {
	defer wg.Done()

	scrapped, _ := scrapStock(browser, stock.Bvmf)

	var response model.Share

	response.Price = scrapped.Price
	response.Variation = scrapped.Variation

	*stockResponse = append(*stockResponse, response)
}

func worker(wg *sync.WaitGroup, browser playwright.Browser, stock model.Share, stockResponse *[]model.Share) {
	defer wg.Done()

	scrapped, _ := scrapStock(browser, stock.Bvmf)

	var response model.Share

	response.Price = scrapped.Price
	response.Variation = scrapped.Variation

	*stockResponse = append(*stockResponse, response)
}
