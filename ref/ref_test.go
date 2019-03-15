package ref

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	name string
}

var _abc string

func TestRefInstance(t *testing.T) {
	userMap := make(map[User]string)

	u1 := User{"u1"}
	u12 := User{"u1"}

	userMap[u1] = "1"
	if _, ok := userMap[u1]; !ok {
		assert.FailNow(t, "should find in map")
	}
	if _, ok := userMap[u12]; !ok {
		assert.FailNow(t, "should find in map")
	}
}

func TestRefPointer(t *testing.T) {
	userMap := make(map[*User]string)

	u1 := &User{"u1"}
	u12 := &User{"u1"}

	userMap[u1] = "1"
	if _, ok := userMap[u1]; !ok {
		assert.FailNow(t, "should find in map")
	}

	if _, ok := userMap[u12]; ok {
		assert.FailNow(t, "should not find in map")
	}
}

func TestRefInterface(t *testing.T) {
	userMap := make(map[interface{}]string)

	u1 := User{"u1"}
	u12 := User{"u1"}

	userMap[u1] = "1"
	if _, ok := userMap[u1]; !ok {
		assert.FailNow(t, "should find in map")
	}

	if _, ok := userMap[u12]; !ok {
		assert.FailNow(t, "should find in map")
	}

	u2 := &User{"u2"}
	u22 := &User{"u2"}

	userMap[u2] = "1"
	if _, ok := userMap[u2]; !ok {
		assert.FailNow(t, "u2 should find in map")
	}

	if _, ok := userMap[u22]; ok {
		assert.FailNow(t, "u22 should not find in map")
	}

	u3 := &User{"u3"}
	u32 := User{"u3"}

	userMap[u3] = "1"
	if _, ok := userMap[u3]; !ok {
		assert.FailNow(t, "u3 should find in map")
	}

	if _, ok := userMap[u32]; ok {
		assert.FailNow(t, "u32 should not find in map")
	}
}
