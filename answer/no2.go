package answer

import (
	"fmt"
	"time"
)

func Experiment2() {
	fmt.Println("\n\nExperiment 2")
	company := "Telkom"

	largeCompanyName := [1000]string{company}

	findCompanyExperiment2(largeCompanyName, company)
}

func findCompanyExperiment2(companies [1000]string, company string) {
	tx := time.Now()

	for i := 0; i < len(companies); i++ {
		if companies[i] == company {
			fmt.Println("Company found!")
		}
	}

	dur := time.Since(tx)
	fmt.Printf("The performance is: %v ms", dur.Milliseconds())
}
