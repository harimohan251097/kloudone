import "github.com/jinzhu/gorm"

type Cycle struct {
	gorm.Model
	Order_Id      int64 `gorm:"autoIncrement"`
	Frame         string
	Frame_price   int
	Brakes        string
	Brakes_price  int
	Wheel         string
	Wheel_price   int
	chain         string
	chain_price   int
	Seating       string
	Seating_price int
	total_price   int
}