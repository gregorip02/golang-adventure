package main

import (
	"testing"
)

type UserSuite struct {
	a Authorizable
	r bool
}

type AuthorizableMock struct {}
type UnauthorizableMock struct {}

func (a AuthorizableMock) authorize() bool {
	return true;
}

func (a UnauthorizableMock) authorize() bool {
	return false;
}

var suites []UserSuite = []UserSuite {
	{AuthorizableMock{}, true},
	{UnauthorizableMock{}, false},
}

func TestAuthorize(t *testing.T) {
	for _, suite := range suites {
		result := suite.a.authorize();

		if result != suite.r {
			t.Fatalf("Unexpected result, got %t expected %t", result, suite.r)
		}
	}
}
