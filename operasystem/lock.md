<!--
 * @Date: 2021-08-12 14:30:50
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-08-12 15:17:51
 * @FilePath: /interview/operasystem/lock.md
-->
存储数据的媒介

开发人员访问的变量都是放在硬件是内存条的内存里面.

cpu大多数时间在访问相同或者地址相邻的数据，于是在cpu于内存之间加了一个cache 寄存器

分为一级cache  二级cache
但是加了cache的同时，引起了数据的一致性
cache一致性包括三个问题
1 一个cpu核心中指令cache和数据cache的一致性问题
2 多个cpu核心和各自

