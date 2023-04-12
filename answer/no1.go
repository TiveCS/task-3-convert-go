package answer

import (
	"fmt"
	"time"
)

func Experiment1() {
	fmt.Println("Experiment 1")
	company := "Telkom"

	largeCompanyName := [100]string{company}

	findCompanyExperiment1(largeCompanyName, company)
}

func findCompanyExperiment1(companies [100]string, company string) {
	tx := time.Now()

	for i := 0; i < len(companies); i++ {
		if companies[i] == company {
			fmt.Println("Company found!")
		}
	}

	dur := time.Since(tx)
	fmt.Printf("The performance is: %v ms", dur.Milliseconds())
}
