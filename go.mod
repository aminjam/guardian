module code.cloudfoundry.org/guardian

go 1.12

require (
	code.cloudfoundry.org/archiver v0.0.0-20180525162158-e135af3d5a2a
	code.cloudfoundry.org/clock v0.0.0-20180518195852-02e53af36e6c
	code.cloudfoundry.org/commandrunner v0.0.0-20180212143422-501fd662150b
	code.cloudfoundry.org/debugserver v0.0.0-20170501225606-70715da12ee9
	code.cloudfoundry.org/garden v0.0.0-20190410122303-a4e51d29c0e5
	code.cloudfoundry.org/grootfs v0.30.1-0.20190603153917-fafcaf9eabf8
	code.cloudfoundry.org/idmapper v0.0.0-20190509093205-06cf188c2966
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/localip v0.0.0-20170223024724-b88ad0dea95c
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/go-winio v0.4.12 // indirect
	github.com/Microsoft/hcsshim v0.6.11 // indirect
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/cloudfoundry/dropsonde v1.0.0
	github.com/cloudfoundry/gosigar v1.1.0
	github.com/containerd/cgroups v0.0.0-20190226200435-dbea6f2bd416 // indirect
	github.com/containerd/console v0.0.0-20181022165439-0650fd9eeb50 // indirect
	github.com/containerd/containerd v1.2.6
	github.com/containerd/continuity v0.0.0-20190426062206-aaeac12a7ffc // indirect
	github.com/containerd/fifo v0.0.0-20180307165137-3d5202aec260 // indirect
	github.com/containerd/go-runc v0.0.0-20180907222934-5a6d9f37cfa3 // indirect
	github.com/containerd/typeurl v0.0.0-20170912152501-f6943554a7e7
	github.com/containers/storage v1.12.8 // indirect
	github.com/coreos/go-systemd v0.0.0-20190212144455-93d5ec2c7f76 // indirect
	github.com/dgryski/go-sip13 v0.0.0-20190329191031-25c5027a8c7b // indirect
	github.com/docker/docker v0.7.3-0.20190329212828-d7ab8ad145fa
	github.com/docker/docker-credential-helpers v0.6.2 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-events v0.0.0-20170721190031-9461782956ad // indirect
	github.com/docker/go-metrics v0.0.0-20181218153428-b84716841b82 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/eapache/go-resiliency v1.1.0
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/gogo/googleapis v1.0.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/gorilla/mux v1.7.2 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/kisielk/gotool v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/mitchellh/copystructure v0.0.0-20170525013902-d23ffcb85de3
	github.com/mitchellh/go-ps v0.0.0-20170309133038-4fdf99ab2936
	github.com/mitchellh/reflectwalk v0.0.0-20170726202117-63d60e9d0dbc // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v1.0.0-rc7
	github.com/opencontainers/runtime-spec v1.0.1
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.3 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/prometheus/procfs v0.0.1 // indirect
	github.com/prometheus/tsdb v0.8.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/st3v/glager v0.3.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20180223013746-33e07d32887e // indirect
	github.com/tedsuo/ifrit v0.0.0-20180410193936-e89a512c3162
	github.com/urfave/cli v1.20.0
	github.com/ventu-io/go-shortid v0.0.0-20171029131806-771a37caa5cf // indirect
	github.com/vishvananda/netlink v1.0.0
	github.com/vishvananda/netns v0.0.0-20171111001504-be1fbeda1936
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	golang.org/x/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/tools v0.0.0-20190603231351-8aaa1484dc10 // indirect
	google.golang.org/genproto v0.0.0-20180518175338-11a468237815 // indirect
	google.golang.org/grpc v1.12.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/docker/distribution => github.com/docker/distribution v0.0.0-20190205005809-0d3efadf0154
