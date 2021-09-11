package main

import "time"

type Person struct {
	id int
	name string
}

type Authorizable interface {
	authorize() bool
}

type Admin struct {
	Person
}

type User struct {
	Person
}

var findDatabasePermission = func(authorization bool) bool {
	time.Sleep(120 * time.Millisecond)
	return authorization
}

func (a Admin) authorize() bool {
	return findDatabasePermission(true)
}

func (u User) authorize() bool {
	return findDatabasePermission(false)
}
