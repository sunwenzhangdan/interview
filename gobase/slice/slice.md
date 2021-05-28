<!--
 * @Date: 2021-05-08 14:07:55
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 08:37:26
 * @FilePath: /面试题/gobase/slice/slice.md
-->
1. 切片的创建
   

2. 切记另一点在go语言中所有的传参都是传拷贝只不过引用类型传的是指针的拷贝

3. 传指针为什么能改变呢，因为指针发生了拷贝，通过拷贝的指针修改了值，但是等于被拷贝的指针指向的东西也发生了变化

4. 在此要引出make与new的区别
new是用来做什么呢
对于基本类型的值，我们初始化话就自动赋了0值，但是对于引用类型的变量，我们需要自己申请一块该类型的内存并返回指向该内存的指针
new(T) 的操作是从内存申请一个T类型内存大小的空间，并且一个指针指向这个空间，把这个指针返回来这就是new干的事,new只清除内存并不初始化
make(T) make函数只适用于slice map channel 在创建map的时候就是调用的makemap函数，make就是map的工厂函数，根据不同的kv创建不同类型的map
同时初始化map的大小。

5. slice 基于数组，就是slice有一个指向数组的指针，所以说slice的底层数据是相等的，除非底层数据空间不够才会去扩容,容量是从数组的首地址到数组的末地址
   利用数组生成切片array[start:end]

6. append函数扩容
扩容容量的选择遵循以下规则：
如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍；
如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；
使用append()向Slice添加一个元素的实现步骤如下：
假如Slice容量够用，则将新元素追加进去，Slice.len++，返回原Slice
原Slice容量不够，则将Slice先扩容，扩容后得到新Slice
将新元素追加进新Slice，Slice.len++，返回新的Slice。

7. 切片的拷贝



8. 切片的迭代



9. 切片的删除



10. 特殊切片
对于切片一般都是[start:end] 长度是end-start 容量是到数组的最后，但是我们也可以传参数直接指定，切片的容量
sliceA:=make(chan int,5,10) 创建一个长度为5容量为10的切片
sliceB:=sliceA[0:5]  sliceB的长度是5容量还是为10
sliceC:=sliceA[0:5:5] sliceC的长度是5容量也是5










