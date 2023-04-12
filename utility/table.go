package utility

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(data [][]string, header []string) {
	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)

	// Set the table headers
	table.SetHeader(header)

	// Add the data to the table
	for _, row := range data {
		table.Append(row)
	}

	// Render the table
	table.Render()
}
