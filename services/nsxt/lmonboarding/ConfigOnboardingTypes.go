/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ConfigOnboarding.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lmonboarding

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func configOnboardingCompleteonboardingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_LM_onboarding_status"] = bindings.NewReferenceType(model.PolicyLMOnboardingStatusBindingType)
	fieldNameMap["policy_LM_onboarding_status"] = "PolicyLMOnboardingStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func configOnboardingCompleteonboardingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyLMOnboardingStatusBindingType)
}

func configOnboardingCompleteonboardingRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["policy_LM_onboarding_status"] = bindings.NewReferenceType(model.PolicyLMOnboardingStatusBindingType)
	fieldNameMap["policy_LM_onboarding_status"] = "PolicyLMOnboardingStatus"
	paramsTypeMap["policy_LM_onboarding_status"] = bindings.NewReferenceType(model.PolicyLMOnboardingStatusBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=complete_onboarding",
		"policy_LM_onboarding_status",
		"POST",
		"/policy/api/v1/lmonboarding/config-onboarding",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func configOnboardingStartonboardingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy_LM_onboarding_metadata"] = bindings.NewReferenceType(model.PolicyLMOnboardingMetadataBindingType)
	fieldNameMap["policy_LM_onboarding_metadata"] = "PolicyLMOnboardingMetadata"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func configOnboardingStartonboardingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyLMOnboardingStatusBindingType)
}

func configOnboardingStartonboardingRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["policy_LM_onboarding_metadata"] = bindings.NewReferenceType(model.PolicyLMOnboardingMetadataBindingType)
	fieldNameMap["policy_LM_onboarding_metadata"] = "PolicyLMOnboardingMetadata"
	paramsTypeMap["policy_LM_onboarding_metadata"] = bindings.NewReferenceType(model.PolicyLMOnboardingMetadataBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=start_onboarding",
		"policy_LM_onboarding_metadata",
		"POST",
		"/policy/api/v1/lmonboarding/config-onboarding",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

