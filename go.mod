module nRIC

go 1.15

require (
	aqwari.net/xml v0.0.0-20201130015819-c3f085bce04e
	cloud.google.com/go v0.72.0 // indirect
	cloud.google.com/go/storage v1.12.0 // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/apache/thrift v0.13.0 // indirect
	github.com/aryann/difflib v0.0.0-20170710044230-e206f873d14a // indirect
	github.com/aws/aws-lambda-go v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2 v0.18.0 // indirect
	github.com/casbin/casbin/v2 v2.1.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/clbanning/x2j v0.0.0-20191024224557-825249438eec // indirect
	github.com/cncf/udpa/go v0.0.0-20201120205902-5459f2c99403 // indirect
	github.com/cockroachdb/datadriven v0.0.0-20190809214429-80d97fb3cbaa // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/dop251/goja v0.0.0-20210126164150-f5884268f0c0
	github.com/dustin/go-humanize v0.0.0-20171111073723-bb3d318650d4 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/envoyproxy/go-control-plane v0.9.7 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1 // indirect
	github.com/franela/goblin v0.0.0-20200105215937-c9ffbefa60db // indirect
	github.com/franela/goreq v0.0.0-20171204163338-bcd34c9993f8 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-delve/delve v1.5.1 // indirect
	github.com/go-echarts/go-echarts v1.0.0
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-openapi/errors v0.19.9
	github.com/go-openapi/loads v0.20.0
	github.com/go-openapi/runtime v0.19.24
	github.com/go-openapi/spec v0.20.1
	github.com/go-openapi/strfmt v0.20.0
	github.com/go-openapi/swag v0.19.13
	github.com/go-openapi/validate v0.20.1
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.1
	github.com/google/go-cmp v0.5.4
	github.com/google/pprof v0.0.0-20201117184057-ae444373da19 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/hashicorp/consul/api v1.3.0 // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hudl/fargo v1.3.0 // indirect
	github.com/iancoleman/strcase v0.1.2 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d // indirect
	github.com/ivarg/goxsd v0.0.0-20160720221513-9a7af8e4e443 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/klauspost/compress v1.11.4
	github.com/kr/text v0.2.0 // indirect
	github.com/lightstep/lightstep-tracer-go v0.18.1 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.4
	github.com/lyft/protoc-gen-star v0.5.2 // indirect
	github.com/magiconair/properties v1.8.1
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mattn/go-runewidth v0.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/nats-io/nats-server/v2 v2.1.2 // indirect
	github.com/oklog/oklog v0.3.2
	github.com/oklog/run v1.0.0 // indirect
	github.com/olekukonko/tablewriter v0.0.0-20170122224234-a0225b3f23b5 // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/opentracing/basictracer-go v1.0.0 // indirect
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/pact-foundation/pact-go v1.0.4 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/performancecopilot/speed v3.0.0+incompatible // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.12.0 // indirect
	github.com/prometheus/client_golang v1.3.0
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/sony/gobreaker v0.4.1 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271 // indirect
	github.com/streadway/handy v0.0.0-20190108123426-d5acb3125c2a // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c
	github.com/xdg/stringprep v1.0.0 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/mod v0.3.1-0.20200828183125-ce943fd02449 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58 // indirect
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930 // indirect
	golang.org/x/tools v0.0.0-20201201030018-7470481624a7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.25 // indirect
	gopkg.in/gcfg.v1 v1.2.3 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009
	sigs.k8s.io/yaml v1.1.0 // indirect
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0 // indirect
)

replace golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
