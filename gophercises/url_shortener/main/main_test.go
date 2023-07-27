package main

import "testing"

// Tests successful open of yaml file provided
func TestReadYAMLOk(t *testing.T) {
	// We don't care about the data returned
	// Just want to check the file is opened ok
	// Yaml parsing is checked in handler_test
	res, err := readYAML("paths_test.yaml")
	if err != nil {
		t.Errorf("TestReadYAMLOk failed. Err is not nil. Data read: %v", res)
	}
}

func TestReadYAMLErr(t *testing.T) {
	// We don't care about the data returned
	// Just want to check the file is opened ok
	// Yaml parsing is checked in handler_test
	res, err := readYAML("some_non_existet_path.yaml")
	if err == nil {
		t.Errorf("TestReadYAMLErr failed. Err needs to be nil. Data read: %v", res)
	}
}
