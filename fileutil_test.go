// Package fileutil provides functions for doing things with files, like reading them
// line by line, etc
package fileutil

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Test_ReadPropertiesFile creates a test properties file and reads it into memory
func Test_ReadPropertiesFile(t *testing.T) {

    //
    // Create the test properties file
    //
	filename := "test.properties"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal("Error creating test properties file", err)
	}
	_, err = io.WriteString(file, `
# This is a test properties file
maxRows = 20     
somethingEnabled = true   
error.message = There was an error   
`)
	if err != nil {
		t.Fatal("Error writing test properties file", err)
	}
	file.Close()

    //
    // Read the properties from the file
    //
    var properties []Property
	properties, err = ReadPropertiesFile("test.properties")
	if err != nil {
		t.Fatal("Error reading test properties file", err)
	}
	if properties[0].Name != "maxRows" {
		t.Fail()
	}
	if properties[2].Name != "error.message" {
		t.Fail()
	}
	if properties[2].Value != "There was an error" {
        t.Fatalf("value is wrong for error.message: '%s'", properties[2].Value)
	}

}
