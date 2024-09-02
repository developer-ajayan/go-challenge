package main

import(
	"fmt"
)
func main(){
	jobs:=make(chan int,100)
	result:=make(chan int,100)
	go worker(jobs,result)
	for item:=0;item<100;item++{
		jobs<-item
	}
	close(jobs)
	for item:=range result{
		fmt.Println("result:::",item)	
	}
}

func worker(jobs <-chan int,result chan int){
	for n:=range jobs {
		result<-fibinacct(n)

	}
	close(result)

}
func fibinacct(n int) int{
	if n<=1{
		return n
	}
	return n*fibinacct(n-1)
}