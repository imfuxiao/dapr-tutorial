package id_server

import v1 "github.com/imfuxiao/dapr-tutorial/pkg/id-server/v1"

type IdService struct {
}

func (i IdService) NewId() int64 {
	panic("implement me")
}

func (i IdService) ExplainId(id int64) v1.ID {
	panic("implement me")
}

func (i IdService) ManualId(time, seq, machineId int64) int64 {
	panic("implement me")
}
