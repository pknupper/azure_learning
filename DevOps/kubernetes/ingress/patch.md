# Setup

```bash
controller_tag=$(curl -s https://api.github.com/repos/kubernetes/ingress-nginx/releases/latest | grep tag_name | cut -d '"' -f 4)
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/${controller_tag}/deploy/static/provider/baremetal/deploy.yaml

mv deploy.yaml nginx-ingress-controller-deploy.yaml

kubectl apply -f nginx-ingress-controller-deploy.yaml
```

## Patch

```bash
kubectl get svc -n ingress-nginx

kubectl -n ingress-nginx patch svc ingress-nginx-controller --type='json' -p '[{"op":"replace","path":"/spec/type","value":"LoadBalancer"}]'

kubectl get service ingress-nginx-controller --namespace=ingress-nginx
```
