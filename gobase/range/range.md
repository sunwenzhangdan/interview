<!--
 * @Date: 2021-05-10 16:31:22
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 16:44:53
 * @FilePath: /面试题/gobase/range/range.md
-->
range 是golang中遍历数据的一种手段，可以遍历数组,slice,map,channel
在迭代切片时可以不适用value 而用slice[index]提高性能

1. 遍历过程中可以适情况放弃接收index或value，可以一定程度上提升性能
2. 遍历channel时，如果channel中没有数据，可能会阻塞
3. 尽量避免遍历过程中修改原数据