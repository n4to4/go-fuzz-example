// +build gofuzzbeta

package main

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func FuzzParseQuery(f *testing.F) {
	f.Add("x=1&y=2")
	f.Fuzz(func(t *testing.T, queryStr string) {
		query, err := url.ParseQuery(queryStr)
		if err != nil {
			t.Skip()
		}
		queryStr2 := query.Encode()
		query2, err := url.ParseQuery(queryStr2)

		fmt.Printf("==> %v, %v\n", query, query2)

		if err != nil {
			t.Fatalf("ParseQuery failed to decode a valid encoded query %s: %v", queryStr2, err)
		}
		if !reflect.DeepEqual(query, query2) {
			t.Errorf("ParseQuery gave different query after being encoded\nbefore: %v\nafter: %v", query, query2)
		}
	})
}
