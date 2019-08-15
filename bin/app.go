package main

import (
	"cloudfunction/domain"
	"cloudfunction/scrape"
	"cloudfunction/util"
	"fmt"
)

func main() {

	funcCh := make(chan func())
	util.StartPool(20, funcCh)
	final := make(chan string)

	funcCh <- func() {
		oemCh := scrape.GetAllOEMs(domain.GetOEMURL())
		for oemRow := range oemCh {

			funcCh <- func() {
				oemPageCh := scrape.GetAllOEMPages(domain.GetOEMPagesURL(), oemRow)
				for oemPageRow := range oemPageCh {
					urlsCh := scrape.GetAllOEMPageResultUrls(oemPageRow)
					for url := range urlsCh {
						fmt.Println(url)
						final <- string(url)
					}
				}
			}
		}
	}

	for s := range final {
		fmt.Println(s)
	}

}
