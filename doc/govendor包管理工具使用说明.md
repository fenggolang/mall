#### 1. 下载
```bash
go get -v github.com/kardianos/govendor
```
#### 2. 初始化项目为vendor管理的目录
```bash
# 进入项目根目录执行
# 下面命令作用是：在当前项目中创建vendor目录和vendor/vendor.json文件(vendor.json存储包的版本信息)
govendor init

# 引用外部包
# 下面命令作用是：GOPATH中去找本项目依赖的包到vendor目录中
# 把gopath下的包(只会是被项目引用到的包)添加到vendor目录下 必须vendor目录下没有，
# 且vendor.json中没有记录这个包的时候才会添加，其中一个存在则命令无效亦不报错
govendor add +external

```
#### 3. govendor其他命令
```bash
# 查看govendor版本
govendor -version

# 移除vendor目录中没有使用的包
govendor remove +unused # 或者govendor remove +u

# 查看当前项目依赖包：包括列出项目本身的包和vendor中的外部包
govendor list

# Look at what is using a package
govendor list -v fmt

# Test your repository only
govendor test +local

# 拉取包(如果是vendor.json里面手写的包，但是又没在GOPATH里面下载下来，那么执行下面命令可以直接下载依赖包到vendor目录
# 此时gopath里面不会有下载的包，只会存在于vendor目录)
govendor sync

# 帮助文档
govendor -h

```
#### 参考
[govendor使用手册](https://my.oschina.net/u/3628490/blog/2245119)
