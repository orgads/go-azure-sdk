package share

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetType string

const (
	DataSetTypeAdlsGenOneFile               DataSetType = "AdlsGen1File"
	DataSetTypeAdlsGenOneFolder             DataSetType = "AdlsGen1Folder"
	DataSetTypeAdlsGenTwoFile               DataSetType = "AdlsGen2File"
	DataSetTypeAdlsGenTwoFileSystem         DataSetType = "AdlsGen2FileSystem"
	DataSetTypeAdlsGenTwoFolder             DataSetType = "AdlsGen2Folder"
	DataSetTypeBlob                         DataSetType = "Blob"
	DataSetTypeBlobFolder                   DataSetType = "BlobFolder"
	DataSetTypeContainer                    DataSetType = "Container"
	DataSetTypeKustoCluster                 DataSetType = "KustoCluster"
	DataSetTypeKustoDatabase                DataSetType = "KustoDatabase"
	DataSetTypeKustoTable                   DataSetType = "KustoTable"
	DataSetTypeSqlDBTable                   DataSetType = "SqlDBTable"
	DataSetTypeSqlDWTable                   DataSetType = "SqlDWTable"
	DataSetTypeSynapseWorkspaceSqlPoolTable DataSetType = "SynapseWorkspaceSqlPoolTable"
)

func PossibleValuesForDataSetType() []string {
	return []string{
		string(DataSetTypeAdlsGenOneFile),
		string(DataSetTypeAdlsGenOneFolder),
		string(DataSetTypeAdlsGenTwoFile),
		string(DataSetTypeAdlsGenTwoFileSystem),
		string(DataSetTypeAdlsGenTwoFolder),
		string(DataSetTypeBlob),
		string(DataSetTypeBlobFolder),
		string(DataSetTypeContainer),
		string(DataSetTypeKustoCluster),
		string(DataSetTypeKustoDatabase),
		string(DataSetTypeKustoTable),
		string(DataSetTypeSqlDBTable),
		string(DataSetTypeSqlDWTable),
		string(DataSetTypeSynapseWorkspaceSqlPoolTable),
	}
}

func parseDataSetType(input string) (*DataSetType, error) {
	vals := map[string]DataSetType{
		"adlsgen1file":                 DataSetTypeAdlsGenOneFile,
		"adlsgen1folder":               DataSetTypeAdlsGenOneFolder,
		"adlsgen2file":                 DataSetTypeAdlsGenTwoFile,
		"adlsgen2filesystem":           DataSetTypeAdlsGenTwoFileSystem,
		"adlsgen2folder":               DataSetTypeAdlsGenTwoFolder,
		"blob":                         DataSetTypeBlob,
		"blobfolder":                   DataSetTypeBlobFolder,
		"container":                    DataSetTypeContainer,
		"kustocluster":                 DataSetTypeKustoCluster,
		"kustodatabase":                DataSetTypeKustoDatabase,
		"kustotable":                   DataSetTypeKustoTable,
		"sqldbtable":                   DataSetTypeSqlDBTable,
		"sqldwtable":                   DataSetTypeSqlDWTable,
		"synapseworkspacesqlpooltable": DataSetTypeSynapseWorkspaceSqlPoolTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSetType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"moving":    ProvisioningStateMoving,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ShareKind string

const (
	ShareKindCopyBased ShareKind = "CopyBased"
	ShareKindInPlace   ShareKind = "InPlace"
)

func PossibleValuesForShareKind() []string {
	return []string{
		string(ShareKindCopyBased),
		string(ShareKindInPlace),
	}
}

func parseShareKind(input string) (*ShareKind, error) {
	vals := map[string]ShareKind{
		"copybased": ShareKindCopyBased,
		"inplace":   ShareKindInPlace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareKind(input)
	return &out, nil
}

type ShareSubscriptionStatus string

const (
	ShareSubscriptionStatusActive        ShareSubscriptionStatus = "Active"
	ShareSubscriptionStatusRevoked       ShareSubscriptionStatus = "Revoked"
	ShareSubscriptionStatusRevoking      ShareSubscriptionStatus = "Revoking"
	ShareSubscriptionStatusSourceDeleted ShareSubscriptionStatus = "SourceDeleted"
)

func PossibleValuesForShareSubscriptionStatus() []string {
	return []string{
		string(ShareSubscriptionStatusActive),
		string(ShareSubscriptionStatusRevoked),
		string(ShareSubscriptionStatusRevoking),
		string(ShareSubscriptionStatusSourceDeleted),
	}
}

func parseShareSubscriptionStatus(input string) (*ShareSubscriptionStatus, error) {
	vals := map[string]ShareSubscriptionStatus{
		"active":        ShareSubscriptionStatusActive,
		"revoked":       ShareSubscriptionStatusRevoked,
		"revoking":      ShareSubscriptionStatusRevoking,
		"sourcedeleted": ShareSubscriptionStatusSourceDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareSubscriptionStatus(input)
	return &out, nil
}

type Status string

const (
	StatusAccepted         Status = "Accepted"
	StatusCanceled         Status = "Canceled"
	StatusFailed           Status = "Failed"
	StatusInProgress       Status = "InProgress"
	StatusSucceeded        Status = "Succeeded"
	StatusTransientFailure Status = "TransientFailure"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAccepted),
		string(StatusCanceled),
		string(StatusFailed),
		string(StatusInProgress),
		string(StatusSucceeded),
		string(StatusTransientFailure),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"accepted":         StatusAccepted,
		"canceled":         StatusCanceled,
		"failed":           StatusFailed,
		"inprogress":       StatusInProgress,
		"succeeded":        StatusSucceeded,
		"transientfailure": StatusTransientFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type SynchronizationMode string

const (
	SynchronizationModeFullSync    SynchronizationMode = "FullSync"
	SynchronizationModeIncremental SynchronizationMode = "Incremental"
)

func PossibleValuesForSynchronizationMode() []string {
	return []string{
		string(SynchronizationModeFullSync),
		string(SynchronizationModeIncremental),
	}
}

func parseSynchronizationMode(input string) (*SynchronizationMode, error) {
	vals := map[string]SynchronizationMode{
		"fullsync":    SynchronizationModeFullSync,
		"incremental": SynchronizationModeIncremental,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMode(input)
	return &out, nil
}
