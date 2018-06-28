// authors: wangoo
// created: 2018-06-22
// test interface inherit

package interfacc

import (
	"fmt"
	"testing"
)

type Ider interface {
	GetId() string
}

type Passworder interface {
	GetPassword() string
}

type Account interface {
	Ider
	Passworder
	GetDesc() string
}

type User struct {
	Id       string
	Password string
	Desc     string
}

func (u User) GetId() string {
	return u.Id
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetDesc() string {
	return u.Desc
}

func TestAccount(t *testing.T) {
	u := User{
		Id:       "u1",
		Password: "pass",
		Desc:     "manager",
	}

	fmt.Println(u.GetId(), u.GetPassword(), u.GetDesc())

	testIder(u)
	testPassword(u)
	testConvert(u)
}

func testIder(ider interface{ Ider }) {
	fmt.Println(ider.GetId())
}

func testPassword(pass interface{ Passworder }) {
	fmt.Println(pass.GetPassword())
}

func testConvert(o interface{}) {
	if _, ok := o.(Ider); ok {
		fmt.Println("it's Ider")
	}
	if _, ok := o.(Passworder); ok {
		fmt.Println("it's Pasworder")
	}
	if _, ok := o.(Account); ok {
		fmt.Println("it's Account")
	}
	if _, ok := o.(User); ok {
		fmt.Println("it's User")
	}
}
