package repository

//repository connect data
//Class
type Customer struct {
	CustomerID int `db:"customer_id"`
	Name string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City string `db:"city"`
	ZipCode string `db:"zipcode"`
	Status int `db:"status"`
}

//สร้าง port เป็น อินเตอร์เฟส
type CustomerRepository interface {
	//funtion
	//retrun array of customer
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}