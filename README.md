
# 1、说明

本项目是上交所流行情MDGW网关接入解析程序，通过Go语言实现。

上交所通过EzSR和MDGW发布Level1行情，后续会逐步转向MDGW流模式行情。

## 1.2 MDGW行情网关

MDGW（Market Data GateWay）行情网关，固定提供行情任务和文件任务。对于行情接收任务，以Binary或者STEP接口规范进行转发，也可以选择是否落地mktdt格式的行情文件。技术架构如下：

![MDGW技术架构](jpg\MDGW技术架构.jpg)

## 1.3 为什么用Go语言

笔者最近在学习Go语言，学习编程最好的方法就是去实践，在这个过程中”踩坑“和实现各种需求是最”便捷“的一条路。Just Do Coding!



## #2、功能

ssemdgw实现的主要功能有：

- 连接MDGW行情网关，完成会话；
- 根据规范完成行情数据的接收和解析；
- 选择



# 附录

[上海证券交易所 行情网关技术指引及接口开发指南](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/policy/c/SSE_MDGW_Interface_0.6_20191111.pdf)



[IS120 上海证券交易所行情网关 BINARY 数据接口规范](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/data/c/IS120_BINARY_Interface_CV0.42_20210315.pdf)





