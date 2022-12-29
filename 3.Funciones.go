package  main
        
import (
    "fmt"
    "strconv"
)
/*-----------------------------------------------Numeros------------------------------------*/
func Serie(n int){ //void
  for n>0 {  //while
  	var d int = n % 10
  	n = n / 10
  	fmt.Println(d)
  }
}

func digitoMenor(n int) int {
  var men int = 10
  var d int
  for n>0 {
  	d = n % 10
  	n = n / 10
  	if d<men {
 		men = d 	
  	}
  }
  return men
}

func digitoMayor(n int) int {
  var may int = 0
  var d int
  for n>0 {
  	d = n % 10
  	n = n / 10
  	if d>may {
 		may = d 	
  	}
  }
  return may
}

func Posicion(n,x int) byte {
  var b bool = false
  var d int 
  var pos byte = 0
  for n>0 && b ==false {
  		d = n % 10
  		n = n / 10
  		pos++
  		if d == x {
  			b = true
  		}
  }

  if (b==false){
  	 pos = 0
  }
  return pos
}

func CntDigitos(n int) int {
	var c int = 0
	for n>0 {
		n = n / 10
		c++
	}
	return c
}

func Power(n int,x byte) int {
	var p int
	var i byte
	p=1
	for i = 0; i<x; i++{
		p = p * n
	}
	return p
}

func EliminarDigito(n int,x byte) int{
	var n1,n2,k int
	k= Power(10,x-1)
	n1 = n / k      
	n2 = n % k      
	n1 = n1 / 10   		
	return n1  * k + n2
}

func OrdenarDeAdentroAfuera (n int) int {
   var c int =  CntDigitos(n)
   var acum int = 0
   var k int = 1
   var t bool = true
   var d int
   var pos byte
   for c>0 {
   		d = digitoMenor(n)
   		if t==true {
   			acum =acum * 10 + d
   			t = false 
   		}else{
   			acum = d * k + acum
   			t = true
   		}
   		k=k * 10
   		pos = Posicion(n,d)
   		n = EliminarDigito(n,pos)
   		c--
   }
   return acum
}
/*-----------------------------------------------Cadenas------------------------------------*/
func siguientePalabra(i *int,x string) string {
 
	for *i<len(x) && string(x[*i])==" " {
	  *i++
	}

	p :=""
	for *i<len(x) && string(x[*i])!=" " {
	  	p=p + string(x[*i])
	  	*i++
	}

  return p
}

func ToString(x string) string {
  //var x string = "hola mundo programador"
  p := ""
  i:=0
  fmt.Println("Longitud:",len(x))
  for j:=0 ; j < len(x) ;j++ {
  	p = p + siguientePalabra(&i,x) + " "
  }
  return p 
}

/*-----------------------------------------------VECTORES-----------------------------------*/
func cargarVector() {
  c := "[ " 
  var v [10]int
  var p int = 0
	for i:=0 ; i<10;i++{
		fmt.Scanln(&p) 
		v[i]=p
		if i<10-1 {
			c += strconv.Itoa(v[i]) + " , "
		}else{
			c += strconv.Itoa(v[i]) + " ]"
		}
	}
	fmt.Println(c)
}

/*-----------------------------------------------MATRICES-----------------------------------*/
func cargarMatriz() {
  c:= "" 
  var v [3][3]int
  var p int = 0
	for i:=0 ; i<3;i++{
		c += "| "
		for j:=0 ; j<3;j++{
			fmt.Scanln(&p) 
			v[i][j]=p
			c += strconv.Itoa(v[i][j]) + " "
		}
		c += "|\n"
	}
	fmt.Println(c)
}

func main(){
	x:=""
	fmt.Print("Introduzca cadena: ")
	fmt.Scanf("%s",&x)
  fmt.Println(ToString(x))
}

