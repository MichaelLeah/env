package env

import (
	"os"
	"testing"
)

func setupSingle() error {
	f, err := os.Create("./.env")
	if err != nil {
		return err
	}

	f.WriteString("API_KEY=123456789")
	f.Close()

	return nil
}

func setupMulti() error {
	f, err := os.Create("./.env")
	if err != nil {
		return err
	}

	f.WriteString("API_KEY=123456789\nANOTHER_KEY=54321")
	f.Close()

	return nil
}

func TestEnv(t *testing.T) {
	err := setupSingle()
	if err != nil {
		t.Error("Unable to create .env file for tests.", err)
	}

	apiKey, err := Get("API_KEY")
	if err != nil {
		t.Error("Unable to get key's values.", err)
	}

	if apiKey != "123456789" {
		t.Error("API_KEY did not equal expected case: 123456789, got: ", apiKey)
	}

	os.Remove("./.env")
}

func TestEnvMultiline(t *testing.T) {
	err := setupMulti()
	if err != nil {
		t.Error("Unable to create .env file for tests.", err)
	}

	apiKey, err := Get("API_KEY")
	if err != nil {
		t.Error("Unable to get key's values.", err)
	}

	anotherKey, err := Get("ANOTHER_KEY")
	if err != nil {
		t.Error("Unable to get key's values.", err)
	}

	if apiKey != "123456789" {
		t.Error("API_KEY did not equal expected case: 123456789, got: ", apiKey)
	}

	if anotherKey != "54321" {
		t.Error("ANOTHER_KEY did not equal expected case: 54321, got: ", anotherKey)
	}

	os.Remove("./.env")
}
