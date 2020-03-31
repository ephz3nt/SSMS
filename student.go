package main

import "fmt"

// define student struct
type student struct {
	id          int
	name, class string
}

// define constructor of student
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// define student manager slice struct
type studentMgr struct {
	allStudents []*student
}

// define constructor of studentMgr
func newStudentMgr(newStu student) *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// define add student method of studentMgr
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// define modify student method of studentMgr
func (s *studentMgr) modifyStudent(newStu *student) {
	for k, v := range s.allStudents {
		if v.id == newStu.id {
			s.allStudents[k] = newStu
			return
		}
	}
	fmt.Printf("can't find id = %d student in system...\n", newStu.id)
}

// define show student method of studentMgr
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Println(v.id, v.name, v.class)
	}
}
