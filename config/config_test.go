package config_test

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"gitpct.epam.com/epmd-aepr/aos_vis/config"
)

/*******************************************************************************
 * Private
 ******************************************************************************/

func createConfigFile() (err error) {
	configContent := `{
"ServerUrl": "localhost:8088",
"VISCert": "wwwivi.crt.pem",
"VISKey": "wwwivi.key.pem",
"Adapters":[{
		"Plugin": "test1"
	}, {
		"Plugin": "test2"
	}, {
		"Plugin": "test3"
	}]
}`

	if err := ioutil.WriteFile(path.Join("tmp", "visconfig.json"), []byte(configContent), 0644); err != nil {
		return err
	}

	return nil
}

func setup() (err error) {
	if err := os.MkdirAll("tmp", 0755); err != nil {
		return err
	}

	if err = createConfigFile(); err != nil {
		return err
	}

	return nil
}

func cleanup() (err error) {
	if err := os.RemoveAll("tmp"); err != nil {
		return err
	}

	return nil
}

/*******************************************************************************
 * Main
 ******************************************************************************/

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatalf("Error creating service images: %s", err)
	}

	ret := m.Run()

	if err := cleanup(); err != nil {
		log.Fatalf("Error cleaning up: %s", err)
	}

	os.Exit(ret)
}

/*******************************************************************************
 * Tests
 ******************************************************************************/

func TestGetCredentials(t *testing.T) {
	config, err := config.New("tmp/visconfig.json")
	if err != nil {
		t.Fatalf("Error opening config file: %s", err)
	}

	if config.ServerURL != "localhost:8088" {
		t.Errorf("Wrong ServerURL value: %s", config.ServerURL)
	}

	if config.VISCert != "wwwivi.crt.pem" {
		t.Errorf("Wrong VISCert value: %s", config.VISCert)
	}

	if config.VISKey != "wwwivi.key.pem" {
		t.Errorf("Wrong VISKey value: %s", config.VISKey)
	}
}

func TestAdapters(t *testing.T) {
	config, err := config.New("tmp/visconfig.json")
	if err != nil {
		t.Fatalf("Error opening config file: %s", err)
	}

	if len(config.Adapters) != 3 {
		t.Errorf("Wrong adapters len: %d", len(config.Adapters))
	}

	if config.Adapters[0].Plugin != "test1" || config.Adapters[1].Plugin != "test2" || config.Adapters[2].Plugin != "test3" {
		t.Error("Wrong adapter name")
	}
}
