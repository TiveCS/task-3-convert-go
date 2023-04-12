package answer

import "fmt"

func Experiment13() {
	fmt.Println("\n\nExperiment 13")
	dataPromo := generateDataExperiment13(100000000)
	dataDiscount := generateDataExperiment13(1000)
	sendPromoDiscountExperiment13(dataPromo, dataDiscount)
}

func generateDataExperiment13(n int) []string {
	baseNumber := "082"
	customers := []string{}

	var mobileNumber string

	for i := 0; i < n; i++ {
		mobileNumber = baseNumber + fmt.Sprintf("%09d", i)
		customers = append(customers, mobileNumber)
	}

	return customers
}

func sendPromoDiscountExperiment13(input1 []string, input2 []string) {
	for i := 0; i < len(input1); i++ {
		fmt.Printf("Sending promo to %s\n", input1[i])
	}
	fmt.Printf("Its finished send promo to %d customers\n", len(input1))

	for i := 0; i < len(input2); i++ {
		fmt.Printf("Sending promo to choosen customer %d\n", (i + 1))
	}
	fmt.Printf("Its finished send discount to %d customers\n", len(input2))
}
