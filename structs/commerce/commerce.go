package commerce

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) Total() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.TotalPrice()
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
