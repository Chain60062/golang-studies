package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee //Embedded field
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// Mock database of employees (in a real app this might come from a DB)
	candidates := []Employee{
		{ID: "E004", Name: "Alice"},
		{ID: "E005", Name: "Bob"},
		{ID: "E006", Name: "Charlie"},
	}

	// Simulated logic: return employees who aren't already in Reports
	existingIDs := make(map[string]bool)
	for _, report := range m.Reports {
		existingIDs[report.ID] = true
	}

	var newEmployees []Employee
	for _, c := range candidates {
		if !existingIDs[c.ID] {
			newEmployees = append(newEmployees, c)
		}
	}

	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{Name: "Vinicius", ID: "12345"},
		Reports:  []Employee{},
	}
	fmt.Println(m.Name) //you can direct access fields from a embedded field (Employee)
	fmt.Println(m.Description())
}
