package main

import (
	"strings"
	"testing"
)

func TestAddParseTimeParam(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantContains string
		wantErr      bool
	}{
		{"With Scheme", "http://example.com", "parseTime=true", false},
		{"Without Scheme", "example.com", "parseTime=true", false},
		{"Invalid URL", "this is not a URL", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addParseTimeParam(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("addParseTimeParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, tt.wantContains) {
				t.Errorf("addParseTimeParam() = %v, want %v", got, tt.wantContains)
			}
		})
	}
}
