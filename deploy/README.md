## 部署

### HELM

* 镜像
```shell
docker.io/daprio/dashboard:0.6.0
# operator/placement/sentry
docker.io/daprio/dapr:1.0.1
# docker.io/daprio/placement:1.0.1
# docker.io/daprio/sentry:1.0.1
# sidecar
docker.io/daprio/daprd:1.0.1
```

替换为:

```shell
registry.cn-hangzhou.aliyuncs.com/morse_dapr/dashboard:0.6.0
registry.cn-hangzhou.aliyuncs.com/morse_dapr/dapr:1.0.1
registry.cn-hangzhou.aliyuncs.com/morse_dapr/daprd:1.0.1
```