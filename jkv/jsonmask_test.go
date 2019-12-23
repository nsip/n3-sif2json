package jkv

import (
	"io/ioutil"
	"sync"
	"testing"
	"time"

	cmn "github.com/nsip/n3-privacy/common"
	pp "github.com/nsip/n3-privacy/preprocess"
)

func TestJSONPolicy(t *testing.T) {
	defer cmn.TmTrack(time.Now())
	data := pp.FmtJSONFile("../../JSON-Mask/data/NAPCodeFrame.json", "../preprocess/utils/")
	mask := pp.FmtJSONFile("../../JSON-Mask/data/NAPCodeFrameMaskP.json", "../preprocess/utils/")

	if data == "" {
		panic("input data is empty, check its path")
	}
	if mask == "" {
		panic("input mask is empty, check its path")
	}

	jkvM := NewJKV(mask, "root")

	if IsJSONArr(data) {
		jsonArr := SplitJSONArr(data)
		wg := sync.WaitGroup{}
		wg.Add(len(jsonArr))
		jsonList := make([]string, len(jsonArr))
		for i, json := range jsonArr {
			go func(i int, json string) {
				defer wg.Done()
				jkvD := NewJKV(json, "root")
				maskroot, _ := jkvD.Unfold(0, jkvM)
				jkvMR := NewJKV(maskroot, "")
				jkvMR.Wrapped = jkvD.Wrapped
				jsonList[i] = jkvMR.UnwrapDefault().JSON
			}(i, json)
		}
		wg.Wait()
		ioutil.WriteFile("array.json", []byte(MergeJSON(jsonList...)), 0666)

	} else {
		jkvD := NewJKV(data, "root")
		maskroot, _ := jkvD.Unfold(0, jkvM)
		jkvMR := NewJKV(maskroot, "")
		jkvMR.Wrapped = jkvD.Wrapped
		json := jkvMR.UnwrapDefault().JSON
		ioutil.WriteFile("single.json", []byte(json), 0666)
	}
}
