package main

func maxSlidingWindow(nums []int, k int) []int {
  if(k == 0) {return []int{}}

	// veirificar se existe elemento na fila
	// caso exista, verifica se ele e maior que o item atual
	// se for, retira o item da fila e coloca o item atual
	// caso contrario, apenas insere

	// retira o primeiro item na proxima janela
	// 
	

	return []int{}
}

func maxInSlice(nums []int) int {
	if(len(nums) == 1) {
		return nums[0]
	}

	initialValue := nums[0]

	for _, value := range nums {
		if value > initialValue {
			initialValue = value
		}
	}

	return initialValue
}
