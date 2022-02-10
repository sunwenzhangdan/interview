/*
 * @Date: 2021-11-02 19:50:00
 * @LastEditors: seven sun
 * @LastEditTime: 2021-11-04 11:47:16
 * @FilePath: /interview/gobase/template/template_test.go
 */
package template_test

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

type StrAndFunc struct {
	Name   string
	module string
	Funcc  func(string) string
}

func TestTemplate(t *testing.T) {
	t.Log("hello")
	strAndf := StrAndFunc{
		Name: "sunshiwen",
		Funcc: func(world string) string {
			return "hello" + world
		},
	}
	t.Log(strAndf.Funcc("dfdf"))
	funcMap := map[string]interface{}{"mkSlice": mkSlice}
	// tc := `
	// {{"************"}}
	// {{$a = mkSlice "abc"}}
	// {{ range $i, $v := $a}}
	// {{$i}} {{$v}}
	// {{end }}
	// {{"^^^^^^^^^^^^^^^"}}
	// `
	ui := `
	{{23}} < {{45}}
	{{23}} < {{- 45}}
	{{23 -}} < {{45 }}
	{{- 23 -}} < {{- 45}}
	`
	tmpl := template.New("test").Funcs(template.FuncMap(funcMap))
	template.Must(tmpl.Parse(ui))
	tmpl.Execute(os.Stdout, strAndf)
}

func SayYouName(name string) string { // 有参数
	return "my name is : " + name
}

func mkSlice(args ...interface{}) []interface{} {
	return args
}

func TestModleByTmpl(t *testing.T) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	csc := strings.Fields(string(data))
	fmt.Println(csc)
	fmt.Println(len(csc))
	csMap := make(map[string]interface{})
	returnType := func(typed string) string {
		var typec string
		switch typed {
		case "1":
			typec = "string"
		case "2":
			typec = "int32"
		}
		return typec
	}
	for i := 0; i < len(csc); i += 2 {
		csMap[csc[i]] = returnType(csc[i+1])
	}

	p1 := Pc{Name: "user", Module: "model", Mc: csMap}
	fmt.Println(csMap)
	var b *template.Template
	b, err = template.ParseFiles("./test.tmpl")
	dstFile, err := os.Create("./test.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	b.Execute(dstFile, p1)
}

type Pc struct {
	Name   string
	Module string
	Mc     map[string]interface{}
}
