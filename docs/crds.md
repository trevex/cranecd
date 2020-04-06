# CRDs

## `Pipeline`

```yaml
apiVersion: cranecd.io/v1alpha1
kind: Pipeline
metadata:
  name: myapp
  namespace: mynamespace
spec:
  serviceAccount: cranecd-worker
  secretName: mysecret # needs to at least contain shared secret, can contain git credentials
  git:
    repository: git@github.com:hello/world
    branch: master
    secretName: mygitsecret # can be separate secret, but falls back to above otherwise
  image: cranecd/helm
  resources: # optional
    limits:
      cpu: 1
      memory: 512m
    requests:
      cpu: 0.25
      memory: 128m
  env:
  - name: IMAGE_REPOSITORY
    required: true # optional
    provided: true # optional
  - name: MY_VAR
    value: test # optional
  config: # worker specific, env variable can be used
    chart: stable/nginx
    release: dev
    values:
    - deploy/values.yaml
    overrides: # optional
     image.tag: 0.1.0-dev1
     image.repository: ${IMAGE_REPOSITORY}

```
