package input

import "fmt"

//.
func Input(arrSize uint) ([]int, error) {
	arr := make([]int, arrSize)
	for i := 0; i < int(arrSize); i++ {
		fmt.Printf("arr[%d]> ", i)
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			return nil, err
		}
	}

	return arr, nil
}
