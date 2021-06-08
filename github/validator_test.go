package github

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
)

// switch english brain
//
func TestGson1(t *testing.T) {
	type Params struct {
		A string `json:"a" validate:"required"`
		B string `json:"b"`
	}
	j := `{"a":"1","b":"2"}`
	param := new(Params)
	err := json.Unmarshal([]byte(j), &param)
	if err != nil {
		t.Fatal(err)
	}
	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(param)
	t.Log(string(b))
}
