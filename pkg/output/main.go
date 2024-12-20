package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/logrusorgru/aurora"
)

// Table .
type Table struct {
	Header []string `json:"header"`
	Data   []Data   `json:"data"`
}

// Data .
type Data []string

// Options .
type Options struct {
	Header    bool
	CSV       bool
	JSON      bool
	Pretty    bool
	Debug     bool
	Error     string
	MultiLine bool
}

// Result .
type Result struct {
	ResultData map[string]interface{} `json:"data,omitempty"`
	Result     string                 `json:"result,omitempty"`
	Error      string                 `json:"error,omitempty"`
	Info       string                 `json:"info,omitempty"`
}

// RenderJSON .
func RenderJSON(data interface{}, opts Options) string {
	var jsonBytes []byte
	var err error
	if opts.Pretty {
		jsonBytes, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}
	} else {
		jsonBytes, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
	}
	return string(jsonBytes)
}

// RenderError .
func RenderError(errorMsg string, opts Options) {
	if opts.JSON {
		jsonData := Result{
			Error: trimQuotes(errorMsg),
		}
		RenderJSON(jsonData, opts)
	} else {
		os.Stderr.WriteString(fmt.Sprintf("Error: %s", trimQuotes(errorMsg)))
	}
}

// RenderInfo .
func RenderInfo(infoMsg string, opts Options) {
	if opts.JSON {
		jsonData := Result{
			Info: trimQuotes(infoMsg),
		}
		RenderJSON(jsonData, opts)
	} else {
		os.Stderr.WriteString(fmt.Sprintf("Info: %s", trimQuotes(infoMsg)))
	}
}

// RenderResult .
func RenderResult(result Result, opts Options) string {
	var out bytes.Buffer
	if opts.JSON {
		return RenderJSON(result, opts)
	} else {
		if trimQuotes(result.Result) == "success" {
			out.WriteString(fmt.Sprintf("Result: %s\n", aurora.Green(trimQuotes(result.Result))))
			if len(result.ResultData) != 0 {
				for k, v := range result.ResultData {
					out.WriteString(fmt.Sprintf("%s: %v\n", k, v))
				}
			}
		} else {
			fmt.Printf("Result: %s\n", aurora.Yellow(trimQuotes(result.Result)))
			if len(result.ResultData) != 0 {
				for k, v := range result.ResultData {
					out.WriteString(fmt.Sprintf("%s: %v\n", k, v))
				}
			}
		}
	}
	return out.String()
}

// RenderOutput .
func RenderOutput(data Table, opts Options) string {
	var out bytes.Buffer
	if opts.Debug {
		out.WriteString(fmt.Sprintf("%s\n", aurora.Yellow("Final result:")))
	}
	if opts.JSON {
		// really basic tabledata to json implementation
		var rawData []interface{}
		for _, dataValues := range data.Data {
			jsonData := make(map[string]interface{})
			for indexID, dataValue := range dataValues {
				dataHeader := strings.ReplaceAll(strings.ToLower(data.Header[indexID]), " ", "-")
				jsonData[dataHeader] = dataValue
			}
			rawData = append(rawData, jsonData)
		}
		returnedData := map[string]interface{}{
			"data": rawData,
		}
		return RenderJSON(returnedData, opts)
	} else {
		// otherwise render a table
		if opts.Error != "" {
			os.Stderr.WriteString(opts.Error)
		}
		t := table.NewWriter()
		opts.Header = !opts.Header
		if opts.Header {
			var hRow table.Row
			for _, k := range data.Header {
				hRow = append(hRow, k)
			}
			t.AppendHeader(hRow)
		}
		t.SetOutputMirror(&out)
		for _, rowData := range data.Data {
			var dRow table.Row
			for _, k := range rowData {
				dRow = append(dRow, k)
			}
			t.AppendRow(dRow)
		}
		t.SetStyle(table.StyleDefault)
		t.Style().Options = table.OptionsNoBordersAndSeparators
		t.Style().Box.PaddingLeft = ""    // trim left space
		t.Style().Box.PaddingRight = "\t" // pad right with tab
		if !opts.MultiLine {
			t.SuppressTrailingSpaces() // suppress the trailing spaces if not multiline
		}
		if opts.MultiLine {
			// stops multiline values bleeding into other columns
			t.SetColumnConfigs([]table.ColumnConfig{
				{Name: "Value", WidthMax: 75}, // Set specific width for "Value" column if multiline
				{Name: "Token", WidthMax: 50}, // Set specific width for "Token" column if multiline
			})
		}

		if opts.CSV {
			t.RenderCSV()
			return out.String()
		}
		t.Render()
		return out.String()
	}
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
