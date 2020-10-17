package main

import (
	"testing"

	xj "github.com/basgys/goxml2json"
	"github.com/davecgh/go-spew/spew"
	cfg "github.com/nsip/n3-sif2json/Config/cfg"
)

func TestMain(t *testing.T) {
	main()
}

func TestLoad(t *testing.T) {
	c := cfg.NewCfg(
		"Config",
		map[string]string{
			"[s]": "Service",
			"[v]": "Version",
		},
		"../Config/config.toml",
	).(*cfg.Config)
	spew.Dump(*c)
}

func TestInit(t *testing.T) {
	c := cfg.NewCfg(
		"Config",
		map[string]string{
			"[s]":    "Service",
			"[v]":    "Version",
			"[port]": "WebService.Port",
		},
		"../Config/config.toml",
	).(*cfg.Config)
	spew.Dump(*c)

	c = env2Struct("Config", &cfg.Config{}).(*cfg.Config)
	spew.Dump(*c)
}

func TestXMLLvl0(t *testing.T) {
	xml := `<Activity RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A">
	<Title>Shakespeare Essay - Much Ado About Nothing</Title>
	<Preamble>This is a very funny comedy - students should have passing familiarity with Shakespeare</Preamble>
	<LearningStandards>
	<LearningStandardItemRefId>9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4</LearningStandardItemRefId>
	</LearningStandards>
	<SourceObjects>
	<SourceObject SIF_RefObject="Lesson">A71ADBD3-D93D-A64B-7166-E420D50EDABC</SourceObject>
	</SourceObjects>
	<Points>50</Points>
	<ActivityTime>
	<CreationDate>2002-06-15</CreationDate>
	<Duration Units="minute">30</Duration>
	<StartDate>2002-09-10</StartDate>
	<FinishDate>2002-09-12</FinishDate>
	<DueDate>2002-09-12</DueDate>
	</ActivityTime>
	<AssessmentRefId>03EDB29E-8116-B450-0435-FA87E42A0AD2</AssessmentRefId>
	<MaxAttemptsAllowed>3</MaxAttemptsAllowed>
	<ActivityWeight>5</ActivityWeight>
	<Evaluation EvaluationType="Inline">
	<Description>Students should be able to correctly identify all major characters.</Description>
	</Evaluation>
	<LearningResources>
	<LearningResourceRefId>B7337698-BF6D-B193-7F79-A07B87211B93</LearningResourceRefId>
	</LearningResources>
	</Activity>`

	name, out, in := XMLLvl0(xml)
	fPln(name)
	fPln(out)
	jsonBuf, _ := xj.Convert(sNewReader(out))
	fPln(jsonBuf.String())
	fPln(" ---------------------------------- ")
	fPln(in)

	sr, _ := XMLBreakCont(in)
	for _, r := range sr {
		fPln(r)
	}
}

func TestJSONBlkCont(t *testing.T) {
	jstr := `{
		"Activity": {
		  "-RefId": "C27E1FCF-C163-485F-BEF0-F36F18A0493A",
		  "ActivityTime": {
			"CreationDate": "2002-06-15",
			"DueDate": "2002-09-12",
			"Duration": {
			  "#content": 30,
			  "-Units": "minute"
			},
			"FinishDate": "2002-09-12",
			"StartDate": "2002-09-10"
		  },
		  "ActivityWeight": 5,
		  "AssessmentRefId": "03EDB29E-8116-B450-0435-FA87E42A0AD2",
		  "Evaluation": {
			"-EvaluationType": "Inline",
			"Description": "Students should be able to correctly identify all major characters."
		  },
		  "LearningResources": {
			"LearningResourceRefId": [
			  "B7337698-BF6D-B193-7F79-A07B87211B93"
			]
		  },
		  "LearningStandards": {
			"LearningStandardItemRefId": [
			  "9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4"
			]
		  },
		  "MaxAttemptsAllowed": 3,
		  "Points": 50,
		  "Preamble": "This is a very funny comedy - students should have passing familiarity with Shakespeare",
		  "SourceObjects": {
			"SourceObject": [
			  {
				"#content": "A71ADBD3-D93D-A64B-7166-E420D50EDABC",
				"-SIF_RefObject": "Lesson"
			  }
			]
		  },
		  "Title": "Shakespeare Essay - Much Ado About Nothing"
		}
	  }`

	name, cont := JSONBlkCont(jstr)
	fPln(name)
	fPln(cont)
}
