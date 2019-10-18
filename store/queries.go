package store

import "fmt"

// AllUsers - ...
func (s *Store) AllUsers() []User {
	row, err := s.Session.Query(`
	SELECT id, "FirstName", "MiddleName","LastName", "DepartmentId", position 
	FROM public.users`)
	if err != nil {
		fmt.Println(err.Error())
		return make([]User, 0)
	}
	var usersArray []User
	for row.Next() {
		var u User
		row.Scan(&u.ID, &u.FirstName, &u.MiddleName, &u.LastName, &u.DepartmentID, &u.Position)
		usersArray = append(usersArray, u)
	}
	return usersArray
}
