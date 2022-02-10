/*
 * @Date: 2021-12-01 15:16:46
 * @LastEditors: seven sun
 * @LastEditTime: 2021-12-01 15:16:46
 * @FilePath: /interview/设计模式/责任链模式/responsibility_test.go
 */

package reponsibility

import (
	"fmt"
	"testing"
)


func TestRes(t *testing.T){
	fmt.Println("dfdf")
}

//go也是面向接口编程，将动作抽象

func TestNewPatient(t *testing.T) {
    p:=&Patient{}
	n:=&Nurse{}
	b:=&BaoAn{}
	d:=&Doctor{}
	d.SetNext(b)
	n.SetNext(d)
	n.Execute(p)

	fmt.Println(p)
}

type Patient struct {
	name string
	doctorCheckupDone bool
	nurseCheckUpDone  bool
	xrapCheckUpDone bool
}

func NewPatient(name string, doctorCheckupDone bool) *Patient {
	return &Patient{name: name, doctorCheckupDone: doctorCheckupDone}
}

type Doctor struct {
	handlers Handlers
}

type Nurse struct {
	handlers Handlers
}

type BaoAn struct {
	handlers Handlers
}


type Handlers interface {
	Execute(*Patient)
	SetNext(Handlers)
}

func (d *Doctor) Execute(patient *Patient) {
		patient.doctorCheckupDone = true
	d.handlers.Execute(patient)
	return
}

func (d *Doctor) SetNext(handlers Handlers) {
	d.handlers=handlers
}

func (n *Nurse) Execute(patient *Patient) {
		patient.nurseCheckUpDone=true
	n.handlers.Execute(patient)
	return

}

func (n *Nurse) SetNext(handlers Handlers) {
	n.handlers=handlers
}

func (b  *BaoAn) Execute(patient *Patient) {
		//patient.xrapCheckUpDone = true
	return
}

func (b *BaoAn) SetNext(handlers Handlers) {
	b.handlers=handlers
}

