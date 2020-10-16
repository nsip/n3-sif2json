package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3log"
	"github.com/cdutwhu/n3-util/rest"
)

var (
	fPln             = fmt.Println
	fSf              = fmt.Sprintf
	fPf              = fmt.Printf
	sReplaceAll      = strings.ReplaceAll
	sTrimRight       = strings.TrimRight
	sTrimLeft        = strings.TrimLeft
	sIndex           = strings.Index
	sLastIndex       = strings.LastIndex
	sJoin            = strings.Join
	sNewReader       = strings.NewReader
	sHasSuffix       = strings.HasSuffix
	rxMustCompile    = regexp.MustCompile
	failOnErr        = fn.FailOnErr
	failOnErrWhen    = fn.FailOnErrWhen
	enableLog2F      = fn.EnableLog2F
	enableWarnDetail = fn.EnableWarnDetail
	logWhen          = fn.LoggerWhen
	logger           = fn.Logger
	warnOnErr        = fn.WarnOnErr
	warner           = fn.Warner
	localIP          = net.LocalIP
	env2Struct       = rflx.Env2Struct
	struct2Map       = rflx.Struct2Map
	tryInvoke        = rflx.TryInvoke
	loggly           = n3log.Loggly
	logBind          = n3log.Bind
	setLoggly        = n3log.SetLoggly
	syncBindLog      = n3log.SyncBindLog
	isXML            = judge.IsXML
	isJSON           = judge.IsJSON
	mustWriteFile    = io.MustWriteFile
	url1Value        = rest.URL1Value
)

var (
	logGrp  = logBind(logger) // logBind(logger, loggly("info"))
	warnGrp = logBind(warner) // logBind(warner, loggly("warn"))

	rxTag  = rxMustCompile(`<\w+[\s/>]`)
	rxHead = rxMustCompile(`<\w+(\s+[\w:]+\s*=\s*"[^"]*"\s*)*\s*/?>`)
)

func initMutex(route interface{}) map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range struct2Map(route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

// XMLBreakCont :
func XMLBreakCont(xml string) (roots, subs []string) {

	remain := xml
AGAIN:
	if loc := rxTag.FindStringIndex(remain); loc != nil {
		s, e := loc[0], loc[1]
		root := remain[s+1 : e-1]
		roots = append(roots, root)

		remain = remain[s:] // from first '<tag>'
		// fPln("remain:", remain)

		end1, end2 := -1, -1
		if loc := rxMustCompile(fSf(`</%s\s*>`, root)).FindStringIndex(remain); loc != nil {
			_, end1 = loc[0], loc[1] // update e to '</tag>' end
			// fPln("end:", remain[s:end1]) // end tag
		}
		if i := sIndex(remain, "/>"); i >= 0 {
			end2 = i + 2 // update e to '/>' end
		}

		// if '/>' is found, and before '</tag>', and this part is valid XML
		switch {
		case end1 >= 0 && end2 < 0:
			e = end1
		case end1 < 0 && end2 >= 0:
			e = end2
		case end1 >= 0 && end2 >= 0:
			if end2 < end1 && isXML(remain[:end2]) {
				e = end2
			} else {
				e = end1
			}
		default:
			panic("invalid sub xml")
		}

		sub := remain[:e]
		// fPln("sub:", sub)
		subs = append(subs, sub)

		remain = remain[e:] // from end of first '</tag>' or '/>'
		goto AGAIN
	}

	return
}

// XMLLvl0 :
func XMLLvl0(xml string) (string, string, string) {
	sTag, name, eTag := "", "", ""
	end1, end2 := 0, 0

	if loc := rxHead.FindStringIndex(xml); loc != nil {
		s, e := loc[0], loc[1]
		sTag = xml[s:e]
		// fPln(1, sTag)
		end1 = e

		if loc := rxTag.FindStringIndex(sTag); loc != nil {
			s, e := loc[0], loc[1]
			name = sTrimRight(sTag[s+1:e], " \t/>")
			// fPln(2, name)

			eTag = fSf(`</%s>`, name)
			// fPln(3, eTag)
			end2 = sLastIndex(xml, eTag)
		}
	}
	return name, sJoin([]string{sTag, eTag}, "\n"), xml[end1:end2]
}
