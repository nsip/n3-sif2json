package main

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/burntsushi/toml"
)

// !!! toml file name must be identical to config struct name !!!

// List2JSON :
type List2JSON struct {
	Path          string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	Identity struct { LIST []string }
	RoomInfo struct { LIST []string }
	StudentParticipation struct { LIST []string }
	TermInfo struct { LIST []string }
	StaffAssignment struct { LIST []string }
	StudentContactRelationship struct { LIST []string }
	TimeTableCell struct { LIST []string }
	TimeTableContainer struct { LIST []string }
	GradingAssignmentScore struct { LIST []string }
	NAPStudentResponseSet struct { LIST []string }
	NAPTestItem struct { LIST []string }
	ScheduledActivity struct { LIST []string }
	MarkValueInfo struct { LIST []string }
	SchoolInfo struct { LIST []string }
	StudentPeriodAttendance struct { LIST []string }
	AggregateCharacteristicInfo struct { LIST []string }
	AggregateStatisticFact struct { LIST []string }
	PersonPicture struct { LIST []string }
	PersonalisedPlan struct { LIST []string }
	Invoice struct { LIST []string }
	SectionInfo struct { LIST []string }
	StudentSchoolEnrollment struct { LIST []string }
	VendorInfo struct { LIST []string }
	WellbeingAlert struct { LIST []string }
	CalendarDate struct { LIST []string }
	NAPEventStudentLink struct { LIST []string }
	PaymentReceipt struct { LIST []string }
	SchoolPrograms struct { LIST []string }
	StudentActivityInfo struct { LIST []string }
	ChargedLocationInfo struct { LIST []string }
	Journal struct { LIST []string }
	NAPTestlet struct { LIST []string }
	ResourceBooking struct { LIST []string }
	CalendarSummary struct { LIST []string }
	NAPCodeFrame struct { LIST []string }
	SchoolCourseInfo struct { LIST []string }
	StaffPersonal struct { LIST []string }
	StudentSectionEnrollment struct { LIST []string }
	TeachingGroup struct { LIST []string }
	FinancialAccount struct { LIST []string }
	SessionInfo struct { LIST []string }
	StudentPersonal struct { LIST []string }
	StudentScoreJudgementAgainstStandard struct { LIST []string }
	TimeTableSubject struct { LIST []string }
	WellbeingCharacteristic struct { LIST []string }
	WellbeingResponse struct { LIST []string }
	Activity struct { LIST []string }
	LearningResource struct { LIST []string }
	ResourceUsage struct { LIST []string }
	StudentContactPersonal struct { LIST []string }
	NAPTestScoreSummary struct { LIST []string }
	TimeTable struct { LIST []string }
	WellbeingAppeal struct { LIST []string }
	CollectionRound struct { LIST []string }
	EquipmentInfo struct { LIST []string }
	FinancialQuestionnaireCollection struct { LIST []string }
	LearningStandardItem struct { LIST []string }
	CensusCollection struct { LIST []string }
	StudentAttendanceSummary struct { LIST []string }
	StudentAttendanceTimeList struct { LIST []string }
	SystemRole struct { LIST []string }
	AddressCollection struct { LIST []string }
	NAPTest struct { LIST []string }
	CollectionStatus struct { LIST []string }
	WellbeingPersonLink struct { LIST []string }
	AggregateStatisticInfo struct { LIST []string }
	Debtor struct { LIST []string }
	StudentGrade struct { LIST []string }
	WellbeingEvent struct { LIST []string }
	StudentActivityParticipation struct { LIST []string }
	StudentDailyAttendance struct { LIST []string }
	GradingAssignment struct { LIST []string }
	LEAInfo struct { LIST []string }
	LearningStandardDocument struct { LIST []string }
	PurchaseOrder struct { LIST []string }
	
}

// Num2JSON :
type Num2JSON struct {
	Path          string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	LEAInfo struct { NUMERIC []string }
	PersonPicture struct { NUMERIC []string }
	SessionInfo struct { NUMERIC []string }
	StudentActivityParticipation struct { NUMERIC []string }
	CalendarDate struct { NUMERIC []string }
	ChargedLocationInfo struct { NUMERIC []string }
	GradingAssignment struct { NUMERIC []string }
	Journal struct { NUMERIC []string }
	ResourceBooking struct { NUMERIC []string }
	ResourceUsage struct { NUMERIC []string }
	SectionInfo struct { NUMERIC []string }
	StaffPersonal struct { NUMERIC []string }
	Debtor struct { NUMERIC []string }
	EquipmentInfo struct { NUMERIC []string }
	StudentAttendanceSummary struct { NUMERIC []string }
	WellbeingPersonLink struct { NUMERIC []string }
	TimeTableSubject struct { NUMERIC []string }
	StudentParticipation struct { NUMERIC []string }
	TermInfo struct { NUMERIC []string }
	WellbeingEvent struct { NUMERIC []string }
	SystemRole struct { NUMERIC []string }
	TimeTable struct { NUMERIC []string }
	StudentContactPersonal struct { NUMERIC []string }
	TeachingGroup struct { NUMERIC []string }
	TimeTableContainer struct { NUMERIC []string }
	WellbeingCharacteristic struct { NUMERIC []string }
	MarkValueInfo struct { NUMERIC []string }
	SchoolCourseInfo struct { NUMERIC []string }
	StudentScoreJudgementAgainstStandard struct { NUMERIC []string }
	WellbeingResponse struct { NUMERIC []string }
	PersonalisedPlan struct { NUMERIC []string }
	StaffAssignment struct { NUMERIC []string }
	CollectionRound struct { NUMERIC []string }
	Identity struct { NUMERIC []string }
	PurchaseOrder struct { NUMERIC []string }
	WellbeingAppeal struct { NUMERIC []string }
	Activity struct { NUMERIC []string }
	AggregateStatisticFact struct { NUMERIC []string }
	FinancialQuestionnaireCollection struct { NUMERIC []string }
	NAPTestItem struct { NUMERIC []string }
	ScheduledActivity struct { NUMERIC []string }
	StudentContactRelationship struct { NUMERIC []string }
	StudentGrade struct { NUMERIC []string }
	StudentPersonal struct { NUMERIC []string }
	CalendarSummary struct { NUMERIC []string }
	CollectionStatus struct { NUMERIC []string }
	TimeTableCell struct { NUMERIC []string }
	SchoolInfo struct { NUMERIC []string }
	SchoolPrograms struct { NUMERIC []string }
	StudentDailyAttendance struct { NUMERIC []string }
	NAPCodeFrame struct { NUMERIC []string }
	NAPTest struct { NUMERIC []string }
	NAPStudentResponseSet struct { NUMERIC []string }
	FinancialAccount struct { NUMERIC []string }
	NAPEventStudentLink struct { NUMERIC []string }
	RoomInfo struct { NUMERIC []string }
	VendorInfo struct { NUMERIC []string }
	WellbeingAlert struct { NUMERIC []string }
	AddressCollection struct { NUMERIC []string }
	Invoice struct { NUMERIC []string }
	LearningStandardDocument struct { NUMERIC []string }
	StudentPeriodAttendance struct { NUMERIC []string }
	AggregateCharacteristicInfo struct { NUMERIC []string }
	CensusCollection struct { NUMERIC []string }
	LearningStandardItem struct { NUMERIC []string }
	StudentActivityInfo struct { NUMERIC []string }
	StudentAttendanceTimeList struct { NUMERIC []string }
	AggregateStatisticInfo struct { NUMERIC []string }
	GradingAssignmentScore struct { NUMERIC []string }
	StudentSectionEnrollment struct { NUMERIC []string }
	NAPTestScoreSummary struct { NUMERIC []string }
	StudentSchoolEnrollment struct { NUMERIC []string }
	PaymentReceipt struct { NUMERIC []string }
	LearningResource struct { NUMERIC []string }
	NAPTestlet struct { NUMERIC []string }
	
}

// Bool2JSON :
type Bool2JSON struct {
	Path          string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	SystemRole struct { BOOLEAN []string }
	WellbeingResponse struct { BOOLEAN []string }
	GradingAssignment struct { BOOLEAN []string }
	Invoice struct { BOOLEAN []string }
	PaymentReceipt struct { BOOLEAN []string }
	StudentDailyAttendance struct { BOOLEAN []string }
	SchoolPrograms struct { BOOLEAN []string }
	StudentAttendanceSummary struct { BOOLEAN []string }
	WellbeingEvent struct { BOOLEAN []string }
	EquipmentInfo struct { BOOLEAN []string }
	MarkValueInfo struct { BOOLEAN []string }
	ResourceBooking struct { BOOLEAN []string }
	RoomInfo struct { BOOLEAN []string }
	StudentPeriodAttendance struct { BOOLEAN []string }
	CalendarDate struct { BOOLEAN []string }
	CensusCollection struct { BOOLEAN []string }
	FinancialAccount struct { BOOLEAN []string }
	SchoolCourseInfo struct { BOOLEAN []string }
	ChargedLocationInfo struct { BOOLEAN []string }
	PurchaseOrder struct { BOOLEAN []string }
	StudentParticipation struct { BOOLEAN []string }
	Debtor struct { BOOLEAN []string }
	NAPTestItem struct { BOOLEAN []string }
	TimeTableContainer struct { BOOLEAN []string }
	AddressCollection struct { BOOLEAN []string }
	AggregateCharacteristicInfo struct { BOOLEAN []string }
	AggregateStatisticInfo struct { BOOLEAN []string }
	StudentContactPersonal struct { BOOLEAN []string }
	StudentActivityParticipation struct { BOOLEAN []string }
	NAPCodeFrame struct { BOOLEAN []string }
	NAPTestlet struct { BOOLEAN []string }
	ScheduledActivity struct { BOOLEAN []string }
	StaffAssignment struct { BOOLEAN []string }
	CalendarSummary struct { BOOLEAN []string }
	GradingAssignmentScore struct { BOOLEAN []string }
	Identity struct { BOOLEAN []string }
	Journal struct { BOOLEAN []string }
	Activity struct { BOOLEAN []string }
	CollectionStatus struct { BOOLEAN []string }
	LEAInfo struct { BOOLEAN []string }
	NAPTestScoreSummary struct { BOOLEAN []string }
	ResourceUsage struct { BOOLEAN []string }
	SectionInfo struct { BOOLEAN []string }
	StudentContactRelationship struct { BOOLEAN []string }
	StudentGrade struct { BOOLEAN []string }
	LearningResource struct { BOOLEAN []string }
	LearningStandardDocument struct { BOOLEAN []string }
	NAPEventStudentLink struct { BOOLEAN []string }
	NAPStudentResponseSet struct { BOOLEAN []string }
	TimeTableSubject struct { BOOLEAN []string }
	TimeTable struct { BOOLEAN []string }
	LearningStandardItem struct { BOOLEAN []string }
	PersonPicture struct { BOOLEAN []string }
	StudentAttendanceTimeList struct { BOOLEAN []string }
	StudentSectionEnrollment struct { BOOLEAN []string }
	StudentScoreJudgementAgainstStandard struct { BOOLEAN []string }
	TimeTableCell struct { BOOLEAN []string }
	VendorInfo struct { BOOLEAN []string }
	WellbeingAppeal struct { BOOLEAN []string }
	FinancialQuestionnaireCollection struct { BOOLEAN []string }
	NAPTest struct { BOOLEAN []string }
	SchoolInfo struct { BOOLEAN []string }
	StudentSchoolEnrollment struct { BOOLEAN []string }
	WellbeingPersonLink struct { BOOLEAN []string }
	CollectionRound struct { BOOLEAN []string }
	PersonalisedPlan struct { BOOLEAN []string }
	TeachingGroup struct { BOOLEAN []string }
	WellbeingCharacteristic struct { BOOLEAN []string }
	StudentPersonal struct { BOOLEAN []string }
	WellbeingAlert struct { BOOLEAN []string }
	AggregateStatisticFact struct { BOOLEAN []string }
	SessionInfo struct { BOOLEAN []string }
	StaffPersonal struct { BOOLEAN []string }
	StudentActivityInfo struct { BOOLEAN []string }
	TermInfo struct { BOOLEAN []string }
	
}

var (
	// toml file name must be identical to config struct name
	lsCfg = []interface{}{
		&List2JSON{},
		&Num2JSON{},
		&Bool2JSON{},		
	}
)

// ------------------------------------------------- //

// NewCfg :
func NewCfg(configs ...string) interface{} {
	for _, f := range configs {
		if _, e := os.Stat(f); e == nil {
			if abs, e := filepath.Abs(f); e == nil {
				return set(f, abs)
			}
		}
	}
	return nil
}

func set(f, abs string) interface{} {
	for _, cfg := range lsCfg {
		name := reflect.TypeOf(cfg).Elem().Name()
		if sHasSuffix(f, "/"+name+".toml") {
			if _, e := toml.DecodeFile(f, cfg); e == nil {
				reflect.ValueOf(cfg).Elem().FieldByName("Path").SetString(abs)
				save(f, cfg)
				// modify for runtime
				return cfg
			}
		}
	}
	return nil
}

func save(path string, cfg interface{}) {
	if f, e := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, os.ModePerm); e == nil {
		defer f.Close()
		toml.NewEncoder(f).Encode(cfg)
	}
}
