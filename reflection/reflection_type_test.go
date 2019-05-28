package reflection

import (
	"reflect"
	"testing"
)

func TestTypeOfInt(t *testing.T) {

	type tt int
	var t1 tt = 123

	t.Logf("type tt int: %v", reflect.TypeOf(t1).Kind())
}
