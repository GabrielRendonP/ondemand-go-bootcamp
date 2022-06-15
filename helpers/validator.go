package helpers

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type QueryStruct struct {
	Ipw   int    `schema:"ipw"`
	Items int    `schema:"items"`
	Type  string `schema:"type"`
}

func ValidateParams(values url.Values) error {
	required := []string{
		"ipw",
		"items",
		"type",
	}

	var queryStruct QueryStruct

	err := decoder.Decode(&queryStruct, values)
	if err != nil {
		err := fmt.Errorf("%s", err)
		return err
	}

	if len(values) != len(required) {
		err := errors.New("wrong number of params")
		return err
	}

	for _, v := range required {
		if !values.Has(v) {
			err := fmt.Errorf("missing param: %s", v)
			return err
		}
	}

	return nil
}
