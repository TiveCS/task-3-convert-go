package answer

import (
	"fmt"

	"github.com/TiveCS/task-3-convert-go/utility"
)

var resultsExperiment9 = [][]string{}

func Experiment9() {
	fmt.Println("\n\nExperiment 9")
	candidates := []string{"Baswedan", "Subianto", "Maharani", "Ganjar"}

	chairs := arrangeExperiment9(candidates, nil)

	utility.RenderTable(chairs, []string{"1", "2", "3", "4"})
}

func arrangeExperiment9(array []string, memory []string) [][]string {
	var current string

	if memory == nil {
		memory = []string{}
	}

	for i := 0; i < len(array); i++ {
		current, array[i] = array[i], array[len(array)-1]
		array = array[:len(array)-1]

		if len(array) == 0 {
			resultsExperiment9 = append(resultsExperiment9, append(memory, current))
		}

		arrangeExperiment9(append([]string{}, array...), append(memory, current))

		array = append(array, current)
		array[i], array[len(array)-1] = array[len(array)-1], array[i]
	}

	return resultsExperiment9
}
