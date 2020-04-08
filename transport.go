package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(svc CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createCustomerRequest)
		v, err := svc.CreateCustomer(req.Name)
		if err != nil {
			return createCustomerResponse{v, err.Error()}, nil
		}
		return createCustomerResponse{v, ""}, nil
	}
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type createCustomerRequest struct {
	Name string `json:"name"`
}

type createCustomerResponse struct {
	ID  string `json:"id"`
	Err string `json:"err,omitempty"`
}
