# \DefaultAPI

All URIs are relative to *https://api.us2.sumologic.com/api/sec/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntityNote**](DefaultAPI.md#AddEntityNote) | **Post** /entities/{entityId}/notes | Add a new Entity Note on an Entity
[**AddEntityTags**](DefaultAPI.md#AddEntityTags) | **Post** /entities/{entityId}/tags | Add tags to an Entity
[**AddIndicatorToThreatIntelSource**](DefaultAPI.md#AddIndicatorToThreatIntelSource) | **Post** /threat-intel-sources/{threatIntelSourceId}/items | Add Indicators to a Threat Intel Source
[**AddInsightComment**](DefaultAPI.md#AddInsightComment) | **Post** /insights/{id}/comments | Add a new comment on an Insight
[**AddItemsToMatchList**](DefaultAPI.md#AddItemsToMatchList) | **Post** /match-lists/{id}/items | Add Match List Items to a Match List
[**AddItemsToSuppressList**](DefaultAPI.md#AddItemsToSuppressList) | **Post** /suppress-lists/{id}/items | Add Suppress List Items to a Suppress List
[**AddSignalsToInsight**](DefaultAPI.md#AddSignalsToInsight) | **Put** /insights/{insightId}/signals | Add Signals to an existing Insight
[**AddTagToInsight**](DefaultAPI.md#AddTagToInsight) | **Post** /insights/{id}/tags | Add a tag to an Insight
[**BulkAddEntityTags**](DefaultAPI.md#BulkAddEntityTags) | **Post** /entities/bulk-add-tags | Bulk add tags to entities
[**BulkDeleteMatchListItems**](DefaultAPI.md#BulkDeleteMatchListItems) | **Post** /match-list-items/bulk-delete | Bulk delete Match List Item
[**BulkRemoveEntityTags**](DefaultAPI.md#BulkRemoveEntityTags) | **Post** /entities/bulk-remove-tags | Bulk remove tags on entities
[**BulkUpdateEntityCriticality**](DefaultAPI.md#BulkUpdateEntityCriticality) | **Post** /entities/bulk-update-criticality | Bulk updates criticality entity field
[**BulkUpdateEntityTags**](DefaultAPI.md#BulkUpdateEntityTags) | **Post** /entities/bulk-update-tags | Bulk update tags on entities
[**BulkUpdateEntityWhitelist**](DefaultAPI.md#BulkUpdateEntityWhitelist) | **Post** /entities/bulk-update-suppressed | Bulk updates suppressed entity field
[**BulkUpsertNetworkBlocks**](DefaultAPI.md#BulkUpsertNetworkBlocks) | **Post** /network-blocks/bulk | Add or update multiple Network Blocks in one request
[**CreateAggregationRule**](DefaultAPI.md#CreateAggregationRule) | **Post** /rules/aggregation | Create a Aggregation Rule
[**CreateAutomation**](DefaultAPI.md#CreateAutomation) | **Post** /automations | Create an Automation
[**CreateChainRule**](DefaultAPI.md#CreateChainRule) | **Post** /rules/chain | Create a Chain Rule
[**CreateContextAction**](DefaultAPI.md#CreateContextAction) | **Post** /context-actions | Create a Context Action
[**CreateCustomEntityType**](DefaultAPI.md#CreateCustomEntityType) | **Post** /custom-entity-types | Create a Custom Entity Type
[**CreateCustomInsight**](DefaultAPI.md#CreateCustomInsight) | **Post** /custom-insights | Create a Custom Insight
[**CreateCustomMatchListColumn**](DefaultAPI.md#CreateCustomMatchListColumn) | **Post** /custom-match-list-columns | Create a Custom Match List Column
[**CreateCustomerSourcedEntityLookupTable**](DefaultAPI.md#CreateCustomerSourcedEntityLookupTable) | **Post** /entity-lookup-tables/customer-sourced | Create a customer sourced entity lookup table
[**CreateEntityCriticalityConfig**](DefaultAPI.md#CreateEntityCriticalityConfig) | **Post** /entity-criticality-configs | Create an Entity Criticality Configuration
[**CreateEntityValue**](DefaultAPI.md#CreateEntityValue) | **Post** /entity-group-configurations/entity_value | Create an Entity Value Group
[**CreateFirstSeenRule**](DefaultAPI.md#CreateFirstSeenRule) | **Post** /rules/first-seen | Create a First Seen Rule
[**CreateInsightFromSignals**](DefaultAPI.md#CreateInsightFromSignals) | **Post** /insights | Create Insight from Signals
[**CreateInsightResolution**](DefaultAPI.md#CreateInsightResolution) | **Post** /insight-resolutions | Create a Insight Resolution
[**CreateInsightStatusOption**](DefaultAPI.md#CreateInsightStatusOption) | **Post** /insight-status | Create an Insight Status
[**CreateInventory**](DefaultAPI.md#CreateInventory) | **Post** /entity-group-configurations/inventory | Create an Inventory Group
[**CreateLogMapping**](DefaultAPI.md#CreateLogMapping) | **Post** /log-mappings | Create a Log Mapping
[**CreateMatchList**](DefaultAPI.md#CreateMatchList) | **Post** /match-lists | Create a Match List
[**CreateMatchRule**](DefaultAPI.md#CreateMatchRule) | **Post** /rules/match | Create a Legacy Match Rule
[**CreateNetworkBlock**](DefaultAPI.md#CreateNetworkBlock) | **Post** /network-blocks | Create a Network Block
[**CreateRuleTuningExpression**](DefaultAPI.md#CreateRuleTuningExpression) | **Post** /rule-tuning-expressions | Create a Rule Tuning Expression
[**CreateSuppressList**](DefaultAPI.md#CreateSuppressList) | **Post** /suppress-lists | Create a Suppress List
[**CreateTagSchema**](DefaultAPI.md#CreateTagSchema) | **Post** /tag-schemas | Create a Tag Schema
[**CreateTemplatedMatchRule**](DefaultAPI.md#CreateTemplatedMatchRule) | **Post** /rules/templated | Create a Match Rule
[**CreateThreatIntelSource**](DefaultAPI.md#CreateThreatIntelSource) | **Post** /threat-intel-sources | Create a Threat Intel Source
[**CreateThresholdRule**](DefaultAPI.md#CreateThresholdRule) | **Post** /rules/threshold | Create a Threshold Rule
[**DeleteAutomation**](DefaultAPI.md#DeleteAutomation) | **Delete** /automations/{id} | Delete an Automation
[**DeleteContextAction**](DefaultAPI.md#DeleteContextAction) | **Delete** /context-actions/{id} | Delete a Context Action
[**DeleteCustomEntityType**](DefaultAPI.md#DeleteCustomEntityType) | **Delete** /custom-entity-types/{id} | Delete a Custom Entity Type
[**DeleteCustomInsight**](DefaultAPI.md#DeleteCustomInsight) | **Delete** /custom-insights/{id} | Delete a Custom Insight
[**DeleteCustomMatchListColumn**](DefaultAPI.md#DeleteCustomMatchListColumn) | **Delete** /custom-match-list-columns/{id} | Delete a Custom Match List Column
[**DeleteCustomerSourcedEntityLookupTable**](DefaultAPI.md#DeleteCustomerSourcedEntityLookupTable) | **Delete** /entity-lookup-tables/customer-source/{id} | Remove a customer sourced entity lookup table
[**DeleteEntityCriticalityConfig**](DefaultAPI.md#DeleteEntityCriticalityConfig) | **Delete** /entity-criticality-configs/{id} | Delete an Entity Criticality Configuration
[**DeleteEntityGroup**](DefaultAPI.md#DeleteEntityGroup) | **Delete** /entity-group-configurations/{id} | Delete an Entity Group Configuration
[**DeleteEntityNote**](DefaultAPI.md#DeleteEntityNote) | **Delete** /entities/{entityId}/notes/{entityNoteId} | Remove an Entity Note from an Entity
[**DeleteInsightResolution**](DefaultAPI.md#DeleteInsightResolution) | **Delete** /insight-resolutions/{id} | Delete a Insight Resolution
[**DeleteInsightStatusOption**](DefaultAPI.md#DeleteInsightStatusOption) | **Delete** /insight-status/{id} | Delete an Insight Status
[**DeleteLogMapping**](DefaultAPI.md#DeleteLogMapping) | **Delete** /log-mappings/{id} | Delete a Log Mapping
[**DeleteMatchList**](DefaultAPI.md#DeleteMatchList) | **Delete** /match-lists/{id} | Delete a Match List
[**DeleteMatchListItem**](DefaultAPI.md#DeleteMatchListItem) | **Delete** /match-list-items/{id} | Delete a Match List Item
[**DeleteNetworkBlock**](DefaultAPI.md#DeleteNetworkBlock) | **Delete** /network-blocks/{id} | Delete a Network Block
[**DeleteRule**](DefaultAPI.md#DeleteRule) | **Delete** /rules/{id} | Delete a Rule
[**DeleteRuleTuningExpression**](DefaultAPI.md#DeleteRuleTuningExpression) | **Delete** /rule-tuning-expressions/{id} | Delete a Rule Tuning Expression
[**DeleteSuppressList**](DefaultAPI.md#DeleteSuppressList) | **Delete** /suppress-lists/{id} | Delete a Suppress List
[**DeleteSuppressListItem**](DefaultAPI.md#DeleteSuppressListItem) | **Delete** /suppress-list-items/{id} | Delete a Suppress List Item
[**DeleteTagSchema**](DefaultAPI.md#DeleteTagSchema) | **Delete** /tag-schemas/{key} | Delete a Tag Schema
[**DeleteThreatIntelIndicator**](DefaultAPI.md#DeleteThreatIntelIndicator) | **Delete** /threat-intel-indicators/{id} | Delete a Threat Intel Indicator
[**ExecuteAutomation**](DefaultAPI.md#ExecuteAutomation) | **Post** /automations/execute | Execute a list of automations to a list of resources
[**GetAllEntities**](DefaultAPI.md#GetAllEntities) | **Get** /entities/all | Get the list of all Entities
[**GetAllInsights**](DefaultAPI.md#GetAllInsights) | **Get** /insights/all | Get the list of all Insights
[**GetAllSignals**](DefaultAPI.md#GetAllSignals) | **Get** /signals/all | Get the list of all Signals
[**GetAllThreatIntelIndicators**](DefaultAPI.md#GetAllThreatIntelIndicators) | **Get** /threat-intel-indicators/all | Get a list of all Threat Intel Indicators
[**GetAutomation**](DefaultAPI.md#GetAutomation) | **Get** /automations/{id} | Get an Automation
[**GetAutomations**](DefaultAPI.md#GetAutomations) | **Get** /automations | Get the list of Automations
[**GetContextAction**](DefaultAPI.md#GetContextAction) | **Get** /context-actions/{id} | Get a Context Action
[**GetContextActions**](DefaultAPI.md#GetContextActions) | **Get** /context-actions | Get the list of Context Actions
[**GetCustomEntityType**](DefaultAPI.md#GetCustomEntityType) | **Get** /custom-entity-types/{id} | Get a Custom Entity Type
[**GetCustomEntityTypes**](DefaultAPI.md#GetCustomEntityTypes) | **Get** /custom-entity-types | Get the list of Custom Entity Types
[**GetCustomInsight**](DefaultAPI.md#GetCustomInsight) | **Get** /custom-insights/{id} | Get a Custom Insight
[**GetCustomInsights**](DefaultAPI.md#GetCustomInsights) | **Get** /custom-insights | Get the list of Custom Insights
[**GetCustomMatchListColumn**](DefaultAPI.md#GetCustomMatchListColumn) | **Get** /custom-match-list-columns/{id} | Get a Custom Match List Column
[**GetCustomMatchListColumns**](DefaultAPI.md#GetCustomMatchListColumns) | **Get** /custom-match-list-columns | Get the list of Custom Match List Columns
[**GetCustomerSourcedLookupTable**](DefaultAPI.md#GetCustomerSourcedLookupTable) | **Get** /entity-lookup-tables/customer-sourced/{id} | Get a customer sourced entity lookup table
[**GetCustomerSourcedLookupTables**](DefaultAPI.md#GetCustomerSourcedLookupTables) | **Get** /entity-lookup-tables/customer-sourced | Get the list of customer sourced entity lookup tables
[**GetEntities**](DefaultAPI.md#GetEntities) | **Get** /entities | Get the list of Entities
[**GetEntity**](DefaultAPI.md#GetEntity) | **Get** /entities/{id} | Get an Entity
[**GetEntityCriticalityConfig**](DefaultAPI.md#GetEntityCriticalityConfig) | **Get** /entity-criticality-configs/{id} | Get an Entity Criticality Configuration
[**GetEntityCriticalityConfigs**](DefaultAPI.md#GetEntityCriticalityConfigs) | **Get** /entity-criticality-configs | Get the list of Entity Criticality Configurations
[**GetEntityDomainConfiguration**](DefaultAPI.md#GetEntityDomainConfiguration) | **Get** /entity-normalization/domain-configuration | Get the Entity Domain Configuration
[**GetEntityEnrichments**](DefaultAPI.md#GetEntityEnrichments) | **Get** /entities/{id}/enrichments | Get an Entity&#39;s enrichments
[**GetEntityExecutedAutomations**](DefaultAPI.md#GetEntityExecutedAutomations) | **Get** /entities/{id}/executed-automations | Get entity executed automations
[**GetEntityGroup**](DefaultAPI.md#GetEntityGroup) | **Get** /entity-group-configurations/{id} | Get an Entity Group Configuration
[**GetEntityGroups**](DefaultAPI.md#GetEntityGroups) | **Get** /entity-group-configurations | Get the list of Entity Group Configurations for a given query
[**GetEntityNotes**](DefaultAPI.md#GetEntityNotes) | **Get** /entities/{entityId}/notes | Get the list of entity notes
[**GetInsight**](DefaultAPI.md#GetInsight) | **Get** /insights/{id} | Get an Insight
[**GetInsightComments**](DefaultAPI.md#GetInsightComments) | **Get** /insights/{id}/comments | Get an Insight&#39;s comments
[**GetInsightCounts**](DefaultAPI.md#GetInsightCounts) | **Get** /insights/counts | Get the count of Insights over time
[**GetInsightEnrichments**](DefaultAPI.md#GetInsightEnrichments) | **Get** /insights/{id}/enrichments | Get an Insights&#39;s enrichments
[**GetInsightExecutedAutomations**](DefaultAPI.md#GetInsightExecutedAutomations) | **Get** /insights/{id}/executed-automations | Get insight executed automations
[**GetInsightHistory**](DefaultAPI.md#GetInsightHistory) | **Get** /insights/{id}/history | Get an Insight&#39;s history
[**GetInsightRelatedEntities**](DefaultAPI.md#GetInsightRelatedEntities) | **Get** /insights/{id}/involved-entities | Get an Insight&#39;s list of involved entities
[**GetInsightResolution**](DefaultAPI.md#GetInsightResolution) | **Get** /insight-resolutions/{id} | Get a Insight Resolution
[**GetInsightResolutions**](DefaultAPI.md#GetInsightResolutions) | **Get** /insight-resolutions | Get the list of Insight Resolutions
[**GetInsightStatus**](DefaultAPI.md#GetInsightStatus) | **Get** /insight-status/{id} | Get an Insight Status
[**GetInsightStatuses**](DefaultAPI.md#GetInsightStatuses) | **Get** /insight-status | Get the list of Insight Statuses
[**GetInsights**](DefaultAPI.md#GetInsights) | **Get** /insights | Get the list of Insights for a given query
[**GetInsightsConfiguration**](DefaultAPI.md#GetInsightsConfiguration) | **Get** /insights-configuration | Get Insight Configuration
[**GetLogMapping**](DefaultAPI.md#GetLogMapping) | **Get** /log-mappings/{id} | Get a Log Mapping
[**GetLogMappingVendorsAndProducts**](DefaultAPI.md#GetLogMappingVendorsAndProducts) | **Get** /vendors-and-products | Get a List of Log Mapping Vendors and Products
[**GetLogMappings**](DefaultAPI.md#GetLogMappings) | **Get** /log-mappings | Get the list of Log Mappings
[**GetMatchList**](DefaultAPI.md#GetMatchList) | **Get** /match-lists/{id} | Get a Match List
[**GetMatchListItem**](DefaultAPI.md#GetMatchListItem) | **Get** /match-list-items/{id} | Get a Match List Item
[**GetMatchListItems**](DefaultAPI.md#GetMatchListItems) | **Get** /match-list-items | Get a list of Match List Items
[**GetMatchLists**](DefaultAPI.md#GetMatchLists) | **Get** /match-lists | Get the list of Match Lists
[**GetNetworkBlock**](DefaultAPI.md#GetNetworkBlock) | **Get** /network-blocks/{id} | Get a Network Block
[**GetNetworkBlocks**](DefaultAPI.md#GetNetworkBlocks) | **Get** /network-blocks | Get the list of Network Blocks
[**GetRecordCounts**](DefaultAPI.md#GetRecordCounts) | **Get** /records/counts | Get the count of Records over time
[**GetRelatedEntitiesById**](DefaultAPI.md#GetRelatedEntitiesById) | **Get** /entities/{id}/related-entities | Get the list of Related Entities
[**GetRule**](DefaultAPI.md#GetRule) | **Get** /rules/{id} | Get a Rule
[**GetRuleTuningExpression**](DefaultAPI.md#GetRuleTuningExpression) | **Get** /rule-tuning-expressions/{id} | Get a Rule Tuning Expression
[**GetRuleTuningExpressions**](DefaultAPI.md#GetRuleTuningExpressions) | **Get** /rule-tuning-expressions | Get the list of Rule Tuning Expressions
[**GetRules**](DefaultAPI.md#GetRules) | **Get** /rules | Get the list of Rules for a given query
[**GetSignal**](DefaultAPI.md#GetSignal) | **Get** /signals/{id} | Get a Signal
[**GetSignalCounts**](DefaultAPI.md#GetSignalCounts) | **Get** /signals/counts | Get the count of Signals over time
[**GetSignalEnrichments**](DefaultAPI.md#GetSignalEnrichments) | **Get** /signals/{id}/enrichments | Get a Signal&#39;s enrichments
[**GetSignals**](DefaultAPI.md#GetSignals) | **Get** /signals | Get the list of Signals for a given query
[**GetSuppressList**](DefaultAPI.md#GetSuppressList) | **Get** /suppress-lists/{id} | Get a Suppress List
[**GetSuppressListItem**](DefaultAPI.md#GetSuppressListItem) | **Get** /suppress-list-items/{id} | Get a Suppress List Item
[**GetSuppressListItems**](DefaultAPI.md#GetSuppressListItems) | **Get** /suppress-list-items | Get a list of Suppress List Items
[**GetSuppressLists**](DefaultAPI.md#GetSuppressLists) | **Get** /suppress-lists | Get the list of Suppress Lists
[**GetTagSchema**](DefaultAPI.md#GetTagSchema) | **Get** /tag-schemas/{key} | Get a Tag Schemas
[**GetTagSchemas**](DefaultAPI.md#GetTagSchemas) | **Get** /tag-schemas | Get the list of Tag Schemas
[**GetThreatIntelIndicator**](DefaultAPI.md#GetThreatIntelIndicator) | **Get** /threat-intel-indicators/{id} | Get a Threat Intel Indicator
[**GetThreatIntelIndicators**](DefaultAPI.md#GetThreatIntelIndicators) | **Get** /threat-intel-indicators | Get a list of Threat Intel Indicators
[**GetThreatIntelSource**](DefaultAPI.md#GetThreatIntelSource) | **Get** /threat-intel-sources/{id} | Get a Threat Intel Source
[**GetThreatIntelligenceSources**](DefaultAPI.md#GetThreatIntelligenceSources) | **Get** /threat-intel-sources | Get the list of Threat Intel Sources
[**LookupNetworkBlocksByIp**](DefaultAPI.md#LookupNetworkBlocksByIp) | **Get** /network-blocks/ip-lookup | Lookup Network Blocks that match a specific IP address
[**OverrideAggregationRule**](DefaultAPI.md#OverrideAggregationRule) | **Put** /rules/aggregation/{id}/override | Override a Aggregation Rule
[**OverrideChainRule**](DefaultAPI.md#OverrideChainRule) | **Put** /rules/chain/{id}/override | Override a Chain Rule
[**OverrideFirstSeenRule**](DefaultAPI.md#OverrideFirstSeenRule) | **Put** /rules/first-seen/{id}/override | Override a First Seen Rule
[**OverrideMatchRule**](DefaultAPI.md#OverrideMatchRule) | **Put** /rules/match/{id}/override | Override a Legacy Match Rule
[**OverrideOutlierRule**](DefaultAPI.md#OverrideOutlierRule) | **Put** /rules/outlier/{id}/override | Override a Outlier Rule
[**OverrideTemplatedMatchRule**](DefaultAPI.md#OverrideTemplatedMatchRule) | **Put** /rules/templated/{id}/override | Override a Match Rule
[**OverrideThresholdRule**](DefaultAPI.md#OverrideThresholdRule) | **Put** /rules/threshold/{id}/override | Override a Threshold Rule
[**RemoveEntityTags**](DefaultAPI.md#RemoveEntityTags) | **Delete** /entities/{entityId}/tags | Remove tags from an Entity
[**RemoveInsightAssignee**](DefaultAPI.md#RemoveInsightAssignee) | **Delete** /insights/{id}/assignee | Remove the assignee from an Insight
[**RemoveSignalsFromInsight**](DefaultAPI.md#RemoveSignalsFromInsight) | **Delete** /insights/{insightId}/signals | Remove existing Signals from an Insight
[**RemoveTagFromInsight**](DefaultAPI.md#RemoveTagFromInsight) | **Delete** /insights/{id}/tags/{tagName} | Remove a tag from an Insight
[**SaveEntityEnrichment**](DefaultAPI.md#SaveEntityEnrichment) | **Put** /entities/{id}/enrichments/{enrichmentType} | Create or update an Enrichment on an Entity
[**SaveInsightEnrichment**](DefaultAPI.md#SaveInsightEnrichment) | **Put** /insights/{id}/enrichments/{enrichmentType} | Create or update an Enrichment on an Insight
[**SaveSignalEnrichment**](DefaultAPI.md#SaveSignalEnrichment) | **Put** /signals/{id}/enrichments/{enrichmentType} | Create or update an Enrichment on a Signal
[**UpdateAggregationRule**](DefaultAPI.md#UpdateAggregationRule) | **Put** /rules/aggregation/{id} | Update a Aggregation Rule
[**UpdateAutomation**](DefaultAPI.md#UpdateAutomation) | **Put** /automations/{id} | Update an Automation
[**UpdateChainRule**](DefaultAPI.md#UpdateChainRule) | **Put** /rules/chain/{id} | Update a Chain Rule
[**UpdateContextAction**](DefaultAPI.md#UpdateContextAction) | **Put** /context-actions/{id} | Update a Context Action
[**UpdateCustomEntityType**](DefaultAPI.md#UpdateCustomEntityType) | **Put** /custom-entity-types/{id} | Update a Custom Entity Type
[**UpdateCustomInsight**](DefaultAPI.md#UpdateCustomInsight) | **Put** /custom-insights/{id} | Update a Custom Insight
[**UpdateCustomMatchListColumn**](DefaultAPI.md#UpdateCustomMatchListColumn) | **Put** /custom-match-list-columns/{id} | Update a Custom Match List Column
[**UpdateCustomerSourcedEntityLookupTable**](DefaultAPI.md#UpdateCustomerSourcedEntityLookupTable) | **Put** /entity-lookup-tables/customer-sourced/{id} | Update a customer sourced entity lookup table
[**UpdateEntityCriticality**](DefaultAPI.md#UpdateEntityCriticality) | **Put** /entities/{entityId}/criticality | Update an Entity&#39;s criticality
[**UpdateEntityCriticalityConfig**](DefaultAPI.md#UpdateEntityCriticalityConfig) | **Put** /entity-criticality-configs/{id} | Update an Entity Criticality Configuration
[**UpdateEntityDomainConfiguration**](DefaultAPI.md#UpdateEntityDomainConfiguration) | **Put** /entity-normalization/domain-configuration | Update the Entity Domain Configuration
[**UpdateEntityNote**](DefaultAPI.md#UpdateEntityNote) | **Put** /entities/{entityId}/notes/{entityNoteId} | Update an existing Entity Note
[**UpdateEntitySuppressed**](DefaultAPI.md#UpdateEntitySuppressed) | **Put** /entities/{entityId}/suppressed | Suppress or un-suppress an Entity&#39;
[**UpdateEntityTags**](DefaultAPI.md#UpdateEntityTags) | **Put** /entities/{entityId}/tags | Update an Entity&#39;s tags, replacing any existing tags
[**UpdateEntityValue**](DefaultAPI.md#UpdateEntityValue) | **Put** /entity-group-configurations/entity_value/{id} | Update an Entity Value Group
[**UpdateFirstSeenRule**](DefaultAPI.md#UpdateFirstSeenRule) | **Put** /rules/first-seen/{id} | Update a First Seen Rule
[**UpdateInsightAssignee**](DefaultAPI.md#UpdateInsightAssignee) | **Put** /insights/{id}/assignee | Update the assignee of an Insight
[**UpdateInsightConfiguration**](DefaultAPI.md#UpdateInsightConfiguration) | **Put** /insights-configuration | Update Insight Configuration
[**UpdateInsightResolution**](DefaultAPI.md#UpdateInsightResolution) | **Put** /insight-resolutions/{id} | Update a Insight Resolution
[**UpdateInsightSeverity**](DefaultAPI.md#UpdateInsightSeverity) | **Put** /insights/{id}/severity | Update the severity of an Insight
[**UpdateInsightStatus**](DefaultAPI.md#UpdateInsightStatus) | **Put** /insights/{id}/status | Update the status of an Insight
[**UpdateInsightStatusOption**](DefaultAPI.md#UpdateInsightStatusOption) | **Put** /insight-status/{id} | Update an Insight Status
[**UpdateInventory**](DefaultAPI.md#UpdateInventory) | **Put** /entity-group-configurations/inventory/{id} | Update an Inventory Group
[**UpdateLogMapping**](DefaultAPI.md#UpdateLogMapping) | **Put** /log-mappings/{id} | Update a Log Mapping
[**UpdateLogMappingEnabled**](DefaultAPI.md#UpdateLogMappingEnabled) | **Put** /log-mappings/{id}/enabled | Update enabled field of a Log Mapping
[**UpdateMatchList**](DefaultAPI.md#UpdateMatchList) | **Put** /match-lists/{id} | Update a Match List
[**UpdateMatchListItem**](DefaultAPI.md#UpdateMatchListItem) | **Put** /match-list-items/{id} | Update a Match List Item
[**UpdateMatchRule**](DefaultAPI.md#UpdateMatchRule) | **Put** /rules/match/{id} | Update a Legacy Match Rule
[**UpdateNetworkBlock**](DefaultAPI.md#UpdateNetworkBlock) | **Put** /network-blocks/{id} | Update a Network Block
[**UpdateRuleEnabled**](DefaultAPI.md#UpdateRuleEnabled) | **Put** /rules/{id}/enabled | Enable or disable a Rule
[**UpdateRuleTuningExpression**](DefaultAPI.md#UpdateRuleTuningExpression) | **Put** /rule-tuning-expressions/{id} | Update a Rule Tuning Expression
[**UpdateSuppressList**](DefaultAPI.md#UpdateSuppressList) | **Put** /suppress-lists/{id} | Update a Suppress List
[**UpdateSuppressListItem**](DefaultAPI.md#UpdateSuppressListItem) | **Put** /suppress-list-items/{id} | Update a Suppress List Item
[**UpdateTagSchema**](DefaultAPI.md#UpdateTagSchema) | **Put** /tag-schemas | Update a Tag Schema
[**UpdateTemplatedMatchRule**](DefaultAPI.md#UpdateTemplatedMatchRule) | **Put** /rules/templated/{id} | Update a Match Rule
[**UpdateThreatIntelIndicator**](DefaultAPI.md#UpdateThreatIntelIndicator) | **Put** /threat-intel-indicators/{id} | Update a Threat Intel Indicator
[**UpdateThresholdRule**](DefaultAPI.md#UpdateThresholdRule) | **Put** /rules/threshold/{id} | Update a Threshold Rule



## AddEntityNote

> AddEntityNoteResponse AddEntityNote(ctx, entityId).AddEntityNoteRequestBody(addEntityNoteRequestBody).Execute()

Add a new Entity Note on an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    addEntityNoteRequestBody := *openapiclient.NewAddEntityNoteRequestBody("Note_example") // AddEntityNoteRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddEntityNote(context.Background(), entityId).AddEntityNoteRequestBody(addEntityNoteRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddEntityNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntityNote`: AddEntityNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddEntityNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntityNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addEntityNoteRequestBody** | [**AddEntityNoteRequestBody**](AddEntityNoteRequestBody.md) |  | 

### Return type

[**AddEntityNoteResponse**](AddEntityNoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddEntityTags

> AddEntityTagsResponse AddEntityTags(ctx, entityId).AddEntityTagsRequestBody(addEntityTagsRequestBody).Execute()

Add tags to an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    addEntityTagsRequestBody := *openapiclient.NewAddEntityTagsRequestBody([]string{"Tags_example"}) // AddEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddEntityTags(context.Background(), entityId).AddEntityTagsRequestBody(addEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntityTags`: AddEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddEntityTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addEntityTagsRequestBody** | [**AddEntityTagsRequestBody**](AddEntityTagsRequestBody.md) |  | 

### Return type

[**AddEntityTagsResponse**](AddEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIndicatorToThreatIntelSource

> AddIndicatorToThreatIntelSourceResponse AddIndicatorToThreatIntelSource(ctx, threatIntelSourceId).AddIndicatorToThreatIntelSourceRequestBody(addIndicatorToThreatIntelSourceRequestBody).Execute()

Add Indicators to a Threat Intel Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    threatIntelSourceId := "threatIntelSourceId_example" // string | 
    addIndicatorToThreatIntelSourceRequestBody := *openapiclient.NewAddIndicatorToThreatIntelSourceRequestBody([]openapiclient.AddItemsToSuppressListRequestBodyItemsInner{*openapiclient.NewAddItemsToSuppressListRequestBodyItemsInner("Value_example", false, "Description_example")}) // AddIndicatorToThreatIntelSourceRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddIndicatorToThreatIntelSource(context.Background(), threatIntelSourceId).AddIndicatorToThreatIntelSourceRequestBody(addIndicatorToThreatIntelSourceRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddIndicatorToThreatIntelSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIndicatorToThreatIntelSource`: AddIndicatorToThreatIntelSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddIndicatorToThreatIntelSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatIntelSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIndicatorToThreatIntelSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addIndicatorToThreatIntelSourceRequestBody** | [**AddIndicatorToThreatIntelSourceRequestBody**](AddIndicatorToThreatIntelSourceRequestBody.md) |  | 

### Return type

[**AddIndicatorToThreatIntelSourceResponse**](AddIndicatorToThreatIntelSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddInsightComment

> AddInsightCommentResponse AddInsightComment(ctx, id).AddInsightCommentRequestBody(addInsightCommentRequestBody).Execute()

Add a new comment on an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    addInsightCommentRequestBody := *openapiclient.NewAddInsightCommentRequestBody("Body_example") // AddInsightCommentRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddInsightComment(context.Background(), id).AddInsightCommentRequestBody(addInsightCommentRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddInsightComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInsightComment`: AddInsightCommentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddInsightComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInsightCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addInsightCommentRequestBody** | [**AddInsightCommentRequestBody**](AddInsightCommentRequestBody.md) |  | 

### Return type

[**AddInsightCommentResponse**](AddInsightCommentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddItemsToMatchList

> AddItemsToMatchListResponse AddItemsToMatchList(ctx, id).AddItemsToMatchListRequestBody(addItemsToMatchListRequestBody).Execute()

Add Match List Items to a Match List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    addItemsToMatchListRequestBody := *openapiclient.NewAddItemsToMatchListRequestBody([]openapiclient.AddItemsToSuppressListRequestBodyItemsInner{*openapiclient.NewAddItemsToSuppressListRequestBodyItemsInner("Value_example", false, "Description_example")}) // AddItemsToMatchListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddItemsToMatchList(context.Background(), id).AddItemsToMatchListRequestBody(addItemsToMatchListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddItemsToMatchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemsToMatchList`: AddItemsToMatchListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddItemsToMatchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddItemsToMatchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addItemsToMatchListRequestBody** | [**AddItemsToMatchListRequestBody**](AddItemsToMatchListRequestBody.md) |  | 

### Return type

[**AddItemsToMatchListResponse**](AddItemsToMatchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddItemsToSuppressList

> AddItemsToSuppressListResponse AddItemsToSuppressList(ctx, id).AddItemsToSuppressListRequestBody(addItemsToSuppressListRequestBody).Execute()

Add Suppress List Items to a Suppress List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    addItemsToSuppressListRequestBody := *openapiclient.NewAddItemsToSuppressListRequestBody([]openapiclient.AddItemsToSuppressListRequestBodyItemsInner{*openapiclient.NewAddItemsToSuppressListRequestBodyItemsInner("Value_example", false, "Description_example")}) // AddItemsToSuppressListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddItemsToSuppressList(context.Background(), id).AddItemsToSuppressListRequestBody(addItemsToSuppressListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddItemsToSuppressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemsToSuppressList`: AddItemsToSuppressListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddItemsToSuppressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddItemsToSuppressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addItemsToSuppressListRequestBody** | [**AddItemsToSuppressListRequestBody**](AddItemsToSuppressListRequestBody.md) |  | 

### Return type

[**AddItemsToSuppressListResponse**](AddItemsToSuppressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSignalsToInsight

> AddSignalsToInsightResponse AddSignalsToInsight(ctx, insightId).AddSignalsToInsightRequestBody(addSignalsToInsightRequestBody).Execute()

Add Signals to an existing Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    insightId := "insightId_example" // string | 
    addSignalsToInsightRequestBody := *openapiclient.NewAddSignalsToInsightRequestBody([]string{"SignalIds_example"}) // AddSignalsToInsightRequestBody | A list of Signal IDs to be added to the Insight

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddSignalsToInsight(context.Background(), insightId).AddSignalsToInsightRequestBody(addSignalsToInsightRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddSignalsToInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSignalsToInsight`: AddSignalsToInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddSignalsToInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSignalsToInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSignalsToInsightRequestBody** | [**AddSignalsToInsightRequestBody**](AddSignalsToInsightRequestBody.md) | A list of Signal IDs to be added to the Insight | 

### Return type

[**AddSignalsToInsightResponse**](AddSignalsToInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTagToInsight

> AddTagToInsightResponse AddTagToInsight(ctx, id).AddTagToInsightRequestBody(addTagToInsightRequestBody).Execute()

Add a tag to an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    addTagToInsightRequestBody := *openapiclient.NewAddTagToInsightRequestBody("TagName_example") // AddTagToInsightRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.AddTagToInsight(context.Background(), id).AddTagToInsightRequestBody(addTagToInsightRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddTagToInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTagToInsight`: AddTagToInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddTagToInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagToInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTagToInsightRequestBody** | [**AddTagToInsightRequestBody**](AddTagToInsightRequestBody.md) |  | 

### Return type

[**AddTagToInsightResponse**](AddTagToInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAddEntityTags

> BulkAddEntityTagsResponse BulkAddEntityTags(ctx).BulkAddEntityTagsRequestBody(bulkAddEntityTagsRequestBody).Execute()

Bulk add tags to entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkAddEntityTagsRequestBody := *openapiclient.NewBulkAddEntityTagsRequestBody([]string{"EntityIds_example"}, []string{"Tags_example"}) // BulkAddEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkAddEntityTags(context.Background()).BulkAddEntityTagsRequestBody(bulkAddEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkAddEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkAddEntityTags`: BulkAddEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkAddEntityTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAddEntityTagsRequestBody** | [**BulkAddEntityTagsRequestBody**](BulkAddEntityTagsRequestBody.md) |  | 

### Return type

[**BulkAddEntityTagsResponse**](BulkAddEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteMatchListItems

> BulkDeleteMatchListItemsResponse BulkDeleteMatchListItems(ctx).BulkDeleteMatchListItemsRequestBody(bulkDeleteMatchListItemsRequestBody).Execute()

Bulk delete Match List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkDeleteMatchListItemsRequestBody := *openapiclient.NewBulkDeleteMatchListItemsRequestBody([]string{"Ids_example"}) // BulkDeleteMatchListItemsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkDeleteMatchListItems(context.Background()).BulkDeleteMatchListItemsRequestBody(bulkDeleteMatchListItemsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkDeleteMatchListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteMatchListItems`: BulkDeleteMatchListItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkDeleteMatchListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteMatchListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkDeleteMatchListItemsRequestBody** | [**BulkDeleteMatchListItemsRequestBody**](BulkDeleteMatchListItemsRequestBody.md) |  | 

### Return type

[**BulkDeleteMatchListItemsResponse**](BulkDeleteMatchListItemsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRemoveEntityTags

> BulkRemoveEntityTagsResponse BulkRemoveEntityTags(ctx).BulkRemoveEntityTagsRequestBody(bulkRemoveEntityTagsRequestBody).Execute()

Bulk remove tags on entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkRemoveEntityTagsRequestBody := *openapiclient.NewBulkRemoveEntityTagsRequestBody([]string{"EntityIds_example"}, []string{"Tags_example"}) // BulkRemoveEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkRemoveEntityTags(context.Background()).BulkRemoveEntityTagsRequestBody(bulkRemoveEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkRemoveEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkRemoveEntityTags`: BulkRemoveEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkRemoveEntityTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkRemoveEntityTagsRequestBody** | [**BulkRemoveEntityTagsRequestBody**](BulkRemoveEntityTagsRequestBody.md) |  | 

### Return type

[**BulkRemoveEntityTagsResponse**](BulkRemoveEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateEntityCriticality

> BulkUpdateEntityCriticalityResponse BulkUpdateEntityCriticality(ctx).BulkUpdateEntityCriticalityRequestBody(bulkUpdateEntityCriticalityRequestBody).Execute()

Bulk updates criticality entity field



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkUpdateEntityCriticalityRequestBody := *openapiclient.NewBulkUpdateEntityCriticalityRequestBody([]string{"EntityIds_example"}) // BulkUpdateEntityCriticalityRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkUpdateEntityCriticality(context.Background()).BulkUpdateEntityCriticalityRequestBody(bulkUpdateEntityCriticalityRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateEntityCriticality``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateEntityCriticality`: BulkUpdateEntityCriticalityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpdateEntityCriticality`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateEntityCriticalityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateEntityCriticalityRequestBody** | [**BulkUpdateEntityCriticalityRequestBody**](BulkUpdateEntityCriticalityRequestBody.md) |  | 

### Return type

[**BulkUpdateEntityCriticalityResponse**](BulkUpdateEntityCriticalityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateEntityTags

> BulkUpdateEntityTagsResponse BulkUpdateEntityTags(ctx).BulkUpdateEntityTagsRequestBody(bulkUpdateEntityTagsRequestBody).Execute()

Bulk update tags on entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkUpdateEntityTagsRequestBody := *openapiclient.NewBulkUpdateEntityTagsRequestBody([]string{"EntityIds_example"}, []string{"Tags_example"}) // BulkUpdateEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkUpdateEntityTags(context.Background()).BulkUpdateEntityTagsRequestBody(bulkUpdateEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateEntityTags`: BulkUpdateEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpdateEntityTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateEntityTagsRequestBody** | [**BulkUpdateEntityTagsRequestBody**](BulkUpdateEntityTagsRequestBody.md) |  | 

### Return type

[**BulkUpdateEntityTagsResponse**](BulkUpdateEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateEntityWhitelist

> BulkUpdateEntityWhitelistResponse BulkUpdateEntityWhitelist(ctx).BulkUpdateEntityWhitelistRequestBody(bulkUpdateEntityWhitelistRequestBody).Execute()

Bulk updates suppressed entity field



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkUpdateEntityWhitelistRequestBody := *openapiclient.NewBulkUpdateEntityWhitelistRequestBody(false, []string{"EntityIds_example"}) // BulkUpdateEntityWhitelistRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkUpdateEntityWhitelist(context.Background()).BulkUpdateEntityWhitelistRequestBody(bulkUpdateEntityWhitelistRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpdateEntityWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateEntityWhitelist`: BulkUpdateEntityWhitelistResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpdateEntityWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateEntityWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateEntityWhitelistRequestBody** | [**BulkUpdateEntityWhitelistRequestBody**](BulkUpdateEntityWhitelistRequestBody.md) |  | 

### Return type

[**BulkUpdateEntityWhitelistResponse**](BulkUpdateEntityWhitelistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpsertNetworkBlocks

> BulkUpsertNetworkBlocksResponse BulkUpsertNetworkBlocks(ctx).BulkUpsertNetworkBlocksRequestBody(bulkUpsertNetworkBlocksRequestBody).Execute()

Add or update multiple Network Blocks in one request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    bulkUpsertNetworkBlocksRequestBody := *openapiclient.NewBulkUpsertNetworkBlocksRequestBody([]openapiclient.CreateNetworkBlockRequestBodyFields{*openapiclient.NewCreateNetworkBlockRequestBodyFields("AddressBlock_example")}) // BulkUpsertNetworkBlocksRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.BulkUpsertNetworkBlocks(context.Background()).BulkUpsertNetworkBlocksRequestBody(bulkUpsertNetworkBlocksRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BulkUpsertNetworkBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpsertNetworkBlocks`: BulkUpsertNetworkBlocksResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BulkUpsertNetworkBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpsertNetworkBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpsertNetworkBlocksRequestBody** | [**BulkUpsertNetworkBlocksRequestBody**](BulkUpsertNetworkBlocksRequestBody.md) |  | 

### Return type

[**BulkUpsertNetworkBlocksResponse**](BulkUpsertNetworkBlocksResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAggregationRule

> CreateAggregationRuleResponse CreateAggregationRule(ctx).CreateAggregationRuleRequestBody(createAggregationRuleRequestBody).Execute()

Create a Aggregation Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createAggregationRuleRequestBody := *openapiclient.NewCreateAggregationRuleRequestBody(*openapiclient.NewCreateAggregationRuleRequestBodyFields(false, "Name_example", []openapiclient.GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner{*openapiclient.NewGetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner("Name_example", "Function_example", []string{"Arguments_example"})}, "DescriptionExpression_example", false, []string{"GroupByFields_example"}, "MatchExpression_example", "NameExpression_example", *openapiclient.NewGetRulesResponseDataObjectsInnerOneOf2ScoreMapping("Type_example"), "Stream_example", "TriggerExpression_example", "WindowSize_example")) // CreateAggregationRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateAggregationRule(context.Background()).CreateAggregationRuleRequestBody(createAggregationRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAggregationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAggregationRule`: CreateAggregationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAggregationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAggregationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAggregationRuleRequestBody** | [**CreateAggregationRuleRequestBody**](CreateAggregationRuleRequestBody.md) |  | 

### Return type

[**CreateAggregationRuleResponse**](CreateAggregationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutomation

> CreateAutomationResponse CreateAutomation(ctx).CreateAutomationRequestBody(createAutomationRequestBody).Execute()

Create an Automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createAutomationRequestBody := *openapiclient.NewCreateAutomationRequestBody(*openapiclient.NewCreateAutomationRequestBodyFields("PlaybookId_example", "CseResourceType_example", []string{"ExecutionTypes_example"}, false)) // CreateAutomationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateAutomation(context.Background()).CreateAutomationRequestBody(createAutomationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomation`: CreateAutomationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutomationRequestBody** | [**CreateAutomationRequestBody**](CreateAutomationRequestBody.md) |  | 

### Return type

[**CreateAutomationResponse**](CreateAutomationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChainRule

> CreateChainRuleResponse CreateChainRule(ctx).CreateChainRuleRequestBody(createChainRuleRequestBody).Execute()

Create a Chain Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createChainRuleRequestBody := *openapiclient.NewCreateChainRuleRequestBody(*openapiclient.NewCreateChainRuleRequestBodyFields(false, "Name_example", "Description_example", []openapiclient.CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner{*openapiclient.NewCreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner("Expression_example", int32(123))}, int32(123), "Stream_example", "WindowSize_example")) // CreateChainRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateChainRule(context.Background()).CreateChainRuleRequestBody(createChainRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateChainRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChainRule`: CreateChainRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateChainRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChainRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createChainRuleRequestBody** | [**CreateChainRuleRequestBody**](CreateChainRuleRequestBody.md) |  | 

### Return type

[**CreateChainRuleResponse**](CreateChainRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContextAction

> CreateContextActionResponse CreateContextAction(ctx).CreateContextActionRequestBody(createContextActionRequestBody).Execute()

Create a Context Action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createContextActionRequestBody := *openapiclient.NewCreateContextActionRequestBody("Name_example", []string{"IocTypes_example"}) // CreateContextActionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateContextAction(context.Background()).CreateContextActionRequestBody(createContextActionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateContextAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContextAction`: CreateContextActionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateContextAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContextActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContextActionRequestBody** | [**CreateContextActionRequestBody**](CreateContextActionRequestBody.md) |  | 

### Return type

[**CreateContextActionResponse**](CreateContextActionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomEntityType

> CreateCustomEntityTypeResponse CreateCustomEntityType(ctx).CreateCustomEntityTypeRequestBody(createCustomEntityTypeRequestBody).Execute()

Create a Custom Entity Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createCustomEntityTypeRequestBody := *openapiclient.NewCreateCustomEntityTypeRequestBody(*openapiclient.NewCreateCustomEntityTypeRequestBodyFields("Name_example", []string{"Fields_example"}, "Identifier_example")) // CreateCustomEntityTypeRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateCustomEntityType(context.Background()).CreateCustomEntityTypeRequestBody(createCustomEntityTypeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCustomEntityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomEntityType`: CreateCustomEntityTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCustomEntityType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomEntityTypeRequestBody** | [**CreateCustomEntityTypeRequestBody**](CreateCustomEntityTypeRequestBody.md) |  | 

### Return type

[**CreateCustomEntityTypeResponse**](CreateCustomEntityTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomInsight

> CreateCustomInsightResponse CreateCustomInsight(ctx).CreateCustomInsightRequestBody(createCustomInsightRequestBody).Execute()

Create a Custom Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createCustomInsightRequestBody := *openapiclient.NewCreateCustomInsightRequestBody(*openapiclient.NewCreateCustomInsightRequestBodyFields("Name_example", "Description_example", "Severity_example", false, false, []string{"Tags_example"})) // CreateCustomInsightRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateCustomInsight(context.Background()).CreateCustomInsightRequestBody(createCustomInsightRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCustomInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomInsight`: CreateCustomInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCustomInsight`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomInsightRequestBody** | [**CreateCustomInsightRequestBody**](CreateCustomInsightRequestBody.md) |  | 

### Return type

[**CreateCustomInsightResponse**](CreateCustomInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomMatchListColumn

> CreateCustomMatchListColumnResponse CreateCustomMatchListColumn(ctx).CreateCustomMatchListColumnRequestBody(createCustomMatchListColumnRequestBody).Execute()

Create a Custom Match List Column



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createCustomMatchListColumnRequestBody := *openapiclient.NewCreateCustomMatchListColumnRequestBody(*openapiclient.NewUpdateCustomMatchListColumnRequestBodyFields("Name_example", []string{"Fields_example"})) // CreateCustomMatchListColumnRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateCustomMatchListColumn(context.Background()).CreateCustomMatchListColumnRequestBody(createCustomMatchListColumnRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCustomMatchListColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomMatchListColumn`: CreateCustomMatchListColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCustomMatchListColumn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomMatchListColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomMatchListColumnRequestBody** | [**CreateCustomMatchListColumnRequestBody**](CreateCustomMatchListColumnRequestBody.md) |  | 

### Return type

[**CreateCustomMatchListColumnResponse**](CreateCustomMatchListColumnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerSourcedEntityLookupTable

> CreateCustomerSourcedEntityLookupTableResponse CreateCustomerSourcedEntityLookupTable(ctx).CreateCustomerSourcedEntityLookupTableRequestBody(createCustomerSourcedEntityLookupTableRequestBody).Execute()

Create a customer sourced entity lookup table



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createCustomerSourcedEntityLookupTableRequestBody := *openapiclient.NewCreateCustomerSourcedEntityLookupTableRequestBody(*openapiclient.NewCreateCustomerSourcedEntityLookupTableRequestBodyFields("Type_example", "KeyColumnName_example", "ValueColumnName_example", "TablePath_example")) // CreateCustomerSourcedEntityLookupTableRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateCustomerSourcedEntityLookupTable(context.Background()).CreateCustomerSourcedEntityLookupTableRequestBody(createCustomerSourcedEntityLookupTableRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCustomerSourcedEntityLookupTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerSourcedEntityLookupTable`: CreateCustomerSourcedEntityLookupTableResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCustomerSourcedEntityLookupTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerSourcedEntityLookupTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerSourcedEntityLookupTableRequestBody** | [**CreateCustomerSourcedEntityLookupTableRequestBody**](CreateCustomerSourcedEntityLookupTableRequestBody.md) |  | 

### Return type

[**CreateCustomerSourcedEntityLookupTableResponse**](CreateCustomerSourcedEntityLookupTableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntityCriticalityConfig

> CreateEntityCriticalityConfigResponse CreateEntityCriticalityConfig(ctx).CreateEntityCriticalityConfigRequestBody(createEntityCriticalityConfigRequestBody).Execute()

Create an Entity Criticality Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createEntityCriticalityConfigRequestBody := *openapiclient.NewCreateEntityCriticalityConfigRequestBody(*openapiclient.NewCreateEntityCriticalityConfigRequestBodyFields("Name_example", "SeverityExpression_example")) // CreateEntityCriticalityConfigRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateEntityCriticalityConfig(context.Background()).CreateEntityCriticalityConfigRequestBody(createEntityCriticalityConfigRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEntityCriticalityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntityCriticalityConfig`: CreateEntityCriticalityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEntityCriticalityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityCriticalityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntityCriticalityConfigRequestBody** | [**CreateEntityCriticalityConfigRequestBody**](CreateEntityCriticalityConfigRequestBody.md) |  | 

### Return type

[**CreateEntityCriticalityConfigResponse**](CreateEntityCriticalityConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntityValue

> CreateEntityValueResponse CreateEntityValue(ctx).CreateEntityValueRequestBody(createEntityValueRequestBody).Execute()

Create an Entity Value Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createEntityValueRequestBody := *openapiclient.NewCreateEntityValueRequestBody(*openapiclient.NewCreateEntityValueRequestBodyFields("Name_example", "ConfigurationType_example")) // CreateEntityValueRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateEntityValue(context.Background()).CreateEntityValueRequestBody(createEntityValueRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEntityValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntityValue`: CreateEntityValueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEntityValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntityValueRequestBody** | [**CreateEntityValueRequestBody**](CreateEntityValueRequestBody.md) |  | 

### Return type

[**CreateEntityValueResponse**](CreateEntityValueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirstSeenRule

> CreateFirstSeenRuleResponse CreateFirstSeenRule(ctx).CreateFirstSeenRuleRequestBody(createFirstSeenRuleRequestBody).Execute()

Create a First Seen Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createFirstSeenRuleRequestBody := *openapiclient.NewCreateFirstSeenRuleRequestBody(*openapiclient.NewCreateFirstSeenRuleRequestBodyFields(false, "Name_example", "DescriptionExpression_example", "NameExpression_example", "FilterExpression_example", []string{"GroupByFields_example"}, int32(123), int32(123), "BaselineWindowSize_example", "RetentionWindowSize_example")) // CreateFirstSeenRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateFirstSeenRule(context.Background()).CreateFirstSeenRuleRequestBody(createFirstSeenRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateFirstSeenRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirstSeenRule`: CreateFirstSeenRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateFirstSeenRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirstSeenRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirstSeenRuleRequestBody** | [**CreateFirstSeenRuleRequestBody**](CreateFirstSeenRuleRequestBody.md) |  | 

### Return type

[**CreateFirstSeenRuleResponse**](CreateFirstSeenRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInsightFromSignals

> CreateInsightFromSignalsResponse CreateInsightFromSignals(ctx).CreateInsightFromSignalsRequestBody(createInsightFromSignalsRequestBody).Execute()

Create Insight from Signals



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createInsightFromSignalsRequestBody := *openapiclient.NewCreateInsightFromSignalsRequestBody([]string{"SignalIds_example"}) // CreateInsightFromSignalsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateInsightFromSignals(context.Background()).CreateInsightFromSignalsRequestBody(createInsightFromSignalsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateInsightFromSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInsightFromSignals`: CreateInsightFromSignalsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateInsightFromSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInsightFromSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsightFromSignalsRequestBody** | [**CreateInsightFromSignalsRequestBody**](CreateInsightFromSignalsRequestBody.md) |  | 

### Return type

[**CreateInsightFromSignalsResponse**](CreateInsightFromSignalsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInsightResolution

> CreateInsightResolutionResponse CreateInsightResolution(ctx).CreateInsightResolutionRequestBody(createInsightResolutionRequestBody).Execute()

Create a Insight Resolution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createInsightResolutionRequestBody := *openapiclient.NewCreateInsightResolutionRequestBody() // CreateInsightResolutionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateInsightResolution(context.Background()).CreateInsightResolutionRequestBody(createInsightResolutionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateInsightResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInsightResolution`: CreateInsightResolutionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateInsightResolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInsightResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsightResolutionRequestBody** | [**CreateInsightResolutionRequestBody**](CreateInsightResolutionRequestBody.md) |  | 

### Return type

[**CreateInsightResolutionResponse**](CreateInsightResolutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInsightStatusOption

> CreateInsightStatusOptionResponse CreateInsightStatusOption(ctx).CreateInsightStatusOptionRequestBody(createInsightStatusOptionRequestBody).Execute()

Create an Insight Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createInsightStatusOptionRequestBody := *openapiclient.NewCreateInsightStatusOptionRequestBody() // CreateInsightStatusOptionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateInsightStatusOption(context.Background()).CreateInsightStatusOptionRequestBody(createInsightStatusOptionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateInsightStatusOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInsightStatusOption`: CreateInsightStatusOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateInsightStatusOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInsightStatusOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInsightStatusOptionRequestBody** | [**CreateInsightStatusOptionRequestBody**](CreateInsightStatusOptionRequestBody.md) |  | 

### Return type

[**CreateInsightStatusOptionResponse**](CreateInsightStatusOptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInventory

> CreateInventoryResponse CreateInventory(ctx).CreateInventoryRequestBody(createInventoryRequestBody).Execute()

Create an Inventory Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createInventoryRequestBody := *openapiclient.NewCreateInventoryRequestBody(*openapiclient.NewCreateInventoryRequestBodyFields("Name_example", "ConfigurationType_example", "InventorySource_example")) // CreateInventoryRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateInventory(context.Background()).CreateInventoryRequestBody(createInventoryRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInventory`: CreateInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInventoryRequestBody** | [**CreateInventoryRequestBody**](CreateInventoryRequestBody.md) |  | 

### Return type

[**CreateInventoryResponse**](CreateInventoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLogMapping

> CreateLogMappingResponse CreateLogMapping(ctx).CreateLogMappingRequestBody(createLogMappingRequestBody).Execute()

Create a Log Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createLogMappingRequestBody := *openapiclient.NewCreateLogMappingRequestBody(*openapiclient.NewCreateLogMappingRequestBodyFields("Name_example", []openapiclient.CreateLogMappingRequestBodyFieldsFieldsInner{*openapiclient.NewCreateLogMappingRequestBodyFieldsFieldsInner("Name_example")}, "RecordType_example", "ProductGuid_example", false)) // CreateLogMappingRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateLogMapping(context.Background()).CreateLogMappingRequestBody(createLogMappingRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLogMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogMapping`: CreateLogMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLogMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLogMappingRequestBody** | [**CreateLogMappingRequestBody**](CreateLogMappingRequestBody.md) |  | 

### Return type

[**CreateLogMappingResponse**](CreateLogMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMatchList

> CreateMatchListResponse CreateMatchList(ctx).CreateMatchListRequestBody(createMatchListRequestBody).Execute()

Create a Match List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createMatchListRequestBody := *openapiclient.NewCreateMatchListRequestBody(*openapiclient.NewCreateSuppressListRequestBodyFields("Name_example", "TargetColumn_example")) // CreateMatchListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateMatchList(context.Background()).CreateMatchListRequestBody(createMatchListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateMatchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMatchList`: CreateMatchListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateMatchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMatchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMatchListRequestBody** | [**CreateMatchListRequestBody**](CreateMatchListRequestBody.md) |  | 

### Return type

[**CreateMatchListResponse**](CreateMatchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMatchRule

> CreateMatchRuleResponse CreateMatchRule(ctx).CreateMatchRuleRequestBody(createMatchRuleRequestBody).Execute()

Create a Legacy Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createMatchRuleRequestBody := *openapiclient.NewCreateMatchRuleRequestBody(*openapiclient.NewCreateMatchRuleRequestBodyFields(false, "Name_example", "Description_example", "Expression_example", int32(123), "Stream_example")) // CreateMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateMatchRule(context.Background()).CreateMatchRuleRequestBody(createMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMatchRule`: CreateMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateMatchRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMatchRuleRequestBody** | [**CreateMatchRuleRequestBody**](CreateMatchRuleRequestBody.md) |  | 

### Return type

[**CreateMatchRuleResponse**](CreateMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkBlock

> CreateNetworkBlockResponse CreateNetworkBlock(ctx).CreateNetworkBlockRequestBody(createNetworkBlockRequestBody).Execute()

Create a Network Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createNetworkBlockRequestBody := *openapiclient.NewCreateNetworkBlockRequestBody(*openapiclient.NewCreateNetworkBlockRequestBodyFields("AddressBlock_example")) // CreateNetworkBlockRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateNetworkBlock(context.Background()).CreateNetworkBlockRequestBody(createNetworkBlockRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateNetworkBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkBlock`: CreateNetworkBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateNetworkBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkBlockRequestBody** | [**CreateNetworkBlockRequestBody**](CreateNetworkBlockRequestBody.md) |  | 

### Return type

[**CreateNetworkBlockResponse**](CreateNetworkBlockResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleTuningExpression

> CreateRuleTuningExpressionResponse CreateRuleTuningExpression(ctx).CreateRuleTuningExpressionRequestBody(createRuleTuningExpressionRequestBody).Execute()

Create a Rule Tuning Expression



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createRuleTuningExpressionRequestBody := *openapiclient.NewCreateRuleTuningExpressionRequestBody(*openapiclient.NewCreateRuleTuningExpressionRequestBodyFields("Name_example", "Description_example", "Expression_example", false, false, false, []string{"RuleIds_example"})) // CreateRuleTuningExpressionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateRuleTuningExpression(context.Background()).CreateRuleTuningExpressionRequestBody(createRuleTuningExpressionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateRuleTuningExpression``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRuleTuningExpression`: CreateRuleTuningExpressionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateRuleTuningExpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleTuningExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRuleTuningExpressionRequestBody** | [**CreateRuleTuningExpressionRequestBody**](CreateRuleTuningExpressionRequestBody.md) |  | 

### Return type

[**CreateRuleTuningExpressionResponse**](CreateRuleTuningExpressionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSuppressList

> CreateSuppressListResponse CreateSuppressList(ctx).CreateSuppressListRequestBody(createSuppressListRequestBody).Execute()

Create a Suppress List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createSuppressListRequestBody := *openapiclient.NewCreateSuppressListRequestBody(*openapiclient.NewCreateSuppressListRequestBodyFields("Name_example", "TargetColumn_example")) // CreateSuppressListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateSuppressList(context.Background()).CreateSuppressListRequestBody(createSuppressListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSuppressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSuppressList`: CreateSuppressListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateSuppressList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSuppressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSuppressListRequestBody** | [**CreateSuppressListRequestBody**](CreateSuppressListRequestBody.md) |  | 

### Return type

[**CreateSuppressListResponse**](CreateSuppressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagSchema

> CreateTagSchemaResponse CreateTagSchema(ctx).CreateTagSchemaRequestBody(createTagSchemaRequestBody).Execute()

Create a Tag Schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createTagSchemaRequestBody := *openapiclient.NewCreateTagSchemaRequestBody(*openapiclient.NewUpdateTagSchemaRequestBodyFields("Key_example", "Label_example", false)) // CreateTagSchemaRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateTagSchema(context.Background()).CreateTagSchemaRequestBody(createTagSchemaRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTagSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagSchema`: CreateTagSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTagSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagSchemaRequestBody** | [**CreateTagSchemaRequestBody**](CreateTagSchemaRequestBody.md) |  | 

### Return type

[**CreateTagSchemaResponse**](CreateTagSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplatedMatchRule

> CreateTemplatedMatchRuleResponse CreateTemplatedMatchRule(ctx).CreateTemplatedMatchRuleRequestBody(createTemplatedMatchRuleRequestBody).Execute()

Create a Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createTemplatedMatchRuleRequestBody := *openapiclient.NewCreateTemplatedMatchRuleRequestBody(*openapiclient.NewCreateTemplatedMatchRuleRequestBodyFields(false, "Name_example", "DescriptionExpression_example", "Expression_example", "NameExpression_example", *openapiclient.NewGetRulesResponseDataObjectsInnerOneOf2ScoreMapping("Type_example"), "Stream_example")) // CreateTemplatedMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateTemplatedMatchRule(context.Background()).CreateTemplatedMatchRuleRequestBody(createTemplatedMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTemplatedMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplatedMatchRule`: CreateTemplatedMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTemplatedMatchRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplatedMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTemplatedMatchRuleRequestBody** | [**CreateTemplatedMatchRuleRequestBody**](CreateTemplatedMatchRuleRequestBody.md) |  | 

### Return type

[**CreateTemplatedMatchRuleResponse**](CreateTemplatedMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThreatIntelSource

> CreateThreatIntelSourceResponse CreateThreatIntelSource(ctx).CreateThreatIntelSourceRequestBody(createThreatIntelSourceRequestBody).Execute()

Create a Threat Intel Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createThreatIntelSourceRequestBody := *openapiclient.NewCreateThreatIntelSourceRequestBody(*openapiclient.NewCreateThreatIntelSourceRequestBodyFields("Name_example")) // CreateThreatIntelSourceRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateThreatIntelSource(context.Background()).CreateThreatIntelSourceRequestBody(createThreatIntelSourceRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateThreatIntelSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateThreatIntelSource`: CreateThreatIntelSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateThreatIntelSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThreatIntelSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createThreatIntelSourceRequestBody** | [**CreateThreatIntelSourceRequestBody**](CreateThreatIntelSourceRequestBody.md) |  | 

### Return type

[**CreateThreatIntelSourceResponse**](CreateThreatIntelSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThresholdRule

> CreateThresholdRuleResponse CreateThresholdRule(ctx).CreateThresholdRuleRequestBody(createThresholdRuleRequestBody).Execute()

Create a Threshold Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createThresholdRuleRequestBody := *openapiclient.NewCreateThresholdRuleRequestBody(*openapiclient.NewCreateThresholdRuleRequestBodyFields(false, "Name_example", "Description_example", false, "Expression_example", int32(123), int32(123), "Stream_example", int32(123), "WindowSize_example")) // CreateThresholdRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateThresholdRule(context.Background()).CreateThresholdRuleRequestBody(createThresholdRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateThresholdRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateThresholdRule`: CreateThresholdRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateThresholdRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThresholdRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createThresholdRuleRequestBody** | [**CreateThresholdRuleRequestBody**](CreateThresholdRuleRequestBody.md) |  | 

### Return type

[**CreateThresholdRuleResponse**](CreateThresholdRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutomation

> DeleteAutomationResponse DeleteAutomation(ctx, id).Execute()

Delete an Automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteAutomation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAutomation`: DeleteAutomationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteAutomation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAutomationResponse**](DeleteAutomationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContextAction

> DeleteContextActionResponse DeleteContextAction(ctx, id).Execute()

Delete a Context Action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteContextAction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteContextAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContextAction`: DeleteContextActionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteContextAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteContextActionResponse**](DeleteContextActionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomEntityType

> DeleteCustomEntityTypeResponse DeleteCustomEntityType(ctx, id).Execute()

Delete a Custom Entity Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteCustomEntityType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCustomEntityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomEntityType`: DeleteCustomEntityTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteCustomEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomEntityTypeResponse**](DeleteCustomEntityTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomInsight

> DeleteCustomInsightResponse DeleteCustomInsight(ctx, id).Execute()

Delete a Custom Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteCustomInsight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCustomInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomInsight`: DeleteCustomInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteCustomInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomInsightResponse**](DeleteCustomInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomMatchListColumn

> DeleteCustomMatchListColumnResponse DeleteCustomMatchListColumn(ctx, id).Execute()

Delete a Custom Match List Column



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteCustomMatchListColumn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCustomMatchListColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomMatchListColumn`: DeleteCustomMatchListColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteCustomMatchListColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomMatchListColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomMatchListColumnResponse**](DeleteCustomMatchListColumnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerSourcedEntityLookupTable

> DeleteCustomerSourcedEntityLookupTableResponse DeleteCustomerSourcedEntityLookupTable(ctx, id).Execute()

Remove a customer sourced entity lookup table



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteCustomerSourcedEntityLookupTable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteCustomerSourcedEntityLookupTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomerSourcedEntityLookupTable`: DeleteCustomerSourcedEntityLookupTableResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteCustomerSourcedEntityLookupTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerSourcedEntityLookupTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCustomerSourcedEntityLookupTableResponse**](DeleteCustomerSourcedEntityLookupTableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntityCriticalityConfig

> DeleteEntityCriticalityConfigResponse DeleteEntityCriticalityConfig(ctx, id).Execute()

Delete an Entity Criticality Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteEntityCriticalityConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEntityCriticalityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntityCriticalityConfig`: DeleteEntityCriticalityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteEntityCriticalityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityCriticalityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEntityCriticalityConfigResponse**](DeleteEntityCriticalityConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntityGroup

> DeleteEntityGroupResponse DeleteEntityGroup(ctx, id).Execute()

Delete an Entity Group Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteEntityGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEntityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntityGroup`: DeleteEntityGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteEntityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEntityGroupResponse**](DeleteEntityGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntityNote

> DeleteEntityNoteResponse DeleteEntityNote(ctx, entityId, entityNoteId).Execute()

Remove an Entity Note from an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    entityNoteId := "entityNoteId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteEntityNote(context.Background(), entityId, entityNoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteEntityNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntityNote`: DeleteEntityNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteEntityNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 
**entityNoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteEntityNoteResponse**](DeleteEntityNoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInsightResolution

> DeleteInsightResolutionResponse DeleteInsightResolution(ctx, id).Execute()

Delete a Insight Resolution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteInsightResolution(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteInsightResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInsightResolution`: DeleteInsightResolutionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteInsightResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInsightResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInsightResolutionResponse**](DeleteInsightResolutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInsightStatusOption

> DeleteInsightStatusOptionResponse DeleteInsightStatusOption(ctx, id).Execute()

Delete an Insight Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteInsightStatusOption(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteInsightStatusOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInsightStatusOption`: DeleteInsightStatusOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteInsightStatusOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInsightStatusOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteInsightStatusOptionResponse**](DeleteInsightStatusOptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogMapping

> DeleteLogMappingResponse DeleteLogMapping(ctx, id).Execute()

Delete a Log Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteLogMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteLogMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogMapping`: DeleteLogMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteLogMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteLogMappingResponse**](DeleteLogMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMatchList

> DeleteMatchListResponse DeleteMatchList(ctx, id).Execute()

Delete a Match List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteMatchList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMatchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMatchList`: DeleteMatchListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteMatchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMatchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMatchListResponse**](DeleteMatchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMatchListItem

> DeleteMatchListItemResponse DeleteMatchListItem(ctx, id).Execute()

Delete a Match List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteMatchListItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteMatchListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMatchListItem`: DeleteMatchListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteMatchListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMatchListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMatchListItemResponse**](DeleteMatchListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkBlock

> DeleteNetworkBlockResponse DeleteNetworkBlock(ctx, id).Execute()

Delete a Network Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteNetworkBlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteNetworkBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkBlock`: DeleteNetworkBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteNetworkBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteNetworkBlockResponse**](DeleteNetworkBlockResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRuleResponse DeleteRule(ctx, id).Execute()

Delete a Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRule`: DeleteRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRuleResponse**](DeleteRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleTuningExpression

> DeleteRuleTuningExpressionResponse DeleteRuleTuningExpression(ctx, id).Execute()

Delete a Rule Tuning Expression



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteRuleTuningExpression(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteRuleTuningExpression``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRuleTuningExpression`: DeleteRuleTuningExpressionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteRuleTuningExpression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleTuningExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteRuleTuningExpressionResponse**](DeleteRuleTuningExpressionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSuppressList

> DeleteSuppressListResponse DeleteSuppressList(ctx, id).Execute()

Delete a Suppress List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteSuppressList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSuppressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSuppressList`: DeleteSuppressListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteSuppressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuppressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSuppressListResponse**](DeleteSuppressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSuppressListItem

> DeleteSuppressListItemResponse DeleteSuppressListItem(ctx, id).Execute()

Delete a Suppress List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteSuppressListItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteSuppressListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSuppressListItem`: DeleteSuppressListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteSuppressListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSuppressListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteSuppressListItemResponse**](DeleteSuppressListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagSchema

> DeleteTagSchemaResponse DeleteTagSchema(ctx, key).Execute()

Delete a Tag Schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteTagSchema(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteTagSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTagSchema`: DeleteTagSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteTagSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteTagSchemaResponse**](DeleteTagSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThreatIntelIndicator

> DeleteThreatIntelIndicatorResponse DeleteThreatIntelIndicator(ctx, id).Execute()

Delete a Threat Intel Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.DeleteThreatIntelIndicator(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteThreatIntelIndicator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteThreatIntelIndicator`: DeleteThreatIntelIndicatorResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteThreatIntelIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThreatIntelIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteThreatIntelIndicatorResponse**](DeleteThreatIntelIndicatorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteAutomation

> ExecuteAutomationResponse ExecuteAutomation(ctx).ExecuteAutomationRequestBody(executeAutomationRequestBody).Execute()

Execute a list of automations to a list of resources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    executeAutomationRequestBody := *openapiclient.NewExecuteAutomationRequestBody(*openapiclient.NewExecuteAutomationRequestBodyFields([]openapiclient.ExecuteAutomationRequestBodyFieldsExecutionCseResourcesInner{*openapiclient.NewExecuteAutomationRequestBodyFieldsExecutionCseResourcesInner("Id_example", "CseResourceType_example")}, []string{"AutomationIds_example"})) // ExecuteAutomationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ExecuteAutomation(context.Background()).ExecuteAutomationRequestBody(executeAutomationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ExecuteAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteAutomation`: ExecuteAutomationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ExecuteAutomation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeAutomationRequestBody** | [**ExecuteAutomationRequestBody**](ExecuteAutomationRequestBody.md) |  | 

### Return type

[**ExecuteAutomationResponse**](ExecuteAutomationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllEntities

> GetAllEntitiesResponse GetAllEntities(ctx).Q(q).NextPageToken(nextPageToken).Expand(expand).Execute()

Get the list of all Entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `ip` - `hostname` - `username` - `value` - `sensorZone` - `whitelisted` - `type` - `tag` - `activityScore` - `recentSignalSeverity` - `lastSeen` - `criticality` - `reputation`      (optional)
    nextPageToken := "nextPageToken_example" // string |  (optional)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAllEntities(context.Background()).Q(q).NextPageToken(nextPageToken).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllEntities`: GetAllEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;ip&#x60; - &#x60;hostname&#x60; - &#x60;username&#x60; - &#x60;value&#x60; - &#x60;sensorZone&#x60; - &#x60;whitelisted&#x60; - &#x60;type&#x60; - &#x60;tag&#x60; - &#x60;activityScore&#x60; - &#x60;recentSignalSeverity&#x60; - &#x60;lastSeen&#x60; - &#x60;criticality&#x60; - &#x60;reputation&#x60;      | 
 **nextPageToken** | **string** |  | 
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetAllEntitiesResponse**](GetAllEntitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInsights

> GetAllInsightsResponse GetAllInsights(ctx).RecordSummaryFields(recordSummaryFields).Q(q).NextPageToken(nextPageToken).Expand(expand).Execute()

Get the list of all Insights



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    recordSummaryFields := []string{"Inner_example"} // []string | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `readableId` - `status` - `statusId` - `name` - `insightId` - `serialId` - `description` - `created` - `timestamp` - `closed` - `assignee` - `entity.id` - `entity.ip` - `entity.hostname` - `entity.username` - `entity.sensorZone` - `entity.type` - `entity.value` - `enrichment` - `sensorZone` - `tag` - `severity` - `resolution` - `subResolution` - `ruleId` - `records` - `confidence`      (optional)
    nextPageToken := "nextPageToken_example" // string |  (optional)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAllInsights(context.Background()).RecordSummaryFields(recordSummaryFields).Q(q).NextPageToken(nextPageToken).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInsights`: GetAllInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordSummaryFields** | **[]string** | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;readableId&#x60; - &#x60;status&#x60; - &#x60;statusId&#x60; - &#x60;name&#x60; - &#x60;insightId&#x60; - &#x60;serialId&#x60; - &#x60;description&#x60; - &#x60;created&#x60; - &#x60;timestamp&#x60; - &#x60;closed&#x60; - &#x60;assignee&#x60; - &#x60;entity.id&#x60; - &#x60;entity.ip&#x60; - &#x60;entity.hostname&#x60; - &#x60;entity.username&#x60; - &#x60;entity.sensorZone&#x60; - &#x60;entity.type&#x60; - &#x60;entity.value&#x60; - &#x60;enrichment&#x60; - &#x60;sensorZone&#x60; - &#x60;tag&#x60; - &#x60;severity&#x60; - &#x60;resolution&#x60; - &#x60;subResolution&#x60; - &#x60;ruleId&#x60; - &#x60;records&#x60; - &#x60;confidence&#x60;      | 
 **nextPageToken** | **string** |  | 
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetAllInsightsResponse**](GetAllInsightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSignals

> GetAllSignalsResponse GetAllSignals(ctx).Q(q).NextPageToken(nextPageToken).Execute()

Get the list of all Signals



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `stage` - `contentType` - `name` - `description` - `created` - `timestamp` - `severity` - `entity.id` - `entity.ip` - `entity.hostname` - `entity.username` - `entity.value` - `entity.type` - `entity.sensorZone` - `involved_entities` - `sensorZone` - `suppressed` - `ruleId` - `ruleType` - `prototype` - `records` - `tag` - `vendor` - `product`      (optional)
    nextPageToken := "nextPageToken_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAllSignals(context.Background()).Q(q).NextPageToken(nextPageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSignals`: GetAllSignalsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;stage&#x60; - &#x60;contentType&#x60; - &#x60;name&#x60; - &#x60;description&#x60; - &#x60;created&#x60; - &#x60;timestamp&#x60; - &#x60;severity&#x60; - &#x60;entity.id&#x60; - &#x60;entity.ip&#x60; - &#x60;entity.hostname&#x60; - &#x60;entity.username&#x60; - &#x60;entity.value&#x60; - &#x60;entity.type&#x60; - &#x60;entity.sensorZone&#x60; - &#x60;involved_entities&#x60; - &#x60;sensorZone&#x60; - &#x60;suppressed&#x60; - &#x60;ruleId&#x60; - &#x60;ruleType&#x60; - &#x60;prototype&#x60; - &#x60;records&#x60; - &#x60;tag&#x60; - &#x60;vendor&#x60; - &#x60;product&#x60;      | 
 **nextPageToken** | **string** |  | 

### Return type

[**GetAllSignalsResponse**](GetAllSignalsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllThreatIntelIndicators

> GetAllThreatIntelIndicatorsResponse GetAllThreatIntelIndicators(ctx).Value(value).Q(q).SourceIds(sourceIds).SourceName(sourceName).NextPageToken(nextPageToken).Expand(expand).Execute()

Get a list of all Threat Intel Indicators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    value := "value_example" // string | A value to search for (optional)
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `targetColumn` - `value` - `active` - `expirationDate` - `listName` - `description` - `created`      (optional)
    sourceIds := []string{"Inner_example"} // []string |  (optional)
    sourceName := "sourceName_example" // string |  (optional)
    nextPageToken := "nextPageToken_example" // string |  (optional)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAllThreatIntelIndicators(context.Background()).Value(value).Q(q).SourceIds(sourceIds).SourceName(sourceName).NextPageToken(nextPageToken).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAllThreatIntelIndicators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllThreatIntelIndicators`: GetAllThreatIntelIndicatorsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAllThreatIntelIndicators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllThreatIntelIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | A value to search for | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;targetColumn&#x60; - &#x60;value&#x60; - &#x60;active&#x60; - &#x60;expirationDate&#x60; - &#x60;listName&#x60; - &#x60;description&#x60; - &#x60;created&#x60;      | 
 **sourceIds** | **[]string** |  | 
 **sourceName** | **string** |  | 
 **nextPageToken** | **string** |  | 
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetAllThreatIntelIndicatorsResponse**](GetAllThreatIntelIndicatorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomation

> GetAutomationResponse GetAutomation(ctx, id).Execute()

Get an Automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAutomation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomation`: GetAutomationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAutomation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAutomationResponse**](GetAutomationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomations

> GetAutomationsResponse GetAutomations(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Automations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetAutomations(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetAutomations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomations`: GetAutomationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetAutomations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetAutomationsResponse**](GetAutomationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextAction

> GetContextActionResponse GetContextAction(ctx, id).Execute()

Get a Context Action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetContextAction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContextAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextAction`: GetContextActionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContextAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContextActionResponse**](GetContextActionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextActions

> GetContextActionsResponse GetContextActions(ctx).IocType(iocType).Enabled(enabled).Offset(offset).Limit(limit).Execute()

Get the list of Context Actions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    iocType := "iocType_example" // string |  (optional)
    enabled := true // bool |  (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetContextActions(context.Background()).IocType(iocType).Enabled(enabled).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContextActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextActions`: GetContextActionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContextActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContextActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iocType** | **string** |  | 
 **enabled** | **bool** |  | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetContextActionsResponse**](GetContextActionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEntityType

> GetCustomEntityTypeResponse GetCustomEntityType(ctx, id).Execute()

Get a Custom Entity Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomEntityType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomEntityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomEntityType`: GetCustomEntityTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomEntityTypeResponse**](GetCustomEntityTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEntityTypes

> GetCustomEntityTypesResponse GetCustomEntityTypes(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Custom Entity Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomEntityTypes(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomEntityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomEntityTypes`: GetCustomEntityTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomEntityTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEntityTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetCustomEntityTypesResponse**](GetCustomEntityTypesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomInsight

> GetCustomInsightResponse GetCustomInsight(ctx, id).Execute()

Get a Custom Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomInsight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomInsight`: GetCustomInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomInsightResponse**](GetCustomInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomInsights

> GetCustomInsightsResponse GetCustomInsights(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Custom Insights



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomInsights(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomInsights`: GetCustomInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetCustomInsightsResponse**](GetCustomInsightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomMatchListColumn

> GetCustomMatchListColumnResponse GetCustomMatchListColumn(ctx, id).Execute()

Get a Custom Match List Column



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomMatchListColumn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomMatchListColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomMatchListColumn`: GetCustomMatchListColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomMatchListColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomMatchListColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomMatchListColumnResponse**](GetCustomMatchListColumnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomMatchListColumns

> GetCustomMatchListColumnsResponse GetCustomMatchListColumns(ctx).Execute()

Get the list of Custom Match List Columns



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomMatchListColumns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomMatchListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomMatchListColumns`: GetCustomMatchListColumnsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomMatchListColumns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomMatchListColumnsRequest struct via the builder pattern


### Return type

[**GetCustomMatchListColumnsResponse**](GetCustomMatchListColumnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerSourcedLookupTable

> GetCustomerSourcedLookupTableResponse GetCustomerSourcedLookupTable(ctx, id).Execute()

Get a customer sourced entity lookup table



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomerSourcedLookupTable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomerSourcedLookupTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerSourcedLookupTable`: GetCustomerSourcedLookupTableResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomerSourcedLookupTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerSourcedLookupTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCustomerSourcedLookupTableResponse**](GetCustomerSourcedLookupTableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerSourcedLookupTables

> GetCustomerSourcedLookupTablesResponse GetCustomerSourcedLookupTables(ctx).Execute()

Get the list of customer sourced entity lookup tables



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetCustomerSourcedLookupTables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCustomerSourcedLookupTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerSourcedLookupTables`: GetCustomerSourcedLookupTablesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCustomerSourcedLookupTables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerSourcedLookupTablesRequest struct via the builder pattern


### Return type

[**GetCustomerSourcedLookupTablesResponse**](GetCustomerSourcedLookupTablesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntities

> GetEntitiesResponse GetEntities(ctx).Q(q).Offset(offset).Limit(limit).Expand(expand).Execute()

Get the list of Entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `ip` - `hostname` - `username` - `value` - `sensorZone` - `whitelisted` - `type` - `tag` - `activityScore` - `recentSignalSeverity` - `lastSeen` - `criticality` - `reputation`      (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntities(context.Background()).Q(q).Offset(offset).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntities`: GetEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;ip&#x60; - &#x60;hostname&#x60; - &#x60;username&#x60; - &#x60;value&#x60; - &#x60;sensorZone&#x60; - &#x60;whitelisted&#x60; - &#x60;type&#x60; - &#x60;tag&#x60; - &#x60;activityScore&#x60; - &#x60;recentSignalSeverity&#x60; - &#x60;lastSeen&#x60; - &#x60;criticality&#x60; - &#x60;reputation&#x60;      | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetEntitiesResponse**](GetEntitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntity

> GetEntityResponse GetEntity(ctx, id).Expand(expand).Execute()

Get an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntity(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntity`: GetEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetEntityResponse**](GetEntityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityCriticalityConfig

> GetEntityCriticalityConfigResponse GetEntityCriticalityConfig(ctx, id).Execute()

Get an Entity Criticality Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityCriticalityConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityCriticalityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityCriticalityConfig`: GetEntityCriticalityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityCriticalityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityCriticalityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityCriticalityConfigResponse**](GetEntityCriticalityConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityCriticalityConfigs

> GetEntityCriticalityConfigsResponse GetEntityCriticalityConfigs(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Entity Criticality Configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityCriticalityConfigs(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityCriticalityConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityCriticalityConfigs`: GetEntityCriticalityConfigsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityCriticalityConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityCriticalityConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetEntityCriticalityConfigsResponse**](GetEntityCriticalityConfigsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityDomainConfiguration

> GetEntityDomainConfigurationResponse GetEntityDomainConfiguration(ctx).Execute()

Get the Entity Domain Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityDomainConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityDomainConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityDomainConfiguration`: GetEntityDomainConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityDomainConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityDomainConfigurationRequest struct via the builder pattern


### Return type

[**GetEntityDomainConfigurationResponse**](GetEntityDomainConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityEnrichments

> GetEntityEnrichmentsResponse GetEntityEnrichments(ctx, id).Execute()

Get an Entity's enrichments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityEnrichments(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityEnrichments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityEnrichments`: GetEntityEnrichmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityEnrichments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityEnrichmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityEnrichmentsResponse**](GetEntityEnrichmentsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityExecutedAutomations

> GetEntityExecutedAutomationsResponse GetEntityExecutedAutomations(ctx, id).Execute()

Get entity executed automations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityExecutedAutomations(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityExecutedAutomations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityExecutedAutomations`: GetEntityExecutedAutomationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityExecutedAutomations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityExecutedAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityExecutedAutomationsResponse**](GetEntityExecutedAutomationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityGroup

> GetEntityGroupResponse GetEntityGroup(ctx, id).Execute()

Get an Entity Group Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityGroup`: GetEntityGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityGroupResponse**](GetEntityGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityGroups

> GetEntityGroupsResponse GetEntityGroups(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Entity Group Configurations for a given query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityGroups(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityGroups`: GetEntityGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetEntityGroupsResponse**](GetEntityGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityNotes

> GetEntityNotesResponse GetEntityNotes(ctx, entityId).Execute()

Get the list of entity notes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetEntityNotes(context.Background(), entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetEntityNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityNotes`: GetEntityNotesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetEntityNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityNotesResponse**](GetEntityNotesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsight

> GetInsightResponse GetInsight(ctx, id).RecordSummaryFields(recordSummaryFields).Exclude(exclude).Execute()

Get an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    recordSummaryFields := []string{"Inner_example"} // []string | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response
    exclude := []string{"Exclude_example"} // []string | A comma-separated list of subfields to be excluded from the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsight(context.Background(), id).RecordSummaryFields(recordSummaryFields).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsight`: GetInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordSummaryFields** | **[]string** | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response | 
 **exclude** | **[]string** | A comma-separated list of subfields to be excluded from the response | 

### Return type

[**GetInsightResponse**](GetInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightComments

> GetInsightCommentsResponse GetInsightComments(ctx, id).Execute()

Get an Insight's comments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightComments(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightComments`: GetInsightCommentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightCommentsResponse**](GetInsightCommentsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightCounts

> GetInsightCountsResponse GetInsightCounts(ctx).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()

Get the count of Insights over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    startTimestamp := time.Now() // time.Time | 
    bucketDuration := int32(56) // int32 | The duration of the buckets in seconds
    endTimestamp := time.Now() // time.Time |  (optional)
    timezone := "timezone_example" // string | The timezone to use for creating the bucket cutoffs (optional) (default to "UTC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightCounts(context.Background()).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightCounts`: GetInsightCountsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **time.Time** |  | 
 **bucketDuration** | **int32** | The duration of the buckets in seconds | 
 **endTimestamp** | **time.Time** |  | 
 **timezone** | **string** | The timezone to use for creating the bucket cutoffs | [default to &quot;UTC&quot;]

### Return type

[**GetInsightCountsResponse**](GetInsightCountsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightEnrichments

> GetInsightEnrichmentsResponse GetInsightEnrichments(ctx, id).Execute()

Get an Insights's enrichments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightEnrichments(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightEnrichments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightEnrichments`: GetInsightEnrichmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightEnrichments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightEnrichmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightEnrichmentsResponse**](GetInsightEnrichmentsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightExecutedAutomations

> GetInsightExecutedAutomationsResponse GetInsightExecutedAutomations(ctx, id).Execute()

Get insight executed automations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightExecutedAutomations(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightExecutedAutomations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightExecutedAutomations`: GetInsightExecutedAutomationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightExecutedAutomations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightExecutedAutomationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightExecutedAutomationsResponse**](GetInsightExecutedAutomationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightHistory

> GetInsightHistoryResponse GetInsightHistory(ctx, id).Execute()

Get an Insight's history



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightHistory`: GetInsightHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightHistoryResponse**](GetInsightHistoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightRelatedEntities

> GetInsightRelatedEntitiesResponse GetInsightRelatedEntities(ctx, id).Execute()

Get an Insight's list of involved entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightRelatedEntities(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightRelatedEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightRelatedEntities`: GetInsightRelatedEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightRelatedEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightRelatedEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightRelatedEntitiesResponse**](GetInsightRelatedEntitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightResolution

> GetInsightResolutionResponse GetInsightResolution(ctx, id).Execute()

Get a Insight Resolution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightResolution(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightResolution`: GetInsightResolutionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightResolutionResponse**](GetInsightResolutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightResolutions

> GetInsightResolutionsResponse GetInsightResolutions(ctx).Offset(offset).Limit(limit).Execute()

Get the list of Insight Resolutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightResolutions(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightResolutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightResolutions`: GetInsightResolutionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightResolutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightResolutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetInsightResolutionsResponse**](GetInsightResolutionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightStatus

> GetInsightStatusResponse GetInsightStatus(ctx, id).Execute()

Get an Insight Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightStatus`: GetInsightStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInsightStatusResponse**](GetInsightStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightStatuses

> GetInsightStatusesResponse GetInsightStatuses(ctx).IncludeDisabled(includeDisabled).Execute()

Get the list of Insight Statuses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    includeDisabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightStatuses(context.Background()).IncludeDisabled(includeDisabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightStatuses`: GetInsightStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDisabled** | **bool** |  | 

### Return type

[**GetInsightStatusesResponse**](GetInsightStatusesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsights

> GetInsightsResponse GetInsights(ctx).RecordSummaryFields(recordSummaryFields).Q(q).Offset(offset).Limit(limit).Exclude(exclude).Execute()

Get the list of Insights for a given query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    recordSummaryFields := []string{"Inner_example"} // []string | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `readableId` - `status` - `statusId` - `name` - `insightId` - `serialId` - `description` - `created` - `timestamp` - `closed` - `assignee` - `entity.id` - `entity.ip` - `entity.hostname` - `entity.username` - `entity.sensorZone` - `entity.type` - `entity.value` - `enrichment` - `sensorZone` - `tag` - `severity` - `resolution` - `subResolution` - `ruleId` - `records` - `confidence`      (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 10)
    exclude := []string{"Exclude_example"} // []string | A comma-separated list of subfields to be excluded from the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsights(context.Background()).RecordSummaryFields(recordSummaryFields).Q(q).Offset(offset).Limit(limit).Exclude(exclude).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsights`: GetInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordSummaryFields** | **[]string** | A list of fields to aggregate from the records of each Insight into a summarized list directly on the Insight object of the response | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;readableId&#x60; - &#x60;status&#x60; - &#x60;statusId&#x60; - &#x60;name&#x60; - &#x60;insightId&#x60; - &#x60;serialId&#x60; - &#x60;description&#x60; - &#x60;created&#x60; - &#x60;timestamp&#x60; - &#x60;closed&#x60; - &#x60;assignee&#x60; - &#x60;entity.id&#x60; - &#x60;entity.ip&#x60; - &#x60;entity.hostname&#x60; - &#x60;entity.username&#x60; - &#x60;entity.sensorZone&#x60; - &#x60;entity.type&#x60; - &#x60;entity.value&#x60; - &#x60;enrichment&#x60; - &#x60;sensorZone&#x60; - &#x60;tag&#x60; - &#x60;severity&#x60; - &#x60;resolution&#x60; - &#x60;subResolution&#x60; - &#x60;ruleId&#x60; - &#x60;records&#x60; - &#x60;confidence&#x60;      | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 10]
 **exclude** | **[]string** | A comma-separated list of subfields to be excluded from the response | 

### Return type

[**GetInsightsResponse**](GetInsightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightsConfiguration

> GetInsightsConfigurationResponse GetInsightsConfiguration(ctx).Execute()

Get Insight Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetInsightsConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetInsightsConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightsConfiguration`: GetInsightsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetInsightsConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightsConfigurationRequest struct via the builder pattern


### Return type

[**GetInsightsConfigurationResponse**](GetInsightsConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogMapping

> GetLogMappingResponse GetLogMapping(ctx, id).Execute()

Get a Log Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetLogMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLogMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogMapping`: GetLogMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLogMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLogMappingResponse**](GetLogMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogMappingVendorsAndProducts

> GetLogMappingVendorsAndProductsResponse GetLogMappingVendorsAndProducts(ctx).Product(product).Vendor(vendor).Execute()

Get a List of Log Mapping Vendors and Products



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    product := "product_example" // string |  (optional)
    vendor := "vendor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetLogMappingVendorsAndProducts(context.Background()).Product(product).Vendor(vendor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLogMappingVendorsAndProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogMappingVendorsAndProducts`: GetLogMappingVendorsAndProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLogMappingVendorsAndProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogMappingVendorsAndProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** |  | 
 **vendor** | **string** |  | 

### Return type

[**GetLogMappingVendorsAndProductsResponse**](GetLogMappingVendorsAndProductsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogMappings

> GetLogMappingsResponse GetLogMappings(ctx).Q(q).Offset(offset).Limit(limit).Execute()

Get the list of Log Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `parentId` - `name` - `created` - `createdBy` - `lastUpdated` - `lastUpdatedBy` - `source` - `isCustom` - `vendor` - `product` - `recordType` - `enabled` - `logFormat` - `eventIdPattern` - `patternName` - `isStructured` - `recordCount07D` - `recordCount24H`      (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetLogMappings(context.Background()).Q(q).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetLogMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogMappings`: GetLogMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetLogMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;parentId&#x60; - &#x60;name&#x60; - &#x60;created&#x60; - &#x60;createdBy&#x60; - &#x60;lastUpdated&#x60; - &#x60;lastUpdatedBy&#x60; - &#x60;source&#x60; - &#x60;isCustom&#x60; - &#x60;vendor&#x60; - &#x60;product&#x60; - &#x60;recordType&#x60; - &#x60;enabled&#x60; - &#x60;logFormat&#x60; - &#x60;eventIdPattern&#x60; - &#x60;patternName&#x60; - &#x60;isStructured&#x60; - &#x60;recordCount07D&#x60; - &#x60;recordCount24H&#x60;      | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetLogMappingsResponse**](GetLogMappingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchList

> GetMatchListResponse GetMatchList(ctx, id).GetTargetCustomColumnName(getTargetCustomColumnName).Execute()

Get a Match List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    getTargetCustomColumnName := true // bool | Return the target_column values as names instead of IDs when they are custom columns (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetMatchList(context.Background(), id).GetTargetCustomColumnName(getTargetCustomColumnName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMatchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMatchList`: GetMatchListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMatchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getTargetCustomColumnName** | **bool** | Return the target_column values as names instead of IDs when they are custom columns | [default to false]

### Return type

[**GetMatchListResponse**](GetMatchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchListItem

> GetMatchListItemResponse GetMatchListItem(ctx, id).Execute()

Get a Match List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetMatchListItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMatchListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMatchListItem`: GetMatchListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMatchListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMatchListItemResponse**](GetMatchListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchListItems

> GetMatchListItemsResponse GetMatchListItems(ctx).Offset(offset).Limit(limit).Value(value).Q(q).ListIds(listIds).Execute()

Get a list of Match List Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    value := "value_example" // string | A value to search for (optional)
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `targetColumn` - `value` - `active` - `expirationDate` - `listName` - `description` - `created`      (optional)
    listIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetMatchListItems(context.Background()).Offset(offset).Limit(limit).Value(value).Q(q).ListIds(listIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMatchListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMatchListItems`: GetMatchListItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMatchListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **value** | **string** | A value to search for | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;targetColumn&#x60; - &#x60;value&#x60; - &#x60;active&#x60; - &#x60;expirationDate&#x60; - &#x60;listName&#x60; - &#x60;description&#x60; - &#x60;created&#x60;      | 
 **listIds** | **[]string** |  | 

### Return type

[**GetMatchListItemsResponse**](GetMatchListItemsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatchLists

> GetMatchListsResponse GetMatchLists(ctx).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()

Get the list of Match Lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    sort := "sort_example" // string |  (optional)
    sortDir := "sortDir_example" // string |  (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetMatchLists(context.Background()).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetMatchLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMatchLists`: GetMatchListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetMatchLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMatchListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **sort** | **string** |  | 
 **sortDir** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**GetMatchListsResponse**](GetMatchListsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlock

> GetNetworkBlockResponse GetNetworkBlock(ctx, id).Execute()

Get a Network Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetNetworkBlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNetworkBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBlock`: GetNetworkBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNetworkBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkBlockResponse**](GetNetworkBlockResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBlocks

> GetNetworkBlocksResponse GetNetworkBlocks(ctx).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()

Get the list of Network Blocks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    sort := "sort_example" // string |  (optional) (default to "ADDRESS_BLOCK")
    sortDir := "sortDir_example" // string |  (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetNetworkBlocks(context.Background()).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNetworkBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkBlocks`: GetNetworkBlocksResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNetworkBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **sort** | **string** |  | [default to &quot;ADDRESS_BLOCK&quot;]
 **sortDir** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**GetNetworkBlocksResponse**](GetNetworkBlocksResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordCounts

> GetRecordCountsResponse GetRecordCounts(ctx).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()

Get the count of Records over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    startTimestamp := time.Now() // time.Time | 
    bucketDuration := int32(56) // int32 | 
    endTimestamp := time.Now() // time.Time |  (optional)
    timezone := "timezone_example" // string | The timezone to use for creating the bucket cutoffs (optional) (default to "UTC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRecordCounts(context.Background()).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRecordCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecordCounts`: GetRecordCountsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRecordCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **time.Time** |  | 
 **bucketDuration** | **int32** |  | 
 **endTimestamp** | **time.Time** |  | 
 **timezone** | **string** | The timezone to use for creating the bucket cutoffs | [default to &quot;UTC&quot;]

### Return type

[**GetRecordCountsResponse**](GetRecordCountsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedEntitiesById

> GetRelatedEntitiesByIdResponse GetRelatedEntitiesById(ctx, id).Start(start).End(end).Execute()

Get the list of Related Entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    start := time.Now() // time.Time | The start of the time range for which to search.
    end := time.Now() // time.Time | The end of the time range for which to search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRelatedEntitiesById(context.Background(), id).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRelatedEntitiesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelatedEntitiesById`: GetRelatedEntitiesByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRelatedEntitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedEntitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | The start of the time range for which to search. | 
 **end** | **time.Time** | The end of the time range for which to search. | 

### Return type

[**GetRelatedEntitiesByIdResponse**](GetRelatedEntitiesByIdResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> GetRuleResponse GetRule(ctx, id).Expand(expand).Execute()

Get a Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRule(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: GetRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetRuleResponse**](GetRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleTuningExpression

> GetRuleTuningExpressionResponse GetRuleTuningExpression(ctx, id).Execute()

Get a Rule Tuning Expression



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRuleTuningExpression(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRuleTuningExpression``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleTuningExpression`: GetRuleTuningExpressionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRuleTuningExpression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleTuningExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRuleTuningExpressionResponse**](GetRuleTuningExpressionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleTuningExpressions

> GetRuleTuningExpressionsResponse GetRuleTuningExpressions(ctx).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()

Get the list of Rule Tuning Expressions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    sort := "sort_example" // string |  (optional) (default to "NAME")
    sortDir := "sortDir_example" // string |  (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRuleTuningExpressions(context.Background()).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRuleTuningExpressions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleTuningExpressions`: GetRuleTuningExpressionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRuleTuningExpressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleTuningExpressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **sort** | **string** |  | [default to &quot;NAME&quot;]
 **sortDir** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**GetRuleTuningExpressionsResponse**](GetRuleTuningExpressionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRules

> GetRulesResponse GetRules(ctx).Q(q).Offset(offset).Limit(limit).Expand(expand).Execute()

Get the list of Rules for a given query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `category` - `ruleSource` - `ruleType` - `stream` - `status` - `name` - `severity` - `score` - `enabled` - `created` - `createdBy` - `lastUpdated` - `lastUpdatedBy` - `signalCount07D` - `signalCount24H` - `isPrototype` - `tag` - `hasOverride`      (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetRules(context.Background()).Q(q).Offset(offset).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRules`: GetRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;category&#x60; - &#x60;ruleSource&#x60; - &#x60;ruleType&#x60; - &#x60;stream&#x60; - &#x60;status&#x60; - &#x60;name&#x60; - &#x60;severity&#x60; - &#x60;score&#x60; - &#x60;enabled&#x60; - &#x60;created&#x60; - &#x60;createdBy&#x60; - &#x60;lastUpdated&#x60; - &#x60;lastUpdatedBy&#x60; - &#x60;signalCount07D&#x60; - &#x60;signalCount24H&#x60; - &#x60;isPrototype&#x60; - &#x60;tag&#x60; - &#x60;hasOverride&#x60;      | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetRulesResponse**](GetRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignal

> GetSignalResponse GetSignal(ctx, id).Execute()

Get a Signal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSignal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSignal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignal`: GetSignalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSignal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSignalResponse**](GetSignalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignalCounts

> GetSignalCountsResponse GetSignalCounts(ctx).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()

Get the count of Signals over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    startTimestamp := time.Now() // time.Time | 
    bucketDuration := int32(56) // int32 | The duration of the buckets in seconds
    endTimestamp := time.Now() // time.Time |  (optional)
    timezone := "timezone_example" // string | The timezone to use for creating the bucket cutoffs (See TZ database name for valid values https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) (optional) (default to "UTC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSignalCounts(context.Background()).StartTimestamp(startTimestamp).BucketDuration(bucketDuration).EndTimestamp(endTimestamp).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSignalCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignalCounts`: GetSignalCountsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSignalCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimestamp** | **time.Time** |  | 
 **bucketDuration** | **int32** | The duration of the buckets in seconds | 
 **endTimestamp** | **time.Time** |  | 
 **timezone** | **string** | The timezone to use for creating the bucket cutoffs (See TZ database name for valid values https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) | [default to &quot;UTC&quot;]

### Return type

[**GetSignalCountsResponse**](GetSignalCountsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignalEnrichments

> GetSignalEnrichmentsResponse GetSignalEnrichments(ctx, id).Execute()

Get a Signal's enrichments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSignalEnrichments(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSignalEnrichments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignalEnrichments`: GetSignalEnrichmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSignalEnrichments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalEnrichmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSignalEnrichmentsResponse**](GetSignalEnrichmentsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignals

> GetSignalsResponse GetSignals(ctx).Q(q).Offset(offset).Limit(limit).Execute()

Get the list of Signals for a given query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `stage` - `contentType` - `name` - `description` - `created` - `timestamp` - `severity` - `entity.id` - `entity.ip` - `entity.hostname` - `entity.username` - `entity.value` - `entity.type` - `entity.sensorZone` - `involved_entities` - `sensorZone` - `suppressed` - `ruleId` - `ruleType` - `prototype` - `records` - `tag` - `vendor` - `product`      (optional)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSignals(context.Background()).Q(q).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignals`: GetSignalsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;stage&#x60; - &#x60;contentType&#x60; - &#x60;name&#x60; - &#x60;description&#x60; - &#x60;created&#x60; - &#x60;timestamp&#x60; - &#x60;severity&#x60; - &#x60;entity.id&#x60; - &#x60;entity.ip&#x60; - &#x60;entity.hostname&#x60; - &#x60;entity.username&#x60; - &#x60;entity.value&#x60; - &#x60;entity.type&#x60; - &#x60;entity.sensorZone&#x60; - &#x60;involved_entities&#x60; - &#x60;sensorZone&#x60; - &#x60;suppressed&#x60; - &#x60;ruleId&#x60; - &#x60;ruleType&#x60; - &#x60;prototype&#x60; - &#x60;records&#x60; - &#x60;tag&#x60; - &#x60;vendor&#x60; - &#x60;product&#x60;      | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]

### Return type

[**GetSignalsResponse**](GetSignalsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressList

> GetSuppressListResponse GetSuppressList(ctx, id).Execute()

Get a Suppress List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSuppressList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSuppressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppressList`: GetSuppressListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSuppressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSuppressListResponse**](GetSuppressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressListItem

> GetSuppressListItemResponse GetSuppressListItem(ctx, id).Execute()

Get a Suppress List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSuppressListItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSuppressListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppressListItem`: GetSuppressListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSuppressListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSuppressListItemResponse**](GetSuppressListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressListItems

> GetSuppressListItemsResponse GetSuppressListItems(ctx).Offset(offset).Limit(limit).Value(value).Q(q).ListIds(listIds).Execute()

Get a list of Suppress List Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    value := "value_example" // string | A value to search for (optional)
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `targetColumn` - `value` - `active` - `expirationDate` - `listName` - `description` - `created`      (optional)
    listIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSuppressListItems(context.Background()).Offset(offset).Limit(limit).Value(value).Q(q).ListIds(listIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSuppressListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppressListItems`: GetSuppressListItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSuppressListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **value** | **string** | A value to search for | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;targetColumn&#x60; - &#x60;value&#x60; - &#x60;active&#x60; - &#x60;expirationDate&#x60; - &#x60;listName&#x60; - &#x60;description&#x60; - &#x60;created&#x60;      | 
 **listIds** | **[]string** |  | 

### Return type

[**GetSuppressListItemsResponse**](GetSuppressListItemsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressLists

> GetSuppressListsResponse GetSuppressLists(ctx).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()

Get the list of Suppress Lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    sort := "sort_example" // string |  (optional)
    sortDir := "sortDir_example" // string |  (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetSuppressLists(context.Background()).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSuppressLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppressLists`: GetSuppressListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSuppressLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **sort** | **string** |  | 
 **sortDir** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**GetSuppressListsResponse**](GetSuppressListsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagSchema

> GetTagSchemaResponse GetTagSchema(ctx, key).Execute()

Get a Tag Schemas



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetTagSchema(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTagSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagSchema`: GetTagSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTagSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTagSchemaResponse**](GetTagSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagSchemas

> GetTagSchemasResponse GetTagSchemas(ctx).Execute()

Get the list of Tag Schemas



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetTagSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTagSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagSchemas`: GetTagSchemasResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTagSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagSchemasRequest struct via the builder pattern


### Return type

[**GetTagSchemasResponse**](GetTagSchemasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatIntelIndicator

> GetThreatIntelIndicatorResponse GetThreatIntelIndicator(ctx, id).Expand(expand).Execute()

Get a Threat Intel Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetThreatIntelIndicator(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetThreatIntelIndicator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreatIntelIndicator`: GetThreatIntelIndicatorResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetThreatIntelIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatIntelIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetThreatIntelIndicatorResponse**](GetThreatIntelIndicatorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatIntelIndicators

> GetThreatIntelIndicatorsResponse GetThreatIntelIndicators(ctx).Offset(offset).Limit(limit).Value(value).Q(q).SourceIds(sourceIds).SourceName(sourceName).Expand(expand).Execute()

Get a list of Threat Intel Indicators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    value := "value_example" // string | A value to search for (optional)
    q := "q_example" // string |      The search query string in our custom DSL that is used to filter the results.      Operators:     - `exampleField:\"bar\"`: The value of the field is equal to \"bar\".     - `exampleField:in(\"bar\", \"baz\", \"qux\")`: The value of the field is equal to either \"bar\", \"baz\", or \"qux\".     - `exampleTextField:contains(\"foo bar\")`: The value of the field contains the phrase \"foo bar\".     - `exampleNumField:>5`: The value of the field is greater than 5. There are similar `<`, `<=`, and `>=` operators.     - `exampleNumField:5..10`: The value of the field is between 5 and 10 (inclusive).     - `exampleDateField:>2019-02-01T05:00:00+00:00`: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - `exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00`: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - `id` - `targetColumn` - `value` - `active` - `expirationDate` - `listName` - `description` - `created`      (optional)
    sourceIds := []string{"Inner_example"} // []string |  (optional)
    sourceName := "sourceName_example" // string |  (optional)
    expand := []string{"Expand_example"} // []string | A comma-separated list of subfields to be returned in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetThreatIntelIndicators(context.Background()).Offset(offset).Limit(limit).Value(value).Q(q).SourceIds(sourceIds).SourceName(sourceName).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetThreatIntelIndicators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreatIntelIndicators`: GetThreatIntelIndicatorsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetThreatIntelIndicators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatIntelIndicatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **value** | **string** | A value to search for | 
 **q** | **string** |      The search query string in our custom DSL that is used to filter the results.      Operators:     - &#x60;exampleField:\&quot;bar\&quot;&#x60;: The value of the field is equal to \&quot;bar\&quot;.     - &#x60;exampleField:in(\&quot;bar\&quot;, \&quot;baz\&quot;, \&quot;qux\&quot;)&#x60;: The value of the field is equal to either \&quot;bar\&quot;, \&quot;baz\&quot;, or \&quot;qux\&quot;.     - &#x60;exampleTextField:contains(\&quot;foo bar\&quot;)&#x60;: The value of the field contains the phrase \&quot;foo bar\&quot;.     - &#x60;exampleNumField:&gt;5&#x60;: The value of the field is greater than 5. There are similar &#x60;&lt;&#x60;, &#x60;&lt;&#x3D;&#x60;, and &#x60;&gt;&#x3D;&#x60; operators.     - &#x60;exampleNumField:5..10&#x60;: The value of the field is between 5 and 10 (inclusive).     - &#x60;exampleDateField:&gt;2019-02-01T05:00:00+00:00&#x60;: The value of the date field is after 5 a.m. UTC time on February 2,         2019.     - &#x60;exampleDateField:2019-02-01T05:00:00+00:00..2019-02-01T08:00:00+00:00&#x60;: The value of the date field is between 5 a.m.         and 8 a.m. UTC time on February 2, 2019.      Fields:     - &#x60;id&#x60; - &#x60;targetColumn&#x60; - &#x60;value&#x60; - &#x60;active&#x60; - &#x60;expirationDate&#x60; - &#x60;listName&#x60; - &#x60;description&#x60; - &#x60;created&#x60;      | 
 **sourceIds** | **[]string** |  | 
 **sourceName** | **string** |  | 
 **expand** | **[]string** | A comma-separated list of subfields to be returned in the response | 

### Return type

[**GetThreatIntelIndicatorsResponse**](GetThreatIntelIndicatorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatIntelSource

> GetThreatIntelSourceResponse GetThreatIntelSource(ctx, id).Execute()

Get a Threat Intel Source



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetThreatIntelSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetThreatIntelSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreatIntelSource`: GetThreatIntelSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetThreatIntelSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatIntelSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetThreatIntelSourceResponse**](GetThreatIntelSourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatIntelligenceSources

> GetThreatIntelligenceSourcesResponse GetThreatIntelligenceSources(ctx).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()

Get the list of Threat Intel Sources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set (optional) (default to 0)
    limit := int32(56) // int32 | The numbers of items to return (optional) (default to 50)
    sort := "sort_example" // string |  (optional)
    sortDir := "sortDir_example" // string |  (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetThreatIntelligenceSources(context.Background()).Offset(offset).Limit(limit).Sort(sort).SortDir(sortDir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetThreatIntelligenceSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreatIntelligenceSources`: GetThreatIntelligenceSourcesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetThreatIntelligenceSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatIntelligenceSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **int32** | The numbers of items to return | [default to 50]
 **sort** | **string** |  | 
 **sortDir** | **string** |  | [default to &quot;ASC&quot;]

### Return type

[**GetThreatIntelligenceSourcesResponse**](GetThreatIntelligenceSourcesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupNetworkBlocksByIp

> LookupNetworkBlocksByIpResponse LookupNetworkBlocksByIp(ctx).Address(address).Execute()

Lookup Network Blocks that match a specific IP address



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    address := "address_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.LookupNetworkBlocksByIp(context.Background()).Address(address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LookupNetworkBlocksByIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupNetworkBlocksByIp`: LookupNetworkBlocksByIpResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LookupNetworkBlocksByIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupNetworkBlocksByIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | 

### Return type

[**LookupNetworkBlocksByIpResponse**](LookupNetworkBlocksByIpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideAggregationRule

> OverrideAggregationRuleResponse OverrideAggregationRule(ctx, id).OverrideAggregationRuleRequestBody(overrideAggregationRuleRequestBody).Execute()

Override a Aggregation Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideAggregationRuleRequestBody := *openapiclient.NewOverrideAggregationRuleRequestBody(*openapiclient.NewOverrideAggregationRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideAggregationRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideAggregationRule(context.Background(), id).OverrideAggregationRuleRequestBody(overrideAggregationRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideAggregationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideAggregationRule`: OverrideAggregationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideAggregationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideAggregationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideAggregationRuleRequestBody** | [**OverrideAggregationRuleRequestBody**](OverrideAggregationRuleRequestBody.md) |  | 

### Return type

[**OverrideAggregationRuleResponse**](OverrideAggregationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideChainRule

> OverrideChainRuleResponse OverrideChainRule(ctx, id).OverrideChainRuleRequestBody(overrideChainRuleRequestBody).Execute()

Override a Chain Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideChainRuleRequestBody := *openapiclient.NewOverrideChainRuleRequestBody(*openapiclient.NewOverrideChainRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideChainRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideChainRule(context.Background(), id).OverrideChainRuleRequestBody(overrideChainRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideChainRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideChainRule`: OverrideChainRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideChainRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideChainRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideChainRuleRequestBody** | [**OverrideChainRuleRequestBody**](OverrideChainRuleRequestBody.md) |  | 

### Return type

[**OverrideChainRuleResponse**](OverrideChainRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideFirstSeenRule

> OverrideFirstSeenRuleResponse OverrideFirstSeenRule(ctx, id).OverrideFirstSeenRuleRequestBody(overrideFirstSeenRuleRequestBody).Execute()

Override a First Seen Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideFirstSeenRuleRequestBody := *openapiclient.NewOverrideFirstSeenRuleRequestBody(*openapiclient.NewOverrideFirstSeenRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideFirstSeenRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideFirstSeenRule(context.Background(), id).OverrideFirstSeenRuleRequestBody(overrideFirstSeenRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideFirstSeenRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideFirstSeenRule`: OverrideFirstSeenRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideFirstSeenRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideFirstSeenRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideFirstSeenRuleRequestBody** | [**OverrideFirstSeenRuleRequestBody**](OverrideFirstSeenRuleRequestBody.md) |  | 

### Return type

[**OverrideFirstSeenRuleResponse**](OverrideFirstSeenRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideMatchRule

> OverrideMatchRuleResponse OverrideMatchRule(ctx, id).OverrideMatchRuleRequestBody(overrideMatchRuleRequestBody).Execute()

Override a Legacy Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideMatchRuleRequestBody := *openapiclient.NewOverrideMatchRuleRequestBody(*openapiclient.NewOverrideMatchRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideMatchRule(context.Background(), id).OverrideMatchRuleRequestBody(overrideMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideMatchRule`: OverrideMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideMatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideMatchRuleRequestBody** | [**OverrideMatchRuleRequestBody**](OverrideMatchRuleRequestBody.md) |  | 

### Return type

[**OverrideMatchRuleResponse**](OverrideMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideOutlierRule

> OverrideOutlierRuleResponse OverrideOutlierRule(ctx, id).OverrideOutlierRuleRequestBody(overrideOutlierRuleRequestBody).Execute()

Override a Outlier Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideOutlierRuleRequestBody := *openapiclient.NewOverrideOutlierRuleRequestBody(*openapiclient.NewOverrideOutlierRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideOutlierRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideOutlierRule(context.Background(), id).OverrideOutlierRuleRequestBody(overrideOutlierRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideOutlierRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideOutlierRule`: OverrideOutlierRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideOutlierRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideOutlierRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideOutlierRuleRequestBody** | [**OverrideOutlierRuleRequestBody**](OverrideOutlierRuleRequestBody.md) |  | 

### Return type

[**OverrideOutlierRuleResponse**](OverrideOutlierRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideTemplatedMatchRule

> OverrideTemplatedMatchRuleResponse OverrideTemplatedMatchRule(ctx, id).OverrideTemplatedMatchRuleRequestBody(overrideTemplatedMatchRuleRequestBody).Execute()

Override a Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideTemplatedMatchRuleRequestBody := *openapiclient.NewOverrideTemplatedMatchRuleRequestBody(*openapiclient.NewOverrideTemplatedMatchRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideTemplatedMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideTemplatedMatchRule(context.Background(), id).OverrideTemplatedMatchRuleRequestBody(overrideTemplatedMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideTemplatedMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideTemplatedMatchRule`: OverrideTemplatedMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideTemplatedMatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideTemplatedMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideTemplatedMatchRuleRequestBody** | [**OverrideTemplatedMatchRuleRequestBody**](OverrideTemplatedMatchRuleRequestBody.md) |  | 

### Return type

[**OverrideTemplatedMatchRuleResponse**](OverrideTemplatedMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideThresholdRule

> OverrideThresholdRuleResponse OverrideThresholdRule(ctx, id).OverrideThresholdRuleRequestBody(overrideThresholdRuleRequestBody).Execute()

Override a Threshold Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    overrideThresholdRuleRequestBody := *openapiclient.NewOverrideThresholdRuleRequestBody(*openapiclient.NewOverrideThresholdRuleRequestBodyFields([]string{"TuningExpressionIds_example"})) // OverrideThresholdRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.OverrideThresholdRule(context.Background(), id).OverrideThresholdRuleRequestBody(overrideThresholdRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.OverrideThresholdRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverrideThresholdRule`: OverrideThresholdRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.OverrideThresholdRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideThresholdRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideThresholdRuleRequestBody** | [**OverrideThresholdRuleRequestBody**](OverrideThresholdRuleRequestBody.md) |  | 

### Return type

[**OverrideThresholdRuleResponse**](OverrideThresholdRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEntityTags

> RemoveEntityTagsResponse RemoveEntityTags(ctx, entityId).RemoveEntityTagsRequestBody(removeEntityTagsRequestBody).Execute()

Remove tags from an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    removeEntityTagsRequestBody := *openapiclient.NewRemoveEntityTagsRequestBody([]string{"Tags_example"}) // RemoveEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RemoveEntityTags(context.Background(), entityId).RemoveEntityTagsRequestBody(removeEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveEntityTags`: RemoveEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RemoveEntityTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeEntityTagsRequestBody** | [**RemoveEntityTagsRequestBody**](RemoveEntityTagsRequestBody.md) |  | 

### Return type

[**RemoveEntityTagsResponse**](RemoveEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveInsightAssignee

> RemoveInsightAssigneeResponse RemoveInsightAssignee(ctx, id).Execute()

Remove the assignee from an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RemoveInsightAssignee(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveInsightAssignee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveInsightAssignee`: RemoveInsightAssigneeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RemoveInsightAssignee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveInsightAssigneeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoveInsightAssigneeResponse**](RemoveInsightAssigneeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSignalsFromInsight

> RemoveSignalsFromInsightResponse RemoveSignalsFromInsight(ctx, insightId).RemoveSignalsFromInsightRequestBody(removeSignalsFromInsightRequestBody).Execute()

Remove existing Signals from an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    insightId := "insightId_example" // string | 
    removeSignalsFromInsightRequestBody := *openapiclient.NewRemoveSignalsFromInsightRequestBody([]string{"SignalIds_example"}) // RemoveSignalsFromInsightRequestBody | A list of Signal IDs to be removed from the Insight

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RemoveSignalsFromInsight(context.Background(), insightId).RemoveSignalsFromInsightRequestBody(removeSignalsFromInsightRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveSignalsFromInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSignalsFromInsight`: RemoveSignalsFromInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RemoveSignalsFromInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSignalsFromInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeSignalsFromInsightRequestBody** | [**RemoveSignalsFromInsightRequestBody**](RemoveSignalsFromInsightRequestBody.md) | A list of Signal IDs to be removed from the Insight | 

### Return type

[**RemoveSignalsFromInsightResponse**](RemoveSignalsFromInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTagFromInsight

> RemoveTagFromInsightResponse RemoveTagFromInsight(ctx, id, tagName).Execute()

Remove a tag from an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    tagName := "tagName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.RemoveTagFromInsight(context.Background(), id, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RemoveTagFromInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTagFromInsight`: RemoveTagFromInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RemoveTagFromInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**tagName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTagFromInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoveTagFromInsightResponse**](RemoveTagFromInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEntityEnrichment

> SaveEntityEnrichmentResponse SaveEntityEnrichment(ctx, id, enrichmentType).SaveEntityEnrichmentRequestBody(saveEntityEnrichmentRequestBody).Execute()

Create or update an Enrichment on an Entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    enrichmentType := "enrichmentType_example" // string | 
    saveEntityEnrichmentRequestBody := *openapiclient.NewSaveEntityEnrichmentRequestBody(map[string]interface{}{"key": interface{}(123)}) // SaveEntityEnrichmentRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SaveEntityEnrichment(context.Background(), id, enrichmentType).SaveEntityEnrichmentRequestBody(saveEntityEnrichmentRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveEntityEnrichment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveEntityEnrichment`: SaveEntityEnrichmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SaveEntityEnrichment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**enrichmentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveEntityEnrichmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saveEntityEnrichmentRequestBody** | [**SaveEntityEnrichmentRequestBody**](SaveEntityEnrichmentRequestBody.md) |  | 

### Return type

[**SaveEntityEnrichmentResponse**](SaveEntityEnrichmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveInsightEnrichment

> SaveInsightEnrichmentResponse SaveInsightEnrichment(ctx, id, enrichmentType).SaveInsightEnrichmentRequestBody(saveInsightEnrichmentRequestBody).Execute()

Create or update an Enrichment on an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    enrichmentType := "enrichmentType_example" // string | 
    saveInsightEnrichmentRequestBody := *openapiclient.NewSaveInsightEnrichmentRequestBody(map[string]interface{}{"key": interface{}(123)}) // SaveInsightEnrichmentRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SaveInsightEnrichment(context.Background(), id, enrichmentType).SaveInsightEnrichmentRequestBody(saveInsightEnrichmentRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveInsightEnrichment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveInsightEnrichment`: SaveInsightEnrichmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SaveInsightEnrichment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**enrichmentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveInsightEnrichmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saveInsightEnrichmentRequestBody** | [**SaveInsightEnrichmentRequestBody**](SaveInsightEnrichmentRequestBody.md) |  | 

### Return type

[**SaveInsightEnrichmentResponse**](SaveInsightEnrichmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSignalEnrichment

> SaveSignalEnrichmentResponse SaveSignalEnrichment(ctx, id, enrichmentType).SaveSignalEnrichmentRequestBody(saveSignalEnrichmentRequestBody).Execute()

Create or update an Enrichment on a Signal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    enrichmentType := "enrichmentType_example" // string | 
    saveSignalEnrichmentRequestBody := *openapiclient.NewSaveSignalEnrichmentRequestBody(map[string]interface{}{"key": interface{}(123)}) // SaveSignalEnrichmentRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.SaveSignalEnrichment(context.Background(), id, enrichmentType).SaveSignalEnrichmentRequestBody(saveSignalEnrichmentRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SaveSignalEnrichment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveSignalEnrichment`: SaveSignalEnrichmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SaveSignalEnrichment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**enrichmentType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveSignalEnrichmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **saveSignalEnrichmentRequestBody** | [**SaveSignalEnrichmentRequestBody**](SaveSignalEnrichmentRequestBody.md) |  | 

### Return type

[**SaveSignalEnrichmentResponse**](SaveSignalEnrichmentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAggregationRule

> UpdateAggregationRuleResponse UpdateAggregationRule(ctx, id).UpdateAggregationRuleRequestBody(updateAggregationRuleRequestBody).Execute()

Update a Aggregation Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateAggregationRuleRequestBody := *openapiclient.NewUpdateAggregationRuleRequestBody(*openapiclient.NewUpdateAggregationRuleRequestBodyFields("Name_example", []openapiclient.GetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner{*openapiclient.NewGetRulesResponseDataObjectsInnerOneOf4AggregationFunctionsInner("Name_example", "Function_example", []string{"Arguments_example"})}, "DescriptionExpression_example", false, []string{"GroupByFields_example"}, "MatchExpression_example", "NameExpression_example", *openapiclient.NewGetRulesResponseDataObjectsInnerOneOf2ScoreMapping("Type_example"), "Stream_example", "TriggerExpression_example", "WindowSize_example")) // UpdateAggregationRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateAggregationRule(context.Background(), id).UpdateAggregationRuleRequestBody(updateAggregationRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAggregationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAggregationRule`: UpdateAggregationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAggregationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAggregationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAggregationRuleRequestBody** | [**UpdateAggregationRuleRequestBody**](UpdateAggregationRuleRequestBody.md) |  | 

### Return type

[**UpdateAggregationRuleResponse**](UpdateAggregationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutomation

> UpdateAutomationResponse UpdateAutomation(ctx, id).UpdateAutomationRequestBody(updateAutomationRequestBody).Execute()

Update an Automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateAutomationRequestBody := *openapiclient.NewUpdateAutomationRequestBody(*openapiclient.NewUpdateAutomationRequestBodyFields()) // UpdateAutomationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateAutomation(context.Background(), id).UpdateAutomationRequestBody(updateAutomationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateAutomation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAutomation`: UpdateAutomationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateAutomation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutomationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAutomationRequestBody** | [**UpdateAutomationRequestBody**](UpdateAutomationRequestBody.md) |  | 

### Return type

[**UpdateAutomationResponse**](UpdateAutomationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChainRule

> UpdateChainRuleResponse UpdateChainRule(ctx, id).UpdateChainRuleRequestBody(updateChainRuleRequestBody).Execute()

Update a Chain Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateChainRuleRequestBody := *openapiclient.NewUpdateChainRuleRequestBody(*openapiclient.NewUpdateChainRuleRequestBodyFields("Name_example", "Description_example", []openapiclient.CreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner{*openapiclient.NewCreateChainRuleRequestBodyFieldsExpressionsAndLimitsInner("Expression_example", int32(123))}, int32(123), "Stream_example", "WindowSize_example")) // UpdateChainRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateChainRule(context.Background(), id).UpdateChainRuleRequestBody(updateChainRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateChainRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChainRule`: UpdateChainRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateChainRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChainRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChainRuleRequestBody** | [**UpdateChainRuleRequestBody**](UpdateChainRuleRequestBody.md) |  | 

### Return type

[**UpdateChainRuleResponse**](UpdateChainRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContextAction

> UpdateContextActionResponse UpdateContextAction(ctx, id).UpdateContextActionRequestBody(updateContextActionRequestBody).Execute()

Update a Context Action



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateContextActionRequestBody := *openapiclient.NewUpdateContextActionRequestBody(*openapiclient.NewUpdateContextActionRequestBodyFields("Name_example", []string{"IocTypes_example"})) // UpdateContextActionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateContextAction(context.Background(), id).UpdateContextActionRequestBody(updateContextActionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateContextAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContextAction`: UpdateContextActionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateContextAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContextActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContextActionRequestBody** | [**UpdateContextActionRequestBody**](UpdateContextActionRequestBody.md) |  | 

### Return type

[**UpdateContextActionResponse**](UpdateContextActionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomEntityType

> UpdateCustomEntityTypeResponse UpdateCustomEntityType(ctx, id).UpdateCustomEntityTypeRequestBody(updateCustomEntityTypeRequestBody).Execute()

Update a Custom Entity Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateCustomEntityTypeRequestBody := *openapiclient.NewUpdateCustomEntityTypeRequestBody(*openapiclient.NewUpdateCustomEntityTypeRequestBodyFields("Name_example", []string{"Fields_example"})) // UpdateCustomEntityTypeRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateCustomEntityType(context.Background(), id).UpdateCustomEntityTypeRequestBody(updateCustomEntityTypeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCustomEntityType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomEntityType`: UpdateCustomEntityTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCustomEntityType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomEntityTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomEntityTypeRequestBody** | [**UpdateCustomEntityTypeRequestBody**](UpdateCustomEntityTypeRequestBody.md) |  | 

### Return type

[**UpdateCustomEntityTypeResponse**](UpdateCustomEntityTypeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomInsight

> UpdateCustomInsightResponse UpdateCustomInsight(ctx, id).UpdateCustomInsightRequestBody(updateCustomInsightRequestBody).Execute()

Update a Custom Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateCustomInsightRequestBody := *openapiclient.NewUpdateCustomInsightRequestBody(*openapiclient.NewCreateCustomInsightRequestBodyFields("Name_example", "Description_example", "Severity_example", false, false, []string{"Tags_example"})) // UpdateCustomInsightRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateCustomInsight(context.Background(), id).UpdateCustomInsightRequestBody(updateCustomInsightRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCustomInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomInsight`: UpdateCustomInsightResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCustomInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomInsightRequestBody** | [**UpdateCustomInsightRequestBody**](UpdateCustomInsightRequestBody.md) |  | 

### Return type

[**UpdateCustomInsightResponse**](UpdateCustomInsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomMatchListColumn

> UpdateCustomMatchListColumnResponse UpdateCustomMatchListColumn(ctx, id).UpdateCustomMatchListColumnRequestBody(updateCustomMatchListColumnRequestBody).Execute()

Update a Custom Match List Column



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateCustomMatchListColumnRequestBody := *openapiclient.NewUpdateCustomMatchListColumnRequestBody(*openapiclient.NewUpdateCustomMatchListColumnRequestBodyFields("Name_example", []string{"Fields_example"})) // UpdateCustomMatchListColumnRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateCustomMatchListColumn(context.Background(), id).UpdateCustomMatchListColumnRequestBody(updateCustomMatchListColumnRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCustomMatchListColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomMatchListColumn`: UpdateCustomMatchListColumnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCustomMatchListColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomMatchListColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomMatchListColumnRequestBody** | [**UpdateCustomMatchListColumnRequestBody**](UpdateCustomMatchListColumnRequestBody.md) |  | 

### Return type

[**UpdateCustomMatchListColumnResponse**](UpdateCustomMatchListColumnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSourcedEntityLookupTable

> UpdateCustomerSourcedEntityLookupTableResponse UpdateCustomerSourcedEntityLookupTable(ctx, id).UpdateCustomerSourcedEntityLookupTableRequestBody(updateCustomerSourcedEntityLookupTableRequestBody).Execute()

Update a customer sourced entity lookup table



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateCustomerSourcedEntityLookupTableRequestBody := *openapiclient.NewUpdateCustomerSourcedEntityLookupTableRequestBody(*openapiclient.NewUpdateCustomerSourcedEntityLookupTableRequestBodyFields("SourceCategory_example")) // UpdateCustomerSourcedEntityLookupTableRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateCustomerSourcedEntityLookupTable(context.Background(), id).UpdateCustomerSourcedEntityLookupTableRequestBody(updateCustomerSourcedEntityLookupTableRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateCustomerSourcedEntityLookupTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerSourcedEntityLookupTable`: UpdateCustomerSourcedEntityLookupTableResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateCustomerSourcedEntityLookupTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSourcedEntityLookupTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerSourcedEntityLookupTableRequestBody** | [**UpdateCustomerSourcedEntityLookupTableRequestBody**](UpdateCustomerSourcedEntityLookupTableRequestBody.md) |  | 

### Return type

[**UpdateCustomerSourcedEntityLookupTableResponse**](UpdateCustomerSourcedEntityLookupTableResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityCriticality

> UpdateEntityCriticalityResponse UpdateEntityCriticality(ctx, entityId).UpdateEntityCriticalityRequestBody(updateEntityCriticalityRequestBody).Execute()

Update an Entity's criticality



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    updateEntityCriticalityRequestBody := *openapiclient.NewUpdateEntityCriticalityRequestBody() // UpdateEntityCriticalityRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityCriticality(context.Background(), entityId).UpdateEntityCriticalityRequestBody(updateEntityCriticalityRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityCriticality``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityCriticality`: UpdateEntityCriticalityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityCriticality`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityCriticalityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityCriticalityRequestBody** | [**UpdateEntityCriticalityRequestBody**](UpdateEntityCriticalityRequestBody.md) |  | 

### Return type

[**UpdateEntityCriticalityResponse**](UpdateEntityCriticalityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityCriticalityConfig

> UpdateEntityCriticalityConfigResponse UpdateEntityCriticalityConfig(ctx, id).UpdateEntityCriticalityConfigRequestBody(updateEntityCriticalityConfigRequestBody).Execute()

Update an Entity Criticality Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateEntityCriticalityConfigRequestBody := *openapiclient.NewUpdateEntityCriticalityConfigRequestBody(*openapiclient.NewUpdateEntityCriticalityConfigRequestBodyFields("SeverityExpression_example")) // UpdateEntityCriticalityConfigRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityCriticalityConfig(context.Background(), id).UpdateEntityCriticalityConfigRequestBody(updateEntityCriticalityConfigRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityCriticalityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityCriticalityConfig`: UpdateEntityCriticalityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityCriticalityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityCriticalityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityCriticalityConfigRequestBody** | [**UpdateEntityCriticalityConfigRequestBody**](UpdateEntityCriticalityConfigRequestBody.md) |  | 

### Return type

[**UpdateEntityCriticalityConfigResponse**](UpdateEntityCriticalityConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityDomainConfiguration

> UpdateEntityDomainConfigurationResponse UpdateEntityDomainConfiguration(ctx).UpdateEntityDomainConfigurationRequestBody(updateEntityDomainConfigurationRequestBody).Execute()

Update the Entity Domain Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateEntityDomainConfigurationRequestBody := *openapiclient.NewUpdateEntityDomainConfigurationRequestBody(*openapiclient.NewGetEntityDomainConfigurationResponseData(false, false, []openapiclient.GetEntityDomainConfigurationResponseDataDomainMappingsInner{*openapiclient.NewGetEntityDomainConfigurationResponseDataDomainMappingsInner("RawDomain_example", "NormalizedDomain_example")}, false, false, false)) // UpdateEntityDomainConfigurationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityDomainConfiguration(context.Background()).UpdateEntityDomainConfigurationRequestBody(updateEntityDomainConfigurationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityDomainConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityDomainConfiguration`: UpdateEntityDomainConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityDomainConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityDomainConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntityDomainConfigurationRequestBody** | [**UpdateEntityDomainConfigurationRequestBody**](UpdateEntityDomainConfigurationRequestBody.md) |  | 

### Return type

[**UpdateEntityDomainConfigurationResponse**](UpdateEntityDomainConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityNote

> UpdateEntityNoteResponse UpdateEntityNote(ctx, entityId, entityNoteId).UpdateEntityNoteRequestBody(updateEntityNoteRequestBody).Execute()

Update an existing Entity Note



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    entityNoteId := "entityNoteId_example" // string | 
    updateEntityNoteRequestBody := *openapiclient.NewUpdateEntityNoteRequestBody("Note_example") // UpdateEntityNoteRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityNote(context.Background(), entityId, entityNoteId).UpdateEntityNoteRequestBody(updateEntityNoteRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityNote`: UpdateEntityNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 
**entityNoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEntityNoteRequestBody** | [**UpdateEntityNoteRequestBody**](UpdateEntityNoteRequestBody.md) |  | 

### Return type

[**UpdateEntityNoteResponse**](UpdateEntityNoteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitySuppressed

> UpdateEntitySuppressedResponse UpdateEntitySuppressed(ctx, entityId).UpdateEntitySuppressedRequestBody(updateEntitySuppressedRequestBody).Execute()

Suppress or un-suppress an Entity'



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    updateEntitySuppressedRequestBody := *openapiclient.NewUpdateEntitySuppressedRequestBody(false) // UpdateEntitySuppressedRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntitySuppressed(context.Background(), entityId).UpdateEntitySuppressedRequestBody(updateEntitySuppressedRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntitySuppressed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntitySuppressed`: UpdateEntitySuppressedResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntitySuppressed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitySuppressedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntitySuppressedRequestBody** | [**UpdateEntitySuppressedRequestBody**](UpdateEntitySuppressedRequestBody.md) |  | 

### Return type

[**UpdateEntitySuppressedResponse**](UpdateEntitySuppressedResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityTags

> UpdateEntityTagsResponse UpdateEntityTags(ctx, entityId).UpdateEntityTagsRequestBody(updateEntityTagsRequestBody).Execute()

Update an Entity's tags, replacing any existing tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityId := "entityId_example" // string | 
    updateEntityTagsRequestBody := *openapiclient.NewUpdateEntityTagsRequestBody([]string{"Tags_example"}) // UpdateEntityTagsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityTags(context.Background(), entityId).UpdateEntityTagsRequestBody(updateEntityTagsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityTags`: UpdateEntityTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityTagsRequestBody** | [**UpdateEntityTagsRequestBody**](UpdateEntityTagsRequestBody.md) |  | 

### Return type

[**UpdateEntityTagsResponse**](UpdateEntityTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityValue

> UpdateEntityValueResponse UpdateEntityValue(ctx, id).UpdateEntityValueRequestBody(updateEntityValueRequestBody).Execute()

Update an Entity Value Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateEntityValueRequestBody := *openapiclient.NewUpdateEntityValueRequestBody(*openapiclient.NewCreateEntityValueRequestBodyFields("Name_example", "ConfigurationType_example")) // UpdateEntityValueRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateEntityValue(context.Background(), id).UpdateEntityValueRequestBody(updateEntityValueRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateEntityValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityValue`: UpdateEntityValueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateEntityValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityValueRequestBody** | [**UpdateEntityValueRequestBody**](UpdateEntityValueRequestBody.md) |  | 

### Return type

[**UpdateEntityValueResponse**](UpdateEntityValueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirstSeenRule

> UpdateFirstSeenRuleResponse UpdateFirstSeenRule(ctx, id).UpdateFirstSeenRuleRequestBody(updateFirstSeenRuleRequestBody).Execute()

Update a First Seen Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateFirstSeenRuleRequestBody := *openapiclient.NewUpdateFirstSeenRuleRequestBody(*openapiclient.NewUpdateFirstSeenRuleRequestBodyFields("Name_example", "DescriptionExpression_example", "NameExpression_example", "FilterExpression_example", []string{"GroupByFields_example"}, int32(123), int32(123), "BaselineWindowSize_example", "RetentionWindowSize_example")) // UpdateFirstSeenRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateFirstSeenRule(context.Background(), id).UpdateFirstSeenRuleRequestBody(updateFirstSeenRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateFirstSeenRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFirstSeenRule`: UpdateFirstSeenRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateFirstSeenRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirstSeenRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirstSeenRuleRequestBody** | [**UpdateFirstSeenRuleRequestBody**](UpdateFirstSeenRuleRequestBody.md) |  | 

### Return type

[**UpdateFirstSeenRuleResponse**](UpdateFirstSeenRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightAssignee

> UpdateInsightAssigneeResponse UpdateInsightAssignee(ctx, id).UpdateInsightAssigneeRequestBody(updateInsightAssigneeRequestBody).Execute()

Update the assignee of an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateInsightAssigneeRequestBody := *openapiclient.NewUpdateInsightAssigneeRequestBody() // UpdateInsightAssigneeRequestBody | The \"type\" of the \"assignee\" should be either \"TEAM\" or \"USER\", and the \"value\" of the \"assignee\" should be the username or team name of the given user/team to be assigned.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightAssignee(context.Background(), id).UpdateInsightAssigneeRequestBody(updateInsightAssigneeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightAssignee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightAssignee`: UpdateInsightAssigneeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightAssignee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightAssigneeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightAssigneeRequestBody** | [**UpdateInsightAssigneeRequestBody**](UpdateInsightAssigneeRequestBody.md) | The \&quot;type\&quot; of the \&quot;assignee\&quot; should be either \&quot;TEAM\&quot; or \&quot;USER\&quot;, and the \&quot;value\&quot; of the \&quot;assignee\&quot; should be the username or team name of the given user/team to be assigned. | 

### Return type

[**UpdateInsightAssigneeResponse**](UpdateInsightAssigneeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightConfiguration

> UpdateInsightConfigurationResponse UpdateInsightConfiguration(ctx).UpdateInsightConfigurationRequestBody(updateInsightConfigurationRequestBody).Execute()

Update Insight Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateInsightConfigurationRequestBody := *openapiclient.NewUpdateInsightConfigurationRequestBody(*openapiclient.NewUpdateInsightConfigurationRequestBodyConfig()) // UpdateInsightConfigurationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightConfiguration(context.Background()).UpdateInsightConfigurationRequestBody(updateInsightConfigurationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightConfiguration`: UpdateInsightConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInsightConfigurationRequestBody** | [**UpdateInsightConfigurationRequestBody**](UpdateInsightConfigurationRequestBody.md) |  | 

### Return type

[**UpdateInsightConfigurationResponse**](UpdateInsightConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightResolution

> UpdateInsightResolutionResponse UpdateInsightResolution(ctx, id).UpdateInsightResolutionRequestBody(updateInsightResolutionRequestBody).Execute()

Update a Insight Resolution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 
    updateInsightResolutionRequestBody := *openapiclient.NewUpdateInsightResolutionRequestBody() // UpdateInsightResolutionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightResolution(context.Background(), id).UpdateInsightResolutionRequestBody(updateInsightResolutionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightResolution`: UpdateInsightResolutionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightResolutionRequestBody** | [**UpdateInsightResolutionRequestBody**](UpdateInsightResolutionRequestBody.md) |  | 

### Return type

[**UpdateInsightResolutionResponse**](UpdateInsightResolutionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightSeverity

> UpdateInsightSeverityResponse UpdateInsightSeverity(ctx, id).UpdateInsightSeverityRequestBody(updateInsightSeverityRequestBody).Execute()

Update the severity of an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateInsightSeverityRequestBody := *openapiclient.NewUpdateInsightSeverityRequestBody() // UpdateInsightSeverityRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightSeverity(context.Background(), id).UpdateInsightSeverityRequestBody(updateInsightSeverityRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightSeverity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightSeverity`: UpdateInsightSeverityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightSeverity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightSeverityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightSeverityRequestBody** | [**UpdateInsightSeverityRequestBody**](UpdateInsightSeverityRequestBody.md) |  | 

### Return type

[**UpdateInsightSeverityResponse**](UpdateInsightSeverityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightStatus

> UpdateInsightStatusResponse UpdateInsightStatus(ctx, id).UpdateInsightStatusRequestBody(updateInsightStatusRequestBody).Execute()

Update the status of an Insight



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateInsightStatusRequestBody := *openapiclient.NewUpdateInsightStatusRequestBody() // UpdateInsightStatusRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightStatus(context.Background(), id).UpdateInsightStatusRequestBody(updateInsightStatusRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightStatus`: UpdateInsightStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightStatusRequestBody** | [**UpdateInsightStatusRequestBody**](UpdateInsightStatusRequestBody.md) |  | 

### Return type

[**UpdateInsightStatusResponse**](UpdateInsightStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightStatusOption

> UpdateInsightStatusOptionResponse UpdateInsightStatusOption(ctx, id).UpdateInsightStatusOptionRequestBody(updateInsightStatusOptionRequestBody).Execute()

Update an Insight Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateInsightStatusOptionRequestBody := *openapiclient.NewUpdateInsightStatusOptionRequestBody() // UpdateInsightStatusOptionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInsightStatusOption(context.Background(), id).UpdateInsightStatusOptionRequestBody(updateInsightStatusOptionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInsightStatusOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInsightStatusOption`: UpdateInsightStatusOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInsightStatusOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInsightStatusOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInsightStatusOptionRequestBody** | [**UpdateInsightStatusOptionRequestBody**](UpdateInsightStatusOptionRequestBody.md) |  | 

### Return type

[**UpdateInsightStatusOptionResponse**](UpdateInsightStatusOptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInventory

> UpdateInventoryResponse UpdateInventory(ctx, id).UpdateInventoryRequestBody(updateInventoryRequestBody).Execute()

Update an Inventory Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateInventoryRequestBody := *openapiclient.NewUpdateInventoryRequestBody(*openapiclient.NewCreateInventoryRequestBodyFields("Name_example", "ConfigurationType_example", "InventorySource_example")) // UpdateInventoryRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateInventory(context.Background(), id).UpdateInventoryRequestBody(updateInventoryRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInventory`: UpdateInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInventoryRequestBody** | [**UpdateInventoryRequestBody**](UpdateInventoryRequestBody.md) |  | 

### Return type

[**UpdateInventoryResponse**](UpdateInventoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogMapping

> UpdateLogMappingResponse UpdateLogMapping(ctx, id).UpdateLogMappingRequestBody(updateLogMappingRequestBody).Execute()

Update a Log Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateLogMappingRequestBody := *openapiclient.NewUpdateLogMappingRequestBody(*openapiclient.NewUpdateLogMappingRequestBodyFields("Name_example", []openapiclient.CreateLogMappingRequestBodyFieldsFieldsInner{*openapiclient.NewCreateLogMappingRequestBodyFieldsFieldsInner("Name_example")}, "RecordType_example", "ProductGuid_example")) // UpdateLogMappingRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateLogMapping(context.Background(), id).UpdateLogMappingRequestBody(updateLogMappingRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateLogMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogMapping`: UpdateLogMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateLogMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogMappingRequestBody** | [**UpdateLogMappingRequestBody**](UpdateLogMappingRequestBody.md) |  | 

### Return type

[**UpdateLogMappingResponse**](UpdateLogMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogMappingEnabled

> UpdateLogMappingEnabledResponse UpdateLogMappingEnabled(ctx, id).UpdateLogMappingEnabledRequestBody(updateLogMappingEnabledRequestBody).Execute()

Update enabled field of a Log Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateLogMappingEnabledRequestBody := *openapiclient.NewUpdateLogMappingEnabledRequestBody(false) // UpdateLogMappingEnabledRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateLogMappingEnabled(context.Background(), id).UpdateLogMappingEnabledRequestBody(updateLogMappingEnabledRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateLogMappingEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogMappingEnabled`: UpdateLogMappingEnabledResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateLogMappingEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogMappingEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogMappingEnabledRequestBody** | [**UpdateLogMappingEnabledRequestBody**](UpdateLogMappingEnabledRequestBody.md) |  | 

### Return type

[**UpdateLogMappingEnabledResponse**](UpdateLogMappingEnabledResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMatchList

> UpdateMatchListResponse UpdateMatchList(ctx, id).UpdateMatchListRequestBody(updateMatchListRequestBody).Execute()

Update a Match List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateMatchListRequestBody := *openapiclient.NewUpdateMatchListRequestBody(*openapiclient.NewUpdateSuppressListRequestBodyFields("Description_example")) // UpdateMatchListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateMatchList(context.Background(), id).UpdateMatchListRequestBody(updateMatchListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMatchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMatchList`: UpdateMatchListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMatchList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMatchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMatchListRequestBody** | [**UpdateMatchListRequestBody**](UpdateMatchListRequestBody.md) |  | 

### Return type

[**UpdateMatchListResponse**](UpdateMatchListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMatchListItem

> UpdateMatchListItemResponse UpdateMatchListItem(ctx, id).UpdateMatchListItemRequestBody(updateMatchListItemRequestBody).Execute()

Update a Match List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateMatchListItemRequestBody := *openapiclient.NewUpdateMatchListItemRequestBody() // UpdateMatchListItemRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateMatchListItem(context.Background(), id).UpdateMatchListItemRequestBody(updateMatchListItemRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMatchListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMatchListItem`: UpdateMatchListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMatchListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMatchListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMatchListItemRequestBody** | [**UpdateMatchListItemRequestBody**](UpdateMatchListItemRequestBody.md) |  | 

### Return type

[**UpdateMatchListItemResponse**](UpdateMatchListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMatchRule

> UpdateMatchRuleResponse UpdateMatchRule(ctx, id).UpdateMatchRuleRequestBody(updateMatchRuleRequestBody).Execute()

Update a Legacy Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateMatchRuleRequestBody := *openapiclient.NewUpdateMatchRuleRequestBody(*openapiclient.NewUpdateMatchRuleRequestBodyFields("Name_example", "Description_example", "Expression_example", int32(123), "Stream_example")) // UpdateMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateMatchRule(context.Background(), id).UpdateMatchRuleRequestBody(updateMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMatchRule`: UpdateMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateMatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMatchRuleRequestBody** | [**UpdateMatchRuleRequestBody**](UpdateMatchRuleRequestBody.md) |  | 

### Return type

[**UpdateMatchRuleResponse**](UpdateMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkBlock

> UpdateNetworkBlockResponse UpdateNetworkBlock(ctx, id).UpdateNetworkBlockRequestBody(updateNetworkBlockRequestBody).Execute()

Update a Network Block



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateNetworkBlockRequestBody := *openapiclient.NewUpdateNetworkBlockRequestBody(*openapiclient.NewUpdateNetworkBlockRequestBodyFields()) // UpdateNetworkBlockRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateNetworkBlock(context.Background(), id).UpdateNetworkBlockRequestBody(updateNetworkBlockRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateNetworkBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkBlock`: UpdateNetworkBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateNetworkBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkBlockRequestBody** | [**UpdateNetworkBlockRequestBody**](UpdateNetworkBlockRequestBody.md) |  | 

### Return type

[**UpdateNetworkBlockResponse**](UpdateNetworkBlockResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleEnabled

> UpdateRuleEnabledResponse UpdateRuleEnabled(ctx, id).UpdateRuleEnabledRequestBody(updateRuleEnabledRequestBody).Execute()

Enable or disable a Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateRuleEnabledRequestBody := *openapiclient.NewUpdateRuleEnabledRequestBody(false) // UpdateRuleEnabledRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateRuleEnabled(context.Background(), id).UpdateRuleEnabledRequestBody(updateRuleEnabledRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateRuleEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRuleEnabled`: UpdateRuleEnabledResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateRuleEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRuleEnabledRequestBody** | [**UpdateRuleEnabledRequestBody**](UpdateRuleEnabledRequestBody.md) |  | 

### Return type

[**UpdateRuleEnabledResponse**](UpdateRuleEnabledResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleTuningExpression

> UpdateRuleTuningExpressionResponse UpdateRuleTuningExpression(ctx, id).UpdateRuleTuningExpressionRequestBody(updateRuleTuningExpressionRequestBody).Execute()

Update a Rule Tuning Expression



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateRuleTuningExpressionRequestBody := *openapiclient.NewUpdateRuleTuningExpressionRequestBody(*openapiclient.NewCreateRuleTuningExpressionRequestBodyFields("Name_example", "Description_example", "Expression_example", false, false, false, []string{"RuleIds_example"})) // UpdateRuleTuningExpressionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateRuleTuningExpression(context.Background(), id).UpdateRuleTuningExpressionRequestBody(updateRuleTuningExpressionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateRuleTuningExpression``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRuleTuningExpression`: UpdateRuleTuningExpressionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateRuleTuningExpression`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleTuningExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRuleTuningExpressionRequestBody** | [**UpdateRuleTuningExpressionRequestBody**](UpdateRuleTuningExpressionRequestBody.md) |  | 

### Return type

[**UpdateRuleTuningExpressionResponse**](UpdateRuleTuningExpressionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuppressList

> UpdateSuppressListResponse UpdateSuppressList(ctx, id).UpdateSuppressListRequestBody(updateSuppressListRequestBody).Execute()

Update a Suppress List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateSuppressListRequestBody := *openapiclient.NewUpdateSuppressListRequestBody(*openapiclient.NewUpdateSuppressListRequestBodyFields("Description_example")) // UpdateSuppressListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSuppressList(context.Background(), id).UpdateSuppressListRequestBody(updateSuppressListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSuppressList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSuppressList`: UpdateSuppressListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSuppressList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuppressListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSuppressListRequestBody** | [**UpdateSuppressListRequestBody**](UpdateSuppressListRequestBody.md) |  | 

### Return type

[**UpdateSuppressListResponse**](UpdateSuppressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSuppressListItem

> UpdateSuppressListItemResponse UpdateSuppressListItem(ctx, id).UpdateSuppressListItemRequestBody(updateSuppressListItemRequestBody).Execute()

Update a Suppress List Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateSuppressListItemRequestBody := *openapiclient.NewUpdateSuppressListItemRequestBody() // UpdateSuppressListItemRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateSuppressListItem(context.Background(), id).UpdateSuppressListItemRequestBody(updateSuppressListItemRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateSuppressListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSuppressListItem`: UpdateSuppressListItemResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateSuppressListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSuppressListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSuppressListItemRequestBody** | [**UpdateSuppressListItemRequestBody**](UpdateSuppressListItemRequestBody.md) |  | 

### Return type

[**UpdateSuppressListItemResponse**](UpdateSuppressListItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagSchema

> UpdateTagSchemaResponse UpdateTagSchema(ctx).UpdateTagSchemaRequestBody(updateTagSchemaRequestBody).Execute()

Update a Tag Schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateTagSchemaRequestBody := *openapiclient.NewUpdateTagSchemaRequestBody(*openapiclient.NewUpdateTagSchemaRequestBodyFields("Key_example", "Label_example", false)) // UpdateTagSchemaRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateTagSchema(context.Background()).UpdateTagSchemaRequestBody(updateTagSchemaRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTagSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagSchema`: UpdateTagSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTagSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTagSchemaRequestBody** | [**UpdateTagSchemaRequestBody**](UpdateTagSchemaRequestBody.md) |  | 

### Return type

[**UpdateTagSchemaResponse**](UpdateTagSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplatedMatchRule

> UpdateTemplatedMatchRuleResponse UpdateTemplatedMatchRule(ctx, id).UpdateTemplatedMatchRuleRequestBody(updateTemplatedMatchRuleRequestBody).Execute()

Update a Match Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateTemplatedMatchRuleRequestBody := *openapiclient.NewUpdateTemplatedMatchRuleRequestBody(*openapiclient.NewUpdateTemplatedMatchRuleRequestBodyFields("Name_example", "DescriptionExpression_example", "Expression_example", "NameExpression_example", *openapiclient.NewGetRulesResponseDataObjectsInnerOneOf2ScoreMapping("Type_example"), "Stream_example")) // UpdateTemplatedMatchRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateTemplatedMatchRule(context.Background(), id).UpdateTemplatedMatchRuleRequestBody(updateTemplatedMatchRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateTemplatedMatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTemplatedMatchRule`: UpdateTemplatedMatchRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateTemplatedMatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplatedMatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTemplatedMatchRuleRequestBody** | [**UpdateTemplatedMatchRuleRequestBody**](UpdateTemplatedMatchRuleRequestBody.md) |  | 

### Return type

[**UpdateTemplatedMatchRuleResponse**](UpdateTemplatedMatchRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThreatIntelIndicator

> UpdateThreatIntelIndicatorResponse UpdateThreatIntelIndicator(ctx, id).UpdateThreatIntelIndicatorRequestBody(updateThreatIntelIndicatorRequestBody).Execute()

Update a Threat Intel Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateThreatIntelIndicatorRequestBody := *openapiclient.NewUpdateThreatIntelIndicatorRequestBody() // UpdateThreatIntelIndicatorRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateThreatIntelIndicator(context.Background(), id).UpdateThreatIntelIndicatorRequestBody(updateThreatIntelIndicatorRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateThreatIntelIndicator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateThreatIntelIndicator`: UpdateThreatIntelIndicatorResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateThreatIntelIndicator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThreatIntelIndicatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateThreatIntelIndicatorRequestBody** | [**UpdateThreatIntelIndicatorRequestBody**](UpdateThreatIntelIndicatorRequestBody.md) |  | 

### Return type

[**UpdateThreatIntelIndicatorResponse**](UpdateThreatIntelIndicatorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThresholdRule

> UpdateThresholdRuleResponse UpdateThresholdRule(ctx, id).UpdateThresholdRuleRequestBody(updateThresholdRuleRequestBody).Execute()

Update a Threshold Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    updateThresholdRuleRequestBody := *openapiclient.NewUpdateThresholdRuleRequestBody(*openapiclient.NewUpdateThresholdRuleRequestBodyFields("Name_example", "Description_example", false, "Expression_example", int32(123), int32(123), "Stream_example", int32(123), "WindowSize_example")) // UpdateThresholdRuleRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.UpdateThresholdRule(context.Background(), id).UpdateThresholdRuleRequestBody(updateThresholdRuleRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateThresholdRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateThresholdRule`: UpdateThresholdRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateThresholdRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThresholdRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateThresholdRuleRequestBody** | [**UpdateThresholdRuleRequestBody**](UpdateThresholdRuleRequestBody.md) |  | 

### Return type

[**UpdateThresholdRuleResponse**](UpdateThresholdRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

