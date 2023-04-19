package runtime

import (
	"reflect"
	"testing"
)

// TestNewRuntime tests the NewRuntime function
func TestNewRuntime(t *testing.T) {
	// Define test cases as a slice of structs
	var tests []struct {
		name    string
		want    *Runtime
		wantErr bool
	}
	// Loop through the test cases
	for _, tt := range tests {
		// Run a subtest for each test case
		t.Run(tt.name, func(t *testing.T) {
			// Call the NewRuntime function
			got, err := NewRuntime()
			// Check if an error occurred (or didn't occur) as expected
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRuntime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if the output of the NewRuntime function matches the expected output
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuntime() got = %v, want %v", got, tt.want)
			}
		})
	}

}
