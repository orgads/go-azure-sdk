package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotPremierAddonId{}

// SlotPremierAddonId is a struct representing the Resource ID for a Slot Premier Addon
type SlotPremierAddonId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	PremierAddonName  string
}

// NewSlotPremierAddonID returns a new SlotPremierAddonId struct
func NewSlotPremierAddonID(subscriptionId string, resourceGroupName string, siteName string, slotName string, premierAddonName string) SlotPremierAddonId {
	return SlotPremierAddonId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		PremierAddonName:  premierAddonName,
	}
}

// ParseSlotPremierAddonID parses 'input' into a SlotPremierAddonId
func ParseSlotPremierAddonID(input string) (*SlotPremierAddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotPremierAddonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotPremierAddonId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.PremierAddonName, ok = parsed.Parsed["premierAddonName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "premierAddonName", *parsed)
	}

	return &id, nil
}

// ParseSlotPremierAddonIDInsensitively parses 'input' case-insensitively into a SlotPremierAddonId
// note: this method should only be used for API response data and not user input
func ParseSlotPremierAddonIDInsensitively(input string) (*SlotPremierAddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotPremierAddonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotPremierAddonId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.PremierAddonName, ok = parsed.Parsed["premierAddonName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "premierAddonName", *parsed)
	}

	return &id, nil
}

// ValidateSlotPremierAddonID checks that 'input' can be parsed as a Slot Premier Addon ID
func ValidateSlotPremierAddonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotPremierAddonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Premier Addon ID
func (id SlotPremierAddonId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/premierAddons/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.PremierAddonName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Premier Addon ID
func (id SlotPremierAddonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotValue"),
		resourceids.StaticSegment("staticPremierAddons", "premierAddons", "premierAddons"),
		resourceids.UserSpecifiedSegment("premierAddonName", "premierAddonValue"),
	}
}

// String returns a human-readable description of this Slot Premier Addon ID
func (id SlotPremierAddonId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Premier Addon Name: %q", id.PremierAddonName),
	}
	return fmt.Sprintf("Slot Premier Addon (%s)", strings.Join(components, "\n"))
}
