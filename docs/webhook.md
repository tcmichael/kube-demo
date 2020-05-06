# Webhook

Admission webhooks are HTTP callbacks that receive admission requests and do something with them.
- Mutating Webhook: Webhook that can change a request as well as accept/reject
- Validating Webhook: ebhook that cannot change request, but can accept or reject.

## How to implement a webhook (use demo as an example)
1. Develop an webhook server
   ```
   make
   docker build -t tcmichael128/smartnews:v1.0.0 .
   docker push tcmichael128/smartnews:v1.0.0
   ```
2. Create `Secret` in K8S
   ```
   hack/webhook-create-signed-cert.sh --service ingress-mutate-webhook --namespace default --secret webhook
   ```
3. Create the `Service` and the webhook `Deployment`
   ```
   kubectl apply -f  deploy/webhook/webhook.yaml
   ```
4. Create `WebhookConfiguration`
   ```
   bash hack/inject-ca-bundle.sh
   kubectl apply -f deploy/webhook/mutatingwebhook-ca-bundle.yaml 
   ```
5. Test
   ```
   kubectl apply -f deploy/webhook/ingress.yaml
   ```


