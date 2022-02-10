<!--
 * @Date: 2021-11-02 11:32:44
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-11-02 12:12:26
 * @FilePath: /interview/gobase/generator_wire/wire.md
-->
go中的依赖注入
go可以在运行时依赖注入，但是这样性能不好。于是采用编译时依赖注入，提前生成号代码。

其实目的就是自己不写new struct()代码了，代码用wire生成

