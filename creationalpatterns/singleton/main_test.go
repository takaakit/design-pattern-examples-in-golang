package singleton

import (
	"fmt"
	"testing"
)

/*
Check whether the same instance is obtained.
*/

func TestMain(m *testing.M) {
	obj1 := NewSingleton()
	obj2 := NewSingleton()
	if obj1 == obj2 {
		fmt.Println("obj1 and obj2 are the same instance.")
	} else {
		fmt.Println("obj1 and obj2 are different instances.")
	}
}
