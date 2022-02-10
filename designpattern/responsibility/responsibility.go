/*
 * @Date: 2021-12-01 15:05:17
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-12-01 15:16:21
 * @FilePath: /interview/设计模式/责任链模式/responsibility.go
 */
package reponsibility

//  责任链模式就是几个相同的handler


//type Patient struct {
//	name string
//	doctorCheckupDone bool
//	nurseCheckUpDone  bool
//	xrapCheckUpDone bool
//}
//
//func NewPatient(name string, doctorCheckupDone bool) *Patient {
//	return &Patient{name: name, doctorCheckupDone: doctorCheckupDone}
//}
//
//type Doctor struct {
//	Handlers
//}
//
//type Nurse struct {
//   Handlers
//}
//
//type BaoAn struct {
//	Handlers
//}
//
//
//type Handlers interface {
//	Execute(*Patient)
//	SetNext(Handlers)
//}
//
//func (d Doctor) Execute(patient *Patient) {
//	patient.doctorCheckupDone=true
//	d.Handlers.Execute(patient)
//	return
//}
//
//func (d Doctor) SetNext(handlers Handlers) {
//	d.Handlers=handlers
//}
//
//func (n Nurse) Execute(patient *Patient) {
//	patient.nurseCheckUpDone=true
//	n.Handlers.Execute(patient)
//	return
//
//}
//
//func (n Nurse) SetNext(handlers Handlers) {
//	n.Handlers=handlers
//}
//
//func (b  BaoAn) Execute(patient *Patient) {
//	patient.xrapCheckUpDone=true
//	b.Handlers.Execute(patient)
//	return
//}
//
//func (b BaoAn) SetNext(handlers Handlers) {
//	b.Handlers=handlers
//}
//
