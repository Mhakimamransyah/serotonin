package pilgan

type Service interface {
	SyncNewDataPilgan(pilgan *PilganTone) error
	ReadDataPilgan(params map[string]interface{}) (*[]PilganTone, error)
}

type Repository interface {
	InsertDataPilgan(pilgan *PilganTone) error
	GetDataPilganById(id int) (*PilganTone, error)
	GetAllDataPilgan(params map[string]interface{}) (*[]PilganTone, error)
}
