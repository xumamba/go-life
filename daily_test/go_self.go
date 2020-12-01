package main

import (
	"fmt"
	"sync"
)

const c1 = 100

var v1 = 123

type Student struct {
	Name string
}

func main() {
	fmt.Println("======slice=======")
	slice := make([]int, 1)
	slice = append(slice, 1)
	fmt.Println(slice)

	fmt.Println("======address=======")
	fmt.Println(&v1, v1)
	// println(&c1, c1)  // Cannot take the address of 'c1'  常量是无法取出地址的，因为字面量符号(常量)并没有地址而言。

	fmt.Println("======Map的Value赋值=======")
	m := make(map[string]Student)
	student := Student{"name"}
	m["student"] = student // 值拷贝
	// m["student"].Name = "newName"  // m["student"]是一个值引用。那么值引用的特点是只读。

	tempStudent := m["student"]
	tempStudent.Name = "newName"
	m["student"] = tempStudent
	fmt.Println(m["student"]) // 发生2次结构体值拷贝，性能很差。

	m2 := make(map[string]*Student)
	m2["student"] = &student // 值拷贝
	m2["student"].Name = "newName"
	fmt.Println(m2["student"]) // 每次修改的都是指针所指向的Student空间，指针本身是常指针，不能修改，只读属性，但是指向的Student是可以随便修改的，而且这里并不需要值拷贝。只是一个指针的赋值。

	fmt.Println("======map的遍历赋值=======")
	m2["student2"] = &Student{"name2"}
	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("---------")
	stus := []Student{
		{Name: "zhou"},
		{Name: "li"},
		{Name: "wang"},
	}
	for _, stu := range stus {
		m2[stu.Name] = &stu // 实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
	}
	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("---------")
	for i := 0; i < len(stus); i++ {
		m2[stus[i].Name] = &stus[i]
	}
	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("======channel=======")
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			if i == 9 {
				close(ch)
			}
		}
	}()
	for {
		value, ok := <-ch
		fmt.Println(value, ok)
		if !ok {
			break
		}
	}

	fmt.Println("======WaitGroup=======")
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()

}
