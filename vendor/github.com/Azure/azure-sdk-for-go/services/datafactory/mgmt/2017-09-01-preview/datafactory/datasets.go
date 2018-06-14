package datafactory

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// DatasetsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that interact
// with Azure Data Factory V2 services.
type DatasetsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// NewDatasetsClient creates an instance of the DatasetsClient client.
func NewDatasetsClient(subscriptionID string) DatasetsClient {
	return NewDatasetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// NewDatasetsClientWithBaseURI creates an instance of the DatasetsClient client.
func NewDatasetsClientWithBaseURI(baseURI string, subscriptionID string) DatasetsClient {
	return DatasetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// CreateOrUpdate creates or updates a dataset.
//
// resourceGroupName is the resource group name. factoryName is the factory name. datasetName is the dataset name.
// dataset is dataset resource definition. ifMatch is eTag of the dataset entity.  Should only be specified for
// update, for which it should match existing entity or can be * for unconditional update.
func (client DatasetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, ifMatch string) (result DatasetResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: datasetName,
			Constraints: []validation.Constraint{{Target: "datasetName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "datasetName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "datasetName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}},
		{TargetValue: dataset,
			Constraints: []validation.Constraint{{Target: "dataset.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "dataset.Properties.LinkedServiceName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "dataset.Properties.LinkedServiceName.Type", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "dataset.Properties.LinkedServiceName.ReferenceName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("datafactory.DatasetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, datasetName, dataset, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatasetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"datasetName":       autorest.Encode("path", datasetName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", pathParameters),
		autorest.WithJSON(dataset),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatasetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DatasetsClient) CreateOrUpdateResponder(resp *http.Response) (result DatasetResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// Delete deletes a dataset.
//
// resourceGroupName is the resource group name. factoryName is the factory name. datasetName is the dataset name.
func (client DatasetsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, datasetName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: datasetName,
			Constraints: []validation.Constraint{{Target: "datasetName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "datasetName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "datasetName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DatasetsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, datasetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// DeletePreparer prepares the Delete request.
func (client DatasetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, datasetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"datasetName":       autorest.Encode("path", datasetName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DatasetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DatasetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// Get gets a dataset.
//
// resourceGroupName is the resource group name. factoryName is the factory name. datasetName is the dataset name.
func (client DatasetsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, datasetName string) (result DatasetResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: datasetName,
			Constraints: []validation.Constraint{{Target: "datasetName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "datasetName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "datasetName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DatasetsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, datasetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetPreparer prepares the Get request.
func (client DatasetsClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, datasetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"datasetName":       autorest.Encode("path", datasetName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatasetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatasetsClient) GetResponder(resp *http.Response) (result DatasetResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// ListByFactory lists datasets.
//
// resourceGroupName is the resource group name. factoryName is the factory name.
func (client DatasetsClient) ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result DatasetListResponsePage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.DatasetsClient", "ListByFactory", err.Error())
	}

	result.fn = client.listByFactoryNextResults
	req, err := client.ListByFactoryPreparer(ctx, resourceGroupName, factoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "ListByFactory", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "ListByFactory", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// ListByFactoryPreparer prepares the ListByFactory request.
func (client DatasetsClient) ListByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// ListByFactorySender sends the ListByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client DatasetsClient) ListByFactorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// ListByFactoryResponder handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (client DatasetsClient) ListByFactoryResponder(resp *http.Response) (result DatasetListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFactoryNextResults retrieves the next set of results, if any.
func (client DatasetsClient) listByFactoryNextResults(lastResults DatasetListResponse) (result DatasetListResponse, err error) {
	req, err := lastResults.datasetListResponsePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "listByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "listByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.DatasetsClient", "listByFactoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/datafactory/mgmt/2017-09-01-preview/datafactory instead.
// ListByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client DatasetsClient) ListByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string) (result DatasetListResponseIterator, err error) {
	result.page, err = client.ListByFactory(ctx, resourceGroupName, factoryName)
	return
}
