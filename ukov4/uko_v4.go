/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.80.0-29334a73-20230925-151553
 */

// Package ukov4 : Operations and models for the UkoV4 service
package ukov4

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/ibm-hpcs-uko-sdk/common"
	"github.com/go-openapi/strfmt"
)

// UkoV4 : API for UKO used for key management.
//
// API Version: 4.14.5
type UkoV4 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "uko"

// UkoV4Options : Service options
type UkoV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUkoV4UsingExternalConfig : constructs an instance of UkoV4 with passed in options and external configuration.
func NewUkoV4UsingExternalConfig(options *UkoV4Options) (uko *UkoV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	uko, err = NewUkoV4(options)
	if err != nil {
		return
	}

	err = uko.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = uko.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUkoV4 : constructs an instance of UkoV4 with passed in options.
func NewUkoV4(options *UkoV4Options) (service *UkoV4, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &UkoV4{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "uko" suitable for processing requests.
func (uko *UkoV4) Clone() *UkoV4 {
	if core.IsNil(uko) {
		return nil
	}
	clone := *uko
	clone.Service = uko.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (uko *UkoV4) SetServiceURL(url string) error {
	return uko.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (uko *UkoV4) GetServiceURL() string {
	return uko.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (uko *UkoV4) SetDefaultHeaders(headers http.Header) {
	uko.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (uko *UkoV4) SetEnableGzipCompression(enableGzip bool) {
	uko.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (uko *UkoV4) GetEnableGzipCompression() bool {
	return uko.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (uko *UkoV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	uko.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (uko *UkoV4) DisableRetries() {
	uko.Service.DisableRetries()
}

// ListManagedKeys : List managed keys
// List all managed keys in the instance. It is possible to sort by the following parameters: label, algorithm, state,
// activation_date, deactivation_date, created_at, updated_at, size, vault.id.
func (uko *UkoV4) ListManagedKeys(listManagedKeysOptions *ListManagedKeysOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	return uko.ListManagedKeysWithContext(context.Background(), listManagedKeysOptions)
}

// ListManagedKeysWithContext is an alternate form of the ListManagedKeys method which supports a Context parameter
func (uko *UkoV4) ListManagedKeysWithContext(ctx context.Context, listManagedKeysOptions *ListManagedKeysOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listManagedKeysOptions, "listManagedKeysOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listManagedKeysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListManagedKeys")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listManagedKeysOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listManagedKeysOptions.Accept))
	}

	if listManagedKeysOptions.VaultID != nil {
		builder.AddQuery("vault.id", strings.Join(listManagedKeysOptions.VaultID, ","))
	}
	if listManagedKeysOptions.Algorithm != nil {
		builder.AddQuery("algorithm", strings.Join(listManagedKeysOptions.Algorithm, ","))
	}
	if listManagedKeysOptions.State != nil {
		builder.AddQuery("state", strings.Join(listManagedKeysOptions.State, ","))
	}
	if listManagedKeysOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listManagedKeysOptions.Limit))
	}
	if listManagedKeysOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listManagedKeysOptions.Offset))
	}
	if listManagedKeysOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listManagedKeysOptions.Sort, ","))
	}
	if listManagedKeysOptions.Label != nil {
		builder.AddQuery("label", fmt.Sprint(*listManagedKeysOptions.Label))
	}
	if listManagedKeysOptions.ActivationDate != nil {
		builder.AddQuery("activation_date", fmt.Sprint(*listManagedKeysOptions.ActivationDate))
	}
	if listManagedKeysOptions.ActivationDateMin != nil {
		builder.AddQuery("activation_date_min", fmt.Sprint(*listManagedKeysOptions.ActivationDateMin))
	}
	if listManagedKeysOptions.ActivationDateMax != nil {
		builder.AddQuery("activation_date_max", fmt.Sprint(*listManagedKeysOptions.ActivationDateMax))
	}
	if listManagedKeysOptions.DeactivationDate != nil {
		builder.AddQuery("deactivation_date", fmt.Sprint(*listManagedKeysOptions.DeactivationDate))
	}
	if listManagedKeysOptions.DeactivationDateMin != nil {
		builder.AddQuery("deactivation_date_min", fmt.Sprint(*listManagedKeysOptions.DeactivationDateMin))
	}
	if listManagedKeysOptions.DeactivationDateMax != nil {
		builder.AddQuery("deactivation_date_max", fmt.Sprint(*listManagedKeysOptions.DeactivationDateMax))
	}
	if listManagedKeysOptions.ExpirationDate != nil {
		builder.AddQuery("expiration_date", fmt.Sprint(*listManagedKeysOptions.ExpirationDate))
	}
	if listManagedKeysOptions.ExpirationDateMin != nil {
		builder.AddQuery("expiration_date_min", fmt.Sprint(*listManagedKeysOptions.ExpirationDateMin))
	}
	if listManagedKeysOptions.ExpirationDateMax != nil {
		builder.AddQuery("expiration_date_max", fmt.Sprint(*listManagedKeysOptions.ExpirationDateMax))
	}
	if listManagedKeysOptions.CreatedAt != nil {
		builder.AddQuery("created_at", fmt.Sprint(*listManagedKeysOptions.CreatedAt))
	}
	if listManagedKeysOptions.CreatedAtMin != nil {
		builder.AddQuery("created_at_min", fmt.Sprint(*listManagedKeysOptions.CreatedAtMin))
	}
	if listManagedKeysOptions.CreatedAtMax != nil {
		builder.AddQuery("created_at_max", fmt.Sprint(*listManagedKeysOptions.CreatedAtMax))
	}
	if listManagedKeysOptions.UpdatedAt != nil {
		builder.AddQuery("updated_at", fmt.Sprint(*listManagedKeysOptions.UpdatedAt))
	}
	if listManagedKeysOptions.UpdatedAtMin != nil {
		builder.AddQuery("updated_at_min", fmt.Sprint(*listManagedKeysOptions.UpdatedAtMin))
	}
	if listManagedKeysOptions.UpdatedAtMax != nil {
		builder.AddQuery("updated_at_max", fmt.Sprint(*listManagedKeysOptions.UpdatedAtMax))
	}
	if listManagedKeysOptions.RotatedAtMin != nil {
		builder.AddQuery("rotated_at_min", fmt.Sprint(*listManagedKeysOptions.RotatedAtMin))
	}
	if listManagedKeysOptions.RotatedAtMax != nil {
		builder.AddQuery("rotated_at_max", fmt.Sprint(*listManagedKeysOptions.RotatedAtMax))
	}
	if listManagedKeysOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*listManagedKeysOptions.Size))
	}
	if listManagedKeysOptions.SizeMin != nil {
		builder.AddQuery("size_min", fmt.Sprint(*listManagedKeysOptions.SizeMin))
	}
	if listManagedKeysOptions.SizeMax != nil {
		builder.AddQuery("size_max", fmt.Sprint(*listManagedKeysOptions.SizeMax))
	}
	if listManagedKeysOptions.ReferencedKeystoresType != nil {
		builder.AddQuery("referenced_keystores[].type", strings.Join(listManagedKeysOptions.ReferencedKeystoresType, ","))
	}
	if listManagedKeysOptions.ReferencedKeystoresName != nil {
		builder.AddQuery("referenced_keystores[].name", strings.Join(listManagedKeysOptions.ReferencedKeystoresName, ","))
	}
	if listManagedKeysOptions.InstancesKeystoreType != nil {
		builder.AddQuery("instances[].keystore.type", strings.Join(listManagedKeysOptions.InstancesKeystoreType, ","))
	}
	if listManagedKeysOptions.TemplateName != nil {
		builder.AddQuery("template.name", fmt.Sprint(*listManagedKeysOptions.TemplateName))
	}
	if listManagedKeysOptions.TemplateID != nil {
		builder.AddQuery("template.id", strings.Join(listManagedKeysOptions.TemplateID, ","))
	}
	if listManagedKeysOptions.TemplateType != nil {
		builder.AddQuery("template.type[]", strings.Join(listManagedKeysOptions.TemplateType, ","))
	}
	if listManagedKeysOptions.StatusInKeystoresKeystoreSyncFlag != nil {
		builder.AddQuery("status_in_keystores[].keystore_sync_flag", strings.Join(listManagedKeysOptions.StatusInKeystoresKeystoreSyncFlag, ","))
	}
	if listManagedKeysOptions.TemplateAlignmentStatus != nil {
		builder.AddQuery("template.alignment_status", fmt.Sprint(*listManagedKeysOptions.TemplateAlignmentStatus))
	}
	if listManagedKeysOptions.KeyMaterialPresent != nil {
		builder.AddQuery("key_material_present", strings.Join(listManagedKeysOptions.KeyMaterialPresent, ","))
	}
	if listManagedKeysOptions.ManagingSystems != nil {
		builder.AddQuery("managing_systems", strings.Join(listManagedKeysOptions.ManagingSystems, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKeyList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateManagedKey : Create a managed key
// Creates a new key based on the supplied template. The template must exist in the system prior to this call.
func (uko *UkoV4) CreateManagedKey(createManagedKeyOptions *CreateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.CreateManagedKeyWithContext(context.Background(), createManagedKeyOptions)
}

// CreateManagedKeyWithContext is an alternate form of the CreateManagedKey method which supports a Context parameter
func (uko *UkoV4) CreateManagedKeyWithContext(ctx context.Context, createManagedKeyOptions *CreateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createManagedKeyOptions, "createManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createManagedKeyOptions, "createManagedKeyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "CreateManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createManagedKeyOptions.TemplateName != nil {
		body["template_name"] = createManagedKeyOptions.TemplateName
	}
	if createManagedKeyOptions.Vault != nil {
		body["vault"] = createManagedKeyOptions.Vault
	}
	if createManagedKeyOptions.Label != nil {
		body["label"] = createManagedKeyOptions.Label
	}
	if createManagedKeyOptions.Description != nil {
		body["description"] = createManagedKeyOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteManagedKey : Delete a managed key
// Delete a managed key by ID from the vault. A key must be in a 'destroyed' state for it to be eligible for deletion.
func (uko *UkoV4) DeleteManagedKey(deleteManagedKeyOptions *DeleteManagedKeyOptions) (response *core.DetailedResponse, err error) {
	return uko.DeleteManagedKeyWithContext(context.Background(), deleteManagedKeyOptions)
}

// DeleteManagedKeyWithContext is an alternate form of the DeleteManagedKey method which supports a Context parameter
func (uko *UkoV4) DeleteManagedKeyWithContext(ctx context.Context, deleteManagedKeyOptions *DeleteManagedKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteManagedKeyOptions, "deleteManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteManagedKeyOptions, "deleteManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DeleteManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*deleteManagedKeyOptions.IfMatch))
	}

	if deleteManagedKeyOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*deleteManagedKeyOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = uko.Service.Request(request, nil)

	return
}

// GetManagedKey : Retrieve a managed key
// Retrieve a managed key and its details by specifying the ID.
func (uko *UkoV4) GetManagedKey(getManagedKeyOptions *GetManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.GetManagedKeyWithContext(context.Background(), getManagedKeyOptions)
}

// GetManagedKeyWithContext is an alternate form of the GetManagedKey method which supports a Context parameter
func (uko *UkoV4) GetManagedKeyWithContext(ctx context.Context, getManagedKeyOptions *GetManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getManagedKeyOptions, "getManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getManagedKeyOptions, "getManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateManagedKey : Update a managed key
// Update attributes of a managed key. It is only possible to modify the key's state separately from other changes.
// Changing a key's state affects its availablity for crypto operations in keystores.
func (uko *UkoV4) UpdateManagedKey(updateManagedKeyOptions *UpdateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.UpdateManagedKeyWithContext(context.Background(), updateManagedKeyOptions)
}

// UpdateManagedKeyWithContext is an alternate form of the UpdateManagedKey method which supports a Context parameter
func (uko *UkoV4) UpdateManagedKeyWithContext(ctx context.Context, updateManagedKeyOptions *UpdateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateManagedKeyOptions, "updateManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateManagedKeyOptions, "updateManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UpdateManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateManagedKeyOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateManagedKeyOptions.Label != nil {
		body["label"] = updateManagedKeyOptions.Label
	}
	if updateManagedKeyOptions.ActivationDate != nil {
		body["activation_date"] = updateManagedKeyOptions.ActivationDate
	}
	if updateManagedKeyOptions.ExpirationDate != nil {
		body["expiration_date"] = updateManagedKeyOptions.ExpirationDate
	}
	if updateManagedKeyOptions.Description != nil {
		body["description"] = updateManagedKeyOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAssociatedResourcesForManagedKey : List associated resources for a managed key
// You can use this endpoint to obtain a list of resources associated with this managed key in IBM Cloud; which cloud
// resources are protected by the key you specify.
func (uko *UkoV4) ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptions *ListAssociatedResourcesForManagedKeyOptions) (result *AssociatedResourceList, response *core.DetailedResponse, err error) {
	return uko.ListAssociatedResourcesForManagedKeyWithContext(context.Background(), listAssociatedResourcesForManagedKeyOptions)
}

// ListAssociatedResourcesForManagedKeyWithContext is an alternate form of the ListAssociatedResourcesForManagedKey method which supports a Context parameter
func (uko *UkoV4) ListAssociatedResourcesForManagedKeyWithContext(ctx context.Context, listAssociatedResourcesForManagedKeyOptions *ListAssociatedResourcesForManagedKeyOptions) (result *AssociatedResourceList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAssociatedResourcesForManagedKeyOptions, "listAssociatedResourcesForManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAssociatedResourcesForManagedKeyOptions, "listAssociatedResourcesForManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listAssociatedResourcesForManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/associated_resources`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAssociatedResourcesForManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListAssociatedResourcesForManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAssociatedResourcesForManagedKeyOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAssociatedResourcesForManagedKeyOptions.Limit))
	}
	if listAssociatedResourcesForManagedKeyOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listAssociatedResourcesForManagedKeyOptions.Offset))
	}
	if listAssociatedResourcesForManagedKeyOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listAssociatedResourcesForManagedKeyOptions.Sort, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAssociatedResourceList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListManagedKeyVersions : List managed key versions
// List all managed key versions in the instance.
func (uko *UkoV4) ListManagedKeyVersions(listManagedKeyVersionsOptions *ListManagedKeyVersionsOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	return uko.ListManagedKeyVersionsWithContext(context.Background(), listManagedKeyVersionsOptions)
}

// ListManagedKeyVersionsWithContext is an alternate form of the ListManagedKeyVersions method which supports a Context parameter
func (uko *UkoV4) ListManagedKeyVersionsWithContext(ctx context.Context, listManagedKeyVersionsOptions *ListManagedKeyVersionsOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listManagedKeyVersionsOptions, "listManagedKeyVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listManagedKeyVersionsOptions, "listManagedKeyVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listManagedKeyVersionsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listManagedKeyVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListManagedKeyVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listManagedKeyVersionsOptions.Algorithm != nil {
		builder.AddQuery("algorithm", strings.Join(listManagedKeyVersionsOptions.Algorithm, ","))
	}
	if listManagedKeyVersionsOptions.State != nil {
		builder.AddQuery("state", strings.Join(listManagedKeyVersionsOptions.State, ","))
	}
	if listManagedKeyVersionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listManagedKeyVersionsOptions.Limit))
	}
	if listManagedKeyVersionsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listManagedKeyVersionsOptions.Offset))
	}
	if listManagedKeyVersionsOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listManagedKeyVersionsOptions.Sort, ","))
	}
	if listManagedKeyVersionsOptions.Label != nil {
		builder.AddQuery("label", fmt.Sprint(*listManagedKeyVersionsOptions.Label))
	}
	if listManagedKeyVersionsOptions.ActivationDate != nil {
		builder.AddQuery("activation_date", fmt.Sprint(*listManagedKeyVersionsOptions.ActivationDate))
	}
	if listManagedKeyVersionsOptions.ActivationDateMin != nil {
		builder.AddQuery("activation_date_min", fmt.Sprint(*listManagedKeyVersionsOptions.ActivationDateMin))
	}
	if listManagedKeyVersionsOptions.ActivationDateMax != nil {
		builder.AddQuery("activation_date_max", fmt.Sprint(*listManagedKeyVersionsOptions.ActivationDateMax))
	}
	if listManagedKeyVersionsOptions.DeactivationDate != nil {
		builder.AddQuery("deactivation_date", fmt.Sprint(*listManagedKeyVersionsOptions.DeactivationDate))
	}
	if listManagedKeyVersionsOptions.DeactivationDateMin != nil {
		builder.AddQuery("deactivation_date_min", fmt.Sprint(*listManagedKeyVersionsOptions.DeactivationDateMin))
	}
	if listManagedKeyVersionsOptions.DeactivationDateMax != nil {
		builder.AddQuery("deactivation_date_max", fmt.Sprint(*listManagedKeyVersionsOptions.DeactivationDateMax))
	}
	if listManagedKeyVersionsOptions.ExpirationDate != nil {
		builder.AddQuery("expiration_date", fmt.Sprint(*listManagedKeyVersionsOptions.ExpirationDate))
	}
	if listManagedKeyVersionsOptions.ExpirationDateMin != nil {
		builder.AddQuery("expiration_date_min", fmt.Sprint(*listManagedKeyVersionsOptions.ExpirationDateMin))
	}
	if listManagedKeyVersionsOptions.ExpirationDateMax != nil {
		builder.AddQuery("expiration_date_max", fmt.Sprint(*listManagedKeyVersionsOptions.ExpirationDateMax))
	}
	if listManagedKeyVersionsOptions.CreatedAt != nil {
		builder.AddQuery("created_at", fmt.Sprint(*listManagedKeyVersionsOptions.CreatedAt))
	}
	if listManagedKeyVersionsOptions.CreatedAtMin != nil {
		builder.AddQuery("created_at_min", fmt.Sprint(*listManagedKeyVersionsOptions.CreatedAtMin))
	}
	if listManagedKeyVersionsOptions.CreatedAtMax != nil {
		builder.AddQuery("created_at_max", fmt.Sprint(*listManagedKeyVersionsOptions.CreatedAtMax))
	}
	if listManagedKeyVersionsOptions.UpdatedAt != nil {
		builder.AddQuery("updated_at", fmt.Sprint(*listManagedKeyVersionsOptions.UpdatedAt))
	}
	if listManagedKeyVersionsOptions.UpdatedAtMin != nil {
		builder.AddQuery("updated_at_min", fmt.Sprint(*listManagedKeyVersionsOptions.UpdatedAtMin))
	}
	if listManagedKeyVersionsOptions.UpdatedAtMax != nil {
		builder.AddQuery("updated_at_max", fmt.Sprint(*listManagedKeyVersionsOptions.UpdatedAtMax))
	}
	if listManagedKeyVersionsOptions.RotatedAtMin != nil {
		builder.AddQuery("rotated_at_min", fmt.Sprint(*listManagedKeyVersionsOptions.RotatedAtMin))
	}
	if listManagedKeyVersionsOptions.RotatedAtMax != nil {
		builder.AddQuery("rotated_at_max", fmt.Sprint(*listManagedKeyVersionsOptions.RotatedAtMax))
	}
	if listManagedKeyVersionsOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*listManagedKeyVersionsOptions.Size))
	}
	if listManagedKeyVersionsOptions.SizeMin != nil {
		builder.AddQuery("size_min", fmt.Sprint(*listManagedKeyVersionsOptions.SizeMin))
	}
	if listManagedKeyVersionsOptions.SizeMax != nil {
		builder.AddQuery("size_max", fmt.Sprint(*listManagedKeyVersionsOptions.SizeMax))
	}
	if listManagedKeyVersionsOptions.ReferencedKeystoresType != nil {
		builder.AddQuery("referenced_keystores[].type", strings.Join(listManagedKeyVersionsOptions.ReferencedKeystoresType, ","))
	}
	if listManagedKeyVersionsOptions.ReferencedKeystoresName != nil {
		builder.AddQuery("referenced_keystores[].name", strings.Join(listManagedKeyVersionsOptions.ReferencedKeystoresName, ","))
	}
	if listManagedKeyVersionsOptions.InstancesKeystoreType != nil {
		builder.AddQuery("instances[].keystore.type", strings.Join(listManagedKeyVersionsOptions.InstancesKeystoreType, ","))
	}
	if listManagedKeyVersionsOptions.TemplateName != nil {
		builder.AddQuery("template.name", fmt.Sprint(*listManagedKeyVersionsOptions.TemplateName))
	}
	if listManagedKeyVersionsOptions.TemplateID != nil {
		builder.AddQuery("template.id", strings.Join(listManagedKeyVersionsOptions.TemplateID, ","))
	}
	if listManagedKeyVersionsOptions.TemplateType != nil {
		builder.AddQuery("template.type[]", strings.Join(listManagedKeyVersionsOptions.TemplateType, ","))
	}
	if listManagedKeyVersionsOptions.StatusInKeystoresKeystoreSyncFlag != nil {
		builder.AddQuery("status_in_keystores[].keystore_sync_flag", strings.Join(listManagedKeyVersionsOptions.StatusInKeystoresKeystoreSyncFlag, ","))
	}
	if listManagedKeyVersionsOptions.TemplateAlignmentStatus != nil {
		builder.AddQuery("template.alignment_status", fmt.Sprint(*listManagedKeyVersionsOptions.TemplateAlignmentStatus))
	}
	if listManagedKeyVersionsOptions.KeyMaterialPresent != nil {
		builder.AddQuery("key_material_present", strings.Join(listManagedKeyVersionsOptions.KeyMaterialPresent, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKeyList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetKeyDistributionStatusForKeystores : Retrieve distribution status for all keystores
// Return distribution status for all keystores for a key instance. If there's any problems reading the keystore status
// of the key instance, http code 200 will still be returned, and the error code will be returned alongside an 'error'
// keystore status.
func (uko *UkoV4) GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptions *GetKeyDistributionStatusForKeystoresOptions) (result *StatusInKeystores, response *core.DetailedResponse, err error) {
	return uko.GetKeyDistributionStatusForKeystoresWithContext(context.Background(), getKeyDistributionStatusForKeystoresOptions)
}

// GetKeyDistributionStatusForKeystoresWithContext is an alternate form of the GetKeyDistributionStatusForKeystores method which supports a Context parameter
func (uko *UkoV4) GetKeyDistributionStatusForKeystoresWithContext(ctx context.Context, getKeyDistributionStatusForKeystoresOptions *GetKeyDistributionStatusForKeystoresOptions) (result *StatusInKeystores, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getKeyDistributionStatusForKeystoresOptions, "getKeyDistributionStatusForKeystoresOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getKeyDistributionStatusForKeystoresOptions, "getKeyDistributionStatusForKeystoresOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getKeyDistributionStatusForKeystoresOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/status_in_keystores`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getKeyDistributionStatusForKeystoresOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetKeyDistributionStatusForKeystores")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatusInKeystores)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateManagedKeyFromTemplate : Update a managed key to match the key template
// Update a managed key to match the latest version of the associated key template. It will install, activate, or
// deactivate the key on target keystores in the group defined by the key template.
func (uko *UkoV4) UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptions *UpdateManagedKeyFromTemplateOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.UpdateManagedKeyFromTemplateWithContext(context.Background(), updateManagedKeyFromTemplateOptions)
}

// UpdateManagedKeyFromTemplateWithContext is an alternate form of the UpdateManagedKeyFromTemplate method which supports a Context parameter
func (uko *UkoV4) UpdateManagedKeyFromTemplateWithContext(ctx context.Context, updateManagedKeyFromTemplateOptions *UpdateManagedKeyFromTemplateOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateManagedKeyFromTemplateOptions, "updateManagedKeyFromTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateManagedKeyFromTemplateOptions, "updateManagedKeyFromTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateManagedKeyFromTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/update_from_template`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateManagedKeyFromTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UpdateManagedKeyFromTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if updateManagedKeyFromTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateManagedKeyFromTemplateOptions.IfMatch))
	}

	if updateManagedKeyFromTemplateOptions.DryRun != nil {
		builder.AddQuery("dry_run", fmt.Sprint(*updateManagedKeyFromTemplateOptions.DryRun))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ActivateManagedKey : Activate a managed key
// Activate a managed key and perform key installation or activation operations on keystores in the keystore group
// associated with the managed key.
func (uko *UkoV4) ActivateManagedKey(activateManagedKeyOptions *ActivateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.ActivateManagedKeyWithContext(context.Background(), activateManagedKeyOptions)
}

// ActivateManagedKeyWithContext is an alternate form of the ActivateManagedKey method which supports a Context parameter
func (uko *UkoV4) ActivateManagedKeyWithContext(ctx context.Context, activateManagedKeyOptions *ActivateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(activateManagedKeyOptions, "activateManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(activateManagedKeyOptions, "activateManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *activateManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/activate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range activateManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ActivateManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if activateManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*activateManagedKeyOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeactivateManagedKey : Deactivate a managed key
// Deactivates a managed key and performs key deactivation operations on keystores in the keystore group associated with
// the managed key.
func (uko *UkoV4) DeactivateManagedKey(deactivateManagedKeyOptions *DeactivateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.DeactivateManagedKeyWithContext(context.Background(), deactivateManagedKeyOptions)
}

// DeactivateManagedKeyWithContext is an alternate form of the DeactivateManagedKey method which supports a Context parameter
func (uko *UkoV4) DeactivateManagedKeyWithContext(ctx context.Context, deactivateManagedKeyOptions *DeactivateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deactivateManagedKeyOptions, "deactivateManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deactivateManagedKeyOptions, "deactivateManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deactivateManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/deactivate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deactivateManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DeactivateManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deactivateManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*deactivateManagedKeyOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DestroyManagedKey : Destroy a managed key
// Destroy a managed key and perform key destruction operations on keystores in the keystore group associated with the
// managed key. This operation cannot be undone. The managed key must be in a 'deactivated' state.
func (uko *UkoV4) DestroyManagedKey(destroyManagedKeyOptions *DestroyManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.DestroyManagedKeyWithContext(context.Background(), destroyManagedKeyOptions)
}

// DestroyManagedKeyWithContext is an alternate form of the DestroyManagedKey method which supports a Context parameter
func (uko *UkoV4) DestroyManagedKeyWithContext(ctx context.Context, destroyManagedKeyOptions *DestroyManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(destroyManagedKeyOptions, "destroyManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(destroyManagedKeyOptions, "destroyManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *destroyManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/destroy`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range destroyManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DestroyManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if destroyManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*destroyManagedKeyOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SyncManagedKey : Sync a managed key in keystores
// Perform the synchronization operation on a managed key to align the states in the associated keystores.
func (uko *UkoV4) SyncManagedKey(syncManagedKeyOptions *SyncManagedKeyOptions) (result *StatusInKeystores, response *core.DetailedResponse, err error) {
	return uko.SyncManagedKeyWithContext(context.Background(), syncManagedKeyOptions)
}

// SyncManagedKeyWithContext is an alternate form of the SyncManagedKey method which supports a Context parameter
func (uko *UkoV4) SyncManagedKeyWithContext(ctx context.Context, syncManagedKeyOptions *SyncManagedKeyOptions) (result *StatusInKeystores, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(syncManagedKeyOptions, "syncManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(syncManagedKeyOptions, "syncManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *syncManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/sync_status_in_keystores`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range syncManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "SyncManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if syncManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*syncManagedKeyOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatusInKeystores)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RotateManagedKey : Rotate the managed key
// Rotate the managed key.
func (uko *UkoV4) RotateManagedKey(rotateManagedKeyOptions *RotateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	return uko.RotateManagedKeyWithContext(context.Background(), rotateManagedKeyOptions)
}

// RotateManagedKeyWithContext is an alternate form of the RotateManagedKey method which supports a Context parameter
func (uko *UkoV4) RotateManagedKeyWithContext(ctx context.Context, rotateManagedKeyOptions *RotateManagedKeyOptions) (result *ManagedKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(rotateManagedKeyOptions, "rotateManagedKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(rotateManagedKeyOptions, "rotateManagedKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *rotateManagedKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/managed_keys/{id}/rotate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range rotateManagedKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "RotateManagedKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if rotateManagedKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*rotateManagedKeyOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKey)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListKeyTemplates : List key templates
// List all key templates in the instance.
func (uko *UkoV4) ListKeyTemplates(listKeyTemplatesOptions *ListKeyTemplatesOptions) (result *TemplateList, response *core.DetailedResponse, err error) {
	return uko.ListKeyTemplatesWithContext(context.Background(), listKeyTemplatesOptions)
}

// ListKeyTemplatesWithContext is an alternate form of the ListKeyTemplates method which supports a Context parameter
func (uko *UkoV4) ListKeyTemplatesWithContext(ctx context.Context, listKeyTemplatesOptions *ListKeyTemplatesOptions) (result *TemplateList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listKeyTemplatesOptions, "listKeyTemplatesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listKeyTemplatesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListKeyTemplates")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listKeyTemplatesOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listKeyTemplatesOptions.Accept))
	}

	if listKeyTemplatesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listKeyTemplatesOptions.Name))
	}
	if listKeyTemplatesOptions.NamingScheme != nil {
		builder.AddQuery("naming_scheme", fmt.Sprint(*listKeyTemplatesOptions.NamingScheme))
	}
	if listKeyTemplatesOptions.VaultID != nil {
		builder.AddQuery("vault.id", strings.Join(listKeyTemplatesOptions.VaultID, ","))
	}
	if listKeyTemplatesOptions.KeyAlgorithm != nil {
		builder.AddQuery("key.algorithm", strings.Join(listKeyTemplatesOptions.KeyAlgorithm, ","))
	}
	if listKeyTemplatesOptions.KeySize != nil {
		builder.AddQuery("key.size", fmt.Sprint(*listKeyTemplatesOptions.KeySize))
	}
	if listKeyTemplatesOptions.KeySizeMin != nil {
		builder.AddQuery("key.size_min", fmt.Sprint(*listKeyTemplatesOptions.KeySizeMin))
	}
	if listKeyTemplatesOptions.KeySizeMax != nil {
		builder.AddQuery("key.size_max", fmt.Sprint(*listKeyTemplatesOptions.KeySizeMax))
	}
	if listKeyTemplatesOptions.KeystoresType != nil {
		builder.AddQuery("keystores[].type", strings.Join(listKeyTemplatesOptions.KeystoresType, ","))
	}
	if listKeyTemplatesOptions.KeystoresGroup != nil {
		builder.AddQuery("keystores[].group", strings.Join(listKeyTemplatesOptions.KeystoresGroup, ","))
	}
	if listKeyTemplatesOptions.CreatedAt != nil {
		builder.AddQuery("created_at", fmt.Sprint(*listKeyTemplatesOptions.CreatedAt))
	}
	if listKeyTemplatesOptions.CreatedAtMin != nil {
		builder.AddQuery("created_at_min", fmt.Sprint(*listKeyTemplatesOptions.CreatedAtMin))
	}
	if listKeyTemplatesOptions.CreatedAtMax != nil {
		builder.AddQuery("created_at_max", fmt.Sprint(*listKeyTemplatesOptions.CreatedAtMax))
	}
	if listKeyTemplatesOptions.UpdatedAt != nil {
		builder.AddQuery("updated_at", fmt.Sprint(*listKeyTemplatesOptions.UpdatedAt))
	}
	if listKeyTemplatesOptions.UpdatedAtMin != nil {
		builder.AddQuery("updated_at_min", fmt.Sprint(*listKeyTemplatesOptions.UpdatedAtMin))
	}
	if listKeyTemplatesOptions.UpdatedAtMax != nil {
		builder.AddQuery("updated_at_max", fmt.Sprint(*listKeyTemplatesOptions.UpdatedAtMax))
	}
	if listKeyTemplatesOptions.Type != nil {
		builder.AddQuery("type[]", strings.Join(listKeyTemplatesOptions.Type, ","))
	}
	if listKeyTemplatesOptions.State != nil {
		builder.AddQuery("state", strings.Join(listKeyTemplatesOptions.State, ","))
	}
	if listKeyTemplatesOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listKeyTemplatesOptions.Sort, ","))
	}
	if listKeyTemplatesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listKeyTemplatesOptions.Limit))
	}
	if listKeyTemplatesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listKeyTemplatesOptions.Offset))
	}
	if listKeyTemplatesOptions.ManagingSystems != nil {
		builder.AddQuery("managing_systems", strings.Join(listKeyTemplatesOptions.ManagingSystems, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplateList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateKeyTemplate : Create a key template
// Create a new key template. Key templates are used to combine information necessary when creating a key that allow
// easy subsequent key creation, without needing to specify any of its details.
func (uko *UkoV4) CreateKeyTemplate(createKeyTemplateOptions *CreateKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.CreateKeyTemplateWithContext(context.Background(), createKeyTemplateOptions)
}

// CreateKeyTemplateWithContext is an alternate form of the CreateKeyTemplate method which supports a Context parameter
func (uko *UkoV4) CreateKeyTemplateWithContext(ctx context.Context, createKeyTemplateOptions *CreateKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createKeyTemplateOptions, "createKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createKeyTemplateOptions, "createKeyTemplateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "CreateKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createKeyTemplateOptions.Vault != nil {
		body["vault"] = createKeyTemplateOptions.Vault
	}
	if createKeyTemplateOptions.Name != nil {
		body["name"] = createKeyTemplateOptions.Name
	}
	if createKeyTemplateOptions.Key != nil {
		body["key"] = createKeyTemplateOptions.Key
	}
	if createKeyTemplateOptions.Keystores != nil {
		body["keystores"] = createKeyTemplateOptions.Keystores
	}
	if createKeyTemplateOptions.Description != nil {
		body["description"] = createKeyTemplateOptions.Description
	}
	if createKeyTemplateOptions.NamingScheme != nil {
		body["naming_scheme"] = createKeyTemplateOptions.NamingScheme
	}
	if createKeyTemplateOptions.Type != nil {
		body["type"] = createKeyTemplateOptions.Type
	}
	if createKeyTemplateOptions.State != nil {
		body["state"] = createKeyTemplateOptions.State
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteKeyTemplate : Delete a template
// Delete a key template from the vault. The key template must not have any managed keys associated with it for it to be
// eligible for deletion.
func (uko *UkoV4) DeleteKeyTemplate(deleteKeyTemplateOptions *DeleteKeyTemplateOptions) (response *core.DetailedResponse, err error) {
	return uko.DeleteKeyTemplateWithContext(context.Background(), deleteKeyTemplateOptions)
}

// DeleteKeyTemplateWithContext is an alternate form of the DeleteKeyTemplate method which supports a Context parameter
func (uko *UkoV4) DeleteKeyTemplateWithContext(ctx context.Context, deleteKeyTemplateOptions *DeleteKeyTemplateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteKeyTemplateOptions, "deleteKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteKeyTemplateOptions, "deleteKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DeleteKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteKeyTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*deleteKeyTemplateOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = uko.Service.Request(request, nil)

	return
}

// GetKeyTemplate : Retrieve a key template
// Retrieve a key template and its details by specifying the ID.
func (uko *UkoV4) GetKeyTemplate(getKeyTemplateOptions *GetKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.GetKeyTemplateWithContext(context.Background(), getKeyTemplateOptions)
}

// GetKeyTemplateWithContext is an alternate form of the GetKeyTemplate method which supports a Context parameter
func (uko *UkoV4) GetKeyTemplateWithContext(ctx context.Context, getKeyTemplateOptions *GetKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getKeyTemplateOptions, "getKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getKeyTemplateOptions, "getKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateKeyTemplate : Update a key template
// Update attributes of a key template.
func (uko *UkoV4) UpdateKeyTemplate(updateKeyTemplateOptions *UpdateKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.UpdateKeyTemplateWithContext(context.Background(), updateKeyTemplateOptions)
}

// UpdateKeyTemplateWithContext is an alternate form of the UpdateKeyTemplate method which supports a Context parameter
func (uko *UkoV4) UpdateKeyTemplateWithContext(ctx context.Context, updateKeyTemplateOptions *UpdateKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateKeyTemplateOptions, "updateKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateKeyTemplateOptions, "updateKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UpdateKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateKeyTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateKeyTemplateOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateKeyTemplateOptions.Name != nil {
		body["name"] = updateKeyTemplateOptions.Name
	}
	if updateKeyTemplateOptions.Keystores != nil {
		body["keystores"] = updateKeyTemplateOptions.Keystores
	}
	if updateKeyTemplateOptions.Description != nil {
		body["description"] = updateKeyTemplateOptions.Description
	}
	if updateKeyTemplateOptions.Key != nil {
		body["key"] = updateKeyTemplateOptions.Key
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListKeystores : List all target keystores
// List all target keystores in the instance.
func (uko *UkoV4) ListKeystores(listKeystoresOptions *ListKeystoresOptions) (result *KeystoreList, response *core.DetailedResponse, err error) {
	return uko.ListKeystoresWithContext(context.Background(), listKeystoresOptions)
}

// ListKeystoresWithContext is an alternate form of the ListKeystores method which supports a Context parameter
func (uko *UkoV4) ListKeystoresWithContext(ctx context.Context, listKeystoresOptions *ListKeystoresOptions) (result *KeystoreList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listKeystoresOptions, "listKeystoresOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listKeystoresOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListKeystores")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listKeystoresOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listKeystoresOptions.Accept))
	}

	if listKeystoresOptions.Type != nil {
		builder.AddQuery("type", strings.Join(listKeystoresOptions.Type, ","))
	}
	if listKeystoresOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listKeystoresOptions.Name))
	}
	if listKeystoresOptions.Description != nil {
		builder.AddQuery("description", fmt.Sprint(*listKeystoresOptions.Description))
	}
	if listKeystoresOptions.Group != nil {
		builder.AddQuery("group", fmt.Sprint(*listKeystoresOptions.Group))
	}
	if listKeystoresOptions.Groups != nil {
		builder.AddQuery("groups[]", fmt.Sprint(*listKeystoresOptions.Groups))
	}
	if listKeystoresOptions.VaultID != nil {
		builder.AddQuery("vault.id", strings.Join(listKeystoresOptions.VaultID, ","))
	}
	if listKeystoresOptions.Location != nil {
		builder.AddQuery("location", strings.Join(listKeystoresOptions.Location, ","))
	}
	if listKeystoresOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listKeystoresOptions.Limit))
	}
	if listKeystoresOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listKeystoresOptions.Offset))
	}
	if listKeystoresOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listKeystoresOptions.Sort, ","))
	}
	if listKeystoresOptions.StatusHealthStatus != nil {
		builder.AddQuery("status.health_status", strings.Join(listKeystoresOptions.StatusHealthStatus, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKeystoreList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateKeystore : Create an internal keystore or a keystore connection
// Create a new internal keystore or a connection to an external keystore of the requested type.  If the `dry_run` query
// parameter is used, then a new keystore is not created in the database,  only a test is performed to verify if the
// connection information is correct. It is possible to sort by the following parameters: name, created_at, updated_at,
// vault.id.
func (uko *UkoV4) CreateKeystore(createKeystoreOptions *CreateKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	return uko.CreateKeystoreWithContext(context.Background(), createKeystoreOptions)
}

// CreateKeystoreWithContext is an alternate form of the CreateKeystore method which supports a Context parameter
func (uko *UkoV4) CreateKeystoreWithContext(ctx context.Context, createKeystoreOptions *CreateKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createKeystoreOptions, "createKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createKeystoreOptions, "createKeystoreOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "CreateKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createKeystoreOptions.DryRun != nil {
		builder.AddQuery("dry_run", fmt.Sprint(*createKeystoreOptions.DryRun))
	}

	_, err = builder.SetBodyContentJSON(createKeystoreOptions.KeystoreBody)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKeystore)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteKeystore : Delete an internal keystore or a connection to an external keystore
// Delete an internal keystore or a connection to an external keystore (in that case, the keystore on the remote system
// is unchanged).
func (uko *UkoV4) DeleteKeystore(deleteKeystoreOptions *DeleteKeystoreOptions) (response *core.DetailedResponse, err error) {
	return uko.DeleteKeystoreWithContext(context.Background(), deleteKeystoreOptions)
}

// DeleteKeystoreWithContext is an alternate form of the DeleteKeystore method which supports a Context parameter
func (uko *UkoV4) DeleteKeystoreWithContext(ctx context.Context, deleteKeystoreOptions *DeleteKeystoreOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteKeystoreOptions, "deleteKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteKeystoreOptions, "deleteKeystoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteKeystoreOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DeleteKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteKeystoreOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*deleteKeystoreOptions.IfMatch))
	}

	if deleteKeystoreOptions.Mode != nil {
		builder.AddQuery("mode", fmt.Sprint(*deleteKeystoreOptions.Mode))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = uko.Service.Request(request, nil)

	return
}

// GetKeystore : Retrieve a target keystore
// Retrieve a target keystore (either an internal keystore or a keystore connection) and its details by specifying the
// ID.
func (uko *UkoV4) GetKeystore(getKeystoreOptions *GetKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	return uko.GetKeystoreWithContext(context.Background(), getKeystoreOptions)
}

// GetKeystoreWithContext is an alternate form of the GetKeystore method which supports a Context parameter
func (uko *UkoV4) GetKeystoreWithContext(ctx context.Context, getKeystoreOptions *GetKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getKeystoreOptions, "getKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getKeystoreOptions, "getKeystoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getKeystoreOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKeystore)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateKeystore : Update an internal keystore or a keystore connection
// Updates attributes of an internal keystore or a keystore connection.
func (uko *UkoV4) UpdateKeystore(updateKeystoreOptions *UpdateKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	return uko.UpdateKeystoreWithContext(context.Background(), updateKeystoreOptions)
}

// UpdateKeystoreWithContext is an alternate form of the UpdateKeystore method which supports a Context parameter
func (uko *UkoV4) UpdateKeystoreWithContext(ctx context.Context, updateKeystoreOptions *UpdateKeystoreOptions) (result KeystoreIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateKeystoreOptions, "updateKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateKeystoreOptions, "updateKeystoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateKeystoreOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UpdateKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateKeystoreOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateKeystoreOptions.IfMatch))
	}

	_, err = builder.SetBodyContentJSON(updateKeystoreOptions.KeystoreBody)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKeystore)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAssociatedResourcesForTargetKeystore : List associated resources for a target keystore
// You can use this endpoint to obtain a list of resources associated with all keys referencing this keystore.
func (uko *UkoV4) ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptions *ListAssociatedResourcesForTargetKeystoreOptions) (result *AssociatedResourceList, response *core.DetailedResponse, err error) {
	return uko.ListAssociatedResourcesForTargetKeystoreWithContext(context.Background(), listAssociatedResourcesForTargetKeystoreOptions)
}

// ListAssociatedResourcesForTargetKeystoreWithContext is an alternate form of the ListAssociatedResourcesForTargetKeystore method which supports a Context parameter
func (uko *UkoV4) ListAssociatedResourcesForTargetKeystoreWithContext(ctx context.Context, listAssociatedResourcesForTargetKeystoreOptions *ListAssociatedResourcesForTargetKeystoreOptions) (result *AssociatedResourceList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAssociatedResourcesForTargetKeystoreOptions, "listAssociatedResourcesForTargetKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAssociatedResourcesForTargetKeystoreOptions, "listAssociatedResourcesForTargetKeystoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listAssociatedResourcesForTargetKeystoreOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}/associated_resources`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAssociatedResourcesForTargetKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListAssociatedResourcesForTargetKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAssociatedResourcesForTargetKeystoreOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAssociatedResourcesForTargetKeystoreOptions.Limit))
	}
	if listAssociatedResourcesForTargetKeystoreOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listAssociatedResourcesForTargetKeystoreOptions.Offset))
	}
	if listAssociatedResourcesForTargetKeystoreOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listAssociatedResourcesForTargetKeystoreOptions.Sort, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAssociatedResourceList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetKeystoreStatus : Retrieve target keystore status
// Retrieve status of a single target keystore (either a keystore connection or an internal keystore).
func (uko *UkoV4) GetKeystoreStatus(getKeystoreStatusOptions *GetKeystoreStatusOptions) (result *KeystoreStatus, response *core.DetailedResponse, err error) {
	return uko.GetKeystoreStatusWithContext(context.Background(), getKeystoreStatusOptions)
}

// GetKeystoreStatusWithContext is an alternate form of the GetKeystoreStatus method which supports a Context parameter
func (uko *UkoV4) GetKeystoreStatusWithContext(ctx context.Context, getKeystoreStatusOptions *GetKeystoreStatusOptions) (result *KeystoreStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getKeystoreStatusOptions, "getKeystoreStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getKeystoreStatusOptions, "getKeystoreStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getKeystoreStatusOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}/status`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getKeystoreStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetKeystoreStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKeystoreStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListManagedKeysFromKeystore : List managed keys on the target keystore
// Lists all managed keys installed on the target keystore (either a keystore connection or an internal keystore). Note
// that `pre_activation` and `destroyed` keys are not installed.
func (uko *UkoV4) ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptions *ListManagedKeysFromKeystoreOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	return uko.ListManagedKeysFromKeystoreWithContext(context.Background(), listManagedKeysFromKeystoreOptions)
}

// ListManagedKeysFromKeystoreWithContext is an alternate form of the ListManagedKeysFromKeystore method which supports a Context parameter
func (uko *UkoV4) ListManagedKeysFromKeystoreWithContext(ctx context.Context, listManagedKeysFromKeystoreOptions *ListManagedKeysFromKeystoreOptions) (result *ManagedKeyList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listManagedKeysFromKeystoreOptions, "listManagedKeysFromKeystoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listManagedKeysFromKeystoreOptions, "listManagedKeysFromKeystoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listManagedKeysFromKeystoreOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/keystores/{id}/managed_keys`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listManagedKeysFromKeystoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListManagedKeysFromKeystore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listManagedKeysFromKeystoreOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listManagedKeysFromKeystoreOptions.Accept))
	}

	if listManagedKeysFromKeystoreOptions.Algorithm != nil {
		builder.AddQuery("algorithm", strings.Join(listManagedKeysFromKeystoreOptions.Algorithm, ","))
	}
	if listManagedKeysFromKeystoreOptions.State != nil {
		builder.AddQuery("state", strings.Join(listManagedKeysFromKeystoreOptions.State, ","))
	}
	if listManagedKeysFromKeystoreOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listManagedKeysFromKeystoreOptions.Limit))
	}
	if listManagedKeysFromKeystoreOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listManagedKeysFromKeystoreOptions.Offset))
	}
	if listManagedKeysFromKeystoreOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listManagedKeysFromKeystoreOptions.Sort, ","))
	}
	if listManagedKeysFromKeystoreOptions.Label != nil {
		builder.AddQuery("label", fmt.Sprint(*listManagedKeysFromKeystoreOptions.Label))
	}
	if listManagedKeysFromKeystoreOptions.ActivationDate != nil {
		builder.AddQuery("activation_date", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ActivationDate))
	}
	if listManagedKeysFromKeystoreOptions.ActivationDateMin != nil {
		builder.AddQuery("activation_date_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ActivationDateMin))
	}
	if listManagedKeysFromKeystoreOptions.ActivationDateMax != nil {
		builder.AddQuery("activation_date_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ActivationDateMax))
	}
	if listManagedKeysFromKeystoreOptions.DeactivationDate != nil {
		builder.AddQuery("deactivation_date", fmt.Sprint(*listManagedKeysFromKeystoreOptions.DeactivationDate))
	}
	if listManagedKeysFromKeystoreOptions.DeactivationDateMin != nil {
		builder.AddQuery("deactivation_date_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.DeactivationDateMin))
	}
	if listManagedKeysFromKeystoreOptions.DeactivationDateMax != nil {
		builder.AddQuery("deactivation_date_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.DeactivationDateMax))
	}
	if listManagedKeysFromKeystoreOptions.ExpirationDate != nil {
		builder.AddQuery("expiration_date", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ExpirationDate))
	}
	if listManagedKeysFromKeystoreOptions.ExpirationDateMin != nil {
		builder.AddQuery("expiration_date_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ExpirationDateMin))
	}
	if listManagedKeysFromKeystoreOptions.ExpirationDateMax != nil {
		builder.AddQuery("expiration_date_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.ExpirationDateMax))
	}
	if listManagedKeysFromKeystoreOptions.CreatedAt != nil {
		builder.AddQuery("created_at", fmt.Sprint(*listManagedKeysFromKeystoreOptions.CreatedAt))
	}
	if listManagedKeysFromKeystoreOptions.CreatedAtMin != nil {
		builder.AddQuery("created_at_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.CreatedAtMin))
	}
	if listManagedKeysFromKeystoreOptions.CreatedAtMax != nil {
		builder.AddQuery("created_at_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.CreatedAtMax))
	}
	if listManagedKeysFromKeystoreOptions.UpdatedAt != nil {
		builder.AddQuery("updated_at", fmt.Sprint(*listManagedKeysFromKeystoreOptions.UpdatedAt))
	}
	if listManagedKeysFromKeystoreOptions.UpdatedAtMin != nil {
		builder.AddQuery("updated_at_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.UpdatedAtMin))
	}
	if listManagedKeysFromKeystoreOptions.UpdatedAtMax != nil {
		builder.AddQuery("updated_at_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.UpdatedAtMax))
	}
	if listManagedKeysFromKeystoreOptions.RotatedAtMin != nil {
		builder.AddQuery("rotated_at_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.RotatedAtMin))
	}
	if listManagedKeysFromKeystoreOptions.RotatedAtMax != nil {
		builder.AddQuery("rotated_at_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.RotatedAtMax))
	}
	if listManagedKeysFromKeystoreOptions.Size != nil {
		builder.AddQuery("size", fmt.Sprint(*listManagedKeysFromKeystoreOptions.Size))
	}
	if listManagedKeysFromKeystoreOptions.SizeMin != nil {
		builder.AddQuery("size_min", fmt.Sprint(*listManagedKeysFromKeystoreOptions.SizeMin))
	}
	if listManagedKeysFromKeystoreOptions.SizeMax != nil {
		builder.AddQuery("size_max", fmt.Sprint(*listManagedKeysFromKeystoreOptions.SizeMax))
	}
	if listManagedKeysFromKeystoreOptions.TemplateName != nil {
		builder.AddQuery("template.name", fmt.Sprint(*listManagedKeysFromKeystoreOptions.TemplateName))
	}
	if listManagedKeysFromKeystoreOptions.TemplateID != nil {
		builder.AddQuery("template.id", strings.Join(listManagedKeysFromKeystoreOptions.TemplateID, ","))
	}
	if listManagedKeysFromKeystoreOptions.TemplateType != nil {
		builder.AddQuery("template.type[]", strings.Join(listManagedKeysFromKeystoreOptions.TemplateType, ","))
	}
	if listManagedKeysFromKeystoreOptions.StatusInKeystoresKeystoreSyncFlag != nil {
		builder.AddQuery("status_in_keystores[].keystore_sync_flag", strings.Join(listManagedKeysFromKeystoreOptions.StatusInKeystoresKeystoreSyncFlag, ","))
	}
	if listManagedKeysFromKeystoreOptions.TemplateAlignmentStatus != nil {
		builder.AddQuery("template.alignment_status", fmt.Sprint(*listManagedKeysFromKeystoreOptions.TemplateAlignmentStatus))
	}
	if listManagedKeysFromKeystoreOptions.KeyMaterialPresent != nil {
		builder.AddQuery("key_material_present", strings.Join(listManagedKeysFromKeystoreOptions.KeyMaterialPresent, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalManagedKeyList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListVaults : List all vaults
// List all vaults in the instance.
func (uko *UkoV4) ListVaults(listVaultsOptions *ListVaultsOptions) (result *VaultList, response *core.DetailedResponse, err error) {
	return uko.ListVaultsWithContext(context.Background(), listVaultsOptions)
}

// ListVaultsWithContext is an alternate form of the ListVaults method which supports a Context parameter
func (uko *UkoV4) ListVaultsWithContext(ctx context.Context, listVaultsOptions *ListVaultsOptions) (result *VaultList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVaultsOptions, "listVaultsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/vaults`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVaultsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ListVaults")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listVaultsOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listVaultsOptions.Accept))
	}

	if listVaultsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listVaultsOptions.Limit))
	}
	if listVaultsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listVaultsOptions.Offset))
	}
	if listVaultsOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(listVaultsOptions.Sort, ","))
	}
	if listVaultsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listVaultsOptions.Name))
	}
	if listVaultsOptions.Description != nil {
		builder.AddQuery("description", fmt.Sprint(*listVaultsOptions.Description))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVaultList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateVault : Create a vault
// Create a new vault in the instance with the specified name and description.
func (uko *UkoV4) CreateVault(createVaultOptions *CreateVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	return uko.CreateVaultWithContext(context.Background(), createVaultOptions)
}

// CreateVaultWithContext is an alternate form of the CreateVault method which supports a Context parameter
func (uko *UkoV4) CreateVaultWithContext(ctx context.Context, createVaultOptions *CreateVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVaultOptions, "createVaultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createVaultOptions, "createVaultOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/vaults`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createVaultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "CreateVault")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVaultOptions.Name != nil {
		body["name"] = createVaultOptions.Name
	}
	if createVaultOptions.Description != nil {
		body["description"] = createVaultOptions.Description
	}
	if createVaultOptions.RecoveryKeyLabel != nil {
		body["recovery_key_label"] = createVaultOptions.RecoveryKeyLabel
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVault)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteVault : Delete an existing vault
// Delete an existing vault from the system. A vault must be empty (that is, no managed keys or keystores remain in the
// vault) before the vault can be deleted.
func (uko *UkoV4) DeleteVault(deleteVaultOptions *DeleteVaultOptions) (response *core.DetailedResponse, err error) {
	return uko.DeleteVaultWithContext(context.Background(), deleteVaultOptions)
}

// DeleteVaultWithContext is an alternate form of the DeleteVault method which supports a Context parameter
func (uko *UkoV4) DeleteVaultWithContext(ctx context.Context, deleteVaultOptions *DeleteVaultOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVaultOptions, "deleteVaultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteVaultOptions, "deleteVaultOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteVaultOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/vaults/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteVaultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "DeleteVault")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteVaultOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*deleteVaultOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = uko.Service.Request(request, nil)

	return
}

// GetVault : Retrieve a vault
// Retrieve a vault and its details by specifying the ID.
func (uko *UkoV4) GetVault(getVaultOptions *GetVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	return uko.GetVaultWithContext(context.Background(), getVaultOptions)
}

// GetVaultWithContext is an alternate form of the GetVault method which supports a Context parameter
func (uko *UkoV4) GetVaultWithContext(ctx context.Context, getVaultOptions *GetVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVaultOptions, "getVaultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVaultOptions, "getVaultOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getVaultOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/vaults/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVaultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "GetVault")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVault)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateVault : Update a vault
// Updates attributes of a vault.
func (uko *UkoV4) UpdateVault(updateVaultOptions *UpdateVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	return uko.UpdateVaultWithContext(context.Background(), updateVaultOptions)
}

// UpdateVaultWithContext is an alternate form of the UpdateVault method which supports a Context parameter
func (uko *UkoV4) UpdateVaultWithContext(ctx context.Context, updateVaultOptions *UpdateVaultOptions) (result *Vault, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateVaultOptions, "updateVaultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateVaultOptions, "updateVaultOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateVaultOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/vaults/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateVaultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UpdateVault")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateVaultOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateVaultOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateVaultOptions.Name != nil {
		body["name"] = updateVaultOptions.Name
	}
	if updateVaultOptions.Description != nil {
		body["description"] = updateVaultOptions.Description
	}
	if updateVaultOptions.RecoveryKeyLabel != nil {
		body["recovery_key_label"] = updateVaultOptions.RecoveryKeyLabel
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVault)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UnarchiveKeyTemplate : Unarchive a key template
// Unarchive a key template.
func (uko *UkoV4) UnarchiveKeyTemplate(unarchiveKeyTemplateOptions *UnarchiveKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.UnarchiveKeyTemplateWithContext(context.Background(), unarchiveKeyTemplateOptions)
}

// UnarchiveKeyTemplateWithContext is an alternate form of the UnarchiveKeyTemplate method which supports a Context parameter
func (uko *UkoV4) UnarchiveKeyTemplateWithContext(ctx context.Context, unarchiveKeyTemplateOptions *UnarchiveKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unarchiveKeyTemplateOptions, "unarchiveKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unarchiveKeyTemplateOptions, "unarchiveKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *unarchiveKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}/unarchive`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range unarchiveKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "UnarchiveKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if unarchiveKeyTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*unarchiveKeyTemplateOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ArchiveKeyTemplate : Archive a key template
// Archive a key template.
func (uko *UkoV4) ArchiveKeyTemplate(archiveKeyTemplateOptions *ArchiveKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.ArchiveKeyTemplateWithContext(context.Background(), archiveKeyTemplateOptions)
}

// ArchiveKeyTemplateWithContext is an alternate form of the ArchiveKeyTemplate method which supports a Context parameter
func (uko *UkoV4) ArchiveKeyTemplateWithContext(ctx context.Context, archiveKeyTemplateOptions *ArchiveKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(archiveKeyTemplateOptions, "archiveKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(archiveKeyTemplateOptions, "archiveKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *archiveKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}/archive`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range archiveKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ArchiveKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if archiveKeyTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*archiveKeyTemplateOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExposeKeyTemplate : Expose a key template that was previously a shadow template
// Expose a key template.
func (uko *UkoV4) ExposeKeyTemplate(exposeKeyTemplateOptions *ExposeKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	return uko.ExposeKeyTemplateWithContext(context.Background(), exposeKeyTemplateOptions)
}

// ExposeKeyTemplateWithContext is an alternate form of the ExposeKeyTemplate method which supports a Context parameter
func (uko *UkoV4) ExposeKeyTemplateWithContext(ctx context.Context, exposeKeyTemplateOptions *ExposeKeyTemplateOptions) (result *Template, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(exposeKeyTemplateOptions, "exposeKeyTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(exposeKeyTemplateOptions, "exposeKeyTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *exposeKeyTemplateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = uko.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(uko.Service.Options.URL, `/api/v4/templates/{id}/expose`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range exposeKeyTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("uko", "V4", "ExposeKeyTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if exposeKeyTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*exposeKeyTemplateOptions.IfMatch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = uko.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ActivateManagedKeyOptions : The ActivateManagedKey options.
type ActivateManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewActivateManagedKeyOptions : Instantiate ActivateManagedKeyOptions
func (*UkoV4) NewActivateManagedKeyOptions(id string, ifMatch string) *ActivateManagedKeyOptions {
	return &ActivateManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *ActivateManagedKeyOptions) SetID(id string) *ActivateManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *ActivateManagedKeyOptions) SetIfMatch(ifMatch string) *ActivateManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ActivateManagedKeyOptions) SetHeaders(param map[string]string) *ActivateManagedKeyOptions {
	options.Headers = param
	return options
}

// ArchiveKeyTemplateOptions : The ArchiveKeyTemplate options.
type ArchiveKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewArchiveKeyTemplateOptions : Instantiate ArchiveKeyTemplateOptions
func (*UkoV4) NewArchiveKeyTemplateOptions(id string, ifMatch string) *ArchiveKeyTemplateOptions {
	return &ArchiveKeyTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *ArchiveKeyTemplateOptions) SetID(id string) *ArchiveKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *ArchiveKeyTemplateOptions) SetIfMatch(ifMatch string) *ArchiveKeyTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ArchiveKeyTemplateOptions) SetHeaders(param map[string]string) *ArchiveKeyTemplateOptions {
	options.Headers = param
	return options
}

// CreateKeyTemplateOptions : The CreateKeyTemplate options.
type CreateKeyTemplateOptions struct {
	// ID of the Vault where the entity is to be created in.
	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// A human-readable name to assign to your template.
	Name *string `json:"name" validate:"required"`

	// Properties describing the properties of the managed key.
	Key *KeyProperties `json:"key" validate:"required"`

	// An array describing the type and group of target keystores the managed key is to be installed in.
	Keystores []KeystoresPropertiesCreateIntf `json:"keystores" validate:"required"`

	// Description of the key template.
	Description *string `json:"description,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	Type []string `json:"type,omitempty"`

	// State of the template which determines if the template is archived or unarchived.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateKeyTemplateOptions.Type property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	CreateKeyTemplateOptions_Type_Shadow = "shadow"
	CreateKeyTemplateOptions_Type_System = "system"
	CreateKeyTemplateOptions_Type_UserDefined = "user_defined"
)

// Constants associated with the CreateKeyTemplateOptions.State property.
// State of the template which determines if the template is archived or unarchived.
const (
	CreateKeyTemplateOptions_State_Archived = "archived"
	CreateKeyTemplateOptions_State_Unarchived = "unarchived"
)

// NewCreateKeyTemplateOptions : Instantiate CreateKeyTemplateOptions
func (*UkoV4) NewCreateKeyTemplateOptions(vault *VaultReferenceInCreationRequest, name string, key *KeyProperties, keystores []KeystoresPropertiesCreateIntf) *CreateKeyTemplateOptions {
	return &CreateKeyTemplateOptions{
		Vault: vault,
		Name: core.StringPtr(name),
		Key: key,
		Keystores: keystores,
	}
}

// SetVault : Allow user to set Vault
func (_options *CreateKeyTemplateOptions) SetVault(vault *VaultReferenceInCreationRequest) *CreateKeyTemplateOptions {
	_options.Vault = vault
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateKeyTemplateOptions) SetName(name string) *CreateKeyTemplateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetKey : Allow user to set Key
func (_options *CreateKeyTemplateOptions) SetKey(key *KeyProperties) *CreateKeyTemplateOptions {
	_options.Key = key
	return _options
}

// SetKeystores : Allow user to set Keystores
func (_options *CreateKeyTemplateOptions) SetKeystores(keystores []KeystoresPropertiesCreateIntf) *CreateKeyTemplateOptions {
	_options.Keystores = keystores
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateKeyTemplateOptions) SetDescription(description string) *CreateKeyTemplateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetNamingScheme : Allow user to set NamingScheme
func (_options *CreateKeyTemplateOptions) SetNamingScheme(namingScheme string) *CreateKeyTemplateOptions {
	_options.NamingScheme = core.StringPtr(namingScheme)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateKeyTemplateOptions) SetType(typeVar []string) *CreateKeyTemplateOptions {
	_options.Type = typeVar
	return _options
}

// SetState : Allow user to set State
func (_options *CreateKeyTemplateOptions) SetState(state string) *CreateKeyTemplateOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateKeyTemplateOptions) SetHeaders(param map[string]string) *CreateKeyTemplateOptions {
	options.Headers = param
	return options
}

// CreateKeystoreOptions : The CreateKeystore options.
type CreateKeystoreOptions struct {
	// Keystore properties to update.
	KeystoreBody KeystoreCreationRequestIntf `json:"keystoreBody" validate:"required"`

	// Do not create/update/delete a resource, only verify and validate if resource can be created/updated/deleted with
	// given request successfully.
	DryRun *bool `json:"dry_run,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateKeystoreOptions : Instantiate CreateKeystoreOptions
func (*UkoV4) NewCreateKeystoreOptions(keystoreBody KeystoreCreationRequestIntf) *CreateKeystoreOptions {
	return &CreateKeystoreOptions{
		KeystoreBody: keystoreBody,
	}
}

// SetKeystoreBody : Allow user to set KeystoreBody
func (_options *CreateKeystoreOptions) SetKeystoreBody(keystoreBody KeystoreCreationRequestIntf) *CreateKeystoreOptions {
	_options.KeystoreBody = keystoreBody
	return _options
}

// SetDryRun : Allow user to set DryRun
func (_options *CreateKeystoreOptions) SetDryRun(dryRun bool) *CreateKeystoreOptions {
	_options.DryRun = core.BoolPtr(dryRun)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateKeystoreOptions) SetHeaders(param map[string]string) *CreateKeystoreOptions {
	options.Headers = param
	return options
}

// CreateManagedKeyOptions : The CreateManagedKey options.
type CreateManagedKeyOptions struct {
	// Name of the key template to use when creating a key.
	TemplateName *string `json:"template_name" validate:"required"`

	// ID of the Vault where the entity is to be created in.
	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// The label of the key.
	Label *string `json:"label,omitempty"`

	// Description of the managed key.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateManagedKeyOptions : Instantiate CreateManagedKeyOptions
func (*UkoV4) NewCreateManagedKeyOptions(templateName string, vault *VaultReferenceInCreationRequest) *CreateManagedKeyOptions {
	return &CreateManagedKeyOptions{
		TemplateName: core.StringPtr(templateName),
		Vault: vault,
	}
}

// SetTemplateName : Allow user to set TemplateName
func (_options *CreateManagedKeyOptions) SetTemplateName(templateName string) *CreateManagedKeyOptions {
	_options.TemplateName = core.StringPtr(templateName)
	return _options
}

// SetVault : Allow user to set Vault
func (_options *CreateManagedKeyOptions) SetVault(vault *VaultReferenceInCreationRequest) *CreateManagedKeyOptions {
	_options.Vault = vault
	return _options
}

// SetLabel : Allow user to set Label
func (_options *CreateManagedKeyOptions) SetLabel(label string) *CreateManagedKeyOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateManagedKeyOptions) SetDescription(description string) *CreateManagedKeyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateManagedKeyOptions) SetHeaders(param map[string]string) *CreateManagedKeyOptions {
	options.Headers = param
	return options
}

// CreateVaultOptions : The CreateVault options.
type CreateVaultOptions struct {
	// A human-readable name to assign to your vault. To protect your privacy, do not use personal data, such as your name
	// or location.
	Name *string `json:"name" validate:"required"`

	// Description of the vault.
	Description *string `json:"description,omitempty"`

	// The label of the recovery key to use for this vault.
	RecoveryKeyLabel *string `json:"recovery_key_label,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateVaultOptions : Instantiate CreateVaultOptions
func (*UkoV4) NewCreateVaultOptions(name string) *CreateVaultOptions {
	return &CreateVaultOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *CreateVaultOptions) SetName(name string) *CreateVaultOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateVaultOptions) SetDescription(description string) *CreateVaultOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetRecoveryKeyLabel : Allow user to set RecoveryKeyLabel
func (_options *CreateVaultOptions) SetRecoveryKeyLabel(recoveryKeyLabel string) *CreateVaultOptions {
	_options.RecoveryKeyLabel = core.StringPtr(recoveryKeyLabel)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVaultOptions) SetHeaders(param map[string]string) *CreateVaultOptions {
	options.Headers = param
	return options
}

// DeactivateManagedKeyOptions : The DeactivateManagedKey options.
type DeactivateManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeactivateManagedKeyOptions : Instantiate DeactivateManagedKeyOptions
func (*UkoV4) NewDeactivateManagedKeyOptions(id string, ifMatch string) *DeactivateManagedKeyOptions {
	return &DeactivateManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DeactivateManagedKeyOptions) SetID(id string) *DeactivateManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DeactivateManagedKeyOptions) SetIfMatch(ifMatch string) *DeactivateManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeactivateManagedKeyOptions) SetHeaders(param map[string]string) *DeactivateManagedKeyOptions {
	options.Headers = param
	return options
}

// DeleteKeyTemplateOptions : The DeleteKeyTemplate options.
type DeleteKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteKeyTemplateOptions : Instantiate DeleteKeyTemplateOptions
func (*UkoV4) NewDeleteKeyTemplateOptions(id string, ifMatch string) *DeleteKeyTemplateOptions {
	return &DeleteKeyTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteKeyTemplateOptions) SetID(id string) *DeleteKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DeleteKeyTemplateOptions) SetIfMatch(ifMatch string) *DeleteKeyTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteKeyTemplateOptions) SetHeaders(param map[string]string) *DeleteKeyTemplateOptions {
	options.Headers = param
	return options
}

// DeleteKeystoreOptions : The DeleteKeystore options.
type DeleteKeystoreOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Mode of disconnecting from keystore.
	Mode *string `json:"mode,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteKeystoreOptions.Mode property.
// Mode of disconnecting from keystore.
const (
	DeleteKeystoreOptions_Mode_Deactivate = "deactivate"
	DeleteKeystoreOptions_Mode_Destroy = "destroy"
	DeleteKeystoreOptions_Mode_Disconnect = "disconnect"
	DeleteKeystoreOptions_Mode_Restrict = "restrict"
)

// NewDeleteKeystoreOptions : Instantiate DeleteKeystoreOptions
func (*UkoV4) NewDeleteKeystoreOptions(id string, ifMatch string) *DeleteKeystoreOptions {
	return &DeleteKeystoreOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteKeystoreOptions) SetID(id string) *DeleteKeystoreOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DeleteKeystoreOptions) SetIfMatch(ifMatch string) *DeleteKeystoreOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetMode : Allow user to set Mode
func (_options *DeleteKeystoreOptions) SetMode(mode string) *DeleteKeystoreOptions {
	_options.Mode = core.StringPtr(mode)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteKeystoreOptions) SetHeaders(param map[string]string) *DeleteKeystoreOptions {
	options.Headers = param
	return options
}

// DeleteManagedKeyOptions : The DeleteManagedKey options.
type DeleteManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Remove Managed Key from UKO even if it's still pending destruction.
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteManagedKeyOptions : Instantiate DeleteManagedKeyOptions
func (*UkoV4) NewDeleteManagedKeyOptions(id string, ifMatch string) *DeleteManagedKeyOptions {
	return &DeleteManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteManagedKeyOptions) SetID(id string) *DeleteManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DeleteManagedKeyOptions) SetIfMatch(ifMatch string) *DeleteManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetForce : Allow user to set Force
func (_options *DeleteManagedKeyOptions) SetForce(force bool) *DeleteManagedKeyOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteManagedKeyOptions) SetHeaders(param map[string]string) *DeleteManagedKeyOptions {
	options.Headers = param
	return options
}

// DeleteVaultOptions : The DeleteVault options.
type DeleteVaultOptions struct {
	// UUID of the vault.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteVaultOptions : Instantiate DeleteVaultOptions
func (*UkoV4) NewDeleteVaultOptions(id string, ifMatch string) *DeleteVaultOptions {
	return &DeleteVaultOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteVaultOptions) SetID(id string) *DeleteVaultOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DeleteVaultOptions) SetIfMatch(ifMatch string) *DeleteVaultOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVaultOptions) SetHeaders(param map[string]string) *DeleteVaultOptions {
	options.Headers = param
	return options
}

// DestroyManagedKeyOptions : The DestroyManagedKey options.
type DestroyManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDestroyManagedKeyOptions : Instantiate DestroyManagedKeyOptions
func (*UkoV4) NewDestroyManagedKeyOptions(id string, ifMatch string) *DestroyManagedKeyOptions {
	return &DestroyManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *DestroyManagedKeyOptions) SetID(id string) *DestroyManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *DestroyManagedKeyOptions) SetIfMatch(ifMatch string) *DestroyManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DestroyManagedKeyOptions) SetHeaders(param map[string]string) *DestroyManagedKeyOptions {
	options.Headers = param
	return options
}

// ExposeKeyTemplateOptions : The ExposeKeyTemplate options.
type ExposeKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExposeKeyTemplateOptions : Instantiate ExposeKeyTemplateOptions
func (*UkoV4) NewExposeKeyTemplateOptions(id string, ifMatch string) *ExposeKeyTemplateOptions {
	return &ExposeKeyTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *ExposeKeyTemplateOptions) SetID(id string) *ExposeKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *ExposeKeyTemplateOptions) SetIfMatch(ifMatch string) *ExposeKeyTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExposeKeyTemplateOptions) SetHeaders(param map[string]string) *ExposeKeyTemplateOptions {
	options.Headers = param
	return options
}

// GetKeyDistributionStatusForKeystoresOptions : The GetKeyDistributionStatusForKeystores options.
type GetKeyDistributionStatusForKeystoresOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetKeyDistributionStatusForKeystoresOptions : Instantiate GetKeyDistributionStatusForKeystoresOptions
func (*UkoV4) NewGetKeyDistributionStatusForKeystoresOptions(id string) *GetKeyDistributionStatusForKeystoresOptions {
	return &GetKeyDistributionStatusForKeystoresOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetKeyDistributionStatusForKeystoresOptions) SetID(id string) *GetKeyDistributionStatusForKeystoresOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetKeyDistributionStatusForKeystoresOptions) SetHeaders(param map[string]string) *GetKeyDistributionStatusForKeystoresOptions {
	options.Headers = param
	return options
}

// GetKeyTemplateOptions : The GetKeyTemplate options.
type GetKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetKeyTemplateOptions : Instantiate GetKeyTemplateOptions
func (*UkoV4) NewGetKeyTemplateOptions(id string) *GetKeyTemplateOptions {
	return &GetKeyTemplateOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetKeyTemplateOptions) SetID(id string) *GetKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetKeyTemplateOptions) SetHeaders(param map[string]string) *GetKeyTemplateOptions {
	options.Headers = param
	return options
}

// GetKeystoreOptions : The GetKeystore options.
type GetKeystoreOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetKeystoreOptions : Instantiate GetKeystoreOptions
func (*UkoV4) NewGetKeystoreOptions(id string) *GetKeystoreOptions {
	return &GetKeystoreOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetKeystoreOptions) SetID(id string) *GetKeystoreOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetKeystoreOptions) SetHeaders(param map[string]string) *GetKeystoreOptions {
	options.Headers = param
	return options
}

// GetKeystoreStatusOptions : The GetKeystoreStatus options.
type GetKeystoreStatusOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetKeystoreStatusOptions : Instantiate GetKeystoreStatusOptions
func (*UkoV4) NewGetKeystoreStatusOptions(id string) *GetKeystoreStatusOptions {
	return &GetKeystoreStatusOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetKeystoreStatusOptions) SetID(id string) *GetKeystoreStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetKeystoreStatusOptions) SetHeaders(param map[string]string) *GetKeystoreStatusOptions {
	options.Headers = param
	return options
}

// GetManagedKeyOptions : The GetManagedKey options.
type GetManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetManagedKeyOptions : Instantiate GetManagedKeyOptions
func (*UkoV4) NewGetManagedKeyOptions(id string) *GetManagedKeyOptions {
	return &GetManagedKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetManagedKeyOptions) SetID(id string) *GetManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetManagedKeyOptions) SetHeaders(param map[string]string) *GetManagedKeyOptions {
	options.Headers = param
	return options
}

// GetVaultOptions : The GetVault options.
type GetVaultOptions struct {
	// UUID of the vault.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVaultOptions : Instantiate GetVaultOptions
func (*UkoV4) NewGetVaultOptions(id string) *GetVaultOptions {
	return &GetVaultOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetVaultOptions) SetID(id string) *GetVaultOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVaultOptions) SetHeaders(param map[string]string) *GetVaultOptions {
	options.Headers = param
	return options
}

// ListAssociatedResourcesForManagedKeyOptions : The ListAssociatedResourcesForManagedKey options.
type ListAssociatedResourcesForManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAssociatedResourcesForManagedKeyOptions : Instantiate ListAssociatedResourcesForManagedKeyOptions
func (*UkoV4) NewListAssociatedResourcesForManagedKeyOptions(id string) *ListAssociatedResourcesForManagedKeyOptions {
	return &ListAssociatedResourcesForManagedKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListAssociatedResourcesForManagedKeyOptions) SetID(id string) *ListAssociatedResourcesForManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAssociatedResourcesForManagedKeyOptions) SetLimit(limit int64) *ListAssociatedResourcesForManagedKeyOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListAssociatedResourcesForManagedKeyOptions) SetOffset(offset int64) *ListAssociatedResourcesForManagedKeyOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListAssociatedResourcesForManagedKeyOptions) SetSort(sort []string) *ListAssociatedResourcesForManagedKeyOptions {
	_options.Sort = sort
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAssociatedResourcesForManagedKeyOptions) SetHeaders(param map[string]string) *ListAssociatedResourcesForManagedKeyOptions {
	options.Headers = param
	return options
}

// ListAssociatedResourcesForTargetKeystoreOptions : The ListAssociatedResourcesForTargetKeystore options.
type ListAssociatedResourcesForTargetKeystoreOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAssociatedResourcesForTargetKeystoreOptions : Instantiate ListAssociatedResourcesForTargetKeystoreOptions
func (*UkoV4) NewListAssociatedResourcesForTargetKeystoreOptions(id string) *ListAssociatedResourcesForTargetKeystoreOptions {
	return &ListAssociatedResourcesForTargetKeystoreOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListAssociatedResourcesForTargetKeystoreOptions) SetID(id string) *ListAssociatedResourcesForTargetKeystoreOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAssociatedResourcesForTargetKeystoreOptions) SetLimit(limit int64) *ListAssociatedResourcesForTargetKeystoreOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListAssociatedResourcesForTargetKeystoreOptions) SetOffset(offset int64) *ListAssociatedResourcesForTargetKeystoreOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListAssociatedResourcesForTargetKeystoreOptions) SetSort(sort []string) *ListAssociatedResourcesForTargetKeystoreOptions {
	_options.Sort = sort
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAssociatedResourcesForTargetKeystoreOptions) SetHeaders(param map[string]string) *ListAssociatedResourcesForTargetKeystoreOptions {
	options.Headers = param
	return options
}

// ListKeyTemplatesOptions : The ListKeyTemplates options.
type ListKeyTemplatesOptions struct {
	// The type of the response: application/json, application/vnd.ibm.uko.key-template-list.v4.1+json,
	// application/vnd.ibm.uko.key-template-list.v4.1.json+zip, or application/vnd.ibm.uko.key-template-list.v4.1.csv+zip.
	Accept *string `json:"Accept,omitempty"`

	// Return only templates whose name begin with the string.
	Name *string `json:"name,omitempty"`

	// Return only templates whose naming scheme contains the string.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// The UUID of the Vault.
	VaultID []string `json:"vault.id,omitempty"`

	// The algorithm of a returned key template.
	KeyAlgorithm []string `json:"key.algorithm,omitempty"`

	// The size of the key.
	KeySize *string `json:"key.size,omitempty"`

	// The minimum size of the key. This query parameter cannot be used in conjunction with the 'key.size' query parameter.
	KeySizeMin *string `json:"key.size_min,omitempty"`

	// The maximum size of the key. This query parameter cannot be used in conjunction with the 'key.size' query parameter.
	KeySizeMax *string `json:"key.size_max,omitempty"`

	// Type of referenced keystore.
	KeystoresType []string `json:"keystores[].type,omitempty"`

	// Group of referenced keystore.
	KeystoresGroup []string `json:"keystores[].group,omitempty"`

	// Return only managed keys whose created_at matches the parameter.
	CreatedAt *string `json:"created_at,omitempty"`

	// Return only managed keys whose created_at is at or after the parameter value. This query parameter cannot be used in
	// conjunction with the 'created_at' query parameter.
	CreatedAtMin *string `json:"created_at_min,omitempty"`

	// Return only managed keys whose created_at is at or before the parameter value. This query parameter cannot be used
	// in conjunction with the 'created_at' query parameter.
	CreatedAtMax *string `json:"created_at_max,omitempty"`

	// Return only managed keys whose updated_at matches the parameter.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Return only managed keys whose updated_at is after the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMin *string `json:"updated_at_min,omitempty"`

	// Return only managed keys whose updated_at is before the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMax *string `json:"updated_at_max,omitempty"`

	// The types of returned templates.
	Type []string `json:"type[],omitempty"`

	// Return only template whose state contains the string.
	State []string `json:"state,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Return only managed keys with the given managing systems.
	ManagingSystems []string `json:"managing_systems,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListKeyTemplatesOptions.KeyAlgorithm property.
// The algorithm of the key.
const (
	ListKeyTemplatesOptions_KeyAlgorithm_Aes = "aes"
	ListKeyTemplatesOptions_KeyAlgorithm_Des = "des"
	ListKeyTemplatesOptions_KeyAlgorithm_Dilithium = "dilithium"
	ListKeyTemplatesOptions_KeyAlgorithm_Ec = "ec"
	ListKeyTemplatesOptions_KeyAlgorithm_Hmac = "hmac"
	ListKeyTemplatesOptions_KeyAlgorithm_Rsa = "rsa"
)

// Constants associated with the ListKeyTemplatesOptions.KeystoresType property.
// Type of keystore.
const (
	ListKeyTemplatesOptions_KeystoresType_AwsKms = "aws_kms"
	ListKeyTemplatesOptions_KeystoresType_AzureKeyVault = "azure_key_vault"
	ListKeyTemplatesOptions_KeystoresType_Cca = "cca"
	ListKeyTemplatesOptions_KeystoresType_GoogleKms = "google_kms"
	ListKeyTemplatesOptions_KeystoresType_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListKeyTemplatesOptions.Type property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	ListKeyTemplatesOptions_Type_Shadow = "shadow"
	ListKeyTemplatesOptions_Type_System = "system"
	ListKeyTemplatesOptions_Type_UserDefined = "user_defined"
)

// Constants associated with the ListKeyTemplatesOptions.State property.
// State of the template which determines if the template is archived or unarchived.
const (
	ListKeyTemplatesOptions_State_Archived = "archived"
	ListKeyTemplatesOptions_State_Unarchived = "unarchived"
)

// Constants associated with the ListKeyTemplatesOptions.ManagingSystems property.
// Managing system of templates and keys.
const (
	ListKeyTemplatesOptions_ManagingSystems_Web = "web"
	ListKeyTemplatesOptions_ManagingSystems_Workstation = "workstation"
)

// NewListKeyTemplatesOptions : Instantiate ListKeyTemplatesOptions
func (*UkoV4) NewListKeyTemplatesOptions() *ListKeyTemplatesOptions {
	return &ListKeyTemplatesOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *ListKeyTemplatesOptions) SetAccept(accept string) *ListKeyTemplatesOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListKeyTemplatesOptions) SetName(name string) *ListKeyTemplatesOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetNamingScheme : Allow user to set NamingScheme
func (_options *ListKeyTemplatesOptions) SetNamingScheme(namingScheme string) *ListKeyTemplatesOptions {
	_options.NamingScheme = core.StringPtr(namingScheme)
	return _options
}

// SetVaultID : Allow user to set VaultID
func (_options *ListKeyTemplatesOptions) SetVaultID(vaultID []string) *ListKeyTemplatesOptions {
	_options.VaultID = vaultID
	return _options
}

// SetKeyAlgorithm : Allow user to set KeyAlgorithm
func (_options *ListKeyTemplatesOptions) SetKeyAlgorithm(keyAlgorithm []string) *ListKeyTemplatesOptions {
	_options.KeyAlgorithm = keyAlgorithm
	return _options
}

// SetKeySize : Allow user to set KeySize
func (_options *ListKeyTemplatesOptions) SetKeySize(keySize string) *ListKeyTemplatesOptions {
	_options.KeySize = core.StringPtr(keySize)
	return _options
}

// SetKeySizeMin : Allow user to set KeySizeMin
func (_options *ListKeyTemplatesOptions) SetKeySizeMin(keySizeMin string) *ListKeyTemplatesOptions {
	_options.KeySizeMin = core.StringPtr(keySizeMin)
	return _options
}

// SetKeySizeMax : Allow user to set KeySizeMax
func (_options *ListKeyTemplatesOptions) SetKeySizeMax(keySizeMax string) *ListKeyTemplatesOptions {
	_options.KeySizeMax = core.StringPtr(keySizeMax)
	return _options
}

// SetKeystoresType : Allow user to set KeystoresType
func (_options *ListKeyTemplatesOptions) SetKeystoresType(keystoresType []string) *ListKeyTemplatesOptions {
	_options.KeystoresType = keystoresType
	return _options
}

// SetKeystoresGroup : Allow user to set KeystoresGroup
func (_options *ListKeyTemplatesOptions) SetKeystoresGroup(keystoresGroup []string) *ListKeyTemplatesOptions {
	_options.KeystoresGroup = keystoresGroup
	return _options
}

// SetCreatedAt : Allow user to set CreatedAt
func (_options *ListKeyTemplatesOptions) SetCreatedAt(createdAt string) *ListKeyTemplatesOptions {
	_options.CreatedAt = core.StringPtr(createdAt)
	return _options
}

// SetCreatedAtMin : Allow user to set CreatedAtMin
func (_options *ListKeyTemplatesOptions) SetCreatedAtMin(createdAtMin string) *ListKeyTemplatesOptions {
	_options.CreatedAtMin = core.StringPtr(createdAtMin)
	return _options
}

// SetCreatedAtMax : Allow user to set CreatedAtMax
func (_options *ListKeyTemplatesOptions) SetCreatedAtMax(createdAtMax string) *ListKeyTemplatesOptions {
	_options.CreatedAtMax = core.StringPtr(createdAtMax)
	return _options
}

// SetUpdatedAt : Allow user to set UpdatedAt
func (_options *ListKeyTemplatesOptions) SetUpdatedAt(updatedAt string) *ListKeyTemplatesOptions {
	_options.UpdatedAt = core.StringPtr(updatedAt)
	return _options
}

// SetUpdatedAtMin : Allow user to set UpdatedAtMin
func (_options *ListKeyTemplatesOptions) SetUpdatedAtMin(updatedAtMin string) *ListKeyTemplatesOptions {
	_options.UpdatedAtMin = core.StringPtr(updatedAtMin)
	return _options
}

// SetUpdatedAtMax : Allow user to set UpdatedAtMax
func (_options *ListKeyTemplatesOptions) SetUpdatedAtMax(updatedAtMax string) *ListKeyTemplatesOptions {
	_options.UpdatedAtMax = core.StringPtr(updatedAtMax)
	return _options
}

// SetType : Allow user to set Type
func (_options *ListKeyTemplatesOptions) SetType(typeVar []string) *ListKeyTemplatesOptions {
	_options.Type = typeVar
	return _options
}

// SetState : Allow user to set State
func (_options *ListKeyTemplatesOptions) SetState(state []string) *ListKeyTemplatesOptions {
	_options.State = state
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListKeyTemplatesOptions) SetSort(sort []string) *ListKeyTemplatesOptions {
	_options.Sort = sort
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListKeyTemplatesOptions) SetLimit(limit int64) *ListKeyTemplatesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListKeyTemplatesOptions) SetOffset(offset int64) *ListKeyTemplatesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetManagingSystems : Allow user to set ManagingSystems
func (_options *ListKeyTemplatesOptions) SetManagingSystems(managingSystems []string) *ListKeyTemplatesOptions {
	_options.ManagingSystems = managingSystems
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListKeyTemplatesOptions) SetHeaders(param map[string]string) *ListKeyTemplatesOptions {
	options.Headers = param
	return options
}

// ListKeystoresOptions : The ListKeystores options.
type ListKeystoresOptions struct {
	// The type of the response: application/json, application/vnd.ibm.uko.keystore-list.v4.1+json,
	// application/vnd.ibm.uko.keystore-list.v4.1.json+zip, or application/vnd.ibm.uko.keystore-list.v4.1.csv+zip.
	Accept *string `json:"Accept,omitempty"`

	// Keystore type.
	Type []string `json:"type,omitempty"`

	// Return only keystores whose name contains the string.
	Name *string `json:"name,omitempty"`

	// Return only keystores whose description contains the string.
	Description *string `json:"description,omitempty"`

	// A Keystore group. This query parameter cannot be used in conjunction with the 'groups[]' query parameter.
	Group *string `json:"group,omitempty"`

	// Keystore groups.
	Groups *string `json:"groups[],omitempty"`

	// The UUID of the Vault.
	VaultID []string `json:"vault.id,omitempty"`

	// Keystore location.
	Location []string `json:"location,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// Keystore Health status.
	StatusHealthStatus []string `json:"status.health_status,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListKeystoresOptions.Type property.
// Type of keystore.
const (
	ListKeystoresOptions_Type_AwsKms = "aws_kms"
	ListKeystoresOptions_Type_AzureKeyVault = "azure_key_vault"
	ListKeystoresOptions_Type_Cca = "cca"
	ListKeystoresOptions_Type_GoogleKms = "google_kms"
	ListKeystoresOptions_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListKeystoresOptions.StatusHealthStatus property.
// Possible states of a keystore.
const (
	ListKeystoresOptions_StatusHealthStatus_ConfigurationError = "configuration_error"
	ListKeystoresOptions_StatusHealthStatus_NotResponding = "not_responding"
	ListKeystoresOptions_StatusHealthStatus_Ok = "ok"
	ListKeystoresOptions_StatusHealthStatus_PendingCheck = "pending_check"
)

// NewListKeystoresOptions : Instantiate ListKeystoresOptions
func (*UkoV4) NewListKeystoresOptions() *ListKeystoresOptions {
	return &ListKeystoresOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *ListKeystoresOptions) SetAccept(accept string) *ListKeystoresOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetType : Allow user to set Type
func (_options *ListKeystoresOptions) SetType(typeVar []string) *ListKeystoresOptions {
	_options.Type = typeVar
	return _options
}

// SetName : Allow user to set Name
func (_options *ListKeystoresOptions) SetName(name string) *ListKeystoresOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ListKeystoresOptions) SetDescription(description string) *ListKeystoresOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetGroup : Allow user to set Group
func (_options *ListKeystoresOptions) SetGroup(group string) *ListKeystoresOptions {
	_options.Group = core.StringPtr(group)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *ListKeystoresOptions) SetGroups(groups string) *ListKeystoresOptions {
	_options.Groups = core.StringPtr(groups)
	return _options
}

// SetVaultID : Allow user to set VaultID
func (_options *ListKeystoresOptions) SetVaultID(vaultID []string) *ListKeystoresOptions {
	_options.VaultID = vaultID
	return _options
}

// SetLocation : Allow user to set Location
func (_options *ListKeystoresOptions) SetLocation(location []string) *ListKeystoresOptions {
	_options.Location = location
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListKeystoresOptions) SetLimit(limit int64) *ListKeystoresOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListKeystoresOptions) SetOffset(offset int64) *ListKeystoresOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListKeystoresOptions) SetSort(sort []string) *ListKeystoresOptions {
	_options.Sort = sort
	return _options
}

// SetStatusHealthStatus : Allow user to set StatusHealthStatus
func (_options *ListKeystoresOptions) SetStatusHealthStatus(statusHealthStatus []string) *ListKeystoresOptions {
	_options.StatusHealthStatus = statusHealthStatus
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListKeystoresOptions) SetHeaders(param map[string]string) *ListKeystoresOptions {
	options.Headers = param
	return options
}

// ListManagedKeyVersionsOptions : The ListManagedKeyVersions options.
type ListManagedKeyVersionsOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// The algorithm of a returned key.
	Algorithm []string `json:"algorithm,omitempty"`

	// The state that returned keys are to be in.
	State []string `json:"state,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// The label of the key.
	Label *string `json:"label,omitempty"`

	// Return only managed keys whose activation_date matches the parameter.
	ActivationDate *string `json:"activation_date,omitempty"`

	// Return only managed keys whose activation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMin *string `json:"activation_date_min,omitempty"`

	// Return only managed keys whose activation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMax *string `json:"activation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter. This query parameter cannot be used in
	// conjunction with the 'expiration_date' query parameter.
	DeactivationDate *string `json:"deactivation_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMin *string `json:"deactivation_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMax *string `json:"deactivation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter.
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMin *string `json:"expiration_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMax *string `json:"expiration_date_max,omitempty"`

	// Return only managed keys whose created_at matches the parameter.
	CreatedAt *string `json:"created_at,omitempty"`

	// Return only managed keys whose created_at is at or after the parameter value. This query parameter cannot be used in
	// conjunction with the 'created_at' query parameter.
	CreatedAtMin *string `json:"created_at_min,omitempty"`

	// Return only managed keys whose created_at is at or before the parameter value. This query parameter cannot be used
	// in conjunction with the 'created_at' query parameter.
	CreatedAtMax *string `json:"created_at_max,omitempty"`

	// Return only managed keys whose updated_at matches the parameter.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Return only managed keys whose updated_at is after the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMin *string `json:"updated_at_min,omitempty"`

	// Return only managed keys whose updated_at is before the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMax *string `json:"updated_at_max,omitempty"`

	// Return only managed keys whose rotated_at is after the parameter value.
	RotatedAtMin *string `json:"rotated_at_min,omitempty"`

	// Return only managed keys whose rotated_at is before the parameter value.
	RotatedAtMax *string `json:"rotated_at_max,omitempty"`

	// The size of the key.
	Size *string `json:"size,omitempty"`

	// The minimum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMin *string `json:"size_min,omitempty"`

	// The maximum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMax *string `json:"size_max,omitempty"`

	// Type of referenced keystore. This query parameter cannot be used in conjunction with the 'instances[].keystore.type'
	// query parameter.
	ReferencedKeystoresType []string `json:"referenced_keystores[].type,omitempty"`

	// Name of referenced keystore.
	ReferencedKeystoresName []string `json:"referenced_keystores[].name,omitempty"`

	// Type of keystore supported by one of the instances. This query parameter cannot be used in conjunction with the
	// 'referenced_keystores[].type' query parameter.
	InstancesKeystoreType []string `json:"instances[].keystore.type,omitempty"`

	// Return only managed keys whose template name begins with the string.
	TemplateName *string `json:"template.name,omitempty"`

	// Return only managed keys with the given template UUID.
	TemplateID []string `json:"template.id,omitempty"`

	// Return only managed keys with the given template type.
	TemplateType []string `json:"template.type[],omitempty"`

	// Return only Managed Keys whose status_in_keystores contains one of the specified keystore_sync_flag.
	StatusInKeystoresKeystoreSyncFlag []string `json:"status_in_keystores[].keystore_sync_flag,omitempty"`

	// Return only managed keys with the given alignment status.
	TemplateAlignmentStatus *string `json:"template.alignment_status,omitempty"`

	// Return only managed keys which key material is present on given set locations.
	KeyMaterialPresent []string `json:"key_material_present,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListManagedKeyVersionsOptions.Algorithm property.
// The algorithm of the key.
const (
	ListManagedKeyVersionsOptions_Algorithm_Aes = "aes"
	ListManagedKeyVersionsOptions_Algorithm_Des = "des"
	ListManagedKeyVersionsOptions_Algorithm_Dilithium = "dilithium"
	ListManagedKeyVersionsOptions_Algorithm_Ec = "ec"
	ListManagedKeyVersionsOptions_Algorithm_Hmac = "hmac"
	ListManagedKeyVersionsOptions_Algorithm_Rsa = "rsa"
)

// Constants associated with the ListManagedKeyVersionsOptions.State property.
// The state of the key.
const (
	ListManagedKeyVersionsOptions_State_Active = "active"
	ListManagedKeyVersionsOptions_State_Compromised = "compromised"
	ListManagedKeyVersionsOptions_State_Deactivated = "deactivated"
	ListManagedKeyVersionsOptions_State_Destroyed = "destroyed"
	ListManagedKeyVersionsOptions_State_DestroyedCompromised = "destroyed_compromised"
	ListManagedKeyVersionsOptions_State_PreActivation = "pre_activation"
)

// Constants associated with the ListManagedKeyVersionsOptions.ReferencedKeystoresType property.
// Type of keystore.
const (
	ListManagedKeyVersionsOptions_ReferencedKeystoresType_AwsKms = "aws_kms"
	ListManagedKeyVersionsOptions_ReferencedKeystoresType_AzureKeyVault = "azure_key_vault"
	ListManagedKeyVersionsOptions_ReferencedKeystoresType_Cca = "cca"
	ListManagedKeyVersionsOptions_ReferencedKeystoresType_GoogleKms = "google_kms"
	ListManagedKeyVersionsOptions_ReferencedKeystoresType_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListManagedKeyVersionsOptions.InstancesKeystoreType property.
// Type of keystore.
const (
	ListManagedKeyVersionsOptions_InstancesKeystoreType_AwsKms = "aws_kms"
	ListManagedKeyVersionsOptions_InstancesKeystoreType_AzureKeyVault = "azure_key_vault"
	ListManagedKeyVersionsOptions_InstancesKeystoreType_Cca = "cca"
	ListManagedKeyVersionsOptions_InstancesKeystoreType_GoogleKms = "google_kms"
	ListManagedKeyVersionsOptions_InstancesKeystoreType_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListManagedKeyVersionsOptions.TemplateType property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	ListManagedKeyVersionsOptions_TemplateType_Shadow = "shadow"
	ListManagedKeyVersionsOptions_TemplateType_System = "system"
	ListManagedKeyVersionsOptions_TemplateType_UserDefined = "user_defined"
)

// Constants associated with the ListManagedKeyVersionsOptions.StatusInKeystoresKeystoreSyncFlag property.
// Flag to represent synchronization status between UKO Managed Key and Target Keystore. Possible status flags. ok:
// managed key state is the same as target keystore state, out_of_sync: managed key state is different than target
// keystore state.
const (
	ListManagedKeyVersionsOptions_StatusInKeystoresKeystoreSyncFlag_Error = "error"
	ListManagedKeyVersionsOptions_StatusInKeystoresKeystoreSyncFlag_Ok = "ok"
	ListManagedKeyVersionsOptions_StatusInKeystoresKeystoreSyncFlag_OutOfSync = "out_of_sync"
	ListManagedKeyVersionsOptions_StatusInKeystoresKeystoreSyncFlag_VerifyingSync = "verifying_sync"
)

// Constants associated with the ListManagedKeyVersionsOptions.TemplateAlignmentStatus property.
// Return only managed keys with the given alignment status.
const (
	ListManagedKeyVersionsOptions_TemplateAlignmentStatus_Aligned = "aligned"
	ListManagedKeyVersionsOptions_TemplateAlignmentStatus_Unaligned = "unaligned"
)

// Constants associated with the ListManagedKeyVersionsOptions.KeyMaterialPresent property.
// Location where key material is present.
const (
	ListManagedKeyVersionsOptions_KeyMaterialPresent_Keystores = "keystores"
	ListManagedKeyVersionsOptions_KeyMaterialPresent_Repository = "repository"
)

// NewListManagedKeyVersionsOptions : Instantiate ListManagedKeyVersionsOptions
func (*UkoV4) NewListManagedKeyVersionsOptions(id string) *ListManagedKeyVersionsOptions {
	return &ListManagedKeyVersionsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListManagedKeyVersionsOptions) SetID(id string) *ListManagedKeyVersionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetAlgorithm : Allow user to set Algorithm
func (_options *ListManagedKeyVersionsOptions) SetAlgorithm(algorithm []string) *ListManagedKeyVersionsOptions {
	_options.Algorithm = algorithm
	return _options
}

// SetState : Allow user to set State
func (_options *ListManagedKeyVersionsOptions) SetState(state []string) *ListManagedKeyVersionsOptions {
	_options.State = state
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListManagedKeyVersionsOptions) SetLimit(limit int64) *ListManagedKeyVersionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListManagedKeyVersionsOptions) SetOffset(offset int64) *ListManagedKeyVersionsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListManagedKeyVersionsOptions) SetSort(sort []string) *ListManagedKeyVersionsOptions {
	_options.Sort = sort
	return _options
}

// SetLabel : Allow user to set Label
func (_options *ListManagedKeyVersionsOptions) SetLabel(label string) *ListManagedKeyVersionsOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetActivationDate : Allow user to set ActivationDate
func (_options *ListManagedKeyVersionsOptions) SetActivationDate(activationDate string) *ListManagedKeyVersionsOptions {
	_options.ActivationDate = core.StringPtr(activationDate)
	return _options
}

// SetActivationDateMin : Allow user to set ActivationDateMin
func (_options *ListManagedKeyVersionsOptions) SetActivationDateMin(activationDateMin string) *ListManagedKeyVersionsOptions {
	_options.ActivationDateMin = core.StringPtr(activationDateMin)
	return _options
}

// SetActivationDateMax : Allow user to set ActivationDateMax
func (_options *ListManagedKeyVersionsOptions) SetActivationDateMax(activationDateMax string) *ListManagedKeyVersionsOptions {
	_options.ActivationDateMax = core.StringPtr(activationDateMax)
	return _options
}

// SetDeactivationDate : Allow user to set DeactivationDate
func (_options *ListManagedKeyVersionsOptions) SetDeactivationDate(deactivationDate string) *ListManagedKeyVersionsOptions {
	_options.DeactivationDate = core.StringPtr(deactivationDate)
	return _options
}

// SetDeactivationDateMin : Allow user to set DeactivationDateMin
func (_options *ListManagedKeyVersionsOptions) SetDeactivationDateMin(deactivationDateMin string) *ListManagedKeyVersionsOptions {
	_options.DeactivationDateMin = core.StringPtr(deactivationDateMin)
	return _options
}

// SetDeactivationDateMax : Allow user to set DeactivationDateMax
func (_options *ListManagedKeyVersionsOptions) SetDeactivationDateMax(deactivationDateMax string) *ListManagedKeyVersionsOptions {
	_options.DeactivationDateMax = core.StringPtr(deactivationDateMax)
	return _options
}

// SetExpirationDate : Allow user to set ExpirationDate
func (_options *ListManagedKeyVersionsOptions) SetExpirationDate(expirationDate string) *ListManagedKeyVersionsOptions {
	_options.ExpirationDate = core.StringPtr(expirationDate)
	return _options
}

// SetExpirationDateMin : Allow user to set ExpirationDateMin
func (_options *ListManagedKeyVersionsOptions) SetExpirationDateMin(expirationDateMin string) *ListManagedKeyVersionsOptions {
	_options.ExpirationDateMin = core.StringPtr(expirationDateMin)
	return _options
}

// SetExpirationDateMax : Allow user to set ExpirationDateMax
func (_options *ListManagedKeyVersionsOptions) SetExpirationDateMax(expirationDateMax string) *ListManagedKeyVersionsOptions {
	_options.ExpirationDateMax = core.StringPtr(expirationDateMax)
	return _options
}

// SetCreatedAt : Allow user to set CreatedAt
func (_options *ListManagedKeyVersionsOptions) SetCreatedAt(createdAt string) *ListManagedKeyVersionsOptions {
	_options.CreatedAt = core.StringPtr(createdAt)
	return _options
}

// SetCreatedAtMin : Allow user to set CreatedAtMin
func (_options *ListManagedKeyVersionsOptions) SetCreatedAtMin(createdAtMin string) *ListManagedKeyVersionsOptions {
	_options.CreatedAtMin = core.StringPtr(createdAtMin)
	return _options
}

// SetCreatedAtMax : Allow user to set CreatedAtMax
func (_options *ListManagedKeyVersionsOptions) SetCreatedAtMax(createdAtMax string) *ListManagedKeyVersionsOptions {
	_options.CreatedAtMax = core.StringPtr(createdAtMax)
	return _options
}

// SetUpdatedAt : Allow user to set UpdatedAt
func (_options *ListManagedKeyVersionsOptions) SetUpdatedAt(updatedAt string) *ListManagedKeyVersionsOptions {
	_options.UpdatedAt = core.StringPtr(updatedAt)
	return _options
}

// SetUpdatedAtMin : Allow user to set UpdatedAtMin
func (_options *ListManagedKeyVersionsOptions) SetUpdatedAtMin(updatedAtMin string) *ListManagedKeyVersionsOptions {
	_options.UpdatedAtMin = core.StringPtr(updatedAtMin)
	return _options
}

// SetUpdatedAtMax : Allow user to set UpdatedAtMax
func (_options *ListManagedKeyVersionsOptions) SetUpdatedAtMax(updatedAtMax string) *ListManagedKeyVersionsOptions {
	_options.UpdatedAtMax = core.StringPtr(updatedAtMax)
	return _options
}

// SetRotatedAtMin : Allow user to set RotatedAtMin
func (_options *ListManagedKeyVersionsOptions) SetRotatedAtMin(rotatedAtMin string) *ListManagedKeyVersionsOptions {
	_options.RotatedAtMin = core.StringPtr(rotatedAtMin)
	return _options
}

// SetRotatedAtMax : Allow user to set RotatedAtMax
func (_options *ListManagedKeyVersionsOptions) SetRotatedAtMax(rotatedAtMax string) *ListManagedKeyVersionsOptions {
	_options.RotatedAtMax = core.StringPtr(rotatedAtMax)
	return _options
}

// SetSize : Allow user to set Size
func (_options *ListManagedKeyVersionsOptions) SetSize(size string) *ListManagedKeyVersionsOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetSizeMin : Allow user to set SizeMin
func (_options *ListManagedKeyVersionsOptions) SetSizeMin(sizeMin string) *ListManagedKeyVersionsOptions {
	_options.SizeMin = core.StringPtr(sizeMin)
	return _options
}

// SetSizeMax : Allow user to set SizeMax
func (_options *ListManagedKeyVersionsOptions) SetSizeMax(sizeMax string) *ListManagedKeyVersionsOptions {
	_options.SizeMax = core.StringPtr(sizeMax)
	return _options
}

// SetReferencedKeystoresType : Allow user to set ReferencedKeystoresType
func (_options *ListManagedKeyVersionsOptions) SetReferencedKeystoresType(referencedKeystoresType []string) *ListManagedKeyVersionsOptions {
	_options.ReferencedKeystoresType = referencedKeystoresType
	return _options
}

// SetReferencedKeystoresName : Allow user to set ReferencedKeystoresName
func (_options *ListManagedKeyVersionsOptions) SetReferencedKeystoresName(referencedKeystoresName []string) *ListManagedKeyVersionsOptions {
	_options.ReferencedKeystoresName = referencedKeystoresName
	return _options
}

// SetInstancesKeystoreType : Allow user to set InstancesKeystoreType
func (_options *ListManagedKeyVersionsOptions) SetInstancesKeystoreType(instancesKeystoreType []string) *ListManagedKeyVersionsOptions {
	_options.InstancesKeystoreType = instancesKeystoreType
	return _options
}

// SetTemplateName : Allow user to set TemplateName
func (_options *ListManagedKeyVersionsOptions) SetTemplateName(templateName string) *ListManagedKeyVersionsOptions {
	_options.TemplateName = core.StringPtr(templateName)
	return _options
}

// SetTemplateID : Allow user to set TemplateID
func (_options *ListManagedKeyVersionsOptions) SetTemplateID(templateID []string) *ListManagedKeyVersionsOptions {
	_options.TemplateID = templateID
	return _options
}

// SetTemplateType : Allow user to set TemplateType
func (_options *ListManagedKeyVersionsOptions) SetTemplateType(templateType []string) *ListManagedKeyVersionsOptions {
	_options.TemplateType = templateType
	return _options
}

// SetStatusInKeystoresKeystoreSyncFlag : Allow user to set StatusInKeystoresKeystoreSyncFlag
func (_options *ListManagedKeyVersionsOptions) SetStatusInKeystoresKeystoreSyncFlag(statusInKeystoresKeystoreSyncFlag []string) *ListManagedKeyVersionsOptions {
	_options.StatusInKeystoresKeystoreSyncFlag = statusInKeystoresKeystoreSyncFlag
	return _options
}

// SetTemplateAlignmentStatus : Allow user to set TemplateAlignmentStatus
func (_options *ListManagedKeyVersionsOptions) SetTemplateAlignmentStatus(templateAlignmentStatus string) *ListManagedKeyVersionsOptions {
	_options.TemplateAlignmentStatus = core.StringPtr(templateAlignmentStatus)
	return _options
}

// SetKeyMaterialPresent : Allow user to set KeyMaterialPresent
func (_options *ListManagedKeyVersionsOptions) SetKeyMaterialPresent(keyMaterialPresent []string) *ListManagedKeyVersionsOptions {
	_options.KeyMaterialPresent = keyMaterialPresent
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListManagedKeyVersionsOptions) SetHeaders(param map[string]string) *ListManagedKeyVersionsOptions {
	options.Headers = param
	return options
}

// ListManagedKeysFromKeystoreOptions : The ListManagedKeysFromKeystore options.
type ListManagedKeysFromKeystoreOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// The type of the response: application/json, application/vnd.ibm.uko.managed-key-list.v4.1+json,
	// application/vnd.ibm.uko.managed-key-list.v4.1.json+zip, or application/vnd.ibm.uko.managed-key-list.v4.1.csv+zip.
	Accept *string `json:"Accept,omitempty"`

	// The algorithm of a returned key.
	Algorithm []string `json:"algorithm,omitempty"`

	// The state that returned keys are to be in.
	State []string `json:"state,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// The label of the key.
	Label *string `json:"label,omitempty"`

	// Return only managed keys whose activation_date matches the parameter.
	ActivationDate *string `json:"activation_date,omitempty"`

	// Return only managed keys whose activation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMin *string `json:"activation_date_min,omitempty"`

	// Return only managed keys whose activation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMax *string `json:"activation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter. This query parameter cannot be used in
	// conjunction with the 'expiration_date' query parameter.
	DeactivationDate *string `json:"deactivation_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMin *string `json:"deactivation_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMax *string `json:"deactivation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter.
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMin *string `json:"expiration_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMax *string `json:"expiration_date_max,omitempty"`

	// Return only managed keys whose created_at matches the parameter.
	CreatedAt *string `json:"created_at,omitempty"`

	// Return only managed keys whose created_at is at or after the parameter value. This query parameter cannot be used in
	// conjunction with the 'created_at' query parameter.
	CreatedAtMin *string `json:"created_at_min,omitempty"`

	// Return only managed keys whose created_at is at or before the parameter value. This query parameter cannot be used
	// in conjunction with the 'created_at' query parameter.
	CreatedAtMax *string `json:"created_at_max,omitempty"`

	// Return only managed keys whose updated_at matches the parameter.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Return only managed keys whose updated_at is after the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMin *string `json:"updated_at_min,omitempty"`

	// Return only managed keys whose updated_at is before the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMax *string `json:"updated_at_max,omitempty"`

	// Return only managed keys whose rotated_at is after the parameter value.
	RotatedAtMin *string `json:"rotated_at_min,omitempty"`

	// Return only managed keys whose rotated_at is before the parameter value.
	RotatedAtMax *string `json:"rotated_at_max,omitempty"`

	// The size of the key.
	Size *string `json:"size,omitempty"`

	// The minimum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMin *string `json:"size_min,omitempty"`

	// The maximum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMax *string `json:"size_max,omitempty"`

	// Return only managed keys whose template name begins with the string.
	TemplateName *string `json:"template.name,omitempty"`

	// Return only managed keys with the given template UUID.
	TemplateID []string `json:"template.id,omitempty"`

	// Return only managed keys with the given template type.
	TemplateType []string `json:"template.type[],omitempty"`

	// Return only Managed Keys whose status_in_keystores contains one of the specified keystore_sync_flag.
	StatusInKeystoresKeystoreSyncFlag []string `json:"status_in_keystores[].keystore_sync_flag,omitempty"`

	// Return only managed keys with the given alignment status.
	TemplateAlignmentStatus *string `json:"template.alignment_status,omitempty"`

	// Return only managed keys which key material is present on given set locations.
	KeyMaterialPresent []string `json:"key_material_present,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListManagedKeysFromKeystoreOptions.Algorithm property.
// The algorithm of the key.
const (
	ListManagedKeysFromKeystoreOptions_Algorithm_Aes = "aes"
	ListManagedKeysFromKeystoreOptions_Algorithm_Des = "des"
	ListManagedKeysFromKeystoreOptions_Algorithm_Dilithium = "dilithium"
	ListManagedKeysFromKeystoreOptions_Algorithm_Ec = "ec"
	ListManagedKeysFromKeystoreOptions_Algorithm_Hmac = "hmac"
	ListManagedKeysFromKeystoreOptions_Algorithm_Rsa = "rsa"
)

// Constants associated with the ListManagedKeysFromKeystoreOptions.State property.
// The state of the key.
const (
	ListManagedKeysFromKeystoreOptions_State_Active = "active"
	ListManagedKeysFromKeystoreOptions_State_Compromised = "compromised"
	ListManagedKeysFromKeystoreOptions_State_Deactivated = "deactivated"
	ListManagedKeysFromKeystoreOptions_State_Destroyed = "destroyed"
	ListManagedKeysFromKeystoreOptions_State_DestroyedCompromised = "destroyed_compromised"
	ListManagedKeysFromKeystoreOptions_State_PreActivation = "pre_activation"
)

// Constants associated with the ListManagedKeysFromKeystoreOptions.TemplateType property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	ListManagedKeysFromKeystoreOptions_TemplateType_Shadow = "shadow"
	ListManagedKeysFromKeystoreOptions_TemplateType_System = "system"
	ListManagedKeysFromKeystoreOptions_TemplateType_UserDefined = "user_defined"
)

// Constants associated with the ListManagedKeysFromKeystoreOptions.StatusInKeystoresKeystoreSyncFlag property.
// Flag to represent synchronization status between UKO Managed Key and Target Keystore. Possible status flags. ok:
// managed key state is the same as target keystore state, out_of_sync: managed key state is different than target
// keystore state.
const (
	ListManagedKeysFromKeystoreOptions_StatusInKeystoresKeystoreSyncFlag_Error = "error"
	ListManagedKeysFromKeystoreOptions_StatusInKeystoresKeystoreSyncFlag_Ok = "ok"
	ListManagedKeysFromKeystoreOptions_StatusInKeystoresKeystoreSyncFlag_OutOfSync = "out_of_sync"
	ListManagedKeysFromKeystoreOptions_StatusInKeystoresKeystoreSyncFlag_VerifyingSync = "verifying_sync"
)

// Constants associated with the ListManagedKeysFromKeystoreOptions.TemplateAlignmentStatus property.
// Return only managed keys with the given alignment status.
const (
	ListManagedKeysFromKeystoreOptions_TemplateAlignmentStatus_Aligned = "aligned"
	ListManagedKeysFromKeystoreOptions_TemplateAlignmentStatus_Unaligned = "unaligned"
)

// Constants associated with the ListManagedKeysFromKeystoreOptions.KeyMaterialPresent property.
// Location where key material is present.
const (
	ListManagedKeysFromKeystoreOptions_KeyMaterialPresent_Keystores = "keystores"
	ListManagedKeysFromKeystoreOptions_KeyMaterialPresent_Repository = "repository"
)

// NewListManagedKeysFromKeystoreOptions : Instantiate ListManagedKeysFromKeystoreOptions
func (*UkoV4) NewListManagedKeysFromKeystoreOptions(id string) *ListManagedKeysFromKeystoreOptions {
	return &ListManagedKeysFromKeystoreOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListManagedKeysFromKeystoreOptions) SetID(id string) *ListManagedKeysFromKeystoreOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *ListManagedKeysFromKeystoreOptions) SetAccept(accept string) *ListManagedKeysFromKeystoreOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetAlgorithm : Allow user to set Algorithm
func (_options *ListManagedKeysFromKeystoreOptions) SetAlgorithm(algorithm []string) *ListManagedKeysFromKeystoreOptions {
	_options.Algorithm = algorithm
	return _options
}

// SetState : Allow user to set State
func (_options *ListManagedKeysFromKeystoreOptions) SetState(state []string) *ListManagedKeysFromKeystoreOptions {
	_options.State = state
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListManagedKeysFromKeystoreOptions) SetLimit(limit int64) *ListManagedKeysFromKeystoreOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListManagedKeysFromKeystoreOptions) SetOffset(offset int64) *ListManagedKeysFromKeystoreOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListManagedKeysFromKeystoreOptions) SetSort(sort []string) *ListManagedKeysFromKeystoreOptions {
	_options.Sort = sort
	return _options
}

// SetLabel : Allow user to set Label
func (_options *ListManagedKeysFromKeystoreOptions) SetLabel(label string) *ListManagedKeysFromKeystoreOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetActivationDate : Allow user to set ActivationDate
func (_options *ListManagedKeysFromKeystoreOptions) SetActivationDate(activationDate string) *ListManagedKeysFromKeystoreOptions {
	_options.ActivationDate = core.StringPtr(activationDate)
	return _options
}

// SetActivationDateMin : Allow user to set ActivationDateMin
func (_options *ListManagedKeysFromKeystoreOptions) SetActivationDateMin(activationDateMin string) *ListManagedKeysFromKeystoreOptions {
	_options.ActivationDateMin = core.StringPtr(activationDateMin)
	return _options
}

// SetActivationDateMax : Allow user to set ActivationDateMax
func (_options *ListManagedKeysFromKeystoreOptions) SetActivationDateMax(activationDateMax string) *ListManagedKeysFromKeystoreOptions {
	_options.ActivationDateMax = core.StringPtr(activationDateMax)
	return _options
}

// SetDeactivationDate : Allow user to set DeactivationDate
func (_options *ListManagedKeysFromKeystoreOptions) SetDeactivationDate(deactivationDate string) *ListManagedKeysFromKeystoreOptions {
	_options.DeactivationDate = core.StringPtr(deactivationDate)
	return _options
}

// SetDeactivationDateMin : Allow user to set DeactivationDateMin
func (_options *ListManagedKeysFromKeystoreOptions) SetDeactivationDateMin(deactivationDateMin string) *ListManagedKeysFromKeystoreOptions {
	_options.DeactivationDateMin = core.StringPtr(deactivationDateMin)
	return _options
}

// SetDeactivationDateMax : Allow user to set DeactivationDateMax
func (_options *ListManagedKeysFromKeystoreOptions) SetDeactivationDateMax(deactivationDateMax string) *ListManagedKeysFromKeystoreOptions {
	_options.DeactivationDateMax = core.StringPtr(deactivationDateMax)
	return _options
}

// SetExpirationDate : Allow user to set ExpirationDate
func (_options *ListManagedKeysFromKeystoreOptions) SetExpirationDate(expirationDate string) *ListManagedKeysFromKeystoreOptions {
	_options.ExpirationDate = core.StringPtr(expirationDate)
	return _options
}

// SetExpirationDateMin : Allow user to set ExpirationDateMin
func (_options *ListManagedKeysFromKeystoreOptions) SetExpirationDateMin(expirationDateMin string) *ListManagedKeysFromKeystoreOptions {
	_options.ExpirationDateMin = core.StringPtr(expirationDateMin)
	return _options
}

// SetExpirationDateMax : Allow user to set ExpirationDateMax
func (_options *ListManagedKeysFromKeystoreOptions) SetExpirationDateMax(expirationDateMax string) *ListManagedKeysFromKeystoreOptions {
	_options.ExpirationDateMax = core.StringPtr(expirationDateMax)
	return _options
}

// SetCreatedAt : Allow user to set CreatedAt
func (_options *ListManagedKeysFromKeystoreOptions) SetCreatedAt(createdAt string) *ListManagedKeysFromKeystoreOptions {
	_options.CreatedAt = core.StringPtr(createdAt)
	return _options
}

// SetCreatedAtMin : Allow user to set CreatedAtMin
func (_options *ListManagedKeysFromKeystoreOptions) SetCreatedAtMin(createdAtMin string) *ListManagedKeysFromKeystoreOptions {
	_options.CreatedAtMin = core.StringPtr(createdAtMin)
	return _options
}

// SetCreatedAtMax : Allow user to set CreatedAtMax
func (_options *ListManagedKeysFromKeystoreOptions) SetCreatedAtMax(createdAtMax string) *ListManagedKeysFromKeystoreOptions {
	_options.CreatedAtMax = core.StringPtr(createdAtMax)
	return _options
}

// SetUpdatedAt : Allow user to set UpdatedAt
func (_options *ListManagedKeysFromKeystoreOptions) SetUpdatedAt(updatedAt string) *ListManagedKeysFromKeystoreOptions {
	_options.UpdatedAt = core.StringPtr(updatedAt)
	return _options
}

// SetUpdatedAtMin : Allow user to set UpdatedAtMin
func (_options *ListManagedKeysFromKeystoreOptions) SetUpdatedAtMin(updatedAtMin string) *ListManagedKeysFromKeystoreOptions {
	_options.UpdatedAtMin = core.StringPtr(updatedAtMin)
	return _options
}

// SetUpdatedAtMax : Allow user to set UpdatedAtMax
func (_options *ListManagedKeysFromKeystoreOptions) SetUpdatedAtMax(updatedAtMax string) *ListManagedKeysFromKeystoreOptions {
	_options.UpdatedAtMax = core.StringPtr(updatedAtMax)
	return _options
}

// SetRotatedAtMin : Allow user to set RotatedAtMin
func (_options *ListManagedKeysFromKeystoreOptions) SetRotatedAtMin(rotatedAtMin string) *ListManagedKeysFromKeystoreOptions {
	_options.RotatedAtMin = core.StringPtr(rotatedAtMin)
	return _options
}

// SetRotatedAtMax : Allow user to set RotatedAtMax
func (_options *ListManagedKeysFromKeystoreOptions) SetRotatedAtMax(rotatedAtMax string) *ListManagedKeysFromKeystoreOptions {
	_options.RotatedAtMax = core.StringPtr(rotatedAtMax)
	return _options
}

// SetSize : Allow user to set Size
func (_options *ListManagedKeysFromKeystoreOptions) SetSize(size string) *ListManagedKeysFromKeystoreOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetSizeMin : Allow user to set SizeMin
func (_options *ListManagedKeysFromKeystoreOptions) SetSizeMin(sizeMin string) *ListManagedKeysFromKeystoreOptions {
	_options.SizeMin = core.StringPtr(sizeMin)
	return _options
}

// SetSizeMax : Allow user to set SizeMax
func (_options *ListManagedKeysFromKeystoreOptions) SetSizeMax(sizeMax string) *ListManagedKeysFromKeystoreOptions {
	_options.SizeMax = core.StringPtr(sizeMax)
	return _options
}

// SetTemplateName : Allow user to set TemplateName
func (_options *ListManagedKeysFromKeystoreOptions) SetTemplateName(templateName string) *ListManagedKeysFromKeystoreOptions {
	_options.TemplateName = core.StringPtr(templateName)
	return _options
}

// SetTemplateID : Allow user to set TemplateID
func (_options *ListManagedKeysFromKeystoreOptions) SetTemplateID(templateID []string) *ListManagedKeysFromKeystoreOptions {
	_options.TemplateID = templateID
	return _options
}

// SetTemplateType : Allow user to set TemplateType
func (_options *ListManagedKeysFromKeystoreOptions) SetTemplateType(templateType []string) *ListManagedKeysFromKeystoreOptions {
	_options.TemplateType = templateType
	return _options
}

// SetStatusInKeystoresKeystoreSyncFlag : Allow user to set StatusInKeystoresKeystoreSyncFlag
func (_options *ListManagedKeysFromKeystoreOptions) SetStatusInKeystoresKeystoreSyncFlag(statusInKeystoresKeystoreSyncFlag []string) *ListManagedKeysFromKeystoreOptions {
	_options.StatusInKeystoresKeystoreSyncFlag = statusInKeystoresKeystoreSyncFlag
	return _options
}

// SetTemplateAlignmentStatus : Allow user to set TemplateAlignmentStatus
func (_options *ListManagedKeysFromKeystoreOptions) SetTemplateAlignmentStatus(templateAlignmentStatus string) *ListManagedKeysFromKeystoreOptions {
	_options.TemplateAlignmentStatus = core.StringPtr(templateAlignmentStatus)
	return _options
}

// SetKeyMaterialPresent : Allow user to set KeyMaterialPresent
func (_options *ListManagedKeysFromKeystoreOptions) SetKeyMaterialPresent(keyMaterialPresent []string) *ListManagedKeysFromKeystoreOptions {
	_options.KeyMaterialPresent = keyMaterialPresent
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListManagedKeysFromKeystoreOptions) SetHeaders(param map[string]string) *ListManagedKeysFromKeystoreOptions {
	options.Headers = param
	return options
}

// ListManagedKeysOptions : The ListManagedKeys options.
type ListManagedKeysOptions struct {
	// The type of the response: application/json, application/vnd.ibm.uko.managed-key-list.v4.1+json,
	// application/vnd.ibm.uko.managed-key-list.v4.1.json+zip, or application/vnd.ibm.uko.managed-key-list.v4.1.csv+zip.
	Accept *string `json:"Accept,omitempty"`

	// The UUID of the Vault.
	VaultID []string `json:"vault.id,omitempty"`

	// The algorithm of a returned key.
	Algorithm []string `json:"algorithm,omitempty"`

	// The state that returned keys are to be in.
	State []string `json:"state,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// The label of the key.
	Label *string `json:"label,omitempty"`

	// Return only managed keys whose activation_date matches the parameter.
	ActivationDate *string `json:"activation_date,omitempty"`

	// Return only managed keys whose activation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMin *string `json:"activation_date_min,omitempty"`

	// Return only managed keys whose activation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'activation_date' query parameter.
	ActivationDateMax *string `json:"activation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter. This query parameter cannot be used in
	// conjunction with the 'expiration_date' query parameter.
	DeactivationDate *string `json:"deactivation_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMin *string `json:"deactivation_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'expiration_date_min' and 'expiration_date_max'
	// query parameters.
	DeactivationDateMax *string `json:"deactivation_date_max,omitempty"`

	// Return only managed keys whose deactivation_date matches the parameter.
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// Return only managed keys whose deactivation_date is at or after the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMin *string `json:"expiration_date_min,omitempty"`

	// Return only managed keys whose deactivation_date is at or before the parameter value. This query parameter cannot be
	// used in conjunction with the 'deactivation_date', 'expiration_date', 'deactivation_date_min' and
	// 'deactivation_date_max' query parameters.
	ExpirationDateMax *string `json:"expiration_date_max,omitempty"`

	// Return only managed keys whose created_at matches the parameter.
	CreatedAt *string `json:"created_at,omitempty"`

	// Return only managed keys whose created_at is at or after the parameter value. This query parameter cannot be used in
	// conjunction with the 'created_at' query parameter.
	CreatedAtMin *string `json:"created_at_min,omitempty"`

	// Return only managed keys whose created_at is at or before the parameter value. This query parameter cannot be used
	// in conjunction with the 'created_at' query parameter.
	CreatedAtMax *string `json:"created_at_max,omitempty"`

	// Return only managed keys whose updated_at matches the parameter.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Return only managed keys whose updated_at is after the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMin *string `json:"updated_at_min,omitempty"`

	// Return only managed keys whose updated_at is before the parameter value. This query parameter cannot be used in
	// conjunction with the 'updated_at' query parameter.
	UpdatedAtMax *string `json:"updated_at_max,omitempty"`

	// Return only managed keys whose rotated_at is after the parameter value.
	RotatedAtMin *string `json:"rotated_at_min,omitempty"`

	// Return only managed keys whose rotated_at is before the parameter value.
	RotatedAtMax *string `json:"rotated_at_max,omitempty"`

	// The size of the key.
	Size *string `json:"size,omitempty"`

	// The minimum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMin *string `json:"size_min,omitempty"`

	// The maximum size of the key. This query parameter cannot be used in conjunction with the 'size' query parameter.
	SizeMax *string `json:"size_max,omitempty"`

	// Type of referenced keystore. This query parameter cannot be used in conjunction with the 'instances[].keystore.type'
	// query parameter.
	ReferencedKeystoresType []string `json:"referenced_keystores[].type,omitempty"`

	// Name of referenced keystore.
	ReferencedKeystoresName []string `json:"referenced_keystores[].name,omitempty"`

	// Type of keystore supported by one of the instances. This query parameter cannot be used in conjunction with the
	// 'referenced_keystores[].type' query parameter.
	InstancesKeystoreType []string `json:"instances[].keystore.type,omitempty"`

	// Return only managed keys whose template name begins with the string.
	TemplateName *string `json:"template.name,omitempty"`

	// Return only managed keys with the given template UUID.
	TemplateID []string `json:"template.id,omitempty"`

	// Return only managed keys with the given template type.
	TemplateType []string `json:"template.type[],omitempty"`

	// Return only Managed Keys whose status_in_keystores contains one of the specified keystore_sync_flag.
	StatusInKeystoresKeystoreSyncFlag []string `json:"status_in_keystores[].keystore_sync_flag,omitempty"`

	// Return only managed keys with the given alignment status.
	TemplateAlignmentStatus *string `json:"template.alignment_status,omitempty"`

	// Return only managed keys which key material is present on given set locations.
	KeyMaterialPresent []string `json:"key_material_present,omitempty"`

	// Return only managed keys with the given managing systems.
	ManagingSystems []string `json:"managing_systems,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListManagedKeysOptions.Algorithm property.
// The algorithm of the key.
const (
	ListManagedKeysOptions_Algorithm_Aes = "aes"
	ListManagedKeysOptions_Algorithm_Des = "des"
	ListManagedKeysOptions_Algorithm_Dilithium = "dilithium"
	ListManagedKeysOptions_Algorithm_Ec = "ec"
	ListManagedKeysOptions_Algorithm_Hmac = "hmac"
	ListManagedKeysOptions_Algorithm_Rsa = "rsa"
)

// Constants associated with the ListManagedKeysOptions.State property.
// The state of the key.
const (
	ListManagedKeysOptions_State_Active = "active"
	ListManagedKeysOptions_State_Compromised = "compromised"
	ListManagedKeysOptions_State_Deactivated = "deactivated"
	ListManagedKeysOptions_State_Destroyed = "destroyed"
	ListManagedKeysOptions_State_DestroyedCompromised = "destroyed_compromised"
	ListManagedKeysOptions_State_PreActivation = "pre_activation"
)

// Constants associated with the ListManagedKeysOptions.ReferencedKeystoresType property.
// Type of keystore.
const (
	ListManagedKeysOptions_ReferencedKeystoresType_AwsKms = "aws_kms"
	ListManagedKeysOptions_ReferencedKeystoresType_AzureKeyVault = "azure_key_vault"
	ListManagedKeysOptions_ReferencedKeystoresType_Cca = "cca"
	ListManagedKeysOptions_ReferencedKeystoresType_GoogleKms = "google_kms"
	ListManagedKeysOptions_ReferencedKeystoresType_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListManagedKeysOptions.InstancesKeystoreType property.
// Type of keystore.
const (
	ListManagedKeysOptions_InstancesKeystoreType_AwsKms = "aws_kms"
	ListManagedKeysOptions_InstancesKeystoreType_AzureKeyVault = "azure_key_vault"
	ListManagedKeysOptions_InstancesKeystoreType_Cca = "cca"
	ListManagedKeysOptions_InstancesKeystoreType_GoogleKms = "google_kms"
	ListManagedKeysOptions_InstancesKeystoreType_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the ListManagedKeysOptions.TemplateType property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	ListManagedKeysOptions_TemplateType_Shadow = "shadow"
	ListManagedKeysOptions_TemplateType_System = "system"
	ListManagedKeysOptions_TemplateType_UserDefined = "user_defined"
)

// Constants associated with the ListManagedKeysOptions.StatusInKeystoresKeystoreSyncFlag property.
// Flag to represent synchronization status between UKO Managed Key and Target Keystore. Possible status flags. ok:
// managed key state is the same as target keystore state, out_of_sync: managed key state is different than target
// keystore state.
const (
	ListManagedKeysOptions_StatusInKeystoresKeystoreSyncFlag_Error = "error"
	ListManagedKeysOptions_StatusInKeystoresKeystoreSyncFlag_Ok = "ok"
	ListManagedKeysOptions_StatusInKeystoresKeystoreSyncFlag_OutOfSync = "out_of_sync"
	ListManagedKeysOptions_StatusInKeystoresKeystoreSyncFlag_VerifyingSync = "verifying_sync"
)

// Constants associated with the ListManagedKeysOptions.TemplateAlignmentStatus property.
// Return only managed keys with the given alignment status.
const (
	ListManagedKeysOptions_TemplateAlignmentStatus_Aligned = "aligned"
	ListManagedKeysOptions_TemplateAlignmentStatus_Unaligned = "unaligned"
)

// Constants associated with the ListManagedKeysOptions.KeyMaterialPresent property.
// Location where key material is present.
const (
	ListManagedKeysOptions_KeyMaterialPresent_Keystores = "keystores"
	ListManagedKeysOptions_KeyMaterialPresent_Repository = "repository"
)

// Constants associated with the ListManagedKeysOptions.ManagingSystems property.
// Managing system of templates and keys.
const (
	ListManagedKeysOptions_ManagingSystems_Web = "web"
	ListManagedKeysOptions_ManagingSystems_Workstation = "workstation"
)

// NewListManagedKeysOptions : Instantiate ListManagedKeysOptions
func (*UkoV4) NewListManagedKeysOptions() *ListManagedKeysOptions {
	return &ListManagedKeysOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *ListManagedKeysOptions) SetAccept(accept string) *ListManagedKeysOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetVaultID : Allow user to set VaultID
func (_options *ListManagedKeysOptions) SetVaultID(vaultID []string) *ListManagedKeysOptions {
	_options.VaultID = vaultID
	return _options
}

// SetAlgorithm : Allow user to set Algorithm
func (_options *ListManagedKeysOptions) SetAlgorithm(algorithm []string) *ListManagedKeysOptions {
	_options.Algorithm = algorithm
	return _options
}

// SetState : Allow user to set State
func (_options *ListManagedKeysOptions) SetState(state []string) *ListManagedKeysOptions {
	_options.State = state
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListManagedKeysOptions) SetLimit(limit int64) *ListManagedKeysOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListManagedKeysOptions) SetOffset(offset int64) *ListManagedKeysOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListManagedKeysOptions) SetSort(sort []string) *ListManagedKeysOptions {
	_options.Sort = sort
	return _options
}

// SetLabel : Allow user to set Label
func (_options *ListManagedKeysOptions) SetLabel(label string) *ListManagedKeysOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetActivationDate : Allow user to set ActivationDate
func (_options *ListManagedKeysOptions) SetActivationDate(activationDate string) *ListManagedKeysOptions {
	_options.ActivationDate = core.StringPtr(activationDate)
	return _options
}

// SetActivationDateMin : Allow user to set ActivationDateMin
func (_options *ListManagedKeysOptions) SetActivationDateMin(activationDateMin string) *ListManagedKeysOptions {
	_options.ActivationDateMin = core.StringPtr(activationDateMin)
	return _options
}

// SetActivationDateMax : Allow user to set ActivationDateMax
func (_options *ListManagedKeysOptions) SetActivationDateMax(activationDateMax string) *ListManagedKeysOptions {
	_options.ActivationDateMax = core.StringPtr(activationDateMax)
	return _options
}

// SetDeactivationDate : Allow user to set DeactivationDate
func (_options *ListManagedKeysOptions) SetDeactivationDate(deactivationDate string) *ListManagedKeysOptions {
	_options.DeactivationDate = core.StringPtr(deactivationDate)
	return _options
}

// SetDeactivationDateMin : Allow user to set DeactivationDateMin
func (_options *ListManagedKeysOptions) SetDeactivationDateMin(deactivationDateMin string) *ListManagedKeysOptions {
	_options.DeactivationDateMin = core.StringPtr(deactivationDateMin)
	return _options
}

// SetDeactivationDateMax : Allow user to set DeactivationDateMax
func (_options *ListManagedKeysOptions) SetDeactivationDateMax(deactivationDateMax string) *ListManagedKeysOptions {
	_options.DeactivationDateMax = core.StringPtr(deactivationDateMax)
	return _options
}

// SetExpirationDate : Allow user to set ExpirationDate
func (_options *ListManagedKeysOptions) SetExpirationDate(expirationDate string) *ListManagedKeysOptions {
	_options.ExpirationDate = core.StringPtr(expirationDate)
	return _options
}

// SetExpirationDateMin : Allow user to set ExpirationDateMin
func (_options *ListManagedKeysOptions) SetExpirationDateMin(expirationDateMin string) *ListManagedKeysOptions {
	_options.ExpirationDateMin = core.StringPtr(expirationDateMin)
	return _options
}

// SetExpirationDateMax : Allow user to set ExpirationDateMax
func (_options *ListManagedKeysOptions) SetExpirationDateMax(expirationDateMax string) *ListManagedKeysOptions {
	_options.ExpirationDateMax = core.StringPtr(expirationDateMax)
	return _options
}

// SetCreatedAt : Allow user to set CreatedAt
func (_options *ListManagedKeysOptions) SetCreatedAt(createdAt string) *ListManagedKeysOptions {
	_options.CreatedAt = core.StringPtr(createdAt)
	return _options
}

// SetCreatedAtMin : Allow user to set CreatedAtMin
func (_options *ListManagedKeysOptions) SetCreatedAtMin(createdAtMin string) *ListManagedKeysOptions {
	_options.CreatedAtMin = core.StringPtr(createdAtMin)
	return _options
}

// SetCreatedAtMax : Allow user to set CreatedAtMax
func (_options *ListManagedKeysOptions) SetCreatedAtMax(createdAtMax string) *ListManagedKeysOptions {
	_options.CreatedAtMax = core.StringPtr(createdAtMax)
	return _options
}

// SetUpdatedAt : Allow user to set UpdatedAt
func (_options *ListManagedKeysOptions) SetUpdatedAt(updatedAt string) *ListManagedKeysOptions {
	_options.UpdatedAt = core.StringPtr(updatedAt)
	return _options
}

// SetUpdatedAtMin : Allow user to set UpdatedAtMin
func (_options *ListManagedKeysOptions) SetUpdatedAtMin(updatedAtMin string) *ListManagedKeysOptions {
	_options.UpdatedAtMin = core.StringPtr(updatedAtMin)
	return _options
}

// SetUpdatedAtMax : Allow user to set UpdatedAtMax
func (_options *ListManagedKeysOptions) SetUpdatedAtMax(updatedAtMax string) *ListManagedKeysOptions {
	_options.UpdatedAtMax = core.StringPtr(updatedAtMax)
	return _options
}

// SetRotatedAtMin : Allow user to set RotatedAtMin
func (_options *ListManagedKeysOptions) SetRotatedAtMin(rotatedAtMin string) *ListManagedKeysOptions {
	_options.RotatedAtMin = core.StringPtr(rotatedAtMin)
	return _options
}

// SetRotatedAtMax : Allow user to set RotatedAtMax
func (_options *ListManagedKeysOptions) SetRotatedAtMax(rotatedAtMax string) *ListManagedKeysOptions {
	_options.RotatedAtMax = core.StringPtr(rotatedAtMax)
	return _options
}

// SetSize : Allow user to set Size
func (_options *ListManagedKeysOptions) SetSize(size string) *ListManagedKeysOptions {
	_options.Size = core.StringPtr(size)
	return _options
}

// SetSizeMin : Allow user to set SizeMin
func (_options *ListManagedKeysOptions) SetSizeMin(sizeMin string) *ListManagedKeysOptions {
	_options.SizeMin = core.StringPtr(sizeMin)
	return _options
}

// SetSizeMax : Allow user to set SizeMax
func (_options *ListManagedKeysOptions) SetSizeMax(sizeMax string) *ListManagedKeysOptions {
	_options.SizeMax = core.StringPtr(sizeMax)
	return _options
}

// SetReferencedKeystoresType : Allow user to set ReferencedKeystoresType
func (_options *ListManagedKeysOptions) SetReferencedKeystoresType(referencedKeystoresType []string) *ListManagedKeysOptions {
	_options.ReferencedKeystoresType = referencedKeystoresType
	return _options
}

// SetReferencedKeystoresName : Allow user to set ReferencedKeystoresName
func (_options *ListManagedKeysOptions) SetReferencedKeystoresName(referencedKeystoresName []string) *ListManagedKeysOptions {
	_options.ReferencedKeystoresName = referencedKeystoresName
	return _options
}

// SetInstancesKeystoreType : Allow user to set InstancesKeystoreType
func (_options *ListManagedKeysOptions) SetInstancesKeystoreType(instancesKeystoreType []string) *ListManagedKeysOptions {
	_options.InstancesKeystoreType = instancesKeystoreType
	return _options
}

// SetTemplateName : Allow user to set TemplateName
func (_options *ListManagedKeysOptions) SetTemplateName(templateName string) *ListManagedKeysOptions {
	_options.TemplateName = core.StringPtr(templateName)
	return _options
}

// SetTemplateID : Allow user to set TemplateID
func (_options *ListManagedKeysOptions) SetTemplateID(templateID []string) *ListManagedKeysOptions {
	_options.TemplateID = templateID
	return _options
}

// SetTemplateType : Allow user to set TemplateType
func (_options *ListManagedKeysOptions) SetTemplateType(templateType []string) *ListManagedKeysOptions {
	_options.TemplateType = templateType
	return _options
}

// SetStatusInKeystoresKeystoreSyncFlag : Allow user to set StatusInKeystoresKeystoreSyncFlag
func (_options *ListManagedKeysOptions) SetStatusInKeystoresKeystoreSyncFlag(statusInKeystoresKeystoreSyncFlag []string) *ListManagedKeysOptions {
	_options.StatusInKeystoresKeystoreSyncFlag = statusInKeystoresKeystoreSyncFlag
	return _options
}

// SetTemplateAlignmentStatus : Allow user to set TemplateAlignmentStatus
func (_options *ListManagedKeysOptions) SetTemplateAlignmentStatus(templateAlignmentStatus string) *ListManagedKeysOptions {
	_options.TemplateAlignmentStatus = core.StringPtr(templateAlignmentStatus)
	return _options
}

// SetKeyMaterialPresent : Allow user to set KeyMaterialPresent
func (_options *ListManagedKeysOptions) SetKeyMaterialPresent(keyMaterialPresent []string) *ListManagedKeysOptions {
	_options.KeyMaterialPresent = keyMaterialPresent
	return _options
}

// SetManagingSystems : Allow user to set ManagingSystems
func (_options *ListManagedKeysOptions) SetManagingSystems(managingSystems []string) *ListManagedKeysOptions {
	_options.ManagingSystems = managingSystems
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListManagedKeysOptions) SetHeaders(param map[string]string) *ListManagedKeysOptions {
	options.Headers = param
	return options
}

// ListVaultsOptions : The ListVaults options.
type ListVaultsOptions struct {
	// The type of the response: application/json, application/vnd.ibm.uko.vault-list.v4.1+json,
	// application/vnd.ibm.uko.vault-list.v4.1.json+zip, or application/vnd.ibm.uko.vault-list.v4.1.csv+zip.
	Accept *string `json:"Accept,omitempty"`

	// The number of resources to retrieve.
	Limit *int64 `json:"limit,omitempty"`

	// The number of resources to skip.
	Offset *int64 `json:"offset,omitempty"`

	// Define sorting order.
	Sort []string `json:"sort,omitempty"`

	// Return only vaults whose name begin with the string.
	Name *string `json:"name,omitempty"`

	// Return only vaults whose description contains the string.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVaultsOptions : Instantiate ListVaultsOptions
func (*UkoV4) NewListVaultsOptions() *ListVaultsOptions {
	return &ListVaultsOptions{}
}

// SetAccept : Allow user to set Accept
func (_options *ListVaultsOptions) SetAccept(accept string) *ListVaultsOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListVaultsOptions) SetLimit(limit int64) *ListVaultsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListVaultsOptions) SetOffset(offset int64) *ListVaultsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListVaultsOptions) SetSort(sort []string) *ListVaultsOptions {
	_options.Sort = sort
	return _options
}

// SetName : Allow user to set Name
func (_options *ListVaultsOptions) SetName(name string) *ListVaultsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ListVaultsOptions) SetDescription(description string) *ListVaultsOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVaultsOptions) SetHeaders(param map[string]string) *ListVaultsOptions {
	options.Headers = param
	return options
}

// RotateManagedKeyOptions : The RotateManagedKey options.
type RotateManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRotateManagedKeyOptions : Instantiate RotateManagedKeyOptions
func (*UkoV4) NewRotateManagedKeyOptions(id string, ifMatch string) *RotateManagedKeyOptions {
	return &RotateManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *RotateManagedKeyOptions) SetID(id string) *RotateManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *RotateManagedKeyOptions) SetIfMatch(ifMatch string) *RotateManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RotateManagedKeyOptions) SetHeaders(param map[string]string) *RotateManagedKeyOptions {
	options.Headers = param
	return options
}

// SyncManagedKeyOptions : The SyncManagedKey options.
type SyncManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSyncManagedKeyOptions : Instantiate SyncManagedKeyOptions
func (*UkoV4) NewSyncManagedKeyOptions(id string, ifMatch string) *SyncManagedKeyOptions {
	return &SyncManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *SyncManagedKeyOptions) SetID(id string) *SyncManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *SyncManagedKeyOptions) SetIfMatch(ifMatch string) *SyncManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SyncManagedKeyOptions) SetHeaders(param map[string]string) *SyncManagedKeyOptions {
	options.Headers = param
	return options
}

// UnarchiveKeyTemplateOptions : The UnarchiveKeyTemplate options.
type UnarchiveKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnarchiveKeyTemplateOptions : Instantiate UnarchiveKeyTemplateOptions
func (*UkoV4) NewUnarchiveKeyTemplateOptions(id string, ifMatch string) *UnarchiveKeyTemplateOptions {
	return &UnarchiveKeyTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *UnarchiveKeyTemplateOptions) SetID(id string) *UnarchiveKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UnarchiveKeyTemplateOptions) SetIfMatch(ifMatch string) *UnarchiveKeyTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UnarchiveKeyTemplateOptions) SetHeaders(param map[string]string) *UnarchiveKeyTemplateOptions {
	options.Headers = param
	return options
}

// UpdateKeyTemplateOptions : The UpdateKeyTemplate options.
type UpdateKeyTemplateOptions struct {
	// UUID of the template.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// A human-readable name to assign to your template.
	Name *string `json:"name,omitempty"`

	// Updated keystore related properties.
	Keystores []KeystoresPropertiesUpdateIntf `json:"keystores,omitempty"`

	// Updated description of the key template.
	Description *string `json:"description,omitempty"`

	// Updated key related properties.
	Key *KeyPropertiesUpdate `json:"key,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateKeyTemplateOptions : Instantiate UpdateKeyTemplateOptions
func (*UkoV4) NewUpdateKeyTemplateOptions(id string, ifMatch string) *UpdateKeyTemplateOptions {
	return &UpdateKeyTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateKeyTemplateOptions) SetID(id string) *UpdateKeyTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateKeyTemplateOptions) SetIfMatch(ifMatch string) *UpdateKeyTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateKeyTemplateOptions) SetName(name string) *UpdateKeyTemplateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetKeystores : Allow user to set Keystores
func (_options *UpdateKeyTemplateOptions) SetKeystores(keystores []KeystoresPropertiesUpdateIntf) *UpdateKeyTemplateOptions {
	_options.Keystores = keystores
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateKeyTemplateOptions) SetDescription(description string) *UpdateKeyTemplateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetKey : Allow user to set Key
func (_options *UpdateKeyTemplateOptions) SetKey(key *KeyPropertiesUpdate) *UpdateKeyTemplateOptions {
	_options.Key = key
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateKeyTemplateOptions) SetHeaders(param map[string]string) *UpdateKeyTemplateOptions {
	options.Headers = param
	return options
}

// UpdateKeystoreOptions : The UpdateKeystore options.
type UpdateKeystoreOptions struct {
	// UUID of the keystore.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Keystore properties to update.
	KeystoreBody KeystoreUpdateRequestIntf `json:"keystoreBody" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateKeystoreOptions : Instantiate UpdateKeystoreOptions
func (*UkoV4) NewUpdateKeystoreOptions(id string, ifMatch string, keystoreBody KeystoreUpdateRequestIntf) *UpdateKeystoreOptions {
	return &UpdateKeystoreOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
		KeystoreBody: keystoreBody,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateKeystoreOptions) SetID(id string) *UpdateKeystoreOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateKeystoreOptions) SetIfMatch(ifMatch string) *UpdateKeystoreOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetKeystoreBody : Allow user to set KeystoreBody
func (_options *UpdateKeystoreOptions) SetKeystoreBody(keystoreBody KeystoreUpdateRequestIntf) *UpdateKeystoreOptions {
	_options.KeystoreBody = keystoreBody
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateKeystoreOptions) SetHeaders(param map[string]string) *UpdateKeystoreOptions {
	options.Headers = param
	return options
}

// UpdateManagedKeyFromTemplateOptions : The UpdateManagedKeyFromTemplate options.
type UpdateManagedKeyFromTemplateOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Do not create/update/delete a resource, only verify and validate if resource can be created/updated/deleted with
	// given request successfully.
	DryRun *bool `json:"dry_run,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateManagedKeyFromTemplateOptions : Instantiate UpdateManagedKeyFromTemplateOptions
func (*UkoV4) NewUpdateManagedKeyFromTemplateOptions(id string, ifMatch string) *UpdateManagedKeyFromTemplateOptions {
	return &UpdateManagedKeyFromTemplateOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateManagedKeyFromTemplateOptions) SetID(id string) *UpdateManagedKeyFromTemplateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateManagedKeyFromTemplateOptions) SetIfMatch(ifMatch string) *UpdateManagedKeyFromTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetDryRun : Allow user to set DryRun
func (_options *UpdateManagedKeyFromTemplateOptions) SetDryRun(dryRun bool) *UpdateManagedKeyFromTemplateOptions {
	_options.DryRun = core.BoolPtr(dryRun)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateManagedKeyFromTemplateOptions) SetHeaders(param map[string]string) *UpdateManagedKeyFromTemplateOptions {
	options.Headers = param
	return options
}

// UpdateManagedKeyOptions : The UpdateManagedKey options.
type UpdateManagedKeyOptions struct {
	// UUID of the key.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The label of the key.
	Label *string `json:"label,omitempty"`

	// Activation date provided in format: YYYY-MM-DD.
	ActivationDate *strfmt.Date `json:"activation_date,omitempty"`

	// Expiration date provided in format: YYYY-MM-DD.
	ExpirationDate *strfmt.Date `json:"expiration_date,omitempty"`

	// Updated description of the managed key.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateManagedKeyOptions : Instantiate UpdateManagedKeyOptions
func (*UkoV4) NewUpdateManagedKeyOptions(id string, ifMatch string) *UpdateManagedKeyOptions {
	return &UpdateManagedKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateManagedKeyOptions) SetID(id string) *UpdateManagedKeyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateManagedKeyOptions) SetIfMatch(ifMatch string) *UpdateManagedKeyOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetLabel : Allow user to set Label
func (_options *UpdateManagedKeyOptions) SetLabel(label string) *UpdateManagedKeyOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetActivationDate : Allow user to set ActivationDate
func (_options *UpdateManagedKeyOptions) SetActivationDate(activationDate *strfmt.Date) *UpdateManagedKeyOptions {
	_options.ActivationDate = activationDate
	return _options
}

// SetExpirationDate : Allow user to set ExpirationDate
func (_options *UpdateManagedKeyOptions) SetExpirationDate(expirationDate *strfmt.Date) *UpdateManagedKeyOptions {
	_options.ExpirationDate = expirationDate
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateManagedKeyOptions) SetDescription(description string) *UpdateManagedKeyOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateManagedKeyOptions) SetHeaders(param map[string]string) *UpdateManagedKeyOptions {
	options.Headers = param
	return options
}

// UpdateVaultOptions : The UpdateVault options.
type UpdateVaultOptions struct {
	// UUID of the vault.
	ID *string `json:"id" validate:"required,ne="`

	// Precondition of the update; Value of the ETag from the header on a GET request.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Updated name of the vault.
	Name *string `json:"name,omitempty"`

	// Updated description of the vault.
	Description *string `json:"description,omitempty"`

	// The label of the recovery key to use for this vault.
	RecoveryKeyLabel *string `json:"recovery_key_label,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateVaultOptions : Instantiate UpdateVaultOptions
func (*UkoV4) NewUpdateVaultOptions(id string, ifMatch string) *UpdateVaultOptions {
	return &UpdateVaultOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateVaultOptions) SetID(id string) *UpdateVaultOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateVaultOptions) SetIfMatch(ifMatch string) *UpdateVaultOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateVaultOptions) SetName(name string) *UpdateVaultOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateVaultOptions) SetDescription(description string) *UpdateVaultOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetRecoveryKeyLabel : Allow user to set RecoveryKeyLabel
func (_options *UpdateVaultOptions) SetRecoveryKeyLabel(recoveryKeyLabel string) *UpdateVaultOptions {
	_options.RecoveryKeyLabel = core.StringPtr(recoveryKeyLabel)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateVaultOptions) SetHeaders(param map[string]string) *UpdateVaultOptions {
	options.Headers = param
	return options
}

// ApiError : An error encountered while using the application.
type ApiError struct {
	// The HTTP status code used for the response.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A unique identifier that is attached to the request and the message that allows to refer to the specific transaction
	// or the event chain.
	Trace *string `json:"trace,omitempty"`

	// A list of errors.
	Errors []ErrorModel `json:"errors" validate:"required"`
}

// UnmarshalApiError unmarshals an instance of ApiError from the specified map of raw messages.
func UnmarshalApiError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiError)
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssociatedResource : Associated resource is any object or entity that is using or referencing a Managed Key in any cloud.
type AssociatedResource struct {
	// An identifier uniquely identifing this associated resource.
	ID *string `json:"id" validate:"required"`

	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// Reference to a managed key.
	ManagedKey *ManagedKeyReference `json:"managed_key,omitempty"`

	// Reference to a target keystore.
	ReferencedKeystore *TargetKeystoreReference `json:"referenced_keystore,omitempty"`

	// ID of the key in this keystore. Every keystore may use different format such as: UUID, GUID, CRN, URI.
	KeyIdInKeystore *string `json:"key_id_in_keystore" validate:"required"`

	// Name of the associated resource.
	Name *string `json:"name" validate:"required"`

	// Type of the associated resource, in reverse domain name notation. Currently only 'com_ibm_cloud_kms_registration' is
	// supported.
	Type *string `json:"type" validate:"required"`

	// Properties of an associated resource of type IBM Cloud KMS Registration; com_ibm_cloud_kms_registration.
	ComIbmCloudKmsRegistration *IbmCloudKmsRegistration `json:"com_ibm_cloud_kms_registration,omitempty"`
}

// UnmarshalAssociatedResource unmarshals an instance of AssociatedResource from the specified map of raw messages.
func UnmarshalAssociatedResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssociatedResource)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "managed_key", &obj.ManagedKey, UnmarshalManagedKeyReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "referenced_keystore", &obj.ReferencedKeystore, UnmarshalTargetKeystoreReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_id_in_keystore", &obj.KeyIdInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "com_ibm_cloud_kms_registration", &obj.ComIbmCloudKmsRegistration, UnmarshalIbmCloudKmsRegistration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssociatedResourceList : The base schema for listing associated resources.
type AssociatedResourceList struct {
	// The total count of all objects in the entire collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The number of resources that were requested in this query.
	Limit *int64 `json:"limit" validate:"required"`

	// The number of resources that were skipped in this query.
	Offset *int64 `json:"offset" validate:"required"`

	// URL of a resource.
	First *HrefObject `json:"first,omitempty"`

	// URL of a resource.
	Last *HrefObject `json:"last,omitempty"`

	// URL of a resource.
	Previous *HrefObject `json:"previous,omitempty"`

	// URL of a resource.
	Next *HrefObject `json:"next,omitempty"`

	// A list of associated resources.
	AssociatedResources []AssociatedResource `json:"associated_resources" validate:"required"`
}

// UnmarshalAssociatedResourceList unmarshals an instance of AssociatedResourceList from the specified map of raw messages.
func UnmarshalAssociatedResourceList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssociatedResourceList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "associated_resources", &obj.AssociatedResources, UnmarshalAssociatedResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *AssociatedResourceList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// ErrorModel : error model.
type ErrorModel struct {
	// Identifier of the error.
	Code *string `json:"code" validate:"required"`

	// A message explaining the problem, with potential suggestions how to address them.
	Message *string `json:"message" validate:"required"`

	// A publicly-accessible URL where information about the error can be read in a web browser. Since more infomration is
	// not always available, this field is optional.
	MoreInfo *string `json:"more_info,omitempty"`

	// Parameters of the message that can be used e.g. for i18n purposes in conjunction with the code. Since the message
	// may not contain any parameters, this field is optional.
	MessageParams []string `json:"message_params,omitempty"`

	// Target of the error.
	Target *Target `json:"target,omitempty"`
}

// UnmarshalErrorModel unmarshals an instance of ErrorModel from the specified map of raw messages.
func UnmarshalErrorModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorModel)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_params", &obj.MessageParams)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HrefObject : URL of a resource.
type HrefObject struct {
	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`
}

// UnmarshalHrefObject unmarshals an instance of HrefObject from the specified map of raw messages.
func UnmarshalHrefObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HrefObject)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IbmCloudKmsRegistration : Properties of an associated resource of type IBM Cloud KMS Registration; com_ibm_cloud_kms_registration.
type IbmCloudKmsRegistration struct {
	// A boolean that determines whether Key Protect must prevent deletion of a root key.
	PreventsKeyDeletion *bool `json:"prevents_key_deletion" validate:"required"`

	// Name of the IBM Cloud service, derived from the CRN. It will be empty if UKO is unable to contact the resource
	// controller.
	ServiceName *string `json:"service_name" validate:"required"`

	// Name of the IBM Cloud service's instance, derived from the CRN. It will be empty if UKO is unable to contact the
	// resource controller.
	ServiceInstanceName *string `json:"service_instance_name" validate:"required"`

	// The Cloud Resource Name (CRN) that represents the cloud resource, such as a Cloud Object Storage bucket, that is
	// associated with the key.
	Crn *string `json:"crn" validate:"required"`

	// Description of the purpose of the registration.
	Description *string `json:"description" validate:"required"`
}

// UnmarshalIbmCloudKmsRegistration unmarshals an instance of IbmCloudKmsRegistration from the specified map of raw messages.
func UnmarshalIbmCloudKmsRegistration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IbmCloudKmsRegistration)
	err = core.UnmarshalPrimitive(m, "prevents_key_deletion", &obj.PreventsKeyDeletion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_instance_name", &obj.ServiceInstanceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceInKeystore : Description of properties of a key within the context of keystores.
type InstanceInKeystore struct {
	Group *string `json:"group" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the InstanceInKeystore.Type property.
// Type of keystore.
const (
	InstanceInKeystore_Type_AwsKms = "aws_kms"
	InstanceInKeystore_Type_AzureKeyVault = "azure_key_vault"
	InstanceInKeystore_Type_Cca = "cca"
	InstanceInKeystore_Type_GoogleKms = "google_kms"
	InstanceInKeystore_Type_IbmCloudKms = "ibm_cloud_kms"
)

// UnmarshalInstanceInKeystore unmarshals an instance of InstanceInKeystore from the specified map of raw messages.
func UnmarshalInstanceInKeystore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceInKeystore)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceReference : InstanceReference struct
type InstanceReference struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`
}

// UnmarshalInstanceReference unmarshals an instance of InstanceReference from the specified map of raw messages.
func UnmarshalInstanceReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstance : KeyInstance struct
// Models which "extend" this model:
// - KeyInstanceGoogleKms
// - KeyInstanceAwsKms
// - KeyInstanceIbmCloudKms
// - KeyInstanceAzure
// - KeyInstanceCca
type KeyInstance struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore,omitempty"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore,omitempty"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level,omitempty"`

	GoogleKeyPurpose *string `json:"google_key_purpose,omitempty"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm,omitempty"`

	AzureKeyProtectionLevel *string `json:"azure_key_protection_level,omitempty"`

	AzureKeyOperations []string `json:"azure_key_operations,omitempty"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeyInstance.Type property.
// Type of the key instance.
const (
	KeyInstance_Type_KeyPair = "key_pair"
	KeyInstance_Type_PrivateKey = "private_key"
	KeyInstance_Type_PublicKey = "public_key"
	KeyInstance_Type_SecretKey = "secret_key"
)

// Constants associated with the KeyInstance.GoogleKeyProtectionLevel property.
const (
	KeyInstance_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeyInstance_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeyInstance.GoogleKeyPurpose property.
const (
	KeyInstance_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeyInstance_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeyInstance_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeyInstance_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeyInstance.GoogleKmsAlgorithm property.
const (
	KeyInstance_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeyInstance_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeyInstance_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeyInstance_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeyInstance_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeyInstance_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeyInstance_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeyInstance_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeyInstance_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

// Constants associated with the KeyInstance.AzureKeyProtectionLevel property.
const (
	KeyInstance_AzureKeyProtectionLevel_Hsm = "hsm"
	KeyInstance_AzureKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeyInstance.AzureKeyOperations property.
const (
	KeyInstance_AzureKeyOperations_Decrypt = "decrypt"
	KeyInstance_AzureKeyOperations_Encrypt = "encrypt"
	KeyInstance_AzureKeyOperations_Sign = "sign"
	KeyInstance_AzureKeyOperations_UnwrapKey = "unwrap_key"
	KeyInstance_AzureKeyOperations_Verify = "verify"
	KeyInstance_AzureKeyOperations_WrapKey = "wrap_key"
)

// Constants associated with the KeyInstance.CcaUsageControl property.
const (
	KeyInstance_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeyInstance_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeyInstance_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeyInstance.CcaKeyType property.
const (
	KeyInstance_CcaKeyType_Cipher = "cipher"
	KeyInstance_CcaKeyType_Data = "data"
	KeyInstance_CcaKeyType_Exporter = "exporter"
	KeyInstance_CcaKeyType_Importer = "importer"
)
func (*KeyInstance) isaKeyInstance() bool {
	return true
}

type KeyInstanceIntf interface {
	isaKeyInstance() bool
}

// UnmarshalKeyInstance unmarshals an instance of KeyInstance from the specified map of raw messages.
func UnmarshalKeyInstance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstance)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_protection_level", &obj.AzureKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_operations", &obj.AzureKeyOperations)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyProperties : Properties describing the properties of the managed key.
type KeyProperties struct {
	// The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
	Size *string `json:"size" validate:"required"`

	// The algorithm of the key.
	Algorithm *string `json:"algorithm" validate:"required"`

	// Key activation date can be provided as a period definition (e.g. P1Y means 1 year).
	ActivationDate *string `json:"activation_date" validate:"required"`

	// Key expiration date can be provided as a period definition (e.g. P1Y means 1 year).
	ExpirationDate *string `json:"expiration_date" validate:"required"`

	// The state that the key will be in after generation.
	State *string `json:"state" validate:"required"`

	// indicates whether to deactivate older versions of a key on rotation.
	DeactivateOnRotation *bool `json:"deactivate_on_rotation,omitempty"`
}

// Constants associated with the KeyProperties.Algorithm property.
// The algorithm of the key.
const (
	KeyProperties_Algorithm_Aes = "aes"
	KeyProperties_Algorithm_Des = "des"
	KeyProperties_Algorithm_Dilithium = "dilithium"
	KeyProperties_Algorithm_Ec = "ec"
	KeyProperties_Algorithm_Hmac = "hmac"
	KeyProperties_Algorithm_Rsa = "rsa"
)

// Constants associated with the KeyProperties.State property.
// The state that the key will be in after generation.
const (
	KeyProperties_State_Active = "active"
	KeyProperties_State_PreActivation = "pre_activation"
)

// NewKeyProperties : Instantiate KeyProperties (Generic Model Constructor)
func (*UkoV4) NewKeyProperties(size string, algorithm string, activationDate string, expirationDate string, state string) (_model *KeyProperties, err error) {
	_model = &KeyProperties{
		Size: core.StringPtr(size),
		Algorithm: core.StringPtr(algorithm),
		ActivationDate: core.StringPtr(activationDate),
		ExpirationDate: core.StringPtr(expirationDate),
		State: core.StringPtr(state),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalKeyProperties unmarshals an instance of KeyProperties from the specified map of raw messages.
func UnmarshalKeyProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyProperties)
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "algorithm", &obj.Algorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "activation_date", &obj.ActivationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration_date", &obj.ExpirationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deactivate_on_rotation", &obj.DeactivateOnRotation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyPropertiesUpdate : Updated key related properties.
type KeyPropertiesUpdate struct {
	// The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
	Size *string `json:"size,omitempty"`

	// Key activation date can be provided as a period definition (e.g. P1Y means 1 year).
	ActivationDate *string `json:"activation_date,omitempty"`

	// Key expiration date can be provided as a period definition (e.g. P1Y means 1 year).
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// The state that the key will be in after generation.
	State *string `json:"state,omitempty"`

	// indicates whether to deactivate older versions of a key on rotation.
	DeactivateOnRotation *bool `json:"deactivate_on_rotation,omitempty"`
}

// Constants associated with the KeyPropertiesUpdate.State property.
// The state that the key will be in after generation.
const (
	KeyPropertiesUpdate_State_Active = "active"
	KeyPropertiesUpdate_State_PreActivation = "pre_activation"
)

// UnmarshalKeyPropertiesUpdate unmarshals an instance of KeyPropertiesUpdate from the specified map of raw messages.
func UnmarshalKeyPropertiesUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyPropertiesUpdate)
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "activation_date", &obj.ActivationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration_date", &obj.ExpirationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deactivate_on_rotation", &obj.DeactivateOnRotation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyVerificationPattern : Key verification pattern is used to identify and distinguish cryptographic keys.
type KeyVerificationPattern struct {
	// The method used for calculating the verification pattern.
	Method *string `json:"method" validate:"required"`

	// The calculated value.
	Value *string `json:"value" validate:"required"`
}

// UnmarshalKeyVerificationPattern unmarshals an instance of KeyVerificationPattern from the specified map of raw messages.
func UnmarshalKeyVerificationPattern(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyVerificationPattern)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Keystore : A target keystore is keystore that is assigned to a vault. If it is an internal keystore, it can be created only in a
// vault. If it is an external keystore, you need to assign the external keystore to a vault when you connect your
// service instance to it.
// Models which "extend" this model:
// - KeystoreTypeGoogleKms
// - KeystoreTypeAwsKms
// - KeystoreTypeAzure
// - KeystoreTypeIbmCloudKms
// - KeystoreTypeCca
type Keystore struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name,omitempty"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status,omitempty"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials,omitempty"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location,omitempty"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring,omitempty"`

	// AWS Region.
	AwsRegion *string `json:"aws_region,omitempty"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name,omitempty"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group,omitempty"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id,omitempty"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password,omitempty"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant,omitempty"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id,omitempty"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint,omitempty"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint,omitempty"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key,omitempty"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id,omitempty"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant,omitempty"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls,omitempty"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host,omitempty"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port,omitempty"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash,omitempty"`
}

// Constants associated with the Keystore.Type property.
// Type of keystore.
const (
	Keystore_Type_AwsKms = "aws_kms"
	Keystore_Type_AzureKeyVault = "azure_key_vault"
	Keystore_Type_Cca = "cca"
	Keystore_Type_GoogleKms = "google_kms"
	Keystore_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the Keystore.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	Keystore_AzureEnvironment_Azure = "azure"
	Keystore_AzureEnvironment_AzureChina = "azure_china"
	Keystore_AzureEnvironment_AzureGermany = "azure_germany"
	Keystore_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the Keystore.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	Keystore_AzureVariant_Premium = "premium"
	Keystore_AzureVariant_Standard = "standard"
)

// Constants associated with the Keystore.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	Keystore_IbmVariant_Hpcs = "hpcs"
	Keystore_IbmVariant_Internal = "internal"
	Keystore_IbmVariant_KeyProtect = "key_protect"
)
func (*Keystore) isaKeystore() bool {
	return true
}

type KeystoreIntf interface {
	isaKeystore() bool
}

// UnmarshalKeystore unmarshals an instance of Keystore from the specified map of raw messages.
func UnmarshalKeystore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Keystore)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_credentials", &obj.GoogleCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_location", &obj.GoogleLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_project_id", &obj.GoogleProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_private_key_id", &obj.GooglePrivateKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_ring", &obj.GoogleKeyRing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_region", &obj.AwsRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_access_key_id", &obj.AwsAccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_secret_access_key", &obj.AwsSecretAccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_name", &obj.AzureServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_resource_group", &obj.AzureResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_location", &obj.AzureLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_client_id", &obj.AzureServicePrincipalClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_password", &obj.AzureServicePrincipalPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_tenant", &obj.AzureTenant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_subscription_id", &obj.AzureSubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_environment", &obj.AzureEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_variant", &obj.AzureVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_use_tls", &obj.CcaUseTls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_trusted_issuer", &obj.CcaTrustedIssuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_host", &obj.CcaHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_port", &obj.CcaPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_public_key_hash", &obj.CcaPublicKeyHash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequest : Properties required to create a keystore.
// Models which "extend" this model:
// - KeystoreCreationRequestKeystoreTypeAwsKmsCreate
// - KeystoreCreationRequestKeystoreTypeGoogleKmsCreate
// - KeystoreCreationRequestKeystoreTypeAzureCreate
// - KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate
// - KeystoreCreationRequestKeystoreTypeCcaCreate
type KeystoreCreationRequest struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// ID of the Vault where the entity is to be created in.
	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// AWS Region.
	AwsRegion *string `json:"aws_region,omitempty"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials,omitempty"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location,omitempty"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring,omitempty"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name,omitempty"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group,omitempty"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id,omitempty"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password,omitempty"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant,omitempty"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id,omitempty"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint,omitempty"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint,omitempty"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key,omitempty"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id,omitempty"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls,omitempty"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host,omitempty"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port,omitempty"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash,omitempty"`
}

// Constants associated with the KeystoreCreationRequest.Type property.
// Type of keystore.
const (
	KeystoreCreationRequest_Type_AwsKms = "aws_kms"
	KeystoreCreationRequest_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequest_Type_Cca = "cca"
	KeystoreCreationRequest_Type_GoogleKms = "google_kms"
	KeystoreCreationRequest_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequest.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	KeystoreCreationRequest_AzureEnvironment_Azure = "azure"
	KeystoreCreationRequest_AzureEnvironment_AzureChina = "azure_china"
	KeystoreCreationRequest_AzureEnvironment_AzureGermany = "azure_germany"
	KeystoreCreationRequest_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the KeystoreCreationRequest.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	KeystoreCreationRequest_AzureVariant_Premium = "premium"
	KeystoreCreationRequest_AzureVariant_Standard = "standard"
)

// Constants associated with the KeystoreCreationRequest.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequest_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequest_IbmVariant_Internal = "internal"
	KeystoreCreationRequest_IbmVariant_KeyProtect = "key_protect"
)
func (*KeystoreCreationRequest) isaKeystoreCreationRequest() bool {
	return true
}

type KeystoreCreationRequestIntf interface {
	isaKeystoreCreationRequest() bool
}

// UnmarshalKeystoreCreationRequest unmarshals an instance of KeystoreCreationRequest from the specified map of raw messages.
func UnmarshalKeystoreCreationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'type' not found in JSON object")
		return
	}
	if discValue == "google_kms" {
		err = core.UnmarshalModel(m, "", result, UnmarshalKeystoreCreationRequestKeystoreTypeGoogleKmsCreate)
	} else if discValue == "aws_kms" {
		err = core.UnmarshalModel(m, "", result, UnmarshalKeystoreCreationRequestKeystoreTypeAwsKmsCreate)
	} else if discValue == "azure_key_vault" {
		err = core.UnmarshalModel(m, "", result, UnmarshalKeystoreCreationRequestKeystoreTypeAzureCreate)
	} else if discValue == "ibm_cloud_kms" {
		err = core.UnmarshalModel(m, "", result, UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate)
	} else if discValue == "cca" {
		err = core.UnmarshalModel(m, "", result, UnmarshalKeystoreCreationRequestKeystoreTypeCcaCreate)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'type': %s", discValue)
	}
	return
}

// KeystoreList : The base schema for listing target keystores.
type KeystoreList struct {
	// The total count of all objects in the entire collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The number of resources that were requested in this query.
	Limit *int64 `json:"limit" validate:"required"`

	// The number of resources that were skipped in this query.
	Offset *int64 `json:"offset" validate:"required"`

	// URL of a resource.
	First *HrefObject `json:"first,omitempty"`

	// URL of a resource.
	Last *HrefObject `json:"last,omitempty"`

	// URL of a resource.
	Previous *HrefObject `json:"previous,omitempty"`

	// URL of a resource.
	Next *HrefObject `json:"next,omitempty"`

	// A list of target keystores.
	Keystores []KeystoreIntf `json:"keystores" validate:"required"`
}

// UnmarshalKeystoreList unmarshals an instance of KeystoreList from the specified map of raw messages.
func UnmarshalKeystoreList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystores", &obj.Keystores, UnmarshalKeystore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *KeystoreList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// KeystoreStatus : The status of the connection to the keystore.
type KeystoreStatus struct {
	// Date of last successful communication with a keystore.
	LastHeartbeat *strfmt.DateTime `json:"last_heartbeat,omitempty"`

	// Possible states of a keystore.
	HealthStatus *string `json:"health_status,omitempty"`

	// Message returned with the status.
	Message *string `json:"message,omitempty"`
}

// Constants associated with the KeystoreStatus.HealthStatus property.
// Possible states of a keystore.
const (
	KeystoreStatus_HealthStatus_ConfigurationError = "configuration_error"
	KeystoreStatus_HealthStatus_NotResponding = "not_responding"
	KeystoreStatus_HealthStatus_Ok = "ok"
	KeystoreStatus_HealthStatus_PendingCheck = "pending_check"
)

// UnmarshalKeystoreStatus unmarshals an instance of KeystoreStatus from the specified map of raw messages.
func UnmarshalKeystoreStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreStatus)
	err = core.UnmarshalPrimitive(m, "last_heartbeat", &obj.LastHeartbeat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health_status", &obj.HealthStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequest : Properties of a keystore that can be updated.
// Models which "extend" this model:
// - KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate
// - KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate
// - KeystoreUpdateRequestKeystoreTypeAzureUpdate
// - KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate
// - KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate
// - KeystoreUpdateRequestKeystoreTypeCcaUpdate
type KeystoreUpdateRequest struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials,omitempty"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location,omitempty"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring,omitempty"`

	// AWS Region.
	AwsRegion *string `json:"aws_region,omitempty"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name,omitempty"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group,omitempty"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id,omitempty"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password,omitempty"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant,omitempty"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id,omitempty"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint,omitempty"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint,omitempty"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key,omitempty"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id,omitempty"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls,omitempty"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host,omitempty"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port,omitempty"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash,omitempty"`
}

// Constants associated with the KeystoreUpdateRequest.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	KeystoreUpdateRequest_AzureEnvironment_Azure = "azure"
	KeystoreUpdateRequest_AzureEnvironment_AzureChina = "azure_china"
	KeystoreUpdateRequest_AzureEnvironment_AzureGermany = "azure_germany"
	KeystoreUpdateRequest_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the KeystoreUpdateRequest.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	KeystoreUpdateRequest_AzureVariant_Premium = "premium"
	KeystoreUpdateRequest_AzureVariant_Standard = "standard"
)
func (*KeystoreUpdateRequest) isaKeystoreUpdateRequest() bool {
	return true
}

type KeystoreUpdateRequestIntf interface {
	isaKeystoreUpdateRequest() bool
}

// UnmarshalKeystoreUpdateRequest unmarshals an instance of KeystoreUpdateRequest from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequest)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_credentials", &obj.GoogleCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_location", &obj.GoogleLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_project_id", &obj.GoogleProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_private_key_id", &obj.GooglePrivateKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_ring", &obj.GoogleKeyRing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_region", &obj.AwsRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_access_key_id", &obj.AwsAccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_secret_access_key", &obj.AwsSecretAccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_name", &obj.AzureServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_resource_group", &obj.AzureResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_location", &obj.AzureLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_client_id", &obj.AzureServicePrincipalClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_password", &obj.AzureServicePrincipalPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_tenant", &obj.AzureTenant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_subscription_id", &obj.AzureSubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_environment", &obj.AzureEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_variant", &obj.AzureVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_use_tls", &obj.CcaUseTls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_trusted_issuer", &obj.CcaTrustedIssuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_host", &obj.CcaHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_port", &obj.CcaPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_public_key_hash", &obj.CcaPublicKeyHash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreate : KeystoresPropertiesCreate struct
// Models which "extend" this model:
// - KeystoresPropertiesCreateGoogleKms
// - KeystoresPropertiesCreateAwsKms
// - KeystoresPropertiesCreateIbmCloudKms
// - KeystoresPropertiesCreateAzure
// - KeystoresPropertiesCreateCca
type KeystoresPropertiesCreate struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level,omitempty"`

	GoogleKeyPurpose *string `json:"google_key_purpose,omitempty"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm,omitempty"`

	AzureKeyProtectionLevel *string `json:"azure_key_protection_level,omitempty"`

	AzureKeyOperations []string `json:"azure_key_operations,omitempty"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreate.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreate_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreate_Type_Cca = "cca"
	KeystoresPropertiesCreate_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoresPropertiesCreate.GoogleKeyProtectionLevel property.
const (
	KeystoresPropertiesCreate_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesCreate_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesCreate.GoogleKeyPurpose property.
const (
	KeystoresPropertiesCreate_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeystoresPropertiesCreate_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeystoresPropertiesCreate_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeystoresPropertiesCreate_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeystoresPropertiesCreate.GoogleKmsAlgorithm property.
const (
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeystoresPropertiesCreate_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

// Constants associated with the KeystoresPropertiesCreate.AzureKeyProtectionLevel property.
const (
	KeystoresPropertiesCreate_AzureKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesCreate_AzureKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesCreate.AzureKeyOperations property.
const (
	KeystoresPropertiesCreate_AzureKeyOperations_Decrypt = "decrypt"
	KeystoresPropertiesCreate_AzureKeyOperations_Encrypt = "encrypt"
	KeystoresPropertiesCreate_AzureKeyOperations_Sign = "sign"
	KeystoresPropertiesCreate_AzureKeyOperations_UnwrapKey = "unwrap_key"
	KeystoresPropertiesCreate_AzureKeyOperations_Verify = "verify"
	KeystoresPropertiesCreate_AzureKeyOperations_WrapKey = "wrap_key"
)

// Constants associated with the KeystoresPropertiesCreate.CcaUsageControl property.
const (
	KeystoresPropertiesCreate_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeystoresPropertiesCreate_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeystoresPropertiesCreate_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeystoresPropertiesCreate.CcaKeyType property.
const (
	KeystoresPropertiesCreate_CcaKeyType_Cipher = "cipher"
	KeystoresPropertiesCreate_CcaKeyType_Data = "data"
	KeystoresPropertiesCreate_CcaKeyType_Exporter = "exporter"
	KeystoresPropertiesCreate_CcaKeyType_Importer = "importer"
)
func (*KeystoresPropertiesCreate) isaKeystoresPropertiesCreate() bool {
	return true
}

type KeystoresPropertiesCreateIntf interface {
	isaKeystoresPropertiesCreate() bool
}

// UnmarshalKeystoresPropertiesCreate unmarshals an instance of KeystoresPropertiesCreate from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreate)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_protection_level", &obj.AzureKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_operations", &obj.AzureKeyOperations)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdate : KeystoresPropertiesUpdate struct
// Models which "extend" this model:
// - KeystoresPropertiesUpdateGoogleKms
// - KeystoresPropertiesUpdateAwsKms
// - KeystoresPropertiesUpdateIbmCloudKms
// - KeystoresPropertiesUpdateAzure
// - KeystoresPropertiesUpdateCca
type KeystoresPropertiesUpdate struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level,omitempty"`

	GoogleKeyPurpose *string `json:"google_key_purpose,omitempty"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm,omitempty"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeystoresPropertiesUpdate.GoogleKeyProtectionLevel property.
const (
	KeystoresPropertiesUpdate_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesUpdate_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesUpdate.GoogleKeyPurpose property.
const (
	KeystoresPropertiesUpdate_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeystoresPropertiesUpdate_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeystoresPropertiesUpdate_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeystoresPropertiesUpdate_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeystoresPropertiesUpdate.GoogleKmsAlgorithm property.
const (
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeystoresPropertiesUpdate_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

// Constants associated with the KeystoresPropertiesUpdate.CcaUsageControl property.
const (
	KeystoresPropertiesUpdate_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeystoresPropertiesUpdate_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeystoresPropertiesUpdate_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeystoresPropertiesUpdate.CcaKeyType property.
const (
	KeystoresPropertiesUpdate_CcaKeyType_Cipher = "cipher"
	KeystoresPropertiesUpdate_CcaKeyType_Data = "data"
	KeystoresPropertiesUpdate_CcaKeyType_Exporter = "exporter"
	KeystoresPropertiesUpdate_CcaKeyType_Importer = "importer"
)
func (*KeystoresPropertiesUpdate) isaKeystoresPropertiesUpdate() bool {
	return true
}

type KeystoresPropertiesUpdateIntf interface {
	isaKeystoresPropertiesUpdate() bool
}

// UnmarshalKeystoresPropertiesUpdate unmarshals an instance of KeystoresPropertiesUpdate from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdate)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ManagedKey : A managed key is a key that is created in and assigned to a vault. You can manage the lifecycle of a managed key and
// install it to multiple keystores in the same vault. You can use a managed key for encryption and decryption only when
// it is installed in at least one target keystore. Installing a managed key in multiple keystores in the same vault
// enables key redundancy. To use a managed key for encryption and decryption, install in one or more keystores within
// the same vault first.
type ManagedKey struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// Reference to a key template.
	Template *TemplateReference `json:"template,omitempty"`

	Version *int64 `json:"version,omitempty"`

	// Description of the managed key.
	Description *string `json:"description,omitempty"`

	// The label of the key.
	Label *string `json:"label" validate:"required"`

	// The state of the key.
	State *string `json:"state" validate:"required"`

	// Array of locations where key material is present.
	KeyMaterialPresent []string `json:"key_material_present,omitempty"`

	// The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
	Size *string `json:"size,omitempty"`

	// The algorithm of the key.
	Algorithm *string `json:"algorithm" validate:"required"`

	// A list of verification patterns of the key (e.g. public key hash for RSA keys).
	VerificationPatterns []KeyVerificationPattern `json:"verification_patterns,omitempty"`

	// First day when the key is active.
	ActivationDate *strfmt.Date `json:"activation_date,omitempty"`

	// Last day when the key is active.
	ExpirationDate *strfmt.Date `json:"expiration_date,omitempty"`

	LabelTags []Tag `json:"label_tags,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	// Boolean value which indicates if key can be rotated.
	IsRotatable *bool `json:"is_rotatable,omitempty"`

	// Date and time when the key was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the key was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// referenced keystores.
	ReferencedKeystores []TargetKeystoreReference `json:"referenced_keystores" validate:"required"`

	// key instances.
	Instances []KeyInstanceIntf `json:"instances" validate:"required"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// Date and time when the key was rotated.
	RotatedAt *strfmt.DateTime `json:"rotated_at,omitempty"`

	// list of key statuses in keystores.
	StatusInKeystores []StatusInKeystore `json:"status_in_keystores" validate:"required"`

	// indicates whether to deactivate older versions of a key on rotation.
	DeactivateOnRotation *bool `json:"deactivate_on_rotation,omitempty"`

	// Managing systems of templates and keys.
	ManagingSystems []string `json:"managing_systems,omitempty"`
}

// Constants associated with the ManagedKey.State property.
// The state of the key.
const (
	ManagedKey_State_Active = "active"
	ManagedKey_State_Compromised = "compromised"
	ManagedKey_State_Deactivated = "deactivated"
	ManagedKey_State_Destroyed = "destroyed"
	ManagedKey_State_DestroyedCompromised = "destroyed_compromised"
	ManagedKey_State_PreActivation = "pre_activation"
)

// Constants associated with the ManagedKey.KeyMaterialPresent property.
// Location where key material is present.
const (
	ManagedKey_KeyMaterialPresent_Keystores = "keystores"
	ManagedKey_KeyMaterialPresent_Repository = "repository"
)

// Constants associated with the ManagedKey.Algorithm property.
// The algorithm of the key.
const (
	ManagedKey_Algorithm_Aes = "aes"
	ManagedKey_Algorithm_Des = "des"
	ManagedKey_Algorithm_Dilithium = "dilithium"
	ManagedKey_Algorithm_Ec = "ec"
	ManagedKey_Algorithm_Hmac = "hmac"
	ManagedKey_Algorithm_Rsa = "rsa"
)

// Constants associated with the ManagedKey.ManagingSystems property.
// Managing system of templates and keys.
const (
	ManagedKey_ManagingSystems_Web = "web"
	ManagedKey_ManagingSystems_Workstation = "workstation"
)

// UnmarshalManagedKey unmarshals an instance of ManagedKey from the specified map of raw messages.
func UnmarshalManagedKey(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManagedKey)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "template", &obj.Template, UnmarshalTemplateReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_material_present", &obj.KeyMaterialPresent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "algorithm", &obj.Algorithm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "verification_patterns", &obj.VerificationPatterns, UnmarshalKeyVerificationPattern)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "activation_date", &obj.ActivationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration_date", &obj.ExpirationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "label_tags", &obj.LabelTags, UnmarshalTag)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalTag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_rotatable", &obj.IsRotatable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "referenced_keystores", &obj.ReferencedKeystores, UnmarshalTargetKeystoreReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "instances", &obj.Instances, UnmarshalKeyInstance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rotated_at", &obj.RotatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status_in_keystores", &obj.StatusInKeystores, UnmarshalStatusInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deactivate_on_rotation", &obj.DeactivateOnRotation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managing_systems", &obj.ManagingSystems)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ManagedKeyList : The base schema for listing managed keys.
type ManagedKeyList struct {
	// The total count of all objects in the entire collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The number of resources that were requested in this query.
	Limit *int64 `json:"limit" validate:"required"`

	// The number of resources that were skipped in this query.
	Offset *int64 `json:"offset" validate:"required"`

	// URL of a resource.
	First *HrefObject `json:"first,omitempty"`

	// URL of a resource.
	Last *HrefObject `json:"last,omitempty"`

	// URL of a resource.
	Previous *HrefObject `json:"previous,omitempty"`

	// URL of a resource.
	Next *HrefObject `json:"next,omitempty"`

	// A list of managed keys.
	ManagedKeys []ManagedKey `json:"managed_keys" validate:"required"`
}

// UnmarshalManagedKeyList unmarshals an instance of ManagedKeyList from the specified map of raw messages.
func UnmarshalManagedKeyList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManagedKeyList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "managed_keys", &obj.ManagedKeys, UnmarshalManagedKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ManagedKeyList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// ManagedKeyReference : Reference to a managed key.
type ManagedKeyReference struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// The label of the key.
	Label *string `json:"label" validate:"required"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`
}

// UnmarshalManagedKeyReference unmarshals an instance of ManagedKeyReference from the specified map of raw messages.
func UnmarshalManagedKeyReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ManagedKeyReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusInKeystore : Describes the status of a key in a single keystore.
type StatusInKeystore struct {
	// Reference to a target keystore.
	Keystore *TargetKeystoreReference `json:"keystore,omitempty"`

	Instance *InstanceReference `json:"instance,omitempty"`

	// Possible states of a key in keystore.  not_present: the key is not in the target keystore at all,  active: the key
	// is in the target keystore, and can be used for its intended purpose not_active: the key is in the target keystore,
	// but cannot be used for its intended purpose wrong_key: there is a key in the target keystore, but it doesn't have
	// the value that is expected error: there was an error checking the status of the key in the target keystore.
	Status *string `json:"status" validate:"required"`

	// Destruction policy for this key in this keystore expressed as a duration (e.g. PT24H means 24 hours). If not
	// possible to determine, an empty string is returned.
	DestructionPolicyDuration *string `json:"destruction_policy_duration,omitempty"`

	// If present, indicates when the key will be destroyed in the target keystore, as a timestamp in UTC. If not possible
	// to determine, or does not apply, an empty string is returned.
	DestructionAt *string `json:"destruction_at,omitempty"`

	// Flag to represent synchronization status between UKO Managed Key and Target Keystore. Possible status flags. ok:
	// managed key state is the same as target keystore state, out_of_sync: managed key state is different than target
	// keystore state.
	KeystoreSyncFlag *string `json:"keystore_sync_flag" validate:"required"`

	// Detailed description to represents every possible state combination or mismatch between UKO Managed Key and Target
	// Keystore.
	KeystoreSyncFlagDetail *string `json:"keystore_sync_flag_detail" validate:"required"`

	// An error encountered while using the application.
	Error *ApiError `json:"error,omitempty"`

	// ID of the key in this keystore. Every keystore may use different format such as: UUID, GUID, CRN, URI.
	KeyIdInKeystore *string `json:"key_id_in_keystore,omitempty"`
}

// Constants associated with the StatusInKeystore.Status property.
// Possible states of a key in keystore.  not_present: the key is not in the target keystore at all,  active: the key is
// in the target keystore, and can be used for its intended purpose not_active: the key is in the target keystore, but
// cannot be used for its intended purpose wrong_key: there is a key in the target keystore, but it doesn't have the
// value that is expected error: there was an error checking the status of the key in the target keystore.
const (
	StatusInKeystore_Status_Active = "active"
	StatusInKeystore_Status_Error = "error"
	StatusInKeystore_Status_NotActive = "not_active"
	StatusInKeystore_Status_NotPresent = "not_present"
	StatusInKeystore_Status_PendingDestruction = "pending_destruction"
	StatusInKeystore_Status_Unknown = "unknown"
	StatusInKeystore_Status_WrongKey = "wrong_key"
)

// Constants associated with the StatusInKeystore.KeystoreSyncFlag property.
// Flag to represent synchronization status between UKO Managed Key and Target Keystore. Possible status flags. ok:
// managed key state is the same as target keystore state, out_of_sync: managed key state is different than target
// keystore state.
const (
	StatusInKeystore_KeystoreSyncFlag_Error = "error"
	StatusInKeystore_KeystoreSyncFlag_Ok = "ok"
	StatusInKeystore_KeystoreSyncFlag_OutOfSync = "out_of_sync"
	StatusInKeystore_KeystoreSyncFlag_VerifyingSync = "verifying_sync"
)

// Constants associated with the StatusInKeystore.KeystoreSyncFlagDetail property.
// Detailed description to represents every possible state combination or mismatch between UKO Managed Key and Target
// Keystore.
const (
	StatusInKeystore_KeystoreSyncFlagDetail_ActiveKeyIsActiveInKeystore = "active_key_is_active_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_ActiveKeyIsNotActiveInKeystore = "active_key_is_not_active_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_ConnectionError = "connection_error"
	StatusInKeystore_KeystoreSyncFlagDetail_DeactivatedKeyIsDeactivatedInKeystore = "deactivated_key_is_deactivated_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_DeactivatedKeyIsNotDeactivatedInKeystore = "deactivated_key_is_not_deactivated_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_DestroyedKeyIsNotPresentInKeystore = "destroyed_key_is_not_present_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_DestroyedKeyIsPresentInKeystore = "destroyed_key_is_present_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_DestroyedPendingKeyIsNotPresentInKeystore = "destroyed_pending_key_is_not_present_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_DestroyedPendingKeyIsPendingDestructionInKeystore = "destroyed_pending_key_is_pending_destruction_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_PreActiveKeyIsNotPresentInKeystore = "pre_active_key_is_not_present_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_PreActiveKeyIsPresentInKeystore = "pre_active_key_is_present_in_keystore"
	StatusInKeystore_KeystoreSyncFlagDetail_TargetKeystoreRemovedByUser = "target_keystore_removed_by_user"
	StatusInKeystore_KeystoreSyncFlagDetail_TargetKeystoreRemovedByUserContainsAnActiveKey = "target_keystore_removed_by_user_contains_an_active_key"
	StatusInKeystore_KeystoreSyncFlagDetail_Unknown = "unknown"
)

// UnmarshalStatusInKeystore unmarshals an instance of StatusInKeystore from the specified map of raw messages.
func UnmarshalStatusInKeystore(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusInKeystore)
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalTargetKeystoreReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "instance", &obj.Instance, UnmarshalInstanceReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destruction_policy_duration", &obj.DestructionPolicyDuration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destruction_at", &obj.DestructionAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keystore_sync_flag", &obj.KeystoreSyncFlag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keystore_sync_flag_detail", &obj.KeystoreSyncFlagDetail)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalApiError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_id_in_keystore", &obj.KeyIdInKeystore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusInKeystores : Status of a key in keystores.
type StatusInKeystores struct {
	// list of key statuses in keystores.
	StatusInKeystores []StatusInKeystore `json:"status_in_keystores" validate:"required"`
}

// UnmarshalStatusInKeystores unmarshals an instance of StatusInKeystores from the specified map of raw messages.
func UnmarshalStatusInKeystores(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusInKeystores)
	err = core.UnmarshalModel(m, "status_in_keystores", &obj.StatusInKeystores, UnmarshalStatusInKeystore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tag : A key tag, as used when creating keys and searching.
type Tag struct {
	// Name of a tag.
	Name *string `json:"name" validate:"required"`

	// Value of a tag.
	Value *string `json:"value" validate:"required"`
}

// NewTag : Instantiate Tag (Generic Model Constructor)
func (*UkoV4) NewTag(name string, value string) (_model *Tag, err error) {
	_model = &Tag{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTag unmarshals an instance of Tag from the specified map of raw messages.
func UnmarshalTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tag)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Target : Target of the error.
type Target struct {
	// Type, one of 'field', 'parameter', or 'header'.
	Type *string `json:"type,omitempty"`

	// Name of the field (with dot-syntax if necessary), query parameter, or header.
	Name *string `json:"name,omitempty"`
}

// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetKeystoreReference : Reference to a target keystore.
type TargetKeystoreReference struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore.
	Name *string `json:"name,omitempty"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`
}

// Constants associated with the TargetKeystoreReference.Type property.
// Type of keystore.
const (
	TargetKeystoreReference_Type_AwsKms = "aws_kms"
	TargetKeystoreReference_Type_AzureKeyVault = "azure_key_vault"
	TargetKeystoreReference_Type_Cca = "cca"
	TargetKeystoreReference_Type_GoogleKms = "google_kms"
	TargetKeystoreReference_Type_IbmCloudKms = "ibm_cloud_kms"
)

// UnmarshalTargetKeystoreReference unmarshals an instance of TargetKeystoreReference from the specified map of raw messages.
func UnmarshalTargetKeystoreReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetKeystoreReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Template : A template and all its properties.
type Template struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Version of the key template. Every time the key template is updated, the version will be updated automatically.
	Version *int64 `json:"version,omitempty"`

	// Name of the key template.
	Name *string `json:"name,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	Type []string `json:"type" validate:"required"`

	// State of the template which determines if the template is archived or unarchived.
	State *string `json:"state" validate:"required"`

	// The total count of keys created with this template.
	KeysCount *int64 `json:"keys_count" validate:"required"`

	// Properties describing the properties of the managed key.
	Key *KeyProperties `json:"key" validate:"required"`

	// Description of the key template.
	Description *string `json:"description" validate:"required"`

	// Date and time when the key template was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the key template was updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key template.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	Keystores []KeystoresPropertiesCreateIntf `json:"keystores" validate:"required"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// Managing systems of templates and keys.
	ManagingSystems []string `json:"managing_systems,omitempty"`
}

// Constants associated with the Template.Type property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	Template_Type_Shadow = "shadow"
	Template_Type_System = "system"
	Template_Type_UserDefined = "user_defined"
)

// Constants associated with the Template.State property.
// State of the template which determines if the template is archived or unarchived.
const (
	Template_State_Archived = "archived"
	Template_State_Unarchived = "unarchived"
)

// Constants associated with the Template.ManagingSystems property.
// Managing system of templates and keys.
const (
	Template_ManagingSystems_Web = "web"
	Template_ManagingSystems_Workstation = "workstation"
)

// UnmarshalTemplate unmarshals an instance of Template from the specified map of raw messages.
func UnmarshalTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Template)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keys_count", &obj.KeysCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "key", &obj.Key, UnmarshalKeyProperties)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystores", &obj.Keystores, UnmarshalKeystoresPropertiesCreate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managing_systems", &obj.ManagingSystems)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TemplateList : The base schema for listing key templates.
type TemplateList struct {
	// The total count of all objects in the entire collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The number of resources that were requested in this query.
	Limit *int64 `json:"limit" validate:"required"`

	// The number of resources that were skipped in this query.
	Offset *int64 `json:"offset" validate:"required"`

	// URL of a resource.
	First *HrefObject `json:"first,omitempty"`

	// URL of a resource.
	Last *HrefObject `json:"last,omitempty"`

	// URL of a resource.
	Previous *HrefObject `json:"previous,omitempty"`

	// URL of a resource.
	Next *HrefObject `json:"next,omitempty"`

	// A list of key templates.
	Templates []Template `json:"templates" validate:"required"`
}

// UnmarshalTemplateList unmarshals an instance of TemplateList from the specified map of raw messages.
func UnmarshalTemplateList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "templates", &obj.Templates, UnmarshalTemplate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *TemplateList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// TemplateReference : Reference to a key template.
type TemplateReference struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the key template.
	Name *string `json:"name,omitempty"`

	Type []string `json:"type,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// Status describing if a key is aligned with the latest key template version.
	AlignmentStatus *string `json:"alignment_status" validate:"required"`
}

// Constants associated with the TemplateReference.Type property.
// Type of the template which determines template origins. Ones created by user are 'user_defined' where 'shadow' means
// it was created under the hood by the UKO.
const (
	TemplateReference_Type_Shadow = "shadow"
	TemplateReference_Type_System = "system"
	TemplateReference_Type_UserDefined = "user_defined"
)

// Constants associated with the TemplateReference.AlignmentStatus property.
// Status describing if a key is aligned with the latest key template version.
const (
	TemplateReference_AlignmentStatus_Aligned = "aligned"
	TemplateReference_AlignmentStatus_Unaligned = "unaligned"
)

// UnmarshalTemplateReference unmarshals an instance of TemplateReference from the specified map of raw messages.
func UnmarshalTemplateReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alignment_status", &obj.AlignmentStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Vault : Vaults are secure repositories for your cryptographic keys and keystores. A managed key or keystore can only be in
// one vault at a time.
type Vault struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the vault.
	Name *string `json:"name" validate:"required"`

	// Description of the vault.
	Description *string `json:"description" validate:"required"`

	// Date and time when the vault was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the vault was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the vault.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the vault.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// The label of the recovery key for this vault.
	RecoveryKeyLabel *string `json:"recovery_key_label,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The number of keys inside the vault.
	KeysCount *int64 `json:"keys_count" validate:"required"`

	// The number of key templates inside the vault.
	KeyTemplatesCount *int64 `json:"key_templates_count" validate:"required"`

	// The number of keystores inside the vault.
	KeystoresCount *int64 `json:"keystores_count" validate:"required"`
}

// UnmarshalVault unmarshals an instance of Vault from the specified map of raw messages.
func UnmarshalVault(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Vault)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "recovery_key_label", &obj.RecoveryKeyLabel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keys_count", &obj.KeysCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_templates_count", &obj.KeyTemplatesCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keystores_count", &obj.KeystoresCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultList : The base schema for listing vaults.
type VaultList struct {
	// The total count of all objects in the entire collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The number of resources that were requested in this query.
	Limit *int64 `json:"limit" validate:"required"`

	// The number of resources that were skipped in this query.
	Offset *int64 `json:"offset" validate:"required"`

	// URL of a resource.
	First *HrefObject `json:"first,omitempty"`

	// URL of a resource.
	Last *HrefObject `json:"last,omitempty"`

	// URL of a resource.
	Previous *HrefObject `json:"previous,omitempty"`

	// URL of a resource.
	Next *HrefObject `json:"next,omitempty"`

	// A list of vaults.
	Vaults []Vault `json:"vaults" validate:"required"`
}

// UnmarshalVaultList unmarshals an instance of VaultList from the specified map of raw messages.
func UnmarshalVaultList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vaults", &obj.Vaults, UnmarshalVault)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *VaultList) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// VaultReference : Reference to a vault.
type VaultReference struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the referenced vault.
	Name *string `json:"name,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`
}

// UnmarshalVaultReference unmarshals an instance of VaultReference from the specified map of raw messages.
func UnmarshalVaultReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VaultReferenceInCreationRequest : ID of the Vault where the entity is to be created in.
type VaultReferenceInCreationRequest struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`
}

// NewVaultReferenceInCreationRequest : Instantiate VaultReferenceInCreationRequest (Generic Model Constructor)
func (*UkoV4) NewVaultReferenceInCreationRequest(id string) (_model *VaultReferenceInCreationRequest, err error) {
	_model = &VaultReferenceInCreationRequest{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalVaultReferenceInCreationRequest unmarshals an instance of VaultReferenceInCreationRequest from the specified map of raw messages.
func UnmarshalVaultReferenceInCreationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VaultReferenceInCreationRequest)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstanceAwsKms : The instance of a managed key for a specific keystore.
// This model "extends" KeyInstance
type KeyInstanceAwsKms struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore" validate:"required"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore" validate:"required"`
}

// Constants associated with the KeyInstanceAwsKms.Type property.
// Type of the key instance.
const (
	KeyInstanceAwsKms_Type_KeyPair = "key_pair"
	KeyInstanceAwsKms_Type_PrivateKey = "private_key"
	KeyInstanceAwsKms_Type_PublicKey = "public_key"
	KeyInstanceAwsKms_Type_SecretKey = "secret_key"
)

func (*KeyInstanceAwsKms) isaKeyInstance() bool {
	return true
}

// UnmarshalKeyInstanceAwsKms unmarshals an instance of KeyInstanceAwsKms from the specified map of raw messages.
func UnmarshalKeyInstanceAwsKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstanceAwsKms)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstanceAzure : The instance of a managed key for a specific keystore.
// This model "extends" KeyInstance
type KeyInstanceAzure struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore" validate:"required"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore" validate:"required"`

	AzureKeyProtectionLevel *string `json:"azure_key_protection_level" validate:"required"`

	AzureKeyOperations []string `json:"azure_key_operations,omitempty"`
}

// Constants associated with the KeyInstanceAzure.Type property.
// Type of the key instance.
const (
	KeyInstanceAzure_Type_KeyPair = "key_pair"
	KeyInstanceAzure_Type_PrivateKey = "private_key"
	KeyInstanceAzure_Type_PublicKey = "public_key"
	KeyInstanceAzure_Type_SecretKey = "secret_key"
)

// Constants associated with the KeyInstanceAzure.AzureKeyProtectionLevel property.
const (
	KeyInstanceAzure_AzureKeyProtectionLevel_Hsm = "hsm"
	KeyInstanceAzure_AzureKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeyInstanceAzure.AzureKeyOperations property.
const (
	KeyInstanceAzure_AzureKeyOperations_Decrypt = "decrypt"
	KeyInstanceAzure_AzureKeyOperations_Encrypt = "encrypt"
	KeyInstanceAzure_AzureKeyOperations_Sign = "sign"
	KeyInstanceAzure_AzureKeyOperations_UnwrapKey = "unwrap_key"
	KeyInstanceAzure_AzureKeyOperations_Verify = "verify"
	KeyInstanceAzure_AzureKeyOperations_WrapKey = "wrap_key"
)

func (*KeyInstanceAzure) isaKeyInstance() bool {
	return true
}

// UnmarshalKeyInstanceAzure unmarshals an instance of KeyInstanceAzure from the specified map of raw messages.
func UnmarshalKeyInstanceAzure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstanceAzure)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_protection_level", &obj.AzureKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_operations", &obj.AzureKeyOperations)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstanceCca : The instance of a managed key for a specific keystore.
// This model "extends" KeyInstance
type KeyInstanceCca struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore" validate:"required"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore" validate:"required"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeyInstanceCca.Type property.
// Type of the key instance.
const (
	KeyInstanceCca_Type_KeyPair = "key_pair"
	KeyInstanceCca_Type_PrivateKey = "private_key"
	KeyInstanceCca_Type_PublicKey = "public_key"
	KeyInstanceCca_Type_SecretKey = "secret_key"
)

// Constants associated with the KeyInstanceCca.CcaUsageControl property.
const (
	KeyInstanceCca_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeyInstanceCca_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeyInstanceCca_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeyInstanceCca.CcaKeyType property.
const (
	KeyInstanceCca_CcaKeyType_Cipher = "cipher"
	KeyInstanceCca_CcaKeyType_Data = "data"
	KeyInstanceCca_CcaKeyType_Exporter = "exporter"
	KeyInstanceCca_CcaKeyType_Importer = "importer"
)

func (*KeyInstanceCca) isaKeyInstance() bool {
	return true
}

// UnmarshalKeyInstanceCca unmarshals an instance of KeyInstanceCca from the specified map of raw messages.
func UnmarshalKeyInstanceCca(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstanceCca)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstanceGoogleKms : The instance of a managed key for a specific keystore.
// This model "extends" KeyInstance
type KeyInstanceGoogleKms struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore" validate:"required"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore" validate:"required"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level" validate:"required"`

	GoogleKeyPurpose *string `json:"google_key_purpose" validate:"required"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm" validate:"required"`
}

// Constants associated with the KeyInstanceGoogleKms.Type property.
// Type of the key instance.
const (
	KeyInstanceGoogleKms_Type_KeyPair = "key_pair"
	KeyInstanceGoogleKms_Type_PrivateKey = "private_key"
	KeyInstanceGoogleKms_Type_PublicKey = "public_key"
	KeyInstanceGoogleKms_Type_SecretKey = "secret_key"
)

// Constants associated with the KeyInstanceGoogleKms.GoogleKeyProtectionLevel property.
const (
	KeyInstanceGoogleKms_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeyInstanceGoogleKms_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeyInstanceGoogleKms.GoogleKeyPurpose property.
const (
	KeyInstanceGoogleKms_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeyInstanceGoogleKms_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeyInstanceGoogleKms_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeyInstanceGoogleKms_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeyInstanceGoogleKms.GoogleKmsAlgorithm property.
const (
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeyInstanceGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

func (*KeyInstanceGoogleKms) isaKeyInstance() bool {
	return true
}

// UnmarshalKeyInstanceGoogleKms unmarshals an instance of KeyInstanceGoogleKms from the specified map of raw messages.
func UnmarshalKeyInstanceGoogleKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstanceGoogleKms)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeyInstanceIbmCloudKms : The instance of a managed key for a specific keystore.
// This model "extends" KeyInstance
type KeyInstanceIbmCloudKms struct {
	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id" validate:"required"`

	// The label of the key.
	LabelInKeystore *string `json:"label_in_keystore" validate:"required"`

	// Type of the key instance.
	Type *string `json:"type,omitempty"`

	// Description of properties of a key within the context of keystores.
	Keystore *InstanceInKeystore `json:"keystore" validate:"required"`
}

// Constants associated with the KeyInstanceIbmCloudKms.Type property.
// Type of the key instance.
const (
	KeyInstanceIbmCloudKms_Type_KeyPair = "key_pair"
	KeyInstanceIbmCloudKms_Type_PrivateKey = "private_key"
	KeyInstanceIbmCloudKms_Type_PublicKey = "public_key"
	KeyInstanceIbmCloudKms_Type_SecretKey = "secret_key"
)

func (*KeyInstanceIbmCloudKms) isaKeyInstance() bool {
	return true
}

// UnmarshalKeyInstanceIbmCloudKms unmarshals an instance of KeyInstanceIbmCloudKms from the specified map of raw messages.
func UnmarshalKeyInstanceIbmCloudKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeyInstanceIbmCloudKms)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_in_keystore", &obj.LabelInKeystore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "keystore", &obj.Keystore, UnmarshalInstanceInKeystore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeAwsKmsCreate : AWS KMS is a managed service for you to create and manage cryptographic keys across a wide range of AWS services.
// This model "extends" KeystoreCreationRequest
type KeystoreCreationRequestKeystoreTypeAwsKmsCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// AWS Region.
	AwsRegion *string `json:"aws_region" validate:"required"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id" validate:"required"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key" validate:"required"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeAwsKmsCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeAwsKmsCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeAwsKmsCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeAwsKmsCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeAwsKmsCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeAwsKmsCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// NewKeystoreCreationRequestKeystoreTypeAwsKmsCreate : Instantiate KeystoreCreationRequestKeystoreTypeAwsKmsCreate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeAwsKmsCreate(typeVar string, vault *VaultReferenceInCreationRequest, name string, awsRegion string, awsAccessKeyID string, awsSecretAccessKey string) (_model *KeystoreCreationRequestKeystoreTypeAwsKmsCreate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeAwsKmsCreate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		Name: core.StringPtr(name),
		AwsRegion: core.StringPtr(awsRegion),
		AwsAccessKeyID: core.StringPtr(awsAccessKeyID),
		AwsSecretAccessKey: core.StringPtr(awsSecretAccessKey),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeAwsKmsCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeAwsKmsCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeAwsKmsCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeAwsKmsCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_region", &obj.AwsRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_access_key_id", &obj.AwsAccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_secret_access_key", &obj.AwsSecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeAzureCreate : Microsoft Azure Key Vault is a cloud service for you to create and manage cryptographic keys and other sensitive
// information.
// This model "extends" KeystoreCreationRequest
type KeystoreCreationRequestKeystoreTypeAzureCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name" validate:"required"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group" validate:"required"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id" validate:"required"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password" validate:"required"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant" validate:"required"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id" validate:"required"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeAzureCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeAzureCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeAzureCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeAzureCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeAzureCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeAzureCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeAzureCreate.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureEnvironment_Azure = "azure"
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureEnvironment_AzureChina = "azure_china"
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureEnvironment_AzureGermany = "azure_germany"
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeAzureCreate.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureVariant_Premium = "premium"
	KeystoreCreationRequestKeystoreTypeAzureCreate_AzureVariant_Standard = "standard"
)

// NewKeystoreCreationRequestKeystoreTypeAzureCreate : Instantiate KeystoreCreationRequestKeystoreTypeAzureCreate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeAzureCreate(typeVar string, vault *VaultReferenceInCreationRequest, azureServiceName string, azureResourceGroup string, azureServicePrincipalClientID string, azureServicePrincipalPassword string, azureTenant string, azureSubscriptionID string) (_model *KeystoreCreationRequestKeystoreTypeAzureCreate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeAzureCreate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		AzureServiceName: core.StringPtr(azureServiceName),
		AzureResourceGroup: core.StringPtr(azureResourceGroup),
		AzureServicePrincipalClientID: core.StringPtr(azureServicePrincipalClientID),
		AzureServicePrincipalPassword: core.StringPtr(azureServicePrincipalPassword),
		AzureTenant: core.StringPtr(azureTenant),
		AzureSubscriptionID: core.StringPtr(azureSubscriptionID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeAzureCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeAzureCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeAzureCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeAzureCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeAzureCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_name", &obj.AzureServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_resource_group", &obj.AzureResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_location", &obj.AzureLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_client_id", &obj.AzureServicePrincipalClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_password", &obj.AzureServicePrincipalPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_tenant", &obj.AzureTenant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_subscription_id", &obj.AzureSubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_environment", &obj.AzureEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_variant", &obj.AzureVariant)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeCcaCreate : Properties related to CCA keystore.
// This model "extends" KeystoreCreationRequest
type KeystoreCreationRequestKeystoreTypeCcaCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls,omitempty"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host" validate:"required"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port" validate:"required"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash" validate:"required"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeCcaCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeCcaCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeCcaCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeCcaCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeCcaCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeCcaCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// NewKeystoreCreationRequestKeystoreTypeCcaCreate : Instantiate KeystoreCreationRequestKeystoreTypeCcaCreate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeCcaCreate(typeVar string, vault *VaultReferenceInCreationRequest, name string, ccaHost string, ccaPort int64, ccaPublicKeyHash string) (_model *KeystoreCreationRequestKeystoreTypeCcaCreate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeCcaCreate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		Name: core.StringPtr(name),
		CcaHost: core.StringPtr(ccaHost),
		CcaPort: core.Int64Ptr(ccaPort),
		CcaPublicKeyHash: core.StringPtr(ccaPublicKeyHash),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeCcaCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeCcaCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeCcaCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeCcaCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeCcaCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_use_tls", &obj.CcaUseTls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_trusted_issuer", &obj.CcaTrustedIssuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_host", &obj.CcaHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_port", &obj.CcaPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_public_key_hash", &obj.CcaPublicKeyHash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeGoogleKmsCreate : Google Cloud KMS is a managed service for you to create and manage cryptographic keys across a wide range of Google
// Cloud services.
// This model "extends" KeystoreCreationRequest
type KeystoreCreationRequestKeystoreTypeGoogleKmsCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials" validate:"required"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location,omitempty"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeGoogleKmsCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeGoogleKmsCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeGoogleKmsCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeGoogleKmsCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeGoogleKmsCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeGoogleKmsCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// NewKeystoreCreationRequestKeystoreTypeGoogleKmsCreate : Instantiate KeystoreCreationRequestKeystoreTypeGoogleKmsCreate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeGoogleKmsCreate(typeVar string, vault *VaultReferenceInCreationRequest, name string, googleCredentials string) (_model *KeystoreCreationRequestKeystoreTypeGoogleKmsCreate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeGoogleKmsCreate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		Name: core.StringPtr(name),
		GoogleCredentials: core.StringPtr(googleCredentials),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeGoogleKmsCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeGoogleKmsCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeGoogleKmsCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeGoogleKmsCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeGoogleKmsCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_credentials", &obj.GoogleCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_location", &obj.GoogleLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_project_id", &obj.GoogleProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_private_key_id", &obj.GooglePrivateKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_ring", &obj.GoogleKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate : Properties required to create an IBM Cloud keystore.
// Models which "extend" this model:
// - KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate
// - KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate
// This model "extends" KeystoreCreationRequest
type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint,omitempty"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint,omitempty"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key,omitempty"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id,omitempty"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_IbmVariant_Internal = "internal"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate_IbmVariant_KeyProtect = "key_protect"
)
func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool {
	return true
}

type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateIntf interface {
	KeystoreCreationRequestIntf
	isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreTypeAwsKms : Properties related to AWS KMS.
// This model "extends" Keystore
type KeystoreTypeAwsKms struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name" validate:"required"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description" validate:"required"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status" validate:"required"`

	// AWS Region.
	AwsRegion *string `json:"aws_region" validate:"required"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id" validate:"required"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key" validate:"required"`
}

// Constants associated with the KeystoreTypeAwsKms.Type property.
// Type of keystore.
const (
	KeystoreTypeAwsKms_Type_AwsKms = "aws_kms"
	KeystoreTypeAwsKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoreTypeAwsKms_Type_Cca = "cca"
	KeystoreTypeAwsKms_Type_GoogleKms = "google_kms"
	KeystoreTypeAwsKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

func (*KeystoreTypeAwsKms) isaKeystore() bool {
	return true
}

// UnmarshalKeystoreTypeAwsKms unmarshals an instance of KeystoreTypeAwsKms from the specified map of raw messages.
func UnmarshalKeystoreTypeAwsKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreTypeAwsKms)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_region", &obj.AwsRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_access_key_id", &obj.AwsAccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_secret_access_key", &obj.AwsSecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreTypeAzure : Proxy for connecting to keystore.
// This model "extends" Keystore
type KeystoreTypeAzure struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name" validate:"required"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description" validate:"required"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status" validate:"required"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name" validate:"required"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group" validate:"required"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id" validate:"required"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password" validate:"required"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant" validate:"required"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id" validate:"required"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`
}

// Constants associated with the KeystoreTypeAzure.Type property.
// Type of keystore.
const (
	KeystoreTypeAzure_Type_AwsKms = "aws_kms"
	KeystoreTypeAzure_Type_AzureKeyVault = "azure_key_vault"
	KeystoreTypeAzure_Type_Cca = "cca"
	KeystoreTypeAzure_Type_GoogleKms = "google_kms"
	KeystoreTypeAzure_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreTypeAzure.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	KeystoreTypeAzure_AzureEnvironment_Azure = "azure"
	KeystoreTypeAzure_AzureEnvironment_AzureChina = "azure_china"
	KeystoreTypeAzure_AzureEnvironment_AzureGermany = "azure_germany"
	KeystoreTypeAzure_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the KeystoreTypeAzure.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	KeystoreTypeAzure_AzureVariant_Premium = "premium"
	KeystoreTypeAzure_AzureVariant_Standard = "standard"
)

func (*KeystoreTypeAzure) isaKeystore() bool {
	return true
}

// UnmarshalKeystoreTypeAzure unmarshals an instance of KeystoreTypeAzure from the specified map of raw messages.
func UnmarshalKeystoreTypeAzure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreTypeAzure)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_name", &obj.AzureServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_resource_group", &obj.AzureResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_location", &obj.AzureLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_client_id", &obj.AzureServicePrincipalClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_password", &obj.AzureServicePrincipalPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_tenant", &obj.AzureTenant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_subscription_id", &obj.AzureSubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_environment", &obj.AzureEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_variant", &obj.AzureVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreTypeCca : Properties related to CCA keystore.
// This model "extends" Keystore
type KeystoreTypeCca struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name" validate:"required"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description" validate:"required"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status" validate:"required"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls" validate:"required"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host" validate:"required"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port" validate:"required"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash" validate:"required"`
}

// Constants associated with the KeystoreTypeCca.Type property.
// Type of keystore.
const (
	KeystoreTypeCca_Type_AwsKms = "aws_kms"
	KeystoreTypeCca_Type_AzureKeyVault = "azure_key_vault"
	KeystoreTypeCca_Type_Cca = "cca"
	KeystoreTypeCca_Type_GoogleKms = "google_kms"
	KeystoreTypeCca_Type_IbmCloudKms = "ibm_cloud_kms"
)

func (*KeystoreTypeCca) isaKeystore() bool {
	return true
}

// UnmarshalKeystoreTypeCca unmarshals an instance of KeystoreTypeCca from the specified map of raw messages.
func UnmarshalKeystoreTypeCca(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreTypeCca)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_use_tls", &obj.CcaUseTls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_trusted_issuer", &obj.CcaTrustedIssuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_host", &obj.CcaHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_port", &obj.CcaPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_public_key_hash", &obj.CcaPublicKeyHash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreTypeGoogleKms : Properties related to Google Cloud KMS.
// This model "extends" Keystore
type KeystoreTypeGoogleKms struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name" validate:"required"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description" validate:"required"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status" validate:"required"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials" validate:"required"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location" validate:"required"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring" validate:"required"`
}

// Constants associated with the KeystoreTypeGoogleKms.Type property.
// Type of keystore.
const (
	KeystoreTypeGoogleKms_Type_AwsKms = "aws_kms"
	KeystoreTypeGoogleKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoreTypeGoogleKms_Type_Cca = "cca"
	KeystoreTypeGoogleKms_Type_GoogleKms = "google_kms"
	KeystoreTypeGoogleKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

func (*KeystoreTypeGoogleKms) isaKeystore() bool {
	return true
}

// UnmarshalKeystoreTypeGoogleKms unmarshals an instance of KeystoreTypeGoogleKms from the specified map of raw messages.
func UnmarshalKeystoreTypeGoogleKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreTypeGoogleKms)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_credentials", &obj.GoogleCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_location", &obj.GoogleLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_project_id", &obj.GoogleProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_private_key_id", &obj.GooglePrivateKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_ring", &obj.GoogleKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreTypeIbmCloudKms : properties of a ibm cloud keystore.
// This model "extends" Keystore
type KeystoreTypeIbmCloudKms struct {
	// Reference to a vault.
	Vault *VaultReference `json:"vault,omitempty"`

	// The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	ID *string `json:"id,omitempty"`

	// Name of the target keystore. It can be changed in the future.
	Name *string `json:"name" validate:"required"`

	// Geographic location of the keystore, if available.
	Location *string `json:"location" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description" validate:"required"`

	// List of groups that this keystore belongs to.
	Groups []string `json:"groups" validate:"required"`

	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	// Date and time when the target keystore was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Date and time when the target keystore was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// ID of the user that created the key.
	CreatedBy *string `json:"created_by,omitempty"`

	// ID of the user that last updated the key.
	UpdatedBy *string `json:"updated_by,omitempty"`

	// A URL that uniquely identifies your cloud resource.
	Href *string `json:"href,omitempty"`

	// The status of the connection to the keystore.
	Status *KeystoreStatus `json:"status" validate:"required"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint" validate:"required"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint" validate:"required"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key" validate:"required"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`
}

// Constants associated with the KeystoreTypeIbmCloudKms.Type property.
// Type of keystore.
const (
	KeystoreTypeIbmCloudKms_Type_AwsKms = "aws_kms"
	KeystoreTypeIbmCloudKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoreTypeIbmCloudKms_Type_Cca = "cca"
	KeystoreTypeIbmCloudKms_Type_GoogleKms = "google_kms"
	KeystoreTypeIbmCloudKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreTypeIbmCloudKms.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreTypeIbmCloudKms_IbmVariant_Hpcs = "hpcs"
	KeystoreTypeIbmCloudKms_IbmVariant_Internal = "internal"
	KeystoreTypeIbmCloudKms_IbmVariant_KeyProtect = "key_protect"
)

func (*KeystoreTypeIbmCloudKms) isaKeystore() bool {
	return true
}

// UnmarshalKeystoreTypeIbmCloudKms unmarshals an instance of KeystoreTypeIbmCloudKms from the specified map of raw messages.
func UnmarshalKeystoreTypeIbmCloudKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreTypeIbmCloudKms)
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalKeystoreStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate : AWS KMS is a managed service for you to create and manage cryptographic keys across a wide range of AWS services.
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// AWS Region.
	AwsRegion *string `json:"aws_region,omitempty"`

	// The access key id used for connecting to this instance of AWS KMS.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`

	// The secret access key used for connecting to this instance of AWS KMS.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
}

func (*KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeAwsKmsUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeAwsKmsUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_region", &obj.AwsRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_access_key_id", &obj.AwsAccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "aws_secret_access_key", &obj.AwsSecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeAzureUpdate : Microsoft Azure Key Vault is a cloud service for you to create and manage cryptographic keys and other sensitive
// information.
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeAzureUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// Service name of the key vault instance from the Azure portal.
	AzureServiceName *string `json:"azure_service_name,omitempty"`

	// Resource group in Azure.
	AzureResourceGroup *string `json:"azure_resource_group,omitempty"`

	// Location of the Azure Key Vault.
	AzureLocation *string `json:"azure_location,omitempty"`

	// Azure service principal client ID.
	AzureServicePrincipalClientID *string `json:"azure_service_principal_client_id,omitempty"`

	// Azure service principal password.
	AzureServicePrincipalPassword *string `json:"azure_service_principal_password,omitempty"`

	// Azure tenant that the Key Vault is associated with,.
	AzureTenant *string `json:"azure_tenant,omitempty"`

	// Subscription ID in Azure.
	AzureSubscriptionID *string `json:"azure_subscription_id,omitempty"`

	// Azure environment, usually 'Azure'.
	AzureEnvironment *string `json:"azure_environment,omitempty"`

	// Variant of the Azure Key Vault.
	AzureVariant *string `json:"azure_variant,omitempty"`
}

// Constants associated with the KeystoreUpdateRequestKeystoreTypeAzureUpdate.AzureEnvironment property.
// Azure environment, usually 'Azure'.
const (
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureEnvironment_Azure = "azure"
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureEnvironment_AzureChina = "azure_china"
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureEnvironment_AzureGermany = "azure_germany"
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureEnvironment_AzureUsGovernment = "azure_us_government"
)

// Constants associated with the KeystoreUpdateRequestKeystoreTypeAzureUpdate.AzureVariant property.
// Variant of the Azure Key Vault.
const (
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureVariant_Premium = "premium"
	KeystoreUpdateRequestKeystoreTypeAzureUpdate_AzureVariant_Standard = "standard"
)

func (*KeystoreUpdateRequestKeystoreTypeAzureUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeAzureUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeAzureUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeAzureUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeAzureUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_name", &obj.AzureServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_resource_group", &obj.AzureResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_location", &obj.AzureLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_client_id", &obj.AzureServicePrincipalClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_service_principal_password", &obj.AzureServicePrincipalPassword)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_tenant", &obj.AzureTenant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_subscription_id", &obj.AzureSubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_environment", &obj.AzureEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_variant", &obj.AzureVariant)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeCcaUpdate : Properties related to CCA keystore.
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeCcaUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// indicates whether to use TLS when connecting to an EKMF agent.
	CcaUseTls *bool `json:"cca_use_tls,omitempty"`

	// Base64 encoded PEM representation of a trusted issuer when using TLS.
	CcaTrustedIssuer *string `json:"cca_trusted_issuer,omitempty"`

	// a host of the keystore.
	CcaHost *string `json:"cca_host,omitempty"`

	// a port of the keystore.
	CcaPort *int64 `json:"cca_port,omitempty"`

	// HEX encoded string contained hash of signature key.
	CcaPublicKeyHash *string `json:"cca_public_key_hash,omitempty"`
}

func (*KeystoreUpdateRequestKeystoreTypeCcaUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeCcaUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeCcaUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeCcaUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeCcaUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_use_tls", &obj.CcaUseTls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_trusted_issuer", &obj.CcaTrustedIssuer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_host", &obj.CcaHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_port", &obj.CcaPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_public_key_hash", &obj.CcaPublicKeyHash)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate : Google Cloud KMS is a managed service for you to create and manage cryptographic keys across a wide range of Google
// Cloud services.
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// The value of the JSON key represented in the Base64 format.
	GoogleCredentials *string `json:"google_credentials,omitempty"`

	// Location represents the geographical region where a Cloud KMS resource is stored and can be accessed. A key's
	// location impacts the performance of applications using the key.
	GoogleLocation *string `json:"google_location,omitempty"`

	// The project id associated with this keystore.
	GoogleProjectID *string `json:"google_project_id,omitempty"`

	// The private key id associated with this keystore.
	GooglePrivateKeyID *string `json:"google_private_key_id,omitempty"`

	// A key ring organizes keys in a specific Google Cloud location and allows you to manage access control on groups of
	// keys.
	GoogleKeyRing *string `json:"google_key_ring,omitempty"`
}

func (*KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_credentials", &obj.GoogleCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_location", &obj.GoogleLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_project_id", &obj.GoogleProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_private_key_id", &obj.GooglePrivateKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_ring", &obj.GoogleKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate : KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate struct
// Models which "extend" this model:
// - KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`
}
func (*KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate() bool {
	return true
}

type KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateIntf interface {
	KeystoreUpdateRequestIntf
	isaKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate() bool
}

func (*KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate : You can connect your Hyper Protect Crypto Services instance to the keystores of another Hyper Protect Crypto Services
// or Key Protect instance, and manage its KMS keys using the current service instance.
// This model "extends" KeystoreUpdateRequest
type KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint,omitempty"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint,omitempty"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key,omitempty"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id,omitempty"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`
}

func (*KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeIbmCloudKmsUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreateAwsKms : KeystoresPropertiesCreateAwsKms struct
// This model "extends" KeystoresPropertiesCreate
type KeystoresPropertiesCreateAwsKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreateAwsKms.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreateAwsKms_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreateAwsKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreateAwsKms_Type_Cca = "cca"
	KeystoresPropertiesCreateAwsKms_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreateAwsKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

func (*KeystoresPropertiesCreateAwsKms) isaKeystoresPropertiesCreate() bool {
	return true
}

// UnmarshalKeystoresPropertiesCreateAwsKms unmarshals an instance of KeystoresPropertiesCreateAwsKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreateAwsKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreateAwsKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreateAzure : KeystoresPropertiesCreateAzure struct
// This model "extends" KeystoresPropertiesCreate
type KeystoresPropertiesCreateAzure struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`

	AzureKeyProtectionLevel *string `json:"azure_key_protection_level,omitempty"`

	AzureKeyOperations []string `json:"azure_key_operations,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreateAzure.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreateAzure_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreateAzure_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreateAzure_Type_Cca = "cca"
	KeystoresPropertiesCreateAzure_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreateAzure_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoresPropertiesCreateAzure.AzureKeyProtectionLevel property.
const (
	KeystoresPropertiesCreateAzure_AzureKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesCreateAzure_AzureKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesCreateAzure.AzureKeyOperations property.
const (
	KeystoresPropertiesCreateAzure_AzureKeyOperations_Decrypt = "decrypt"
	KeystoresPropertiesCreateAzure_AzureKeyOperations_Encrypt = "encrypt"
	KeystoresPropertiesCreateAzure_AzureKeyOperations_Sign = "sign"
	KeystoresPropertiesCreateAzure_AzureKeyOperations_UnwrapKey = "unwrap_key"
	KeystoresPropertiesCreateAzure_AzureKeyOperations_Verify = "verify"
	KeystoresPropertiesCreateAzure_AzureKeyOperations_WrapKey = "wrap_key"
)

func (*KeystoresPropertiesCreateAzure) isaKeystoresPropertiesCreate() bool {
	return true
}

// UnmarshalKeystoresPropertiesCreateAzure unmarshals an instance of KeystoresPropertiesCreateAzure from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreateAzure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreateAzure)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_protection_level", &obj.AzureKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "azure_key_operations", &obj.AzureKeyOperations)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreateCca : KeystoresPropertiesCreateCca struct
// This model "extends" KeystoresPropertiesCreate
type KeystoresPropertiesCreateCca struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreateCca.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreateCca_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreateCca_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreateCca_Type_Cca = "cca"
	KeystoresPropertiesCreateCca_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreateCca_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoresPropertiesCreateCca.CcaUsageControl property.
const (
	KeystoresPropertiesCreateCca_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeystoresPropertiesCreateCca_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeystoresPropertiesCreateCca_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeystoresPropertiesCreateCca.CcaKeyType property.
const (
	KeystoresPropertiesCreateCca_CcaKeyType_Cipher = "cipher"
	KeystoresPropertiesCreateCca_CcaKeyType_Data = "data"
	KeystoresPropertiesCreateCca_CcaKeyType_Exporter = "exporter"
	KeystoresPropertiesCreateCca_CcaKeyType_Importer = "importer"
)

func (*KeystoresPropertiesCreateCca) isaKeystoresPropertiesCreate() bool {
	return true
}

// UnmarshalKeystoresPropertiesCreateCca unmarshals an instance of KeystoresPropertiesCreateCca from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreateCca(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreateCca)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreateGoogleKms : KeystoresPropertiesCreateGoogleKms struct
// This model "extends" KeystoresPropertiesCreate
type KeystoresPropertiesCreateGoogleKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level,omitempty"`

	GoogleKeyPurpose *string `json:"google_key_purpose,omitempty"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreateGoogleKms.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreateGoogleKms_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreateGoogleKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreateGoogleKms_Type_Cca = "cca"
	KeystoresPropertiesCreateGoogleKms_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreateGoogleKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoresPropertiesCreateGoogleKms.GoogleKeyProtectionLevel property.
const (
	KeystoresPropertiesCreateGoogleKms_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesCreateGoogleKms_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesCreateGoogleKms.GoogleKeyPurpose property.
const (
	KeystoresPropertiesCreateGoogleKms_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeystoresPropertiesCreateGoogleKms_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeystoresPropertiesCreateGoogleKms_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeystoresPropertiesCreateGoogleKms_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeystoresPropertiesCreateGoogleKms.GoogleKmsAlgorithm property.
const (
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeystoresPropertiesCreateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

func (*KeystoresPropertiesCreateGoogleKms) isaKeystoresPropertiesCreate() bool {
	return true
}

// UnmarshalKeystoresPropertiesCreateGoogleKms unmarshals an instance of KeystoresPropertiesCreateGoogleKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreateGoogleKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreateGoogleKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesCreateIbmCloudKms : KeystoresPropertiesCreateIbmCloudKms struct
// This model "extends" KeystoresPropertiesCreate
type KeystoresPropertiesCreateIbmCloudKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	// Managed key naming scheme which will be applied to every key created with this template. Every tag in the naming
	// scheme must be enclosed in angle brackets. For Every tag in the naming scheme, a value will need to be either
	// provided by the user during key creation or computed by the service for the set of special tags.
	NamingScheme *string `json:"naming_scheme,omitempty"`

	// Type of keystore.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the KeystoresPropertiesCreateIbmCloudKms.Type property.
// Type of keystore.
const (
	KeystoresPropertiesCreateIbmCloudKms_Type_AwsKms = "aws_kms"
	KeystoresPropertiesCreateIbmCloudKms_Type_AzureKeyVault = "azure_key_vault"
	KeystoresPropertiesCreateIbmCloudKms_Type_Cca = "cca"
	KeystoresPropertiesCreateIbmCloudKms_Type_GoogleKms = "google_kms"
	KeystoresPropertiesCreateIbmCloudKms_Type_IbmCloudKms = "ibm_cloud_kms"
)

func (*KeystoresPropertiesCreateIbmCloudKms) isaKeystoresPropertiesCreate() bool {
	return true
}

// UnmarshalKeystoresPropertiesCreateIbmCloudKms unmarshals an instance of KeystoresPropertiesCreateIbmCloudKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesCreateIbmCloudKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesCreateIbmCloudKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "naming_scheme", &obj.NamingScheme)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdateAwsKms : KeystoresPropertiesUpdateAwsKms struct
// This model "extends" KeystoresPropertiesUpdate
type KeystoresPropertiesUpdateAwsKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`
}

func (*KeystoresPropertiesUpdateAwsKms) isaKeystoresPropertiesUpdate() bool {
	return true
}

// UnmarshalKeystoresPropertiesUpdateAwsKms unmarshals an instance of KeystoresPropertiesUpdateAwsKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdateAwsKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdateAwsKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdateAzure : KeystoresPropertiesUpdateAzure struct
// This model "extends" KeystoresPropertiesUpdate
type KeystoresPropertiesUpdateAzure struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`
}

func (*KeystoresPropertiesUpdateAzure) isaKeystoresPropertiesUpdate() bool {
	return true
}

// UnmarshalKeystoresPropertiesUpdateAzure unmarshals an instance of KeystoresPropertiesUpdateAzure from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdateAzure(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdateAzure)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdateCca : KeystoresPropertiesUpdateCca struct
// This model "extends" KeystoresPropertiesUpdate
type KeystoresPropertiesUpdateCca struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	CcaUsageControl *string `json:"cca_usage_control,omitempty"`

	CcaKeyType *string `json:"cca_key_type,omitempty"`

	// A list of CCA key words.
	CcaKeyWords []string `json:"cca_key_words,omitempty"`
}

// Constants associated with the KeystoresPropertiesUpdateCca.CcaUsageControl property.
const (
	KeystoresPropertiesUpdateCca_CcaUsageControl_KeyManagementOnly = "key_management_only"
	KeystoresPropertiesUpdateCca_CcaUsageControl_SignatureAndKeyManagement = "signature_and_key_management"
	KeystoresPropertiesUpdateCca_CcaUsageControl_SignatureOnly = "signature_only"
)

// Constants associated with the KeystoresPropertiesUpdateCca.CcaKeyType property.
const (
	KeystoresPropertiesUpdateCca_CcaKeyType_Cipher = "cipher"
	KeystoresPropertiesUpdateCca_CcaKeyType_Data = "data"
	KeystoresPropertiesUpdateCca_CcaKeyType_Exporter = "exporter"
	KeystoresPropertiesUpdateCca_CcaKeyType_Importer = "importer"
)

func (*KeystoresPropertiesUpdateCca) isaKeystoresPropertiesUpdate() bool {
	return true
}

// UnmarshalKeystoresPropertiesUpdateCca unmarshals an instance of KeystoresPropertiesUpdateCca from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdateCca(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdateCca)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_usage_control", &obj.CcaUsageControl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_type", &obj.CcaKeyType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cca_key_words", &obj.CcaKeyWords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdateGoogleKms : KeystoresPropertiesUpdateGoogleKms struct
// This model "extends" KeystoresPropertiesUpdate
type KeystoresPropertiesUpdateGoogleKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`

	GoogleKeyProtectionLevel *string `json:"google_key_protection_level,omitempty"`

	GoogleKeyPurpose *string `json:"google_key_purpose,omitempty"`

	GoogleKmsAlgorithm *string `json:"google_kms_algorithm,omitempty"`
}

// Constants associated with the KeystoresPropertiesUpdateGoogleKms.GoogleKeyProtectionLevel property.
const (
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyProtectionLevel_Hsm = "hsm"
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyProtectionLevel_Software = "software"
)

// Constants associated with the KeystoresPropertiesUpdateGoogleKms.GoogleKeyPurpose property.
const (
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyPurpose_AsymmetricDecrypt = "asymmetric_decrypt"
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyPurpose_AsymmetricSign = "asymmetric_sign"
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyPurpose_EncryptDecrypt = "encrypt_decrypt"
	KeystoresPropertiesUpdateGoogleKms_GoogleKeyPurpose_Mac = "mac"
)

// Constants associated with the KeystoresPropertiesUpdateGoogleKms.GoogleKmsAlgorithm property.
const (
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_EcSignP256Sha256 = "ec_sign_p256_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_EcSignP384Sha384 = "ec_sign_p384_sha384"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_EcSignSecp256k1Sha256 = "ec_sign_secp256k1_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_GoogleSymmetricEncryption = "google_symmetric_encryption"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_HmacSha256 = "hmac_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha1 = "rsa_decrypt_oaep_2048_sha1"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep2048Sha256 = "rsa_decrypt_oaep_2048_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha1 = "rsa_decrypt_oaep_3072_sha1"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep3072Sha256 = "rsa_decrypt_oaep_3072_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha1 = "rsa_decrypt_oaep_4096_sha1"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha256 = "rsa_decrypt_oaep_4096_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaDecryptOaep4096Sha512 = "rsa_decrypt_oaep_4096_sha512"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs12048Sha256 = "rsa_sign_pkcs1_2048_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs13072Sha256 = "rsa_sign_pkcs1_3072_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha256 = "rsa_sign_pkcs1_4096_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPkcs14096Sha512 = "rsa_sign_pkcs1_4096_sha512"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPss2048Sha256 = "rsa_sign_pss_2048_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPss3072Sha256 = "rsa_sign_pss_3072_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha256 = "rsa_sign_pss_4096_sha256"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignPss4096Sha512 = "rsa_sign_pss_4096_sha512"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs12048 = "rsa_sign_raw_pkcs1_2048"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs13072 = "rsa_sign_raw_pkcs1_3072"
	KeystoresPropertiesUpdateGoogleKms_GoogleKmsAlgorithm_RsaSignRawPkcs14096 = "rsa_sign_raw_pkcs1_4096"
)

func (*KeystoresPropertiesUpdateGoogleKms) isaKeystoresPropertiesUpdate() bool {
	return true
}

// UnmarshalKeystoresPropertiesUpdateGoogleKms unmarshals an instance of KeystoresPropertiesUpdateGoogleKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdateGoogleKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdateGoogleKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_protection_level", &obj.GoogleKeyProtectionLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_key_purpose", &obj.GoogleKeyPurpose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "google_kms_algorithm", &obj.GoogleKmsAlgorithm)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoresPropertiesUpdateIbmCloudKms : KeystoresPropertiesUpdateIbmCloudKms struct
// This model "extends" KeystoresPropertiesUpdate
type KeystoresPropertiesUpdateIbmCloudKms struct {
	// Which keystore group to distribute the key to.
	Group *string `json:"group,omitempty"`
}

func (*KeystoresPropertiesUpdateIbmCloudKms) isaKeystoresPropertiesUpdate() bool {
	return true
}

// UnmarshalKeystoresPropertiesUpdateIbmCloudKms unmarshals an instance of KeystoresPropertiesUpdateIbmCloudKms from the specified map of raw messages.
func UnmarshalKeystoresPropertiesUpdateIbmCloudKms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoresPropertiesUpdateIbmCloudKms)
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate : You can connect your Hyper Protect Crypto Services instance to the keystores of another Hyper Protect Crypto Services
// or Key Protect instance, and manage its KMS keys using the current service instance.
// This model "extends" KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate
type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`

	// API endpoint of the IBM Cloud keystore.
	IbmApiEndpoint *string `json:"ibm_api_endpoint" validate:"required"`

	// Endpoint of the IAM service for this IBM Cloud keystore.
	IbmIamEndpoint *string `json:"ibm_iam_endpoint" validate:"required"`

	// The IBM Cloud API key to be used for connecting to this IBM Cloud keystore.
	IbmApiKey *string `json:"ibm_api_key" validate:"required"`

	// The instance ID of the IBM Cloud keystore.
	IbmInstanceID *string `json:"ibm_instance_id" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// The key ring of an IBM Cloud KMS Keystore.
	IbmKeyRing *string `json:"ibm_key_ring,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_IbmVariant_Internal = "internal"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate_IbmVariant_KeyProtect = "key_protect"
)

// NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate : Instantiate KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate(typeVar string, vault *VaultReferenceInCreationRequest, name string, ibmApiEndpoint string, ibmIamEndpoint string, ibmApiKey string, ibmInstanceID string, ibmVariant string) (_model *KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		Name: core.StringPtr(name),
		IbmApiEndpoint: core.StringPtr(ibmApiEndpoint),
		IbmIamEndpoint: core.StringPtr(ibmIamEndpoint),
		IbmApiKey: core.StringPtr(ibmApiKey),
		IbmInstanceID: core.StringPtr(ibmInstanceID),
		IbmVariant: core.StringPtr(ibmVariant),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_endpoint", &obj.IbmApiEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_iam_endpoint", &obj.IbmIamEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_api_key", &obj.IbmApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_instance_id", &obj.IbmInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_key_ring", &obj.IbmKeyRing)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate : An internal keystore generation request describing all information necessary to generate an internal keystore. It
// only require name.
// Models which "extend" this model:
// - KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate
// This model "extends" KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate
type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name" validate:"required"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_IbmVariant_Internal = "internal"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate_IbmVariant_KeyProtect = "key_protect"
)
func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate() bool {
	return true
}

type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateIntf interface {
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateIntf
	isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate() bool
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate unmarshals an instance of KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate : Base of a keystore update.
// This model "extends" KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate
type KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate struct {
	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`
}

func (*KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdate() bool {
	return true
}

func (*KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreUpdateRequest() bool {
	return true
}

// UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate unmarshals an instance of KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate from the specified map of raw messages.
func UnmarshalKeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreUpdateRequestKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate : KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate struct
// Models which "extend" this model:
// - KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate
// This model "extends" KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate
type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_IbmVariant_Internal = "internal"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate_IbmVariant_KeyProtect = "key_protect"
)
func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate() bool {
	return true
}

type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateIntf interface {
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateIntf
	isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate() bool
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate unmarshals an instance of KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate : Base of a keystore update.
// This model "extends" KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate
type KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate struct {
	// Type of keystore.
	Type *string `json:"type" validate:"required"`

	Vault *VaultReferenceInCreationRequest `json:"vault" validate:"required"`

	// Possible IBM Cloud KMS variants.
	IbmVariant *string `json:"ibm_variant" validate:"required"`

	// Name of a target keystore.
	Name *string `json:"name,omitempty"`

	// Description of the keystore.
	Description *string `json:"description,omitempty"`

	// URL of a TLS proxy to use for connecting to private endpoints.
	TlsProxy *string `json:"tls_proxy,omitempty"`

	// A list of groups that this keystore belongs to.
	Groups []string `json:"groups,omitempty"`
}

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate.Type property.
// Type of keystore.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_Type_AwsKms = "aws_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_Type_AzureKeyVault = "azure_key_vault"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_Type_Cca = "cca"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_Type_GoogleKms = "google_kms"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_Type_IbmCloudKms = "ibm_cloud_kms"
)

// Constants associated with the KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate.IbmVariant property.
// Possible IBM Cloud KMS variants.
const (
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_IbmVariant_Hpcs = "hpcs"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_IbmVariant_Internal = "internal"
	KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate_IbmVariant_KeyProtect = "key_protect"
)

// NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate : Instantiate KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate (Generic Model Constructor)
func (*UkoV4) NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate(typeVar string, vault *VaultReferenceInCreationRequest, ibmVariant string) (_model *KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate, err error) {
	_model = &KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate{
		Type: core.StringPtr(typeVar),
		Vault: vault,
		IbmVariant: core.StringPtr(ibmVariant),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreate() bool {
	return true
}

func (*KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate) isaKeystoreCreationRequest() bool {
	return true
}

// UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate unmarshals an instance of KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate from the specified map of raw messages.
func UnmarshalKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "vault", &obj.Vault, UnmarshalVaultReferenceInCreationRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_variant", &obj.IbmVariant)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tls_proxy", &obj.TlsProxy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "groups", &obj.Groups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// ManagedKeysPager can be used to simplify the use of the "ListManagedKeys" method.
//
type ManagedKeysPager struct {
	hasNext bool
	options *ListManagedKeysOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewManagedKeysPager returns a new ManagedKeysPager instance.
func (uko *UkoV4) NewManagedKeysPager(options *ListManagedKeysOptions) (pager *ManagedKeysPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListManagedKeysOptions = *options
	pager = &ManagedKeysPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ManagedKeysPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ManagedKeysPager) GetNextWithContext(ctx context.Context) (page []ManagedKey, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListManagedKeysWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.ManagedKeys

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ManagedKeysPager) GetAllWithContext(ctx context.Context) (allItems []ManagedKey, err error) {
	for pager.HasNext() {
		var nextPage []ManagedKey
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeysPager) GetNext() (page []ManagedKey, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeysPager) GetAll() (allItems []ManagedKey, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// AssociatedResourcesForManagedKeyPager can be used to simplify the use of the "ListAssociatedResourcesForManagedKey" method.
//
type AssociatedResourcesForManagedKeyPager struct {
	hasNext bool
	options *ListAssociatedResourcesForManagedKeyOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewAssociatedResourcesForManagedKeyPager returns a new AssociatedResourcesForManagedKeyPager instance.
func (uko *UkoV4) NewAssociatedResourcesForManagedKeyPager(options *ListAssociatedResourcesForManagedKeyOptions) (pager *AssociatedResourcesForManagedKeyPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListAssociatedResourcesForManagedKeyOptions = *options
	pager = &AssociatedResourcesForManagedKeyPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AssociatedResourcesForManagedKeyPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AssociatedResourcesForManagedKeyPager) GetNextWithContext(ctx context.Context) (page []AssociatedResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListAssociatedResourcesForManagedKeyWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.AssociatedResources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AssociatedResourcesForManagedKeyPager) GetAllWithContext(ctx context.Context) (allItems []AssociatedResource, err error) {
	for pager.HasNext() {
		var nextPage []AssociatedResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AssociatedResourcesForManagedKeyPager) GetNext() (page []AssociatedResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AssociatedResourcesForManagedKeyPager) GetAll() (allItems []AssociatedResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// ManagedKeyVersionsPager can be used to simplify the use of the "ListManagedKeyVersions" method.
//
type ManagedKeyVersionsPager struct {
	hasNext bool
	options *ListManagedKeyVersionsOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewManagedKeyVersionsPager returns a new ManagedKeyVersionsPager instance.
func (uko *UkoV4) NewManagedKeyVersionsPager(options *ListManagedKeyVersionsOptions) (pager *ManagedKeyVersionsPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListManagedKeyVersionsOptions = *options
	pager = &ManagedKeyVersionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ManagedKeyVersionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ManagedKeyVersionsPager) GetNextWithContext(ctx context.Context) (page []ManagedKey, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListManagedKeyVersionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.ManagedKeys

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ManagedKeyVersionsPager) GetAllWithContext(ctx context.Context) (allItems []ManagedKey, err error) {
	for pager.HasNext() {
		var nextPage []ManagedKey
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeyVersionsPager) GetNext() (page []ManagedKey, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeyVersionsPager) GetAll() (allItems []ManagedKey, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// KeyTemplatesPager can be used to simplify the use of the "ListKeyTemplates" method.
//
type KeyTemplatesPager struct {
	hasNext bool
	options *ListKeyTemplatesOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewKeyTemplatesPager returns a new KeyTemplatesPager instance.
func (uko *UkoV4) NewKeyTemplatesPager(options *ListKeyTemplatesOptions) (pager *KeyTemplatesPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListKeyTemplatesOptions = *options
	pager = &KeyTemplatesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *KeyTemplatesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *KeyTemplatesPager) GetNextWithContext(ctx context.Context) (page []Template, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListKeyTemplatesWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Templates

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *KeyTemplatesPager) GetAllWithContext(ctx context.Context) (allItems []Template, err error) {
	for pager.HasNext() {
		var nextPage []Template
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *KeyTemplatesPager) GetNext() (page []Template, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *KeyTemplatesPager) GetAll() (allItems []Template, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// KeystoresPager can be used to simplify the use of the "ListKeystores" method.
//
type KeystoresPager struct {
	hasNext bool
	options *ListKeystoresOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewKeystoresPager returns a new KeystoresPager instance.
func (uko *UkoV4) NewKeystoresPager(options *ListKeystoresOptions) (pager *KeystoresPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListKeystoresOptions = *options
	pager = &KeystoresPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *KeystoresPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *KeystoresPager) GetNextWithContext(ctx context.Context) (page []KeystoreIntf, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListKeystoresWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Keystores

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *KeystoresPager) GetAllWithContext(ctx context.Context) (allItems []KeystoreIntf, err error) {
	for pager.HasNext() {
		var nextPage []KeystoreIntf
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *KeystoresPager) GetNext() (page []KeystoreIntf, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *KeystoresPager) GetAll() (allItems []KeystoreIntf, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// AssociatedResourcesForTargetKeystorePager can be used to simplify the use of the "ListAssociatedResourcesForTargetKeystore" method.
//
type AssociatedResourcesForTargetKeystorePager struct {
	hasNext bool
	options *ListAssociatedResourcesForTargetKeystoreOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewAssociatedResourcesForTargetKeystorePager returns a new AssociatedResourcesForTargetKeystorePager instance.
func (uko *UkoV4) NewAssociatedResourcesForTargetKeystorePager(options *ListAssociatedResourcesForTargetKeystoreOptions) (pager *AssociatedResourcesForTargetKeystorePager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListAssociatedResourcesForTargetKeystoreOptions = *options
	pager = &AssociatedResourcesForTargetKeystorePager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AssociatedResourcesForTargetKeystorePager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AssociatedResourcesForTargetKeystorePager) GetNextWithContext(ctx context.Context) (page []AssociatedResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListAssociatedResourcesForTargetKeystoreWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.AssociatedResources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AssociatedResourcesForTargetKeystorePager) GetAllWithContext(ctx context.Context) (allItems []AssociatedResource, err error) {
	for pager.HasNext() {
		var nextPage []AssociatedResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AssociatedResourcesForTargetKeystorePager) GetNext() (page []AssociatedResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AssociatedResourcesForTargetKeystorePager) GetAll() (allItems []AssociatedResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// ManagedKeysFromKeystorePager can be used to simplify the use of the "ListManagedKeysFromKeystore" method.
//
type ManagedKeysFromKeystorePager struct {
	hasNext bool
	options *ListManagedKeysFromKeystoreOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewManagedKeysFromKeystorePager returns a new ManagedKeysFromKeystorePager instance.
func (uko *UkoV4) NewManagedKeysFromKeystorePager(options *ListManagedKeysFromKeystoreOptions) (pager *ManagedKeysFromKeystorePager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListManagedKeysFromKeystoreOptions = *options
	pager = &ManagedKeysFromKeystorePager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ManagedKeysFromKeystorePager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ManagedKeysFromKeystorePager) GetNextWithContext(ctx context.Context) (page []ManagedKey, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListManagedKeysFromKeystoreWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.ManagedKeys

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ManagedKeysFromKeystorePager) GetAllWithContext(ctx context.Context) (allItems []ManagedKey, err error) {
	for pager.HasNext() {
		var nextPage []ManagedKey
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeysFromKeystorePager) GetNext() (page []ManagedKey, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ManagedKeysFromKeystorePager) GetAll() (allItems []ManagedKey, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// VaultsPager can be used to simplify the use of the "ListVaults" method.
//
type VaultsPager struct {
	hasNext bool
	options *ListVaultsOptions
	client  *UkoV4
	pageContext struct {
		next *int64
	}
}

// NewVaultsPager returns a new VaultsPager instance.
func (uko *UkoV4) NewVaultsPager(options *ListVaultsOptions) (pager *VaultsPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListVaultsOptions = *options
	pager = &VaultsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  uko,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *VaultsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *VaultsPager) GetNextWithContext(ctx context.Context) (page []Vault, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListVaultsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Vaults

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *VaultsPager) GetAllWithContext(ctx context.Context) (allItems []Vault, err error) {
	for pager.HasNext() {
		var nextPage []Vault
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *VaultsPager) GetNext() (page []Vault, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *VaultsPager) GetAll() (allItems []Vault, err error) {
	return pager.GetAllWithContext(context.Background())
}
