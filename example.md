# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ConfluentCloudKafka
metadata:
  name: my-confluent-kafka-cluster
spec:
  confluentCredentialId: my-confluent-credential-id
```

# Example with Environment Information

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ConfluentCloudKafka
metadata:
  name: my-env-confluent-kafka-cluster
spec:
  environmentInfo:
    envId: production-environment
  confluentCredentialId: prod-confluent-credential-id
```

# Example with Stack Job Settings

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ConfluentCloudKafka
metadata:
  name: advanced-confluent-kafka-cluster
spec:
  stackJobSettings:
    pulumiBackendCredentialId: my-pulumi-backend-credential
    stackJobRunnerId: my-stack-job-runner
  confluentCredentialId: advanced-confluent-credential-id
```
