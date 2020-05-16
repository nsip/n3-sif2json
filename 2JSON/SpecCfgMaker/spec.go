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
	CollectionRound struct { LIST []string }
	FinancialQuestionnaireCollection struct { LIST []string }
	ResourceUsage struct { LIST []string }
	ScheduledActivity struct { LIST []string }
	WellbeingPersonLink struct { LIST []string }
	EquipmentInfo struct { LIST []string }
	GradingAssignmentScore struct { LIST []string }
	Identity struct { LIST []string }
	SessionInfo struct { LIST []string }
	VendorInfo struct { LIST []string }
	WellbeingCharacteristic struct { LIST []string }
	AggregateStatisticInfo struct { LIST []string }
	PaymentReceipt struct { LIST []string }
	StudentSchoolEnrollment struct { LIST []string }
	SystemRole struct { LIST []string }
	CalendarDate struct { LIST []string }
	Debtor struct { LIST []string }
	NAPStudentResponseSet struct { LIST []string }
	PurchaseOrder struct { LIST []string }
	SchoolCourseInfo struct { LIST []string }
	StudentPersonal struct { LIST []string }
	WellbeingEvent struct { LIST []string }
	NAPTestScoreSummary struct { LIST []string }
	TimeTable struct { LIST []string }
	TimeTableContainer struct { LIST []string }
	CollectionStatus struct { LIST []string }
	StaffPersonal struct { LIST []string }
	StudentAttendanceTimeList struct { LIST []string }
	StudentPeriodAttendance struct { LIST []string }
	StudentScoreJudgementAgainstStandard struct { LIST []string }
	Activity struct { LIST []string }
	LearningStandardDocument struct { LIST []string }
	ResourceBooking struct { LIST []string }
	RoomInfo struct { LIST []string }
	TermInfo struct { LIST []string }
	LearningStandardItem struct { LIST []string }
	NAPCodeFrame struct { LIST []string }
	NAPEventStudentLink struct { LIST []string }
	NAPTest struct { LIST []string }
	StaffAssignment struct { LIST []string }
	StudentActivityParticipation struct { LIST []string }
	TeachingGroup struct { LIST []string }
	TimeTableCell struct { LIST []string }
	WellbeingAppeal struct { LIST []string }
	AddressCollection struct { LIST []string }
	CensusCollection struct { LIST []string }
	NAPTestItem struct { LIST []string }
	SchoolInfo struct { LIST []string }
	StudentDailyAttendance struct { LIST []string }
	GradingAssignment struct { LIST []string }
	Invoice struct { LIST []string }
	StudentActivityInfo struct { LIST []string }
	StudentAttendanceSummary struct { LIST []string }
	TimeTableSubject struct { LIST []string }
	ChargedLocationInfo struct { LIST []string }
	LEAInfo struct { LIST []string }
	PersonalisedPlan struct { LIST []string }
	SchoolPrograms struct { LIST []string }
	StudentContactPersonal struct { LIST []string }
	WellbeingResponse struct { LIST []string }
	Journal struct { LIST []string }
	SectionInfo struct { LIST []string }
	StudentContactRelationship struct { LIST []string }
	WellbeingAlert struct { LIST []string }
	CalendarSummary struct { LIST []string }
	PersonPicture struct { LIST []string }
	StudentSectionEnrollment struct { LIST []string }
	AggregateCharacteristicInfo struct { LIST []string }
	AggregateStatisticFact struct { LIST []string }
	FinancialAccount struct { LIST []string }
	LearningResource struct { LIST []string }
	MarkValueInfo struct { LIST []string }
	NAPTestlet struct { LIST []string }
	StudentGrade struct { LIST []string }
	StudentParticipation struct { LIST []string }
	
}

// Num2JSON :
type Num2JSON struct {
	Path          string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	ResourceBooking struct { NUMERIC []string }
	StudentSchoolEnrollment struct { NUMERIC []string }
	StudentScoreJudgementAgainstStandard struct { NUMERIC []string }
	StudentSectionEnrollment struct { NUMERIC []string }
	TeachingGroup struct { NUMERIC []string }
	Debtor struct { NUMERIC []string }
	LearningStandardDocument struct { NUMERIC []string }
	PaymentReceipt struct { NUMERIC []string }
	TimeTableSubject struct { NUMERIC []string }
	WellbeingAlert struct { NUMERIC []string }
	SessionInfo struct { NUMERIC []string }
	StudentActivityInfo struct { NUMERIC []string }
	StudentParticipation struct { NUMERIC []string }
	AddressCollection struct { NUMERIC []string }
	GradingAssignmentScore struct { NUMERIC []string }
	PersonalisedPlan struct { NUMERIC []string }
	StudentAttendanceSummary struct { NUMERIC []string }
	AggregateStatisticFact struct { NUMERIC []string }
	GradingAssignment struct { NUMERIC []string }
	NAPTestlet struct { NUMERIC []string }
	StudentPeriodAttendance struct { NUMERIC []string }
	WellbeingEvent struct { NUMERIC []string }
	PersonPicture struct { NUMERIC []string }
	PurchaseOrder struct { NUMERIC []string }
	StaffAssignment struct { NUMERIC []string }
	StudentPersonal struct { NUMERIC []string }
	WellbeingResponse struct { NUMERIC []string }
	CensusCollection struct { NUMERIC []string }
	MarkValueInfo struct { NUMERIC []string }
	ScheduledActivity struct { NUMERIC []string }
	RoomInfo struct { NUMERIC []string }
	StudentContactRelationship struct { NUMERIC []string }
	SystemRole struct { NUMERIC []string }
	TermInfo struct { NUMERIC []string }
	Identity struct { NUMERIC []string }
	LearningStandardItem struct { NUMERIC []string }
	NAPTest struct { NUMERIC []string }
	StudentGrade struct { NUMERIC []string }
	LEAInfo struct { NUMERIC []string }
	NAPTestItem struct { NUMERIC []string }
	StudentDailyAttendance struct { NUMERIC []string }
	WellbeingAppeal struct { NUMERIC []string }
	Invoice struct { NUMERIC []string }
	SchoolCourseInfo struct { NUMERIC []string }
	StudentActivityParticipation struct { NUMERIC []string }
	ResourceUsage struct { NUMERIC []string }
	WellbeingCharacteristic struct { NUMERIC []string }
	WellbeingPersonLink struct { NUMERIC []string }
	ChargedLocationInfo struct { NUMERIC []string }
	EquipmentInfo struct { NUMERIC []string }
	Journal struct { NUMERIC []string }
	SchoolPrograms struct { NUMERIC []string }
	TimeTable struct { NUMERIC []string }
	CollectionRound struct { NUMERIC []string }
	NAPCodeFrame struct { NUMERIC []string }
	NAPEventStudentLink struct { NUMERIC []string }
	FinancialQuestionnaireCollection struct { NUMERIC []string }
	StudentContactPersonal struct { NUMERIC []string }
	NAPTestScoreSummary struct { NUMERIC []string }
	TimeTableContainer struct { NUMERIC []string }
	VendorInfo struct { NUMERIC []string }
	AggregateCharacteristicInfo struct { NUMERIC []string }
	LearningResource struct { NUMERIC []string }
	SectionInfo struct { NUMERIC []string }
	AggregateStatisticInfo struct { NUMERIC []string }
	StaffPersonal struct { NUMERIC []string }
	StudentAttendanceTimeList struct { NUMERIC []string }
	CollectionStatus struct { NUMERIC []string }
	FinancialAccount struct { NUMERIC []string }
	NAPStudentResponseSet struct { NUMERIC []string }
	SchoolInfo struct { NUMERIC []string }
	TimeTableCell struct { NUMERIC []string }
	Activity struct { NUMERIC []string }
	CalendarDate struct { NUMERIC []string }
	CalendarSummary struct { NUMERIC []string }
	
}

// Bool2JSON :
type Bool2JSON struct {
	Path          string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	StudentContactPersonal struct { BOOLEAN []string }
	StudentParticipation struct { BOOLEAN []string }
	StudentSectionEnrollment struct { BOOLEAN []string }
	WellbeingAlert struct { BOOLEAN []string }
	AddressCollection struct { BOOLEAN []string }
	FinancialAccount struct { BOOLEAN []string }
	NAPEventStudentLink struct { BOOLEAN []string }
	NAPTestlet struct { BOOLEAN []string }
	WellbeingAppeal struct { BOOLEAN []string }
	CalendarDate struct { BOOLEAN []string }
	FinancialQuestionnaireCollection struct { BOOLEAN []string }
	StaffPersonal struct { BOOLEAN []string }
	StudentDailyAttendance struct { BOOLEAN []string }
	SystemRole struct { BOOLEAN []string }
	WellbeingPersonLink struct { BOOLEAN []string }
	LEAInfo struct { BOOLEAN []string }
	ResourceBooking struct { BOOLEAN []string }
	RoomInfo struct { BOOLEAN []string }
	SchoolPrograms struct { BOOLEAN []string }
	StaffAssignment struct { BOOLEAN []string }
	StudentAttendanceTimeList struct { BOOLEAN []string }
	StudentScoreJudgementAgainstStandard struct { BOOLEAN []string }
	TimeTable struct { BOOLEAN []string }
	NAPTestItem struct { BOOLEAN []string }
	PersonalisedPlan struct { BOOLEAN []string }
	TimeTableCell struct { BOOLEAN []string }
	TimeTableSubject struct { BOOLEAN []string }
	SectionInfo struct { BOOLEAN []string }
	TimeTableContainer struct { BOOLEAN []string }
	WellbeingCharacteristic struct { BOOLEAN []string }
	WellbeingEvent struct { BOOLEAN []string }
	AggregateStatisticFact struct { BOOLEAN []string }
	PaymentReceipt struct { BOOLEAN []string }
	StudentPeriodAttendance struct { BOOLEAN []string }
	WellbeingResponse struct { BOOLEAN []string }
	AggregateStatisticInfo struct { BOOLEAN []string }
	NAPTestScoreSummary struct { BOOLEAN []string }
	ResourceUsage struct { BOOLEAN []string }
	SchoolCourseInfo struct { BOOLEAN []string }
	StudentSchoolEnrollment struct { BOOLEAN []string }
	TermInfo struct { BOOLEAN []string }
	Identity struct { BOOLEAN []string }
	Invoice struct { BOOLEAN []string }
	SessionInfo struct { BOOLEAN []string }
	StudentActivityParticipation struct { BOOLEAN []string }
	StudentAttendanceSummary struct { BOOLEAN []string }
	Activity struct { BOOLEAN []string }
	LearningStandardDocument struct { BOOLEAN []string }
	NAPCodeFrame struct { BOOLEAN []string }
	StudentContactRelationship struct { BOOLEAN []string }
	CollectionRound struct { BOOLEAN []string }
	Debtor struct { BOOLEAN []string }
	StudentActivityInfo struct { BOOLEAN []string }
	StudentGrade struct { BOOLEAN []string }
	NAPTest struct { BOOLEAN []string }
	PurchaseOrder struct { BOOLEAN []string }
	NAPStudentResponseSet struct { BOOLEAN []string }
	EquipmentInfo struct { BOOLEAN []string }
	MarkValueInfo struct { BOOLEAN []string }
	CollectionStatus struct { BOOLEAN []string }
	Journal struct { BOOLEAN []string }
	StudentPersonal struct { BOOLEAN []string }
	AggregateCharacteristicInfo struct { BOOLEAN []string }
	CalendarSummary struct { BOOLEAN []string }
	ScheduledActivity struct { BOOLEAN []string }
	ChargedLocationInfo struct { BOOLEAN []string }
	GradingAssignment struct { BOOLEAN []string }
	TeachingGroup struct { BOOLEAN []string }
	GradingAssignmentScore struct { BOOLEAN []string }
	LearningStandardItem struct { BOOLEAN []string }
	PersonPicture struct { BOOLEAN []string }
	SchoolInfo struct { BOOLEAN []string }
	CensusCollection struct { BOOLEAN []string }
	LearningResource struct { BOOLEAN []string }
	VendorInfo struct { BOOLEAN []string }
	
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
