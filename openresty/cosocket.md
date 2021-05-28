<!--
 * @Date: 2021-05-10 18:13:12
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-05-10 18:13:51
 * @FilePath: /面试题/openresty/cosocket.md
-->

cosocket是 OpenResty 中的专有名词，是把协程和网络套接字的英文 拼在一起形成的，即 cosocket = coroutine + socket。所以，可以把 cosocket 翻译为“协程套接字”。
cosocket 不仅需要 Lua 协程特性的支持，也需要 Nginx 中非常重要的事件机制的支持，这两者结合在一 起，最终实现了非阻塞网络 I/O。另外，cosocket 支持 TCP、UDP 和 Unix Domain Socket。
在 OpenResty 中调用一个 cosocket 相关函数，内部实现便是下面这张图的样子：

用户的 Lua 脚本每触发一个网络操作，都会有协程的 yield 以及 resume。
遇到网络 I/O 时，它会交出控制权（yield），把网络事件注册到 Nginx 监听列表中，并把权限交给 Nginx；当有 Nginx 事件达到触发条件时，便唤醒对应的协程继续处理（resume）。
OpenResty 正是以此为基础，封装实现 connect、send、receive 等操作，形成了现在的 cosocket API。以处理 TCP 的 API 为例来介绍一下。处理 UDP 和 Unix Domain Socket ，与TCP 的接口基 本是一样的。

