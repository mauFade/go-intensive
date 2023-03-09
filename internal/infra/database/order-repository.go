package database

import (
	"database/sql"

	"github.com/mauFade/go-intensive/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (repository *OrderRepository) Save(order *entity.Order) error {
	_, err := repository.DB.Exec("INSERT INTO orders (id, price, tax, final_price), VALUES(?,?,?,?)", order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (repository *OrderRepository) GetTotalOrdersQuantity() (int, error) {
	var total int

	err := repository.DB.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
