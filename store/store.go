package store

type Store interface {
	Datacenter() DatacenterStore
}

type DatacenterStore interface {
	GetAll() string
}
