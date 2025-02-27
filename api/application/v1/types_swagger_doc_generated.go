/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_App = map[string]string{
	"":     "App is a app bootstrap in TKE.",
	"spec": "Spec defines the desired identities of bootstrap in this set.",
}

func (App) SwaggerDoc() map[string]string {
	return map_App
}

var map_AppHistory = map[string]string{
	"":     "AppHistory is a app history in TKE.",
	"spec": "Spec defines the desired identities of bootstrap in this set.",
}

func (AppHistory) SwaggerDoc() map[string]string {
	return map_AppHistory
}

var map_AppHistorySpec = map[string]string{
	"": "AppHistorySpec is a description of a AppHistory.",
}

func (AppHistorySpec) SwaggerDoc() map[string]string {
	return map_AppHistorySpec
}

var map_AppList = map[string]string{
	"":      "AppList is the whole list of all bootstraps.",
	"items": "List of bootstraps",
}

func (AppList) SwaggerDoc() map[string]string {
	return map_AppList
}

var map_AppResource = map[string]string{
	"":     "AppResource is a app resource in TKE.",
	"spec": "Spec defines the desired identities of bootstrap in this set.",
}

func (AppResource) SwaggerDoc() map[string]string {
	return map_AppResource
}

var map_AppResourceSpec = map[string]string{
	"": "AppResourceSpec is a description of a AppResource.",
}

func (AppResourceSpec) SwaggerDoc() map[string]string {
	return map_AppResourceSpec
}

var map_AppSpec = map[string]string{
	"":       "AppSpec is a description of a project.",
	"values": "Values holds the values for this app.",
}

func (AppSpec) SwaggerDoc() map[string]string {
	return map_AppSpec
}

var map_AppStatus = map[string]string{
	"":                   "AppStatus represents information about the status of a bootstrap.",
	"phase":              "Phase the release is in, one of ('ChartFetched', 'ChartFetchFailed', 'Installing', 'Upgrading', 'Succeeded', 'RollingBack', 'RolledBack', 'RollbackFailed')",
	"observedGeneration": "ObservedGeneration is the most recent generation observed by the operator.",
	"releaseStatus":      "ReleaseStatus is the status as given by Helm for the release managed by this resource.",
	"releaseLastUpdated": "ReleaseLastUpdated is the last updated time for the release",
	"revision":           "Revision holds the Git hash or version of the chart currently deployed.",
	"rollbackRevision":   "RollbackRevision specify the target rollback version of the chart",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
	"manifest":           "Dryrun result.",
}

func (AppStatus) SwaggerDoc() map[string]string {
	return map_AppStatus
}

var map_AppValues = map[string]string{
	"": "AppValues string the values for this app.",
}

func (AppValues) SwaggerDoc() map[string]string {
	return map_AppValues
}

var map_Chart = map[string]string{
	"":                "Chart is a description of a chart.",
	"chartName":       "ChartName is the name of the chart.",
	"chartVersion":    "ChartVersion is the version of the chart.",
	"createNamespace": "CreateNamespace create namespace when install helm release",
	"atomic":          "Atomic, if true, for install case, will uninstall failed release, for upgrade case, will roll back on failure.",
}

func (Chart) SwaggerDoc() map[string]string {
	return map_Chart
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_History = map[string]string{
	"": "History is a history of a app.",
}

func (History) SwaggerDoc() map[string]string {
	return map_History
}

var map_RollbackProxyOptions = map[string]string{
	"": "RollbackProxyOptions is the query options to an app rollback proxy call.",
}

func (RollbackProxyOptions) SwaggerDoc() map[string]string {
	return map_RollbackProxyOptions
}

// AUTO-GENERATED FUNCTIONS END HERE
