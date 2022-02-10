/*
 * @Date: 2021-12-01 15:05:17
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-12-01 15:16:21
 * @FilePath: /interview/设计模式/责任链模式/responsibility.go
 */
package reponsibility

//  责任链模式就是几个相同的handler


type Patient struct {
	name string
	doctorCheckupDone bool
}

func NewPatient(name string, doctorCheckupDone bool) *Patient {
	return &Patient{name: name, doctorCheckupDone: doctorCheckupDone}
}

type Doctor struct {
	Handlers
}


type Handlers interface {
	Execute(*Patient)
	SetNext(Handlers)
}
