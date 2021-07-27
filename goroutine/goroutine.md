<!--
 * @Date: 2021-07-19 15:18:31
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-07-19 16:53:17
 * @FilePath: /面试题/goroutine/goroutine.md
-->
每次go func(){

}就是向队列中提交了一个任务，就是一个生产者，消费者模型，runtime中每次go  func(){},真实的线程就是消费者来消费打包的go func()

P

G

M 

go中线程的数量
每起一个协程就是一个计算任务，

g 一个携程  就是一个计算任务 需要一部分资源

m 真正的线程  真正操作系统的线程  machine 是一个执行实体

P 逻辑调度器   就是一个调度器 


