package pilgan

import (
	"log"
	"serotonin/repositories/messaging"
	logging "serotonin/util/logging"
	"time"

	"github.com/sirupsen/logrus"
)

type PilganService struct {
	repository           Repository
	messaging_repository messaging.MessagingRepository
}

func InitPilganService(repo Repository, repo_messaging messaging.MessagingRepository) *PilganService {
	return &PilganService{
		repository:           repo,
		messaging_repository: repo_messaging,
	}
}

func errSync() {
	data := recover()
	if data != nil {
		// log error here
		logging.GetLogger("errorwhilesyncujian.log").WithFields(logrus.Fields{}).Error(data)
		log.Printf("%v", data)
	}
}

func (service *PilganService) ReadDataPilgan(param map[string]interface{}) (*[]PilganTone, error) {
	list_pilgan, err := service.repository.GetAllDataPilgan(param)
	if err != nil {
		// logging here
		logging.GetLogger("errorreaddata.log").WithFields(logrus.Fields{
			"param": param,
		}).Error(err.Error())
		log.Printf("%s", err)
	}
	return list_pilgan, nil
}

func (service *PilganService) SyncNewDataPilgan(pilgan *PilganTone) error {
	defer errSync()
	id := int((*pilgan)["id_data_core"].(float64)) //read from outside using float64
	res, _ := service.repository.GetDataPilganById(id)
	if res != nil {
		(*pilgan)["version"] = int((*res)["version"].(int32)) + 1 //read from inside using int32
	} else {
		(*pilgan)["version"] = 0
	}
	(*pilgan)["receive_date"] = time.Now()
	err := service.repository.InsertDataPilgan(pilgan)
	if err != nil {
		// log error here while insert to mongoDB
		logging.GetLogger("errorinsertdata.log").WithFields(logrus.Fields{
			"data": (*pilgan),
		}).Error(err.Error())
		log.Printf("%s", err)
		return err
	}
	err = service.messaging_repository.SendMessage("ujian_exchangers", "ack_ujian", "signal.ack.ujian", pilgan)
	if err != nil {
		// log error here while send ack success insert
		logging.GetLogger("errorsendack.log").WithFields(logrus.Fields{
			"data": (*pilgan),
		}).Error(err.Error())
		log.Printf("%s", err)
		return err
	}
	return nil
}
