package cvt2json

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/burntsushi/toml"
)

// !!! toml file name must be identical to config struct definition name !!!

type cfg2json struct {
	Path          string
	PathAbs       string
	JQDir         string
	Sep           string
	CfgJSONOutDir string
	// ----------------------------------------------------------- //
	AggregateStatisticInfo               struct{ ListAttrs []string }
	SchoolPrograms                       struct{ ListAttrs []string }
	SectionInfo                          struct{ ListAttrs []string }
	TimeTableCell                        struct{ ListAttrs []string }
	WellbeingAppeal                      struct{ ListAttrs []string }
	AGStatusReport                       struct{ ListAttrs []string }
	CalendarSummary                      struct{ ListAttrs []string }
	FinancialAccount                     struct{ ListAttrs []string }
	FinancialQuestionnaireSubmission     struct{ ListAttrs []string }
	StudentActivityParticipation         struct{ ListAttrs []string }
	GradingAssignmentScore               struct{ ListAttrs []string }
	StudentContactPersonal               struct{ ListAttrs []string }
	VendorInfo                           struct{ ListAttrs []string }
	LEAInfo                              struct{ ListAttrs []string }
	MarkValueInfo                        struct{ ListAttrs []string }
	NAPEventStudentLink                  struct{ ListAttrs []string }
	NAPTest                              struct{ ListAttrs []string }
	StaffPersonal                        struct{ ListAttrs []string }
	TermInfo                             struct{ ListAttrs []string }
	WellbeingEvent                       struct{ ListAttrs []string }
	Identity                             struct{ ListAttrs []string }
	Invoice                              struct{ ListAttrs []string }
	LearningStandardDocument             struct{ ListAttrs []string }
	WellbeingPersonLink                  struct{ ListAttrs []string }
	LearningStandardItem                 struct{ ListAttrs []string }
	StudentGrade                         struct{ ListAttrs []string }
	StudentSectionEnrollment             struct{ ListAttrs []string }
	StudentPersonal                      struct{ ListAttrs []string }
	AGAddressCollectionSubmission        struct{ ListAttrs []string }
	AGGetRounds                          struct{ ListAttrs []string }
	LearningResource                     struct{ ListAttrs []string }
	NAPTestScoreSummary                  struct{ ListAttrs []string }
	ResourceUsage                        struct{ ListAttrs []string }
	SessionInfo                          struct{ ListAttrs []string }
	StudentAttendanceTimeList            struct{ ListAttrs []string }
	NAPCodeFrame                         struct{ ListAttrs []string }
	StudentScoreJudgementAgainstStandard struct{ ListAttrs []string }
	StudentAttendanceSummary             struct{ ListAttrs []string }
	GradingAssignment                    struct{ ListAttrs []string }
	PersonPrivacyObligation              struct{ ListAttrs []string }
	RoomInfo                             struct{ ListAttrs []string }
	ScheduledActivity                    struct{ ListAttrs []string }
	SchoolInfo                           struct{ ListAttrs []string }
	StaffAssignment                      struct{ ListAttrs []string }
	StudentActivityInfo                  struct{ ListAttrs []string }
	SystemRole                           struct{ ListAttrs []string }
	TimeTableSubject                     struct{ ListAttrs []string }
	WellbeingAlert                       struct{ ListAttrs []string }
	AggregateStatisticFact               struct{ ListAttrs []string }
	Journal                              struct{ ListAttrs []string }
	StudentDailyAttendance               struct{ ListAttrs []string }
	EquipmentInfo                        struct{ ListAttrs []string }
	AggregateCharacteristicInfo          struct{ ListAttrs []string }
	NAPTestItem                          struct{ ListAttrs []string }
	Activity                             struct{ ListAttrs []string }
	PersonalisedPlan                     struct{ ListAttrs []string }
	StudentPeriodAttendance              struct{ ListAttrs []string }
	TimeTableContainer                   struct{ ListAttrs []string }
	WellbeingCharacteristic              struct{ ListAttrs []string }
	PersonPicture                        struct{ ListAttrs []string }
	PurchaseOrder                        struct{ ListAttrs []string }
	StudentContactRelationship           struct{ ListAttrs []string }
	WellbeingResponse                    struct{ ListAttrs []string }
	NAPTestlet                           struct{ ListAttrs []string }
	PaymentReceipt                       struct{ ListAttrs []string }
	ResourceBooking                      struct{ ListAttrs []string }
	SchoolCourseInfo                     struct{ ListAttrs []string }
	StudentSchoolEnrollment              struct{ ListAttrs []string }
	TeachingGroup                        struct{ ListAttrs []string }
	TimeTable                            struct{ ListAttrs []string }
	CalendarDate                         struct{ ListAttrs []string }
	ChargedLocationInfo                  struct{ ListAttrs []string }
	Debtor                               struct{ ListAttrs []string }
	NAPStudentResponseSet                struct{ ListAttrs []string }
	ScheduledActivityContainer           struct{ ListAttrs []string }
	StudentParticipation                 struct{ ListAttrs []string }
}

type sif2json struct {
	Path          string
	PathAbs       string
	JQDir         string
	AttrPrefix    string
	ContentPrefix string
	CfgJSONDir    string
}

var (
	// toml file name must be identical to config struct definition name
	lsCfg = []interface{}{
		&cfg2json{},
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
