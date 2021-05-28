<!--
 * @Date: 2021-05-08 18:21:48
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 17:24:24
 * @FilePath: /面试题/gobase/channel/channel.md
-->
面向面试学习
就想着面试，想着给别人介绍channel

channel的数据结构
closed  chan是否关闭
elemtype chan里面的数据类型
buf 循环队列
sendx 发送索引
recvx 接收索引
sendq 发送队列
recvq 接收队列
lock  锁

chan 的基本原理
chan的数据结构是固定长度的双向循环列表
按顺序往里面写，写满继续从0开始
chan中最重要的两个组件buf双向循环队列，还有waitq就是sendq和recvq 

chan的创建




发送数据到chan
1. 有等待接收的groutinue,如果待接收队列中有等待的接收者的话，说明 channel 的缓冲区为空
   如果recvq中有等待接收数据的groutinue，将把发送的数据给groutinue并且被唤醒
2. 无等待接收的groutinue且环形队列buf未满 将把发送的数据放入buf队列，更新元素个数以及发送数据索引 
3. 有等待接收的groutinue且环形队列buf已满 将数据与发送数据的groutinue打包成sudog，加入等待发送数据队列的尾部 



从chan中读数据
1. 有等待发送的chan数据，如果待发送队列中有等待发送的 goroutine，说明 channel 是非缓冲channel，或者缓冲区已经满了
   非缓冲channel，会将数据从待发送者复制给接收者
   缓冲区已满的话，会先从缓冲区中接收数据，然后将待发送者的数据发送到缓冲区中
   与上面的区别：
   这里和发送操作时，channel 的待接收队列不为空的情况不一样，因为待接收队列不为空，说明缓冲区肯定是没有数据的，可以跳过缓冲区，直接将数据发送到等待接收的 goroutine
2. 就是先从环形队列中读取数据，然后将待发送的数据从队列中拷贝到环形队列中
3. 如果没有数据则放到recv队列中






从chan中读数据


channel的使用
channel分为以下几类
零值通道
非零值但已经关闭的通道
非零值尚未关闭的通道






