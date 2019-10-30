package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// Table .
type Table struct {
	Header []string
	Data   []Data
}

// Data .
type Data []string

// RenderTable .
func RenderTable(tableData Table) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader(tableData.Header)
	for _, rowData := range tableData.Data {
		table.Append(rowData)
	}
	table.Render()
}
