package models

import (
 "testing"
)

func TestUnitTestPassed(t *testing.T) {
    if (UnitTestPassed() != 1) {
        t.Error("unitTest() return and error")
    }
}

/*func TestUnitTestFailed(t *testing.T) {
    t.Fail()
    if (UnitTestFailed() != 1) {
        t.Error("unitTest() return and error")
    }
}*/