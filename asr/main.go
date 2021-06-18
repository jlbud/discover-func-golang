package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"os/exec"
	"strings"
)

type OutputText struct {
	PuncResult   string      `json:"punc_result"`
	PuncText     string      `json:"punc_text"`
	PuncWordInfo interface{} `json:"punc_word_info"`
	Result       string      `json:"result"`
	Status       string      `json:"status"`
	Text         string      `json:"text"`
	Type         string      `json:"type"`
	VarPuncText  string      `json:"var_punc_text"`
}

type LastText struct {
	AllResult []LastAllResult `json:"allResult"`
}

type LastAllResult struct {
	Text OutputText `json:"text"`
}

func main() {
	exe := exec.Command("/bin/bash", "-c", `gurl ws -ac 1 -H "rate:16000" -w -H "appkey:xuqo7pqagqx5gvdbqyfybrusfosbbkjjtfvsr5qx" -H "closeVad:true"  -H "session-id:uuidgen" -H "wordBoundary:1" -binary -p "@./xinwen.pcm" eof -H "eof:eof" -send-rate "8000/250ms" -url ws://192.168.5.213:10036/ws/asr/pcm`)
	out, err := exe.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("exe.Output err %v", err))
		return
	}
	// split by char "}{"
	strArr := strings.Split(string(out), "}{")
	if len(strArr) == 0 {
		fmt.Println(fmt.Errorf("strArr len is 0"))
		return
	}

	// spliced correctly
	var buf bytes.Buffer
	for i, v := range strArr {
		if i == 0 { // first
			buf.WriteString(fmt.Sprintf("[%s},", v))
		} else if i == len(strArr)-1 { // last one
			buf.WriteString(fmt.Sprintf("{%s]", v))
		} else { // others
			buf.WriteString(fmt.Sprintf("{%s},", v))
		}
	}

	otSli := make([]OutputText, 0)
	bufStr := buf.String()
	for i := 0; i < len(strArr); i++ {
		var v gjson.Result
		//isLastBody := gjson.Get(bufStr, fmt.Sprintf("%d.last_body", i)).Exists()
		allArr := gjson.Get(bufStr, fmt.Sprintf("%d.allResult", i)).Array()

		//if isLastBody { // last one all result
		//	all := gjson.Get(bufStr, fmt.Sprintf("%d.last_body", i)).String()
		//	last := LastText{}
		//	err := json.Unmarshal([]byte(all), &last)
		//	if err != nil {
		//		fmt.Println(fmt.Errorf("LastText json.Unmarshal err %v", err))
		//		return
		//	}
		//	otSli = append(otSli, last.AllResult[0].Text)
		//	continue
		//} else
		if len(allArr) > 0 { // all result
			v = gjson.Get(bufStr, fmt.Sprintf("%d.allResult.0.text", i))
		} else { // current result
			v = gjson.Get(bufStr, fmt.Sprintf("%d.currResult.text", i))
		}
		ot := OutputText{}
		ot.PuncResult = gjson.Get(v.String(), "punc_result").String()
		ot.PuncText = gjson.Get(v.String(), "punc_text").String()
		ot.PuncWordInfo = gjson.Get(v.String(), "punc_word_info").Raw
		ot.Result = gjson.Get(v.String(), "result").String()
		ot.Status = gjson.Get(v.String(), "status").String()
		ot.Text = gjson.Get(v.String(), "text").String()
		ot.Type = gjson.Get(v.String(), "type").String()
		ot.VarPuncText = gjson.Get(v.String(), "var_punc_text").String()
		otSli = append(otSli, ot)
	}

	for i, v := range otSli {
		val, err := json.Marshal(v)
		if err != nil {
			fmt.Println(fmt.Errorf("LastText json.Marshal err %v", err))
			return
		}
		fmt.Println(i)
		fmt.Println(string(val))
	}
}
