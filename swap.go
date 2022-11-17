package common

import "encoding/json"

func SwapTo(source, target interface{}) (err error) {
	dataByte, err := json.Marshal(source)
	if err != nil {
		return
	}

	err = json.Unmarshal(dataByte, target)
	return
}
