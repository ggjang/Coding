package main

import "fmt"

func main() {
	//변수 add에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

/*문자열
1. (``) Back Quote로 둘러싸인 문자열은 Raw String Literal이라고 부르며
   1) 이 안에 있는 문자열은 별도록 해석 되지 않고 Raw String 그대로의 값을 갖는다.
   2) (``)은 복수 라인의 문자열을 표현할 때 자주 사용된다.
    ex) 문자열 안에 \n 이 있는 경우 이를 New Line으로 해석하지 않는다.

2. ("") 쌍 따옴표로 둘러 싸인 문자열은 Interpreted String Literal이라고 부르며
   1) 복수 라인에 걸쳐 쓸 수 없다.
   2) ("") 안의 Escape 문자열들은 특별한 의미로 해석됨
   3) ("") 를 이용하여 문자열을 여러 라인에 거려 쓰기 위해서는 + 연사자를 이용 하여 결합
*/
func Literal() {
	rawLiteral := `아리랑\n
아리랑\n
  아라리요`

	interLiteral := "아리랑아리랑\n아리리요" //"아리랑아리랑\n" + "아리리요" 처럼 사용 가능

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}

/*데이타 타입 변환 (Type Conversion)
1. T(v) T: 변환하고자 하는 타입, v: 변환될 값*/

func Type_Conversion() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes, str2)
}

/*연산자
1. 산술연산자: 사칙연산자와 증감연산자를 사용
2. 관계연산자: ==, !=, >=
3. 논리연산자: AND(&&), OR(||), NOT(!)
4. Bitwise 연산자: 비트단위 연산 AND(&), OR(|), XOR(), 쉬프트(<<,>>)
5. 할당연산자: 값을 할당하는 =연산자외에 사칙연산 비트연산
6. 포인터연산자: &혹은 *을 사용하여 해당 변수의 주소를 얻어내거나 이를 반대로 Dereference할때 사용
*/

/*조건문
1. if: 해당 조건이 맞으면 {} 블럭안의 내용을 실행
   1)블럭 시작 브레이스({)를 if와 같은 라인에 두어야 한다.
   2)반드시 Boolean 식으로 표현 되어야 한다.
   3)Optional Statement를 함께 실행 가능하며 이때 생성된 변수는 if문 블럭 안에서만 사용 가능

2. switch: 여러 값을 비교해야 하는 경우 혹은 다수의 조건식을 체크해야하는 경우 switch문을 사용

*/
func Condition_Sentence() {
	var k int = 1
	var i int = 1
	var max int = 3

	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Other")
	}

	if val := i * 2; val < max {
		fmt.Println(val)
	}

}

func Switch_Example() {
	var name string
	var category int = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	fmt.Println(name)

	//No Expression
	var score int = 95
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

/*반복문
for: GoLang의 반복문은 for 하나뿐 for 초기값; 조건식; 증감; {...}와 같은 구조를 가지며 해당 식들을 생략 가능하다.
     ()를 생략하며 사용시 에러 발생
*/

func iterator() {
	//for
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	//조건문만 사용
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)

	//for range
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}

/*함수
파라미터 전달 방식
Pass by Value : 함수의 Parameter 변수에 넘겨주는 값을 복사하는 방식으로
				함수 내부에서 값이 변경 되어도 넘겨준 외부의 값은 변화가 없다.

Pass by Reference : 함수의 Parameter 변수에 넘겨주는 것이 값이 아니라 주소이며
					함수 내부에서 값이 변경 되면 외부에서도 동일하게 값이 변경 된다.

Variadic Function(가변인자함수) : 함수에 고정된 개수의 Parameter를 전달하지 않고
								 다양한 개수의 Parameter를 전달 하고자 할때 ...을 사용여 n개의 동일한 type의 Para를 전달한다.

함수의 Return값 : C언어와 다르게 복수의 반환값을 가질 수 있으며
				 Named Return Parameter 기능을 제공하여 반환 값을 return Paramter에 저장할 수 있다.
*/
func PassByValue(msg string) {
	println(msg)
}

func PassByRef(msg *string) {
	println(*msg)
	*msg = "Changed!"
}

func VariadicFX(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

//함수의 Return값 Example
func sum(num ...int) (total int, count int) {
	for _, i := range num {
		total += i
	}
	count = len(num)

	return
}

/*익명 함수
  함수명을 가지지 않는 함수
  일반적인 사용법은 함수 전체를 변수에 할당하거나 다른 함수의 Paramter에 직접 할당하는 방식을 사용

  sum := func(nums ...int) int {
		total := 0
		for _, i := range nums {
			total += i
		}

		return total
	}

	result := sum(1, 2, 3, 4, 5, 6)
	println(result)
*/

/*일급 함수 (First Class Function)
  Go에서 함수는 일급 함수로 Go의 기본 타입과 동일하게 취급된다.
  따라서 다른 함수의 PAramter로 전달하거나 리턴 값이 될 수 있다.
  함수의 파라미터로 사용하는 방법은 익명함수로 변수에 할당한이후 넘겨주거나 파라미터 변수 자체에 적는 방법이 있다.
  func main() {
	//변수 add에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
*/
