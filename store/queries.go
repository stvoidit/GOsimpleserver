package store

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// AllUsers - ...
func (s *Store) AllUsers(chu []string) []User {
	row, err := func() (*sql.Rows, error) {
		if len(chu) > 0 {
			data, err := s.Session.Query(`
			SELECT id, "FirstName", "MiddleName","LastName", "DepartmentId", position 
			FROM public.users WHERE id = ANY($1)`, pq.Array(chu))
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			return data, nil
		}
		data, err := s.Session.Query(`
			SELECT id, "FirstName", "MiddleName","LastName", "DepartmentId", position 
			FROM public.users`)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return data, nil
	}()
	if err != nil {
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

// AllDepartments - ...
func (s *Store) AllDepartments() []Department {
	row, err := s.Session.Query(`
	SELECT id, "Name", isgeneral, maindepartment 
	FROM public.departments`)
	if err != nil {
		fmt.Println(err.Error())
		return make([]Department, 0)
	}
	var departmentsArray []Department
	for row.Next() {
		var d Department
		row.Scan(&d.ID, &d.Name, &d.IsGeneral, &d.Maindepartment)
		departmentsArray = append(departmentsArray, d)
	}
	return departmentsArray
}

// AllUsersDepartments - ...
func (s *Store) AllUsersDepartments() []UserDepartments {
	row, err := s.Session.Query(`
	SELECT u.id, u."FirstName", u."MiddleName",u."LastName", u."DepartmentId", u.position, d.id, d."Name",d.isgeneral, d.maindepartment 
	FROM public.users as u
	left join public.departments as d on d.id = u."DepartmentId" `)
	if err != nil {
		fmt.Println(err.Error())
		return make([]UserDepartments, 0)
	}
	var userdepsArray []UserDepartments
	for row.Next() {
		var ud UserDepartments
		row.Scan(&ud.User.ID, &ud.FirstName, &ud.MiddleName, &ud.LastName, &ud.DepartmentID, &ud.Position, &ud.Department.ID, &ud.Name, &ud.IsGeneral, &ud.Maindepartment)
		userdepsArray = append(userdepsArray, ud)
	}
	return userdepsArray
}
