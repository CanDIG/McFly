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
	Gender                                   string `json:"GENDER"`
	Metastasis                               string `json:"METASTASIS"`
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
	ERStatus                      string `json:"ER_STATUS"`
	PRStatus                      string `json:"PR_STATUS"`
	HER2Status                    string `json:"HER2_STATUS"`
	TumorStage                    string `json:"TUMOR_STAGE"`
	TumorT1Coded                  string `json:"TUMOR_T1_CODED"`
	Nodes                         string `json:"NODES"`
	NodeCoded                     string `json:"NODE_CODED"`
	MetastasisCoded               string `json:"METASTASIS_CODED"`
	ConvertedStage                string `json:"CONVERTED_STAGE"`
	SurvivalDataForm              string `json:"SURVIVAL_DATA_FORM"`
	PAM50Subtype                  string `json:"PAM50_SUBTYPE"`
	SigclustUnsupervisedMRNA      string `json:"SIGCLUST_UNSUPERVISED_MRNA"`
	SigclustIntrinsicMRNA         string `json:"SIGCLUST_INTRINSIC_MRNA"`
	MirnaCluster                  string `json:"MIRNA_CLUSTER"`
	MethylationCluster            string `json:"METHYLATION_CLUSTER"`
	RPPACluster                   string `json:"RPPA_CLUSTER"`
	CNCluster                     string `json:"CN_CLUSTER"`
	IntegreatedClusterswithPAM50  string `json:"INTEGRATED_CLUSTERS_WITH_PAM50"`
	IntegreatedClustersNoExp      string `json:"INTEGRATED_CLUSTERS_NO_EXP"`
	IntegreatedClustersUnsupExp   string `json:"INTEGRATED_CLUSTERS_UNSUP_EXP"`
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
	// Transcript_ID
	// RefSeq
	// Protein_position
	// Codons
	// Hotspot
	// ALLELE_NUM
	// PICK
	// UNIPARC
	// n_depth
	// ONCOTATOR_TRANSCRIPT_CHANGE_BEST_EFFECT
	// MA:link.PDB
	// Feature
	// CLIN_SIG
	// ONCOTATOR_DBSNP_RS
	// Gene
	// HGNC_ID
	// ExAC_AF_AMR
	// t_depth
	// MA:FIS
	// ONCOTATOR_REFSEQ_MRNA_ID_BEST_EFFECT
	// ONCOTATOR_CODON_CHANGE_BEST_EFFECT
	// DISTANCE
	// SYMBOL_SOURCE
	// type_wu
	// type_WU
	// ONCOTATOR_PROTEIN_POS_END
	// Existing_variation
	// SYMBOL
	// ExAC_AF_SAS
	// VARIANT_CLASS
	// AA_MAF
	// HIGH_INF_POS
	// GENE_PHENO
	// ExAC_AF_AFR
	// ASN_MAF
	// all_domains_wu
	// all_domains_WU
	// PHENO
	// ONCOTATOR_VARIANT_CLASSIFICATION
	// BIOTYPE
	// transcript_name_wu
	// transcript_name_WU
	// MA:link.var
	// AFR_MAF
	// reference_wu
	// reference_WU
	// DOMAINS
	// MOTIF_SCORE_CHANGE
	// ONCOTATOR_UNIPROT_ENTRY_NAME
	// ONCOTATOR_TRANSCRIPT_CHANGE
	// Amino_acids
	// ONCOTATOR_REFSEQ_PROT_ID_BEST_EFFECT
	// EA_MAF
	// Allele
	// transcript_version_wu
	// transcript_version_WU
	// MA:FImpact
	// cDNA_position
	// chromosome_name_wu
	// chromosome_name_WU
	// ExAC_AF_NFE
	// ONCOTATOR_PROTEIN_POS_END_BEST_EFFECT
	// ONCOTATOR_PROTEIN_CHANGE_BEST_EFFECT
	// transcript_source_wu
	// transcript_source_WU
	// MA:link.MSA
	// SIFT
	// ONCOTATOR_GENE_SYMBOL_BEST_EFFECT
	// INTRON
	// ONCOTATOR_EXON_AFFECTED
	// TREMBL
	// transcript_status_wu
	// transcript_status_WU
	// AMR_MAF
	// stop_wu
	// stop_WU
	// ucsc_cons_wu
	// ucsc_cons_WU
	// transcript_species_wu
	// transcript_species_WU
	// ONCOTATOR_UNIPROT_ACCESSION_BEST_EFFECT
	// EAS_MAF
	// gene_name_wu
	// gene_name_WU
	// variant_wu
	// variant_WU
	// ONCOTATOR_PROTEIN_POS_START
	// CANONICAL
	// all_effects
	// ExAC_AF_EAS
	// start_wu
	// start_WU
	// GMAF
	// MOTIF_NAME
	// TSL
	// amino_acid_change_wu
	// amino_acid_change_WU
	// SOMATIC
	// MOTIF_POS
	// IMPACT
	// CDS_position
	// annotation_errors_wu
	// annotation_errors_WU
	// strand_wu
	// strand_WU
	// ONCOTATOR_DBSNP_VAL_STATUS
	// c_position_wu
	// c_position_WU
	// MA:protein.change
	// ONCOTATOR_COSMIC_OVERLAPPING
	// ONCOTATOR_REFSEQ_MRNA_ID
	// ONCOTATOR_CODON_CHANGE
	// ONCOTATOR_PROTEIN_CHANGE
	// ONCOTATOR_VARIANT_CLASSIFICATION_BEST_EFFECT
	// ONCOTATOR_EXON_AFFECTED_BEST_EFFECT
	// ONCOTATOR_GENE_SYMBOL
	// SWISSPROT
	// ExAC_AF_FIN
	// EUR_MAF
	// deletion_substructures_wu
	// deletion_substructures_WU
	// Feature_type
	// HGVS_OFFSET
	// domain_wu
	// domain_WU
	// PolyPhen
	// FILTER
	// ONCOTATOR_UNIPROT_ENTRY_NAME_BEST_EFFECT
	// ENSP
	// ExAC_AF
	// trv_type_wu
	// trv_type_WU
	// ONCOTATOR_REFSEQ_PROT_ID
	// CCDS
	// EXON
	// ExAC_AF_OTH
	// SAS_MAF
	// ONCOTATOR_PROTEIN_POS_START_BEST_EFFECT
	// Exon_Number
	// MINIMISED
	// ONCOTATOR_UNIPROT_ACCESSION
	// PUBMED
}
