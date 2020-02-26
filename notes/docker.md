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