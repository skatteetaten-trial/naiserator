package config

import (
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type KafkaTLS struct {
	CAPath          string `json:"ca-path"`
	CertificatePath string `json:"certificate-path"`
	Enabled         bool   `json:"enabled"`
	Insecure        bool   `json:"insecure"`
	PrivateKeyPath  string `json:"private-key-path"`
}

type Kafka struct {
	Brokers      []string `json:"brokers"`
	Enabled      bool     `json:"enabled"`
	LogVerbosity string   `json:"log-verbosity"`
	TLS          KafkaTLS `json:"tls"`
	Topic        string   `json:"topic"`
}

type Log struct {
	Format string `json:"format"`
	Level  string `json:"level"`
}

type Informer struct {
	FullSyncInterval time.Duration `json:"full-sync-interval"`
}

type Synchronizer struct {
	SynchronizationTimeout time.Duration `json:"synchronization-timeout"`
	RolloutTimeout         time.Duration `json:"rollout-timeout"`
	RolloutCheckInterval   time.Duration `json:"rollout-check-interval"`
}

type Features struct {
	Istio                       bool     `json:"istio"`
	Linkerd                     bool     `json:"linkerd"`
	AccessPolicyNotAllowedCIDRs []string `json:"access-policy-not-allowed-cidrs"`
	NativeSecrets               bool     `json:"native-secrets"`
	NetworkPolicy               bool     `json:"network-policy"`
	Vault                       bool     `json:"vault"`
	Jwker                       bool     `json:"jwker"`
	Azurerator                  bool     `json:"azurerator"`
	Kafkarator                  bool     `json:"kafkarator"`
	Digdirator                  bool     `json:"digdirator"`
}

type Securelogs struct {
	FluentdImage         string `json:"fluentd-image"`
	ConfigMapReloadImage string `json:"configmap-reload-image"`
}

type Proxy struct {
	Address string   `json:"address"`
	Exclude []string `json:"exclude"`
}

type Vault struct {
	Address            string `json:"address"`
	InitContainerImage string `json:"init-container-image"`
	AuthPath           string `json:"auth-path"`
	KeyValuePath       string `json:"kv-path"`
}

type GatewayMapping struct {
	DomainSuffix string `json:"domainSuffix"`
	GatewayName  string `json:"gatewayName"`  // Istio
	IngressClass string `json:"ingressClass"` // Nginx
}

type HostAlias struct {
	Host    string `json:"host"`
	Address string `json:"address"`
}

type Ratelimit struct {
	QPS   int `json:"qps"`
	Burst int `json:"burst"`
}

type ServiceEntryHosts struct {
	Azurerator []string `json:"azurerator"`
	Digdirator []string `json:"digdirator"`
	Jwker      []string `json:"jwker"`
}

type VirtualServiceRegistry struct {
	Enabled        bool   `json:"enabled"`
	Namespace      string `json:"namespace"`
	ApplyOnStartup bool   `json:"apply-on-startup"`
}

type Config struct {
	DryRun                            bool                   `json:"dry-run"`
	Bind                              string                 `json:"bind"`
	Informer                          Informer               `json:"informer"`
	Synchronizer                      Synchronizer           `json:"synchronizer"`
	Kubeconfig                        string                 `json:"kubeconfig"`
	ClusterName                       string                 `json:"cluster-name"`
	GoogleProjectId                   string                 `json:"google-project-id"`
	GoogleCloudSQLProxyContainerImage string                 `json:"google-cloud-sql-proxy-container-image"`
	ApiServerIp                       string                 `json:"api-server-ip"`
	VirtualServiceRegistry            VirtualServiceRegistry `json:"virtual-service-registry"`
	Ratelimit                         Ratelimit              `json:"ratelimit"`
	Log                               Log                    `json:"log"`
	Features                          Features               `json:"features"`
	Securelogs                        Securelogs             `json:"securelogs"`
	Proxy                             Proxy                  `json:"proxy"`
	Vault                             Vault                  `json:"vault"`
	Kafka                             Kafka                  `json:"kafka"`
	HostAliases                       []HostAlias            `json:"host-aliases"`
	GatewayMappings                   []GatewayMapping       `json:"gateway-mappings"`
	ServiceEntryHosts                 ServiceEntryHosts      `json:"service-entry-hosts"`
}

const (
	ApiServerIp                          = "api-server-ip"
	Bind                                 = "bind"
	ClusterName                          = "cluster-name"
	DryRun                               = "dry-run"
	FeaturesAccessPolicyNotAllowedCIDRs  = "features.access-policy-not-allowed-cidrs"
	FeaturesAzurerator                   = "features.azurerator"
	FeaturesDigdirator                   = "features.digdirator"
	FeaturesIstio                        = "features.istio"
	FeaturesJwker                        = "features.jwker"
	FeaturesKafkarator                   = "features.kafkarator"
	FeaturesLinkerd                      = "features.linkerd"
	FeaturesNativeSecrets                = "features.native-secrets"
	FeaturesNetworkPolicy                = "features.network-policy"
	FeaturesVault                        = "features.vault"
	GoogleCloudSQLProxyContainerImage    = "google-cloud-sql-proxy-container-image"
	GoogleProjectId                      = "google-project-id"
	InformerFullSynchronizationInterval  = "informer.full-sync-interval"
	KafkaBrokers                         = "kafka.brokers"
	KafkaEnabled                         = "kafka.enabled"
	KafkaLogVerbosity                    = "kafka.log-verbosity"
	KafkaTLSCAPath                       = "kafka.tls.ca-path"
	KafkaTLSCertificatePath              = "kafka.tls.certificate-path"
	KafkaTLSEnabled                      = "kafka.tls.enabled"
	KafkaTLSInsecure                     = "kafka.tls.insecure"
	KafkaTLSPrivateKeyPath               = "kafka.tls.private-key-path"
	KafkaTopic                           = "kafka.topic"
	KubeConfig                           = "kubeconfig"
	ProxyAddress                         = "proxy.address"
	ProxyExclude                         = "proxy.exclude"
	RateLimitBurst                       = "ratelimit.burst"
	RateLimitQPS                         = "ratelimit.qps"
	SecurelogsConfigMapReloadImage       = "securelogs.configmap-reload-image"
	SecurelogsFluentdImage               = "securelogs.fluentd-image"
	ServiceEntryHostsAzurerator          = "service-entry-hosts.azurerator"
	ServiceEntryHostsDigdirator          = "service-entry-hosts.digdirator"
	ServiceEntryHostsJwker               = "service-entry-hosts.jwker"
	SynchronizerRolloutCheckInterval     = "synchronizer.rollout-check-interval"
	SynchronizerRolloutTimeout           = "synchronizer.rollout-timeout"
	SynchronizerSynchronizationTimeout   = "synchronizer.synchronization-timeout"
	VaultAddress                         = "vault.address"
	VaultAuthPath                        = "vault.auth-path"
	VaultInitContainerImage              = "vault.init-container-image"
	VaultKvPath                          = "vault.kv-path"
	VirtualServiceRegistryApplyOnStartup = "virtual-service-registry.apply-on-startup"
	VirtualServiceRegistryEnabled        = "virtual-service-registry.enabled"
	VirtualServiceRegistryNamespace      = "virtual-service-registry.namespace"
)

func bindNAIS() {
	viper.BindEnv(KafkaBrokers, "KAFKA_BROKERS")
	viper.BindEnv(KafkaTLSCAPath, "KAFKA_CA_PATH")
	viper.BindEnv(KafkaTLSCertificatePath, "KAFKA_CERTIFICATE_PATH")
	viper.BindEnv(KafkaTLSPrivateKeyPath, "KAFKA_PRIVATE_KEY_PATH")
}

func init() {
	// Automatically read configuration options from environment variables.
	// i.e. --proxy.address will be configurable using NAISERATOR_PROXY_ADDRESS.
	viper.SetEnvPrefix("NAISERATOR")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// Read configuration file from working directory and/or /etc.
	// File formats supported include JSON, TOML, YAML, HCL, envfile and Java properties config files
	viper.SetConfigName("naiserator")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc")

	// Ensure Kafkarator variables are used
	bindNAIS()

	// Provide command-line flags
	flag.Bool(DryRun, false, "set to true to run without any actual changes to the cluster")
	flag.String(KubeConfig, "", "path to Kubernetes config file")
	flag.String(Bind, "127.0.0.1:8080", "ip:port where http requests are served")
	flag.String(ClusterName, "cluster-name-unconfigured", "cluster name as presented to deployed applications")
	flag.String(GoogleProjectId, "", "GCP project-id to store google service accounts")
	flag.String(GoogleCloudSQLProxyContainerImage, "", "Docker image of Cloud SQL Proxy container")
	flag.String(ApiServerIp, "", "IP to master in GCP, e.g. 172.16.0.2/32 for GCP")
	flag.Bool(FeaturesIstio, false, "enable creation of Istio-specific resources")
	flag.Bool(FeaturesLinkerd, false, "enable creation of Linkerd-specific resources")
	flag.StringSlice(FeaturesAccessPolicyNotAllowedCIDRs, []string{""}, "CIDRs that should not be included within the allowed IP Block rule for network policy")
	flag.Bool(FeaturesNativeSecrets, false, "enable use of native secrets")
	flag.Bool(FeaturesNetworkPolicy, false, "enable creation of network policies")
	flag.Bool(FeaturesVault, false, "enable use of vault secret injection")
	flag.Bool(FeaturesJwker, false, "enable creation of Jwker resources and secret injection")
	flag.Bool(FeaturesAzurerator, false, "enable creation of AzureAdApplication resources and secret injection")
	flag.Bool(FeaturesKafkarator, false, "enable Kafkarator secret injection")
	flag.Bool(FeaturesDigdirator, false, "enable creation of IDPorten client resources and secret injection")

	flag.Duration(InformerFullSynchronizationInterval, time.Duration(30*time.Minute), "how often to run a full synchronization of all applications")

	flag.Int(RateLimitQPS, 20, "how quickly the rate limit burst bucket is filled per second")
	flag.Int(RateLimitBurst, 200, "how many requests to Kubernetes to allow per second")

	flag.Duration(SynchronizerSynchronizationTimeout, time.Duration(5*time.Second), "how long to allow for resource creation on a single application")
	flag.Duration(SynchronizerRolloutCheckInterval, time.Duration(5*time.Second), "how often to check if a deployment has rolled out successfully")
	flag.Duration(SynchronizerRolloutTimeout, time.Duration(5*time.Minute), "how long to keep checking for a successful deployment rollout")

	flag.String(SecurelogsFluentdImage, "", "Docker image used for secure log fluentd sidecar")
	flag.String(SecurelogsConfigMapReloadImage, "", "Docker image used for secure log configmap reload sidecar")

	flag.StringSlice(ServiceEntryHostsAzurerator, []string{}, "list of hosts to output to ServiceEntry for Applications using Azurerator")
	flag.StringSlice(ServiceEntryHostsDigdirator, []string{}, "list of hosts to output to ServiceEntry for Applications using Digdirator")
	flag.StringSlice(ServiceEntryHostsJwker, []string{}, "list of hosts to output to ServiceEntry for Applications using Jwker")

	flag.String(ProxyAddress, "", "HTTPS?_PROXY environment variable injected into containers")
	flag.StringSlice(ProxyExclude, []string{"localhost"}, "list of hosts or domains injected into NO_PROXY environment variable")

	flag.String(VaultAddress, "", "address of the Vault server")
	flag.String(VaultInitContainerImage, "", "Docker image of init container to use to read secrets from Vault")
	flag.String(VaultAuthPath, "", "path to vault kubernetes auth backend")
	flag.String(VaultKvPath, "", "path to Vault KV mount")

	flag.Bool(VirtualServiceRegistryEnabled, false, "group VirtualService routes per domain instead of multiple per application")
	flag.String(VirtualServiceRegistryNamespace, "default", "where to save VirtualService resources in GCP mode")
	flag.Bool(VirtualServiceRegistryApplyOnStartup, false, "update all VirtualService resources before starting reconciler")

	flag.Bool(KafkaEnabled, false, "Enable connection to kafka")
	flag.Bool(KafkaTLSEnabled, false, "Use TLS for connecting to Kafka.")
	flag.Bool(KafkaTLSInsecure, false, "Allow insecure Kafka TLS connections.")
	flag.String(KafkaLogVerbosity, "trace", "Log verbosity for Kafka client.")
	flag.String(KafkaTLSCAPath, "", "Path to Kafka TLS CA certificate.")
	flag.String(KafkaTLSCertificatePath, "", "Path to Kafka TLS certificate.")
	flag.String(KafkaTLSPrivateKeyPath, "", "Path to Kafka TLS private key.")
	flag.String(KafkaTopic, "deploymentEvents", "Kafka topic for deployment status.")
	flag.StringSlice(KafkaBrokers, []string{"localhost:9092"}, "Comma-separated list of Kafka brokers, HOST:PORT.")
}

// Print out all configuration options except secret stuff.
func Print(redacted []string) {
	ok := func(key string) bool {
		for _, forbiddenKey := range redacted {
			if forbiddenKey == key {
				return false
			}
		}
		return true
	}

	var keys sort.StringSlice = viper.AllKeys()

	keys.Sort()
	for _, key := range keys {
		if ok(key) {
			log.Printf("%s: %v", key, viper.Get(key))
		} else {
			log.Printf("%s: ***REDACTED***", key)
		}
	}

}

func decoderHook(dc *mapstructure.DecoderConfig) {
	dc.TagName = "json"
	dc.ErrorUnused = true
}

func New() (*Config, error) {
	var err error
	var cfg Config

	err = viper.ReadInConfig()
	if err != nil {
		if err.(viper.ConfigFileNotFoundError) != err {
			return nil, err
		}
	}

	flag.Parse()

	err = viper.BindPFlags(flag.CommandLine)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg, decoderHook)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
