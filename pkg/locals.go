package pkg

import (
	"github.com/plantoncloud/planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/confluent/confluentkafka"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	ConfluentKafka *confluentkafka.ConfluentKafka
}

func initializeLocals(ctx *pulumi.Context, stackInput *confluentkafka.ConfluentKafkaStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.ConfluentKafka = stackInput.Target

	return locals
}
