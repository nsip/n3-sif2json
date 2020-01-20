package cvt2json

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/burntsushi/toml"
)

// !!! toml file name must be identical to config struct definition name !!!

type list2json struct {
	Path          string
	PathAbs       string
	JQDir         string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	FinancialAccount                     struct{ LIST []string }
	Invoice                              struct{ LIST []string }
	NAPTestScoreSummary                  struct{ LIST []string }
	ResourceBooking                      struct{ LIST []string }
	SessionInfo                          struct{ LIST []string }
	StudentDailyAttendance               struct{ LIST []string }
	StudentPersonal                      struct{ LIST []string }
	LEAInfo                              struct{ LIST []string }
	LearningResource                     struct{ LIST []string }
	NAPTestlet                           struct{ LIST []string }
	StudentPeriodAttendance              struct{ LIST []string }
	WellbeingEvent                       struct{ LIST []string }
	MarkValueInfo                        struct{ LIST []string }
	NAPStudentResponseSet                struct{ LIST []string }
	PersonalisedPlan                     struct{ LIST []string }
	SchoolPrograms                       struct{ LIST []string }
	StudentActivityParticipation         struct{ LIST []string }
	TimeTable                            struct{ LIST []string }
	AGAddressCollectionSubmission        struct{ LIST []string }
	Journal                              struct{ LIST []string }
	PurchaseOrder                        struct{ LIST []string }
	LearningStandardDocument             struct{ LIST []string }
	ScheduledActivity                    struct{ LIST []string }
	SchoolInfo                           struct{ LIST []string }
	StudentAttendanceSummary             struct{ LIST []string }
	StudentSchoolEnrollment              struct{ LIST []string }
	VendorInfo                           struct{ LIST []string }
	WellbeingAlert                       struct{ LIST []string }
	WellbeingResponse                    struct{ LIST []string }
	PaymentReceipt                       struct{ LIST []string }
	PersonPrivacyObligation              struct{ LIST []string }
	StaffAssignment                      struct{ LIST []string }
	TimeTableContainer                   struct{ LIST []string }
	TimeTableSubject                     struct{ LIST []string }
	NAPTest                              struct{ LIST []string }
	StaffPersonal                        struct{ LIST []string }
	StudentContactRelationship           struct{ LIST []string }
	WellbeingAppeal                      struct{ LIST []string }
	AGStatusReport                       struct{ LIST []string }
	Activity                             struct{ LIST []string }
	StudentSectionEnrollment             struct{ LIST []string }
	SystemRole                           struct{ LIST []string }
	NAPEventStudentLink                  struct{ LIST []string }
	StudentActivityInfo                  struct{ LIST []string }
	StudentParticipation                 struct{ LIST []string }
	TermInfo                             struct{ LIST []string }
	AggregateStatisticFact               struct{ LIST []string }
	LearningStandardItem                 struct{ LIST []string }
	StudentAttendanceTimeList            struct{ LIST []string }
	StudentScoreJudgementAgainstStandard struct{ LIST []string }
	WellbeingPersonLink                  struct{ LIST []string }
	AGGetRounds                          struct{ LIST []string }
	CalendarSummary                      struct{ LIST []string }
	GradingAssignmentScore               struct{ LIST []string }
	PersonPicture                        struct{ LIST []string }
	FinancialQuestionnaireSubmission     struct{ LIST []string }
	NAPCodeFrame                         struct{ LIST []string }
	StudentContactPersonal               struct{ LIST []string }
	AggregateCharacteristicInfo          struct{ LIST []string }
	Debtor                               struct{ LIST []string }
	SchoolCourseInfo                     struct{ LIST []string }
	WellbeingCharacteristic              struct{ LIST []string }
	GradingAssignment                    struct{ LIST []string }
	ResourceUsage                        struct{ LIST []string }
	TimeTableCell                        struct{ LIST []string }
	AggregateStatisticInfo               struct{ LIST []string }
	CalendarDate                         struct{ LIST []string }
	ChargedLocationInfo                  struct{ LIST []string }
	Identity                             struct{ LIST []string }
	RoomInfo                             struct{ LIST []string }
	StudentGrade                         struct{ LIST []string }
	EquipmentInfo                        struct{ LIST []string }
	NAPTestItem                          struct{ LIST []string }
	ScheduledActivityContainer           struct{ LIST []string }
	SectionInfo                          struct{ LIST []string }
	TeachingGroup                        struct{ LIST []string }
}

type num2json struct {
	Path          string
	PathAbs       string
	JQDir         string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	FinancialAccount                     struct{ NUMERIC []string }
	Journal                              struct{ NUMERIC []string }
	SchoolCourseInfo                     struct{ NUMERIC []string }
	StudentActivityParticipation         struct{ NUMERIC []string }
	TimeTable                            struct{ NUMERIC []string }
	EquipmentInfo                        struct{ NUMERIC []string }
	Identity                             struct{ NUMERIC []string }
	NAPStudentResponseSet                struct{ NUMERIC []string }
	StudentContactPersonal               struct{ NUMERIC []string }
	WellbeingEvent                       struct{ NUMERIC []string }
	AGAddressCollectionSubmission        struct{ NUMERIC []string }
	StudentAttendanceTimeList            struct{ NUMERIC []string }
	StudentGrade                         struct{ NUMERIC []string }
	WellbeingAppeal                      struct{ NUMERIC []string }
	WellbeingPersonLink                  struct{ NUMERIC []string }
	SystemRole                           struct{ NUMERIC []string }
	AGStatusReport                       struct{ NUMERIC []string }
	NAPEventStudentLink                  struct{ NUMERIC []string }
	ScheduledActivity                    struct{ NUMERIC []string }
	SchoolPrograms                       struct{ NUMERIC []string }
	SessionInfo                          struct{ NUMERIC []string }
	StudentParticipation                 struct{ NUMERIC []string }
	StudentPeriodAttendance              struct{ NUMERIC []string }
	AGGetRounds                          struct{ NUMERIC []string }
	AggregateStatisticFact               struct{ NUMERIC []string }
	NAPCodeFrame                         struct{ NUMERIC []string }
	StudentActivityInfo                  struct{ NUMERIC []string }
	CalendarDate                         struct{ NUMERIC []string }
	LearningStandardItem                 struct{ NUMERIC []string }
	ResourceUsage                        struct{ NUMERIC []string }
	TimeTableCell                        struct{ NUMERIC []string }
	Activity                             struct{ NUMERIC []string }
	AggregateStatisticInfo               struct{ NUMERIC []string }
	GradingAssignmentScore               struct{ NUMERIC []string }
	NAPTestItem                          struct{ NUMERIC []string }
	ResourceBooking                      struct{ NUMERIC []string }
	StudentPersonal                      struct{ NUMERIC []string }
	TeachingGroup                        struct{ NUMERIC []string }
	AggregateCharacteristicInfo          struct{ NUMERIC []string }
	PersonPicture                        struct{ NUMERIC []string }
	PersonalisedPlan                     struct{ NUMERIC []string }
	SchoolInfo                           struct{ NUMERIC []string }
	StudentScoreJudgementAgainstStandard struct{ NUMERIC []string }
	Invoice                              struct{ NUMERIC []string }
	StudentSectionEnrollment             struct{ NUMERIC []string }
	ChargedLocationInfo                  struct{ NUMERIC []string }
	GradingAssignment                    struct{ NUMERIC []string }
	PaymentReceipt                       struct{ NUMERIC []string }
	PurchaseOrder                        struct{ NUMERIC []string }
	WellbeingAlert                       struct{ NUMERIC []string }
	FinancialQuestionnaireSubmission     struct{ NUMERIC []string }
	LearningResource                     struct{ NUMERIC []string }
	LearningStandardDocument             struct{ NUMERIC []string }
	NAPTestlet                           struct{ NUMERIC []string }
	CalendarSummary                      struct{ NUMERIC []string }
	StudentContactRelationship           struct{ NUMERIC []string }
	TimeTableContainer                   struct{ NUMERIC []string }
	WellbeingCharacteristic              struct{ NUMERIC []string }
	ScheduledActivityContainer           struct{ NUMERIC []string }
	SectionInfo                          struct{ NUMERIC []string }
	StaffAssignment                      struct{ NUMERIC []string }
	StudentSchoolEnrollment              struct{ NUMERIC []string }
	TimeTableSubject                     struct{ NUMERIC []string }
	VendorInfo                           struct{ NUMERIC []string }
	Debtor                               struct{ NUMERIC []string }
	LEAInfo                              struct{ NUMERIC []string }
	RoomInfo                             struct{ NUMERIC []string }
	StudentAttendanceSummary             struct{ NUMERIC []string }
	StudentDailyAttendance               struct{ NUMERIC []string }
	TermInfo                             struct{ NUMERIC []string }
	WellbeingResponse                    struct{ NUMERIC []string }
	MarkValueInfo                        struct{ NUMERIC []string }
	NAPTest                              struct{ NUMERIC []string }
	NAPTestScoreSummary                  struct{ NUMERIC []string }
	StaffPersonal                        struct{ NUMERIC []string }
}

type bool2json struct {
	Path          string
	PathAbs       string
	JQDir         string
	Sep           string
	CfgJSONOutDir string
	CfgJSONValue  string
	// ----------------------------------------------------------- //
	CalendarDate                         struct{ BOOLEAN []string }
	Debtor                               struct{ BOOLEAN []string }
	PersonPrivacyObligation              struct{ BOOLEAN []string }
	StudentGrade                         struct{ BOOLEAN []string }
	StudentParticipation                 struct{ BOOLEAN []string }
	StudentScoreJudgementAgainstStandard struct{ BOOLEAN []string }
	Activity                             struct{ BOOLEAN []string }
	AggregateCharacteristicInfo          struct{ BOOLEAN []string }
	NAPStudentResponseSet                struct{ BOOLEAN []string }
	NAPTestItem                          struct{ BOOLEAN []string }
	StudentPeriodAttendance              struct{ BOOLEAN []string }
	LearningStandardDocument             struct{ BOOLEAN []string }
	LearningStandardItem                 struct{ BOOLEAN []string }
	SchoolInfo                           struct{ BOOLEAN []string }
	StudentActivityInfo                  struct{ BOOLEAN []string }
	StudentContactPersonal               struct{ BOOLEAN []string }
	TimeTable                            struct{ BOOLEAN []string }
	NAPCodeFrame                         struct{ BOOLEAN []string }
	NAPEventStudentLink                  struct{ BOOLEAN []string }
	TeachingGroup                        struct{ BOOLEAN []string }
	TermInfo                             struct{ BOOLEAN []string }
	WellbeingPersonLink                  struct{ BOOLEAN []string }
	AggregateStatisticInfo               struct{ BOOLEAN []string }
	GradingAssignment                    struct{ BOOLEAN []string }
	NAPTest                              struct{ BOOLEAN []string }
	StaffAssignment                      struct{ BOOLEAN []string }
	StudentContactRelationship           struct{ BOOLEAN []string }
	FinancialQuestionnaireSubmission     struct{ BOOLEAN []string }
	LearningResource                     struct{ BOOLEAN []string }
	MarkValueInfo                        struct{ BOOLEAN []string }
	StudentSectionEnrollment             struct{ BOOLEAN []string }
	WellbeingAlert                       struct{ BOOLEAN []string }
	AGGetRounds                          struct{ BOOLEAN []string }
	ChargedLocationInfo                  struct{ BOOLEAN []string }
	EquipmentInfo                        struct{ BOOLEAN []string }
	NAPTestScoreSummary                  struct{ BOOLEAN []string }
	NAPTestlet                           struct{ BOOLEAN []string }
	ScheduledActivityContainer           struct{ BOOLEAN []string }
	SessionInfo                          struct{ BOOLEAN []string }
	Journal                              struct{ BOOLEAN []string }
	RoomInfo                             struct{ BOOLEAN []string }
	AGAddressCollectionSubmission        struct{ BOOLEAN []string }
	AggregateStatisticFact               struct{ BOOLEAN []string }
	FinancialAccount                     struct{ BOOLEAN []string }
	Invoice                              struct{ BOOLEAN []string }
	SchoolCourseInfo                     struct{ BOOLEAN []string }
	AGStatusReport                       struct{ BOOLEAN []string }
	ResourceBooking                      struct{ BOOLEAN []string }
	ScheduledActivity                    struct{ BOOLEAN []string }
	StudentAttendanceSummary             struct{ BOOLEAN []string }
	CalendarSummary                      struct{ BOOLEAN []string }
	GradingAssignmentScore               struct{ BOOLEAN []string }
	PersonPicture                        struct{ BOOLEAN []string }
	StudentPersonal                      struct{ BOOLEAN []string }
	TimeTableSubject                     struct{ BOOLEAN []string }
	VendorInfo                           struct{ BOOLEAN []string }
	LEAInfo                              struct{ BOOLEAN []string }
	StudentAttendanceTimeList            struct{ BOOLEAN []string }
	PersonalisedPlan                     struct{ BOOLEAN []string }
	StudentSchoolEnrollment              struct{ BOOLEAN []string }
	Identity                             struct{ BOOLEAN []string }
	PaymentReceipt                       struct{ BOOLEAN []string }
	PurchaseOrder                        struct{ BOOLEAN []string }
	SchoolPrograms                       struct{ BOOLEAN []string }
	SectionInfo                          struct{ BOOLEAN []string }
	TimeTableCell                        struct{ BOOLEAN []string }
	TimeTableContainer                   struct{ BOOLEAN []string }
	WellbeingEvent                       struct{ BOOLEAN []string }
	StaffPersonal                        struct{ BOOLEAN []string }
	StudentActivityParticipation         struct{ BOOLEAN []string }
	SystemRole                           struct{ BOOLEAN []string }
	ResourceUsage                        struct{ BOOLEAN []string }
	StudentDailyAttendance               struct{ BOOLEAN []string }
	WellbeingAppeal                      struct{ BOOLEAN []string }
	WellbeingCharacteristic              struct{ BOOLEAN []string }
	WellbeingResponse                    struct{ BOOLEAN []string }
}

type sif2json struct {
	Path            string
	PathAbs         string
	JQDir           string
	AttrPrefix      string
	ContentPrefix   string
	CfgJSONDir4LIST string
	CfgJSONDir4NUM  string
	CfgJSONDir4BOOL string
}

var (
	// toml file name must be identical to config struct definition name
	lsCfg = []interface{}{
		&list2json{},
		&num2json{},
		&bool2json{},
		&sif2json{},
	}
)

// ------------------------------------------------- //

// NewCfg :
func NewCfg(cfgPaths ...string) interface{} {
	for _, f := range cfgPaths {
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
				reflect.ValueOf(cfg).Elem().FieldByName("Path").SetString(f)
				reflect.ValueOf(cfg).Elem().FieldByName("PathAbs").SetString(abs)
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
