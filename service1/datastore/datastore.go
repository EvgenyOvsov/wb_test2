package datastore

import (
	"context"
	"draft/grpc"
	"draft/model"
	"log"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MemCache struct {
	cli grpc.BuyersAndShopsClient
}

func NewDatastore(URL string) *MemCache {
	conn, err := ggrpc.Dial(URL, ggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	cli := grpc.NewBuyersAndShopsClient(conn)
	return &MemCache{cli}
}

// Ну а дальше идут просто круды

func (m *MemCache) CallSaveBuyer(buyer *model.Buyer) error {
	object := &grpc.Buyer{
		LastName:    buyer.LastName,
		FirstName:   buyer.FirstName,
		MiddleName:  buyer.MiddleName,
		Age:         int32(buyer.Age),
		Registation: buyer.Registation,
	}
	_, err := m.cli.CreateBuyer(context.Background(), object)
	return err
}

func (m *MemCache) CallSaveShop(shop *model.Shop) error {
	object := &grpc.Shop{
		Name:    shop.Name,
		Address: shop.Address,
		Working: shop.Working,
		Owner:   shop.Owner,
	}
	_, err := m.cli.CreateShop(context.Background(), object)
	return err
}

func (m *MemCache) CallGetBuyer(lastName string) *model.Buyer {
	object, err := m.cli.GetBuyer(context.Background(), &grpc.Buyer{
		LastName: lastName,
	})
	if err != nil {
		return nil
	}
	return &model.Buyer{
		LastName:    object.LastName,
		FirstName:   object.FirstName,
		MiddleName:  object.MiddleName,
		Age:         int(object.Age),
		Registation: object.Registation,
	}
}

func (m *MemCache) CallGetShop(name string) *model.Shop {
	object, err := m.cli.GetShop(context.Background(), &grpc.Shop{
		Name: name,
	})
	if err != nil {
		return nil
	}
	return &model.Shop{
		Name:    object.Name,
		Address: object.Address,
		Working: object.Working,
		Owner:   object.Owner,
	}
}

func (m *MemCache) CallDeleteBuyer(buyer *model.Buyer) error {
	_, err := m.cli.DeleteBuyer(context.Background(), &grpc.Buyer{
		LastName: buyer.LastName,
	})
	return err
}

func (m *MemCache) CallDeleteShop(shop *model.Shop) error {
	_, err := m.cli.DeleteShop(context.Background(), &grpc.Shop{
		Name: shop.Name,
	})
	return err
}
