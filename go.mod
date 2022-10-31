module github.com/prebid/prebid-server

go 1.19

replace gitlab.indexexchange.com/exchange-node/third-party/protobuf-go/protobuf => gitlab.indexexchange.com/exchange-node/third-party/protobuf-go.git v1.25.4

replace gitlab.indexexchange.com/exchange-node/rules-lib => gitlab.indexexchange.com/exchange-node/rules-lib v1.354.0

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/IABTechLab/adscert v0.44.0
	github.com/NYTimes/gziphandler v1.1.1
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/benbjohnson/clock v1.3.0
	github.com/buger/jsonparser v1.1.1
	github.com/chasex/glog v0.0.0-20160217080310-c62392af379c
	github.com/coocood/freecache v1.2.1
	github.com/docker/go-units v0.4.0
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/golang/glog v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab // indirect
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.10.4
	github.com/mitchellh/copystructure v1.2.0
	github.com/prebid/go-gdpr v1.11.0
	github.com/prebid/openrtb/v17 v17.0.0
	github.com/prometheus/client_golang v1.13.0
	github.com/prometheus/client_model v0.2.0
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rs/cors v1.8.2
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
	github.com/vrischmann/go-metrics-influxdb v0.1.1
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/yudai/gojsondiff v1.0.0
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591
	golang.org/x/text v0.3.7
	google.golang.org/grpc v1.46.2
	gopkg.in/evanphx/json-patch.v4 v4.12.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	gitlab.indexexchange.com/exchange-node/rules-lib v1.390.0
	gitlab.indexexchange.com/exchange-node/schema v1.608.0
)

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.2 // indirect
	github.com/aerospike/aerospike-client-go/v5 v5.7.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/prometheus/statsd_exporter v0.22.7 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da // indirect
	gitlab.indexexchange.com/exchange-blocks/schema v1.8.0 // indirect
	gitlab.indexexchange.com/exchange-node/telemetry v0.41.0 // indirect
	gitlab.indexexchange.com/exchange-node/third-party/protobuf-go/protobuf v1.25.5 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
