package util

import (
	dapr "github.com/dapr/go-sdk/client"
	"strconv"
)

type StateItem dapr.StateItem

func NewStateItem(si *dapr.StateItem) *StateItem {
	return (*StateItem)(si)
}

func (s *StateItem) GetIntValue() (int, error) {
	return strconv.Atoi(string(s.Value))
}
