package main
import "fmt"
var arr []int
func Seive(a int,b int){
var i,j,k,l int
for i=1;i<=b;i++{
	if i%2==0{
	arr[i]=1
	} 
	fmt.Println(a+b)	
}



for j=1;j<=b;j++{
	if arr[j] == 0{
	for k=2;k*j<=b;k++{
	arr[k*j]=1
	fmt.Println("bg")
	}
}
}

for l=1;l<=b;l++{
	fmt.Println(arr[l])
}

}

func main(){
	// your code goes here
	var t int
	var a int
	var b int
	fmt.Scan(&t)
	for i:=1;i<=t;i++{
	fmt.Scan(&a,&b)
	Seive(a,b)
	fmt.Println(a+b)
	}
}
