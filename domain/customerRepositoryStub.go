package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func(s CustomerRepositoryStub) FindAll() ([]Customer , error) {
	return s.customers, nil
}

func NewCustomerRepositotyStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001", "Danilo", "Brasilia", "28026140", "1989-07-06", "1"},
		{"1002", "Juliana", "Rio de Janeiro", "28026140", "1991-12-06", "1"},
	}

	return CustomerRepositoryStub{customers}
}