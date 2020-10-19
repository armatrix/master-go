package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// MyFloat 别名对数据进行单独的处理
type MyFloat float64

// MarshalJSON 保留固定的精度
func (mf MyFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", float64(mf))), nil
}

// PricePoint for demo
type PricePoint struct {
	Price MyFloat   `json:"price"`
	From  time.Time `json:"valid_from"`
	To    time.Time `json:"valid_to"`
}

func main() {
	p := PricePoint{
		Price: 12.2889777,
	}
	b, _ := json.Marshal(&p)
	fmt.Println(string(b))
}
