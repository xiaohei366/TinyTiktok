# TinyTiktok项目
## 一、如何测试框架的项目
### 1.首先需要对环境初始化：
- 根据`./cmd/user/initialize/db`内的sql文件进行MYSQL初始化（PS：数据库名建议"TinyTiktok"（shared处可修改））
- 下载etcd，并在一个bash挂机下运行（直接运行即可）。
### 2.启动指令
设定项目路径为当前根目录
#### Run User RPC Server
```bash
cd ./cmd/user
sh build.sh
sh output/bootstrap.sh
```
#### Run API Server
```bash
cd ./cmd/api
go run .
```
### 3.测试
我未经postman测试，根据日志直接在抖声app上测试的。进入抖声直接打开“我的”，随后测试注册、登录以及返回的用户粉丝和关注数(测试值为10)
## 二、service如何插入框架（以我的user模块制作过程为例）
<font color = red>我们默认的根目录为`github.com/xiaohei366/TinyTiktok/`，在开发过程中，后面直接加相对路径</font>
1. 首先在`idl`写好user文件夹下的idl/UserServer.proto文件，对应暴露的接口和服务。随后在cmd文件夹下，创建自己模块的文件夹，并通过定义好的idl文件来自动的生成服务端的微服务框架。