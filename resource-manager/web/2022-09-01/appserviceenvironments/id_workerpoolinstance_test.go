package appserviceenvironments

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkerPoolInstanceId{}

func TestNewWorkerPoolInstanceID(t *testing.T) {
	id := NewWorkerPoolInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentValue", "workerPoolValue", "instanceValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.ResourceGroupName != "example-resource-group" {
		t.Fatalf("Expected %q but got %q for Segment 'ResourceGroupName'", id.ResourceGroupName, "example-resource-group")
	}

	if id.HostingEnvironmentName != "hostingEnvironmentValue" {
		t.Fatalf("Expected %q but got %q for Segment 'HostingEnvironmentName'", id.HostingEnvironmentName, "hostingEnvironmentValue")
	}

	if id.WorkerPoolName != "workerPoolValue" {
		t.Fatalf("Expected %q but got %q for Segment 'WorkerPoolName'", id.WorkerPoolName, "workerPoolValue")
	}

	if id.InstanceName != "instanceValue" {
		t.Fatalf("Expected %q but got %q for Segment 'InstanceName'", id.InstanceName, "instanceValue")
	}
}

func TestFormatWorkerPoolInstanceID(t *testing.T) {
	actual := NewWorkerPoolInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentValue", "workerPoolValue", "instanceValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances/instanceValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseWorkerPoolInstanceID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *WorkerPoolInstanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances/instanceValue",
			Expected: &WorkerPoolInstanceId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				HostingEnvironmentName: "hostingEnvironmentValue",
				WorkerPoolName:         "workerPoolValue",
				InstanceName:           "instanceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances/instanceValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseWorkerPoolInstanceID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.HostingEnvironmentName != v.Expected.HostingEnvironmentName {
			t.Fatalf("Expected %q but got %q for HostingEnvironmentName", v.Expected.HostingEnvironmentName, actual.HostingEnvironmentName)
		}

		if actual.WorkerPoolName != v.Expected.WorkerPoolName {
			t.Fatalf("Expected %q but got %q for WorkerPoolName", v.Expected.WorkerPoolName, actual.WorkerPoolName)
		}

		if actual.InstanceName != v.Expected.InstanceName {
			t.Fatalf("Expected %q but got %q for InstanceName", v.Expected.InstanceName, actual.InstanceName)
		}

	}
}

func TestParseWorkerPoolInstanceIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *WorkerPoolInstanceId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe/wOrKeRpOoLs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe/wOrKeRpOoLs/wOrKeRpOoLvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe/wOrKeRpOoLs/wOrKeRpOoLvAlUe/iNsTaNcEs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances/instanceValue",
			Expected: &WorkerPoolInstanceId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "example-resource-group",
				HostingEnvironmentName: "hostingEnvironmentValue",
				WorkerPoolName:         "workerPoolValue",
				InstanceName:           "instanceValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Web/hostingEnvironments/hostingEnvironmentValue/workerPools/workerPoolValue/instances/instanceValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe/wOrKeRpOoLs/wOrKeRpOoLvAlUe/iNsTaNcEs/iNsTaNcEvAlUe",
			Expected: &WorkerPoolInstanceId{
				SubscriptionId:         "12345678-1234-9876-4563-123456789012",
				ResourceGroupName:      "eXaMpLe-rEsOuRcE-GrOuP",
				HostingEnvironmentName: "hOsTiNgEnViRoNmEnTvAlUe",
				WorkerPoolName:         "wOrKeRpOoLvAlUe",
				InstanceName:           "iNsTaNcEvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/rEsOuRcEgRoUpS/eXaMpLe-rEsOuRcE-GrOuP/pRoViDeRs/mIcRoSoFt.wEb/hOsTiNgEnViRoNmEnTs/hOsTiNgEnViRoNmEnTvAlUe/wOrKeRpOoLs/wOrKeRpOoLvAlUe/iNsTaNcEs/iNsTaNcEvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseWorkerPoolInstanceIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.ResourceGroupName != v.Expected.ResourceGroupName {
			t.Fatalf("Expected %q but got %q for ResourceGroupName", v.Expected.ResourceGroupName, actual.ResourceGroupName)
		}

		if actual.HostingEnvironmentName != v.Expected.HostingEnvironmentName {
			t.Fatalf("Expected %q but got %q for HostingEnvironmentName", v.Expected.HostingEnvironmentName, actual.HostingEnvironmentName)
		}

		if actual.WorkerPoolName != v.Expected.WorkerPoolName {
			t.Fatalf("Expected %q but got %q for WorkerPoolName", v.Expected.WorkerPoolName, actual.WorkerPoolName)
		}

		if actual.InstanceName != v.Expected.InstanceName {
			t.Fatalf("Expected %q but got %q for InstanceName", v.Expected.InstanceName, actual.InstanceName)
		}

	}
}

func TestSegmentsForWorkerPoolInstanceId(t *testing.T) {
	segments := WorkerPoolInstanceId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("WorkerPoolInstanceId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
