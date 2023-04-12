package answer

import (
	"fmt"
	"time"
)

func Experiment5() {
	fmt.Println("\n\nExperiment 5")
	address := "DKI Jakarta"
	addresses := [1000]string{address}

	findAddressExperiment5(addresses, address)
}

func findAddressExperiment5(addresses [1000]string, address string) {
	tx := time.Now()

	fmt.Printf("Default address is: %v\n", addresses[0])

	dur := time.Since(tx)

	fmt.Printf("The performance is: %v ms\n", dur.Milliseconds())
}
