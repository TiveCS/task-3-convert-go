package answer

import (
	"fmt"
	"time"
)

func Experiment4() {
	fmt.Println("\n\nExperiment 4")
	address := "DKI Jakarta"
	addresses := [10]string{address}

	findAddressExperiment4(addresses, address)
}

func findAddressExperiment4(addresses [10]string, address string) {
	tx := time.Now()

	fmt.Printf("Default address is: %v\n", addresses[0])

	dur := time.Since(tx)

	fmt.Printf("The performance is: %v ms\n", dur.Milliseconds())
}
