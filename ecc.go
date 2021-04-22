package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"time"
)
var curve01=elliptic.P256()
var params3=curve01.Params()
func test(ab *big.Int,ac *big.Int)(x,y *big.Int){
	fmt.Println(params3.IsOnCurve(ab,ac))
	d,e:=params3.Add(ab,ac,params3.Gx,params3.Gy)
	for i:=1;i<=1000000;i++{
ab,ac:=params3.Add(ab,ac,d,e)
//print
		//fmt.Println("hello",params3.IsOnCurve(ab,ac))
//if i==50{
	//fmt.Printf ("第五十次椭圆曲线点加操作的X和Y的值%v，%v",ab,ac)

if i!=1000000{
	continue
}else{
return ab,ac
}

	}
return
}
func main(){


	a:=params3.B
	b:=params3.P
	c:=params3.N
	d:=params3.Gx//基点x的坐标
	e:=params3.Gy//基点y坐标

	fmt.Printf("%v,%v,%v \n",a,b,c)
	fmt.Printf("%v,%v\n",d,e)
res:=params3.IsOnCurve(d,e)
fmt.Println(res)
addresx,addresy:=params3.Add(d,e,d,e)
res2:=params3.IsOnCurve(addresx,addresy)
fmt.Println(res2)
	fmt.Printf("%v,%v\n",*addresx,*addresy)

	start:=time.Now()
	ad,ac:=test(d,e)
	fmt.Printf("%v,%v",ad,ac)
/*for i:=0;i<=1000;i++{
	//getrandom,_:=rand.Int(rand.Reader,*b)
d,e:=params3.Add(*d,*e,*d,*e)
	fmt.Printf("%d",*d,*e)



	}*/
	end:=time.Since(start)
	fmt.Printf("abv%v",end )
}
