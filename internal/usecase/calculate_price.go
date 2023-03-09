package usecase

import "github.com/mauFade/go-intensive/internal/entity"

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (calculate *CalculateFinalPrice) Execute(data OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(data.ID, data.Price, data.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = calculate.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
