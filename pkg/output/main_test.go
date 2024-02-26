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

func TestRenderError(t *testing.T) {
	var testData = `Error Message`
	var testSuccess1 = `{"error":"Error Message"}
`
	var testSuccess2 = `Error: Error Message
`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   true,
		Pretty: false,
	}
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	RenderError(testData, outputOptions)
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess1 {
		checkEqual(t, string(out), testSuccess1, " render error json processing failed")
	}

	outputOptions.JSON = false
	rescueStdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w
	RenderError(testData, outputOptions)
	w.Close()
	out, _ = io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess2 {
		checkEqual(t, string(out), testSuccess2, " render error stdout processing failed")
	}
}

func TestRenderInfo(t *testing.T) {
	var testData = `Info Message`
	var testSuccess1 = `{"info":"Info Message"}
`
	var testSuccess2 = `Info: Info Message
`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   true,
		Pretty: false,
	}
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	RenderInfo(testData, outputOptions)
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess1 {
		checkEqual(t, string(out), testSuccess1, " render info json processing failed")
	}

	outputOptions.JSON = false
	rescueStdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w
	RenderInfo(testData, outputOptions)
	w.Close()
	out, _ = io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess2 {
		checkEqual(t, string(out), testSuccess2, " render info stdout processing failed")
	}
}

func TestRenderOutput(t *testing.T) {
	var testData = `{"header":["NID","NotificationName","Channel","Webhook"],"data":[["1","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`
	var testSuccess1 = `{"data":[{"channel":"lagoon-local-ci","nid":"1","notificationname":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]}
`
	var testSuccess2 = `NID	NOTIFICATIONNAME         	CHANNEL        	WEBHOOK
1  	amazeeio--lagoon-local-ci	lagoon-local-ci	https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn
`
	var testSuccess3 = `1	amazeeio--lagoon-local-ci	lagoon-local-ci	https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn
`

	outputOptions := Options{
		Header: false,
		CSV:    false,
		JSON:   true,
		Pretty: false,
	}

	var dataMain Table
	json.Unmarshal([]byte(testData), &dataMain)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	RenderOutput(dataMain, outputOptions)
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess1 {
		checkEqual(t, string(out), testSuccess1, " render output json processing failed")
	}

	outputOptions.JSON = false
	rescueStdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w
	RenderOutput(dataMain, outputOptions)
	w.Close()
	out, _ = io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess2 {
		checkEqual(t, string(out), testSuccess2, " render output table stdout processing failed")
	}

	outputOptions.Header = true
	rescueStdout = os.Stdout
	r, w, _ = os.Pipe()
	os.Stdout = w
	RenderOutput(dataMain, outputOptions)
	w.Close()
	out, _ = io.ReadAll(r)
	os.Stdout = rescueStdout
	if string(out) != testSuccess3 {
		checkEqual(t, string(out), testSuccess3, " render output table stdout no header processing failed")
	}
}
