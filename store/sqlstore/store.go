package sqlstore

import (
	"github.com/xtopcla/dcim/store"
)

type SqlStore struct {
}

func (s *SqlStore) Datacenter() store.DatacenterStore {
	dcstore := &DatacenterStore{}
	return dcstore
}

type DatacenterStore struct{}

func (s *DatacenterStore) GetAll() string {
	return "blah"
}
