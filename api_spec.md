# API 设计规范

[TOC]

## URI 的设计

1. 在RESTful架构中，每个网址代表一种资源（resource），**一般是名词且表示方式为复数**。

2. URI 使用**小写**，如果资源由多个单词组成，使用 `-` 连接。

3. 对于资源的具体操作类型，由HTTP动词表示。

   常用的HTTP动词有下面五个（括号里是对应的SQL命令）。
   > - GET（SELECT）：从服务器取出资源（一项或多项）。
   >
   > - POST（CREATE）：在服务器新建一个资源。
   >
   > - PUT（UPDATE）：在服务器更新资源。
   >
   > - DELETE（DELETE）：从服务器删除资源。

   下面是一些例子。
   > - GET /zoos：列出所有动物园
   > - POST /zoos：新建一个动物园
   > - GET /zoos/ID：获取某个指定动物园的信息
   > - PUT /zoos/ID：更新某个指定动物园的信息
   > - DELETE /zoos/ID：删除某个动物园
   > - GET /zoos/ID/animals：列出某个指定动物园的所有动物
   > - DELETE /zoos/ID/animals/ID：删除某个指定动物园的指定动物
   >

##过滤、排序和搜索等信息

**过滤**

对每一个字段使用一个唯一查询参数，就可以实现过滤。

例如：

```
GET /tickets?state=open // 获取状态为开放的票
```

**排序**

采用泛型参数（`sort`）来描述排序的规则，排序参数采取逗号分隔的字段列表的形式，每一个字段前都可能有一个负号来表示按降序排序。

例如：

```
GET /tickets?sort=-priority // 获取票据列表，按优先级字段降序排序
GET /tickets?sort=-priority,created_at // 获取票据列表，按“priority”字段降序排序。在一个特定的优先级内，较早的票排在前面。
```

**搜索**

当我们需要全文搜索或者需要搜索多个字段时，我们可以采取泛型参数（`q`）的形式来表示。

```
GET /tickets?q=return // 查询印有 "return" 字样的票
```

**分页参数**

使用 `page` 和 `per_page` 来表示当前处于第几页及每页的记录数。

```
GET /tickets?page=1&per_page=10
```



## 响应和错误处理

**响应**

统一采用 `json`   形式。

json 中的字段名采用 **snake_case** 形式，如 `err_code` 而非 `errCode`。

### REST API

**结构和错误处理**

由于在复杂的系统中 http 状态码并不能覆盖所有的错误处理。为此我们采用所有响应的状态码都为 `200`，只是在返回结果中表明状态和错误信息，结构如下：

```json
{
    "status": "success", // success or fail,  用来表明请求状态：成功 或 出错
    "err_code": "",
    "err_msg": "",
    "data": {
        ...
    }
}

// 如果 status 为 success，则可读取 data 里面的内容作后续处理。
// 如果 status 为 fail ，则读取 err_code 和 err_msg。
```

这样设计的好处在于既不用关心 http 状态码，也不用对成功或失败来处理不同的JSON结构。

### WebSocket

```json
{
    "type": "", // 标识消息类型
    "data": {
        ...
    }
}
```



## 版本控制

API-SERVER 里不加版本控制，当出现了不兼容的升级后，所有依赖与此 API 的调用者都必须强制升级。



## 认证

由于 Restful API 是无状态的，当我们需要进行权限认证时，可在请求头里加上请求凭证（access_token），可采用 [OAuth 2.0 ](http://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html)框架。



## 缓存

HTTP提供了自带的缓存框架。你需要做的是在返回的时候加入一些返回头信息，在接受输入的时候加入输入验证。基本两种方法：

- ETag：当生成请求的时候，在HTTP头里面加入ETag，其中包含请求的校验和和哈希值，这个值和在输入变化的时候也应该变化。如果输入的HTTP请求包含IF-NONE-MATCH头以及一个ETag值，那么API应该返回304 not modified状态码，而不是常规的输出结果。
- Last-Modified：和etag一样，只是多了一个时间戳。返回头里的Last-Modified：包含了 [RFC 1123](https://link.zhihu.com/?target=http%3A//www.ietf.org/rfc/rfc1123.txt) 时间戳，它和IF-MODIFIED-SINCE一致。HTTP规范里面有三种date格式，服务器应该都能处理。

## 未完待续...

由于 api 规范的定制不是一个一蹴而就的事情，可能会随着遇到的问题有所增加或调整，后面也会持续的完善和更新。

## 其他规范可参考

[Restful API 最佳实践](https://zhuanlan.zhihu.com/p/25647039)

[Restful API 设计最佳实践](https://www.cnblogs.com/yuzhongwusan/p/3152526.html)



