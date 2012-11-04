package model

import (
    "testing"
)

func TestGetManagerInstance(t *testing.T) {
    m1 := GetManagerInstance()
    m2 := GetManagerInstance()
    if m1 != m2 {
        t.Errorf("two instance are not the same")
    }
}
