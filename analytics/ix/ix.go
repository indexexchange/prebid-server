package ix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/chasex/glog"
	"github.com/prebid/prebid-server/analytics"
	"github.com/prebid/prebid-server/config"
)

const EVENT_VERSION string = "1.0"

//Module that can perform transactional logging
type IXAnalyticsModule struct {
	Logger *glog.Logger
	config config.IX
}

//Format for Data entry in IndexExchangeData
type EventData struct {
	Bidder string `json:"bidder"`
}

//Format for event notification JSON log
type IXAnalyticsData struct {
	EventType string    `json:"type"`
	Version   string    `json:"version"`
	Publisher string    `json:"publisher"`
	BidId     string    `json:"bidId"`
	Timestamp int       `json:"timestamp"`
	Data      EventData `json:"data"`
}

//Method to initialize the analytic module
func NewIXModule(config config.IX) (analytics.PBSAnalyticsModule, error) {

	if len(config.LogOptions.File) == 0 {
		return nil, fmt.Errorf("missing ix.log_options.file configuration")
	}
	if logger, err := glog.New(config.LogOptions); err == nil {
		return &IXAnalyticsModule{
			logger,
			config,
		}, nil
	} else {
		return nil, err
	}
}

//Logs Event to file
func (i *IXAnalyticsModule) LogNotificationEventObject(ne *analytics.NotificationEvent) {
	if ne == nil {
		return
	}

	if !isEventEnabled(ne.Request.Type, i.config.Events) {
		return
	}

	eventObject, err := jsonifyNotificationEventObject(ne)
	if err != nil {
		return
	}
	//Code to parse the object and log in a way required
	var b bytes.Buffer
	b.WriteString(eventObject)
	i.Logger.Info(b.String())
	i.Logger.Flush()
}

//Checks if triggered event is enabled within IX modules configuration
func isEventEnabled(eventTriggered analytics.EventType, configuredEvents config.IXEvents) bool {
	eventType := strings.Title(strings.ToLower(string(eventTriggered)))
	return reflect.Indirect(reflect.ValueOf(configuredEvents)).FieldByName(eventType).Bool()
}

//Converts analytics notification object to standard logging JSON object
func jsonifyNotificationEventObject(ne *analytics.NotificationEvent) (string, error) {
	b, err := json.Marshal(IXAnalyticsData{
		EventType: string(ne.Request.Type),
		Version:   EVENT_VERSION,
		Publisher: ne.Request.AccountID,
		BidId:     ne.Request.BidID,
		Timestamp: int(ne.Request.Timestamp),
		Data:      EventData{Bidder: ne.Request.Bidder},
	})

	if err != nil {
		return "", fmt.Errorf("transactional logs error: notificationEvent object badly formed %v", err)
	}
	return string(b), nil
}

//Writes AuctionObject to file
func (i *IXAnalyticsModule) LogAuctionObject(ao *analytics.AuctionObject) {
	//TODO - Implement as needed in the future
}

//Writes VideoObject to file
func (i *IXAnalyticsModule) LogVideoObject(vo *analytics.VideoObject) {
	//TODO - Implement as needed in the future
}

//Logs SetUIDObject to file
func (i *IXAnalyticsModule) LogSetUIDObject(so *analytics.SetUIDObject) {
	//TODO - Implement as needed in the future
}

//Logs CookieSyncObject to file
func (i *IXAnalyticsModule) LogCookieSyncObject(cso *analytics.CookieSyncObject) {
	//TODO - Implement as needed in the future
}

//Logs AmpObject to file
func (i *IXAnalyticsModule) LogAmpObject(ao *analytics.AmpObject) {
	//TODO - Implement as needed in the future
}
