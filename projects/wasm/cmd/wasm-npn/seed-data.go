package main

import (
	"github.com/kyleu/npn/app"
	"github.com/kyleu/npn/app/request"
)

func seedData(svc *app.Service) error {
	err := svc.Collection.Save("", "seed-a", "Seed A", "Initial seed data, collection A")
	if err != nil {
		return err
	}

	err = svc.Collection.Save("", "seed-b", "Seed B", "Initial seed data, collection B")
	if err != nil {
		return err
	}

	reqA1 := &request.Request{
		Key:         "req-a-1",
		Title:       "Req A1",
		Description: "Initial seed data, request A1",
		Prototype: &request.Prototype{
			Method:   request.MethodGet,
			Protocol: request.ProtocolHTTP,
			Domain:   "localhost",
			Port:     10101,
			Path:     "",
		},
	}
	err = svc.Collection.SaveRequest("seed-a", "", reqA1)
	if err != nil {
		return err
	}

	reqA2 := &request.Request{
		Key:         "req-a-2",
		Title:       "Req A2",
		Description: "Initial seed data, request A2",
		Prototype: &request.Prototype{
			Method:   request.MethodGet,
			Protocol: request.ProtocolHTTPS,
			Domain:   "google.com",
			Path:     "search",
		},
	}
	err = svc.Collection.SaveRequest("seed-a", "", reqA2)
	if err != nil {
		return err
	}

	return nil
}
