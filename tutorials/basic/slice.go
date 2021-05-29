package main

import "fmt"

func main(){
	//cấu trúc slice: reflect.SliceHeader

	oldSlice := []byte{'h', 'e', 'l', 'l', 'o', ' '}

	newSlice := oldSlice[:2]
	fmt.Println(newSlice)
	newSlice[0] = 'H'
	fmt.Println(oldSlice)	

	//memory allocation slice: khi slice ban đầu không đủ chứa khi thêm phần tử thì hàm append
	//sẽ thực hiện cấp phát lại vùng nhớ có độ rộng gấp đôi và thực hiện sao chép
	//hàm append sẽ không cấp lại vùng nhớ khi chưa đạt tới sức chứa tối đa của cap
	fmt.Println(TrimSpace(oldSlice))
}

func TrimSpace(s []byte) []byte{
	b := s[:0]

	for _, v := range s {
		if(v != ' ') {
			b = append(b, v)
		}
	}

	return b
}