package main
import "fmt"
func main() {
	month:=15
	switch month {
	case 1:
		fmt.Println("janaory")
	case 2://agami kal ka mitting//
		fmt.Println("fub")
	case 3:
		fmt.Println("mrc")
	case 4:
		fmt.Println("apr")
	case 5:
		fmt.Println("mey")
	case 6:
		fmt.Println("jun")
	case 7:
		fmt.Println("jul")
	case 8:
		fmt.Println("agu")
	case 9:
		fmt.Println("sap")
	case 10:
		fmt.Println("oac")
	case 11:
		fmt.Println("nov")
	case 12:
		fmt.Println("dac")

	default:
		fmt.Println("this is not eait year")

	}
}


//bijor o jor sonka bar korar cord 1 thaka 10 ar moda
/*package main
import "fmt"
func main() {
	numbar:=9
	switch numbar {
	case 1,3,5,7,9:
		fmt.Println("of numbar")
	case 2,4,6,8,10:
		fmt.Println("evan numbar")
	default:
		fmt.Println("this numbar is not avalabal")

	}
	
}*/

//for loop
/*package main
import "fmt"
func main() {
	for i:=1; i<101; i+=2 {
		fmt.Println(i)
	}
	
}
*/
//continue cord
/*package main
import "fmt"
func main() {
	for i:=1; i<10; i++ { 
		if i ==7{
			continue
		}
		fmt.Println(i)

	}
}*/
//brack cord
/*package main
import "fmt"
func main() {
	for i:=1; i<10; i++ { 
		if i ==7{
			break
		}
		fmt.Println(i)
		}
	}
*/

/*package main
import ("fmt")

func main() {
  //adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
  */
//range cord
/*package main
import("fmt")
func main() {
	alfabat:=[] string {"a","b","c","d","e","f","g","h","i"}
	for idx, val := range alfabat{
		fmt.Printf("%v\t%v\n",idx,val)
	}
}*/

/*package main
import ("fmt")
func main() {
	//case
	//dufault
	//switch
	day:=4
	switch day {
	case 1:
		fmt.Println("saterday")
	case 2:
		fmt.Println("sunday")
	case 3:
		fmt.Println("monday")
	case 4:
		fmt.Println("tusday")
	case 5:
		fmt.Println("wednesday")
	case 6:
		fmt.Println("rhusrday")
	case 7:
		fmt.Println("friday")
	}
	
}
*/
