# 记录一下网上找的一个练手项目

## 项目作用

`logAgent` 中实现了日志分发存储的功能，并通过 `etcd` 进行配置的热更新，用一些中间件比如 `kafka` 来维护可靠性。

`logTransfer` *TODO*.

`resources` 中编写了下载依赖项的 `Bash` 脚本。

## 做出的改进

在 `logAgent` 中，原项目各个模块之间的耦合度过高，我重新设计了部分的类，并通过闭包传递信息的方式降低了耦合度。

添加了 `resources`, 使得项目进行移植变得更方便了。最终目的是便于打包进 `Docker`.

## 遇到过的问题以及解决方案

### 已解决

- WSL(Windows Subsystem Linux) 中无法正常启动 `etcd`:

  借一台 Linux 服务器。（认真）

  推测但是没进行实操的其他解决方案:

  1. 关闭 WSL 中默认的 777 权限，具体操作是修改 `wsl.conf` 文件，内容如下，然后重启 WSL.

     ```bash
     >>> sudo vim /etc/wsl.conf
     [automount]
     enabled = true
     options = "metadata,umask=22,fmask=11"
     mountFsTab = false
     ```

     然后关闭防火墙。（我不想关闭防火墙，所以我借了一台 Linux 服务器，然后问题就解决了。）

  2. 在 Windows 下安装 `etcd`, 然后让 WSL 里的程序连接 Windows 的 `etcd`.

- `go get github.com/coreos/etcd/clientv3` 报错：

  在 `go.mod` 文件内添加

  ```tex
  replace (
  	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
  	go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5
  	google.golang.org/grpc => google.golang.org/grpc v1.26.0
  )
  ```

  然后再执行 `go get`.

### 未解决

- 依赖项下载太慢。
- `start.sh` 编写不够合理，我希望的是每个都单独开一个 `shell` 进行观察，同时 `kafka` 要在 `zookeeper` 之后打开，但是没找到合适的方案。

## 收获与记录

- 学会了 `Golang`.
-  `Git` 使用水平达到了掌握。
- 见识到了各种各样的中间件。