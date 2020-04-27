package main

import (
	"fmt"
	"reflect"
	"strings"

	cmn "github.com/cdutwhu/json-util/common"
	glb "github.com/nsip/n3-sif2json/Client/global"
)

var (
	fPt  = fmt.Print
	fPf  = fmt.Printf
	fPln = fmt.Println
	fSf  = fmt.Sprintf

	sJoin      = strings.Join
	sHasSuffix = strings.HasSuffix
	sTrimRight = strings.TrimRight
	sReplace   = strings.Replace

	isFLog        = cmn.IsFLog
	failOnErrWhen = cmn.FailOnErrWhen
	failOnErr     = cmn.FailOnErr
	warnOnErrWhen = cmn.WarnOnErrWhen
	warnOnErr     = cmn.WarnOnErr
	setLog        = cmn.SetLog
	resetLog      = cmn.ResetLog
	isXML         = cmn.IsXML
	isJSON        = cmn.IsJSON
)

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