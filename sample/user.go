package sample

import (
	"math/rand"
	"time"
)

type User struct {
	Name         string
	BirthDay     *time.Time
	RegisteredAt time.Time
	Phone        string
	Balance      float64
	ShoppingCart ShoppingCart
	Blocked      bool
}

type ShoppingCart map[string]CartItem

type CartItem struct {
	Item     Item
	Quantity int32
}

type Item struct {
	ID        string
	CreatedAt time.Time
	Visible   bool
	OwnerID   uint32

	Name        string
	Description *string
	Keywords    []string

	Price  float64
	Weight float32
}

var Sample User

func init() {
	rand.Seed(1)

	Sample = User{
		Name:         "Joanne Rowling",
		RegisteredAt: time.Now().Truncate(time.Millisecond),
		Phone:        "+1 123 14992 91249",
		Balance:      32929.12223412,
		ShoppingCart: make(map[string]CartItem),
		Blocked:      false,
	}

	var keywords []string
	for i := 0; i < 100; i++ {
		keywords = append(keywords, "k"+genName(10))
	}

	for i := 0; i < 100; i++ {
		item := Item{
			ID:        genName(35),
			CreatedAt: time.Now().Truncate(time.Millisecond),
			Visible:   rand.Intn(10) > 0,
			OwnerID:   rand.Uint32(),
			Name:      genName(50),
			Price:     rand.Float64(),
			Weight:    rand.Float32(),
		}

		if rand.Intn(10) > 0 {
			tmp := genString(700)
			item.Description = &tmp
		}

		for j := 0; j < 5; j++ {
			item.Keywords = append(item.Keywords, keywords[rand.Intn(len(keywords))])
		}

		Sample.ShoppingCart[item.ID] = CartItem{
			Item:     item,
			Quantity: int32(rand.Intn(10) + 1),
		}
	}
}

func genString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = 32 + byte(rand.Intn(125-32))
	}

	return string(b)
}

func genName(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = 'A' + byte(rand.Intn(26))
	}

	return string(b)
}
