package main

import "fmt"

type  Node struct{
	val int
	prev *Node
	next *Node
}
type Stack struct{
	head *Node
	tail *Node
	top *Node
}
func Constructor() Stack{
	return Stack{
	}
}

func (st *Stack) Size()int{
	count:=0
	cur:=st.top
	for cur!=nil{
		count++
		cur= cur.prev
	}
	return count
}
func (st *Stack) IsEmpty() bool {
	return st.top==nil
}
func (st *Stack) ShowAllVal()  {
	cur:= st.top
	for cur!=nil{
		fmt.Println(*cur)
		cur=cur.prev
	}
}

func (st *Stack) Push(val int)error{
	node:=&Node{
		val:  val,
	}
	if st.head==nil{
		st.head=node
		st.tail=node
		st.top=node
	} else if st.tail==st.top{
		st.tail.next=node
		node.prev=st.tail
		st.tail=st.tail.next
		st.top=st.tail
	} else if st.tail!=st.top{
		st.top.next.val=node.val
		st.top=st.top.next
	}
	return nil
}
func (st *Stack) Pop()error{
	if st.top==nil{
		return nil
	} else if st.Size()==1{
		st.head=nil
		st.tail=nil
		st.top=nil
	} else{
		st.top=st.top.prev
	}
	fmt.Println("Front element popped out")
	return nil
}


//func main(){
//	st:=Constructor()
//	st.Push(1)
//	fmt.Println("top=>",st.top,"tail=>",st.tail)
//	st.Push(2)
//	fmt.Println("top=>",st.top,"tail=>",st.tail)
//	st.Push(3)
//	fmt.Println("top=>",st.top,"tail=>",st.tail)
//	st.Pop()
//	fmt.Println("top=>",st.top,"tail=>",st.tail)
//	st.Push(4)
//	fmt.Println("top=>",st.top,"tail=>",st.tail)
//	st.ShowAllVal()
//
//}

