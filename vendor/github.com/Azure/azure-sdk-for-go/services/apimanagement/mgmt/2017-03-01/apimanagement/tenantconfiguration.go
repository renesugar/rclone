package apimanagement

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

// TenantConfigurationClient is the apiManagement Client
type TenantConfigurationClient struct {
	BaseClient
}

// NewTenantConfigurationClient creates an instance of the TenantConfigurationClient client.
func NewTenantConfigurationClient(subscriptionID string) TenantConfigurationClient {
	return NewTenantConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTenantConfigurationClientWithBaseURI creates an instance of the TenantConfigurationClient client.
func NewTenantConfigurationClientWithBaseURI(baseURI string, subscriptionID string) TenantConfigurationClient {
	return TenantConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deploy this operation applies changes from the specified Git branch to the configuration database. This is a long
// running operation and could take several minutes to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// parameters - deploy Configuration parameters.
func (client TenantConfigurationClient) Deploy(ctx context.Context, resourceGroupName string, serviceName string, parameters DeployConfigurationParameters) (result TenantConfigurationDeployFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Deploy", err.Error())
	}

	req, err := client.DeployPreparer(ctx, resourceGroupName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", nil, "Failure preparing request")
		return
	}

	result, err = client.DeploySender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Deploy", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeployPreparer prepares the Deploy request.
func (client TenantConfigurationClient) DeployPreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters DeployConfigurationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/deploy", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeploySender sends the Deploy request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) DeploySender(req *http.Request) (future TenantConfigurationDeployFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// DeployResponder handles the response to the Deploy request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) DeployResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSyncState gets the status of the most recent synchronization between the configuration database and the Git
// repository.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
func (client TenantConfigurationClient) GetSyncState(ctx context.Context, resourceGroupName string, serviceName string) (result TenantConfigurationSyncStateContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "GetSyncState", err.Error())
	}

	req, err := client.GetSyncStatePreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSyncStateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", resp, "Failure sending request")
		return
	}

	result, err = client.GetSyncStateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "GetSyncState", resp, "Failure responding to request")
	}

	return
}

// GetSyncStatePreparer prepares the GetSyncState request.
func (client TenantConfigurationClient) GetSyncStatePreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/syncState", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSyncStateSender sends the GetSyncState request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) GetSyncStateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetSyncStateResponder handles the response to the GetSyncState request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) GetSyncStateResponder(resp *http.Response) (result TenantConfigurationSyncStateContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Save this operation creates a commit with the current configuration snapshot to the specified branch in the
// repository. This is a long running operation and could take several minutes to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// parameters - save Configuration parameters.
func (client TenantConfigurationClient) Save(ctx context.Context, resourceGroupName string, serviceName string, parameters SaveConfigurationParameter) (result TenantConfigurationSaveFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Save", err.Error())
	}

	req, err := client.SavePreparer(ctx, resourceGroupName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", nil, "Failure preparing request")
		return
	}

	result, err = client.SaveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Save", result.Response(), "Failure sending request")
		return
	}

	return
}

// SavePreparer prepares the Save request.
func (client TenantConfigurationClient) SavePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters SaveConfigurationParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/save", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SaveSender sends the Save request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) SaveSender(req *http.Request) (future TenantConfigurationSaveFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// SaveResponder handles the response to the Save request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) SaveResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate this operation validates the changes in the specified Git branch. This is a long running operation and
// could take several minutes to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// parameters - validate Configuration parameters.
func (client TenantConfigurationClient) Validate(ctx context.Context, resourceGroupName string, serviceName string, parameters DeployConfigurationParameters) (result TenantConfigurationValidateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Branch", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.TenantConfigurationClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, resourceGroupName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", nil, "Failure preparing request")
		return
	}

	result, err = client.ValidateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.TenantConfigurationClient", "Validate", result.Response(), "Failure sending request")
		return
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client TenantConfigurationClient) ValidatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters DeployConfigurationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configurationName": autorest.Encode("path", "configuration"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/validate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client TenantConfigurationClient) ValidateSender(req *http.Request) (future TenantConfigurationValidateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client TenantConfigurationClient) ValidateResponder(resp *http.Response) (result OperationResultContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
