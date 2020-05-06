cat deploy/webhook/mutatingwebhook.yaml | hack/webhook-patch-ca-bundle.sh > deploy/webhook/mutatingwebhook-ca-bundle.yaml
