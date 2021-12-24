package helpers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func ToStringJson(i interface{}) string {
	r, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		logrus.Error(err)
	}
	return string(r)
}
