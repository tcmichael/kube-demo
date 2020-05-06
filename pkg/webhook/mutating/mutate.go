package mutating

import (
	"context"
	"fmt"

	"github.com/slok/kubewebhook/pkg/webhook"
	"github.com/slok/kubewebhook/pkg/webhook/mutating"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressMutator struct {
}

const (
	SecurityGroupsAnnoKey   = "alb.ingress.kubernetes.io/security-groups"
	SecurityGroupsAnnoValue = "sg-096a805e76c771869, sg-0f8c19e908cf748dd"
)

func (m *IngressMutator) Mutate(_ context.Context, obj metav1.Object) (bool, error) {
	ingress := obj.(*v1beta1.Ingress)
	fmt.Println(ingress.Namespace + "/" + ingress.Name)
	if ingress.Annotations == nil {
		ingress.Annotations = make(map[string]string)
	}
	ingress.Annotations[SecurityGroupsAnnoKey] = SecurityGroupsAnnoValue
	return true, nil
}

func NewIngressWebhook() (webhook.Webhook, error) {
	ingressMutator := &IngressMutator{}
	cfg := mutating.WebhookConfig{
		Name: "Ingress Mutator Webhook",
		Obj:  &v1beta1.Ingress{},
	}
	return mutating.NewWebhook(cfg, ingressMutator, nil, nil, nil)
}
