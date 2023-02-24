package main

import (
	"fmt"
	"sort"
)

var (
	name string
)

type User struct {
	Login    string
	Password string
}

type Item struct {
	Name        string
	Price       float64
	Rating      float64
	NumOfRates  int
	Description string
}

var users = []User{
	{"moldaspan", "ers123"},
	{"danialbratishka", "bratishka123"},
	{"bahagul", "baha123"},
	{"madimadiyarov", "madi123"}}
var items = []Item{
	{"NIKE", 36500, 4, 1, "sportswear"},
	{"Adidas", 45000, 9, 2, "sneakers"},
	{"Puma", 31990, 9.8, 2, "t-shirt"},
	{"Reebok", 20990, 4.0, 1, "tights"}}

func Registration(login, password string) {
	for i := 0; i < len(users); i++ {
		if users[i].Login == login {
			fmt.Println("User already exists")
			return
		}
	}
	users = append(users, User{Login: login, Password: password})
}
func Authorization(login, password string) *User {
	for i := 0; i < len(users); i++ {
		if users[i].Login == login && users[i].Password == password {
			return &users[i]
		}
	}
	return nil
}

func Searching(name string) *Item {
	for i := 0; i < len(items); i++ {
		if items[i].Name == name {
			return &items[i]
		}
	}
	return nil
}

func FilteringByPrice() {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Price < items[j].Price
	})
}
func FilteringByRating() {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Rating < items[j].Rating
	})
}

func (u *User) GiveRate(name string, rate float64) {
	for i := 0; i < len(items); i++ {
		if items[i].Name == name {
			items[i].Rating = (items[i].Rating*float64(items[i].NumOfRates) + rate) / float64(items[i].NumOfRates+1)
			items[i].NumOfRates++
		}
	}

}

func main() {
	for true {
		fmt.Println("1.Registration\n2.Authorization")
		var choise int
		fmt.Scan(&choise)

		fmt.Println("Enter login: ")
		var login string
		fmt.Scan(&login)

		fmt.Println("Enter password: ")
		var password string
		fmt.Scan(&password)
		var u *User = nil
		if choise == 1 {
			Registration(login, password)
		} else {
			u = Authorization(login, password)
			if u == nil {
				fmt.Println("Wrong login or password")
			} else {
				fmt.Println("Authorised.")
				for true {
					fmt.Println("\n3.Searching")
					fmt.Println("4.Filtering by price")
					fmt.Println("5.Filtering by rating")
					fmt.Println("7.exit")

					var temp string
					fmt.Scan(&temp)
					if temp == "3" {
						fmt.Println("Enter item name:")
						fmt.Scan(&name)
						var it *Item = Searching(name)
						if it == nil {
							fmt.Println("Not found")
						} else {

							fmt.Println(it)
							fmt.Println("Giving Rating")
							var rate float64

							fmt.Scan(&rate)

							u.GiveRate(name, rate)
							for i := 0; i < len(items); i++ {
								fmt.Println(items[i])
							}
						}
					} else if temp == "4" {
						FilteringByPrice()
						for i := 0; i < len(items); i++ {
							fmt.Println(items[i])
						}
					} else if temp == "5" {

						FilteringByRating()
						for i := 0; i < len(items); i++ {
							fmt.Println(items[i])
						}
					} else if temp == "7" {
						break
					} else {
						fmt.Println("Wrong option!")
					}

				}
			}
		}
	}
}
