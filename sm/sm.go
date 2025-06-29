package main

import (
	"fmt"

	"pqan.com/learn-go/sm/services/student"
	"pqan.com/learn-go/sm/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("=== Chuong trinh quan ly ===")
		fmt.Println("1. Quan ly sinh vien")
		fmt.Println("2. Thoat")
		choice := utils.GetPositiveInt("Chon chuc nang: ")

		switch choice {
		case 1:
			student.SutdentMenu()
		case 2:
			return
		default:
			fmt.Println("Lua chon khong hop le")
		}

		utils.ReadInput("\n Nhan phim Enter de tiep tuc...")
	}

}
