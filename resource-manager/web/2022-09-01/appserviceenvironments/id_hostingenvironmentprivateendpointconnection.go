package appserviceenvironments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HostingEnvironmentPrivateEndpointConnectionId{}

// HostingEnvironmentPrivateEndpointConnectionId is a struct representing the Resource ID for a Hosting Environment Private Endpoint Connection
type HostingEnvironmentPrivateEndpointConnectionId struct {
	SubscriptionId                string
	ResourceGroupName             string
	HostingEnvironmentName        string
	PrivateEndpointConnectionName string
}

// NewHostingEnvironmentPrivateEndpointConnectionID returns a new HostingEnvironmentPrivateEndpointConnectionId struct
func NewHostingEnvironmentPrivateEndpointConnectionID(subscriptionId string, resourceGroupName string, hostingEnvironmentName string, privateEndpointConnectionName string) HostingEnvironmentPrivateEndpointConnectionId {
	return HostingEnvironmentPrivateEndpointConnectionId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		HostingEnvironmentName:        hostingEnvironmentName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
	}
}

// ParseHostingEnvironmentPrivateEndpointConnectionID parses 'input' into a HostingEnvironmentPrivateEndpointConnectionId
func ParseHostingEnvironmentPrivateEndpointConnectionID(input string) (*HostingEnvironmentPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentPrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ParseHostingEnvironmentPrivateEndpointConnectionIDInsensitively parses 'input' case-insensitively into a HostingEnvironmentPrivateEndpointConnectionId
// note: this method should only be used for API response data and not user input
func ParseHostingEnvironmentPrivateEndpointConnectionIDInsensitively(input string) (*HostingEnvironmentPrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentPrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentPrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ValidateHostingEnvironmentPrivateEndpointConnectionID checks that 'input' can be parsed as a Hosting Environment Private Endpoint Connection ID
func ValidateHostingEnvironmentPrivateEndpointConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostingEnvironmentPrivateEndpointConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hosting Environment Private Endpoint Connection ID
func (id HostingEnvironmentPrivateEndpointConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/hostingEnvironments/%s/privateEndpointConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HostingEnvironmentName, id.PrivateEndpointConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hosting Environment Private Endpoint Connection ID
func (id HostingEnvironmentPrivateEndpointConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticHostingEnvironments", "hostingEnvironments", "hostingEnvironments"),
		resourceids.UserSpecifiedSegment("hostingEnvironmentName", "hostingEnvironmentValue"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionValue"),
	}
}

// String returns a human-readable description of this Hosting Environment Private Endpoint Connection ID
func (id HostingEnvironmentPrivateEndpointConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Hosting Environment Name: %q", id.HostingEnvironmentName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
	}
	return fmt.Sprintf("Hosting Environment Private Endpoint Connection (%s)", strings.Join(components, "\n"))
}
