package main

import (
	"testing"
)

func TestInitApp(t *testing.T) {
	realMain()
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
