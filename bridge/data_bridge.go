package bridge

import "encoding/json"

func Encode(data []string) string {
	b, _ := json.Marshal(data)
	return string(b)
}
