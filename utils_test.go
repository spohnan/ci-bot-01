// Package bot is the top level container for ci-bot-01 code
package bot

import (
	"testing"
)

func TestMapKeyPresent(t *testing.T) {
	m := make(map[string]interface{})
	m["present"] = "foo"
	if !mapContainsKey(m, "present") {
		t.Fatal("key not found")
	}
	if mapContainsKey(m, "not-present") {
		t.Fatal("non-existant key found")
	}
}
