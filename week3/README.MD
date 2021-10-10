
https://zhuanlan.zhihu.com/p/102450025
```dockerfile
FROM base_image:tag    # 引用基础镜像 *必要*

ARG arg_key[=default_value1]     # 声明变量
ENV env_key=value2     # 声明环境变量

# 构建几乎不变的部分，例如整体的目录结构，build时依赖的文件和工具包等
COPY src dst
RUN command1 && command2 ...

WORKDIR /path/to/work/dir   # 设置工作目录 

# 构建较少变动的部分，例如应用的依赖的文件、依赖的包等
COPY src dst
RUN command3 && command4 ...

# 构建经常变动的部分，例如应用的编译生成
COPY src dst
RUN command5 && command6 ...

# 容器入口  *必要*
ENTRYPOINT ["/entry.app"]  # 指定容器启动时默认执行的命令
CMD ["--options"] # 指定容器启动时默认命令的默认参数
```