package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID int
	Name string
	Email string
	Role string
}

var allowedRoles = []string {"client", "freelancer"}

var allowedEmails = []string { "munachi.sosa@comcast.net", "kfvkcugq3fuhpqkv0j8s@hotmail.com", "lepeng.herrera@ymail.com", "y9efurls16@rediffmail.com", "nikodem.simon@aol.com", "v5vyril4y77@hotmail.com", "jayden.alston@msn.com","k8hmk8k9wvs1p7ec9@rediffmail.com","louie.sanders@hotmail.com", "dc9aigt2mto@ymail.com" }

var names = []string { "Антон" , "Виталий", "Елена", "Константин", "Дмитрий", "Юрий", "Ольга", "Борис", "Илья", "Юлия" }

func getRandomElem[T any](list []T) T {
	listLength := len(list)

	if listLength == 0 {
		return list[0]
	}

	randomIdx := rand.Intn(listLength)
	return list[randomIdx]
}

func createRandomUsers() []User {
	fmt.Println("Generating users...")

	var users []User

	for idx, email := range allowedEmails {
		name := getRandomElem(names)
		role := getRandomElem(allowedRoles)

		var user = User{
			ID: idx + 1,
			Name: name,
			Role: role,
			Email: email,
		}

		users = append(users, user)
	}

	return users
}

func main() {

	var users = createRandomUsers()

	fmt.Println(users)

	total := len(users)
	var freelancers int
	var clients int

	for _, user := range users {
		
		switch {
		case user.Role == "freelancer":
			freelancers++
		case user.Role == "client":
			clients++
		}
	}

	fmt.Printf(`User stats: 
Total users: %d, 
Freelancers: %d,
Clients: %d`, total, freelancers, clients)
}