package main

import (
	"fmt"
)

type  Node struct{
	val int
	prev *Node
	next *Node
}
type Queue struct{
	head *Node
	tail *Node
	front *Node
}
func Constructor() Queue{
	return Queue{
	}
}

func (q *Queue) Size()int{
	count:=0
	cur:=q.front
	for cur!=nil{
		count++
		cur=cur.next
	}
	return count
}
func (q *Queue) IsEmpty() bool {
	return q.front==nil
}
func (q *Queue) ShowAllVal()  {
	cur:=q.front
	for cur!=nil{
		fmt.Println(*cur)
		cur=cur.next
	}
}

func (q *Queue) Push(val int) error {
	node:=&Node{
		val:  val,
	}
	if q.head==nil{
		q.head=node
		q.tail=node
		q.front=node
	} else if q.head==q.front {
		q.tail.next=node
		node.prev=q.tail
		q.tail=q.tail.next
	} else if q.head!=q.front{
		temp:=q.head
		q.head=q.front
		temp.next=nil
		q.tail.next=temp
		temp.prev=q.tail
		q.tail=q.tail.next
		q.tail.val=node.val
	}
	return nil
}


func (q *Queue) Pop() error {
	if q.head==nil{
		return nil
	} else if q.Size()==1{
		q.head=nil
		q.front=nil
		q.tail=nil
		fmt.Println("Front element popped out")
	} else{
		q.front=q.front.next
		fmt.Println("Front element popped out")
	}
	return nil
}

func main(){
	queue:=Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Pop()
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	queue.Pop()
	queue.Push(8)
	queue.ShowAllVal()

}
