package answer

import (
	"fmt"
)

func Experiment12() {
	fmt.Println("\n\nExperiment 12")
	customers := generateDataExperiment12(1000)
	sendPromoDiscountExperiment12(customers)
}

func generateDataExperiment12(n int) []string {
	baseNumber := "082"
	customers := []string{}

	var mobileNumber string

	for i := 0; i < n; i++ {
		mobileNumber = baseNumber + fmt.Sprintf("%09d", i)
		customers = append(customers, mobileNumber)
	}

	return customers
}

func sendPromoDiscountExperiment12(input []string) {
	for i := 0; i < len(input); i++ {
		fmt.Printf("Sending promo to %s\n", input[i])
	}
	fmt.Printf("Its finished send promo to %d customers\n", len(input))

	for i := 0; i < len(input); i++ {
		fmt.Printf("Sending promo to choosen customer %d\n", (i + 1))
	}
	fmt.Printf("Its finished send discount to %d customers\n", len(input))
}
