package mapper

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func MapVars(c *map[string]interface{}) error {
	for s, i := range *c {
		if value, ok := os.LookupEnv(strings.ToUpper(s)); ok {
			switch i.(type) {
			case string:
				(*c)[s] = value
			case int:
				if n, e := strconv.Atoi(value); e != nil {
					return e
				} else {
					(*c)[s] = n
				}
			case float64:
				if n, e := strconv.ParseFloat(value, 64); e != nil {
					return e
				} else {
					(*c)[s] = n
				}
			case bool:
				value = strings.ToLower(value)
				if value == "yes" || value == "true" || value == "1" ||
					value == "si" || value == "hai" || value == "ja" ||
					value == "da" || value == "sim" || value == "ok" {
					(*c)[s] = true
				} else {
					(*c)[s] = false
				}
			default:
				var obj interface{}
				if e := json.Unmarshal([]byte(value), &obj); e != nil {
					return e
				}
				(*c)[s] = obj
			}
		}
	}
	return nil
}
