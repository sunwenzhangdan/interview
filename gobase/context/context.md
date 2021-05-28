<!--
 * @Date: 2021-05-12 11:03:08
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-12 20:00:52
 * @FilePath: /面试题/gobase/context/context.md
-->
为什么要引入context包呢，想象一下这样的一个情况
如果用户发起一个http请求，http调用了一个中间件，但是用户在请求后断开了请求，这样我们已经请求了中间件，但是如何中断呢
协程之间没有父子关系，不会父协程退出导致子协程退出，所以在父协程内没法控制子协程。所以，我们现在需要解决这个问题。


WithCancel：基于父级 context，创建一个可以取消的新 context。
WithDeadline：基于父级 context，创建一个具有截止时间(Deadline)的新 context。
WithTimeout：基于父级 context，创建一个具有超时时间(Timeout)的新 context。
Background：创建一个空的 context，一般常用于作为根的父级 context。


context是一个接口，里面有四个方法

