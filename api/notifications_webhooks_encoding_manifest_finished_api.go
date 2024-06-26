package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// NotificationsWebhooksEncodingManifestFinishedAPI communicates with '/notifications/webhooks/encoding/manifest/finished' endpoints
type NotificationsWebhooksEncodingManifestFinishedAPI struct {
	apiClient *apiclient.APIClient
}

// NewNotificationsWebhooksEncodingManifestFinishedAPI constructor for NotificationsWebhooksEncodingManifestFinishedAPI that takes options as argument
func NewNotificationsWebhooksEncodingManifestFinishedAPI(options ...apiclient.APIClientOption) (*NotificationsWebhooksEncodingManifestFinishedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewNotificationsWebhooksEncodingManifestFinishedAPIWithClient(apiClient), nil
}

// NewNotificationsWebhooksEncodingManifestFinishedAPIWithClient constructor for NotificationsWebhooksEncodingManifestFinishedAPI that takes an APIClient as argument
func NewNotificationsWebhooksEncodingManifestFinishedAPIWithClient(apiClient *apiclient.APIClient) *NotificationsWebhooksEncodingManifestFinishedAPI {
	a := &NotificationsWebhooksEncodingManifestFinishedAPI{apiClient: apiClient}
	return a
}

// Create Add &#39;Manifest Finished Successfully&#39; Webhook (All Manifests)
func (api *NotificationsWebhooksEncodingManifestFinishedAPI) Create(webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/finished", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// CreateByManifestId Add &#39;Manifest Finished Successfully&#39; Webhook Notification (Specific Manifest)
func (api *NotificationsWebhooksEncodingManifestFinishedAPI) CreateByManifestId(manifestId string, webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.Webhook
	err := api.apiClient.Post("/notifications/webhooks/encoding/manifest/{manifest_id}/finished", &webhook, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Manifest Finished Webhook
func (api *NotificationsWebhooksEncodingManifestFinishedAPI) Delete(notificationId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/notifications/webhooks/encoding/manifest/finished/{notification_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get &#39;Manifest Finished&#39; Webhooks (All Manifests)
// Get all webhook notifications triggering when a manifest generation finishes successfully
func (api *NotificationsWebhooksEncodingManifestFinishedAPI) List() (*pagination.WebhooksListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel pagination.WebhooksListPagination
	err := api.apiClient.Get("/notifications/webhooks/encoding/manifest/finished", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update Replace Manifest Finished Webhook Notification
func (api *NotificationsWebhooksEncodingManifestFinishedAPI) Update(notificationId string, webhook model.Webhook) (*model.Webhook, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["notification_id"] = notificationId
	}

	var responseModel model.Webhook
	err := api.apiClient.Put("/notifications/webhooks/encoding/manifest/finished/{notification_id}", &webhook, &responseModel, reqParams)
	return &responseModel, err
}
