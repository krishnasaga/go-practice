package main

func polynomialAdd(poli1 []int,poli2 []int)  []int{

	var result []int;

	for i := 0;i<len(poli1); i++ {
		result[i] = poli1[i] + poli2[i];
	}
	return result;
}
