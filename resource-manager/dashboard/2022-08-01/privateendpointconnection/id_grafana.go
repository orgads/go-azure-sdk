package privateendpointconnection

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = GrafanaId{}

// GrafanaId is a struct representing the Resource ID for a Grafana
type GrafanaId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
}

// NewGrafanaID returns a new GrafanaId struct
func NewGrafanaID(subscriptionId string, resourceGroupName string, workspaceName string) GrafanaId {
	return GrafanaId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
	}
}

// ParseGrafanaID parses 'input' into a GrafanaId
func ParseGrafanaID(input string) (*GrafanaId, error) {
	parser := resourceids.NewParserFromResourceIdType(GrafanaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GrafanaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseGrafanaIDInsensitively parses 'input' case-insensitively into a GrafanaId
// note: this method should only be used for API response data and not user input
func ParseGrafanaIDInsensitively(input string) (*GrafanaId, error) {
	parser := resourceids.NewParserFromResourceIdType(GrafanaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GrafanaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateGrafanaID checks that 'input' can be parsed as a Grafana ID
func ValidateGrafanaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGrafanaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Grafana ID
func (id GrafanaId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Dashboard/grafana/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Grafana ID
func (id GrafanaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDashboard", "Microsoft.Dashboard", "Microsoft.Dashboard"),
		resourceids.StaticSegment("staticGrafana", "grafana", "grafana"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
	}
}

// String returns a human-readable description of this Grafana ID
func (id GrafanaId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
	}
	return fmt.Sprintf("Grafana (%s)", strings.Join(components, "\n"))
}
