package  main

import "fmt"
func main(){

	strs:="abcabdabcabfabcabdabcabcddd"
	str :="abcabc"

	for i:=0;i<=len(strs)-len(str);i++{
		j:=0
		if strs[i] == str[0]{
			for k:=0;k<len(str);k++{
				if strs[k+i]==str[k]{
					j++
				}else{
					break
				}
			}
			if j== len(str){
				fmt.Println("match!")
			}
		}


	}
}
