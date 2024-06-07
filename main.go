package main

import (
	// "fmt"

	"fmt"

	"example.com/pricecalculator/cmdmanager"
	// "example.com/pricecalculator/filemanager"
	"example.com/pricecalculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
