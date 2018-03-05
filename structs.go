package main

//Record type based on data_clinical_patient
type Record struct {
	PatientID                                string `json:"PATIENT_ID"`
	Age                                      string `json:"AGE"`
	SmokingYearStarted                       string `json:"SMOKING_YEAR_STARTED"`
	ClinicalAssessment                       string `json:"CLINICAL_ASSESSMENT"`
	YearOfDiagnosis                          string `json:"YEAR_OF_DIAGNOSIS"`
	Subtype                                  string `json:"SUBTYPE"`
	Ethnicity                                string `json:"ETHNICITY"`
	Sex                                      string `json:"SEX"`
	LostToFollowup                           string `json:"LOST_TO_FOLLOWUP"`
	LymphNodesExamined                       string `json:"LYMPH_NODES_EXAMINED"`
	AngiolymphaticInvasion                   string `json:"ANGIOLYMPHATIC_INVASION"`
	MethodOfInitialPathDiagnosis             string `json:"METHOD_OF_INITIAL_PATH_DIAGNOSIS"`
	NewTumorEventAfterInitialTreatment       string `json:"NEW_TUMOR_EVENT_AFTER_INITIAL_TREATMENT"`
	SmokingPackYears                         string `json:"SMOKING_PACK_YEARS"`
	MethodOfInitialPathDiagnosisOther        string `json:"METHOD_OF_INITIAL_PATH_DIAGNOSIS_OTHER"`
	NewTumorEventOtherSite                   string `json:"NEW_TUMOR_EVENT_OTHER_SITE"`
	PathologicDistantSpread                  string `json:"PATHOLOGIC_DISTANT_SPREAD"`
	PrimaryTumorPathologicSpread             string `json:"PRIMARY_TUMOR_PATHOLOGIC_SPREAD"`
	ProspectiveCollection                    string `json:"PROSPECTIVE_COLLECTION"`
	Race                                     string `json:"RACE"`
	SmokingHistory                           string `json:"SMOKING_HISTORY"`
	SmokingYearStopped                       string `json:"SMOKING_YEAR_STOPPED"`
	OsStatus                                 string `json:"OS_STATUS"`
	OsMonths                                 string `json:"OS_MONTHS"`
	DfsStatus                                string `json:"DFS_STATUS"`
	DfsMonths                                string `json:"DFS_MONTHS"`
	PathologicSpreadIncidentalProstateCancer string `json:"PATHOLOGIC_SPREAD_INCIDENTAL_PROSTATE_CANCER"`
}
