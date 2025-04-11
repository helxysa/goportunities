package schemas

import (
	"time"

	"gorm.io/gorm"
)
  
type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
}


//No go, alem do model, tem que fazer o response para o json.
//omitempty é para não mostrar o deleted_at se ele for nulo
type OpeningResponse struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at, omitempty"`	
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}