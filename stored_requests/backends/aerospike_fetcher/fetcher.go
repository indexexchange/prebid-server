package aerospike_fetcher

import (
	"context"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/prebid/prebid-server/config"
	"gitlab.indexexchange.com/exchange-node/rules-lib/providers"
)

type AerospikeProviders struct {
	StoredRequestProvider                  *providers.StoredRequestCachedProvider
	StoredImpressionProvider               *providers.StoredImpressionCachedProvider
	StoredAccountProvider                  *providers.StoredAccountCachedProvider
	PrimaryAdServerCategoryMappingProvider *providers.PrimaryAdServerCategoryMappingCachedProvider
}

func NewFetcher(providers *AerospikeProviders, datatype config.DataType) *aerospikeFetcher {
	if datatype == config.RequestDataType {
		if providers.StoredRequestProvider == nil {
			glog.Infof("The Aerospike stored request fetcher requires an aerospike provider.")
		}
		if providers.StoredImpressionProvider == nil {
			glog.Infof("The Aerospike stored impression fetcher requires an aerospike provider.")
		}
	}

	if datatype == config.AccountDataType && providers.StoredAccountProvider == nil {
		glog.Infof("The Aerospike stored account fetcher requires an aerospike provider.")
	}

	if datatype == config.CategoryDataType && providers.PrimaryAdServerCategoryMappingProvider == nil {
		glog.Infof("The Aerospike primary ad server category mapping fetcher requires an aerospike provider.")
	}

	return &aerospikeFetcher{
		storedRequestProvider:                  providers.StoredRequestProvider,
		storedImpressionProvider:               providers.StoredImpressionProvider,
		storedAccountProvider:                  providers.StoredAccountProvider,
		primaryAdServerCategoryMappingProvider: providers.PrimaryAdServerCategoryMappingProvider,
	}
}

type aerospikeFetcher struct {
	storedRequestProvider                  *providers.StoredRequestCachedProvider
	storedImpressionProvider               *providers.StoredImpressionCachedProvider
	storedAccountProvider                  *providers.StoredAccountCachedProvider
	primaryAdServerCategoryMappingProvider *providers.PrimaryAdServerCategoryMappingCachedProvider
}

func (fetcher *aerospikeFetcher) FetchRequests(ctx context.Context, requestIDs []string, impIDs []string) (map[string]json.RawMessage, map[string]json.RawMessage, []error) {
	if len(requestIDs) < 1 && len(impIDs) < 1 {
		return nil, nil, nil
	}

	storedRequestData := make(map[string]json.RawMessage, len(requestIDs))
	storedImpressionData := make(map[string]json.RawMessage, len(impIDs))

	if fetcher.storedRequestProvider != nil {
		for _, uuid := range requestIDs {
			storedRequest, err := fetcher.storedRequestProvider.GetStoredRequestObjectByUUID(uuid)
			if err != nil {
				glog.Infof("Failed to retrieve Stored Request object: %v", err)
				return nil, nil, []error{err}
			}

			if err != nil {
				glog.Infof("Failed to marshal Stored Request object: %v", err)
				return nil, nil, []error{err}
			}
			storedRequestData[uuid] = json.RawMessage(storedRequest.Config)
		}
	}

	if fetcher.storedImpressionProvider != nil {
		for _, uuid := range impIDs {
			storedImpression, err := fetcher.storedImpressionProvider.GetStoredImpressionObjectByUUID(uuid)
			if err != nil {
				glog.Infof("Failed to retrieve Stored Impression object: %v", err)
				return nil, nil, []error{err}
			}

			if err != nil {
				glog.Infof("Failed to marshal Stored Impression object: %v", err)
				return nil, nil, []error{err}
			}
			storedImpressionData[uuid] = json.RawMessage(storedImpression.Config)
		}
	}

	return storedRequestData, storedImpressionData, nil
}

func (fetcher *aerospikeFetcher) FetchResponses(ctx context.Context, ids []string) (data map[string]json.RawMessage, errs []error) {
	return nil, nil
}

func (fetcher *aerospikeFetcher) FetchAccount(ctx context.Context, accountID string) (json.RawMessage, []error) {
	if accountID == "" {
		return nil, nil
	}

	storedAccount, err := fetcher.storedAccountProvider.GetStoredAccountObjectByID(accountID)

	if err != nil {
		glog.Infof("Failed to retrieve stored account object: %v", err)
		return nil, []error{err}
	}

	return json.RawMessage(storedAccount.Config), nil
}

func (fetcher *aerospikeFetcher) FetchCategories(ctx context.Context, primaryAdServer, publisherId, iabCategory string) (string, error) {
	if primaryAdServer == "" || publisherId == "" || iabCategory == "" {
		return "", nil
	}

	primaryAdServerCategoryMapping, err := fetcher.primaryAdServerCategoryMappingProvider.GetPrimaryAdServerIabCategoryMappingObjectByMappingID(iabCategory, primaryAdServer, publisherId)

	if err != nil {
		glog.Infof("Failed to retrieve categories object: %v", err)
		return "", err
	}

	if err != nil {
		glog.Infof("Failed to marshal category mapping object: %v", err)
		return "", err
	}

	return string(primaryAdServerCategoryMapping.MappingID), nil
}
