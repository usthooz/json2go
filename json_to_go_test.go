package main

import (
	"testing"
)

func TestJson2To(t *testing.T) {
	genStruct("json2go.json", "file", "")
}
