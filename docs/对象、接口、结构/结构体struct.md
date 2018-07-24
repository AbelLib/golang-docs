# 结构体struct
## 面向对象
经典的面向对象三大特征，go只支持封装，不支持继承和多态，因此go语言没有类class，只有结构体struct。

## struct定义
go语言使用type和struct关键字定义结构体，结构体内只能定义属性，以二叉树的节点定义为例：
```go
type treeNode struct {
    left, right *treeNode
    value int
}
```
与系统内置的类型一样，struct可以通过多种方式进行初始化, go语言没有构造函数和析构函数，如果对象不再使用会被gc回收。
```go
var node1 treeNode
node1 = &treeNode{}
node2 = &treeNode{}
node3 = &treeNode{nil, nil, 4}
node4 = new(treeNode)
```

## struct定义方法
struct的方法不能像类一样定义在struct里面，需要在struct外进行定义： 
```go
func (node treeNode) print() {
    fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
    node.value = value
}
```
如果需要传值引用，需要使用指针。

## 值接收者和指针接收者
* 如果需要改变内容必须使用指针接收者
* 结构过大也考虑使用指针接收者
* 建议是如果有指针接收者，最好都是使用指针
* 值接收者是go语言特有的，其他语言都是指针接收者(Java中被称为引用)
* 函数调用时，会自动进行选择