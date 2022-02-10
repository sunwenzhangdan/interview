<!--
 * @Date: 2021-12-22 14:21:52
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-12-22 14:32:37
 * @FilePath: /interview/消息队列/kafka.md
-->
kafka  rocketmq
kafka 由broker组成，每个broker是一个节点，创建一个topic，这个topic可以划分为partition,每个partition可以存在于不同broker上，
如何保证不重复消费，重复消费不可怕，如何保证幂等性


将消费的数据放到redis
如何防止消息队列丢消息，手动确认 rocketmq事务消息
保证消息的顺序性

