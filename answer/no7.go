package answer

import "fmt"

func Experiment7() {
	fmt.Println("\n\nExperiment 7")
	foods := []string{"Sate", "Bakso", "Dimsum", "Rames"}
	drinks := []string{"Jeruk", "Teh", "Kelapa", "Cendol"}

	logPairs(foods, drinks, "Menu")
}

func logPairs(arr1 []string, arr2 []string, words string) {
	counter := 0

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			counter++
			fmt.Printf("%s %d: %s dan %s\n", words, counter, arr1[i], arr2[j])
		}
	}
}
