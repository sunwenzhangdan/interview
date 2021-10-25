<!--
 * @Date: 2021-07-26 13:35:09
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-08-03 11:41:20
 * @FilePath: /interview/docker/docker.md
-->
使用docker将一些常用的工具构建成镜像
dockerfile 文件编写
FROM 基于是哪个镜像
ADD  user /user 将文件拷贝到目录
ENTRYPOINT 启动命令

docker images 对镜像的管理
docker rmi -f 删除镜像
docker tag 镜像标签
docker  images prune 

docker create   是根据镜像创建一个容器，但是并没有启动
docker  run=create+start
docker 
容器所在系统的多个端口可以绑定容器的一个接口


