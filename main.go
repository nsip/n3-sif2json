package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	pp "./preprocess"
	xj "github.com/basgys/goxml2json"
	"github.com/clbanning/mxj"
)

func main() {

	xmlstr := `
	<?xml version="1.0" encoding="UTF-8"?>
	<osm version="-0.6" generator="CGImap 0.0.2">
	 <bounds minlat="54.0889580" minlon="12.2487570" maxlat="54.0913900" maxlon="12.2524800">True</bounds>
	 <foo>bar</foo>	 
	</osm>
	`

	b, _ := ioutil.ReadFile("test.xml")
	xmlstr = string(b)
	fmt.Println(xmlstr)

	// xml is an io.Reader
	xmlReader := strings.NewReader(xmlstr)
	jsonbuf, err := xj.Convert(
		xmlReader,
		xj.WithTypeConverter(xj.Float, xj.Int, xj.Bool, xj.Null),
		//xj.WithAttrPrefix(""),
		xj.WithContentPrefix("="),
	)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(jsonbuf.String())
	// {"hello": "world"}

	jsonfmt := pp.FmtJSONStr(jsonbuf.String(), "./preprocess/utils")
	fmt.Println(jsonfmt)

	ioutil.WriteFile("test.json", []byte(jsonfmt), 0666)

	// return

	// convert back to xml
	var f interface{}
	if err = json.Unmarshal([]byte(jsonfmt), &f); err != nil {
		panic("1")
	}

	fmt.Println(" --------------- ")
	fmt.Printf("%v", f)

	b, err = mxj.AnyXmlIndent(f, "", "    ", "")

	re1 := regexp.MustCompile("\n[ ]*<=content>")
	re2 := regexp.MustCompile("</=content>\n[ ]*")

	xmlstr = string(b)
	xmlstr = strings.ReplaceAll(xmlstr, "<>", "")
	xmlstr = strings.ReplaceAll(xmlstr, "</>", "")
	//xmlstr = strings.ReplaceAll(xmlstr, " *** <=content>", "")
	//xmlstr = strings.ReplaceAll(xmlstr, "</=content> *** ", "")

	xmlstr = re1.ReplaceAllString(xmlstr, "")
	xmlstr = re2.ReplaceAllString(xmlstr, "")

	// if b, err = xml.Marshal(&f); err != nil {
	// 	panic("2")
	// }
	ioutil.WriteFile("test1.xml", []byte(xmlstr), 0666)

	return

	// jkv := jkv.NewJKV(jsonfmt, "")
	// jsontmp, lvl := jkv.Unfold(-1, nil)
	// fmt.Println(jsontmp, lvl)

	// c := cfg.NewCfg("./config.toml")
	// fmt.Println(c.MustBeArray)
	// fmt.Println(cfg.ListFieldValue(c.MustBeArray))
}
