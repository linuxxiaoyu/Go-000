package main

import (
	"testing"

	"go.uber.org/goleak"
)

// func TestRealMain(t *testing.T) {
// 	realMain()
// }

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
