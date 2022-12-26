# GO Note  

## 基础  

### 在编程之前  

[Go wiki](https://learnku.com/go/wikis/26447)  
[Go Bible](https://books.studygolang.com/gopl-zh/)  

搜寻一下可用库，看看是否能简化你的工作（[pkg.go](https://pkg.go.dev/)和[golang.org](https://golang.org/pkg)）  PS:由于某种不可抗力，可能会进不去  

### Golang main function  

```go
package main

func main() {
  // variables and functions etc.
}
```  

不同于其他的一些编程语言，Golang的main函数无需任何参数和返回值  

### Printf  

Printf有一大堆这种转换，Go程序员称之为动词（verb）。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：  

```bash
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```  

### 并发  

当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑。在[例子](./5-fetch_multi_url.go)中，每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。

### 使用:=和=  

[Go圣经章节](https://books.studygolang.com/gopl-zh/ch2/ch2-03.html)

~~目前还不太了解，~~ go中可以用`:=`也可以用`=`赋值，但是使用`=`有些时候报`undefined`错，用`:=`就可以了  

可以看看关于[:=和=的区别](https://blog.csdn.net/wei242425445/article/details/88390605)，简单概括就是，使用`:=`的语句中要包含一个**新变量**。当然这只是表象，官方一点的说法是，`:=`用于简短变量声明，其语句中必须有一个新变量，这样的变量没有必要要在语句中显示指定类型，多用于定义局部变量，而使用`=`要显示地指定变量类型，或者声明与赋值分开写  

>因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。var形式的声明语句往往是用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。请记住“:=”是一个变量声明语句，而“=”是一个变量赋值操作。也不要混淆多个变量的声明和元组的多重赋值（§2.4.1），后者是将右边各个表达式的值赋值给左边对应位置的各个变量：

```go
i, j = j, i // 交换 i 和 j 的值 
```

>简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。  

### Go返回局部变量的地址是合法操作  

go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。所以不用担心会不会导致memory leak，因为GO语言有强大的垃圾回收机制。go语言声称这样可以释放程序员关于内存的使用限制，更多的让程序员关注于程序功能逻辑本身。  

## 注意事项  

使用cmd执行程序，用powershell事会把\\n替换为\\r\\n，这样unit3就没法生成在Windows下可见的gif图了  
