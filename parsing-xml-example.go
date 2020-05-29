package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml"
)

type sitemapIndex struct{
	Location []Location `xml:"sitemap"`
}


type Location struct{

	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main(){
 resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
 bytes, _ := ioutil.ReadAll(resp.Body)
 //string_body := string(bytes)
 //fmt.Println(string_body)
 resp.Body.Close()
 var s sitemapIndex
 xml.Unmarshal(bytes, &s)
 //fmt.Println(s.Location)	


 for _, Location := range s.Location{
 	fmt.Printf("\n %s",Location)
 }

}
