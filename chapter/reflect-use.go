package main

import(
	"fmt"
	"reflect"
)

func main(){
	// 将接口类型变量 转化为 反射类型对象（似懂非懂）
	x := 3.12
	// 直接输出: type: float64
	fmt.Println("type:",reflect.TypeOf(x))
	fmt.Println("value:",reflect.ValueOf(x))

	// 反射类型变量 转化为 接口类型变量（完全看不懂）
	var name interface{} = "shirdon"
	fmt.Printf("原始接口变量的类型为%T,值为%v \n",name ,name)

	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为%T \n",t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为%T \n",v)
	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为%T值为%v \n",i,i)

	// 如果要修改 反射类型对象 ，其值必须是可写（云里雾里）
	var myName = "zhaoxiongbin"
	m := reflect.ValueOf(myName)
	n := reflect.ValueOf(&myName)
	// Elem()返回指针所指向的数据
	mn := n.Elem()
	// 判断可写性，用CanSet()
	fmt.Println("m可写性：",m.CanSet())
	fmt.Println("n可写性：",n.CanSet())
	fmt.Println("mn写性：",mn.CanSet())
	mn.SetString("招雄彬")
	fmt.Println("通过反射对象进行更新后，真实myName变为:",myName)
}

