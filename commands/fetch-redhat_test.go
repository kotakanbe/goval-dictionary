package commands

import (
	"reflect"
	"testing"
)

func Test_RedhatOvalNamePattern(t *testing.T) {
	type response struct {
		isMatched bool
		matched   []string
	}
	tests := []struct {
		name     string
		osVer    string
		expected response
	}{
		{
			"RHEL 4",
			"4",
			response{isMatched: false, matched: nil},
		},
		{
			"RHEL 5",
			"5",
			response{isMatched: true, matched: []string{"5", "5", "", "", ""}},
		},
		{
			"RHEL 8.1 EUS",
			"8.1-eus",
			response{isMatched: true, matched: []string{"8.1-eus", "", "8", "1", "eus"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected.isMatched != RedhatOvalNamePattern.MatchString(tt.osVer) {
				t.Errorf("RedhatOvalNamePattern.MatchString(%s) = %v, want %v", tt.osVer, RedhatOvalNamePattern.MatchString(tt.osVer), tt.expected.isMatched)
			}

			if !reflect.DeepEqual(tt.expected.matched, RedhatOvalNamePattern.FindStringSubmatch(tt.osVer)) {
				t.Errorf("RedhatOvalNamePattern.FindStringSubmatch(%s) = %v, want %v", tt.osVer, RedhatOvalNamePattern.FindStringSubmatch(tt.osVer), tt.expected.matched)
			}
		})
	}
}
