package config_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/cwansart/mcs-status-go/config"
)

func TestSettings(t *testing.T) {
	t.Run("create new config file", func(t *testing.T) {
		t.Parallel()
		configFile := getTempConfigFileName(t)

		config.NewConfig(configFile)

		if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
			t.Errorf("Test should have created config file '%s' but did not", configFile)
		}
	})

	t.Run("should use default config", func(t *testing.T) {
		t.Parallel()
		configFile := getTempConfigFileName(t)

		got := config.NewConfig(configFile)

		want := "http://localhost:2006"
		if got.ServerUrl != want {
			t.Errorf("ServerUrl should be %s", want)
		}
	})

	t.Run("read stored config file", func(t *testing.T) {
		configFile := getTempConfigFileName(t)
		saveAsFile(t, config.Config{"http://localhost:8888"}, configFile)

		got := config.NewConfig(configFile)

		want := "http://localhost:8888"
		if got.ServerUrl != want {
			t.Errorf("ServerUrl should be %s", want)
		}
	})
}

func getTempConfigFileName(t *testing.T) (c string) {
	t.Helper()
	dir := t.TempDir()
	c = fmt.Sprintf("%s/%s", dir, "config.json")
	return c
}

func saveAsFile(t *testing.T, c config.Config, f string) {
	t.Helper()

	json, err := json.Marshal(c)
	if err != nil {
		t.Error("Failed to create json test data")
	}

	err = os.WriteFile(f, []byte(json), 0600)
	if err != nil {
		t.Error("Failed to write temp config file")
	}
}
