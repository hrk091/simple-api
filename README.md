# simple-api (for Tekton Test)

Simple API Server and k8s manifest that is useful as a knob to switch Tekton Task status Success/Fail.

## Sample Tekton Task

```yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: simple-api-task
spec:
  steps:
  - image: bash:latest
    name: step1
    script: |
      #!/usr/bin/env bash
      wget http://simple-api.simple.svc.cluster.local:8000/health && exit 0
      exit 1
```

Then you can toggle whether this task will succeed or fail by editing STATUS_CODE env var.
