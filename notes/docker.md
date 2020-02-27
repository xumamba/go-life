##Docker

### 镜像原理
 
    镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于环境开发的软件，它包含运行某个软件所需的所有内容，包括代码、运行时、库、环境变量和配置文件。
 
 - 是什么：
 
    - UnionFS(联合文件系统)
    
    - Docker镜像加载原理
    
      - bootfs(boot file system)
            
      - rootfs(root file system)  
    
    - 分层的镜像
    
    - 为什么Docker镜像要采用分层结构
 - 特点：
 
    -  Docker镜像都是只读的
 
 - docker commit -m="提交的描述信息" -a="作者" 容器ID 要创建的目标镜像名:[标签名]

    - -p 主机端口:docker端口
    - -P 随机分配映射端口
    - -i  -t
 
 - Docker容器数据卷
 
    - 是什么
    
        - 类似Redis里面的rdb和aof文件
        
    - 能干嘛
    
        - 容器的持久化
        - 容器间继承+共享数据
     
    - 数据卷
    
        - 容器内添加
            
            - 直接命令添加
            - DockerFile添加