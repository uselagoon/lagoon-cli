package output

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"testing"
)

func checkEqual(t *testing.T, got, want interface{}, msgs ...interface{}) {
	if !reflect.DeepEqual(got, want) {
		buf := bytes.Buffer{}
		buf.WriteString("got:\n[%v]\nwant:\n[%v]\n")
		for _, v := range msgs {
			buf.WriteString(v.(string))
		}
		t.Errorf(buf.String(), got, want)
	}
}

func TestRenderJSON(t *testing.T) {
	var testData = `Error Message`
	var testSuccess = `{
  "error": "Error Message"
}`
	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   true,
		Pretty: true,
	}

	jsonData := Result{
		Error: trimQuotes(testData),
	}
	output := RenderJSON(jsonData, outputOptions)
	if output != testSuccess {
		checkEqual(t, output, testSuccess, " render error json processing failed")
	}
}

func TestRenderError(t *testing.T) {
	var testData = `Error Message`
	var testSuccess = `Error: Error Message`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   false,
		Pretty: false,
	}

	out := captureStderr(func() {
		RenderError(testData, outputOptions)
	})
	checkEqual(t, out, testSuccess, " render error stdout processing failed")

	var testJson = `{"error":"Error Message"}`
	outputOptions.JSON = true
	out = captureStderr(func() {
		RenderError(testData, outputOptions)
	})
	checkEqual(t, out, testJson, " render error stdout processing failed")
}

func TestRenderInfo(t *testing.T) {
	var testData = `Info Message`
	var testSuccess1 = `Info: Info Message`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   false,
		Pretty: false,
	}

	out := captureStderr(func() {
		RenderInfo(testData, outputOptions)
	})
	checkEqual(t, out, testSuccess1, " render info stdout processing failed")

	var testJson = `{"info":"Info Message"}`
	outputOptions.JSON = true
	out = captureStderr(func() {
		RenderInfo(testData, outputOptions)
	})
	checkEqual(t, out, testJson, " render error stdout processing failed")
}

func TestRenderOutput(t *testing.T) {
	var testData = `{"header":["NID","NotificationName","Channel","Webhook"],"data":[["1","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`
	var testSuccess1 = `NID	NOTIFICATIONNAME         	CHANNEL        	WEBHOOK
1  	amazeeio--lagoon-local-ci	lagoon-local-ci	https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn
`
	var testSuccess2 = `1	amazeeio--lagoon-local-ci	lagoon-local-ci	https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn
`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   false,
		Pretty: false,
	}

	var dataMain Table
	_ = json.Unmarshal([]byte(testData), &dataMain)

	output := RenderOutput(dataMain, outputOptions)
	if output != testSuccess1 {
		checkEqual(t, output, testSuccess1, " render output table stdout processing failed")
	}

	outputOptions.Header = true
	output = RenderOutput(dataMain, outputOptions)
	if output != testSuccess2 {
		checkEqual(t, output, testSuccess2, " render output table stdout no header processing failed")
	}
}

func TestRenderString(t *testing.T) {
	var testData = `Plain string message`
	var testSuccess1 = "Plain string message\n"

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   false,
		Pretty: false,
	}

	output := RenderString(testData, outputOptions)
	checkEqual(t, output, testSuccess1, " render string processing failed")

	var testSuccess2 = `"Plain string message"`
	outputOptions.JSON = true

	output = RenderString(testData, outputOptions)
	checkEqual(t, output, testSuccess2, " render string json processing failed")
}

func captureStderr(fn func()) string {
	var out bytes.Buffer
	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	fn()

	w.Close()
	os.Stderr = rescueStderr

	_, _ = io.Copy(&out, r)
	return out.String()
}
