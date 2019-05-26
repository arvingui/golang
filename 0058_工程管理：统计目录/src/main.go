//1、分文件编程（多个源文件），必须放在src目录
//2、设置GOPATH环境变量
//3、同一个目录下的源文件（*.go文件），包名必须一样
//4、go env 查看go相关环境路径，直接调用即可，无需包名引用
//5、同一个目录，调用别的文件

//package main

//忽略此包
/*import _ "fmt"

func main () {

}*/

//给包起别名
/*import io "fmt" //io相当于fmt的别名，可以用io代替fmt

func main () {
	io.Println("this is s test")
}*/

package main

func main() {
	test() //直接调用，无需包名引用
}
