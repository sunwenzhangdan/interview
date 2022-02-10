<!--
 * @Date: 2021-11-02 19:51:08
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-11-03 11:25:12
 * @FilePath: /interview/gobase/template/template.md
-->
template 最基本的几个函数

func New(name string) *template
创建一个名为name的模版

func (*Template) parse(text string) (*Template error)
将字符串解析为模版，

func(*Template) excute(wr io.Writer, data interface{}) (err error)
将模版渲染输出

```
str := `world`
	tmpl, err := template.New("test").Parse("hello,{{.}}\n")
	err = tmpl.Execute(os.Stdout, str)
```
在这里参数为str,为.的值，模版为parse里面的字符串


```
type Inventory struct {
    Material string
    Count    uint
}

sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}\n")//{{.Count}}获取的struct对象中的Count字段的值
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)//返回 17 of wool
```
这里.为Inventory对象

"{{$str1 := .Say}}
 {{$str2 := .SayHello}}
 {{$str3 := .SayYouName .Name}}
 {{$str1}} {{$str2}}\n{{$str3}}\n"

 注意结构体里面的函数与

 ``与 ""区别
 ``除了{{}}里面的东西其他原本输出 ""有转义



 往模版里面传参

 引用外边普通变量 引用变量
 调用外边函数

 模版里面
 变量
 函数
 结构体
 if语句
 and or !语句
 循环语句


模板嵌套


前面为字段,后面为1为string 2 为int32
`a 1
b 1
c 1`
生成一个结构体
