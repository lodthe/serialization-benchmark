package fprotobuf

import (
	"github.com/golang/protobuf/proto"
	"github.com/lodthe/serialization-benchmark/sample"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	obj, ok := input.(sample.User)
	if !ok {
		return nil, errors.New("invalid input type")
	}

	cart := make(map[string]*CartItem, len(obj.ShoppingCart))
	for id, c := range obj.ShoppingCart {
		item := &c.Item
		cart[id] = &CartItem{
			Item: &Item{
				Id:          item.ID,
				CreatedAt:   timestamppb.New(item.CreatedAt),
				Visible:     item.Visible,
				OwnerId:     item.OwnerID,
				Name:        item.Name,
				Description: item.Description,
				Keywords:    item.Keywords,
				Price:       item.Price,
				Weight:      item.Weight,
			},
			Quantity: c.Quantity,
		}
	}

	converted := &User{
		Name:         obj.Name,
		Phone:        obj.Phone,
		Balance:      obj.Balance,
		RegisteredAt: timestamppb.New(obj.RegisteredAt),
		ShoppingCart: cart,
		Blocked:      obj.Blocked,
	}
	if obj.BirthDay != nil {
		converted.Birthday = timestamppb.New(*obj.BirthDay)
	}

	return proto.Marshal(converted)
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	result, ok := output.(*sample.User)
	if !ok {
		return errors.New("invalid output type")
	}

	var user User
	err := proto.Unmarshal(data, &user)
	if err != nil {
		return err
	}

	result.Name = user.Name
	result.Phone = user.Phone
	result.Balance = user.Balance
	result.RegisteredAt = user.RegisteredAt.AsTime()
	result.Blocked = user.Blocked

	if user.GetBirthday() != nil {
		t := user.Birthday.AsTime()
		result.BirthDay = &t
	}

	result.ShoppingCart = make(map[string]sample.CartItem, len(user.ShoppingCart))
	for id, c := range user.ShoppingCart {
		item := c.Item
		result.ShoppingCart[id] = sample.CartItem{
			Item: sample.Item{
				ID:          item.Id,
				CreatedAt:   item.CreatedAt.AsTime(),
				Visible:     item.Visible,
				OwnerID:     item.OwnerId,
				Name:        item.Name,
				Description: item.Description,
				Keywords:    item.Keywords,
				Price:       item.Price,
				Weight:      item.Weight,
			},
			Quantity: c.Quantity,
		}
	}

	return nil
}
