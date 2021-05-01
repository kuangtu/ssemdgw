
# 1、说明

本项目是上交所流行情MDGW网关接入解析程序，通过Go语言实现。

上交所通过EzSR和MDGW发布Level1行情，后续会逐步转向MDGW流模式行情。

## 1.2 MDGW行情网关

MDGW（Market Data GateWay）行情网关，固定提供行情任务和文件任务。对于行情接收任务，以Binary或者STEP接口规范进行转发，也可以选择是否落地mktdt格式的行情文件。技术架构如下：

![MDGW技术架构](jpg\MDGW技术架构.jpg)

## 1.3 为什么用Go语言

笔者最近在学习Go语言，学习编程最好的方法就是去实践，在这个过程中”踩坑“和实现各种需求是最”便捷“的一条路。Just Do Coding!



# 2、功能

## 2.1 基本功能

ssemdgw实现的主要功能有：

- 连接MDGW行情网关，完成会话；
- 根据binary接口规范完成行情数据的接收和解析；
- 行情转发，当其他VSS连接ssemdgw之后，ssemdgw会转发行情数据。

## 2.2 基本架构

ssemdgw基本架构如下：

![基本架构](jpg\基本架构.png)



## 2.3 基本模块

（1）系统配置

​	配置MDGW地址和端口；

​	配置本地监听端口，等待其他VSS进行连接。

（2）socket连接

​	ssemdgw启动之后，根据配置的MDGW地址和端口连接MDGW网关；

   根据协议规范建立会话。

（3）会话管理

​	ssemdgw建立会话之后，如果因为网络或者MDGW故障会话断开，可以根据配置的时间间隔重新连接。

（4）数据接收和解析

​	接收MDGW发送的数据，按照接口规范进行解析。

（5）数据保存

​	对于接收到的数据保存原始以及解析后的明文数据。

（6）行情转发

​	其他行情接收模块连接ssemdgw之后，ssemdgw会转发收到的行情数据。



# 3、基本设计  

## 3.1 配置读取

（1）格式

​	json格式，通过Go语言自带的encoding/json解析和处理；

​	定义结构体:



# 附录

[上海证券交易所 行情网关技术指引及接口开发指南](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/policy/c/SSE_MDGW_Interface_0.6_20191111.pdf)



[IS120 上海证券交易所行情网关 BINARY 数据接口规范](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/data/c/IS120_BINARY_Interface_CV0.42_20210315.pdf)





