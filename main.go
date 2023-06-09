package main

import (
	"fmt"
	"time"
)

// EmployeeRepository interface untuk mengelola data karyawan
type EmployeeRepository interface {
	AddEmployee(employee Employee)
	RemoveEmployee(employeeID int)
	FindEmployeeByID(employeeID int) *Employee
}

// Employee struct untuk merepresentasikan data karyawan
type Employee struct {
	ID        int
	Name      string
	Position  string
	CreatedAt time.Time
}

// InMemoryEmployeeRepository struct untuk mengelola data karyawan
type InMemoryEmployeeRepository struct {
	employees []Employee
}

// AddEmployee metode untuk menambahkan karyawan baru
func (er *InMemoryEmployeeRepository) AddEmployee(employee Employee) {
	er.employees = append(er.employees, employee)
}

// RemoveEmployee metode untuk menghapus karyawan berdasarkan ID
func (er *InMemoryEmployeeRepository) RemoveEmployee(employeeID int) {
	for i, employee := range er.employees {
		if employee.ID == employeeID {
			er.employees = append(er.employees[:i], er.employees[i+1:]...)
			break
		}
	}
}

// FindEmployeeByID metode untuk mencari karyawan berdasarkan ID
func (er *InMemoryEmployeeRepository) FindEmployeeByID(employeeID int) *Employee {
	for _, employee := range er.employees {
		if employee.ID == employeeID {
			return &employee
		}
	}
	return nil
}

// AttendanceManager interface untuk mengelola absensi karyawan
type AttendanceManager interface {
	RecordAttendance(employeeID int)
}

// AttendanceService struct untuk mengelola absensi karyawan
type AttendanceService struct {
	employeeRepository EmployeeRepository
}

// RecordAttendance metode untuk merekam absensi karyawan
func (as *AttendanceService) RecordAttendance(employeeID int) {
	employee := as.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

func main() {
	employeeRepo := &InMemoryEmployeeRepository{}
	attendanceService := &AttendanceService{employeeRepository: employeeRepo}

	employee1 := Employee{ID: 1, Name: "John Doe", Position: "Manager", CreatedAt: time.Now()}
	employee2 := Employee{ID: 2, Name: "Jane Smith", Position: "Staff", CreatedAt: time.Now()}

	employeeRepo.AddEmployee(employee1)
	employeeRepo.AddEmployee(employee2)

	attendanceService.RecordAttendance(1)
	attendanceService.RecordAttendance(2)
	attendanceService.RecordAttendance(3) // Karyawan dengan ID 3 tidak ditemukan
}
