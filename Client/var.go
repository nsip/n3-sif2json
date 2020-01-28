package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"

	glb "github.com/nsip/n3-sif2json/Client/global"
)

var (
	fPf   = fmt.Printf
	fPln  = fmt.Println
	fSf   = fmt.Sprintf
	fEf   = fmt.Errorf
	sJoin = strings.Join
)

func isValidXML(data []byte) bool {
	return xml.Unmarshal(data, new(interface{})) == nil
}

var (
	mFnURL = map[string]string{}
)

func initMapFnURL(protocol, ip string, port int) bool {
	v := reflect.ValueOf(glb.Cfg.Route)
	typeOfT := reflect.ValueOf(&glb.Cfg.Route).Elem().Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfT.Field(i).Name
		value := v.Field(i).Interface().(string)
		mFnURL[field] = fSf("%s://%s:%d%s", protocol, ip, port, value)
	}
	return len(mFnURL) > 0
}

func getCfgRouteFields() (fields []string) {
	v := reflect.ValueOf(glb.Cfg.Route)
	// typeOfT := reflect.ValueOf(&glb.Cfg.Route).Elem().Type()
	typeOfT := reflect.ValueOf(glb.Cfg.Route).Type()
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, typeOfT.Field(i).Name)
	}
	return
}
