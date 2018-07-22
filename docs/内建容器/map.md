# map

## map的定义
map是一种键值对的数据结构，通过key可以很方便的(一般认为其时间复杂度为O(1), 实际复杂度取决于于map的具体实现，一般都是会大于O(1)的)。
go语言中直接使用关键字`map`定义: `map[key_type]value_type`
```go
   m1 := map[string]string {
       "key1": "value1",
       "key2": "value2",
       "key3": "value3",
       "key4": "value4",
   }
   m2 := make(map[string]int)
   var m3 = map[string]int
```
## map的操作
1. map的遍历
map可以通过range对key和value进行遍历，具体方式可以类比成数组range遍历的索引和变量。
```go
   for k, v := range m {
       fmt.Println(k, v)
   }
   for _, v := range m {
       fmt.Println(v)
   }
   for k := range m {
       fmt.Println(k)
   }
```
在go语言中，map默认实现是哈希表，遍历时不会按照插入式的顺序。

2. map获取key值
通过`map[key]`可以获取对应key值的value：
```go
   key1 := m[k1]
```
判断key是否存在map中：
```go
   if value1，ok := m[key1]; ok {
       fmt.Println(value1)
   } else {
       fmt.Println("key does not exist")
   }
```
3. 删除map的键值对
使用delete删除指定key的键值对,如删除key值为key1的键值对：
```go
   delete(m, key1)
```
## map的key
* map底层采用hash表实现，key必须要能比较大小；
* 除了slice，map，function的内建类型都可以作为key；
* Struct类型不包含上述字段，也可以作为key。