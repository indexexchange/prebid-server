package aerospike_fetcher

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/prebid/prebid-server/config"
	"github.com/stretchr/testify/assert"
	"gitlab.indexexchange.com/exchange-node/rules-lib/aerospike"
	"gitlab.indexexchange.com/exchange-node/rules-lib/providers"
	"gitlab.indexexchange.com/exchange-node/schema/rulesmodel"
)

var (
	testStoredRequest = &rulesmodel.StoredRequest{
		Config: "{'key': 'value'}",
	}
	testUUID               = "test"
	testAccountID          = "test"
	testPrimaryAdServerKey = "1_2_3"
	mockValueBinName       = "value"
	testStoredImpression   = &rulesmodel.StoredImpression{
		Config: "{'key': 'value'}",
	}
	testStoredAccount = &rulesmodel.StoredAccount{
		Config: "{'key': 'value'}",
	}
	testPrimaryAdServerCategoryMapping = &rulesmodel.PrimaryAdServerCategoryMapping{
		MappingID: "{'key': 'value'}",
	}
)

const (
	asPrebidStoredRequestsSet           string = "prebid_stored_requests"
	asPrebidStoredImpressionsSet        string = "prebid_stored_impressions"
	asPrebidStoredAccountsSet           string = "prebid_stored_accounts"
	asPrebidAdServerCategoryMappingsSet string = "prebid_ad_server_category_mappings"
)

func getTestFetcher(datatype config.DataType) *aerospikeFetcher {

	mockAerospikeClient := aerospike.NewMockClient()

	mockStoredRequestData, _ := proto.Marshal(testStoredRequest)
	mockStoredRequestRecord := aerospike.MockAerospikeRecord{
		mockValueBinName: mockStoredRequestData,
	}

	mockStoredImpressionData, _ := proto.Marshal(testStoredImpression)
	mockStoredImpressionRecord := aerospike.MockAerospikeRecord{
		mockValueBinName: mockStoredImpressionData,
	}

	mockStoredAccountData, _ := proto.Marshal(testStoredAccount)
	mockStoredAccountRecord := aerospike.MockAerospikeRecord{
		mockValueBinName: mockStoredAccountData,
	}

	mockPrimaryAdServerData, _ := proto.Marshal(testPrimaryAdServerCategoryMapping)
	mockPrimaryAdServerRecord := aerospike.MockAerospikeRecord{
		mockValueBinName: mockPrimaryAdServerData,
	}

	mockAerospikeClient.InsertMockAerospikeRecord("rules", asPrebidStoredRequestsSet, testUUID, mockStoredRequestRecord)
	mockAerospikeClient.InsertMockAerospikeRecord("rules", asPrebidStoredImpressionsSet, testUUID, mockStoredImpressionRecord)
	mockAerospikeClient.InsertMockAerospikeRecord("rules", asPrebidStoredAccountsSet, testAccountID, mockStoredAccountRecord)
	mockAerospikeClient.InsertMockAerospikeRecord("rules", asPrebidAdServerCategoryMappingsSet, testPrimaryAdServerKey, mockPrimaryAdServerRecord)
	storedRequestProvider, _ := providers.NewStoredRequestCachedProvider(mockAerospikeClient)
	storedImpressionProvider, _ := providers.NewStoredImpressionCachedProvider(mockAerospikeClient)
	storedAccountProvider, _ := providers.NewStoredAccountCachedProvider(mockAerospikeClient)
	storedPrimaryCategoryProvider, _ := providers.NewPrimaryAdServerCategoryMappingCachedProvider(mockAerospikeClient)
	aerospikeProviders := &AerospikeProviders{
		StoredRequestProvider:                  storedRequestProvider,
		StoredImpressionProvider:               storedImpressionProvider,
		StoredAccountProvider:                  storedAccountProvider,
		PrimaryAdServerCategoryMappingProvider: storedPrimaryCategoryProvider,
	}
	fetcher := NewFetcher(aerospikeProviders, datatype)
	return fetcher
}

func TestNewFetcher(t *testing.T) {
	fetcher := getTestFetcher(config.RequestDataType)
	assert.NotNil(t, fetcher)
}

func TestEmptyRequestIDs(t *testing.T) {
	fetcher := getTestFetcher(config.RequestDataType)
	storedReqs, storedImps, errs := fetcher.FetchRequests(context.Background(), nil, nil)
	assert.Nil(t, storedReqs)
	assert.Nil(t, storedImps)
	assert.Nil(t, errs)
}

func TestFetchRequests(t *testing.T) {
	fetcher := getTestFetcher(config.RequestDataType)

	testCases := []struct {
		name         string
		Uuid         []string
		expectErr    bool
		expectExists bool
	}{
		{
			name:         "Stored request does not exist",
			Uuid:         []string{"notExist"},
			expectErr:    true,
			expectExists: false,
		},
		{
			name:         "Stored request exists",
			Uuid:         []string{"test"},
			expectErr:    false,
			expectExists: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			storedReqs, storedImps, errs := fetcher.FetchRequests(context.Background(), tc.Uuid, tc.Uuid)

			if tc.expectErr {
				assert.Len(t, errs, 1)
			} else {
				assert.Len(t, errs, 0)
			}

			if tc.expectExists {
				assert.Len(t, storedReqs, 1)
				assert.Len(t, storedImps, 1)
			} else {
				assert.Len(t, storedReqs, 0)
				assert.Len(t, storedImps, 0)
			}
		})
	}
}

func TestFetchAccount(t *testing.T) {
	fetcher := getTestFetcher(config.AccountDataType)

	testCases := []struct {
		name         string
		accountID    string
		expectErr    bool
		expectExists bool
	}{
		{
			name:         "Stored account does not exist",
			accountID:    "notExist",
			expectErr:    true,
			expectExists: false,
		},
		{
			name:         "Stored account does exist",
			accountID:    "test",
			expectErr:    false,
			expectExists: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			storedAccount, errs := fetcher.FetchAccount(context.Background(), tc.accountID)

			if tc.expectErr {
				assert.Len(t, errs, 1)
			} else {
				assert.Len(t, errs, 0)
			}

			if tc.expectExists {
				assert.NotNil(t, storedAccount)
			} else {
				assert.Nil(t, storedAccount)
			}
		})
	}
}

func TestFetchCategories(t *testing.T) {
	fetcher := getTestFetcher(config.CategoryDataType)

	testCases := []struct {
		name              string
		iabCategoryID     string
		primaryAdServerID string
		publisherID       string
		expectErr         bool
		expectExists      bool
	}{
		{
			name:              "Fetch category does not exist",
			iabCategoryID:     "1",
			primaryAdServerID: "2",
			publisherID:       "not_exist",
			expectErr:         true,
			expectExists:      false,
		},
		{
			name:              "Fetch category exists",
			iabCategoryID:     "1",
			primaryAdServerID: "2",
			publisherID:       "3",
			expectErr:         false,
			expectExists:      true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			primaryAdMapping, errs := fetcher.FetchCategories(context.Background(), tc.primaryAdServerID, tc.publisherID, tc.iabCategoryID)

			if tc.expectErr {
				assert.NotNil(t, errs)
			} else {
				assert.Nil(t, errs)
			}

			if tc.expectExists {
				assert.NotEmpty(t, primaryAdMapping)
			} else {
				assert.Empty(t, primaryAdMapping)
			}
		})
	}
}
