package sms

import (
	"encoding/json"
	"errors"
	"net/http"
	"sample-api/internal/repository"
)

type Service interface {
	GetData(w http.ResponseWriter, r *http.Request)
	AddData(w http.ResponseWriter, r *http.Request)
}

type service struct {
	Data           []string
	userDetailRepo repository.UserDetail
}

func NewService(userDetailRepo repository.UserDetail) Service {
	return &service{
		Data:           []string{},
		userDetailRepo: userDetailRepo,
	}
}

type GetDataResp struct {
	Name []string `json:"name"`
}

func (s *service) GetData(w http.ResponseWriter, _ *http.Request) {
	var tmp []string
	for _, val := range s.Data {
		tmp = append(tmp, val)
	}

	resp := GetDataResp{Name: tmp}
	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(&resp)
}

type AddDataReq struct {
	Firstname string `json:"firstname"`
}

type AddDataResp struct {
	Result string `json:"result"`
}

func (s *service) AddData(w http.ResponseWriter, r *http.Request) {
	req, err := validateReq(r)
	if err != nil {
		http.Error(w, "invalid request format", http.StatusBadRequest) //400
		return
	}
	s.Data = append(s.Data, req.Firstname)

	data := repository.UserDetailEntity{}
	err = s.userDetailRepo.Create(data)
	if err != nil {
		http.Error(w, "internal server err", http.StatusInternalServerError) //400
		return
	}

	w.Header().Add("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(&AddDataResp{
		Result: "Add data successfully",
	})
}

func validateReq(r *http.Request) (*AddDataReq, error) {
	var req AddDataReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.New("invalid request format")
	}
	if len(req.Firstname) == 0 {
		return nil, errors.New("firstname is required")
	}
	return &req, nil
}
