package models

import "fmt"

/*User struct holds user */
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

/*GetUsers return users */
func GetUsers() []*User {
	return users
}

/*AddUser add a user*/
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

/*TestPointer The pointer test and amap */
func TestPointer() {
	amap := map[string]int{"foo": 42}
	var fistName *string = new(string)
	*fistName = "Arthur"
	fmt.Println(*fistName)
	fmt.Println(amap)
}
