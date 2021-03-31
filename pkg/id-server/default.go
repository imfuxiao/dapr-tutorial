package id_server

import "github.com/spf13/pflag"

// 最大峰值类型: 版本(1 bit) + ID类型(1 bit) + ID发布方式(2 bit) + 时间(30 bit, 精度到秒) + 序列号(20 bit) + 机器ID(10 bit)
// 最大粒度类型: 版本(1 bit) + ID类型(1 bit) + ID发布方式(2 bit) + 时间(40 bit, 精度到毫秒) + 序列号(10 bit) + 机器ID(10 bit)
var (
	version       Version
	idType        IdType
	generatorType GeneratorType
	storeName     string
)

type (
	Version       uint8
	IdType        uint8
	GeneratorType uint8
)

// ID版本: 1 bit
const (
	DefaultVersion   Version = iota // 0 默认版本
	ExtensionVersion                // 1 扩展版本
)

// ID类型: 1 bit
const (
	MaximumPeakIdType     IdType = iota // 0 最大峰值类型
	MinimumParticleIdType               // 1 最小粒度类型
)

// ID生成方式: 2 bit
const (
	MicroServerGeneratorType GeneratorType = iota // 微服务生成
	SDKGeneratorType                              // SDK生成
	ManualGeneratorType                           // 手工生成
	OtherGeneratorType                            // 其他方式
)

func init() {
	pflag.Uint8Var((*uint8)(&version), "id-version", uint8(DefaultVersion), "id version")
	pflag.Uint8Var((*uint8)(&idType), "id-type", uint8(MaximumPeakIdType), "id type")
	pflag.Uint8Var((*uint8)(&generatorType), "id-generator-type", uint8(MicroServerGeneratorType), "id generator type")
	pflag.StringVar(&storeName, "store-name", DefaultStoreName, "store name")
	pflag.Parse()
}
