package main

func polynomialAdd(poli1 []int,poli2 []int)  []int{

	var result []int;

	var bigPoli []int;

	if(len(poli1) < len(poli2)){
		bigPoli = poli2;
	}

	for i := 0;i<len(bigPoli); i++ {
		secondPoliCofecent := poli2[i];
		result[i] = bigPoli[i]  + secondPoliCofecent;
	}
	return result;
}
