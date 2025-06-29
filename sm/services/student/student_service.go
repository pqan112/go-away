package student

import (
	"fmt"

	"pqan.com/learn-go/sm/utils"
)

var studentList []Student

func addStudent() {
	var scores []float64
	var id int
	fmt.Println("-=-=-=- Them sinh vien -=-=-=-")
	for {
		id = utils.GetPositiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, studentList) {
			break
		}
		fmt.Println("id da ton tai")
	}
	name := utils.GetNonEmptyString("- Nhap ten: ")
	class := utils.GetNonEmptyString("- Nhap lop: ")
	totalPoint := utils.GetPositiveInt("- Nhap so luong diem: ")

	for i := 1; i <= totalPoint; i++ {
		score := utils.GetPositiveFloat(fmt.Sprintf("- Nhap diem %d: ", i))
		scores = append(scores, score)
	}

	student := Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	}

	studentList = append(studentList, student)

	fmt.Println(student.GetInfo())
}

func deleteStudent() {

}

func updateStudent() {
	fmt.Println("-=-=-=- Sua sinh vien -=-=-=-")
	id := utils.GetPositiveInt("Nhap Id sinh vien can sua: ")

	for _, s := range studentList {
		if s.GetId() == id {
			fmt.Println("Nhap thong tin moi")
			name := utils.GetOptionalEmptyString(fmt.Sprintf("- Nhap ten (%s): ", s.Name), s.Name)
			class := utils.GetOptionalEmptyString(fmt.Sprintf("- Nhap lop (%s): ", s.Class), s.Class)

			student := Student{
				Id:    id,
				Name:  name,
				Class: class,
			}

			fmt.Println(student.GetInfo())

		}
	}
}

func listStudent() {
	fmt.Println("-=-=-=- Danh sach sinh vien -=-=-=-")
	if len(studentList) == 0 {
		fmt.Println("Khong co sinh vien nao trong danh sach")
		return
	}
	for _, s := range studentList {
		fmt.Println(s.GetInfo())
	}
}

func searchStudent() {

}

func SutdentMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("Quan ly sinh vien")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("Chon chuc nang: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			listStudent()
		case 5:
			searchStudent()
		case 6:
			return
		default:
			fmt.Println("Lua chon khong hop le")
		}

		utils.ReadInput("\n Nhan phim Enter de tiep tuc...")

	}
}
