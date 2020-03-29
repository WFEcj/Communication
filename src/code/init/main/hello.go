package main
import "fmt"
func main(){
	arr := [5]int{3,4,5,2,1}
	for i := len(arr)-1; i > 0 ; i-- {
		for j := 0 ; j < i ; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
		fmt.Println(arr)
	}
	var answer int
	fmt.Println("请输入您要查找的数：")
	fmt.Scanln(&answer)
	binarySearch(0,len(arr)-1,answer,&arr)
	//创建map的三种方式
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "卢俊义"
	fmt.Println(a)
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	heros := map[string]string {
		"heros1":"宋江",
		"heros2":"吴用",
		"heros3":"武松",
	}
	fmt.Println(heros)
	//删除map操作 当key不存在时，操作不会执行 也不会报错
	delete(cities,"no1")
	//map 的查找
	val , ok := a["no1"];
	if ok {
		fmt.Printf("val=%v",val)
	} else {
		fmt.Println("没有找到!")
	}
	//for range 遍历map
	for i , v := range cities {
		fmt.Printf("i=%v  v=%v",i,v)
	}
	users := make(map[string]map[string]string,10)
	modifyUser(users,"tom")
}
func modifyUser(users map[string]map[string]string,name string) {
	if users[name]!=nil {
		users[name]["pass"] = "888888"
	} else {
		users[name] = make(map[string]string,2)
		users[name]["pass"] = "888888"
		users[name]["nickName"] = "昵称~"+name
	}
}
func binarySearch(start int,end int,  findCValue int,arr *[5]int) {
	if start > end {
		fmt.Println("找不到")
		return
	}
	mid := (start + end) / 2
	if findCValue > (*arr)[mid] {
		binarySearch(mid+1,end,findCValue,arr)
	} else if findCValue < (*arr)[mid] {
		binarySearch(start,mid-1,findCValue,arr)
	} else {
		fmt.Println("找到了")
		fmt.Printf("index=%d",mid)
	}
}