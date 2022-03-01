package repository

type AddressDetailEntity struct {
}

type AddressDetail interface {
	FindAll() ([]AddressDetailEntity, error)
}

type addressDetail struct {
}

func NewAddressDetail() AddressDetail {
	return &addressDetail{}
}

func (repo addressDetail) FindAll() ([]AddressDetailEntity, error) {
	return nil, nil
}
