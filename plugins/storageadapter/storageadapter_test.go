package main

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"aos_vis/dataadaptertest"
)

/*******************************************************************************
 * Var
 ******************************************************************************/

var (
	adapterInfo dataadaptertest.TestAdapterInfo
)

/*******************************************************************************
 * Init
 ******************************************************************************/

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		FullTimestamp:    true})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

/*******************************************************************************
 * Main
 ******************************************************************************/

func TestMain(m *testing.M) {
	configJSON := `{"Data": {
		"Attribute.Vehicle.VehicleIdentification.VIN":    {"Value": "TestVIN", "Public": true,"ReadOnly": true},
		"Attribute.Vehicle.UserIdentification.Users":     {"Value": ["User1", "Provider1"], "Public": true},

		"Signal.Drivetrain.InternalCombustionEngine.RPM": {"Value": 1000, "ReadOnly": true},

		"Signal.Body.Trunk.IsLocked":                     {"Value": false},
		"Signal.Body.Trunk.IsOpen":                       {"Value": true},

		"Signal.Cabin.Door.Row1.Right.IsLocked":          {"Value": true},
		"Signal.Cabin.Door.Row1.Right.Window.Position":   {"Value": 50},
		"Signal.Cabin.Door.Row1.Left.IsLocked":           {"Value": true},
		"Signal.Cabin.Door.Row1.Left.Window.Position":    {"Value": 23},
		"Signal.Cabin.Door.Row2.Right.IsLocked":          {"Value": false},
		"Signal.Cabin.Door.Row2.Right.Window.Position":   {"Value": 100},
		"Signal.Cabin.Door.Row2.Left.IsLocked":           {"Value": true},
		"Signal.Cabin.Door.Row2.Left.Window.Position":    {"Value": 0}
	}}`

	storageAdapter, err := NewAdapter([]byte(configJSON))
	if err != nil {
		log.Fatalf("Can't create storage adapter: %s", err)
	}
	defer storageAdapter.Close()

	adapterInfo = dataadaptertest.TestAdapterInfo{
		Name:        "StorageAdapter",
		PathListLen: 13,
		Adapter:     storageAdapter,
		SetData: map[string]interface{}{
			"Signal.Cabin.Door.Row1.Right.IsLocked":        true,
			"Signal.Cabin.Door.Row1.Right.Window.Position": 200,
			"Signal.Cabin.Door.Row1.Left.IsLocked":         false,
			"Signal.Cabin.Door.Row1.Left.Window.Position":  100,
			"Signal.Cabin.Door.Row2.Right.IsLocked":        true,
			"Signal.Cabin.Door.Row2.Right.Window.Position": 400,
			"Signal.Cabin.Door.Row2.Left.IsLocked":         false,
			"Signal.Cabin.Door.Row2.Left.Window.Position":  50},
		SubscribeList: []string{
			"Signal.Cabin.Door.Row1.Right.IsLocked",
			"Signal.Cabin.Door.Row1.Right.Window.Position",
			"Signal.Cabin.Door.Row1.Left.IsLocked",
			"Signal.Cabin.Door.Row1.Left.Window.Position",
			"Signal.Cabin.Door.Row2.Right.IsLocked",
			"Signal.Cabin.Door.Row2.Right.Window.Position",
			"Signal.Cabin.Door.Row2.Left.IsLocked",
			"Signal.Cabin.Door.Row2.Left.Window.Position"},
		SetSubscribeData: map[string]interface{}{
			"Signal.Cabin.Door.Row1.Right.IsLocked":        false,
			"Signal.Cabin.Door.Row1.Right.Window.Position": 100,
			"Signal.Cabin.Door.Row1.Left.IsLocked":         true,
			"Signal.Cabin.Door.Row1.Left.Window.Position":  50,
			"Signal.Cabin.Door.Row2.Right.IsLocked":        false,
			"Signal.Cabin.Door.Row2.Right.Window.Position": 60,
			"Signal.Cabin.Door.Row2.Left.IsLocked":         true,
			"Signal.Cabin.Door.Row2.Left.Window.Position":  70},
	}

	ret := m.Run()

	os.Exit(ret)
}

/*******************************************************************************
 * Tests
 ******************************************************************************/

func TestGetName(t *testing.T) {
	if err := dataadaptertest.GetName(&adapterInfo); err != nil {
		t.Errorf("Test get name error: %s", err)
	}
}

func TestGetPathList(t *testing.T) {
	if err := dataadaptertest.GetPathList(&adapterInfo); err != nil {
		t.Errorf("Test get path lis error: %s", err)
	}
}

func TestPublicPath(t *testing.T) {
	if err := dataadaptertest.PublicPath(&adapterInfo); err != nil {
		t.Errorf("Test public path error: %s", err)
	}
}

func TestGetSetData(t *testing.T) {
	if err := dataadaptertest.GetSetData(&adapterInfo); err != nil {
		t.Errorf("Test get set data error: %s", err)
	}
}

func TestSubscribeUnsubscribe(t *testing.T) {
	if err := dataadaptertest.SubscribeUnsubscribe(&adapterInfo); err != nil {
		t.Errorf("Test subscribe unsubscribe error: %s", err)
	}
}
