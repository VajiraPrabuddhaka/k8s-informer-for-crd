package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type APIV1Alpha1Interface interface {
	APIs(namespace string) APIInterface
}

type APIV1Alpha1Client struct {
	restClient rest.Interface
}

func (c *APIV1Alpha1Client) APIs(namespace string) APIInterface {
	return &apiClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}

func NewForConfig(c *rest.Config) (*APIV1Alpha1Client, error) {
	AddToScheme(scheme.Scheme)
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GroupVersion.Group, Version: GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &APIV1Alpha1Client{restClient: client}, nil
}
