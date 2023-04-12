package answer

import (
	"fmt"

	"github.com/TiveCS/task-3-convert-go/utility"
)

var resultsExperiment8 = [][]string{}

func Experiment8() {
	fmt.Println("\n\nExperiment 8")
	candidates := []string{"Baswedan", "Subianto", "Maharani"}

	chairs := arrangeExperiment8(candidates, nil)

	utility.RenderTable(chairs, []string{"1", "2", "3"})
}

func arrangeExperiment8(array []string, memory []string) [][]string {
	var current string

	if memory == nil {
		memory = []string{}
	}

	for i := 0; i < len(array); i++ {
		current, array[i] = array[i], array[len(array)-1]
		array = array[:len(array)-1]

		if len(array) == 0 {
			resultsExperiment8 = append(resultsExperiment8, append(memory, current))
		}

		arrangeExperiment8(append([]string{}, array...), append(memory, current))

		array = append(array, current)
		array[i], array[len(array)-1] = array[len(array)-1], array[i]
	}

	return resultsExperiment8
}
