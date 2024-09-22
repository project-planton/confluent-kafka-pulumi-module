package pkg

import (
	"github.com/plantoncloud/planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/confluent/confluentcloudkafka"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	ConfluentCloudKafka *confluentcloudkafka.ConfluentCloudKafka
}

func initializeLocals(ctx *pulumi.Context, stackInput *confluentcloudkafka.ConfluentCloudKafkaStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.ConfluentCloudKafka = stackInput.Target

	return locals
}
