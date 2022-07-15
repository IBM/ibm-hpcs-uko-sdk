/**
 * (C) Copyright IBM Corp. 2022.
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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/cloud-go-sdk/ukov4"
	"github.com/IBM/go-sdk-core/v5/core"
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
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
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
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMax = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMax = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMax = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
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
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMax = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
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
				listManagedKeysOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.SizeMax = core.Int64Ptr(int64(38))
				listManagedKeysOptionsModel.ReferencedKeystoresType = []string{"ibm_cloud_kms"}
				listManagedKeysOptionsModel.ReferencedKeystoresName = []string{"testString"}
				listManagedKeysOptionsModel.InstancesKeystoreType = []string{"ibm_cloud_kms"}
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
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","tags":[{"name":"Name","value":"Value"}],"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"}}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","tags":[{"name":"Name","value":"Value"}],"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"}}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"total_count":2,"limit":1}`)
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
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Size: core.Int64Ptr(int64(38)),
					SizeMin: core.Int64Ptr(int64(38)),
					SizeMax: core.Int64Ptr(int64(38)),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
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
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Size: core.Int64Ptr(int64(38)),
					SizeMin: core.Int64Ptr(int64(38)),
					SizeMax: core.Int64Ptr(int64(38)),
					ReferencedKeystoresType: []string{"ibm_cloud_kms"},
					ReferencedKeystoresName: []string{"testString"},
					InstancesKeystoreType: []string{"ibm_cloud_kms"},
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateManagedKeyOptions model
				createManagedKeyOptionsModel := new(ukov4.CreateManagedKeyOptions)
				createManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				createManagedKeyOptionsModel.TemplateName = core.StringPtr("testString")
				createManagedKeyOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				createManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				deleteManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				deleteManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				getManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				getManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				getManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				getManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateManagedKeyOptions model
				updateManagedKeyOptionsModel := new(ukov4.UpdateManagedKeyOptions)
				updateManagedKeyOptionsModel.ID = core.StringPtr("testString")
				updateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyOptionsModel.IfMatch = core.StringPtr("testString")
				updateManagedKeyOptionsModel.Label = core.StringPtr("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.ActivationDate = CreateMockDate("2020-12-11")
				updateManagedKeyOptionsModel.ExpirationDate = CreateMockDate("2030-11-12")
				updateManagedKeyOptionsModel.Tags = []ukov4.Tag{*tagModel}
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
	Describe(`GetKeyDistributionStatusForKeystores(getKeyDistributionStatusForKeystoresOptions *GetKeyDistributionStatusForKeystoresOptions) - Operation response error`, func() {
		getKeyDistributionStatusForKeystoresPath := "/v4/managed_keys/testString/status_in_keystores"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeyDistributionStatusForKeystoresPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeyDistributionStatusForKeystoresOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
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
				getKeyDistributionStatusForKeystoresOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_in_keystores": [{"keystore": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "status": "active", "error": {"status_code": 400, "trace": "9daee671-916a-4678-850b-10b9110236d", "errors": [{"code": "missing_field", "message": "The algorithm field is required", "more_info": "https://cloud.ibm.com/apidocs/uko#create-managed-key", "message_params": ["My Key Template"], "target": {"type": "field", "name": "first_name"}}]}, "key_id_in_keystore": "123e4567-e89b-12d3-a456-426614174000"}]}`)
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
				getKeyDistributionStatusForKeystoresOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeyDistributionStatusForKeystoresOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeyDistributionStatusForKeystoresOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				updateManagedKeyFromTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				updateManagedKeyFromTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				updateManagedKeyFromTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
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
				updateManagedKeyFromTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
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
				updateManagedKeyFromTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateManagedKeyFromTemplateOptionsModel.IfMatch = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				activateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				activateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				activateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				activateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				activateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				deactivateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				deactivateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				deactivateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				deactivateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				deactivateManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				destroyManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				destroyManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				destroyManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				destroyManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
				destroyManagedKeyOptionsModel.UKOVault = core.StringPtr("testString")
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
	Describe(`ListKeyTemplates(listKeyTemplatesOptions *ListKeyTemplatesOptions) - Operation response error`, func() {
		listKeyTemplatesPath := "/v4/templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listKeyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["key.algorithm"]).To(Equal([]string{"aes"}))
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
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = core.StringPtr("aes")
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

					Expect(req.URL.Query()["key.algorithm"]).To(Equal([]string{"aes"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "templates": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = core.StringPtr("aes")
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

					Expect(req.URL.Query()["key.algorithm"]).To(Equal([]string{"aes"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "templates": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = core.StringPtr("aes")
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
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = core.StringPtr("aes")
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
				listKeyTemplatesOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeyTemplatesOptionsModel.KeyAlgorithm = core.StringPtr("aes")
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
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"templates":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","version":1,"name":"AWS-TEMPLATE","key":{"size":"256","algorithm":"aes","activation_date":"P5Y1M1W2D","expiration_date":"P1Y2M1W4D","state":"active"},"description":"The description of the template","created_at":"2022-02-05T23:00:14.000Z","updated_at":"2022-02-05T23:00:14.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","keystores":[{"group":"Production","type":"ibm_cloud_kms"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"templates":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","version":1,"name":"AWS-TEMPLATE","key":{"size":"256","algorithm":"aes","activation_date":"P5Y1M1W2D","expiration_date":"P1Y2M1W4D","state":"active"},"description":"The description of the template","created_at":"2022-02-05T23:00:14.000Z","updated_at":"2022-02-05T23:00:14.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","keystores":[{"group":"Production","type":"ibm_cloud_kms"}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"limit":1}`)
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
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					KeyAlgorithm: core.StringPtr("aes"),
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
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					KeyAlgorithm: core.StringPtr("aes"),
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresProperties{*keystoresPropertiesModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresProperties{*keystoresPropertiesModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresProperties{*keystoresPropertiesModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
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

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresProperties{*keystoresPropertiesModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
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

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")

				// Construct an instance of the CreateKeyTemplateOptions model
				createKeyTemplateOptionsModel := new(ukov4.CreateKeyTemplateOptions)
				createKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				createKeyTemplateOptionsModel.Vault = vaultReferenceInCreationRequestModel
				createKeyTemplateOptionsModel.Name = core.StringPtr("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.Key = keyPropertiesModel
				createKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresProperties{*keystoresPropertiesModel}
				createKeyTemplateOptionsModel.Description = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				deleteKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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
				deleteKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				getKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
				getKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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

				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "version": 1, "name": "AWS-TEMPLATE", "key": {"size": "256", "algorithm": "aes", "activation_date": "P5Y1M1W2D", "expiration_date": "P1Y2M1W4D", "state": "active"}, "description": "The description of the template", "created_at": "2022-02-05T23:00:14.000Z", "updated_at": "2022-02-05T23:00:14.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "keystores": [{"group": "Production", "type": "ibm_cloud_kms"}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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

				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}
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

				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}
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

				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")

				// Construct an instance of the UpdateKeyTemplateOptions model
				updateKeyTemplateOptionsModel := new(ukov4.UpdateKeyTemplateOptions)
				updateKeyTemplateOptionsModel.ID = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.UKOVault = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				updateKeyTemplateOptionsModel.Keystores = []ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}
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
					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
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
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = core.StringPtr("testString")
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
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

					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "keystores": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}]}`)
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
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = core.StringPtr("testString")
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
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

					Expect(req.URL.Query()["name"]).To(Equal([]string{"Main IBM Cloud"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Keystore Description"}))
					Expect(req.URL.Query()["groups[]"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "keystores": [{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}]}`)
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
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = core.StringPtr("testString")
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
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
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = core.StringPtr("testString")
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
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
				listKeystoresOptionsModel.Type = []string{"ibm_cloud_kms"}
				listKeystoresOptionsModel.Name = core.StringPtr("Main IBM Cloud")
				listKeystoresOptionsModel.Description = core.StringPtr("My Example Keystore Description")
				listKeystoresOptionsModel.Groups = core.StringPtr("testString")
				listKeystoresOptionsModel.VaultID = []string{"123e4567-e89b-12d3-a456-426614174000"}
				listKeystoresOptionsModel.Location = core.StringPtr("testString")
				listKeystoresOptionsModel.Limit = core.Int64Ptr(int64(10))
				listKeystoresOptionsModel.Offset = core.Int64Ptr(int64(0))
				listKeystoresOptionsModel.Sort = []string{"-updated_at"}
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
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"keystores":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Accounting","location":"us-south","description":"IBM Cloud keystore for testing","groups":["Production"],"type":"ibm_cloud_kms","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","aws_region":"af_south_1","aws_access_key_id":"BSDFWERUANLKJDN54AAS","aws_secret_access_key":"6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"keystores":[{"vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Accounting","location":"us-south","description":"IBM Cloud keystore for testing","groups":["Production"],"type":"ibm_cloud_kms","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46","aws_region":"af_south_1","aws_access_key_id":"BSDFWERUANLKJDN54AAS","aws_secret_access_key":"6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}],"limit":1}`)
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
					Type: []string{"ibm_cloud_kms"},
					Name: core.StringPtr("Main IBM Cloud"),
					Description: core.StringPtr("My Example Keystore Description"),
					Groups: core.StringPtr("testString"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Location: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
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
					Type: []string{"ibm_cloud_kms"},
					Name: core.StringPtr("Main IBM Cloud"),
					Description: core.StringPtr("My Example Keystore Description"),
					Groups: core.StringPtr("testString"),
					VaultID: []string{"123e4567-e89b-12d3-a456-426614174000"},
					Location: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: []string{"-updated_at"},
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for dry_run query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for dry_run query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the CreateKeystoreOptions model
				createKeystoreOptionsModel := new(ukov4.CreateKeystoreOptions)
				createKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["mode"]).To(Equal([]string{"restrict"}))
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
				deleteKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				deleteKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				deleteKeystoreOptionsModel.Mode = core.StringPtr("restrict")
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
				deleteKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				deleteKeystoreOptionsModel.IfMatch = core.StringPtr("testString")
				deleteKeystoreOptionsModel.Mode = core.StringPtr("restrict")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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
				getKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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
				getKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Accounting", "location": "us-south", "description": "IBM Cloud keystore for testing", "groups": ["Production"], "type": "ibm_cloud_kms", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46", "aws_region": "af_south_1", "aws_access_key_id": "BSDFWERUANLKJDN54AAS", "aws_secret_access_key": "6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs"}`)
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

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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

				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")

				// Construct an instance of the UpdateKeystoreOptions model
				updateKeystoreOptionsModel := new(ukov4.UpdateKeystoreOptions)
				updateKeystoreOptionsModel.ID = core.StringPtr("testString")
				updateKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
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
	Describe(`GetKeystoreStatus(getKeystoreStatusOptions *GetKeystoreStatusOptions) - Operation response error`, func() {
		getKeystoreStatusPath := "/v4/keystores/testString/status"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKeystoreStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeystoreStatusOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeystoreStatusOptionsModel.UKOVault = core.StringPtr("testString")
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
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
				getKeystoreStatusOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeystoreStatusOptionsModel.UKOVault = core.StringPtr("testString")
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
				getKeystoreStatusOptionsModel.UKOVault = core.StringPtr("testString")
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
					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
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
				listManagedKeysFromKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
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
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.Int64Ptr(int64(38))
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listManagedKeysFromKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
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
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.Int64Ptr(int64(38))
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

					Expect(req.Header["Uko-Vault"]).ToNot(BeNil())
					Expect(req.Header["Uko-Vault"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["label"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["activation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["deactivation_date_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["created_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_min"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["updated_at_max"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["size"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_min"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["size_max"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "managed_keys": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "vault": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "Vault-1", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "template": {"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "AWS-KMS-TEMPLATE", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "description": "Managed key description", "label": "IBM CLOUD KEY", "state": "active", "size": "256", "algorithm": "aes", "verification_patterns": [{"method": "enc-zero", "value": "U3dhZ2dlciByb2Nrcw=="}], "activation_date": "2020-12-11", "expiration_date": "2030-11-12", "tags": [{"name": "Name", "value": "Value"}], "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "referenced_keystores": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "ibm-cloud", "type": "ibm_cloud_kms", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}], "instances": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "label_in_keystore": "IBM CLOUD KEY", "type": "private_key", "keystore": {"group": "Group", "type": "ibm_cloud_kms"}}], "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
				listManagedKeysFromKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
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
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.Int64Ptr(int64(38))
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
				listManagedKeysFromKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
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
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.Int64Ptr(int64(38))
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
				listManagedKeysFromKeystoreOptionsModel.UKOVault = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.ID = core.StringPtr("testString")
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
				listManagedKeysFromKeystoreOptionsModel.CreatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.CreatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAt = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax = core.StringPtr("testString")
				listManagedKeysFromKeystoreOptionsModel.Size = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMin = core.Int64Ptr(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SizeMax = core.Int64Ptr(int64(38))
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
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","tags":[{"name":"Name","value":"Value"}],"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"}}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"managed_keys":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","vault":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"Vault-1","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"template":{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"AWS-KMS-TEMPLATE","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"},"description":"Managed key description","label":"IBM CLOUD KEY","state":"active","size":"256","algorithm":"aes","verification_patterns":[{"method":"enc-zero","value":"U3dhZ2dlciByb2Nrcw=="}],"activation_date":"2020-12-11","expiration_date":"2030-11-12","tags":[{"name":"Name","value":"Value"}],"created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","referenced_keystores":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"ibm-cloud","type":"ibm_cloud_kms","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"instances":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","label_in_keystore":"IBM CLOUD KEY","type":"private_key","keystore":{"group":"Group","type":"ibm_cloud_kms"}}],"href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}],"total_count":2,"limit":1}`)
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
					UKOVault: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
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
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Size: core.Int64Ptr(int64(38)),
					SizeMin: core.Int64Ptr(int64(38)),
					SizeMax: core.Int64Ptr(int64(38)),
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
					UKOVault: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
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
					CreatedAt: core.StringPtr("testString"),
					CreatedAtMin: core.StringPtr("testString"),
					CreatedAtMax: core.StringPtr("testString"),
					UpdatedAt: core.StringPtr("testString"),
					UpdatedAtMin: core.StringPtr("testString"),
					UpdatedAtMax: core.StringPtr("testString"),
					Size: core.Int64Ptr(int64(38)),
					SizeMin: core.Int64Ptr(int64(38)),
					SizeMax: core.Int64Ptr(int64(38)),
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Vault"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Vault Description"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "vaults": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"My Example Vault"}))
					Expect(req.URL.Query()["description"]).To(Equal([]string{"My Example Vault Description"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 3456, "limit": 200, "offset": 100, "first": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "last": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "previous": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "next": {"href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}, "vaults": [{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"vaults":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"My Example Vault","description":"The description of the vault","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"vaults":[{"id":"5295ad47-2ce9-43c3-b9e7-e5a9482c362b","name":"My Example Vault","description":"The description of the vault","created_at":"2022-02-22T10:27:08.000Z","updated_at":"2022-02-22T10:27:08.000Z","created_by":"IBMid-1308197YB4","updated_by":"IBMid-1308197YB4","href":"https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}]}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "5295ad47-2ce9-43c3-b9e7-e5a9482c362b", "name": "My Example Vault", "description": "The description of the vault", "created_at": "2022-02-22T10:27:08.000Z", "updated_at": "2022-02-22T10:27:08.000Z", "created_by": "IBMid-1308197YB4", "updated_by": "IBMid-1308197YB4", "href": "https://uko.us-south.hs-crypto.cloud.ibm.com:9549/api/v4/managed_keys/c2d8d0ee-c333-414f-8e64-af47320e5a46"}`)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ukoService, _ := ukov4.NewUkoV4(&ukov4.UkoV4Options{
				URL:           "http://ukov4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewActivateManagedKeyOptions successfully`, func() {
				// Construct an instance of the ActivateManagedKeyOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				activateManagedKeyOptionsModel := ukoService.NewActivateManagedKeyOptions(id, ukoVault, ifMatch)
				activateManagedKeyOptionsModel.SetID("testString")
				activateManagedKeyOptionsModel.SetUKOVault("testString")
				activateManagedKeyOptionsModel.SetIfMatch("testString")
				activateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(activateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(activateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(activateManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(activateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(activateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				Expect(keyPropertiesModel.Size).To(Equal(core.StringPtr("256")))
				Expect(keyPropertiesModel.Algorithm).To(Equal(core.StringPtr("aes")))
				Expect(keyPropertiesModel.ActivationDate).To(Equal(core.StringPtr("P5Y1M1W2D")))
				Expect(keyPropertiesModel.ExpirationDate).To(Equal(core.StringPtr("P1Y2M1W4D")))
				Expect(keyPropertiesModel.State).To(Equal(core.StringPtr("active")))

				// Construct an instance of the KeystoresProperties model
				keystoresPropertiesModel := new(ukov4.KeystoresProperties)
				Expect(keystoresPropertiesModel).ToNot(BeNil())
				keystoresPropertiesModel.Group = core.StringPtr("Production")
				keystoresPropertiesModel.Type = core.StringPtr("ibm_cloud_kms")
				Expect(keystoresPropertiesModel.Group).To(Equal(core.StringPtr("Production")))
				Expect(keystoresPropertiesModel.Type).To(Equal(core.StringPtr("ibm_cloud_kms")))

				// Construct an instance of the CreateKeyTemplateOptions model
				ukoVault := "testString"
				var createKeyTemplateOptionsVault *ukov4.VaultReferenceInCreationRequest = nil
				createKeyTemplateOptionsName := "EXAMPLE-TEMPLATE"
				var createKeyTemplateOptionsKey *ukov4.KeyProperties = nil
				createKeyTemplateOptionsKeystores := []ukov4.KeystoresProperties{}
				createKeyTemplateOptionsModel := ukoService.NewCreateKeyTemplateOptions(ukoVault, createKeyTemplateOptionsVault, createKeyTemplateOptionsName, createKeyTemplateOptionsKey, createKeyTemplateOptionsKeystores)
				createKeyTemplateOptionsModel.SetUKOVault("testString")
				createKeyTemplateOptionsModel.SetVault(vaultReferenceInCreationRequestModel)
				createKeyTemplateOptionsModel.SetName("EXAMPLE-TEMPLATE")
				createKeyTemplateOptionsModel.SetKey(keyPropertiesModel)
				createKeyTemplateOptionsModel.SetKeystores([]ukov4.KeystoresProperties{*keystoresPropertiesModel})
				createKeyTemplateOptionsModel.SetDescription("testString")
				createKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(createKeyTemplateOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(createKeyTemplateOptionsModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(createKeyTemplateOptionsModel.Name).To(Equal(core.StringPtr("EXAMPLE-TEMPLATE")))
				Expect(createKeyTemplateOptionsModel.Key).To(Equal(keyPropertiesModel))
				Expect(createKeyTemplateOptionsModel.Keystores).To(Equal([]ukov4.KeystoresProperties{*keystoresPropertiesModel}))
				Expect(createKeyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
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
				keystoreCreationRequestModel.Groups = []string{"Production"}
				keystoreCreationRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreCreationRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreCreationRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")
				Expect(keystoreCreationRequestModel.Type).To(Equal(core.StringPtr("aws_kms")))
				Expect(keystoreCreationRequestModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(keystoreCreationRequestModel.Name).To(Equal(core.StringPtr("IBM Cloud Keystore Name")))
				Expect(keystoreCreationRequestModel.Description).To(Equal(core.StringPtr("Azure keystore")))
				Expect(keystoreCreationRequestModel.Groups).To(Equal([]string{"Production"}))
				Expect(keystoreCreationRequestModel.AwsRegion).To(Equal(core.StringPtr("af_south_1")))
				Expect(keystoreCreationRequestModel.AwsAccessKeyID).To(Equal(core.StringPtr("BSDFWERUANLKJDN54AAS")))
				Expect(keystoreCreationRequestModel.AwsSecretAccessKey).To(Equal(core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")))

				// Construct an instance of the CreateKeystoreOptions model
				ukoVault := "testString"
				var keystoreBody ukov4.KeystoreCreationRequestIntf = nil
				createKeystoreOptionsModel := ukoService.NewCreateKeystoreOptions(ukoVault, keystoreBody)
				createKeystoreOptionsModel.SetUKOVault("testString")
				createKeystoreOptionsModel.SetKeystoreBody(keystoreCreationRequestModel)
				createKeystoreOptionsModel.SetDryRun(false)
				createKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createKeystoreOptionsModel).ToNot(BeNil())
				Expect(createKeystoreOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
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

				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				Expect(tagModel).ToNot(BeNil())
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")
				Expect(tagModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(tagModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateManagedKeyOptions model
				ukoVault := "testString"
				createManagedKeyOptionsTemplateName := "testString"
				var createManagedKeyOptionsVault *ukov4.VaultReferenceInCreationRequest = nil
				createManagedKeyOptionsLabel := "IBM CLOUD KEY"
				createManagedKeyOptionsModel := ukoService.NewCreateManagedKeyOptions(ukoVault, createManagedKeyOptionsTemplateName, createManagedKeyOptionsVault, createManagedKeyOptionsLabel)
				createManagedKeyOptionsModel.SetUKOVault("testString")
				createManagedKeyOptionsModel.SetTemplateName("testString")
				createManagedKeyOptionsModel.SetVault(vaultReferenceInCreationRequestModel)
				createManagedKeyOptionsModel.SetLabel("IBM CLOUD KEY")
				createManagedKeyOptionsModel.SetTags([]ukov4.Tag{*tagModel})
				createManagedKeyOptionsModel.SetDescription("testString")
				createManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createManagedKeyOptionsModel).ToNot(BeNil())
				Expect(createManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(createManagedKeyOptionsModel.TemplateName).To(Equal(core.StringPtr("testString")))
				Expect(createManagedKeyOptionsModel.Vault).To(Equal(vaultReferenceInCreationRequestModel))
				Expect(createManagedKeyOptionsModel.Label).To(Equal(core.StringPtr("IBM CLOUD KEY")))
				Expect(createManagedKeyOptionsModel.Tags).To(Equal([]ukov4.Tag{*tagModel}))
				Expect(createManagedKeyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVaultOptions successfully`, func() {
				// Construct an instance of the CreateVaultOptions model
				createVaultOptionsName := "Example Vault"
				createVaultOptionsModel := ukoService.NewCreateVaultOptions(createVaultOptionsName)
				createVaultOptionsModel.SetName("Example Vault")
				createVaultOptionsModel.SetDescription("The description of the creating vault")
				createVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVaultOptionsModel).ToNot(BeNil())
				Expect(createVaultOptionsModel.Name).To(Equal(core.StringPtr("Example Vault")))
				Expect(createVaultOptionsModel.Description).To(Equal(core.StringPtr("The description of the creating vault")))
				Expect(createVaultOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeactivateManagedKeyOptions successfully`, func() {
				// Construct an instance of the DeactivateManagedKeyOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				deactivateManagedKeyOptionsModel := ukoService.NewDeactivateManagedKeyOptions(id, ukoVault, ifMatch)
				deactivateManagedKeyOptionsModel.SetID("testString")
				deactivateManagedKeyOptionsModel.SetUKOVault("testString")
				deactivateManagedKeyOptionsModel.SetIfMatch("testString")
				deactivateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deactivateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(deactivateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deactivateManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(deactivateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deactivateManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteKeyTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteKeyTemplateOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				deleteKeyTemplateOptionsModel := ukoService.NewDeleteKeyTemplateOptions(id, ukoVault, ifMatch)
				deleteKeyTemplateOptionsModel.SetID("testString")
				deleteKeyTemplateOptionsModel.SetUKOVault("testString")
				deleteKeyTemplateOptionsModel.SetIfMatch("testString")
				deleteKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeyTemplateOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteKeystoreOptions successfully`, func() {
				// Construct an instance of the DeleteKeystoreOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				deleteKeystoreOptionsModel := ukoService.NewDeleteKeystoreOptions(id, ukoVault, ifMatch)
				deleteKeystoreOptionsModel.SetID("testString")
				deleteKeystoreOptionsModel.SetUKOVault("testString")
				deleteKeystoreOptionsModel.SetIfMatch("testString")
				deleteKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteKeystoreOptionsModel).ToNot(BeNil())
				Expect(deleteKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeystoreOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeystoreOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(deleteKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteManagedKeyOptions successfully`, func() {
				// Construct an instance of the DeleteManagedKeyOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				deleteManagedKeyOptionsModel := ukoService.NewDeleteManagedKeyOptions(id, ukoVault, ifMatch)
				deleteManagedKeyOptionsModel.SetID("testString")
				deleteManagedKeyOptionsModel.SetUKOVault("testString")
				deleteManagedKeyOptionsModel.SetIfMatch("testString")
				deleteManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteManagedKeyOptionsModel).ToNot(BeNil())
				Expect(deleteManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
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
				ukoVault := "testString"
				ifMatch := "testString"
				destroyManagedKeyOptionsModel := ukoService.NewDestroyManagedKeyOptions(id, ukoVault, ifMatch)
				destroyManagedKeyOptionsModel.SetID("testString")
				destroyManagedKeyOptionsModel.SetUKOVault("testString")
				destroyManagedKeyOptionsModel.SetIfMatch("testString")
				destroyManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(destroyManagedKeyOptionsModel).ToNot(BeNil())
				Expect(destroyManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(destroyManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(destroyManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(destroyManagedKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeyDistributionStatusForKeystoresOptions successfully`, func() {
				// Construct an instance of the GetKeyDistributionStatusForKeystoresOptions model
				id := "testString"
				ukoVault := "testString"
				getKeyDistributionStatusForKeystoresOptionsModel := ukoService.NewGetKeyDistributionStatusForKeystoresOptions(id, ukoVault)
				getKeyDistributionStatusForKeystoresOptionsModel.SetID("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.SetUKOVault("testString")
				getKeyDistributionStatusForKeystoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeyDistributionStatusForKeystoresOptionsModel).ToNot(BeNil())
				Expect(getKeyDistributionStatusForKeystoresOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeyDistributionStatusForKeystoresOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(getKeyDistributionStatusForKeystoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeyTemplateOptions successfully`, func() {
				// Construct an instance of the GetKeyTemplateOptions model
				id := "testString"
				ukoVault := "testString"
				getKeyTemplateOptionsModel := ukoService.NewGetKeyTemplateOptions(id, ukoVault)
				getKeyTemplateOptionsModel.SetID("testString")
				getKeyTemplateOptionsModel.SetUKOVault("testString")
				getKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(getKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeyTemplateOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(getKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeystoreOptions successfully`, func() {
				// Construct an instance of the GetKeystoreOptions model
				id := "testString"
				ukoVault := "testString"
				getKeystoreOptionsModel := ukoService.NewGetKeystoreOptions(id, ukoVault)
				getKeystoreOptionsModel.SetID("testString")
				getKeystoreOptionsModel.SetUKOVault("testString")
				getKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeystoreOptionsModel).ToNot(BeNil())
				Expect(getKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKeystoreStatusOptions successfully`, func() {
				// Construct an instance of the GetKeystoreStatusOptions model
				id := "testString"
				ukoVault := "testString"
				getKeystoreStatusOptionsModel := ukoService.NewGetKeystoreStatusOptions(id, ukoVault)
				getKeystoreStatusOptionsModel.SetID("testString")
				getKeystoreStatusOptionsModel.SetUKOVault("testString")
				getKeystoreStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKeystoreStatusOptionsModel).ToNot(BeNil())
				Expect(getKeystoreStatusOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreStatusOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(getKeystoreStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetManagedKeyOptions successfully`, func() {
				// Construct an instance of the GetManagedKeyOptions model
				id := "testString"
				ukoVault := "testString"
				getManagedKeyOptionsModel := ukoService.NewGetManagedKeyOptions(id, ukoVault)
				getManagedKeyOptionsModel.SetID("testString")
				getManagedKeyOptionsModel.SetUKOVault("testString")
				getManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getManagedKeyOptionsModel).ToNot(BeNil())
				Expect(getManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
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
			It(`Invoke NewListKeyTemplatesOptions successfully`, func() {
				// Construct an instance of the ListKeyTemplatesOptions model
				listKeyTemplatesOptionsModel := ukoService.NewListKeyTemplatesOptions()
				listKeyTemplatesOptionsModel.SetVaultID([]string{"123e4567-e89b-12d3-a456-426614174000"})
				listKeyTemplatesOptionsModel.SetKeyAlgorithm("aes")
				listKeyTemplatesOptionsModel.SetSort([]string{"-updated_at"})
				listKeyTemplatesOptionsModel.SetLimit(int64(10))
				listKeyTemplatesOptionsModel.SetOffset(int64(0))
				listKeyTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listKeyTemplatesOptionsModel).ToNot(BeNil())
				Expect(listKeyTemplatesOptionsModel.VaultID).To(Equal([]string{"123e4567-e89b-12d3-a456-426614174000"}))
				Expect(listKeyTemplatesOptionsModel.KeyAlgorithm).To(Equal(core.StringPtr("aes")))
				Expect(listKeyTemplatesOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listKeyTemplatesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listKeyTemplatesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listKeyTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListKeystoresOptions successfully`, func() {
				// Construct an instance of the ListKeystoresOptions model
				listKeystoresOptionsModel := ukoService.NewListKeystoresOptions()
				listKeystoresOptionsModel.SetType([]string{"ibm_cloud_kms"})
				listKeystoresOptionsModel.SetName("Main IBM Cloud")
				listKeystoresOptionsModel.SetDescription("My Example Keystore Description")
				listKeystoresOptionsModel.SetGroups("testString")
				listKeystoresOptionsModel.SetVaultID([]string{"123e4567-e89b-12d3-a456-426614174000"})
				listKeystoresOptionsModel.SetLocation("testString")
				listKeystoresOptionsModel.SetLimit(int64(10))
				listKeystoresOptionsModel.SetOffset(int64(0))
				listKeystoresOptionsModel.SetSort([]string{"-updated_at"})
				listKeystoresOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listKeystoresOptionsModel).ToNot(BeNil())
				Expect(listKeystoresOptionsModel.Type).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listKeystoresOptionsModel.Name).To(Equal(core.StringPtr("Main IBM Cloud")))
				Expect(listKeystoresOptionsModel.Description).To(Equal(core.StringPtr("My Example Keystore Description")))
				Expect(listKeystoresOptionsModel.Groups).To(Equal(core.StringPtr("testString")))
				Expect(listKeystoresOptionsModel.VaultID).To(Equal([]string{"123e4567-e89b-12d3-a456-426614174000"}))
				Expect(listKeystoresOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(listKeystoresOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listKeystoresOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listKeystoresOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listKeystoresOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListManagedKeysFromKeystoreOptions successfully`, func() {
				// Construct an instance of the ListManagedKeysFromKeystoreOptions model
				ukoVault := "testString"
				id := "testString"
				listManagedKeysFromKeystoreOptionsModel := ukoService.NewListManagedKeysFromKeystoreOptions(ukoVault, id)
				listManagedKeysFromKeystoreOptionsModel.SetUKOVault("testString")
				listManagedKeysFromKeystoreOptionsModel.SetID("testString")
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
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAt("testString")
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAtMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetCreatedAtMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAt("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAtMin("testString")
				listManagedKeysFromKeystoreOptionsModel.SetUpdatedAtMax("testString")
				listManagedKeysFromKeystoreOptionsModel.SetSize(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SetSizeMin(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SetSizeMax(int64(38))
				listManagedKeysFromKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listManagedKeysFromKeystoreOptionsModel).ToNot(BeNil())
				Expect(listManagedKeysFromKeystoreOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
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
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysFromKeystoreOptionsModel.Size).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysFromKeystoreOptionsModel.SizeMin).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysFromKeystoreOptionsModel.SizeMax).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysFromKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListManagedKeysOptions successfully`, func() {
				// Construct an instance of the ListManagedKeysOptions model
				listManagedKeysOptionsModel := ukoService.NewListManagedKeysOptions()
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
				listManagedKeysOptionsModel.SetCreatedAt("testString")
				listManagedKeysOptionsModel.SetCreatedAtMin("testString")
				listManagedKeysOptionsModel.SetCreatedAtMax("testString")
				listManagedKeysOptionsModel.SetUpdatedAt("testString")
				listManagedKeysOptionsModel.SetUpdatedAtMin("testString")
				listManagedKeysOptionsModel.SetUpdatedAtMax("testString")
				listManagedKeysOptionsModel.SetSize(int64(38))
				listManagedKeysOptionsModel.SetSizeMin(int64(38))
				listManagedKeysOptionsModel.SetSizeMax(int64(38))
				listManagedKeysOptionsModel.SetReferencedKeystoresType([]string{"ibm_cloud_kms"})
				listManagedKeysOptionsModel.SetReferencedKeystoresName([]string{"testString"})
				listManagedKeysOptionsModel.SetInstancesKeystoreType([]string{"ibm_cloud_kms"})
				listManagedKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listManagedKeysOptionsModel).ToNot(BeNil())
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
				Expect(listManagedKeysOptionsModel.CreatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.CreatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.CreatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAt).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAtMin).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.UpdatedAtMax).To(Equal(core.StringPtr("testString")))
				Expect(listManagedKeysOptionsModel.Size).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysOptionsModel.SizeMin).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysOptionsModel.SizeMax).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listManagedKeysOptionsModel.ReferencedKeystoresType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeysOptionsModel.ReferencedKeystoresName).To(Equal([]string{"testString"}))
				Expect(listManagedKeysOptionsModel.InstancesKeystoreType).To(Equal([]string{"ibm_cloud_kms"}))
				Expect(listManagedKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVaultsOptions successfully`, func() {
				// Construct an instance of the ListVaultsOptions model
				listVaultsOptionsModel := ukoService.NewListVaultsOptions()
				listVaultsOptionsModel.SetLimit(int64(10))
				listVaultsOptionsModel.SetOffset(int64(0))
				listVaultsOptionsModel.SetSort([]string{"-updated_at"})
				listVaultsOptionsModel.SetName("My Example Vault")
				listVaultsOptionsModel.SetDescription("My Example Vault Description")
				listVaultsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVaultsOptionsModel).ToNot(BeNil())
				Expect(listVaultsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listVaultsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listVaultsOptionsModel.Sort).To(Equal([]string{"-updated_at"}))
				Expect(listVaultsOptionsModel.Name).To(Equal(core.StringPtr("My Example Vault")))
				Expect(listVaultsOptionsModel.Description).To(Equal(core.StringPtr("My Example Vault Description")))
				Expect(listVaultsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateKeyTemplateOptions successfully`, func() {
				// Construct an instance of the KeystoresPropertiesUpdate model
				keystoresPropertiesUpdateModel := new(ukov4.KeystoresPropertiesUpdate)
				Expect(keystoresPropertiesUpdateModel).ToNot(BeNil())
				keystoresPropertiesUpdateModel.Group = core.StringPtr("Production")
				Expect(keystoresPropertiesUpdateModel.Group).To(Equal(core.StringPtr("Production")))

				// Construct an instance of the KeyPropertiesUpdate model
				keyPropertiesUpdateModel := new(ukov4.KeyPropertiesUpdate)
				Expect(keyPropertiesUpdateModel).ToNot(BeNil())
				keyPropertiesUpdateModel.Size = core.StringPtr("256")
				keyPropertiesUpdateModel.ActivationDate = core.StringPtr("P5Y1M1W2D")
				keyPropertiesUpdateModel.ExpirationDate = core.StringPtr("P1Y2M1W4D")
				keyPropertiesUpdateModel.State = core.StringPtr("active")
				Expect(keyPropertiesUpdateModel.Size).To(Equal(core.StringPtr("256")))
				Expect(keyPropertiesUpdateModel.ActivationDate).To(Equal(core.StringPtr("P5Y1M1W2D")))
				Expect(keyPropertiesUpdateModel.ExpirationDate).To(Equal(core.StringPtr("P1Y2M1W4D")))
				Expect(keyPropertiesUpdateModel.State).To(Equal(core.StringPtr("active")))

				// Construct an instance of the UpdateKeyTemplateOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				updateKeyTemplateOptionsModel := ukoService.NewUpdateKeyTemplateOptions(id, ukoVault, ifMatch)
				updateKeyTemplateOptionsModel.SetID("testString")
				updateKeyTemplateOptionsModel.SetUKOVault("testString")
				updateKeyTemplateOptionsModel.SetIfMatch("testString")
				updateKeyTemplateOptionsModel.SetKeystores([]ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel})
				updateKeyTemplateOptionsModel.SetDescription("testString")
				updateKeyTemplateOptionsModel.SetKey(keyPropertiesUpdateModel)
				updateKeyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateKeyTemplateOptionsModel).ToNot(BeNil())
				Expect(updateKeyTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.Keystores).To(Equal([]ukov4.KeystoresPropertiesUpdate{*keystoresPropertiesUpdateModel}))
				Expect(updateKeyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateKeyTemplateOptionsModel.Key).To(Equal(keyPropertiesUpdateModel))
				Expect(updateKeyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateKeystoreOptions successfully`, func() {
				// Construct an instance of the KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate model
				keystoreUpdateRequestModel := new(ukov4.KeystoreUpdateRequestKeystoreTypeAwsKmsUpdate)
				Expect(keystoreUpdateRequestModel).ToNot(BeNil())
				keystoreUpdateRequestModel.Name = core.StringPtr("IBM Cloud Keystore Name")
				keystoreUpdateRequestModel.Description = core.StringPtr("Azure keystore")
				keystoreUpdateRequestModel.Groups = []string{"Production"}
				keystoreUpdateRequestModel.AwsRegion = core.StringPtr("af_south_1")
				keystoreUpdateRequestModel.AwsAccessKeyID = core.StringPtr("BSDFWERUANLKJDN54AAS")
				keystoreUpdateRequestModel.AwsSecretAccessKey = core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")
				Expect(keystoreUpdateRequestModel.Name).To(Equal(core.StringPtr("IBM Cloud Keystore Name")))
				Expect(keystoreUpdateRequestModel.Description).To(Equal(core.StringPtr("Azure keystore")))
				Expect(keystoreUpdateRequestModel.Groups).To(Equal([]string{"Production"}))
				Expect(keystoreUpdateRequestModel.AwsRegion).To(Equal(core.StringPtr("af_south_1")))
				Expect(keystoreUpdateRequestModel.AwsAccessKeyID).To(Equal(core.StringPtr("BSDFWERUANLKJDN54AAS")))
				Expect(keystoreUpdateRequestModel.AwsSecretAccessKey).To(Equal(core.StringPtr("6HSz234KBjMrASFasfg5PasAFGNasg87asdgQzgs")))

				// Construct an instance of the UpdateKeystoreOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				var keystoreBody ukov4.KeystoreUpdateRequestIntf = nil
				updateKeystoreOptionsModel := ukoService.NewUpdateKeystoreOptions(id, ukoVault, ifMatch, keystoreBody)
				updateKeystoreOptionsModel.SetID("testString")
				updateKeystoreOptionsModel.SetUKOVault("testString")
				updateKeystoreOptionsModel.SetIfMatch("testString")
				updateKeystoreOptionsModel.SetKeystoreBody(keystoreUpdateRequestModel)
				updateKeystoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateKeystoreOptionsModel).ToNot(BeNil())
				Expect(updateKeystoreOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateKeystoreOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(updateKeystoreOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateKeystoreOptionsModel.KeystoreBody).To(Equal(keystoreUpdateRequestModel))
				Expect(updateKeystoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateManagedKeyFromTemplateOptions successfully`, func() {
				// Construct an instance of the UpdateManagedKeyFromTemplateOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				updateManagedKeyFromTemplateOptionsModel := ukoService.NewUpdateManagedKeyFromTemplateOptions(id, ukoVault, ifMatch)
				updateManagedKeyFromTemplateOptionsModel.SetID("testString")
				updateManagedKeyFromTemplateOptionsModel.SetUKOVault("testString")
				updateManagedKeyFromTemplateOptionsModel.SetIfMatch("testString")
				updateManagedKeyFromTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateManagedKeyFromTemplateOptionsModel).ToNot(BeNil())
				Expect(updateManagedKeyFromTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyFromTemplateOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyFromTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyFromTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateManagedKeyOptions successfully`, func() {
				// Construct an instance of the Tag model
				tagModel := new(ukov4.Tag)
				Expect(tagModel).ToNot(BeNil())
				tagModel.Name = core.StringPtr("testString")
				tagModel.Value = core.StringPtr("testString")
				Expect(tagModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(tagModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateManagedKeyOptions model
				id := "testString"
				ukoVault := "testString"
				ifMatch := "testString"
				updateManagedKeyOptionsModel := ukoService.NewUpdateManagedKeyOptions(id, ukoVault, ifMatch)
				updateManagedKeyOptionsModel.SetID("testString")
				updateManagedKeyOptionsModel.SetUKOVault("testString")
				updateManagedKeyOptionsModel.SetIfMatch("testString")
				updateManagedKeyOptionsModel.SetLabel("IBM CLOUD KEY")
				updateManagedKeyOptionsModel.SetActivationDate(CreateMockDate("2020-12-11"))
				updateManagedKeyOptionsModel.SetExpirationDate(CreateMockDate("2030-11-12"))
				updateManagedKeyOptionsModel.SetTags([]ukov4.Tag{*tagModel})
				updateManagedKeyOptionsModel.SetDescription("testString")
				updateManagedKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateManagedKeyOptionsModel).ToNot(BeNil())
				Expect(updateManagedKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.UKOVault).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateManagedKeyOptionsModel.Label).To(Equal(core.StringPtr("IBM CLOUD KEY")))
				Expect(updateManagedKeyOptionsModel.ActivationDate).To(Equal(CreateMockDate("2020-12-11")))
				Expect(updateManagedKeyOptionsModel.ExpirationDate).To(Equal(CreateMockDate("2030-11-12")))
				Expect(updateManagedKeyOptionsModel.Tags).To(Equal([]ukov4.Tag{*tagModel}))
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
				updateVaultOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVaultOptionsModel).ToNot(BeNil())
				Expect(updateVaultOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVaultOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateVaultOptionsModel.Name).To(Equal(core.StringPtr("Jakub's Vault")))
				Expect(updateVaultOptionsModel.Description).To(Equal(core.StringPtr("Updated description of the vault")))
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
			It(`Invoke NewKeystoresProperties successfully`, func() {
				group := "Production"
				typeVar := "ibm_cloud_kms"
				_model, err := ukoService.NewKeystoresProperties(group, typeVar)
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
				azureLocation := "europe_north"
				azureServicePrincipalClientID := "91018db5-c756-468e-bd4e-69c99fc1a749"
				azureServicePrincipalPassword := "9wN1YP5XwrrHIdvIYv7imHiC83Q_lSWAWa"
				azureTenant := "b8e1a93c-2449-462f-8fa0-1d00595ea859"
				azureSubscriptionID := "a98667h9b-5fhf-42f3-9392-26856b045g08"
				azureEnvironment := "azure"
				_, err := ukoService.NewKeystoreCreationRequestKeystoreTypeAzureCreate(typeVar, vault, azureServiceName, azureResourceGroup, azureLocation, azureServicePrincipalClientID, azureServicePrincipalPassword, azureTenant, azureSubscriptionID, azureEnvironment)
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
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
