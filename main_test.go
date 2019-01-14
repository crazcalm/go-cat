package main

import (
    "testing"
)

func TestParseFileNames(t *testing.T) {
    tests := []struct{
    	Input string
    	Output []string
    }{
    	{"testing", []string{"testing"}},
    	{"   testing   ", []string{"testing"}},
    	{"testing multiple files", []string{"testing", "multiple", "files"}},
    	{"   testing   multiple   files   ", []string{"testing", "multiple", "files"}},
    }

    for index, test := range tests{
    	result := parseFileNames(test.Input)

    	if len(result) != len(test.Output){
    		t.Errorf("Test Case %d: Expected %d items, but got %d\nresult: %v\nExpected: %v", index + 1, len(test.Output), len(result), result, test.Output)
    	}

    	for i, name := range test.Output {
    		if name != result[i] {
    			t.Errorf("Test Case %d: Expected %s, but got %s", index + 1, name, result[i])
    		}
    	}
    }
}