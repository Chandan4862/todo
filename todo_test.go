package main

import( 
	"testing"
	"todo/handler"
)

var task= handler.New()
var slice []handler.Tasks

func BenchmarkDemo(b *testing.B){
for i:=0;i<b.N;i++{
task.Add(handler.Tasks{"Item 1", "Item 1 Post"})
}
}


func BenchmarkGet( b *testing.B){

for i:=0;i<b.N;i++{
	task.Get()

}
}
