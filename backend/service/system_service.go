package service

import (
	"admin-system/model"
)

type (
	SystemListResp struct {
		Total uint           `json:"total"`
		List  []model.System `json:"list"`
	}

	SystemCreateReq struct {
		Hp  int `json:"hp"`
		Def int `json:"def"`
	}
)

type SystemService struct{}

func (s *SystemService) SystemList(category string) (*SystemListResp, error) {
	users, err := model.SystemList(category)
	if err != nil {
		return nil, err
	}

	return &SystemListResp{
		Total: 1,
		List:  users,
	}, nil
}

func (s *SystemService) SystemCreate(category string, value any) error {
	data := &model.System{
		Name: category,
	}
	return model.SystemCreate(category, data)
}
