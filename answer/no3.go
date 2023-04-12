package answer

import (
	"fmt"
	"time"
)

func Experiment3() {
	fmt.Println("\n\nExperiment 3")
	company := "Telkom"

	largeCompanyName := [100000]string{company}

	findCompanyExperiment3(largeCompanyName, company)
}

func findCompanyExperiment3(companies [100000]string, company string) {
	tx := time.Now()

	for i := 0; i < len(companies); i++ {
		if companies[i] == company {
			fmt.Println("Company found!")
		}
	}

	dur := time.Since(tx)
	fmt.Printf("The performance is: %v ms", dur.Milliseconds())
}
