Pangu Network 监控

[TOC]

# tm-monitor 

tm-monitor 是 Tendermint 官方的监控工具，主要用于监控网络中的一些自己关心的 node。

## 文档及使用说明

Git repository：[tendermint/tools](https://github.com/tendermint/tools)

Docs：[tm-monitor](http://tendermint.readthedocs.io/en/master/tools/benchmarking-and-monitoring.html#tm-monitor)



## 用法实例

Go 交叉编译 Linux 二进制可执行文件或者 Linux 本地编译（tm-monitor），然后直接使用命令 `tm-monitor host1:46657,host2:46657` 即可。

Pangu Testnet Monitor Url：[http://47.104.155.125:46670/](http://47.104.155.125:46670/)

Note：

>  此监控结果的报警程度只能设为 info 或 warning，监控结果里的 `health` 字段并不能完全代表整个网络的健康与否，因为当有一个监控 node 下线时，`health` 就变为了 `1`（不正常）。



# 监控 rainbow api 响应状态

针对监控 api 的响应状态，目前只是做了最基本的监控（api status_code 是否正常）与报警。即挑选特定的 api 为代表，来设置监控与报警。

## 用法实例

[zabbix 监控](http://106.15.33.15:8199/zabbix/zabbix.php?action=dashboard.view)

邮件截图



![](https://raw.githubusercontent.com/kaifei-bianjie/images/master/mail_20180403131302.png)

