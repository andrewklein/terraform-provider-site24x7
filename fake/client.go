package fake

import (
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/fake"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/integration"
	"github.com/site24x7/terraform-provider-site24x7/api/endpoints/monitors"
)

// Client is an implementation of site24x7.Client that stubs out all endpoints
// with mocks. In can be used in unit tests.
type Client struct {
	FakeCurrentStatus          *fake.CurrentStatus
	FakeURLAutomations         *fake.URLAutomations
	FakeLocationTemplate       *fake.LocationTemplate
	FakeLocationProfiles       *fake.LocationProfiles
	FakeMonitorGroups          *fake.MonitorGroups
	FakeTags                   *fake.Tags
	FakeAmazonMonitors         *fake.AmazonMonitors
	FakeWebsiteMonitors        *fake.WebsiteMonitors
	FakeSSLMonitors            *fake.SSLMonitors
	FakeRestApiMonitors        *fake.RestApiMonitors
	FakeNotificationProfiles   *fake.NotificationProfiles
	FakeThresholdProfiles      *fake.ThresholdProfiles
	FakeUserGroups             *fake.UserGroups
	FakeOpsgenieIntegration    *fake.OpsgenieIntegration
	FakeSlackIntegration       *fake.SlackIntegration
	FakeWebhookIntegration     *fake.WebhookIntegration
	FakePagerDutyIntegration   *fake.PagerDutyIntegration
	FakeServiceNowIntegration  *fake.ServiceNowIntegration
	FakeThirdPartyIntegrations *fake.ThirdPartyIntegrations
}

// NewClient creates a new fake site24x7 API client.
func NewClient() *Client {
	return &Client{
		FakeCurrentStatus:          &fake.CurrentStatus{},
		FakeURLAutomations:         &fake.URLAutomations{},
		FakeLocationProfiles:       &fake.LocationProfiles{},
		FakeLocationTemplate:       &fake.LocationTemplate{},
		FakeMonitorGroups:          &fake.MonitorGroups{},
		FakeTags:                   &fake.Tags{},
		FakeAmazonMonitors:         &fake.AmazonMonitors{},
		FakeSSLMonitors:            &fake.SSLMonitors{},
		FakeWebsiteMonitors:        &fake.WebsiteMonitors{},
		FakeRestApiMonitors:        &fake.RestApiMonitors{},
		FakeNotificationProfiles:   &fake.NotificationProfiles{},
		FakeThresholdProfiles:      &fake.ThresholdProfiles{},
		FakeUserGroups:             &fake.UserGroups{},
		FakeOpsgenieIntegration:    &fake.OpsgenieIntegration{},
		FakeSlackIntegration:       &fake.SlackIntegration{},
		FakePagerDutyIntegration:   &fake.PagerDutyIntegration{},
		FakeServiceNowIntegration:  &fake.ServiceNowIntegration{},
		FakeWebhookIntegration:     &fake.WebhookIntegration{},
		FakeThirdPartyIntegrations: &fake.ThirdPartyIntegrations{},
	}
}

// CurrentStatus implements Client.
func (c *Client) CurrentStatus() endpoints.CurrentStatus {
	return c.FakeCurrentStatus
}

// ItAutomations implements Client.
func (c *Client) URLAutomations() endpoints.URLAutomations {
	return c.FakeURLAutomations
}

// LocationProfiles implements Client.
func (c *Client) LocationProfiles() endpoints.LocationProfiles {
	return c.FakeLocationProfiles
}

// LocationTemplate implements Client.
func (c *Client) LocationTemplate() endpoints.LocationTemplate {
	return c.FakeLocationTemplate
}

// Monitors implements Client.
func (c *Client) WebsiteMonitors() monitors.WebsiteMonitors {
	return c.FakeWebsiteMonitors
}

// SSLMonitors implements Client.
func (c *Client) SSLMonitors() monitors.SSLMonitors {
	return c.FakeSSLMonitors
}

// RestApiMonitors implements Client.
func (c *Client) RestApiMonitors() monitors.RestApiMonitors {
	return c.FakeRestApiMonitors
}

// Monitors implements Client.
func (c *Client) AmazonMonitors() monitors.AmazonMonitors {
	return c.FakeAmazonMonitors
}

// MonitorGroups implements Client.
func (c *Client) MonitorGroups() endpoints.MonitorGroups {
	return c.FakeMonitorGroups
}

// Tags implements Client.
func (c *Client) Tags() endpoints.Tags {
	return c.FakeTags
}

// NotificationProfiles implements Client.
func (c *Client) NotificationProfiles() endpoints.NotificationProfiles {
	return c.FakeNotificationProfiles
}

// ThresholdProfiles implements Client.
func (c *Client) ThresholdProfiles() endpoints.ThresholdProfiles {
	return c.FakeThresholdProfiles
}

// UserGroups implements Client.
func (c *Client) UserGroups() endpoints.UserGroups {
	return c.FakeUserGroups
}

// OpsgenieIntegration implements Client.
func (c *Client) OpsgenieIntegration() integration.OpsgenieIntegration {
	return c.FakeOpsgenieIntegration
}

// SlackIntegration implements Client.
func (c *Client) SlackIntegration() integration.SlackIntegration {
	return c.FakeSlackIntegration
}

// PagerDuty implements Client.
func (c *Client) PagerDutyIntegration() integration.PagerDutyIntegration {
	return c.FakePagerDutyIntegration
}

// PagerDuty implements Client.
func (c *Client) ServiceNowIntegration() integration.ServiceNowIntegration {
	return c.FakeServiceNowIntegration
}

// WebhookIntegration implements Client.
func (c *Client) WebhookIntegration() integration.WebhookIntegration {
	return c.FakeWebhookIntegration
}

// ThirdPartyIntegrations implements Client.
func (c *Client) ThirdPartyIntegrations() integration.ThirdpartyIntegrations {
	return c.FakeThirdPartyIntegrations
}
