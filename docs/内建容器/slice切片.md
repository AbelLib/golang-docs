# slice(切片)

## slice的定义
slice是针对数组的封装，数据就是存在放在内部的数组中的。个人感觉可以类比成Java nio中的buffer(从实现的角度来看)，当然在python中也提供了slice，两者的意思基本是相同的。

## slice的底层实现
slice由三个变量组成：ptr，len，cap。
* ptr就是上面说到的存储数据的数组；
* cap是ptr数组的长度；
* len是slice有效数组的长度，即slice包含元素的长度。
需要注意的是：
s[i]访问时不能超过len(s)，这个是我们可以理解的；
slice可以向后扩展，不可以向前扩展，直到cap(s),也就是说`s[low, high]` 中high的取值可以大于len(s)但是不能大于cap(s), 也就是说向后扩展不能越过ptr数组的长度cap(s)。

## slice的操作
1. 创建slice
slice的定义与数组基本类似，最主要的区别是slice定义时是不需要确定长度的;
slice可以直接定义，也可以由数组中截取元素使用`[low:high]`定义；
slice可以继续进行切片生成新的slice。
```go
    s := []int{6, 7, 8, 9, 10}
    arr := [5]int{1, 2, 3, 4, 5}
    s1 := arr[2:6]
    s2 := s[:]
```
2. make
使用make可以创建指定长度的slice，指定的长度值就是slice的capacity。
创建一个长度为10的slice：
```go
    a := make([]int, 10)
```
3. 添加元素-append
```go
    s3 := append(s2, 10)
```
添加元素时如果超过cap，系统会重新分配一个更大的底层数组，并且将slice的ptr变量指向这个新的数组。
由于在执行append的过程中可能会导致底层数组的引用改变，所以属性ptr，len和cap都可能会改变，因此append是必须有一个返回值来接受append的结果的，否则函数的结构会丢失。  
4. 复制slice-copy
将一个slice的内容复制到另一个slice中，如果两个slice大小不同，以长度短的那个为准，复制分别从slice第一个元素开始。
```go
    copy(s1, s2)
```
5. 删除slice的元素
标准库中没有提供删除单个元素的接口，可以通过append函数实现：
```go
    s = append(s[:3], s[4:]...)
```
上面的程序表示删除s[3]这个元素，并且将slice的后面的元素前移。