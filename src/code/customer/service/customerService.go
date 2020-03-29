package service

import "fmt"
import "code/customer/model"
type CustomerService struct {
	customers []model.Customer
	customerNum int
}
func NewCustomerService() *CustomerService  {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	customerService.customers = append(customerService.customers,*customer)
	return customerService
}
func (this *CustomerService) List() []model.Customer{
	return this.customers
}
func (this *CustomerService) Add(customer model.Customer) {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
}
func (this *CustomerService) Delete (id int) bool{
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}
func (this *CustomerService) Update (id int) bool {
	index := this.FindById(id)
	if id == -1 {
		return false
	}
	customer := &this.customers[index]
	fmt.Printf("姓名(%v)<回车表示不修改>:",customer.Name)
	fmt.Scanln(&customer.Name)
	fmt.Printf("性别(%v):",customer.Gender)
	fmt.Scanln(&customer.Gender)
	fmt.Printf("年龄(%v):",customer.Age)
	fmt.Scanln(&customer.Age)
	fmt.Printf("电话(%v):",customer.Phone)
	fmt.Scanln(&customer.Phone)
	fmt.Printf("邮箱(%v):",customer.Email)
	fmt.Scanln(&customer.Email)
	return true
}
func (this *CustomerService) FindById(id int) int{
	index := -1
	for i := 0;i < len(this.customers) ; i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}