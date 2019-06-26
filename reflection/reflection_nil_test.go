package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectNil(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			t.Error("should panic")
		}
	}()
	reflect.SliceOf(nil)
}
