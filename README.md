## Dapr Tutorial

### ping/pong

### id-service 发号服务

发号服务的需求:

* 全局唯一: 利用时间的有序性, 不使用锁(使用锁会降低性能), 并且在时间的某个单元下采用自增序列, 来达到全局唯一.
* 粗略有序: UUID最大的问题是无序. 而要做到完全有序, 就必须使用锁或分布式锁. 所以采用了折中的粗略有序. 秒有序或毫秒有序.
* 可反解: 可通过ID知道它是什么时间生成的等等
* 可制造性: 出了问题, 可手中制造.
* 高性能: 
* 高可用: 
* 可伸缩性: 必须支持水平可伸缩, 满足业务的增长

发布模式:

* REST发布模式: 通过http/gRPC接口对外发布

ID类型:

* 最大峰值类型: 性能好, 但是只能做到秒有序.
* 最小粒度类型: 可以做到毫秒有序, 但是每个毫秒只能生成1024个ID.

ID格式: 总长度63个bit

* 最大峰值类型: 版本(1 bit) + ID类型(1 bit) + ID发布方式(2 bit) + 时间(30 bit, 精度到秒) + 序列号(20 bit) + 机器ID(10 bit)
* 最大粒度类型: 版本(1 bit) + ID类型(1 bit) + ID发布方式(2 bit) + 时间(40 bit, 精度到毫秒) + 序列号(10 bit) + 机器ID(10 bit)

字段说明

* 版本: 1 bit, 用来扩展或者扩容的临时方案

默认值为0, 1表示扩展或扩容. 

* ID类型: 1 bit, 0表示最大峰值类型, 1表示最小粒度类型
* ID发布方式: 2 bit

00: 微服务发布模式
01: 内置sdk发布
10: 手动生成
11: 保留

* 时间

最大峰值类型: 30bit, 精确到秒. 2^30/(365*24*60*60) = 34 即可以使用 30 多年
最小粒度类型: 40bit, 精确到毫秒, 2^40/(365*24*60*60*1000) = 34, 同样可以使用 30 多年

注意: 操作系统层面时间必须保持一致.

* 序列号

最大峰值类型: 20bit, 2^20, 即每秒大约可生成1百万个ID. 即qps大约为百万级
最小粒度类型: 10bit, 2^10, 即每毫秒大约可生成1024个ID.

* 机器ID: 10 bit, 最多可支持1000个服务器组成集群


## go-to-protobuf

`go-to-protobuf`



```sh
> brew install protobuf
> make all WHAT=vendor/k8s.io/code-generator/cmd/go-to-protobuf
> make all WHAT=vendor/k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo
```

go-to-protobuf下包含2个工具：

go-to-protobuf: 用于把package下type生成protobuf message
protoc-gen-gogo: gogo-protobuf 用于生成type的序列化/反序列化代码

生成的顺序大致如下：
package下源码 --(go-to-protobuf)--> generated.proto --(protoc-gen-gogo)--> generated.pb.go


## dapr 存储事务配置

```yaml
- name: actorStateStore
  value: "true"
```

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: statestore
spec:
  type: state.redis
  metadata:
  - name: redisHost
    value: localhost:6379
  - name: redisPassword
    value: ""
  - name: actorStateStore
    value: "true"
```