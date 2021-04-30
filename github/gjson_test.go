package github

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestGJson(t *testing.T) {
	j := `{
    "Version": "2.0.0",
    "Language": "en",
    "AudioCheckShow": 40,
    "Text": "This is a phone test, ee[p:iː]",
    "EvalType": "enWordDetect",
    "AgeGroup": "preschool",
    "ScoreCoefficient": 1.6,
    "PointSystem": 80,
    "Precision": 0.1,
    "Grade": 4,
    "Dialect": "us",
    "Syllable": 0,
    "Phoneme": 1,
    "Match": 1,
    "Stress": true,
    "X-Contentid": "",
    "Phone": true,
    "Grammar": "enumerate",
    "GrammarWeight": "",
    "Reference": ""
}`
	v := gjson.Get(j, "n").String()
	if v == "" {
		t.Error("the 'v' is empty")
	} else {
		t.Log("from 'j' get 'Version' and value is ", v)
	}
}
