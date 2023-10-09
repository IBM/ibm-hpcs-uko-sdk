/**
 * (C) Copyright IBM Corp. 2023.
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

package ukov4_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-hpcs-uko-sdk/ukov4"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`UkoV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ukoService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ukoService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
				URL: "https://ukov4/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ukoService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"UKO_URL": "https://ukov4/api",
				"UKO_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ukoService, serviceErr := ukov4.NewUkoV4UsingExternalConfig(&ukov4.UkoV4Options{
				})
				Expect(ukoService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ukoService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ukoService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ukoService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ukoService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ukoService, serviceErr := ukov4.NewUkoV4UsingExternalConfig(&ukov4.UkoV4Options{
					URL: "https://testService/api",
				})
				Expect(ukoService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ukoService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ukoService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ukoService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ukoService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ukoService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ukoService, serviceErr := ukov4.NewUkoV4UsingExternalConfig(&ukov4.UkoV4Options{
				})
				err := ukoService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ukoService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ukoService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ukoService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ukoService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ukoService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"UKO_URL": "https://ukov4/api",
				"UKO_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ukoService, serviceErr := ukov4.NewUkoV4UsingExternalConfig(&ukov4.UkoV4Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ukoService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"UKO_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ukoService, serviceErr := ukov4.NewUkoV4UsingExternalConfig(&ukov4.UkoV4Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ukoService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ukov4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListManagedKeys(listManagedKeysOptions *ListManagedKeysOptions) - Operation response error`, func() {
		listManagedKeysPath := "/v4/managed_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListManagedKeys with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := new(ukov4.ListManagedKeysOptions)
				listManagedKeysOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listManagedKeysOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListManagedKeys(listManagedKeysOptions *ListManagedKeysOptions)`, func() {
		listManagedKeysPath := "/v4/managed_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeys successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := new(ukov4.ListManagedKeysOptions)
				listManagedKeysOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listManagedKeysOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListManagedKeysWithContext(ctx, listManagedKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListManagedKeysWithContext(ctx, listManagedKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeys successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListManagedKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := new(ukov4.ListManagedKeysOptions)
				listManagedKeysOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listManagedKeysOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListManagedKeys with error: Operation request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := new(ukov4.ListManagedKeysOptions)
				listManagedKeysOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listManagedKeysOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListManagedKeys successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := new(ukov4.ListManagedKeysOptions)
				listManagedKeysOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listManagedKeysOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListManagedKeys(listManagedKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.ManagedKeyList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ManagedKeysPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeysOptionsModel := &ukov4.ListManagedKeysOptions{
					Accept: core.StringPtr("application/json"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Algorithm: []string{"aes"},
					State: []string{"pre_activation", "active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeysPager(listManagedKeysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.ManagedKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ManagedKeysPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeysOptionsModel := &ukov4.ListManagedKeysOptions{
					Accept: core.StringPtr("application/json"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Algorithm: []string{"aes"},
					State: []string{"pre_activation", "active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeysPager(listManagedKeysOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateManagedKey(createManagedKeyOptions *CreateManagedKeyOptions) - Operation response error`, func() {
		createManagedKeyPath := "/v4/managed_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Description = core.StringPtr("testString")
				createManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateManagedKey(createManagedKeyOptions *CreateManagedKeyOptions)`, func() {
		createManagedKeyPath := "/v4/managed_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke CreateManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Description = core.StringPtr("testString")
				createManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.CreateManagedKeyWithContext(ctx, createManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.CreateManagedKeyWithContext(ctx, createManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke CreateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.CreateManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Description = core.StringPtr("testString")
				createManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Description = core.StringPtr("testString")
				createManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateManagedKeyOptions model with no property values
				createManagedKeyOptionsModelNew := new(ukov4.CreateManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.CreateManagedKey(createManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Description = core.StringPtr("testString")
				createManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.CreateManagedKey(createManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteManagedKey(deleteManagedKeyOptions *DeleteManagedKeyOptions)`, func() {
		deleteManagedKeyPath := "/v4/managed_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteManagedKeyPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ukoService.DeleteManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteManagedKeyOptions model
				deleteManagedKeyOptionsModel := new(ukov4.DeleteManagedKeyOptions)
				deleteManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deleteManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deleteManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ukoService.DeleteManagedKey(deleteManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeleteManagedKeyOptions model
				deleteManagedKeyOptionsModel := new(ukov4.DeleteManagedKeyOptions)
				deleteManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deleteManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deleteManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ukoService.DeleteManagedKey(deleteManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteManagedKeyOptions model with no property values
				deleteManagedKeyOptionsModelNew := new(ukov4.DeleteManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ukoService.DeleteManagedKey(deleteManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetManagedKey(getManagedKeyOptions *GetManagedKeyOptions) - Operation response error`, func() {
		getManagedKeyPath := "/v4/managed_keys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetManagedKeyOptions model
				getManagedKeyOptionsModel := new(ukov4.GetManagedKeyOptions)
				getManagedKeyOptionsModel.ID = core.StringPtr("testString")
				getManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetManagedKey(getManagedKeyOptions *GetManagedKeyOptions)`, func() {
		getManagedKeyPath := "/v4/managed_keys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke GetManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetManagedKeyOptions model
				getManagedKeyOptionsModel := new(ukov4.GetManagedKeyOptions)
				getManagedKeyOptionsModel.ID = core.StringPtr("testString")
				getManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetManagedKeyWithContext(ctx, getManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetManagedKeyWithContext(ctx, getManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke GetManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetManagedKeyOptions model
				getManagedKeyOptionsModel := new(ukov4.GetManagedKeyOptions)
				getManagedKeyOptionsModel.ID = core.StringPtr("testString")
				getManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetManagedKeyOptions model
				getManagedKeyOptionsModel := new(ukov4.GetManagedKeyOptions)
				getManagedKeyOptionsModel.ID = core.StringPtr("testString")
				getManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetManagedKeyOptions model with no property values
				getManagedKeyOptionsModelNew := new(ukov4.GetManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetManagedKey(getManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetManagedKeyOptions model
				getManagedKeyOptionsModel := new(ukov4.GetManagedKeyOptions)
				getManagedKeyOptionsModel.ID = core.StringPtr("testString")
				getManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetManagedKey(getManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateManagedKey(updateManagedKeyOptions *UpdateManagedKeyOptions) - Operation response error`, func() {
		updateManagedKeyPath := "/v4/managed_keys/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Description = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateManagedKey(updateManagedKeyOptions *UpdateManagedKeyOptions)`, func() {
		updateManagedKeyPath := "/v4/managed_keys/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke UpdateManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Description = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UpdateManagedKeyWithContext(ctx, updateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UpdateManagedKeyWithContext(ctx, updateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke UpdateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UpdateManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Description = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Description = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateManagedKeyOptions model with no property values
				updateManagedKeyOptionsModelNew := new(ukov4.UpdateManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UpdateManagedKey(updateManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Description = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UpdateManagedKey(updateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptions *ListAssociatedResourcesForManagedKeyOptions) - Operation response error`, func() {
		listAssociatedResourcesForManagedKeyPath := "/v4/managed_keys/testString/associated_resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				listAssociatedResourcesForManagedKeyOptionsModel := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				listAssociatedResourcesForManagedKeyOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptions *ListAssociatedResourcesForManagedKeyOptions)`, func() {
		listAssociatedResourcesForManagedKeyPath := "/v4/managed_keys/testString/associated_resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "associated_resources": [{"id": "crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_key": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label": "IBM CLOUD KEY", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "referenced_keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000", "name": "keyprotecttest", "type": "com_ibm_cloud_kms_registration", "com_ibm_cloud_kms_registration": {"prevents_key_deletion": true, "service_name": "cloud-object-storage", "service_instance_name": "Cloud Object Storage-7s", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest", "description": "Example Description"}}]}`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				listAssociatedResourcesForManagedKeyOptionsModel := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				listAssociatedResourcesForManagedKeyOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListAssociatedResourcesForManagedKeyWithContext(ctx, listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListAssociatedResourcesForManagedKeyWithContext(ctx, listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "associated_resources": [{"id": "crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_key": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label": "IBM CLOUD KEY", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "referenced_keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000", "name": "keyprotecttest", "type": "com_ibm_cloud_kms_registration", "com_ibm_cloud_kms_registration": {"prevents_key_deletion": true, "service_name": "cloud-object-storage", "service_instance_name": "Cloud Object Storage-7s", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest", "description": "Example Description"}}]}`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListAssociatedResourcesForManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				listAssociatedResourcesForManagedKeyOptionsModel := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				listAssociatedResourcesForManagedKeyOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAssociatedResourcesForManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				listAssociatedResourcesForManagedKeyOptionsModel := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				listAssociatedResourcesForManagedKeyOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAssociatedResourcesForManagedKeyOptions model with no property values
				listAssociatedResourcesForManagedKeyOptionsModelNew := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAssociatedResourcesForManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				listAssociatedResourcesForManagedKeyOptionsModel := new(ukov4.ListAssociatedResourcesForManagedKeyOptions)
				listAssociatedResourcesForManagedKeyOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListAssociatedResourcesForManagedKey(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForManagedKeyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"associated_resources":[{"id":"crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"managed_key":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label":"IBM CLOUD KEY","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"referenced_keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000","name":"keyprotecttest","type":"com_ibm_cloud_kms_registration","com_ibm_cloud_kms_registration":{"prevents_key_deletion":true,"service_name":"cloud-object-storage","service_instance_name":"Cloud Object Storage-7s","crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest","description":"Example Description"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"associated_resources":[{"id":"crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"managed_key":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label":"IBM CLOUD KEY","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"referenced_keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000","name":"keyprotecttest","type":"com_ibm_cloud_kms_registration","com_ibm_cloud_kms_registration":{"prevents_key_deletion":true,"service_name":"cloud-object-storage","service_instance_name":"Cloud Object Storage-7s","crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest","description":"Example Description"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AssociatedResourcesForManagedKeyPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listAssociatedResourcesForManagedKeyOptionsModel := &ukov4.ListAssociatedResourcesForManagedKeyOptions{
					ID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"name"},
				}

				pager, err := ukoService.NewAssociatedResourcesForManagedKeyPager(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.AssociatedResource
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AssociatedResourcesForManagedKeyPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listAssociatedResourcesForManagedKeyOptionsModel := &ukov4.ListAssociatedResourcesForManagedKeyOptions{
					ID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"name"},
				}

				pager, err := ukoService.NewAssociatedResourcesForManagedKeyPager(listAssociatedResourcesForManagedKeyOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListManagedKeyVersions(listManagedKeyVersionsOptions *ListManagedKeyVersionsOptions) - Operation response error`, func() {
		listManagedKeyVersionsPath := "/v4/managed_keys/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeyVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListManagedKeyVersions with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeyVersionsOptions model
				listManagedKeyVersionsOptionsModel := new(ukov4.ListManagedKeyVersionsOptions)
				listManagedKeyVersionsOptionsModel.ID = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Algorithm = []string{"aes"}
				listManagedKeyVersionsOptionsModel.State = []string{"active"}
				listManagedKeyVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeyVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeyVersionsOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeyVersionsOptionsModel.Label = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Size = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeyVersionsOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeyVersionsOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeyVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListManagedKeyVersions(listManagedKeyVersionsOptions *ListManagedKeyVersionsOptions)`, func() {
		listManagedKeyVersionsPath := "/v4/managed_keys/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeyVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeyVersions successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListManagedKeyVersionsOptions model
				listManagedKeyVersionsOptionsModel := new(ukov4.ListManagedKeyVersionsOptions)
				listManagedKeyVersionsOptionsModel.ID = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Algorithm = []string{"aes"}
				listManagedKeyVersionsOptionsModel.State = []string{"active"}
				listManagedKeyVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeyVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeyVersionsOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeyVersionsOptionsModel.Label = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Size = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeyVersionsOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeyVersionsOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeyVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListManagedKeyVersionsWithContext(ctx, listManagedKeyVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListManagedKeyVersionsWithContext(ctx, listManagedKeyVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeyVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeyVersions successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListManagedKeyVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListManagedKeyVersionsOptions model
				listManagedKeyVersionsOptionsModel := new(ukov4.ListManagedKeyVersionsOptions)
				listManagedKeyVersionsOptionsModel.ID = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Algorithm = []string{"aes"}
				listManagedKeyVersionsOptionsModel.State = []string{"active"}
				listManagedKeyVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeyVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeyVersionsOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeyVersionsOptionsModel.Label = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Size = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeyVersionsOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeyVersionsOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeyVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListManagedKeyVersions with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeyVersionsOptions model
				listManagedKeyVersionsOptionsModel := new(ukov4.ListManagedKeyVersionsOptions)
				listManagedKeyVersionsOptionsModel.ID = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Algorithm = []string{"aes"}
				listManagedKeyVersionsOptionsModel.State = []string{"active"}
				listManagedKeyVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeyVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeyVersionsOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeyVersionsOptionsModel.Label = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Size = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeyVersionsOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeyVersionsOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeyVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListManagedKeyVersionsOptions model with no property values
				listManagedKeyVersionsOptionsModelNew := new(ukov4.ListManagedKeyVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListManagedKeyVersions successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeyVersionsOptions model
				listManagedKeyVersionsOptionsModel := new(ukov4.ListManagedKeyVersionsOptions)
				listManagedKeyVersionsOptionsModel.ID = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Algorithm = []string{"aes"}
				listManagedKeyVersionsOptionsModel.State = []string{"active"}
				listManagedKeyVersionsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeyVersionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeyVersionsOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeyVersionsOptionsModel.Label = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.Size = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeyVersionsOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
				listManagedKeyVersionsOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeyVersionsOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeyVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListManagedKeyVersions(listManagedKeyVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.ManagedKeyList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeyVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ManagedKeyVersionsPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeyVersionsOptionsModel := &ukov4.ListManagedKeyVersionsOptions{
					ID: core.StringPtr("testString"),
					Algorithm: []string{"aes"},
					State: []string{"active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeyVersionsPager(listManagedKeyVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.ManagedKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ManagedKeyVersionsPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeyVersionsOptionsModel := &ukov4.ListManagedKeyVersionsOptions{
					ID: core.StringPtr("testString"),
					Algorithm: []string{"aes"},
					State: []string{"active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeyVersionsPager(listManagedKeyVersionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptions *GetKeyDistributionStatusForKeystoresOptions) - Operation response error`, func() {
		getKeyDistributionStatusForKeystoresPath := "/v4/managed_keys/testString/status_in_keystores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyDistributionStatusForKeystoresPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetKeyDistributionStatusForKeystores with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				getKeyDistributionStatusForKeystoresOptionsModel := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				getKeyDistributionStatusForKeystoresOptionsModel.ID = core.StringPtr("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptions *GetKeyDistributionStatusForKeystoresOptions)`, func() {
		getKeyDistributionStatusForKeystoresPath := "/v4/managed_keys/testString/status_in_keystores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyDistributionStatusForKeystoresPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
				}))
			})
			It(`Invoke GetKeyDistributionStatusForKeystores successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				getKeyDistributionStatusForKeystoresOptionsModel := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				getKeyDistributionStatusForKeystoresOptionsModel.ID = core.StringPtr("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetKeyDistributionStatusForKeystoresWithContext(ctx, getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetKeyDistributionStatusForKeystoresWithContext(ctx, getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyDistributionStatusForKeystoresPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
				}))
			})
			It(`Invoke GetKeyDistributionStatusForKeystores successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetKeyDistributionStatusForKeystores(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				getKeyDistributionStatusForKeystoresOptionsModel := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				getKeyDistributionStatusForKeystoresOptionsModel.ID = core.StringPtr("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetKeyDistributionStatusForKeystores with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				getKeyDistributionStatusForKeystoresOptionsModel := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				getKeyDistributionStatusForKeystoresOptionsModel.ID = core.StringPtr("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKeyDistributionStatusForKeystoresOptions model with no property values
				getKeyDistributionStatusForKeystoresOptionsModelNew := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetKeyDistributionStatusForKeystores successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				getKeyDistributionStatusForKeystoresOptionsModel := new(ukov4.GetKeyDistributionStatusForKeystoresOptions)
				getKeyDistributionStatusForKeystoresOptionsModel.ID = core.StringPtr("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptions *UpdateManagedKeyFromTemplateOptions) - Operation response error`, func() {
		updateManagedKeyFromTemplatePath := "/v4/managed_keys/testString/update_from_template"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyFromTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for dry_run query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateManagedKeyFromTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				updateManagedKeyFromTemplateOptionsModel := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				updateManagedKeyFromTemplateOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.DryRun = core.BoolPtr(false)
				updateManagedKeyFromTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptions *UpdateManagedKeyFromTemplateOptions)`, func() {
		updateManagedKeyFromTemplatePath := "/v4/managed_keys/testString/update_from_template"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyFromTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for dry_run query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke UpdateManagedKeyFromTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				updateManagedKeyFromTemplateOptionsModel := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				updateManagedKeyFromTemplateOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.DryRun = core.BoolPtr(false)
				updateManagedKeyFromTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UpdateManagedKeyFromTemplateWithContext(ctx, updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UpdateManagedKeyFromTemplateWithContext(ctx, updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateManagedKeyFromTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for dry_run query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke UpdateManagedKeyFromTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UpdateManagedKeyFromTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				updateManagedKeyFromTemplateOptionsModel := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				updateManagedKeyFromTemplateOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.DryRun = core.BoolPtr(false)
				updateManagedKeyFromTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateManagedKeyFromTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				updateManagedKeyFromTemplateOptionsModel := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				updateManagedKeyFromTemplateOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.DryRun = core.BoolPtr(false)
				updateManagedKeyFromTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateManagedKeyFromTemplateOptions model with no property values
				updateManagedKeyFromTemplateOptionsModelNew := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateManagedKeyFromTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				updateManagedKeyFromTemplateOptionsModel := new(ukov4.UpdateManagedKeyFromTemplateOptions)
				updateManagedKeyFromTemplateOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.DryRun = core.BoolPtr(false)
				updateManagedKeyFromTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UpdateManagedKeyFromTemplate(updateManagedKeyFromTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ActivateManagedKey(activateManagedKeyOptions *ActivateManagedKeyOptions) - Operation response error`, func() {
		activateManagedKeyPath := "/v4/managed_keys/testString/activate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(activateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ActivateManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ActivateManagedKeyOptions model
				activateManagedKeyOptionsModel := new(ukov4.ActivateManagedKeyOptions)
				activateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				activateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				activateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ActivateManagedKey(activateManagedKeyOptions *ActivateManagedKeyOptions)`, func() {
		activateManagedKeyPath := "/v4/managed_keys/testString/activate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(activateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke ActivateManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ActivateManagedKeyOptions model
				activateManagedKeyOptionsModel := new(ukov4.ActivateManagedKeyOptions)
				activateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				activateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				activateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ActivateManagedKeyWithContext(ctx, activateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ActivateManagedKeyWithContext(ctx, activateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(activateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke ActivateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ActivateManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ActivateManagedKeyOptions model
				activateManagedKeyOptionsModel := new(ukov4.ActivateManagedKeyOptions)
				activateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				activateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				activateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ActivateManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ActivateManagedKeyOptions model
				activateManagedKeyOptionsModel := new(ukov4.ActivateManagedKeyOptions)
				activateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				activateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				activateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ActivateManagedKeyOptions model with no property values
				activateManagedKeyOptionsModelNew := new(ukov4.ActivateManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ActivateManagedKey(activateManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ActivateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ActivateManagedKeyOptions model
				activateManagedKeyOptionsModel := new(ukov4.ActivateManagedKeyOptions)
				activateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				activateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				activateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ActivateManagedKey(activateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeactivateManagedKey(deactivateManagedKeyOptions *DeactivateManagedKeyOptions) - Operation response error`, func() {
		deactivateManagedKeyPath := "/v4/managed_keys/testString/deactivate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deactivateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeactivateManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeactivateManagedKeyOptions model
				deactivateManagedKeyOptionsModel := new(ukov4.DeactivateManagedKeyOptions)
				deactivateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeactivateManagedKey(deactivateManagedKeyOptions *DeactivateManagedKeyOptions)`, func() {
		deactivateManagedKeyPath := "/v4/managed_keys/testString/deactivate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deactivateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke DeactivateManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the DeactivateManagedKeyOptions model
				deactivateManagedKeyOptionsModel := new(ukov4.DeactivateManagedKeyOptions)
				deactivateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.DeactivateManagedKeyWithContext(ctx, deactivateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.DeactivateManagedKeyWithContext(ctx, deactivateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deactivateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke DeactivateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.DeactivateManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeactivateManagedKeyOptions model
				deactivateManagedKeyOptionsModel := new(ukov4.DeactivateManagedKeyOptions)
				deactivateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeactivateManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeactivateManagedKeyOptions model
				deactivateManagedKeyOptionsModel := new(ukov4.DeactivateManagedKeyOptions)
				deactivateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeactivateManagedKeyOptions model with no property values
				deactivateManagedKeyOptionsModelNew := new(ukov4.DeactivateManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeactivateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeactivateManagedKeyOptions model
				deactivateManagedKeyOptionsModel := new(ukov4.DeactivateManagedKeyOptions)
				deactivateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				deactivateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.DeactivateManagedKey(deactivateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DestroyManagedKey(destroyManagedKeyOptions *DestroyManagedKeyOptions) - Operation response error`, func() {
		destroyManagedKeyPath := "/v4/managed_keys/testString/destroy"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(destroyManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DestroyManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DestroyManagedKeyOptions model
				destroyManagedKeyOptionsModel := new(ukov4.DestroyManagedKeyOptions)
				destroyManagedKeyOptionsModel.ID = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DestroyManagedKey(destroyManagedKeyOptions *DestroyManagedKeyOptions)`, func() {
		destroyManagedKeyPath := "/v4/managed_keys/testString/destroy"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(destroyManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke DestroyManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the DestroyManagedKeyOptions model
				destroyManagedKeyOptionsModel := new(ukov4.DestroyManagedKeyOptions)
				destroyManagedKeyOptionsModel.ID = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.DestroyManagedKeyWithContext(ctx, destroyManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.DestroyManagedKeyWithContext(ctx, destroyManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(destroyManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke DestroyManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.DestroyManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DestroyManagedKeyOptions model
				destroyManagedKeyOptionsModel := new(ukov4.DestroyManagedKeyOptions)
				destroyManagedKeyOptionsModel.ID = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DestroyManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DestroyManagedKeyOptions model
				destroyManagedKeyOptionsModel := new(ukov4.DestroyManagedKeyOptions)
				destroyManagedKeyOptionsModel.ID = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DestroyManagedKeyOptions model with no property values
				destroyManagedKeyOptionsModelNew := new(ukov4.DestroyManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.DestroyManagedKey(destroyManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DestroyManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DestroyManagedKeyOptions model
				destroyManagedKeyOptionsModel := new(ukov4.DestroyManagedKeyOptions)
				destroyManagedKeyOptionsModel.ID = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				destroyManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.DestroyManagedKey(destroyManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SyncManagedKey(syncManagedKeyOptions *SyncManagedKeyOptions) - Operation response error`, func() {
		syncManagedKeyPath := "/v4/managed_keys/testString/sync_status_in_keystores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(syncManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SyncManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the SyncManagedKeyOptions model
				syncManagedKeyOptionsModel := new(ukov4.SyncManagedKeyOptions)
				syncManagedKeyOptionsModel.ID = core.StringPtr("testString")
				syncManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				syncManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SyncManagedKey(syncManagedKeyOptions *SyncManagedKeyOptions)`, func() {
		syncManagedKeyPath := "/v4/managed_keys/testString/sync_status_in_keystores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(syncManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
				}))
			})
			It(`Invoke SyncManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the SyncManagedKeyOptions model
				syncManagedKeyOptionsModel := new(ukov4.SyncManagedKeyOptions)
				syncManagedKeyOptionsModel.ID = core.StringPtr("testString")
				syncManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				syncManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.SyncManagedKeyWithContext(ctx, syncManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.SyncManagedKeyWithContext(ctx, syncManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(syncManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
				}))
			})
			It(`Invoke SyncManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.SyncManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SyncManagedKeyOptions model
				syncManagedKeyOptionsModel := new(ukov4.SyncManagedKeyOptions)
				syncManagedKeyOptionsModel.ID = core.StringPtr("testString")
				syncManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				syncManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SyncManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the SyncManagedKeyOptions model
				syncManagedKeyOptionsModel := new(ukov4.SyncManagedKeyOptions)
				syncManagedKeyOptionsModel.ID = core.StringPtr("testString")
				syncManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				syncManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SyncManagedKeyOptions model with no property values
				syncManagedKeyOptionsModelNew := new(ukov4.SyncManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.SyncManagedKey(syncManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke SyncManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the SyncManagedKeyOptions model
				syncManagedKeyOptionsModel := new(ukov4.SyncManagedKeyOptions)
				syncManagedKeyOptionsModel.ID = core.StringPtr("testString")
				syncManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				syncManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.SyncManagedKey(syncManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RotateManagedKey(rotateManagedKeyOptions *RotateManagedKeyOptions) - Operation response error`, func() {
		rotateManagedKeyPath := "/v4/managed_keys/testString/rotate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rotateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RotateManagedKey with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the RotateManagedKeyOptions model
				rotateManagedKeyOptionsModel := new(ukov4.RotateManagedKeyOptions)
				rotateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RotateManagedKey(rotateManagedKeyOptions *RotateManagedKeyOptions)`, func() {
		rotateManagedKeyPath := "/v4/managed_keys/testString/rotate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rotateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke RotateManagedKey successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the RotateManagedKeyOptions model
				rotateManagedKeyOptionsModel := new(ukov4.RotateManagedKeyOptions)
				rotateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.RotateManagedKeyWithContext(ctx, rotateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.RotateManagedKeyWithContext(ctx, rotateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(rotateManagedKeyPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}`)
				}))
			})
			It(`Invoke RotateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.RotateManagedKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RotateManagedKeyOptions model
				rotateManagedKeyOptionsModel := new(ukov4.RotateManagedKeyOptions)
				rotateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RotateManagedKey with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the RotateManagedKeyOptions model
				rotateManagedKeyOptionsModel := new(ukov4.RotateManagedKeyOptions)
				rotateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RotateManagedKeyOptions model with no property values
				rotateManagedKeyOptionsModelNew := new(ukov4.RotateManagedKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.RotateManagedKey(rotateManagedKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RotateManagedKey successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the RotateManagedKeyOptions model
				rotateManagedKeyOptionsModel := new(ukov4.RotateManagedKeyOptions)
				rotateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				rotateManagedKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.RotateManagedKey(rotateManagedKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListKeyTemplates(listKeyTemplatesOptions *ListKeyTemplatesOptions) - Operation response error`, func() {
		listKeyTemplatesPath := "/v4/templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["naming_scheme"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["key.size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListKeyTemplates with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := new(ukov4.ListKeyTemplatesOptions)
				listKeyTemplatesOptionsModel.Accept = core.StringPtr("application/json")
				listKeyTemplatesOptionsModel.Name = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.NamingScheme = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = []string{"aes"}
				listKeyTemplatesOptionsModel.KeySize = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeystoresType = []string{"ibm_cloud_kms"}
				listKeyTemplatesOptionsModel.KeystoresGroup = []string{"testString"}
				listKeyTemplatesOptionsModel.CreatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.Type = []string{"user_defined"}
				listKeyTemplatesOptionsModel.State = []string{"unarchived"}
				listKeyTemplatesOptionsModel.Sort = []string{"-updated_at"}
				listKeyTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeyTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListKeyTemplates(listKeyTemplatesOptions *ListKeyTemplatesOptions)`, func() {
		listKeyTemplatesPath := "/v4/templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["naming_scheme"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["key.size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "templates": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
				}))
			})
			It(`Invoke ListKeyTemplates successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := new(ukov4.ListKeyTemplatesOptions)
				listKeyTemplatesOptionsModel.Accept = core.StringPtr("application/json")
				listKeyTemplatesOptionsModel.Name = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.NamingScheme = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = []string{"aes"}
				listKeyTemplatesOptionsModel.KeySize = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeystoresType = []string{"ibm_cloud_kms"}
				listKeyTemplatesOptionsModel.KeystoresGroup = []string{"testString"}
				listKeyTemplatesOptionsModel.CreatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.Type = []string{"user_defined"}
				listKeyTemplatesOptionsModel.State = []string{"unarchived"}
				listKeyTemplatesOptionsModel.Sort = []string{"-updated_at"}
				listKeyTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeyTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListKeyTemplatesWithContext(ctx, listKeyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListKeyTemplatesWithContext(ctx, listKeyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["naming_scheme"]).To(Equal([]string{"My Example Template"}))
					Expect(req.URL.Query()["key.size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["key.size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "templates": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
				}))
			})
			It(`Invoke ListKeyTemplates successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListKeyTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := new(ukov4.ListKeyTemplatesOptions)
				listKeyTemplatesOptionsModel.Accept = core.StringPtr("application/json")
				listKeyTemplatesOptionsModel.Name = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.NamingScheme = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = []string{"aes"}
				listKeyTemplatesOptionsModel.KeySize = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeystoresType = []string{"ibm_cloud_kms"}
				listKeyTemplatesOptionsModel.KeystoresGroup = []string{"testString"}
				listKeyTemplatesOptionsModel.CreatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.Type = []string{"user_defined"}
				listKeyTemplatesOptionsModel.State = []string{"unarchived"}
				listKeyTemplatesOptionsModel.Sort = []string{"-updated_at"}
				listKeyTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeyTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListKeyTemplates with error: Operation request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := new(ukov4.ListKeyTemplatesOptions)
				listKeyTemplatesOptionsModel.Accept = core.StringPtr("application/json")
				listKeyTemplatesOptionsModel.Name = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.NamingScheme = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = []string{"aes"}
				listKeyTemplatesOptionsModel.KeySize = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeystoresType = []string{"ibm_cloud_kms"}
				listKeyTemplatesOptionsModel.KeystoresGroup = []string{"testString"}
				listKeyTemplatesOptionsModel.CreatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.Type = []string{"user_defined"}
				listKeyTemplatesOptionsModel.State = []string{"unarchived"}
				listKeyTemplatesOptionsModel.Sort = []string{"-updated_at"}
				listKeyTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeyTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListKeyTemplates successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := new(ukov4.ListKeyTemplatesOptions)
				listKeyTemplatesOptionsModel.Accept = core.StringPtr("application/json")
				listKeyTemplatesOptionsModel.Name = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.NamingScheme = core.StringPtr("My Example Template")
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = []string{"aes"}
				listKeyTemplatesOptionsModel.KeySize = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeySizeMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.KeystoresType = []string{"ibm_cloud_kms"}
				listKeyTemplatesOptionsModel.KeystoresGroup = []string{"testString"}
				listKeyTemplatesOptionsModel.CreatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAt = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listKeyTemplatesOptionsModel.Type = []string{"user_defined"}
				listKeyTemplatesOptionsModel.State = []string{"unarchived"}
				listKeyTemplatesOptionsModel.Sort = []string{"-updated_at"}
				listKeyTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeyTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListKeyTemplates(listKeyTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.TemplateList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.TemplateList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.TemplateList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.TemplateList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"templates":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","version":1,"name":"AWS-TEMPLATE","naming_scheme":"A-<APP>-AES256-<ENV>-<GROUP>","type":["user_defined"],"state":"unarchived","keys_count":3456,"key":{"size":"256","algorithm":"aes","activation_date":"P5Y1M1W2D","expiration_date":"P1Y2M1W4D","state":"active","deactivate_on_rotation":true},"description":"The description of the template","created_at":"2022-02-05T23:00:14.000Z","updated_at":"2022-02-05T23:00:14.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","keystores":[{"group":"Production","naming_scheme":"A-<APP>-AES256-<ENV>-<GROUP>","type":"ibm_cloud_kms","google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"templates":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","version":1,"name":"AWS-TEMPLATE","naming_scheme":"A-<APP>-AES256-<ENV>-<GROUP>","type":["user_defined"],"state":"unarchived","keys_count":3456,"key":{"size":"256","algorithm":"aes","activation_date":"P5Y1M1W2D","expiration_date":"P1Y2M1W4D","state":"active","deactivate_on_rotation":true},"description":"The description of the template","created_at":"2022-02-05T23:00:14.000Z","updated_at":"2022-02-05T23:00:14.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","keystores":[{"group":"Production","naming_scheme":"A-<APP>-AES256-<ENV>-<GROUP>","type":"ibm_cloud_kms","google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use KeyTemplatesPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listKeyTemplatesOptionsModel := &ukov4.ListKeyTemplatesOptions{
					Accept: core.StringPtr("application/json"),
					Name: core.StringPtr("My Example Template"),
					NamingScheme: core.StringPtr("My Example Template"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					KeyAlgorithm: []string{"aes"},
					KeySize: core.StringPtr("testString"),
					KeySizeMin: core.StringPtr("testString"),
					KeySizeMax: core.StringPtr("testString"),
					KeystoresType: []string{"ibm_cloud_kms"},
					KeystoresGroup: []string{"testString"},
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Type: []string{"user_defined"},
					State: []string{"unarchived"},
					Sort: []string{"-updated_at"},
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := ukoService.NewKeyTemplatesPager(listKeyTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.Template
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use KeyTemplatesPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listKeyTemplatesOptionsModel := &ukov4.ListKeyTemplatesOptions{
					Accept: core.StringPtr("application/json"),
					Name: core.StringPtr("My Example Template"),
					NamingScheme: core.StringPtr("My Example Template"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					KeyAlgorithm: []string{"aes"},
					KeySize: core.StringPtr("testString"),
					KeySizeMin: core.StringPtr("testString"),
					KeySizeMax: core.StringPtr("testString"),
					KeystoresType: []string{"ibm_cloud_kms"},
					KeystoresGroup: []string{"testString"},
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Type: []string{"user_defined"},
					State: []string{"unarchived"},
					Sort: []string{"-updated_at"},
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := ukoService.NewKeyTemplatesPager(listKeyTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateKeyTemplate(createKeyTemplateOptions *CreateKeyTemplateOptions) - Operation response error`, func() {
		createKeyTemplatePath := "/v4/templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				createKeyTemplateOptionsModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.Type = []string{"user_defined"}
				createKeyTemplateOptionsModel.State = core.StringPtr("unarchived")
				createKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateKeyTemplate(createKeyTemplateOptions *CreateKeyTemplateOptions)`, func() {
		createKeyTemplatePath := "/v4/templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke CreateKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				createKeyTemplateOptionsModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.Type = []string{"user_defined"}
				createKeyTemplateOptionsModel.State = core.StringPtr("unarchived")
				createKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.CreateKeyTemplateWithContext(ctx, createKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.CreateKeyTemplateWithContext(ctx, createKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke CreateKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.CreateKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				createKeyTemplateOptionsModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.Type = []string{"user_defined"}
				createKeyTemplateOptionsModel.State = core.StringPtr("unarchived")
				createKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				createKeyTemplateOptionsModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.Type = []string{"user_defined"}
				createKeyTemplateOptionsModel.State = core.StringPtr("unarchived")
				createKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateKeyTemplateOptions model with no property values
				createKeyTemplateOptionsModelNew := new(ukov4.CreateKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.CreateKeyTemplate(createKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				createKeyTemplateOptionsModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.Type = []string{"user_defined"}
				createKeyTemplateOptionsModel.State = core.StringPtr("unarchived")
				createKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.CreateKeyTemplate(createKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteKeyTemplate(deleteKeyTemplateOptions *DeleteKeyTemplateOptions)`, func() {
		deleteKeyTemplatePath := "/v4/templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteKeyTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ukoService.DeleteKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteKeyTemplateOptions model
				deleteKeyTemplateOptionsModel := new(ukov4.DeleteKeyTemplateOptions)
				deleteKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				deleteKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				deleteKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ukoService.DeleteKeyTemplate(deleteKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeleteKeyTemplateOptions model
				deleteKeyTemplateOptionsModel := new(ukov4.DeleteKeyTemplateOptions)
				deleteKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				deleteKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				deleteKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ukoService.DeleteKeyTemplate(deleteKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteKeyTemplateOptions model with no property values
				deleteKeyTemplateOptionsModelNew := new(ukov4.DeleteKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ukoService.DeleteKeyTemplate(deleteKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeyTemplate(getKeyTemplateOptions *GetKeyTemplateOptions) - Operation response error`, func() {
		getKeyTemplatePath := "/v4/templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyTemplateOptions model
				getKeyTemplateOptionsModel := new(ukov4.GetKeyTemplateOptions)
				getKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				getKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeyTemplate(getKeyTemplateOptions *GetKeyTemplateOptions)`, func() {
		getKeyTemplatePath := "/v4/templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke GetKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetKeyTemplateOptions model
				getKeyTemplateOptionsModel := new(ukov4.GetKeyTemplateOptions)
				getKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				getKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetKeyTemplateWithContext(ctx, getKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetKeyTemplateWithContext(ctx, getKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke GetKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKeyTemplateOptions model
				getKeyTemplateOptionsModel := new(ukov4.GetKeyTemplateOptions)
				getKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				getKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyTemplateOptions model
				getKeyTemplateOptionsModel := new(ukov4.GetKeyTemplateOptions)
				getKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				getKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKeyTemplateOptions model with no property values
				getKeyTemplateOptionsModelNew := new(ukov4.GetKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetKeyTemplate(getKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeyTemplateOptions model
				getKeyTemplateOptionsModel := new(ukov4.GetKeyTemplateOptions)
				getKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				getKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetKeyTemplate(getKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateKeyTemplate(updateKeyTemplateOptions *UpdateKeyTemplateOptions) - Operation response error`, func() {
		updateKeyTemplatePath := "/v4/templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeyTemplatePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}
				updateKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Key = keyPropertiesUpdateModel
				updateKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateKeyTemplate(updateKeyTemplateOptions *UpdateKeyTemplateOptions)`, func() {
		updateKeyTemplatePath := "/v4/templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeyTemplatePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke UpdateKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}
				updateKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Key = keyPropertiesUpdateModel
				updateKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UpdateKeyTemplateWithContext(ctx, updateKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UpdateKeyTemplateWithContext(ctx, updateKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeyTemplatePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke UpdateKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UpdateKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}
				updateKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Key = keyPropertiesUpdateModel
				updateKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}
				updateKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Key = keyPropertiesUpdateModel
				updateKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateKeyTemplateOptions model with no property values
				updateKeyTemplateOptionsModelNew := new(ukov4.UpdateKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}
				updateKeyTemplateOptionsModel.Description = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Key = keyPropertiesUpdateModel
				updateKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UpdateKeyTemplate(updateKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListKeystores(listKeystoresOptions *ListKeystoresOptions) - Operation response error`, func() {
		listKeystoresPath := "/v4/keystores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeystoresPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["group"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListKeystores with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := new(ukov4.ListKeystoresOptions)
				listKeystoresOptionsModel.Accept = core.StringPtr("application/json")
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Group = core.StringPtr("testString")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = []string{"testString"}
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
				listKeystoresOptionsModel.StatusHealthStatus = []string{"ok"}
				listKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListKeystores(listKeystoresOptions *ListKeystoresOptions)`, func() {
		listKeystoresPath := "/v4/keystores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeystoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["group"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "keystores": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}]}`)
				}))
			})
			It(`Invoke ListKeystores successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := new(ukov4.ListKeystoresOptions)
				listKeystoresOptionsModel.Accept = core.StringPtr("application/json")
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Group = core.StringPtr("testString")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = []string{"testString"}
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
				listKeystoresOptionsModel.StatusHealthStatus = []string{"ok"}
				listKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListKeystoresWithContext(ctx, listKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListKeystoresWithContext(ctx, listKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeystoresPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["group"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "keystores": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}]}`)
				}))
			})
			It(`Invoke ListKeystores successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListKeystores(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := new(ukov4.ListKeystoresOptions)
				listKeystoresOptionsModel.Accept = core.StringPtr("application/json")
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Group = core.StringPtr("testString")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = []string{"testString"}
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
				listKeystoresOptionsModel.StatusHealthStatus = []string{"ok"}
				listKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListKeystores with error: Operation request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := new(ukov4.ListKeystoresOptions)
				listKeystoresOptionsModel.Accept = core.StringPtr("application/json")
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Group = core.StringPtr("testString")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = []string{"testString"}
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
				listKeystoresOptionsModel.StatusHealthStatus = []string{"ok"}
				listKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListKeystores successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := new(ukov4.ListKeystoresOptions)
				listKeystoresOptionsModel.Accept = core.StringPtr("application/json")
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Group = core.StringPtr("testString")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = []string{"testString"}
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
				listKeystoresOptionsModel.StatusHealthStatus = []string{"ok"}
				listKeystoresOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListKeystores(listKeystoresOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.KeystoreList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.KeystoreList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.KeystoreList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.KeystoreList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeystoresPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"keystores":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Accounting","location":"us-south","description":"IBM Cloud keystore for testing","groups":["Production"],"type":"ibm_cloud_kms","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","status":{"last_heartbeat":"2019-01-01T12:00:00.000Z","health_status":"ok","message":"Ping executed successfully."},"google_credentials":"eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=","google_location":"europe-central2","google_project_id":"demo-project","google_private_key_id":"f871b60d0617be19393bb66ea142887fc9621360","google_key_ring":"my-key-ring"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"keystores":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Accounting","location":"us-south","description":"IBM Cloud keystore for testing","groups":["Production"],"type":"ibm_cloud_kms","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","status":{"last_heartbeat":"2019-01-01T12:00:00.000Z","health_status":"ok","message":"Ping executed successfully."},"google_credentials":"eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=","google_location":"europe-central2","google_project_id":"demo-project","google_private_key_id":"f871b60d0617be19393bb66ea142887fc9621360","google_key_ring":"my-key-ring"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use KeystoresPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listKeystoresOptionsModel := &ukov4.ListKeystoresOptions{
					Accept: core.StringPtr("application/json"),
					Type: []string{"ibm_cloud_kms"},
					Name: core.StringPtr("Main IBM Cloud"),
					Description: core.StringPtr("My Example Keystore Description"),
					Group: core.StringPtr("testString"),
					Groups: core.StringPtr("testString"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Location: []string{"testString"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					StatusHealthStatus: []string{"ok"},
				}

				pager, err := ukoService.NewKeystoresPager(listKeystoresOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.KeystoreIntf
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use KeystoresPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listKeystoresOptionsModel := &ukov4.ListKeystoresOptions{
					Accept: core.StringPtr("application/json"),
					Type: []string{"ibm_cloud_kms"},
					Name: core.StringPtr("Main IBM Cloud"),
					Description: core.StringPtr("My Example Keystore Description"),
					Group: core.StringPtr("testString"),
					Groups: core.StringPtr("testString"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Location: []string{"testString"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					StatusHealthStatus: []string{"ok"},
				}

				pager, err := ukoService.NewKeystoresPager(listKeystoresOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateKeystore(createKeystoreOptions *CreateKeystoreOptions) - Operation response error`, func() {
		createKeystorePath := "/v4/keystores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeystorePath))
					Expect(req.Method).To(Equal("POST"))
					// TODO: Add check for dry_run query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateKeystore with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.KeystoreBody = keystoreCreationRequestModel
				createKeystoreOptionsModel.DryRun = core.BoolPtr(false)
				createKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateKeystore(createKeystoreOptions *CreateKeystoreOptions)`, func() {
		createKeystorePath := "/v4/keystores"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeystorePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for dry_run query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke CreateKeystore successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.KeystoreBody = keystoreCreationRequestModel
				createKeystoreOptionsModel.DryRun = core.BoolPtr(false)
				createKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.CreateKeystoreWithContext(ctx, createKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.CreateKeystoreWithContext(ctx, createKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createKeystorePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for dry_run query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke CreateKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.CreateKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.KeystoreBody = keystoreCreationRequestModel
				createKeystoreOptionsModel.DryRun = core.BoolPtr(false)
				createKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.KeystoreBody = keystoreCreationRequestModel
				createKeystoreOptionsModel.DryRun = core.BoolPtr(false)
				createKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateKeystoreOptions model with no property values
				createKeystoreOptionsModelNew := new(ukov4.CreateKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.CreateKeystore(createKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.KeystoreBody = keystoreCreationRequestModel
				createKeystoreOptionsModel.DryRun = core.BoolPtr(false)
				createKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.CreateKeystore(createKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteKeystore(deleteKeystoreOptions *DeleteKeystoreOptions)`, func() {
		deleteKeystorePath := "/v4/keystores/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteKeystorePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ukoService.DeleteKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteKeystoreOptions model
				deleteKeystoreOptionsModel := new(ukov4.DeleteKeystoreOptions)
				deleteKeystoreOptionsModel.ID = core.StringPtr("testString")
				deleteKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ukoService.DeleteKeystore(deleteKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeleteKeystoreOptions model
				deleteKeystoreOptionsModel := new(ukov4.DeleteKeystoreOptions)
				deleteKeystoreOptionsModel.ID = core.StringPtr("testString")
				deleteKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ukoService.DeleteKeystore(deleteKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteKeystoreOptions model with no property values
				deleteKeystoreOptionsModelNew := new(ukov4.DeleteKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ukoService.DeleteKeystore(deleteKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeystore(getKeystoreOptions *GetKeystoreOptions) - Operation response error`, func() {
		getKeystorePath := "/v4/keystores/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystorePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetKeystore with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreOptions model
				getKeystoreOptionsModel := new(ukov4.GetKeystoreOptions)
				getKeystoreOptionsModel.ID = core.StringPtr("testString")
				getKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeystore(getKeystoreOptions *GetKeystoreOptions)`, func() {
		getKeystorePath := "/v4/keystores/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke GetKeystore successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetKeystoreOptions model
				getKeystoreOptionsModel := new(ukov4.GetKeystoreOptions)
				getKeystoreOptionsModel.ID = core.StringPtr("testString")
				getKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetKeystoreWithContext(ctx, getKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetKeystoreWithContext(ctx, getKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke GetKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKeystoreOptions model
				getKeystoreOptionsModel := new(ukov4.GetKeystoreOptions)
				getKeystoreOptionsModel.ID = core.StringPtr("testString")
				getKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreOptions model
				getKeystoreOptionsModel := new(ukov4.GetKeystoreOptions)
				getKeystoreOptionsModel.ID = core.StringPtr("testString")
				getKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKeystoreOptions model with no property values
				getKeystoreOptionsModelNew := new(ukov4.GetKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetKeystore(getKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreOptions model
				getKeystoreOptionsModel := new(ukov4.GetKeystoreOptions)
				getKeystoreOptionsModel.ID = core.StringPtr("testString")
				getKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetKeystore(getKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateKeystore(updateKeystoreOptions *UpdateKeystoreOptions) - Operation response error`, func() {
		updateKeystorePath := "/v4/keystores/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeystorePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateKeystore with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeystoreOptionsModel.KeystoreBody = keystoreUpdateRequestModel
				updateKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateKeystore(updateKeystoreOptions *UpdateKeystoreOptions)`, func() {
		updateKeystorePath := "/v4/keystores/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeystorePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke UpdateKeystore successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeystoreOptionsModel.KeystoreBody = keystoreUpdateRequestModel
				updateKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UpdateKeystoreWithContext(ctx, updateKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UpdateKeystoreWithContext(ctx, updateKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateKeystorePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "status": {"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}, "google_credentials": "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=", "google_location": "europe-central2", "google_project_id": "demo-project", "google_private_key_id": "f871b60d0617be19393bb66ea142887fc9621360", "google_key_ring": "my-key-ring"}`)
				}))
			})
			It(`Invoke UpdateKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UpdateKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeystoreOptionsModel.KeystoreBody = keystoreUpdateRequestModel
				updateKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeystoreOptionsModel.KeystoreBody = keystoreUpdateRequestModel
				updateKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateKeystoreOptions model with no property values
				updateKeystoreOptionsModelNew := new(ukov4.UpdateKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UpdateKeystore(updateKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeystoreOptionsModel.KeystoreBody = keystoreUpdateRequestModel
				updateKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UpdateKeystore(updateKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptions *ListAssociatedResourcesForTargetKeystoreOptions) - Operation response error`, func() {
		listAssociatedResourcesForTargetKeystorePath := "/v4/keystores/testString/associated_resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForTargetKeystorePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForTargetKeystore with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				listAssociatedResourcesForTargetKeystoreOptionsModel := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				listAssociatedResourcesForTargetKeystoreOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForTargetKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptions *ListAssociatedResourcesForTargetKeystoreOptions)`, func() {
		listAssociatedResourcesForTargetKeystorePath := "/v4/keystores/testString/associated_resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForTargetKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "associated_resources": [{"id": "crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_key": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label": "IBM CLOUD KEY", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "referenced_keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000", "name": "keyprotecttest", "type": "com_ibm_cloud_kms_registration", "com_ibm_cloud_kms_registration": {"prevents_key_deletion": true, "service_name": "cloud-object-storage", "service_instance_name": "Cloud Object Storage-7s", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest", "description": "Example Description"}}]}`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForTargetKeystore successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				listAssociatedResourcesForTargetKeystoreOptionsModel := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				listAssociatedResourcesForTargetKeystoreOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForTargetKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListAssociatedResourcesForTargetKeystoreWithContext(ctx, listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListAssociatedResourcesForTargetKeystoreWithContext(ctx, listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForTargetKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "associated_resources": [{"id": "crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_key": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label": "IBM CLOUD KEY", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "referenced_keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000", "name": "keyprotecttest", "type": "com_ibm_cloud_kms_registration", "com_ibm_cloud_kms_registration": {"prevents_key_deletion": true, "service_name": "cloud-object-storage", "service_instance_name": "Cloud Object Storage-7s", "crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest", "description": "Example Description"}}]}`)
				}))
			})
			It(`Invoke ListAssociatedResourcesForTargetKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListAssociatedResourcesForTargetKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				listAssociatedResourcesForTargetKeystoreOptionsModel := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				listAssociatedResourcesForTargetKeystoreOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForTargetKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListAssociatedResourcesForTargetKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				listAssociatedResourcesForTargetKeystoreOptionsModel := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				listAssociatedResourcesForTargetKeystoreOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForTargetKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAssociatedResourcesForTargetKeystoreOptions model with no property values
				listAssociatedResourcesForTargetKeystoreOptionsModelNew := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAssociatedResourcesForTargetKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				listAssociatedResourcesForTargetKeystoreOptionsModel := new(ukov4.ListAssociatedResourcesForTargetKeystoreOptions)
				listAssociatedResourcesForTargetKeystoreOptionsModel.ID = core.StringPtr("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.Sort = []string{"name"}
				listAssociatedResourcesForTargetKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListAssociatedResourcesForTargetKeystore(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.AssociatedResourceList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAssociatedResourcesForTargetKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"associated_resources":[{"id":"crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"managed_key":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label":"IBM CLOUD KEY","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"referenced_keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000","name":"keyprotecttest","type":"com_ibm_cloud_kms_registration","com_ibm_cloud_kms_registration":{"prevents_key_deletion":true,"service_name":"cloud-object-storage","service_instance_name":"Cloud Object Storage-7s","crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest","description":"Example Description"}}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"associated_resources":[{"id":"crn%3Av1%3Abluemix%3Apublic%3Acloud-object-storage%3Aglobal%3Aa%2Fdb995d8d9cc715cd99f13b0671d978b6%3A57da8e3a-a86d-4e01-b840-f22d36e6f23f%3Abucket%3Akeyprotecttest","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"managed_key":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label":"IBM CLOUD KEY","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"referenced_keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000","name":"keyprotecttest","type":"com_ibm_cloud_kms_registration","com_ibm_cloud_kms_registration":{"prevents_key_deletion":true,"service_name":"cloud-object-storage","service_instance_name":"Cloud Object Storage-7s","crn":"crn:v1:bluemix:public:cloud-object-storage:global:a/db995d8d9cc715cd99f13b0671d978b6:57da8e3a-a86d-4e01-b840-f22d36e6f23f:bucket:keyprotecttest","description":"Example Description"}}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AssociatedResourcesForTargetKeystorePager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listAssociatedResourcesForTargetKeystoreOptionsModel := &ukov4.ListAssociatedResourcesForTargetKeystoreOptions{
					ID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"name"},
				}

				pager, err := ukoService.NewAssociatedResourcesForTargetKeystorePager(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.AssociatedResource
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AssociatedResourcesForTargetKeystorePager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listAssociatedResourcesForTargetKeystoreOptionsModel := &ukov4.ListAssociatedResourcesForTargetKeystoreOptions{
					ID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"name"},
				}

				pager, err := ukoService.NewAssociatedResourcesForTargetKeystorePager(listAssociatedResourcesForTargetKeystoreOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetKeystoreStatus(getKeystoreStatusOptions *GetKeystoreStatusOptions) - Operation response error`, func() {
		getKeystoreStatusPath := "/v4/keystores/testString/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystoreStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetKeystoreStatus with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreStatusOptions model
				getKeystoreStatusOptionsModel := new(ukov4.GetKeystoreStatusOptions)
				getKeystoreStatusOptionsModel.ID = core.StringPtr("testString")
				getKeystoreStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetKeystoreStatus(getKeystoreStatusOptions *GetKeystoreStatusOptions)`, func() {
		getKeystoreStatusPath := "/v4/keystores/testString/status"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystoreStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}`)
				}))
			})
			It(`Invoke GetKeystoreStatus successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetKeystoreStatusOptions model
				getKeystoreStatusOptionsModel := new(ukov4.GetKeystoreStatusOptions)
				getKeystoreStatusOptionsModel.ID = core.StringPtr("testString")
				getKeystoreStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetKeystoreStatusWithContext(ctx, getKeystoreStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetKeystoreStatusWithContext(ctx, getKeystoreStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystoreStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"last_heartbeat": "2019-01-01T12:00:00.000Z", "health_status": "ok", "message": "Ping executed successfully."}`)
				}))
			})
			It(`Invoke GetKeystoreStatus successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetKeystoreStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKeystoreStatusOptions model
				getKeystoreStatusOptionsModel := new(ukov4.GetKeystoreStatusOptions)
				getKeystoreStatusOptionsModel.ID = core.StringPtr("testString")
				getKeystoreStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetKeystoreStatus with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreStatusOptions model
				getKeystoreStatusOptionsModel := new(ukov4.GetKeystoreStatusOptions)
				getKeystoreStatusOptionsModel.ID = core.StringPtr("testString")
				getKeystoreStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKeystoreStatusOptions model with no property values
				getKeystoreStatusOptionsModelNew := new(ukov4.GetKeystoreStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetKeystoreStatus successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetKeystoreStatusOptions model
				getKeystoreStatusOptionsModel := new(ukov4.GetKeystoreStatusOptions)
				getKeystoreStatusOptionsModel.ID = core.StringPtr("testString")
				getKeystoreStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetKeystoreStatus(getKeystoreStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptions *ListManagedKeysFromKeystoreOptions) - Operation response error`, func() {
		listManagedKeysFromKeystorePath := "/v4/keystores/testString/managed_keys"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysFromKeystorePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListManagedKeysFromKeystore with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				listManagedKeysFromKeystoreOptionsModel := new(ukov4.ListManagedKeysFromKeystoreOptions)
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysFromKeystoreOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysFromKeystoreOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysFromKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysFromKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysFromKeystoreOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysFromKeystoreOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysFromKeystoreOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysFromKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptions *ListManagedKeysFromKeystoreOptions)`, func() {
		listManagedKeysFromKeystorePath := "/v4/keystores/testString/managed_keys"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysFromKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeysFromKeystore successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				listManagedKeysFromKeystoreOptionsModel := new(ukov4.ListManagedKeysFromKeystoreOptions)
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysFromKeystoreOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysFromKeystoreOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysFromKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysFromKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysFromKeystoreOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysFromKeystoreOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysFromKeystoreOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysFromKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListManagedKeysFromKeystoreWithContext(ctx, listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListManagedKeysFromKeystoreWithContext(ctx, listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysFromKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["expiration_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["rotated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template.name"]).To(Equal([]string{"AWS-TEMPLATE"}))
					Expect(req.URL.Query()["template.alignment_status"]).To(Equal([]string{"aligned"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "type": ["user_defined"], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "alignment_status": "aligned"}, "version": 1, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "label_tags": [{"name": "Name", "value": "Value"}], "tags": [{"name": "Name", "value": "Value"}], "is_rotatable": false, "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}, "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "rotated_at": "2022-02-22T10:27:08.000Z", "status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "keystore_sync_flag": "ok", "keystore_sync_flag_detail": "pre_active_key_not_present_in_keystore", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}], "deactivate_on_rotation": true}]}`)
				}))
			})
			It(`Invoke ListManagedKeysFromKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListManagedKeysFromKeystore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				listManagedKeysFromKeystoreOptionsModel := new(ukov4.ListManagedKeysFromKeystoreOptions)
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysFromKeystoreOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysFromKeystoreOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysFromKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysFromKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysFromKeystoreOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysFromKeystoreOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysFromKeystoreOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysFromKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListManagedKeysFromKeystore with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				listManagedKeysFromKeystoreOptionsModel := new(ukov4.ListManagedKeysFromKeystoreOptions)
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysFromKeystoreOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysFromKeystoreOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysFromKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysFromKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysFromKeystoreOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysFromKeystoreOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysFromKeystoreOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysFromKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListManagedKeysFromKeystoreOptions model with no property values
				listManagedKeysFromKeystoreOptionsModelNew := new(ukov4.ListManagedKeysFromKeystoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListManagedKeysFromKeystore successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				listManagedKeysFromKeystoreOptionsModel := new(ukov4.ListManagedKeysFromKeystoreOptions)
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Accept = core.StringPtr("application/json")
				listManagedKeysFromKeystoreOptionsModel.Algorithm = []string{"aes"}
				listManagedKeysFromKeystoreOptionsModel.State = []string{"pre_activation", "active"}
				listManagedKeysFromKeystoreOptionsModel.Limit = core.Int64Ptr(int64(10))
				listManagedKeysFromKeystoreOptionsModel.Offset = core.Int64Ptr(int64(0))
				listManagedKeysFromKeystoreOptionsModel.Sort = []string{"-updated_at"}
				listManagedKeysFromKeystoreOptionsModel.Label = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ActivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDate = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.RotatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.TemplateName = core.StringPtr("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.TemplateID = []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}
				listManagedKeysFromKeystoreOptionsModel.TemplateType = []string{"user_defined"}
				listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag = []string{"out_of_sync"}
				listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus = core.StringPtr("aligned")
				listManagedKeysFromKeystoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListManagedKeysFromKeystore(listManagedKeysFromKeystoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.ManagedKeyList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.ManagedKeyList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listManagedKeysFromKeystorePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","type":["user_defined"],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","alignment_status":"aligned"},"version":1,"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","label_tags":[{"name":"Name","value":"Value"}],"tags":[{"name":"Name","value":"Value"}],"is_rotatable":false,"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"},"google_key_protection_level":"software","google_key_purpose":"encrypt_decrypt","google_kms_algorithm":"google_symmetric_encryption"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","rotated_at":"2022-02-22T10:27:08.000Z","status_in_keystores":[{"keystore":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"status":"active","keystore_sync_flag":"ok","keystore_sync_flag_detail":"pre_active_key_not_present_in_keystore","error":{"status_code":400,"trace":"9daee671-916a-4678-850b-10b9110236d","errors":[{"code":"missing_field","message":"The algorithm field is required","more_info":"https://cloud.ibm.com/apidocs/uko#create-managed-key","message_params":["My Key Template"],"target":{"type":"field","name":"first_name"}}]},"key_id_in_keystore":"123e4567-e89b-12d3-a456-426614174000"}],"deactivate_on_rotation":true}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ManagedKeysFromKeystorePager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeysFromKeystoreOptionsModel := &ukov4.ListManagedKeysFromKeystoreOptions{
					ID: core.StringPtr("testString"),
					Accept: core.StringPtr("application/json"),
					Algorithm: []string{"aes"},
					State: []string{"pre_activation", "active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeysFromKeystorePager(listManagedKeysFromKeystoreOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.ManagedKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ManagedKeysFromKeystorePager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listManagedKeysFromKeystoreOptionsModel := &ukov4.ListManagedKeysFromKeystoreOptions{
					ID: core.StringPtr("testString"),
					Accept: core.StringPtr("application/json"),
					Algorithm: []string{"aes"},
					State: []string{"pre_activation", "active"},
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Label: core.StringPtr("testString"),
					ActivationDate: core.StringPtr("testString"),
					ActivationDateMin: core.StringPtr("testString"),
					ActivationDateMax: core.StringPtr("testString"),
					DeactivationDate: core.StringPtr("testString"),
					DeactivationDateMin: core.StringPtr("testString"),
					DeactivationDateMax: core.StringPtr("testString"),
					ExpirationDate: core.StringPtr("testString"),
					ExpirationDateMin: core.StringPtr("testString"),
					ExpirationDateMax: core.StringPtr("testString"),
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					RotatedAtMin: core.StringPtr("testString"),
					RotatedAtMax: core.StringPtr("testString"),
					Size: core.StringPtr("testString"),
					SizeMin: core.StringPtr("testString"),
					SizeMax: core.StringPtr("testString"),
					TemplateName: core.StringPtr("AWS-TEMPLATE"),
					TemplateID: []string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"},
					TemplateType: []string{"user_defined"},
					StatusInKeystoresKeystoreSyncFlag: []string{"out_of_sync"},
					TemplateAlignmentStatus: core.StringPtr("aligned"),
				}

				pager, err := ukoService.NewManagedKeysFromKeystorePager(listManagedKeysFromKeystoreOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`ListVaults(listVaultsOptions *ListVaultsOptions) - Operation response error`, func() {
		listVaultsPath := "/v4/vaults"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVaultsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Vault"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Vault Description"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVaults with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := new(ukov4.ListVaultsOptions)
				listVaultsOptionsModel.Accept = core.StringPtr("application/json")
				listVaultsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listVaultsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listVaultsOptionsModel.Sort = []string{"-updated_at"}
				listVaultsOptionsModel.Name = core.StringPtr("My Example Vault")
				listVaultsOptionsModel.Description = core.StringPtr("My Example Vault Description")
				listVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVaults(listVaultsOptions *ListVaultsOptions)`, func() {
		listVaultsPath := "/v4/vaults"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Vault"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Vault Description"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "vaults": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}]}`)
				}))
			})
			It(`Invoke ListVaults successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := new(ukov4.ListVaultsOptions)
				listVaultsOptionsModel.Accept = core.StringPtr("application/json")
				listVaultsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listVaultsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listVaultsOptionsModel.Sort = []string{"-updated_at"}
				listVaultsOptionsModel.Name = core.StringPtr("My Example Vault")
				listVaultsOptionsModel.Description = core.StringPtr("My Example Vault Description")
				listVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ListVaultsWithContext(ctx, listVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ListVaultsWithContext(ctx, listVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Vault"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Vault Description"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "vaults": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}]}`)
				}))
			})
			It(`Invoke ListVaults successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ListVaults(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := new(ukov4.ListVaultsOptions)
				listVaultsOptionsModel.Accept = core.StringPtr("application/json")
				listVaultsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listVaultsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listVaultsOptionsModel.Sort = []string{"-updated_at"}
				listVaultsOptionsModel.Name = core.StringPtr("My Example Vault")
				listVaultsOptionsModel.Description = core.StringPtr("My Example Vault Description")
				listVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVaults with error: Operation request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := new(ukov4.ListVaultsOptions)
				listVaultsOptionsModel.Accept = core.StringPtr("application/json")
				listVaultsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listVaultsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listVaultsOptionsModel.Sort = []string{"-updated_at"}
				listVaultsOptionsModel.Name = core.StringPtr("My Example Vault")
				listVaultsOptionsModel.Description = core.StringPtr("My Example Vault Description")
				listVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVaults successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := new(ukov4.ListVaultsOptions)
				listVaultsOptionsModel.Accept = core.StringPtr("application/json")
				listVaultsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listVaultsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listVaultsOptionsModel.Sort = []string{"-updated_at"}
				listVaultsOptionsModel.Name = core.StringPtr("My Example Vault")
				listVaultsOptionsModel.Description = core.StringPtr("My Example Vault Description")
				listVaultsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ListVaults(listVaultsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(ukov4.VaultList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(ukov4.VaultList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(ukov4.VaultList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(ukov4.VaultList)
				nextObject := new(ukov4.HrefObject)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVaultsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"vaults":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"My Example Vault","description":"The description of the vault","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","recovery_key_label":"TEKMF.AES.RECOVERY.00001","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","keys_count":1023,"key_templates_count":100,"keystores_count":10}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"vaults":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"My Example Vault","description":"The description of the vault","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","recovery_key_label":"TEKMF.AES.RECOVERY.00001","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","keys_count":1023,"key_templates_count":100,"keystores_count":10}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use VaultsPager.GetNext successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listVaultsOptionsModel := &ukov4.ListVaultsOptions{
					Accept: core.StringPtr("application/json"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Name: core.StringPtr("My Example Vault"),
					Description: core.StringPtr("My Example Vault Description"),
				}

				pager, err := ukoService.NewVaultsPager(listVaultsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []ukov4.Vault
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use VaultsPager.GetAll successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				listVaultsOptionsModel := &ukov4.ListVaultsOptions{
					Accept: core.StringPtr("application/json"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
					Name: core.StringPtr("My Example Vault"),
					Description: core.StringPtr("My Example Vault Description"),
				}

				pager, err := ukoService.NewVaultsPager(listVaultsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`CreateVault(createVaultOptions *CreateVaultOptions) - Operation response error`, func() {
		createVaultPath := "/v4/vaults"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVault with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsModel := new(ukov4.CreateVaultOptions)
				createVaultOptionsModel.Name = core.StringPtr("Example Vault")
				createVaultOptionsModel.Description = core.StringPtr("The description of the creating vault")
				createVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("testString")
				createVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVault(createVaultOptions *CreateVaultOptions)`, func() {
		createVaultPath := "/v4/vaults"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke CreateVault successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsModel := new(ukov4.CreateVaultOptions)
				createVaultOptionsModel.Name = core.StringPtr("Example Vault")
				createVaultOptionsModel.Description = core.StringPtr("The description of the creating vault")
				createVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("testString")
				createVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.CreateVaultWithContext(ctx, createVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.CreateVaultWithContext(ctx, createVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVaultPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke CreateVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.CreateVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsModel := new(ukov4.CreateVaultOptions)
				createVaultOptionsModel.Name = core.StringPtr("Example Vault")
				createVaultOptionsModel.Description = core.StringPtr("The description of the creating vault")
				createVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("testString")
				createVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVault with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsModel := new(ukov4.CreateVaultOptions)
				createVaultOptionsModel.Name = core.StringPtr("Example Vault")
				createVaultOptionsModel.Description = core.StringPtr("The description of the creating vault")
				createVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("testString")
				createVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVaultOptions model with no property values
				createVaultOptionsModelNew := new(ukov4.CreateVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.CreateVault(createVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsModel := new(ukov4.CreateVaultOptions)
				createVaultOptionsModel.Name = core.StringPtr("Example Vault")
				createVaultOptionsModel.Description = core.StringPtr("The description of the creating vault")
				createVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("testString")
				createVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.CreateVault(createVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVault(deleteVaultOptions *DeleteVaultOptions)`, func() {
		deleteVaultPath := "/v4/vaults/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVaultPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ukoService.DeleteVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVaultOptions model
				deleteVaultOptionsModel := new(ukov4.DeleteVaultOptions)
				deleteVaultOptionsModel.ID = core.StringPtr("testString")
				deleteVaultOptionsModel.IfMatch = core.StringPtr("testString")
				deleteVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ukoService.DeleteVault(deleteVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVault with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the DeleteVaultOptions model
				deleteVaultOptionsModel := new(ukov4.DeleteVaultOptions)
				deleteVaultOptionsModel.ID = core.StringPtr("testString")
				deleteVaultOptionsModel.IfMatch = core.StringPtr("testString")
				deleteVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ukoService.DeleteVault(deleteVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVaultOptions model with no property values
				deleteVaultOptionsModelNew := new(ukov4.DeleteVaultOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ukoService.DeleteVault(deleteVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVault(getVaultOptions *GetVaultOptions) - Operation response error`, func() {
		getVaultPath := "/v4/vaults/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVaultPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVault with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetVaultOptions model
				getVaultOptionsModel := new(ukov4.GetVaultOptions)
				getVaultOptionsModel.ID = core.StringPtr("testString")
				getVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVault(getVaultOptions *GetVaultOptions)`, func() {
		getVaultPath := "/v4/vaults/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVaultPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke GetVault successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the GetVaultOptions model
				getVaultOptionsModel := new(ukov4.GetVaultOptions)
				getVaultOptionsModel.ID = core.StringPtr("testString")
				getVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.GetVaultWithContext(ctx, getVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.GetVaultWithContext(ctx, getVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVaultPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke GetVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.GetVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVaultOptions model
				getVaultOptionsModel := new(ukov4.GetVaultOptions)
				getVaultOptionsModel.ID = core.StringPtr("testString")
				getVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVault with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetVaultOptions model
				getVaultOptionsModel := new(ukov4.GetVaultOptions)
				getVaultOptionsModel.ID = core.StringPtr("testString")
				getVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVaultOptions model with no property values
				getVaultOptionsModelNew := new(ukov4.GetVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.GetVault(getVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the GetVaultOptions model
				getVaultOptionsModel := new(ukov4.GetVaultOptions)
				getVaultOptionsModel.ID = core.StringPtr("testString")
				getVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.GetVault(getVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVault(updateVaultOptions *UpdateVaultOptions) - Operation response error`, func() {
		updateVaultPath := "/v4/vaults/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVaultPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateVault with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateVaultOptions model
				updateVaultOptionsModel := new(ukov4.UpdateVaultOptions)
				updateVaultOptionsModel.ID = core.StringPtr("testString")
				updateVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateVaultOptionsModel.Name = core.StringPtr("Jakub's Vault")
				updateVaultOptionsModel.Description = core.StringPtr("Updated description of the vault")
				updateVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVault(updateVaultOptions *UpdateVaultOptions)`, func() {
		updateVaultPath := "/v4/vaults/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVaultPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke UpdateVault successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the UpdateVaultOptions model
				updateVaultOptionsModel := new(ukov4.UpdateVaultOptions)
				updateVaultOptionsModel.ID = core.StringPtr("testString")
				updateVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateVaultOptionsModel.Name = core.StringPtr("Jakub's Vault")
				updateVaultOptionsModel.Description = core.StringPtr("Updated description of the vault")
				updateVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UpdateVaultWithContext(ctx, updateVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UpdateVaultWithContext(ctx, updateVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVaultPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "recovery_key_label": "TEKMF.AES.RECOVERY.00001", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "keys_count": 1023, "key_templates_count": 100, "keystores_count": 10}`)
				}))
			})
			It(`Invoke UpdateVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UpdateVault(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVaultOptions model
				updateVaultOptionsModel := new(ukov4.UpdateVaultOptions)
				updateVaultOptionsModel.ID = core.StringPtr("testString")
				updateVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateVaultOptionsModel.Name = core.StringPtr("Jakub's Vault")
				updateVaultOptionsModel.Description = core.StringPtr("Updated description of the vault")
				updateVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateVault with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateVaultOptions model
				updateVaultOptionsModel := new(ukov4.UpdateVaultOptions)
				updateVaultOptionsModel.ID = core.StringPtr("testString")
				updateVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateVaultOptionsModel.Name = core.StringPtr("Jakub's Vault")
				updateVaultOptionsModel.Description = core.StringPtr("Updated description of the vault")
				updateVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateVaultOptions model with no property values
				updateVaultOptionsModelNew := new(ukov4.UpdateVaultOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UpdateVault(updateVaultOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateVault successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UpdateVaultOptions model
				updateVaultOptionsModel := new(ukov4.UpdateVaultOptions)
				updateVaultOptionsModel.ID = core.StringPtr("testString")
				updateVaultOptionsModel.IfMatch = core.StringPtr("testString")
				updateVaultOptionsModel.Name = core.StringPtr("Jakub's Vault")
				updateVaultOptionsModel.Description = core.StringPtr("Updated description of the vault")
				updateVaultOptionsModel.RecoveryKeyLabel = core.StringPtr("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UpdateVault(updateVaultOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnarchiveKeyTemplate(unarchiveKeyTemplateOptions *UnarchiveKeyTemplateOptions) - Operation response error`, func() {
		unarchiveKeyTemplatePath := "/v4/templates/testString/unarchive"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unarchiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UnarchiveKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UnarchiveKeyTemplateOptions model
				unarchiveKeyTemplateOptionsModel := new(ukov4.UnarchiveKeyTemplateOptions)
				unarchiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnarchiveKeyTemplate(unarchiveKeyTemplateOptions *UnarchiveKeyTemplateOptions)`, func() {
		unarchiveKeyTemplatePath := "/v4/templates/testString/unarchive"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unarchiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke UnarchiveKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the UnarchiveKeyTemplateOptions model
				unarchiveKeyTemplateOptionsModel := new(ukov4.UnarchiveKeyTemplateOptions)
				unarchiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.UnarchiveKeyTemplateWithContext(ctx, unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.UnarchiveKeyTemplateWithContext(ctx, unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(unarchiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke UnarchiveKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.UnarchiveKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UnarchiveKeyTemplateOptions model
				unarchiveKeyTemplateOptionsModel := new(ukov4.UnarchiveKeyTemplateOptions)
				unarchiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UnarchiveKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UnarchiveKeyTemplateOptions model
				unarchiveKeyTemplateOptionsModel := new(ukov4.UnarchiveKeyTemplateOptions)
				unarchiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UnarchiveKeyTemplateOptions model with no property values
				unarchiveKeyTemplateOptionsModelNew := new(ukov4.UnarchiveKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UnarchiveKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the UnarchiveKeyTemplateOptions model
				unarchiveKeyTemplateOptionsModel := new(ukov4.UnarchiveKeyTemplateOptions)
				unarchiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				unarchiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.UnarchiveKeyTemplate(unarchiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ArchiveKeyTemplate(archiveKeyTemplateOptions *ArchiveKeyTemplateOptions) - Operation response error`, func() {
		archiveKeyTemplatePath := "/v4/templates/testString/archive"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(archiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ArchiveKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ArchiveKeyTemplateOptions model
				archiveKeyTemplateOptionsModel := new(ukov4.ArchiveKeyTemplateOptions)
				archiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ArchiveKeyTemplate(archiveKeyTemplateOptions *ArchiveKeyTemplateOptions)`, func() {
		archiveKeyTemplatePath := "/v4/templates/testString/archive"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(archiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke ArchiveKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ArchiveKeyTemplateOptions model
				archiveKeyTemplateOptionsModel := new(ukov4.ArchiveKeyTemplateOptions)
				archiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ArchiveKeyTemplateWithContext(ctx, archiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ArchiveKeyTemplateWithContext(ctx, archiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(archiveKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke ArchiveKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ArchiveKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ArchiveKeyTemplateOptions model
				archiveKeyTemplateOptionsModel := new(ukov4.ArchiveKeyTemplateOptions)
				archiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ArchiveKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ArchiveKeyTemplateOptions model
				archiveKeyTemplateOptionsModel := new(ukov4.ArchiveKeyTemplateOptions)
				archiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ArchiveKeyTemplateOptions model with no property values
				archiveKeyTemplateOptionsModelNew := new(ukov4.ArchiveKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ArchiveKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ArchiveKeyTemplateOptions model
				archiveKeyTemplateOptionsModel := new(ukov4.ArchiveKeyTemplateOptions)
				archiveKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				archiveKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ArchiveKeyTemplate(archiveKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExposeKeyTemplate(exposeKeyTemplateOptions *ExposeKeyTemplateOptions) - Operation response error`, func() {
		exposeKeyTemplatePath := "/v4/templates/testString/expose"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exposeKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ExposeKeyTemplate with error: Operation response processing error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ExposeKeyTemplateOptions model
				exposeKeyTemplateOptionsModel := new(ukov4.ExposeKeyTemplateOptions)
				exposeKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ukoService.EnableRetries(0, 0)
				result, response, operationErr = ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExposeKeyTemplate(exposeKeyTemplateOptions *ExposeKeyTemplateOptions)`, func() {
		exposeKeyTemplatePath := "/v4/templates/testString/expose"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exposeKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke ExposeKeyTemplate successfully with retries`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())
				ukoService.EnableRetries(0, 0)

				// Construct an instance of the ExposeKeyTemplateOptions model
				exposeKeyTemplateOptionsModel := new(ukov4.ExposeKeyTemplateOptions)
				exposeKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ukoService.ExposeKeyTemplateWithContext(ctx, exposeKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ukoService.DisableRetries()
				result, response, operationErr := ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ukoService.ExposeKeyTemplateWithContext(ctx, exposeKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exposeKeyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": ["user_defined"], "state": "unarchived", "keys_count": 3456, "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active", "deactivate_on_rotation": true}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "naming_scheme": "A-<APP>-AES256-<ENV>-<GROUP>", "type": "ibm_cloud_kms", "google_key_protection_level": "software", "google_key_purpose": "encrypt_decrypt", "google_kms_algorithm": "google_symmetric_encryption"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
				}))
			})
			It(`Invoke ExposeKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ukoService.ExposeKeyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExposeKeyTemplateOptions model
				exposeKeyTemplateOptionsModel := new(ukov4.ExposeKeyTemplateOptions)
				exposeKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ExposeKeyTemplate with error: Operation validation and request error`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ExposeKeyTemplateOptions model
				exposeKeyTemplateOptionsModel := new(ukov4.ExposeKeyTemplateOptions)
				exposeKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ukoService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExposeKeyTemplateOptions model with no property values
				exposeKeyTemplateOptionsModelNew := new(ukov4.ExposeKeyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ExposeKeyTemplate successfully`, func() {
				ukoService, serviceErr := ukov4.NewUkoV4(&ukov4.UkoV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ukoService).ToNot(BeNil())

				// Construct an instance of the ExposeKeyTemplateOptions model
				exposeKeyTemplateOptionsModel := new(ukov4.ExposeKeyTemplateOptions)
				exposeKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				exposeKeyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ukoService.ExposeKeyTemplate(exposeKeyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ukoService, _ := ukov4.NewUkoV4(&ukov4.UkoV4Options{
				URL:           "http://ukov4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewActivateManagedKeyOptions successfully`, func() {
				// Construct an instance of the ActivateManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				activateManagedKeyOptionsModel := ukoService.NewActivateManagedKeyOptions(id, ifMatch)
				activateManagedKeyOptionsModel.SetID("testString")
				activateManagedKeyOptionsModel.SetIfMatch("testString")
				activateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(activateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(activateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(activateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(activateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewArchiveKeyTemplateOptions successfully`, func() {
				// Construct an instance of the ArchiveKeyTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				archiveKeyTemplateOptionsModel := ukoService.NewArchiveKeyTemplateOptions(id, ifMatch)
				archiveKeyTemplateOptionsModel.SetID("testString")
				archiveKeyTemplateOptionsModel.SetIfMatch("testString")
				archiveKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(archiveKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(archiveKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(archiveKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(archiveKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateKeyTemplateOptions successfully`, func() {
				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				Expect(vaultReferenceInCreationRequestModel).ToNot(BeNil())
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")
				Expect(vaultReferenceInCreationRequestModel.ID).To(Equal(core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")))

				// Construct an instance of the KeyProperties model
				keyPropertiesModel := new(ukov4.KeyProperties)
				Expect(keyPropertiesModel).ToNot(BeNil())
				keyPropertiesModel.Size = core.StringPtr("256")
				keyPropertiesModel.Algorithm = core.StringPtr("aes")
				keyPropertiesModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesModel.State = core.StringPtr("active")
				keyPropertiesModel.DeactivateOnRotation = core.BoolPtr(true)
				Expect(keyPropertiesModel.Size).To(Equal(core.StringPtr("256")))
				Expect(keyPropertiesModel.Algorithm).To(Equal(core.StringPtr("aes")))
				Expect(keyPropertiesModel.ActivationDate).To(Equal(core.StringPtr("P5Y1M1W2D")))
				Expect(keyPropertiesModel.ExpirationDate).To(Equal(core.StringPtr("P1Y2M1W4D")))
				Expect(keyPropertiesModel.State).To(Equal(core.StringPtr("active")))
				Expect(keyPropertiesModel.DeactivateOnRotation).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the KeystoresPropertiesCreateGoogleKms model
				keystoresPropertiesCreateModel := new(ukov4.KeystoresPropertiesCreateGoogleKms)
				Expect(keystoresPropertiesCreateModel).ToNot(BeNil())
				keystoresPropertiesCreateModel.Group = core.StringPtr("Production")
				keystoresPropertiesCreateModel.NamingScheme = core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")
				keystoresPropertiesCreateModel.Type = core.StringPtr("ibm_cloud_kms")
				keystoresPropertiesCreateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesCreateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesCreateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")
				Expect(keystoresPropertiesCreateModel.Group).To(Equal(core.StringPtr("Production")))
				Expect(keystoresPropertiesCreateModel.NamingScheme).To(Equal(core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")))
				Expect(keystoresPropertiesCreateModel.Type).To(Equal(core.StringPtr("ibm_cloud_kms")))
				Expect(keystoresPropertiesCreateModel.GoogleKeyProtectionLevel).To(Equal(core.StringPtr("software")))
				Expect(keystoresPropertiesCreateModel.GoogleKeyPurpose).To(Equal(core.StringPtr("encrypt_decrypt")))
				Expect(keystoresPropertiesCreateModel.GoogleKmsAlgorithm).To(Equal(core.StringPtr("google_symmetric_encryption")))

				// Construct an instance of the CreateKeyTemplateOptions model
				var createKeyTemplateOptionsVault *ukov4.VaultReferenceInCreationRequest = nil
				createKeyTemplateOptionsName := "EXAMPLE-TEMPLATE"
				var createKeyTemplateOptionsKey *ukov4.KeyProperties = nil
				createKeyTemplateOptionsKeystores := []ukov4.KeystoresPropertiesCreateIntf{}
				createKeyTemplateOptionsModel := ukoService.NewCreateKeyTemplateOptions(createKeyTemplateOptionsVault, createKeyTemplateOptionsName, createKeyTemplateOptionsKey, createKeyTemplateOptionsKeystores)
				createKeyTemplateOptionsModel.SetVault(vaultReferenceInCreationRequestModel)
				createKeyTemplateOptionsModel.SetName("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.SetKey(keyPropertiesModel)
				createKeyTemplateOptionsModel.SetKeystores([]ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel})
				createKeyTemplateOptionsModel.SetDescription("testString")
				createKeyTemplateOptionsModel.SetNamingScheme("A-<APP>-AES256-<ENV>-<GROUP>")
				createKeyTemplateOptionsModel.SetType([]string{"user_defined"})
				createKeyTemplateOptionsModel.SetState("unarchived")
				createKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(createKeyTemplateOptionsModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(createKeyTemplateOptionsModel.Name).To(Equal(core.StringPtr("EXAMPLE-TEMPLATE")))
				Expect(createKeyTemplateOptionsModel.Key).To(Equal(keyPropertiesModel))
				Expect(createKeyTemplateOptionsModel.Keystores).To(Equal([]ukov4.KeystoresPropertiesCreateIntf{keystoresPropertiesCreateModel}))
				Expect(createKeyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createKeyTemplateOptionsModel.NamingScheme).To(Equal(core.StringPtr("A-<APP>-AES256-<ENV>-<GROUP>")))
				Expect(createKeyTemplateOptionsModel.Type).To(Equal([]string{"user_defined"}))
				Expect(createKeyTemplateOptionsModel.State).To(Equal(core.StringPtr("unarchived")))
				Expect(createKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateKeystoreOptions successfully`, func() {
				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				Expect(vaultReferenceInCreationRequestModel).ToNot(BeNil())
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")
				Expect(vaultReferenceInCreationRequestModel.ID).To(Equal(core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")))

				// Construct an instance of the KeystoreCreationRequestKeystoreTypeAwsKmsCreate model
				keystoreCreationRequestModel := new(ukov4.KeystoreCreationRequestKeystoreTypeAwsKmsCreate)
				Expect(keystoreCreationRequestModel).ToNot(BeNil())
				keystoreCreationRequestModel.Type = core.StringPtr("aws_kms")
				keystoreCreationRequestModel.Vault = vaultReferenceInCreationRequestModel
				keystoreCreationRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreCreationRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreCreationRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")
				Expect(keystoreCreationRequestModel.Type).To(Equal(core.StringPtr("aws_kms")))
				Expect(keystoreCreationRequestModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(keystoreCreationRequestModel.Name).To(Equal(core.StringPtr("IBM Cloud Keystore Name")))
				Expect(keystoreCreationRequestModel.Description).To(Equal(core.StringPtr("Azure keystore")))
				Expect(keystoreCreationRequestModel.TlsProxy).To(Equal(core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")))
				Expect(keystoreCreationRequestModel.Groups).To(Equal([]string{"Production"}))
				Expect(keystoreCreationRequestModel.AwsRegion).To(Equal(core.StringPtr("af_south_1")))
				Expect(keystoreCreationRequestModel.AwsAccessKeyID).To(Equal(core.StringPtr("BSDFWERUANLKJDN54AAS")))
				Expect(keystoreCreationRequestModel.AwsSecretAccessKey).To(Equal(core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")))

				// Construct an instance of the CreateKeystoreOptions model
				var keystoreBody ukov4.KeystoreCreationRequestIntf = nil
				createKeystoreOptionsModel := ukoService.NewCreateKeystoreOptions(keystoreBody)
				createKeystoreOptionsModel.SetKeystoreBody(keystoreCreationRequestModel)
				createKeystoreOptionsModel.SetDryRun(false)
				createKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createKeystoreOptionsModel).ToNot(BeNil())
				Expect(createKeystoreOptionsModel.KeystoreBody).To(Equal(keystoreCreationRequestModel))
				Expect(createKeystoreOptionsModel.DryRun).To(Equal(core.BoolPtr(false)))
				Expect(createKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateManagedKeyOptions successfully`, func() {
				// Construct an instance of the VaultReferenceInCreationRequest model
				vaultReferenceInCreationRequestModel := new(ukov4.VaultReferenceInCreationRequest)
				Expect(vaultReferenceInCreationRequestModel).ToNot(BeNil())
				vaultReferenceInCreationRequestModel.ID = core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")
				Expect(vaultReferenceInCreationRequestModel.ID).To(Equal(core.StringPtr("5295ad47-2ce9-43c3-b9e7-e5a9482c362b")))

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsTemplateName := "testString"
				var createManagedKeyOptionsVault *ukov4.VaultReferenceInCreationRequest = nil
				createManagedKeyOptionsModel := ukoService.NewCreateManagedKeyOptions(createManagedKeyOptionsTemplateName, createManagedKeyOptionsVault)
				createManagedKeyOptionsModel.SetTemplateName("testString")
				createManagedKeyOptionsModel.SetVault(vaultReferenceInCreationRequestModel)
				createManagedKeyOptionsModel.SetLabel("IBM CLOUD KEY")
				createManagedKeyOptionsModel.SetDescription("testString")
				createManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createManagedKeyOptionsModel).ToNot(BeNil())
				Expect(createManagedKeyOptionsModel.TemplateName).To(Equal(core.StringPtr("testString")))
				Expect(createManagedKeyOptionsModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(createManagedKeyOptionsModel.Label).To(Equal(core.StringPtr("IBM CLOUD KEY")))
				Expect(createManagedKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVaultOptions successfully`, func() {
				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsName := "Example Vault"
				createVaultOptionsModel := ukoService.NewCreateVaultOptions(createVaultOptionsName)
				createVaultOptionsModel.SetName("Example Vault")
				createVaultOptionsModel.SetDescription("The description of the creating vault")
				createVaultOptionsModel.SetRecoveryKeyLabel("testString")
				createVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVaultOptionsModel).ToNot(BeNil())
				Expect(createVaultOptionsModel.Name).To(Equal(core.StringPtr("Example Vault")))
				Expect(createVaultOptionsModel.Description).To(Equal(core.StringPtr("The description of the creating vault")))
				Expect(createVaultOptionsModel.RecoveryKeyLabel).To(Equal(core.StringPtr("testString")))
				Expect(createVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeactivateManagedKeyOptions successfully`, func() {
				// Construct an instance of the DeactivateManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				deactivateManagedKeyOptionsModel := ukoService.NewDeactivateManagedKeyOptions(id, ifMatch)
				deactivateManagedKeyOptionsModel.SetID("testString")
				deactivateManagedKeyOptionsModel.SetIfMatch("testString")
				deactivateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deactivateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(deactivateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deactivateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deactivateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteKeyTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteKeyTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				deleteKeyTemplateOptionsModel := ukoService.NewDeleteKeyTemplateOptions(id, ifMatch)
				deleteKeyTemplateOptionsModel.SetID("testString")
				deleteKeyTemplateOptionsModel.SetIfMatch("testString")
				deleteKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteKeystoreOptions successfully`, func() {
				// Construct an instance of the DeleteKeystoreOptions model
				id := "testString"
				deleteKeystoreOptionsModel := ukoService.NewDeleteKeystoreOptions(id)
				deleteKeystoreOptionsModel.SetID("testString")
				deleteKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteKeystoreOptionsModel).ToNot(BeNil())
				Expect(deleteKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteManagedKeyOptions successfully`, func() {
				// Construct an instance of the DeleteManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				deleteManagedKeyOptionsModel := ukoService.NewDeleteManagedKeyOptions(id, ifMatch)
				deleteManagedKeyOptionsModel.SetID("testString")
				deleteManagedKeyOptionsModel.SetIfMatch("testString")
				deleteManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteManagedKeyOptionsModel).ToNot(BeNil())
				Expect(deleteManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVaultOptions successfully`, func() {
				// Construct an instance of the DeleteVaultOptions model
				id := "testString"
				ifMatch := "testString"
				deleteVaultOptionsModel := ukoService.NewDeleteVaultOptions(id, ifMatch)
				deleteVaultOptionsModel.SetID("testString")
				deleteVaultOptionsModel.SetIfMatch("testString")
				deleteVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVaultOptionsModel).ToNot(BeNil())
				Expect(deleteVaultOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVaultOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDestroyManagedKeyOptions successfully`, func() {
				// Construct an instance of the DestroyManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				destroyManagedKeyOptionsModel := ukoService.NewDestroyManagedKeyOptions(id, ifMatch)
				destroyManagedKeyOptionsModel.SetID("testString")
				destroyManagedKeyOptionsModel.SetIfMatch("testString")
				destroyManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(destroyManagedKeyOptionsModel).ToNot(BeNil())
				Expect(destroyManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(destroyManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(destroyManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExposeKeyTemplateOptions successfully`, func() {
				// Construct an instance of the ExposeKeyTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				exposeKeyTemplateOptionsModel := ukoService.NewExposeKeyTemplateOptions(id, ifMatch)
				exposeKeyTemplateOptionsModel.SetID("testString")
				exposeKeyTemplateOptionsModel.SetIfMatch("testString")
				exposeKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(exposeKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(exposeKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(exposeKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(exposeKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeyDistributionStatusForKeystoresOptions successfully`, func() {
				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				id := "testString"
				getKeyDistributionStatusForKeystoresOptionsModel := ukoService.NewGetKeyDistributionStatusForKeystoresOptions(id)
				getKeyDistributionStatusForKeystoresOptionsModel.SetID("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeyDistributionStatusForKeystoresOptionsModel).ToNot(BeNil())
				Expect(getKeyDistributionStatusForKeystoresOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeyDistributionStatusForKeystoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeyTemplateOptions successfully`, func() {
				// Construct an instance of the GetKeyTemplateOptions model
				id := "testString"
				getKeyTemplateOptionsModel := ukoService.NewGetKeyTemplateOptions(id)
				getKeyTemplateOptionsModel.SetID("testString")
				getKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(getKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeystoreOptions successfully`, func() {
				// Construct an instance of the GetKeystoreOptions model
				id := "testString"
				getKeystoreOptionsModel := ukoService.NewGetKeystoreOptions(id)
				getKeystoreOptionsModel.SetID("testString")
				getKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeystoreOptionsModel).ToNot(BeNil())
				Expect(getKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeystoreStatusOptions successfully`, func() {
				// Construct an instance of the GetKeystoreStatusOptions model
				id := "testString"
				getKeystoreStatusOptionsModel := ukoService.NewGetKeystoreStatusOptions(id)
				getKeystoreStatusOptionsModel.SetID("testString")
				getKeystoreStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeystoreStatusOptionsModel).ToNot(BeNil())
				Expect(getKeystoreStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetManagedKeyOptions successfully`, func() {
				// Construct an instance of the GetManagedKeyOptions model
				id := "testString"
				getManagedKeyOptionsModel := ukoService.NewGetManagedKeyOptions(id)
				getManagedKeyOptionsModel.SetID("testString")
				getManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getManagedKeyOptionsModel).ToNot(BeNil())
				Expect(getManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVaultOptions successfully`, func() {
				// Construct an instance of the GetVaultOptions model
				id := "testString"
				getVaultOptionsModel := ukoService.NewGetVaultOptions(id)
				getVaultOptionsModel.SetID("testString")
				getVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVaultOptionsModel).ToNot(BeNil())
				Expect(getVaultOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAssociatedResourcesForManagedKeyOptions successfully`, func() {
				// Construct an instance of the ListAssociatedResourcesForManagedKeyOptions model
				id := "testString"
				listAssociatedResourcesForManagedKeyOptionsModel := ukoService.NewListAssociatedResourcesForManagedKeyOptions(id)
				listAssociatedResourcesForManagedKeyOptionsModel.SetID("testString")
				listAssociatedResourcesForManagedKeyOptionsModel.SetLimit(int64(10))
				listAssociatedResourcesForManagedKeyOptionsModel.SetOffset(int64(0))
				listAssociatedResourcesForManagedKeyOptionsModel.SetSort([]string{"name"})
				listAssociatedResourcesForManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAssociatedResourcesForManagedKeyOptionsModel).ToNot(BeNil())
				Expect(listAssociatedResourcesForManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listAssociatedResourcesForManagedKeyOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAssociatedResourcesForManagedKeyOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listAssociatedResourcesForManagedKeyOptionsModel.Sort).To(Equal([]string{"name"}))
				Expect(listAssociatedResourcesForManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAssociatedResourcesForTargetKeystoreOptions successfully`, func() {
				// Construct an instance of the ListAssociatedResourcesForTargetKeystoreOptions model
				id := "testString"
				listAssociatedResourcesForTargetKeystoreOptionsModel := ukoService.NewListAssociatedResourcesForTargetKeystoreOptions(id)
				listAssociatedResourcesForTargetKeystoreOptionsModel.SetID("testString")
				listAssociatedResourcesForTargetKeystoreOptionsModel.SetLimit(int64(10))
				listAssociatedResourcesForTargetKeystoreOptionsModel.SetOffset(int64(0))
				listAssociatedResourcesForTargetKeystoreOptionsModel.SetSort([]string{"name"})
				listAssociatedResourcesForTargetKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel).ToNot(BeNil())
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel.Sort).To(Equal([]string{"name"}))
				Expect(listAssociatedResourcesForTargetKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListKeyTemplatesOptions successfully`, func() {
				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := ukoService.NewListKeyTemplatesOptions()
				listKeyTemplatesOptionsModel.SetAccept("application/json")
				listKeyTemplatesOptionsModel.SetName("My Example Template")
				listKeyTemplatesOptionsModel.SetNamingScheme("My Example Template")
				listKeyTemplatesOptionsModel.SetVaultID([]string{"123e4567-e89b-12d3-a456-426614174000"})
				listKeyTemplatesOptionsModel.SetKeyAlgorithm([]string{"aes"})
				listKeyTemplatesOptionsModel.SetKeySize("testString")
				listKeyTemplatesOptionsModel.SetKeySizeMin("testString")
				listKeyTemplatesOptionsModel.SetKeySizeMax("testString")
				listKeyTemplatesOptionsModel.SetKeystoresType([]string{"ibm_cloud_kms"})
				listKeyTemplatesOptionsModel.SetKeystoresGroup([]string{"testString"})
				listKeyTemplatesOptionsModel.SetCreatedAt("testString")
				listKeyTemplatesOptionsModel.SetCreatedAtMin("testString")
				listKeyTemplatesOptionsModel.SetCreatedAtMax("testString")
				listKeyTemplatesOptionsModel.SetUpdatedAt("testString")
				listKeyTemplatesOptionsModel.SetUpdatedAtMin("testString")
				listKeyTemplatesOptionsModel.SetUpdatedAtMax("testString")
				listKeyTemplatesOptionsModel.SetType([]string{"user_defined"})
				listKeyTemplatesOptionsModel.SetState([]string{"unarchived"})
				listKeyTemplatesOptionsModel.SetSort([]string{"-updated_at"})
				listKeyTemplatesOptionsModel.SetLimit(int64(10))
				listKeyTemplatesOptionsModel.SetOffset(int64(0))
				listKeyTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listKeyTemplatesOptionsModel).ToNot(BeNil())
				Expect(listKeyTemplatesOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(listKeyTemplatesOptionsModel.Name).To(Equal(core.StringPtr("My Example Template")))
				Expect(listKeyTemplatesOptionsModel.NamingScheme).To(Equal(core.StringPtr("My Example Template")))
				Expect(listKeyTemplatesOptionsModel.VaultID).To(Equal([]string{"123e4567-e89b-12d3-a456-426614174000"}))
				Expect(listKeyTemplatesOptionsModel.KeyAlgorithm).To(Equal([]string{"aes"}))
				Expect(listKeyTemplatesOptionsModel.KeySize).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.KeySizeMin).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.KeySizeMax).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.KeystoresType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listKeyTemplatesOptionsModel.KeystoresGroup).To(Equal([]string{"testString"}))
				Expect(listKeyTemplatesOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listKeyTemplatesOptionsModel.Type).To(Equal([]string{"user_defined"}))
				Expect(listKeyTemplatesOptionsModel.State).To(Equal([]string{"unarchived"}))
				Expect(listKeyTemplatesOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listKeyTemplatesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listKeyTemplatesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listKeyTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListKeystoresOptions successfully`, func() {
				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := ukoService.NewListKeystoresOptions()
				listKeystoresOptionsModel.SetAccept("application/json")
				listKeystoresOptionsModel.SetType([]string{"ibm_cloud_kms"})
				listKeystoresOptionsModel.SetName("Main IBM Cloud")
				listKeystoresOptionsModel.SetDescription("My Example Keystore Description")
				listKeystoresOptionsModel.SetGroup("testString")
				listKeystoresOptionsModel.SetGroups("testString")
				listKeystoresOptionsModel.SetVaultID([]string{"123e4567-e89b-12d3-a456-426614174000"})
				listKeystoresOptionsModel.SetLocation([]string{"testString"})
				listKeystoresOptionsModel.SetLimit(int64(10))
				listKeystoresOptionsModel.SetOffset(int64(0))
				listKeystoresOptionsModel.SetSort([]string{"-updated_at"})
				listKeystoresOptionsModel.SetStatusHealthStatus([]string{"ok"})
				listKeystoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listKeystoresOptionsModel).ToNot(BeNil())
				Expect(listKeystoresOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(listKeystoresOptionsModel.Type).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listKeystoresOptionsModel.Name).To(Equal(core.StringPtr("Main IBM Cloud")))
				Expect(listKeystoresOptionsModel.Description).To(Equal(core.StringPtr("My Example Keystore Description")))
				Expect(listKeystoresOptionsModel.Group).To(Equal(core.StringPtr("testString")))
				Expect(listKeystoresOptionsModel.Groups).To(Equal(core.StringPtr("testString")))
				Expect(listKeystoresOptionsModel.VaultID).To(Equal([]string{"123e4567-e89b-12d3-a456-426614174000"}))
				Expect(listKeystoresOptionsModel.Location).To(Equal([]string{"testString"}))
				Expect(listKeystoresOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listKeystoresOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listKeystoresOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listKeystoresOptionsModel.StatusHealthStatus).To(Equal([]string{"ok"}))
				Expect(listKeystoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListManagedKeyVersionsOptions successfully`, func() {
				// Construct an instance of the ListManagedKeyVersionsOptions model
				id := "testString"
				listManagedKeyVersionsOptionsModel := ukoService.NewListManagedKeyVersionsOptions(id)
				listManagedKeyVersionsOptionsModel.SetID("testString")
				listManagedKeyVersionsOptionsModel.SetAlgorithm([]string{"aes"})
				listManagedKeyVersionsOptionsModel.SetState([]string{"active"})
				listManagedKeyVersionsOptionsModel.SetLimit(int64(10))
				listManagedKeyVersionsOptionsModel.SetOffset(int64(0))
				listManagedKeyVersionsOptionsModel.SetSort([]string{"-updated_at"})
				listManagedKeyVersionsOptionsModel.SetLabel("testString")
				listManagedKeyVersionsOptionsModel.SetActivationDate("testString")
				listManagedKeyVersionsOptionsModel.SetActivationDateMin("testString")
				listManagedKeyVersionsOptionsModel.SetActivationDateMax("testString")
				listManagedKeyVersionsOptionsModel.SetDeactivationDate("testString")
				listManagedKeyVersionsOptionsModel.SetDeactivationDateMin("testString")
				listManagedKeyVersionsOptionsModel.SetDeactivationDateMax("testString")
				listManagedKeyVersionsOptionsModel.SetExpirationDate("testString")
				listManagedKeyVersionsOptionsModel.SetExpirationDateMin("testString")
				listManagedKeyVersionsOptionsModel.SetExpirationDateMax("testString")
				listManagedKeyVersionsOptionsModel.SetCreatedAt("testString")
				listManagedKeyVersionsOptionsModel.SetCreatedAtMin("testString")
				listManagedKeyVersionsOptionsModel.SetCreatedAtMax("testString")
				listManagedKeyVersionsOptionsModel.SetUpdatedAt("testString")
				listManagedKeyVersionsOptionsModel.SetUpdatedAtMin("testString")
				listManagedKeyVersionsOptionsModel.SetUpdatedAtMax("testString")
				listManagedKeyVersionsOptionsModel.SetRotatedAtMin("testString")
				listManagedKeyVersionsOptionsModel.SetRotatedAtMax("testString")
				listManagedKeyVersionsOptionsModel.SetSize("testString")
				listManagedKeyVersionsOptionsModel.SetSizeMin("testString")
				listManagedKeyVersionsOptionsModel.SetSizeMax("testString")
				listManagedKeyVersionsOptionsModel.SetReferencedKeystoresType([]string{"ibm_cloud_kms"})
				listManagedKeyVersionsOptionsModel.SetReferencedKeystoresName([]string{"testString"})
				listManagedKeyVersionsOptionsModel.SetInstancesKeystoreType([]string{"ibm_cloud_kms"})
				listManagedKeyVersionsOptionsModel.SetTemplateName("AWS-TEMPLATE")
				listManagedKeyVersionsOptionsModel.SetTemplateID([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"})
				listManagedKeyVersionsOptionsModel.SetTemplateType([]string{"user_defined"})
				listManagedKeyVersionsOptionsModel.SetStatusInKeystoresKeystoreSyncFlag([]string{"out_of_sync"})
				listManagedKeyVersionsOptionsModel.SetTemplateAlignmentStatus("aligned")
				listManagedKeyVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listManagedKeyVersionsOptionsModel).ToNot(BeNil())
				Expect(listManagedKeyVersionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.Algorithm).To(Equal([]string{"aes"}))
				Expect(listManagedKeyVersionsOptionsModel.State).To(Equal([]string{"active"}))
				Expect(listManagedKeyVersionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listManagedKeyVersionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listManagedKeyVersionsOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listManagedKeyVersionsOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ActivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ActivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ActivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.DeactivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.DeactivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.DeactivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ExpirationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ExpirationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ExpirationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.RotatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.RotatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.Size).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.SizeMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.SizeMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeyVersionsOptionsModel.ReferencedKeystoresType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeyVersionsOptionsModel.ReferencedKeystoresName).To(Equal([]string{"testString"}))
				Expect(listManagedKeyVersionsOptionsModel.InstancesKeystoreType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeyVersionsOptionsModel.TemplateName).To(Equal(core.StringPtr("AWS-TEMPLATE")))
				Expect(listManagedKeyVersionsOptionsModel.TemplateID).To(Equal([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}))
				Expect(listManagedKeyVersionsOptionsModel.TemplateType).To(Equal([]string{"user_defined"}))
				Expect(listManagedKeyVersionsOptionsModel.StatusInKeystoresKeystoreSyncFlag).To(Equal([]string{"out_of_sync"}))
				Expect(listManagedKeyVersionsOptionsModel.TemplateAlignmentStatus).To(Equal(core.StringPtr("aligned")))
				Expect(listManagedKeyVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListManagedKeysFromKeystoreOptions successfully`, func() {
				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				id := "testString"
				listManagedKeysFromKeystoreOptionsModel := ukoService.NewListManagedKeysFromKeystoreOptions(id)
				listManagedKeysFromKeystoreOptionsModel.SetID("testString")
				listManagedKeysFromKeystoreOptionsModel.SetAccept("application/json")
				listManagedKeysFromKeystoreOptionsModel.SetAlgorithm([]string{"aes"})
				listManagedKeysFromKeystoreOptionsModel.SetState([]string{"pre_activation", "active"})
				listManagedKeysFromKeystoreOptionsModel.SetLimit(int64(10))
				listManagedKeysFromKeystoreOptionsModel.SetOffset(int64(0))
				listManagedKeysFromKeystoreOptionsModel.SetSort([]string{"-updated_at"})
				listManagedKeysFromKeystoreOptionsModel.SetLabel("testString")
				listManagedKeysFromKeystoreOptionsModel.SetActivationDate("testString")
				listManagedKeysFromKeystoreOptionsModel.SetActivationDateMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetActivationDateMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetDeactivationDate("testString")
				listManagedKeysFromKeystoreOptionsModel.SetDeactivationDateMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetDeactivationDateMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetExpirationDate("testString")
				listManagedKeysFromKeystoreOptionsModel.SetExpirationDateMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetExpirationDateMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAt("testString")
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAtMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAtMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAt("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAtMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAtMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetRotatedAtMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetRotatedAtMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetSize("testString")
				listManagedKeysFromKeystoreOptionsModel.SetSizeMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetSizeMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetTemplateName("AWS-TEMPLATE")
				listManagedKeysFromKeystoreOptionsModel.SetTemplateID([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"})
				listManagedKeysFromKeystoreOptionsModel.SetTemplateType([]string{"user_defined"})
				listManagedKeysFromKeystoreOptionsModel.SetStatusInKeystoresKeystoreSyncFlag([]string{"out_of_sync"})
				listManagedKeysFromKeystoreOptionsModel.SetTemplateAlignmentStatus("aligned")
				listManagedKeysFromKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listManagedKeysFromKeystoreOptionsModel).ToNot(BeNil())
				Expect(listManagedKeysFromKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(listManagedKeysFromKeystoreOptionsModel.Algorithm).To(Equal([]string{"aes"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.State).To(Equal([]string{"pre_activation", "active"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listManagedKeysFromKeystoreOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listManagedKeysFromKeystoreOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ActivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ActivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ActivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.DeactivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.DeactivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.DeactivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ExpirationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ExpirationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ExpirationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.RotatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.RotatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.Size).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.SizeMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.SizeMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.TemplateName).To(Equal(core.StringPtr("AWS-TEMPLATE")))
				Expect(listManagedKeysFromKeystoreOptionsModel.TemplateID).To(Equal([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.TemplateType).To(Equal([]string{"user_defined"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.StatusInKeystoresKeystoreSyncFlag).To(Equal([]string{"out_of_sync"}))
				Expect(listManagedKeysFromKeystoreOptionsModel.TemplateAlignmentStatus).To(Equal(core.StringPtr("aligned")))
				Expect(listManagedKeysFromKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListManagedKeysOptions successfully`, func() {
				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := ukoService.NewListManagedKeysOptions()
				listManagedKeysOptionsModel.SetAccept("application/json")
				listManagedKeysOptionsModel.SetVaultID([]string{"123e4567-e89b-12d3-a456-426614174000"})
				listManagedKeysOptionsModel.SetAlgorithm([]string{"aes"})
				listManagedKeysOptionsModel.SetState([]string{"pre_activation", "active"})
				listManagedKeysOptionsModel.SetLimit(int64(10))
				listManagedKeysOptionsModel.SetOffset(int64(0))
				listManagedKeysOptionsModel.SetSort([]string{"-updated_at"})
				listManagedKeysOptionsModel.SetLabel("testString")
				listManagedKeysOptionsModel.SetActivationDate("testString")
				listManagedKeysOptionsModel.SetActivationDateMin("testString")
				listManagedKeysOptionsModel.SetActivationDateMax("testString")
				listManagedKeysOptionsModel.SetDeactivationDate("testString")
				listManagedKeysOptionsModel.SetDeactivationDateMin("testString")
				listManagedKeysOptionsModel.SetDeactivationDateMax("testString")
				listManagedKeysOptionsModel.SetExpirationDate("testString")
				listManagedKeysOptionsModel.SetExpirationDateMin("testString")
				listManagedKeysOptionsModel.SetExpirationDateMax("testString")
				listManagedKeysOptionsModel.SetCreatedAt("testString")
				listManagedKeysOptionsModel.SetCreatedAtMin("testString")
				listManagedKeysOptionsModel.SetCreatedAtMax("testString")
				listManagedKeysOptionsModel.SetUpdatedAt("testString")
				listManagedKeysOptionsModel.SetUpdatedAtMin("testString")
				listManagedKeysOptionsModel.SetUpdatedAtMax("testString")
				listManagedKeysOptionsModel.SetRotatedAtMin("testString")
				listManagedKeysOptionsModel.SetRotatedAtMax("testString")
				listManagedKeysOptionsModel.SetSize("testString")
				listManagedKeysOptionsModel.SetSizeMin("testString")
				listManagedKeysOptionsModel.SetSizeMax("testString")
				listManagedKeysOptionsModel.SetReferencedKeystoresType([]string{"ibm_cloud_kms"})
				listManagedKeysOptionsModel.SetReferencedKeystoresName([]string{"testString"})
				listManagedKeysOptionsModel.SetInstancesKeystoreType([]string{"ibm_cloud_kms"})
				listManagedKeysOptionsModel.SetTemplateName("AWS-TEMPLATE")
				listManagedKeysOptionsModel.SetTemplateID([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"})
				listManagedKeysOptionsModel.SetTemplateType([]string{"user_defined"})
				listManagedKeysOptionsModel.SetStatusInKeystoresKeystoreSyncFlag([]string{"out_of_sync"})
				listManagedKeysOptionsModel.SetTemplateAlignmentStatus("aligned")
				listManagedKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listManagedKeysOptionsModel).ToNot(BeNil())
				Expect(listManagedKeysOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(listManagedKeysOptionsModel.VaultID).To(Equal([]string{"123e4567-e89b-12d3-a456-426614174000"}))
				Expect(listManagedKeysOptionsModel.Algorithm).To(Equal([]string{"aes"}))
				Expect(listManagedKeysOptionsModel.State).To(Equal([]string{"pre_activation", "active"}))
				Expect(listManagedKeysOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listManagedKeysOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listManagedKeysOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listManagedKeysOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ActivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ActivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ActivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.DeactivationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.DeactivationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.DeactivationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ExpirationDate).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ExpirationDateMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ExpirationDateMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.RotatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.RotatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.Size).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.SizeMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.SizeMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.ReferencedKeystoresType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeysOptionsModel.ReferencedKeystoresName).To(Equal([]string{"testString"}))
				Expect(listManagedKeysOptionsModel.InstancesKeystoreType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeysOptionsModel.TemplateName).To(Equal(core.StringPtr("AWS-TEMPLATE")))
				Expect(listManagedKeysOptionsModel.TemplateID).To(Equal([]string{"5295ad47-2ce9-43c3-b9e7-e5a9482c362b"}))
				Expect(listManagedKeysOptionsModel.TemplateType).To(Equal([]string{"user_defined"}))
				Expect(listManagedKeysOptionsModel.StatusInKeystoresKeystoreSyncFlag).To(Equal([]string{"out_of_sync"}))
				Expect(listManagedKeysOptionsModel.TemplateAlignmentStatus).To(Equal(core.StringPtr("aligned")))
				Expect(listManagedKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVaultsOptions successfully`, func() {
				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := ukoService.NewListVaultsOptions()
				listVaultsOptionsModel.SetAccept("application/json")
				listVaultsOptionsModel.SetLimit(int64(10))
				listVaultsOptionsModel.SetOffset(int64(0))
				listVaultsOptionsModel.SetSort([]string{"-updated_at"})
				listVaultsOptionsModel.SetName("My Example Vault")
				listVaultsOptionsModel.SetDescription("My Example Vault Description")
				listVaultsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVaultsOptionsModel).ToNot(BeNil())
				Expect(listVaultsOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(listVaultsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listVaultsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listVaultsOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listVaultsOptionsModel.Name).To(Equal(core.StringPtr("My Example Vault")))
				Expect(listVaultsOptionsModel.Description).To(Equal(core.StringPtr("My Example Vault Description")))
				Expect(listVaultsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRotateManagedKeyOptions successfully`, func() {
				// Construct an instance of the RotateManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				rotateManagedKeyOptionsModel := ukoService.NewRotateManagedKeyOptions(id, ifMatch)
				rotateManagedKeyOptionsModel.SetID("testString")
				rotateManagedKeyOptionsModel.SetIfMatch("testString")
				rotateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(rotateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(rotateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(rotateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(rotateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSyncManagedKeyOptions successfully`, func() {
				// Construct an instance of the SyncManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				syncManagedKeyOptionsModel := ukoService.NewSyncManagedKeyOptions(id, ifMatch)
				syncManagedKeyOptionsModel.SetID("testString")
				syncManagedKeyOptionsModel.SetIfMatch("testString")
				syncManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(syncManagedKeyOptionsModel).ToNot(BeNil())
				Expect(syncManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(syncManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(syncManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnarchiveKeyTemplateOptions successfully`, func() {
				// Construct an instance of the UnarchiveKeyTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				unarchiveKeyTemplateOptionsModel := ukoService.NewUnarchiveKeyTemplateOptions(id, ifMatch)
				unarchiveKeyTemplateOptionsModel.SetID("testString")
				unarchiveKeyTemplateOptionsModel.SetIfMatch("testString")
				unarchiveKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unarchiveKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(unarchiveKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unarchiveKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(unarchiveKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateKeyTemplateOptions successfully`, func() {
				// Construct an instance of the KeystoresPropertiesUpdateGoogleKms model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdateGoogleKms)
				Expect(keystoresPropertiesUpdateModel).ToNot(BeNil())
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel = core.StringPtr("software")
				keystoresPropertiesUpdateModel.GoogleKeyPurpose = core.StringPtr("encrypt_decrypt")
				keystoresPropertiesUpdateModel.GoogleKmsAlgorithm = core.StringPtr("google_symmetric_encryption")
				Expect(keystoresPropertiesUpdateModel.Group).To(Equal(core.StringPtr("Production")))
				Expect(keystoresPropertiesUpdateModel.GoogleKeyProtectionLevel).To(Equal(core.StringPtr("software")))
				Expect(keystoresPropertiesUpdateModel.GoogleKeyPurpose).To(Equal(core.StringPtr("encrypt_decrypt")))
				Expect(keystoresPropertiesUpdateModel.GoogleKmsAlgorithm).To(Equal(core.StringPtr("google_symmetric_encryption")))

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				Expect(keyPropertiesUpdateModel).ToNot(BeNil())
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				keyPropertiesUpdateModel.DeactivateOnRotation = core.BoolPtr(true)
				Expect(keyPropertiesUpdateModel.Size).To(Equal(core.StringPtr("256")))
				Expect(keyPropertiesUpdateModel.ActivationDate).To(Equal(core.StringPtr("P5Y1M1W2D")))
				Expect(keyPropertiesUpdateModel.ExpirationDate).To(Equal(core.StringPtr("P1Y2M1W4D")))
				Expect(keyPropertiesUpdateModel.State).To(Equal(core.StringPtr("active")))
				Expect(keyPropertiesUpdateModel.DeactivateOnRotation).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateKeyTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				updateKeyTemplateOptionsModel := ukoService.NewUpdateKeyTemplateOptions(id, ifMatch)
				updateKeyTemplateOptionsModel.SetID("testString")
				updateKeyTemplateOptionsModel.SetIfMatch("testString")
				updateKeyTemplateOptionsModel.SetName("EXAMPLE-TEMPLATE")
				updateKeyTemplateOptionsModel.SetKeystores([]ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel})
				updateKeyTemplateOptionsModel.SetDescription("testString")
				updateKeyTemplateOptionsModel.SetKey(keyPropertiesUpdateModel)
				updateKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(updateKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.Name).To(Equal(core.StringPtr("EXAMPLE-TEMPLATE")))
				Expect(updateKeyTemplateOptionsModel.Keystores).To(Equal([]ukov4.KeystoresPropertiesUpdateIntf{keystoresPropertiesUpdateModel}))
				Expect(updateKeyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.Key).To(Equal(keyPropertiesUpdateModel))
				Expect(updateKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateKeystoreOptions successfully`, func() {
				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeGoogleKmsUpdate)
				Expect(keystoreUpdateRequestModel).ToNot(BeNil())
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.TlsProxy = core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.GoogleCredentials = core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")
				keystoreUpdateRequestModel.GoogleLocation = core.StringPtr("europe-central2")
				keystoreUpdateRequestModel.GoogleKeyRing = core.StringPtr("my-key-ring")
				Expect(keystoreUpdateRequestModel.Name).To(Equal(core.StringPtr("IBM Cloud Keystore Name")))
				Expect(keystoreUpdateRequestModel.Description).To(Equal(core.StringPtr("Azure keystore")))
				Expect(keystoreUpdateRequestModel.TlsProxy).To(Equal(core.StringPtr("c-04.private.us-east.link.satellite.cloud.ibm.com:12358")))
				Expect(keystoreUpdateRequestModel.Groups).To(Equal([]string{"Production"}))
				Expect(keystoreUpdateRequestModel.GoogleCredentials).To(Equal(core.StringPtr("eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo=")))
				Expect(keystoreUpdateRequestModel.GoogleLocation).To(Equal(core.StringPtr("europe-central2")))
				Expect(keystoreUpdateRequestModel.GoogleKeyRing).To(Equal(core.StringPtr("my-key-ring")))

				// Construct an instance of the UpdateKeystoreOptions model
				id := "testString"
				ifMatch := "testString"
				var keystoreBody ukov4.KeystoreUpdateRequestIntf = nil
				updateKeystoreOptionsModel := ukoService.NewUpdateKeystoreOptions(id, ifMatch, keystoreBody)
				updateKeystoreOptionsModel.SetID("testString")
				updateKeystoreOptionsModel.SetIfMatch("testString")
				updateKeystoreOptionsModel.SetKeystoreBody(keystoreUpdateRequestModel)
				updateKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateKeystoreOptionsModel).ToNot(BeNil())
				Expect(updateKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateKeystoreOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateKeystoreOptionsModel.KeystoreBody).To(Equal(keystoreUpdateRequestModel))
				Expect(updateKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateManagedKeyFromTemplateOptions successfully`, func() {
				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				id := "testString"
				ifMatch := "testString"
				updateManagedKeyFromTemplateOptionsModel := ukoService.NewUpdateManagedKeyFromTemplateOptions(id, ifMatch)
				updateManagedKeyFromTemplateOptionsModel.SetID("testString")
				updateManagedKeyFromTemplateOptionsModel.SetIfMatch("testString")
				updateManagedKeyFromTemplateOptionsModel.SetDryRun(false)
				updateManagedKeyFromTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateManagedKeyFromTemplateOptionsModel).ToNot(BeNil())
				Expect(updateManagedKeyFromTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyFromTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyFromTemplateOptionsModel.DryRun).To(Equal(core.BoolPtr(false)))
				Expect(updateManagedKeyFromTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateManagedKeyOptions successfully`, func() {
				// Construct an instance of the UpdateManagedKeyOptions model
				id := "testString"
				ifMatch := "testString"
				updateManagedKeyOptionsModel := ukoService.NewUpdateManagedKeyOptions(id, ifMatch)
				updateManagedKeyOptionsModel.SetID("testString")
				updateManagedKeyOptionsModel.SetIfMatch("testString")
				updateManagedKeyOptionsModel.SetLabel("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.SetActivationDate(CreateMockDate("2020-12-11"))
				updateManagedKeyOptionsModel.SetExpirationDate(CreateMockDate("2030-11-12"))
				updateManagedKeyOptionsModel.SetDescription("testString")
				updateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(updateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.Label).To(Equal(core.StringPtr("IBM CLOUD KEY")))
				Expect(updateManagedKeyOptionsModel.ActivationDate).To(Equal(CreateMockDate("2020-12-11")))
				Expect(updateManagedKeyOptionsModel.ExpirationDate).To(Equal(CreateMockDate("2030-11-12")))
				Expect(updateManagedKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVaultOptions successfully`, func() {
				// Construct an instance of the UpdateVaultOptions model
				id := "testString"
				ifMatch := "testString"
				updateVaultOptionsModel := ukoService.NewUpdateVaultOptions(id, ifMatch)
				updateVaultOptionsModel.SetID("testString")
				updateVaultOptionsModel.SetIfMatch("testString")
				updateVaultOptionsModel.SetName("Jakub's Vault")
				updateVaultOptionsModel.SetDescription("Updated description of the vault")
				updateVaultOptionsModel.SetRecoveryKeyLabel("TEKMF.AES.RECOVERY.00001")
				updateVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVaultOptionsModel).ToNot(BeNil())
				Expect(updateVaultOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVaultOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateVaultOptionsModel.Name).To(Equal(core.StringPtr("Jakub's Vault")))
				Expect(updateVaultOptionsModel.Description).To(Equal(core.StringPtr("Updated description of the vault")))
				Expect(updateVaultOptionsModel.RecoveryKeyLabel).To(Equal(core.StringPtr("TEKMF.AES.RECOVERY.00001")))
				Expect(updateVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewKeyProperties successfully`, func() {
				size := "256"
				algorithm := "aes"
				activationDate := "P5Y1M1W2D"
				expirationDate := "P1Y2M1W4D"
				state := "active"
				_model, err := ukoService.NewKeyProperties(size, algorithm, activationDate, expirationDate, state)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTag successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := ukoService.NewTag(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVaultReferenceInCreationRequest successfully`, func() {
				id := "5295ad47-2ce9-43c3-b9e7-e5a9482c362b"
				_model, err := ukoService.NewVaultReferenceInCreationRequest(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewKeystoreCreationRequestKeystoreTypeAwsKmsCreate successfully`, func() {
				typeVar := "aws_kms"
				var vault *ukov4.VaultReferenceInCreationRequest = nil
				name := "IBM Cloud Keystore Name"
				awsRegion := "af_south_1"
				awsAccessKeyID := "BSDFWERUANLKJDN54AAS"
				awsSecretAccessKey := "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeAwsKmsCreate(typeVar, vault, name, awsRegion, awsAccessKeyID, awsSecretAccessKey)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewKeystoreCreationRequestKeystoreTypeAzureCreate successfully`, func() {
				typeVar := "azure_key_vault"
				var vault *ukov4.VaultReferenceInCreationRequest = nil
				azureServiceName := "azure-service-in-ibm"
				azureResourceGroup := "Azure-test"
				azureServicePrincipalClientID := "91018db5-c756-468e-bd4e-69c99fc1a749"
				azureServicePrincipalPassword := "9wN1YP5XwrrHIdvIYv7imHiC83Q_lSWAWa"
				azureTenant := "b8e1a93c-2449-462f-8fa0-1d00595ea859"
				azureSubscriptionID := "a98667h9b-5fhf-42f3-9392-26856b045g08"
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeAzureCreate(typeVar, vault, azureServiceName, azureResourceGroup, azureServicePrincipalClientID, azureServicePrincipalPassword, azureTenant, azureSubscriptionID)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewKeystoreCreationRequestKeystoreTypeGoogleKmsCreate successfully`, func() {
				typeVar := "google_kms"
				var vault *ukov4.VaultReferenceInCreationRequest = nil
				name := "IBM Cloud Keystore Name"
				googleCredentials := "eyJleGFtcGxlIjogImdvb2dsZV9jbG91ZF9rbXMifQo="
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeGoogleKmsCreate(typeVar, vault, name, googleCredentials)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate successfully`, func() {
				typeVar := "ibm_cloud_kms"
				var vault *ukov4.VaultReferenceInCreationRequest = nil
				name := "IBM Cloud Keystore Name"
				ibmApiEndpoint := "https://us-south.kms.cloud.ibm.com"
				ibmIamEndpoint := "https://iam.cloud.ibm.com/identity/token"
				ibmApiKey := "9NoWuteprHTtC_-YDkJnbBoHn3qj0gy-rWHUqegOh1DA"
				ibmInstanceID := "d139ea58-a073-441b-ba4e-dcc8bae58be4"
				ibmVariant := "hpcs"
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsCreate(typeVar, vault, name, ibmApiEndpoint, ibmIamEndpoint, ibmApiKey, ibmInstanceID, ibmVariant)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate successfully`, func() {
				typeVar := "ibm_cloud_kms"
				var vault *ukov4.VaultReferenceInCreationRequest = nil
				ibmVariant := "hpcs"
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalExternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalCreateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeIbmCloudKmsInternalUpdateKeystoreTypeBaseUpdate(typeVar, vault, ibmVariant)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
