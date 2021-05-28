<!--
 * @Date: 2021-05-10 17:10:33
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 17:18:38
 * @FilePath: /面试题/gobase/select/select.md
-->
select功能
在多个通道上进行读或写操作，让函数可以处理多个事情，但1次只处理1个。以下特性也都必须熟记于心：

每次执行select，都会只执行其中1个case或者执行default语句。
当没有case或者default可以执行时，select则阻塞，等待直到有1个case可以执行。
当有多个case可以执行时，则随机选择1个case执行。
case后面跟的必须是读或者写通道的操作，否则编译出错。
select长下面这个样子，由select和case组成，default不是必须的，如果没其他事可做，可以省略default。
select语句中除default外，每个case操作一个channel，要么读要么写
select语句中除default外，各case执行顺序是随机的
select语句中如果没有default语句，则会阻塞等待任一case
select语句中读操作要判断是否成功读取，关闭的channel也可以读取

select{} 用于阻塞

select 后面跟的是从channel中读或者是将数据放入channel

