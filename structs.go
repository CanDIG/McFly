package main

//PatientRecord type based on data_clinical_patient
type PatientRecord struct {
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

//SampleRecord type based on data_clinical_sample
type SampleRecord struct {
	PatientID                     string `json:"PATIENT_ID"`
	SampleID                      string `json:"SAMPLE_ID"`
	BatchNumber                   string `json:"BATCH_NUMBER"`
	PrimarySite                   string `json:"PRIMARY_SITE"`
	ExtracapsularSpread           string `json:"EXTRACAPSULAR_SPREAD"`
	IfExtracapsularSpreadPresent  string `json:"IF_EXTRACAPSULAR_SPREAD_PRESENT"`
	IncidentalProstateCancer      string `json:"INCIDENTAL_PROSTATE_CANCER"`
	MetastaticSite                string `json:"METASTATIC_SITE"`
	NodesEx                       string `json:"NODES_EX"`
	MetastaticSiteOther           string `json:"METASTATIC_SITE_OTHER"`
	RegionalNodesPathologicSpread string `json:"REGIONAL_NODES_PATHOLOGIC_SPREAD"`
	PriorTumor                    string `json:"PRIOR_TUMOR"`
	NewTumorEventSite             string `json:"NEW_TUMOR_EVENT_SITE"`
	Grade                         string `json:"GRADE"`
	TStage                        string `json:"T_STAGE"`
	NewTumorEventType             string `json:"NEW_TUMOR_EVENT_TYPE"`
	CompleteTCGAID                string `json:"COMPLETE_TCGA_ID"`
	CancerTypeDetailed            string `json:"CANCER_TYPE_DETAILED"`
	OncotreeCode                  string `json:"ONCOTREE_CODE"`
}

//MutationRecord type based on data_clinical_sample
type MutationRecord struct {
	HugoSymbol                 string `json:"Hugo_Symbol"`
	EntrezGeneID               string `json:"Entrez_Gene_Id"`
	Center                     string `json:"Center"`
	NCBIBuild                  string `json:"NCBI_Build"`
	Chromosome                 string `json:"Chromosome"`
	StartPosition              string `json:"Start_Position"`
	EndPosition                string `json:"End_Position"`
	Strand                     string `json:"Strand"`
	Consequence                string `json:"Consequence"`
	VariantClassification      string `json:"Variant_Classification"`
	VariantType                string `json:"Variant_Type"`
	ReferenceAllele            string `json:"Reference_Allele"`
	TumorSeqAllele1            string `json:"Tumor_Seq_Allele1"`
	TumorSeqAllele2            string `json:"Tumor_Seq_Allele2"`
	DBSNPRS                    string `json:"dbSNP_RS"`
	DBSNPValStatus             string `json:"dbSNP_Val_Status"`
	TumorSampleBarcode         string `json:"Tumor_Sample_Barcode"`
	MatchedNormSampleBarcode   string `json:"Matched_Norm_Sample_Barcode"`
	MatchNormSeqAllele1        string `json:"Match_Norm_Seq_Allele1"`
	MatchNormSeqAllele2        string `json:"Match_Norm_Seq_Allele2"`
	TumorValidationAllele1     string `json:"Tumor_Validation_Allele1"`
	TumorValidationAllele2     string `json:"Tumor_Validation_Allele2"`
	MatchNormValidationAllele1 string `json:"Match_Norm_Validation_Allele1"`
	MatchNormValidationAllele2 string `json:"Match_Norm_Validation_Allele2"`
	VerificationStatus         string `json:"Verification_Status"`
	ValidationStatus           string `json:"Validation_Status"`
	MutationStatus             string `json:"Mutation_Status"`
	SequencingPhase            string `json:"Sequencing_Phase"`
	SequenceSource             string `json:"Sequence_Source"`
	ValidationMethod           string `json:"Validation_Method"`
	Score                      string `json:"Score"`
	BAMFile                    string `json:"BAM_File"`
	Sequencer                  string `json:"Sequencer"`
	TRefCount                  string `json:"t_ref_count"`
	TAltCount                  string `json:"t_alt_count"`
	NRefCount                  string `json:"n_ref_count"`
	NAltCount                  string `json:"n_alt_count"`
	HGVSc                      string `json:"HGVSc"`
	HGVSp                      string `json:"HGVSp"`
	HGVSpShort                 string `json:"HGVSp_Short"`
}
