package v1

type ID struct {
	Version       uint8  `json:"version" protobuf:"int64,1,name=version"`
	Type          uint8  `json:"type" protobuf:"int64,2,name=type"`
	GeneratorType uint8  `json:"generatorType" protobuf:"int64,3,name=generatorType"`
	Time          string `json:"time" protobuf:"string,4,name=time"`
	Timestamp     int64  `json:"timestamp" protobuf:"int64,5,name=timestamp"`
	Seq           int64  `json:"seq" protobuf:"int64,6,name=seq"`
	MachineId     int64  `json:"machineId" protobuf:"int64,7,name=machineId"`
}
