package main

import "fmt"
import "code/customer/service"
import "code/customer/model"
type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}
func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers) ; i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}
func (this *customerView) add() {
	fmt.Println("----------添加客户----------")
	fmt.Printf("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别：")
	sex := ""
	fmt.Scanln(&sex)
	fmt.Printf("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name,sex,age,phone,email)
	this.customerService.Add(*customer)
}
func (this *customerView) delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("a请您选择删除的编号（-1退出）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return 
	}
	fmt.Println("确认是否删除？（Y/N）")
	answer := ""
	fmt.Scanln(&answer)
	if answer == "y" || answer == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("----------删除成功----------")
		} else {
			fmt.Println("----------id不存在----------")
		}
	}
}
func (this *customerView) exit() {
	fmt.Println("您是否要退出？（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}
func (this *customerView) update() {
	fmt.Println("----------修改客户----------")
	fmt.Printf("请输入要修改的id号：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if this.customerService.Update (id) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败----------")
	}
}
func (this *customerView) mainMenu() {
	for{
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("           1 添加客户")
		fmt.Println("           2 修改客户")
		fmt.Println("           3 删除客户")
		fmt.Println("           4 客户列表")
		fmt.Println("           5 退    出")
		fmt.Printf("请输入选择<1-5>:")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入有效的选项！！")
		}
		
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户信息管理软件！")
}
func main()  {
	customerView := customerView{
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}