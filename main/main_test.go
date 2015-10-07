package main

import (
	"testing"

	"github.com/gregpechiro/walkerTR"
)

var FOLDERCOUNT = 3

func TestWalker(t *testing.T) {
	m := walkerTR.Walk("db")
    if len(m) != FOLDERCOUNT {
        t.Error("WalkerTR Failed")
    }
}
