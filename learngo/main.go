package main //오직 컴파일을 위해서만 있다.
// := 는 go 언어에서 타입을 자동으로 지정해준다. 한번 지정해준 타입은 후에 바꿀 수 없다. 한 가지 함수 내에서 지정해준 변수는 그 안에서만 유효하다.
// 상황에 따라 인자 뒤에 type 값을 붙여주자 그리고 리턴 값에도 타입을 지정해주어야 한다. fun multiple(a,b int) int {}
// go.lang.org 에서 package에 관한 함수의 종류를 알 수 있다.
// 함수를 만들고 사용하지 않으면 에러를 발생시킨다.
// 여러개의 타입지정에서 값을 무시하고 진행시키고 싶으면 _ 를 이용하면 된다.
// ...는 모든 인자값(argument)에 대해 타입을 지정해준다.
// naked return의 기능은 함수의 리턴 값을 지정해주지 않아도 작동하게 해주는 것이다. 이때 자동으로 리턴된 값은 처음 리턴값의 타입을 지정할 때의 값이다.
// defer은 함수가 "끝난 후" 실행되는 유용한 기능이다.
// go에서 반복문(loop)는 오직 for만 이용가능하다.
// range는 array("[]")에 loop를 적용시킨다.
// if -else문에서 go만의 특별한 점은 if문에서만의 변수타입 지정이 가능하다는 것이다. 이 장점 때문에 조건문에서 타입에러를 잘 다룰 수 있다.
// Pointer : 변수 앞에 &를 붙여주면 메모리 주소값을 알 수 있다. *은 메모리 주소값에 저장된 값을 알려준다.
// GO에서는 array에 slice라는 []에 길이 값을 넣지 않는 타입이 있다.
// slice 타입에 변수를 넣고 싶으면 append(slice, 값)을 쓰면 된다.
// map 자료구조 용법, 변수명 := map[type값]type값{], map 은 key값과 value값으로 이루어진 자료구조다 파이썬의 dictionary와 비슷하다.
// 변수의 단위로 저장을 위한 구조체(object) 짜는법 type person struct {name string, age int, favfood []string}

import "fmt" //함수를 대문자로 시작해서 쓰면 export 할 수 있다.

func main() {
	fmt.Println("hello world")
}
