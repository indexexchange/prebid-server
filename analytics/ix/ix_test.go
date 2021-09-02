package ix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/chasex/glog"
	"github.com/prebid/prebid-server/analytics"
	"github.com/prebid/prebid-server/config"
)

const (
	TEST_FILE_NAME = "_LogNotificationsTest.log"
	BIDDER         = "ix"
	TYPE           = "imp"
	BID_ID         = "1234"
	ACCOUNT_ID     = "test"
	TIMESTAMP      = 1099785736
)

func TestIXModuleSuccess(t *testing.T) {
	ixConfig := config.IX{
		Enabled: true,
		Events: config.IXEvents{
			Win: true,
			Imp: false,
		},
		LogOptions: glog.LogOptions{
			File: TEST_FILE_NAME,
		},
	}
	_, err := NewIXModule(ixConfig)
	osErr := os.Remove(TEST_FILE_NAME)

	if osErr != nil {
		t.Errorf("Error: unable to remove file  %v", TEST_FILE_NAME)
	}
	if err != nil {
		t.Errorf("Error: ix module was not able to initalize: %v", err)
	}
}

func TestIXModuleWithInvalidConfig(t *testing.T) {
	ixConfig := config.IX{
		Enabled: true,
		Events: config.IXEvents{
			Win: true,
			Imp: false,
		},
		LogOptions: glog.LogOptions{
			File: "",
		},
	}
	_, err := NewIXModule(ixConfig)
	os.Remove(TEST_FILE_NAME)

	if err == nil {
		t.Error("Error: created ix module with invalid config")
	}
}

func TestLogNotificationEventObjectIgnoresDisabledEvent(t *testing.T) {
	const TYPE = "imp"
	neo := &analytics.NotificationEvent{
		Request: &analytics.EventRequest{
			Bidder:    BIDDER,
			Type:      TYPE,
			BidID:     BID_ID,
			AccountID: ACCOUNT_ID,
			Timestamp: TIMESTAMP,
		},
		Account: &config.Account{
			ID:            "id",
			Disabled:      false,
			CacheTTL:      config.DefaultTTLs{},
			EventsEnabled: false,
			CCPA:          config.AccountCCPA{},
			GDPR:          config.AccountGDPR{},
			DebugAllow:    false,
		},
	}

	logger, err := glog.New(glog.LogOptions{
		File:  TEST_FILE_NAME,
		Flag:  glog.LstdNull,
		Level: glog.Ldebug,
		Mode:  glog.LstdNull,
	})
	if err != nil {
		return
	}
	ixmodule := IXAnalyticsModule{
		logger,
		config.IX{
			Enabled: true,
			Events: config.IXEvents{
				Win: true,
				Imp: false,
			},
			LogOptions: glog.LogOptions{
				File: TEST_FILE_NAME,
			},
		},
	}

	ixmodule.LogNotificationEventObject(neo)
	file, _ := ioutil.ReadFile(TEST_FILE_NAME)

	if len(string(file)) != 0 {
		t.Errorf("Error: created log for disabled type:  %v", TYPE)
	}

	osErr := os.Remove(TEST_FILE_NAME)

	if osErr != nil {
		t.Errorf("Error: unable to remove file  %v", TEST_FILE_NAME)
	}
}

func TestLogNotificationEventObjectCreatesCorrectEntryInLog(t *testing.T) {
	neo := &analytics.NotificationEvent{
		Request: &analytics.EventRequest{
			Bidder:    BIDDER,
			Type:      "",
			BidID:     BID_ID,
			AccountID: ACCOUNT_ID,
			Timestamp: TIMESTAMP,
		},
		Account: &config.Account{
			ID:            "id",
			Disabled:      false,
			CacheTTL:      config.DefaultTTLs{},
			EventsEnabled: false,
			CCPA:          config.AccountCCPA{},
			GDPR:          config.AccountGDPR{},
			DebugAllow:    false,
		},
	}

	os.Remove(TEST_FILE_NAME)
	eventTypes := []analytics.EventType{"win", "imp"}
	for _, event := range eventTypes {
		neo.Request.Type = event
		logger, err := glog.New(glog.LogOptions{
			File:  TEST_FILE_NAME,
			Flag:  glog.LstdNull,
			Level: glog.Ldebug,
			Mode:  glog.LstdNull,
		})
		if err != nil {
			return
		}
		ixmodule := IXAnalyticsModule{
			logger,
			config.IX{
				Enabled: true,
				Events: config.IXEvents{
					Win: true,
					Imp: true,
				},
				LogOptions: glog.LogOptions{
					File: TEST_FILE_NAME,
				},
			},
		}

		ixmodule.LogNotificationEventObject(neo)
		file, _ := ioutil.ReadFile(TEST_FILE_NAME)
		data := IXAnalyticsData{}
		err = json.Unmarshal([]byte(file), &data)

		if err != nil {
			t.Errorf("Error: unable to parse JSON: %v", err)
		}

		if data.EventType != string(neo.Request.Type) {
			t.Errorf("Error: incorrect event logged %v", TYPE)
		}

		osErr := os.Remove(TEST_FILE_NAME)

		if osErr != nil {
			t.Errorf("Error: unable to remove file  %v", TEST_FILE_NAME)
		}
	}
}

func TestJsonifyNotificationEventObject(t *testing.T) {
	neo := &analytics.NotificationEvent{
		Request: &analytics.EventRequest{
			Bidder:    BIDDER,
			Type:      TYPE,
			BidID:     BID_ID,
			AccountID: ACCOUNT_ID,
			Timestamp: TIMESTAMP,
		},
		Account: &config.Account{
			ID:            "id",
			Disabled:      false,
			CacheTTL:      config.DefaultTTLs{},
			EventsEnabled: false,
			CCPA:          config.AccountCCPA{},
			GDPR:          config.AccountGDPR{},
			DebugAllow:    false,
		},
	}
	jsonifiedObject, err := jsonifyNotificationEventObject(neo)

	if err != nil {
		t.Errorf("unable to jsonify event object %v", err)
	}

	switch eventString := jsonifiedObject; {
	case regexp.MustCompile(fmt.Sprintf(`\"bidder\":\"%v\"`, BIDDER)).MatchString(eventString) == false:
		t.Errorf("jsonified event object doesn't contain expected bidder code: %v", BIDDER)
	case regexp.MustCompile(fmt.Sprintf(`\"type\":\"%v\"`, TYPE)).MatchString(eventString) == false:
		t.Errorf("jsonified event object doesn't contain expected type: %v", TYPE)
	case regexp.MustCompile(fmt.Sprintf(`\"bidId\":\"%v\"`, BID_ID)).MatchString(eventString) == false:
		t.Errorf("jsonified event object doesn't contain expected bidId: %v", BID_ID)
	case regexp.MustCompile(fmt.Sprintf(`\"publisher\":\"%v\"`, ACCOUNT_ID)).MatchString(eventString) == false:
		t.Errorf("jsonified event object doesn't contain expected account id: %v", ACCOUNT_ID)
	case regexp.MustCompile(fmt.Sprintf(`\"timestamp\":%v`, TIMESTAMP)).MatchString(eventString) == false:
		t.Errorf("jsonified event object doesn't contain expected timestamp: %v", fmt.Sprint(TIMESTAMP))
	}
}
