package answer

import (
	"fmt"
	"strconv"
)

func Experiment11() {
	fmt.Println("\n\nExperiment 11")

	customers := generateDataExperiment11()
	sendPromoDiscountExperiment11(customers)
}

func generateDataExperiment11() []string {
	number := "0821234567"
	customers := []string{}

	var mobileNumber string

	for i := 0; i < 100; i++ {
		if i < 10 {
			mobileNumber = number + "0" + strconv.Itoa(i)
		} else {
			mobileNumber = number + strconv.Itoa(i)
		}

		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscountExperiment11(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("Sending promo to %s\n", array[i])
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Sending promo to choosen customer %d\n", (i + 1))
	}
}
