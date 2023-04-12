package answer

import (
	"fmt"
	"time"
)

func Experiment6() {
	fmt.Println("\n\nExperiment 6")
	address := "DKI Jakarta"
	addresses := [100000]string{address}

	findAddressExperiment6(addresses, address)
}

func findAddressExperiment6(addresses [100000]string, address string) {
	tx := time.Now()

	fmt.Printf("Default address is: %v\n", addresses[0])

	dur := time.Since(tx)

	fmt.Printf("The performance is: %v ms\n", dur.Milliseconds())
}
