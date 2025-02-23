package service

import (
	"fmt"
	db "todo/pkg/database"
	p "todo/pkg/processor"
	mux "todo/pkg/rest"
)

func NewService() *service {

	dbConn := db.NewDatabaseConnection()
	handler := p.NewProcessor(dbConn)
	rest, err := mux.NewServer("http://localhost:55000/", handler) // URL fomr config
	if err != nil {
		panic("error in main service , entrypoint will restart")
	}

	return &service{
		databaseConnection: dbConn,
		processor:          handler,
		restServer:         rest,
	}
}

type service struct {
	databaseConnection database
	processor          processor
	restServer         restServer
}

func (s *service) Run() {
	err := s.restServer.Run()
	if err != nil {
		fmt.Printf("Run: restServer: %s", err.Error())
		return
	}
}

func (s *service) Stop() error {
	err := s.databaseConnection.CloseConnection()
	if err != nil {
		fmt.Printf("Stop: Database close : %s", err.Error())
		return err
	}

	err = s.restServer.Stop()
	if err != nil {
		fmt.Printf("Stop: Database close : %s", err.Error())
		return err
	}
	return nil
}
