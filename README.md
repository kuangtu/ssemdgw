
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

### 3.1.1 格式

​	json格式，通过Go语言自带的encoding/json解析和处理；

​	定义结构体:

```go

type SysConf struct {
    Gatewayip   string `json:"GateWayIP"`
    Gatewayport int    `json:"GateWayPort"`
    Localip     string `json:"LocalIP"`
    Localport   int    `json:"LocalPort"`
    Backdir     string `json:"BackDir"`
}
```



# 3.2 网络处理

https://www.infoq.cn/article/boeavgkiqmvcj8qjnbxk



### 3.2.1 连接MDGW网关

（1）连接过程

- 通过TCP协议连接MDGW网关， ssemdgw发起TCP请求，如果连接失败，等待配置时间间隔后重新连接；

- 如果TCP连接建立之后因网络问题出现中断，也等待配置时间间隔后重新发起连接；
- 或者在配置的时间内收不到消息。

（2）配置项

- 超时时间（timeout），TCP接收超时时间配置；
- 连接时间间隔（conn_inteval），重新发起连接的间隔时间。

（3）实现方法

​	通过golang语言net库实现连接。

### 3.2.2 会话管理

​	MDGW行情网关协议如下：

![协议交互](jpg\协议交互.png)

（1）登录过程

- ssemdgw完成TCP连接之后，发送登录消息，然后接收登录验证消息，如果验证失败（可能是由于用户名、密码等验证错误），解析注销消息获取原因；
- 如果登录成功，解析登录成功消息。

（2）心跳消息

- 用于监控通信连接的状况；
- 当连接的任何一方在心跳时间间隔（由 HeartBtInt 域指定）时间内没有接收或发送任何数据的时候，需要产生一个心跳消息并发送出去；
- 如果接收方在 2 倍心跳时间间隔内都没有收到任何消息的时候，那么可认为行情会话出现异常，可以立即关闭 TCP 连接。

（3）注销消息

- 发起或者确认行情会话终止；
- 未经注销消息交换而断开连接，一律视为非正 常的断开。

（4）消息组成

- 每条消息有消息头、消息体和消息尾组成，消息最大长度为8K字节。
-  头部格式：

![消息头部](jpg\消息头部.png)

- 消息尾:

![消息尾](jpg\消息尾.png)



（5）消息验证算法

```c
uint32 CalcChecksum(const char* buffer, uint32 len)
{
 uint8 checksum = 0;
 uint32 i = 0;
 for (i = 0; i < len; i++)
 {
 checksum += (uint8)buffer[i];
 }
 return (uint32)checksum;
}
```



（2）主要消息体结构设计

​	参照《规范》中的消息字段的类型：

| 类型      | 说明                                                         |
| --------- | ------------------------------------------------------------ |
| char[x]   | 代表该字段为字符串，x 代表该字符串的最大字节数，x 为 大于零的数字，例如 char[5]代表最大长度为 5 字节的字符 串；当最大长度大于实际长度时，右补空格。字符串使用 GBK 编码 |
| int,uint  | 代表该字段为整型数值，如 uint32 表示 32 位无符号整数， int64 表示 64 位有符号整数 |
| Nx、Nx(y) | 与 int、uint 一并使用，用于给出该整型数值实际表示的业 务字段的长度（精度）: Nx 代表最大长度为 x 位数字的整 数；Nx(y)代表最大长度为 x 位数字，其中最末 y 位数字为小数部分 |





golang的net包



# 附录

[上海证券交易所 行情网关技术指引及接口开发指南](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/policy/c/SSE_MDGW_Interface_0.6_20191111.pdf)

[IS120 上海证券交易所行情网关 BINARY 数据接口规范](http://www.sse.com.cn/services/tradingservice/tradingtech/technical/data/c/IS120_BINARY_Interface_CV0.42_20210315.pdf)







