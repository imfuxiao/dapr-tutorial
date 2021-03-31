package id_server

// 环境变量
const (
	EnvPodIp = "POD_IP" // POD IP
)

// dapr store key
const (
	MachineIdSeq    = "machine_id_seq" // 机器ID序列: 从0开始
	MachineIdKeyFmt = "machine_id_%s"  // 机器ID存储key格式, machine_id_POD_IP
	MaxMachineSeq   = 1024             // 机器ID共10个bit, 所以最大1024

	DefaultStoreName = "statestore" // 默认存储模块名称
)
