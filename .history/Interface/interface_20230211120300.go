// interface: mo ta nhung gi can phai lam ma cac class thuc thi interface phai tuan theo
package main

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	emId     int
	basicpay int
	pf       int
}

type Contract struct {
	emId     int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
