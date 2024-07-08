/* Implemente el TAD fecha especificado en el Práctico Uno. ¿La estructura interna
debería poder mutar? Ayuda: Para manejar los topes de días en los meses puede
tener una lista como la siguiente: meses = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31] */

package main

import "fmt";

type Fecha struct {
	day int
	month int
	year int
}

var meses = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func (f *Fecha) setDay(day int) {
	f.day = day
}

func (f *Fecha) setMonth(month int) {
	f.month = month
}

func (f *Fecha) setYear(year int) {
	f.year = year
}

func (f *Fecha) getDay() int {
	return f.day
}

func (f *Fecha) getMonth() int {
	return f.month
}

func (f *Fecha) getYear() int {
	return f.year
}

func (f *Fecha) isLeapYear() bool {
	if f.year % 4 == 0 {
		if f.year % 100 == 0 {
			if f.year % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func (f *Fecha) esValido() bool {
	if f.day > 0 && f.day <= meses[f.month] {
		return true
	}
	if f.month == 2 && f.isLeapYear() && f.day == 29 {
		return true
	}
	return false
}

func main() {
	var fecha Fecha
	fecha.setDay(32)
	fecha.setMonth(3)
	fecha.setYear(2024)
	fmt.Println(fecha.esValido())
}
