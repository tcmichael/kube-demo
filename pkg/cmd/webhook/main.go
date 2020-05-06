package main

import (
	"fmt"
	"net/http"

	whhttp "github.com/slok/kubewebhook/pkg/http"

	"smartnews.com/tcmichael/kube-demo/pkg/webhook/mutating"
)

func main() {
	mpw, err := mutating.NewIngressWebhook()
	if err != nil {
		fmt.Printf("failed to create Ingress Webhook, err: %v", err)
		return
	}

	mpwh, err := whhttp.HandlerFor(mpw)
	if err != nil {
		fmt.Printf("failed to create ingress handler, err: %v", err)
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/webhooks/mutating/ingress", mpwh)
	fmt.Println("start service")
	if err := http.ListenAndServeTLS(":10080", "/etc/webhook/certs/cert.pem", "/etc/webhook/certs/key.pem", mux); err != nil {
		fmt.Printf("server closed, err: %v", err)
	}
}
