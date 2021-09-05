package main

type Post struct {
	id int
	content string
}

type User struct {
	id int
	name string
	password string
}

func (user User) friends() int {
	return len(user.name)
}

func (user User) posts() []Post {
	return []Post {
		{id: 4243, content: "Hello world"},
		{id: 7655, content: "Golang is awesome"},
	}
}

type Admin struct {
	user User
	permissions []string
}

func main() {
	var admin Admin = Admin{
		user: User{name: "Gregori"},
		permissions: []string{"one", "two"},
	}

	println(admin.user.friends())
	println(len(admin.user.posts()))
}
