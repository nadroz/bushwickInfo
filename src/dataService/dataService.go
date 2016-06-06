package dataService

import (
"config"
"fmt"
"dbRepository"
"strconv"
)

type DataService struct {
	Repo *dbRepository.DbRepository
}

func Initialize(config *config.DbConfig) (*DataService, error){
	bRepo, err := dbRepository.Initialize(config)
	if err != nil{
		fmt.Println("no workee")
		return nil, err
	}
	srv := DataService { Repo: bRepo }
	return  &srv, nil
}

func (srv DataService) GetActors () (*dbRepository.DataFrame, error) {
	acts, actsErr := srv.Repo.GetActors()
	if actsErr != nil{
		return nil, actsErr
	}
	return acts, nil
}

func (srv DataService) HealthCheck () (int, error) {
	rowCount, qErr := srv.Repo.HealthCheck()
	if qErr != nil{
		return 0, qErr
	}
	
	fmt.Println(strconv.Itoa(rowCount))
	return rowCount, nil
}

func (srv DataService) CheckConfig(config *config.DbConfig) (string, error) {
	return config.BushwickConn, nil
}
/*
func (srv DataService)  AddActor() {
	srv.repo.AddActor()
}*/

