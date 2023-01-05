package writer

import "encoding/json"

func ToJSON(value any) string {
	j, _ := json.Marshal(value)

	return string(j)
}
