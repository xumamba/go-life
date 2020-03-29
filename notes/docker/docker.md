## Docker

    Docker本身是一个容器运行载体或称之为管理引擎。

### Base
  - 镜像 容器 仓库
  - 安装 <替换镜像源>
  - 为什么Docker比VM快？
    
    - docker有着比VM更少的抽象层。由于docker不需要Hypervisor实现硬件资源虚拟化，运行在docker容器上的程序直接使用的都是实际物理机的硬件资源。因此在CPU和内存利用率上docker将会再效率上有明显优势。
    - docker利用的是宿主机的内核，而不需要Guest OS。因此，当新建一个容器时，docker不需要和虚拟机一样重新加载一个操作系统内核。避免引寻、加载操作系统内核整个比较费时费资源的过程，当新建一个虚拟机时，虚拟机软件需要加载Guest OS，整个新建过程是分钟级别的。而docker由于直接利用宿主机的操作系统，因此新建一个docker容器只需要几秒钟。
   
    |   | docker容器|虚拟机VM|
    |---| ------|-------| 
    |操作系统|与宿主机共享OS|宿主机OS上运行虚拟机OS|
    |存储大小|镜像小，便于存储传输|镜像庞大（vmdk、vdi等）|
    |运行性能|几乎无额外的性能损失|操作系统额外的CPU、内存消耗|
    |移植性|轻便、灵活、适用于Linux|笨重，与虚拟化技术耦合度高|
    |硬件亲和性|面向软件开发者|面向硬件运维者|
    |部署速度|快速、秒级|较慢|
  - Commands
    
    - push restart rm rmi run save search start stop tag top unpause version wait 

### Principle
 
    镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于环境开发的软件，它包含运行某个软件所需的所有内容，包括代码、运行时、库、环境变量和配置文件。
 
 - 是什么：
 
    - UnionFS(联合文件系统)
    
        - Union文件系统是一种分层、轻量级并且高性能的文件系统，它支持对文件系统的修改作为一次提交来一层一层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统下（unite several directories into a single virtual filesystem）；
        - Union文件系统是Docker镜像的基础；
        - 镜像可以通过分层来进行继承，基于基础镜像（天启：scratch),可以制作各种具体的应用镜像；
        - 特性：一次同时加载多个文件系统，但从外表看，只能看到一个文件系统，联合加载会把各层文件系统叠加起来，这样最终的文件系统会包含所有底层的文件和目录。
    
    - Docker镜像加载原理
     
      - docker的镜像实际上是由一层一层的文件系统组成UnionFS；
      - bootfs(boot file system)主要包含bootloader和kernel，bootloader主要时引导加载kernel，Linux刚启动时会加载bootfs文件系统，在Docker镜像的最底层是bootfs。这一层与Linux/Unix系统一样，包含boot加载器和内核。当boot加载完成之后整个内核就在内存中了，此时内存的使用权已由bootfs转交给内核，此时系统也会加载bootfs；
      - rootfs(root file system)，在bootfs之上，包含的就是典型的Linux系统中的/dev /proc /bin /etc等标准目录和文件。rootfs就是各种不同的操作系统发行版，比如Ubuntu,centOS等等；
      - 对于一个精简的OS,rootfs可以很小，只需要包括最基本的命令，工具和程序库就可以了，因此底层直接用Host的kernel，自己只需要提供rootfs就行了。由此可见对于不同的Linux发行版，bootfs基本是一致的，rootfs会有差别，因此不同的发行版可以公用bootfs。  
    
    - 分层的镜像
    
    - 为什么Docker镜像要采用分层结构
        
        - 共享资源
 - 特点：
 
    - Docker镜像都是只读的
    - 当容器启动时，一个新的可写层被加载到镜像的顶部，这一层通常被称作“容器层”，“容器层”之下的都叫“镜像层” 
 
 - docker commit -m="提交的描述信息" -a="作者" 容器ID 要创建的目标镜像名:[标签名]

    - -p 主机端口:docker端口
    - -P 随机分配映射端口
    - -i  -t
 
 - Docker容器数据卷
    
    - docker理念
    
        - 将运用与运行的环境打包形成容器运行，运行可以伴随着容器，但是对数据的要求希望是持久化的；
        - 容器之间希望有可能共享数据。
    
    - 引入
    
        - docker容器产生的数据，如果不通过docker commit生成新的镜像，使得数据作为镜像的一部分保存下来，那么当容器删除后，数据自然也就没有了；
        - 为了能保存数据，在docker中使用卷
     
    - 是什么
        
        - 卷就是目录或文件，存在于一个或多个容器中，由docker挂载到容器，但不属于UnionFS，因此能够绕过UnionFS提供一些用于持续存储或共享数据的特性；
        - 卷的设计目的就是数据的持久化，完全独立于容器的生存周期，因此docker不会在容器删除时删除其挂载的数据卷；
        - 类似Redis里面的rdb和aof文件
        
    - 特点
    
        - 数据卷可以在容器之间共享或重用数据
        - 卷中的更改可以直接生效
        - 数据卷中的更改不会包含在镜像的更新中
        - 数据卷的生命周期一直持续到没有容器使用它为止
        
    - 能干嘛
    
        - 容器的持久化
        - 容器间继承+共享数据
     
    - 数据卷
    
        - 容器内添加
            
            - 直接命令添加
            
            docker run -it -v /宿主机绝对路径目录:/容器内目录[:ro] 镜像名
            
            - 查看数据卷是否挂载成功
            
            $ docker inspect
            
            - DockerFile添加
            
            Hello.java   ---> Hello.class
            Docker images ---> DockerFile
            
            dicker run -it --name child02 --volumes-from child01 centOS 
            共享  容器停止退出后，主机修改后数据依旧同步
       
 - DockerFile
    
    - 执行流程：
    
        - docker从基础镜像运行一个容器
        - 执行一条指令并对容器作出修改
        - 执行类似docker commit的操作提交一个新的镜像层
        - docker再基于刚提交的镜像运行一个新容器
        - 执行dockerfile中的下一条指令直到所有指令执行完成
        
    - 体系结构
    
        - FROM MAINTAINER RUN EXPOSE WORKDIR ENV ADD COPY VOLUME CMD ENTRYPOINT ONBUILD
 
---
-日常命令：
    - 进入后台运行的容器中：$docker exec -it containerID /bin/bash  
    - 临时退出容器 Ctrl+P+Q / exit