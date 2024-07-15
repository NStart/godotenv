package godotenv

import (
	"os"
	"testing"
)

var noopPresets = make(map[string]string)

func parseAndCompare(t *testing.T, rawEnvLine string, expectedKey string, expectedValue string) {
	result, err := Unmarshal(rawEnvLine)

	if err != nil {
		t.Errorf("expected %q to parse as %q: %q, errored %q", rawEnvLine, expectedKey, expectedValue, err)
	}

	if result[expectedKey] != expectedValue {
		t.Errorf("Expected '%v' to parse as '%v' => '%v', got %q instead", rawEnvLine, expectedKey, expectedValue, result)
	}

}

func loadEnvAndCompareValues(t *testing.T, loader func(files ...string) error, envFileName string, expectedValues map[string]string, presets map[string]string) {
	os.Clearenv()

	for k, v := range presets {
		os.Setenv(k, v)
	}

	err := loader(envFileName)
	if err != nil {
		t.Fatalf("Error loading  %v", envFileName)
	}

	for k := range expectedValues {
		envValue := os.Getenv(k)
		v := expectedValues[k]
		if envValue != v {
			t.Errorf("Mismatch for key '%v': expected '%#v' got '%#v'", k, v, envValue)
		}
	}
}

func TestLoadWithNoArgsLoadsDotEnv(t *testing.T) {
	err := Load()
	pathError := err.(*os.PathError)
	if pathError == nil || pathError.Op != "open" || pathError.Path != ".env" {
		t.Errorf("Didn't try and open .env by default")
	}
}
