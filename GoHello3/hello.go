package main

import (
	
	"encoding/xml"	
	"fmt"
	
	"os"
)
type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"Version,attr"`
	Svs  []server `xml:"server"`
}

type server struct{
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}


//
func main() {
	 v:=&Servers{Version:"1"}
	 fmt.Println(v)
	 v.Svs = append(v.Svs, server{"Local_Web","172.0.0.1"})
	 fmt.Println(v)
	 v.Svs = append(v.Svs,server{"Local_DB","172.0.0.2"})
	 fmt.Println(v)
	 output,err := xml.MarshalIndent(v,"   ","	")
	 if err!=nil {
		 fmt.Println("error: %v\n",err)
		 return
	 }
	 os.Stdout.Write([]byte(xml.Header))
	 os.Stdout.Write(output)
}
