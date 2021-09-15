package network_policy

import (
	"sort"

	skatteetaten_no_v1alpha1 "github.com/nais/liberator/pkg/apis/nebula.skatteetaten.no/v1alpha1"
	"github.com/nais/naiserator/pkg/skatteetaten_generator/authorization_policy"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	KubeNamespace = "kube-system"
	MetricsPort   = 15020
	DNSPort       = 53
)

func GenerateNetworkPolicy(application skatteetaten_no_v1alpha1.Application, config skatteetaten_no_v1alpha1.ApplicationSpec) *networkingv1.NetworkPolicy {
	np := generateNetworkPolicy(application)

	// Minimum required policies needed for a pod to start
	np.Spec.Ingress = *generateDefaultIngressRules(application)
	np.Spec.Egress = *generateDefaultEgressRules(application)

	if config.Ingress != nil {
		// Internal ingress
		// Sort to allow fixture testing
		keys := make([]string, 0, len(config.Ingress.Internal))
		for k := range config.Ingress.Internal {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, rule := range keys {
			if config.Ingress.Internal[rule].Enabled {
				np.Spec.Ingress = append(np.Spec.Ingress, *generateNetworkPolicyIngressRule(
					application,
					config.Ingress.Internal[rule]))
			}
		}

		// Public ingress
		for _, ingress := range config.Ingress.Public {
			if ingress.Enabled {
				gateway := ingress.Gateway
				if len(gateway) == 0 {
					gateway = authorization_policy.DefaultIngressGateway
				}

				rule := networkingv1.NetworkPolicyIngressRule{}
				appLabel := map[string]string{
					"app":   gateway,
					"istio": "ingressgateway",
				}

				rule.From = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(application, authorization_policy.IstioNamespace, appLabel)}
				rule.Ports = *generateNetworkPolicyPorts([]skatteetaten_no_v1alpha1.PortConfig{{Port: uint16(ingress.Port), Protocol: "TCP"}})
				np.Spec.Ingress = append(np.Spec.Ingress, rule)
			}
		}
	}

	if config.Egress != nil {
		// Internal egress
		// Sort to allow fixture testing
		keys := make([]string, 0, len(config.Egress.Internal))
		for k := range config.Egress.Internal {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, rule := range keys {
			if config.Egress.Internal[rule].Enabled {
				np.Spec.Egress = append(
					np.Spec.Egress, *generateNetworkPolicyEgressRule(
						application,
						config.Egress.Internal[rule]))
			}
		}

		// External egress
		if len(config.Egress.External) > 0 {
			np.Spec.Egress = append(np.Spec.Egress, *generateNetworkPolicyExternalEgressRule())
		}
	}

	return np
}

func generateNetworkPolicy(application skatteetaten_no_v1alpha1.Application) *networkingv1.NetworkPolicy {
	return &networkingv1.NetworkPolicy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.k8s.io/v1",
			Kind:       "NetworkPolicy",
		},
		ObjectMeta: application.StandardObjectMeta(),
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{"app": application.Name},
			},
		},
	}
}

func generateDefaultIngressRules(application skatteetaten_no_v1alpha1.Application) *[]networkingv1.NetworkPolicyIngressRule {
	var ruleList []networkingv1.NetworkPolicyIngressRule

	// Allow prometheus scraping on the "merged metrics" port on the istio proxy.
	// Istio proxy collects metrics from the app on the configured metrics port and merges with own metrics.
	rule := networkingv1.NetworkPolicyIngressRule{}
	rule.From = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(
		application,
		authorization_policy.IstioNamespace,
		map[string]string{"app": "prometheus", "component": "server"})}
	rule.Ports = *generateNetworkPolicyPorts([]skatteetaten_no_v1alpha1.PortConfig{{Protocol: "TCP", Port: MetricsPort}})
	ruleList = append(ruleList, rule)

	return &ruleList
}

func generateDefaultEgressRules(application skatteetaten_no_v1alpha1.Application) *[]networkingv1.NetworkPolicyEgressRule {
	var ruleList []networkingv1.NetworkPolicyEgressRule
	rule := networkingv1.NetworkPolicyEgressRule{}

	// Allow access to kube-dns
	rule.To = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(
		application,
		KubeNamespace,
		map[string]string{"k8s-app": "kube-dns"})}
	rule.Ports = *generateNetworkPolicyPorts([]skatteetaten_no_v1alpha1.PortConfig{{Protocol: "UDP", Port: DNSPort}})
	ruleList = append(ruleList, rule)

	// Seems like kube-dns isn't enough. And I am not sure why, but some investigation
	// suggests it is only required when starting up sidecar/init-containers in AKS.
	rule = networkingv1.NetworkPolicyEgressRule{}
	rule.Ports = *generateNetworkPolicyPorts([]skatteetaten_no_v1alpha1.PortConfig{{Protocol: "UDP", Port: DNSPort}})
	ruleList = append(ruleList, rule)

	// Istio Proxy needs access to Istio pilot.
	// TODO: Limit on specific ports.
	rule = networkingv1.NetworkPolicyEgressRule{}
	rule.To = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(
		application,
		authorization_policy.IstioNamespace,
		map[string]string{"app": "istiod", "istio": "pilot"})}
	ruleList = append(ruleList, rule)

	// This is needed to reach the cluster's metadata server (169.254.169.254).
	// It's reachable through localhost, so why we need the rule at all is weird.
	rule = networkingv1.NetworkPolicyEgressRule{}
	peer := networkingv1.NetworkPolicyPeer{}
	peer.IPBlock = &networkingv1.IPBlock{CIDR: "127.0.0.1/32"}
	rule.To = []networkingv1.NetworkPolicyPeer{peer}
	ruleList = append(ruleList, rule)

	return &ruleList
}

func generateNetworkPolicyIngressRule(application skatteetaten_no_v1alpha1.Application, inbound skatteetaten_no_v1alpha1.InternalIngressConfig) *networkingv1.NetworkPolicyIngressRule {
	rule := networkingv1.NetworkPolicyIngressRule{}
	appLabel := map[string]string{}

	if inbound.Application != "*" && inbound.Application != "" {
		appLabel["app"] = inbound.Application
	}

	rule.From = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(application, inbound.Namespace, appLabel)}
	rule.Ports = *generateNetworkPolicyPorts(inbound.Ports)

	return &rule
}

func generateNetworkPolicyEgressRule(application skatteetaten_no_v1alpha1.Application, outbound skatteetaten_no_v1alpha1.InternalEgressConfig) *networkingv1.NetworkPolicyEgressRule {
	rule := networkingv1.NetworkPolicyEgressRule{}
	appLabel := map[string]string{}

	if outbound.Application != "*" && outbound.Application != "" {
		appLabel["app"] = outbound.Application
	}

	rule.To = []networkingv1.NetworkPolicyPeer{*generateNetworkPolicyPeer(
		application,
		outbound.Namespace,
		appLabel)}
	rule.Ports = *generateNetworkPolicyPorts(outbound.Ports)

	return &rule
}

func generateNetworkPolicyExternalEgressRule() *networkingv1.NetworkPolicyEgressRule {
	// The Calico version on AKS only supports IP based rules for external hosts. (Calico enterprise
	// supports hostname based filtering). Doing IP-based filtering is not a viable solution, so to
	// allow any external traffic we need accept all. However, we can still force use of Network
	// Policies for any internal traffic. For external egress we use Istio ServiceEntry to handle
	// filtering in Istio. Note that egress has to be configured in Azure firewall (NSG) as well.
	return &networkingv1.NetworkPolicyEgressRule{
		To: []networkingv1.NetworkPolicyPeer{{
			IPBlock: &networkingv1.IPBlock{
				CIDR: "0.0.0.0/0",
				Except: []string{
					"10.0.0.0/8",
					"172.16.0.0/12",
					"192.168.0.0/16",
				}},
		},
		},
	}
}

func generateNetworkPolicyPeer(application skatteetaten_no_v1alpha1.Application, namespace string, appLabel map[string]string) *networkingv1.NetworkPolicyPeer {
	peer := networkingv1.NetworkPolicyPeer{}

	if len(namespace) == 0 {
		namespace = application.Namespace
	}

	peer.NamespaceSelector = &metav1.LabelSelector{
		MatchLabels: map[string]string{"name": namespace},
	}

	if len(appLabel) > 0 {
		peer.PodSelector = &metav1.LabelSelector{
			MatchLabels: appLabel,
		}
	}

	return &peer
}

func generateNetworkPolicyPort(protocol string, port uint16) *networkingv1.NetworkPolicyPort {
	protocolType := v1.Protocol(protocol)

	return &networkingv1.NetworkPolicyPort{
		Protocol: &protocolType,
		Port:     &intstr.IntOrString{Type: intstr.Int, IntVal: int32(port)},
	}
}

func generateNetworkPolicyPorts(portConfig []skatteetaten_no_v1alpha1.PortConfig) *[]networkingv1.NetworkPolicyPort {
	var ports []networkingv1.NetworkPolicyPort
	for _, port := range portConfig {
		ports = append(ports, *generateNetworkPolicyPort(port.Protocol, port.Port))
	}

	return &ports
}