#


## Schema Engine

### Resolver 规则

* command 和 query 必须是struct，这样才能根据filed的tag获取参数名称

* command 和 query 必须放在 以 **_commands 和 **_queries 包名的package中，方便根据包名判断类型
