package mapper

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

const configJson = `{
	"port": 8080,
	"host": "0.0.0.0",
    "user_admin": "admin",
	"useSSL": false,
    "tags": [],
	"alias": {}
}`

func checkMapField(m *map[string]interface{}, k string, v interface{}) bool {
	mv, ok := (*m)[k]
	return ok && mv == v
}
func setEnvVars(m map[string]string) error {
	for n, v := range m {
		if e := os.Setenv(n, v); e != nil {
			return e
		}
	}
	return nil
}

func unsetEnvVars(m map[string]string) error {
	for n := range m {
		if e := os.Unsetenv(n); e != nil {
			return e
		}
	}
	return nil
}

func TestMapVarsWithAllDefaultValues(t *testing.T) {
	var config map[string]interface{}

	if e := json.Unmarshal([]byte(configJson), &config); e != nil {
		t.Fatalf("cannot unmarshal json:\n%s", e)
	}

	if config == nil {
		t.Fatalf("config was not unmarshal")
	}

	if e := MapVars(&config); e != nil {
		t.Fatalf("cannot map the env-vars")
	}

	if !checkMapField(&config, "port", float64(8080)) {
		t.Errorf("the port field is not 8080")
	}

	if !checkMapField(&config, "host", "0.0.0.0") {
		t.Errorf("the host field is not 0.0.0.0")
	}

	if !checkMapField(&config, "user_admin", "admin") {
		t.Errorf("the user field is not admin")
	}

	if !checkMapField(&config, "useSSL", false) {
		t.Errorf("the useSSL field is not false")
	}

	if reflect.TypeOf(config["tags"]).Kind() != reflect.Slice {
		t.Errorf("the tags field is not []")
	}

	if reflect.TypeOf(config["alias"]).Kind() != reflect.Map {
		t.Errorf("the alias field is not {}")
	}
}

func TestMapVarsWithSetEnvVars(t *testing.T) {
	var config map[string]interface{}

	if e := json.Unmarshal([]byte(configJson), &config); e != nil {
		t.Fatalf("cannot unmarshal json:\n%s", e)
	}

	if config == nil {
		t.Fatalf("config was not unmarshal")
	}

	vars := map[string]string{
		"port":       "7000",
		"host":       "localhost",
		"user_admin": "god",
		"useSSL":     "yes",
		"tags":       `["dev","test","prod"]`,
		"alias":      `{"test-1":"www.test.com", "test-2":["dev", "test", "prod"], "test-3":{"key":"value"}}`,
	}

	if e := setEnvVars(vars); e != nil {
		t.Errorf("the env-vars cannot be set:\n%s", e)
	}

	if e := MapVars(&config); e != nil {
		t.Fatalf("cannot map the env-vars")
	}

	if !checkMapField(&config, "port", float64(7000)) {
		t.Errorf("the port field is not 700")
	}

	if !checkMapField(&config, "host", "localhost") {
		t.Errorf("the host field is not localhost")
	}

	if !checkMapField(&config, "user_admin", "god") {
		t.Errorf("the user field is not god")
	}

	if !checkMapField(&config, "useSSL", true) {
		t.Errorf("the useSSL field is not true")
	}

	if reflect.TypeOf(config["tags"]).Kind() != reflect.Slice {
		t.Errorf("the tags field is not []")
	} else {
		tags := config["tags"].([]interface{})
		if len(tags) != 3 {
			t.Errorf("the tags field has not 3 items")
		} else {
			if tags[0] != "dev" {
				t.Errorf("the item 0 from tags field is not dev")
			}
			if tags[1] != "test" {
				t.Errorf("the item 0 from tags field is not test")
			}
			if tags[2] != "prod" {
				t.Errorf("the item 0 from tags field is not prod")
			}
		}
	}

	if reflect.TypeOf(config["alias"]).Kind() != reflect.Map {
		t.Errorf("the alias field is not {}")
	} else {
		alias := config["alias"].(map[string]interface{})
		if len(alias) != 3 {
			t.Errorf("the alias field has not 3 items")
		} else {
			if alias["test-1"] != "www.test.com" {
				t.Errorf("the item test-1 from alias field is not www.test.com")
			}

			test2 := alias["test-2"].([]interface{})
			if len(test2) != 3 {
				t.Errorf("the test2 field of alias has not 3 items")
			} else {
				if test2[0] != "dev" {
					t.Errorf("the item 0 from test2 field of alias is not dev")
				}
				if test2[1] != "test" {
					t.Errorf("the item 0 from test2 field of alias is not test")
				}
				if test2[2] != "prod" {
					t.Errorf("the item 0 from test2 field of alias is not prod")
				}
			}

			test3 := alias["test-3"].(map[string]interface{})
			if len(test3) != 3 {
				t.Errorf("the test3 field of alias has not 3 items")
			} else {
				if test3["key"] != "value" {
					t.Errorf("the key field from test3 field of alias is not value")
				}
			}
		}
	}
}
