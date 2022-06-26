package main

import "fmt"

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// Task 1
func NewResident(name string, age int, address map[string]string) *Resident {
	resident := Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return &resident
}

// Task 2
func (r *Resident) HasRequiredInfo() bool {
	if r.Name != "" {
		fmt.Println(r.Address)
		if r.Address != nil {
			for _, value := range r.Address {
				if value != "" {
					return true
				}
			}
		}
	}
	return false
}

// Task 3
func (r *Resident) Delete() {
	r.Address = nil
	r.Age = 0
	r.Name = ""
}

// Task 4
func Count(residents []*Resident) int {
	var counter int = 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			counter++
		}
	}
	return counter
}
