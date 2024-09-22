# Overview

This Pulumi module provides integration with **Planton Cloud's Unified APIs** for deploying **MicroserviceKubernetes** resources across multiple cloud platforms. The module enables developers to define Kubernetes microservices using simple YAML configurations, closely following Kubernetes API resource conventions. The module is designed to deploy resources automatically based on the configuration in the `spec` of the `MicroserviceKubernetes` API resource. It uses a standardized structure that includes fields such as `apiVersion`, `kind`, `metadata`, `spec`, and `status`, ensuring uniformity across different resource types.

The Pulumi module supports a wide range of cloud providers, allowing developers to deploy microservices on infrastructure hosted by AWS, Azure, GCP, or any other cloud platform supported by Planton Cloud. This module enables developers to easily manage microservices, inject environment variables, utilize secrets management (such as GCP Secrets Manager), and configure resources like containers, ports, and more. Outputs from the deployment are captured in the `status.stackOutputs` for tracking key deployment details.

# Key Features

1. **Unified API Resource Structure**: All API resources follow the Kubernetes-inspired API modeling pattern, ensuring consistency and predictability. Resources are defined with common fields such as `apiVersion`, `kind`, `metadata`, `spec`, and `status`.

2. **Multi-Cloud Support**: The module supports deploying microservices on multiple cloud platforms. Whether you're using AWS, GCP, Azure, or others, the module abstracts cloud provider-specific details and allows you to focus on defining your microservice configuration.

3. **Integrated Pulumi Module**: The Pulumi module provides Golang-based implementations for deploying the resources defined in the YAML configuration. It automates the resource provisioning based on the `spec` defined in the API resource, and outputs essential details in `status.stackOutputs`.

4. **Environment Configuration**: Easily configure environment variables, secrets, and resource limits (CPU, memory, etc.) for your microservices. The module provides a flexible way to inject environment variables and manage secrets securely, including integration with secret managers like GCP Secrets Manager.

5. **Container and Port Management**: Define your container images, ports, and network configurations effortlessly. You can set container ports, enable ingress, and define the container image repository and tag.

6. **Resource Outputs**: After deployment, the module captures important outputs such as the status of the resources, service endpoints, and any other relevant deployment metadata. These outputs are stored in `status.stackOutputs` and can be referenced for future operations or tracking.

# Usage

To deploy resources using the **MicroserviceKubernetes** Pulumi module, create a YAML configuration that defines the required fields for your microservice. The YAML follows the structure of a Kubernetes resource and can be easily adapted for different environments. Once the YAML is created, you can use the Planton Cloud CLI to apply the configuration and trigger the Pulumi module to deploy the resources.

Refer to the **example section** for detailed usage instructions and sample YAML configurations.

```shell
planton pulumi up --stack-input <api-resource.yaml>
```

# Module Implementation

The Pulumi module integrates tightly with the **MicroserviceKubernetes** API resource, utilizing the fields defined in the `spec` to provision cloud resources. The core functionality involves parsing the YAML configuration, initializing the Pulumi stack, and using the provided configurations to deploy the microservices. 

For secrets management, the module supports integrating with external secret stores like GCP Secrets Manager. By defining secrets in the configuration, developers can securely manage and inject them into the microservice's environment variables.

The `metadata` field helps categorize and identify resources, while the `status.stackOutputs` field is used to store key output values from the Pulumi deployment.
