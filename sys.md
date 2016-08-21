### 系统架构

### 目录结构
flood 由三个主目录组成:agent,server,vendor
client 表示客户端代理层,负责和server通信,执行具体的压测任务
server 表示整个压测节点的中心节点,负责和所有的agent,呈现和管理压测任务
vendor server和agent共用库