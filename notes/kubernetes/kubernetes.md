## Kubernetes
    Infrastructure as a Service / Platform as a service / Software as a Service
### 背景
    容器集群管理  Google 容器化基础框架  borg  Go
    
   - 特点：
     - 轻量级，消耗的资源少；
     - 开源；
     - 弹性伸缩；
     - 负载均衡：IPVS
### 组件说明
#### Borg组件说明
   - api server:所有服务统一入口
   - CrontrollerManger:维持副本期望数目
   - Scheduler:负责介绍任务，选择合适的节点进行分配任务
   - ETCD:键值对数据库，存储K8S集群所有重要信息（持久化）
   - Kubelet：直接跟容器引擎交互实现容器的生命周期管理
   - Kube-Proxy:负责写入规则至IPTABLES、IPVS 实现服务映射访问
   - CoreDNS:可以为集群中的SVC创建一个域名IP的对应关系解析
#### k8s结构说明
##### 网络结构

##### 组件结构 



### 关键字解释
