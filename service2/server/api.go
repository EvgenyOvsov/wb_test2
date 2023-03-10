package server

import (
	"context"
	"draft/grpc"
	"draft/model"
)

func (s *server) CreateBuyer(ctx context.Context, req *grpc.Buyer) (*grpc.Status, error) {
	err := s.ds.SaveBuyer(
		&model.Buyer{
			LastName:  req.LastName,
			FirstName: req.FirstName,
			Age:       int(req.Age),
		},
	)
	if err != nil {
		return &grpc.Status{
			Error: err.Error(),
		}, err
	}
	return &grpc.Status{}, nil
}

func (s *server) CreateShop(ctx context.Context, req *grpc.Shop) (*grpc.Status, error) {
	err := s.ds.SaveShop(
		&model.Shop{
			Name:    req.Name,
			Address: req.Address,
			Working: req.Working,
			Owner:   req.Owner,
		},
	)
	if err != nil {
		return &grpc.Status{
			Error: err.Error(),
		}, err
	}
	return &grpc.Status{}, nil
}

func (s *server) GetBuyer(ctx context.Context, req *grpc.Buyer) (*grpc.Buyer, error) {
	buyer := s.ds.GetBuyer(req.LastName)
	if buyer == nil {
		return nil, nil
	}
	return &grpc.Buyer{
		LastName:  buyer.LastName,
		FirstName: buyer.FirstName,
		Age:       int32(buyer.Age),
	}, nil
}

func (s *server) GetShop(ctx context.Context, req *grpc.Shop) (*grpc.Shop, error) {
	shop := s.ds.GetShop(req.Name)
	if shop == nil {
		return nil, nil
	}
	return &grpc.Shop{
		Name:    shop.Name,
		Address: shop.Address,
		Working: shop.Working,
		Owner:   shop.Owner,
	}, nil
}

func (s *server) DeleteBuyer(ctx context.Context, req *grpc.Buyer) (*grpc.Status, error) {
	err := s.ds.DeleteBuyer(
		&model.Buyer{
			LastName:  req.LastName,
			FirstName: req.FirstName,
			Age:       int(req.Age),
		},
	)
	if err != nil {
		return &grpc.Status{
			Error: err.Error(),
		}, err
	}
	return &grpc.Status{}, nil
}

func (s *server) DeleteShop(ctx context.Context, req *grpc.Shop) (*grpc.Status, error) {
	err := s.ds.DeleteShop(
		&model.Shop{
			Name:    req.Name,
			Address: req.Address,
			Working: req.Working,
			Owner:   req.Owner,
		},
	)
	if err != nil {
		return &grpc.Status{
			Error: err.Error(),
		}, err
	}
	return &grpc.Status{}, nil
}
