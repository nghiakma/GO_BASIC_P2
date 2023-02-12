// De tao ra 1 map thong qua type of key and value to make function
// make(map[type of key]type of value)
package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["steve"] = 12000
	employeesalary["jamie"] = 15000
	employeesalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeesalary)

	employee := "jamie"
	salary := employeesalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	//checking if a key exists (tontai)

	newEmp := "joe"

	value, ok := employeesalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	//Iterate over(lap lai) all elements in a map
	fmt.Println("Contents of the map")
	for key, value := range employeesalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}
	//xxoa phan tu trong map
	delete(employeesalary, "steve")
}
