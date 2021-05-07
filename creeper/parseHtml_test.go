package creeper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"testing"
)

func Test_ParseHtml(t *testing.T){
	f, err := os.Open("E:\\go_project\\videoProject\\server\\video\\creeper\\test.html")
	if err != nil {
		t.Error(err)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	p := doc.Find("#p2")
	p.Each(func(i int, s *goquery.Selection){
		fmt.Println(i)
		fmt.Println(s.Attr("id"))
		fmt.Println(s.Text())
		fmt.Println()
	})


}
