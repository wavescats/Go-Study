package main

import "fmt"

func main() {

	var test1 string = "test string" 
	// 변수 이름을 정하고 / string 문자열 / 변수에 들어갈 내용
	
	//--------------------------------------------
	var test2 string // 변수이름과 타입만 지정하고
	test2 = "test string" // 변수내용은 따로 지정 할 수도 있다
	test2 = "test string2" // 👉 위에서 지정한 값을 재지정해도 바뀔수 있다
  //--------------------------------------------

	// var num uint8 = 2900 // 👉256 을 초과하였기에 에러
	var num uint8 = 155 // 0 ~ 255 	

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(num)
	fmt.Println(test1, num)
	// 함수 안에 프린트 명령

	num = num + 100
	fmt.Println(num)
	// 👉 num 재지정 후 값은 255
}

// int vs uint
// int 와 uint의 차이는 마이너스(음수)가 있느냐 없느냐에 차이이다

// int8
// -2^7 ~ 2^7 - 1  (마이너스 2의 7승 부터 ~ 2의 7승 마이너스 1)

//--------------------------------------------

// uint256 = uint8 / uint = uint8 = uint256 같은 의미 / 보통 uint256 으로 많이 쓰인다
// 0 ~ 2^8 - 1 (0부터 ~ 2의 8승 마이너스 1)

//--------------------------------------------
// 자료형 		   저장범위 					설명
// uint8 	|    0 ~ 255  			    	|	부호 없는 8비트 정수형
// uint16 	|    0 ~ 65,535 	  	    	|	부호 없는 16비트 정수형
// unit32 	|    0 ~ 4,294,967,295	            	|	부호 없는 32비트 정수형
// uint64 	|    0 ~ 18,446,744,073,709,551,615 	|	부호 없는 64비트 정수형
// uint 	|				    	|	32비트 시스템에서는 uint32, 64비트 시스템에서는 uint64
// int8 	|    -128 ~ 127 		    	|	부호 있는 8비트 정수형
// int16 	|    -32,768 ~ 32,767 	 	    	|	부호 있는 16비트 정수형
// int32	|    -2,147,483,648 ~ 2,147,483,647 	|	부호 있는 32비트 정수형
// int64 	|    -9,223,372,036,854,775,808
// 		       ~ 9,223,372,036,854,775,807  	|	부호 있는 64비트 정수형
// int 		|								| 32비트 시스템에서는 int32, 64비트 시스템에서는 int64
// float32 	|				    	|	IEEE-754 32비트 부동소수점, 7자리 정밀도
// float64	|	     			    	|	IEEE-754 64비트 부동소수점, 12자리 정밀도
// complex64 	|		 		    	|	float32 크기의 실수부와 허수부로 구성된 복소수
// complex128 	|		 		    	|	float64 크기의 실수부와 허수부로 구성된 복소수
// uintptr 	|		 		    	|	uint와 같은 크기를 갖는 포인터형
// bool 	|				    	|	참, 거짓을 표현하기 위한 8비트 자료형
// byte 	|		 		    	|	8비트 자료형
// rune 	|		 		    	|	유니코드 저장을 위한 자료형, 크기는 int32와 동일
// string 	|		 		    	|	문자열을 저장하기 위한 자료형
