# 环境配置
## etcd：docker里面起的。端口写在shared里面了。
## mysql：docker里面起的。端口写在config里面了。
## minio：docker里面起的。端口写在shared里面了。后续可能改到video里面去。

# 启动
```bash
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```


