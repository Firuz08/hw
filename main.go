package main

import(
	`fmt`
)

type pur struct{
	purchase 		[]int64
	percentCashback int64
	cash 			int64
}

func (new pur) cashback()int64 {
	var sum int64
	for _, i := range new.purchase{
		sum+=i
	}
	var cash int64
	cash = sum * new.percentCashback/100
	return cash
}

func main(){
	purchases := pur{
		purchase: 			[]int64{123,489,652,789,235},
		percentCashback:	20,
		cash:				0,
	}
	
	purchases.cash = purchases.cashback()
	fmt.Println(purchases.cash)
}