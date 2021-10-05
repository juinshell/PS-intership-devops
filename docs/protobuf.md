# XML
**XML 是独立于软件和硬件的信息传输工具**
1. XML 指可扩展标记语言（EXtensible Markup Language）。
2. XML 是一种很像HTML的标记语言。
3. XML 的设计宗旨是传输数据，而不是显示数据。
4. XML 标签没有被预定义。您需要自行定义标签。
5. XML 被设计为具有自我描述性。
6. XML 是 W3C 的推荐标准。\
贴一个简单的XML格式文档：
```xml
<?xml version="1.0" encoding="UTF-8"?>
<note>
  <to>Tove</to>
  <from>Jani</from>
  <heading>Reminder</heading>
  <body>Don't forget me this weekend!</body>
</note>
```
# protobuf
## 优点
### 性能好
1. 时间开销：XML格式化（序列化）的开销还好；但是XML解析（反序列化）的开销一般。
2. 空间开销：熟悉XML语法的同学应该知道，XML格式为了有较好的可读性，引入了一些冗余的文本信息。所以空间开销也不是太好。
### 代码生成机制
For  gRPC services:\
gRPC 使用protoc特殊的 gRPC 插件从 proto 文件生成代码：获得生成的 gRPC 客户端和服务器代码，以及用于填充、序列化和检索消息类型的常规协议缓冲区代码。
For message:\
为每个字段提供了简单的访问器，如name()和set_name()，以及将整个结构序列化/解析为原始字节和从原始字节序列化/解析整个结构的方法。
### 支持“向后兼容”和“向前兼容”
老版本能识别新版本，忽略掉新增属性
新版本能识别老版本，把需要的新增属性设置为非必要或缺省
### 支持多种编程语言
### 二进制编码