> 用于调用翻译接口的作用

百翻译开放平台：[官方文档](http://api.fanyi.baidu.com/doc/21)

因翻译 仅采用get方式 且采用http

--- 

注意：
因为我的业务是普通用户，在百度平台上的并发量为1 所以需要一个锁，每次资源取走之后锁要求自动锁住1s

### 最后

这是提供一个服务形式的平台，url形式参考如下：

```
 /etc?q=hello
 /cte?q=你好
 /translate?query=你好&from=zh&to=en
```