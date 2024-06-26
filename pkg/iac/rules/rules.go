package rules

import (
	trules "github.com/aquasecurity/trivy-checks/pkg/rules"

	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/accessanalyzer"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/apigateway"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/athena"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/cloudfront"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/cloudtrail"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/cloudwatch"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/codebuild"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/config"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/documentdb"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/dynamodb"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/ec2"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/ecr"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/ecs"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/efs"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/eks"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/elasticache"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/elasticsearch"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/elb"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/emr"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/iam"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/kinesis"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/kms"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/lambda"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/mq"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/msk"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/neptune"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/rds"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/redshift"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/s3"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/sam"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/sns"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/sqs"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/ssm"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/aws/workspaces"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/appservice"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/authorization"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/container"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/database"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/datafactory"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/datalake"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/keyvault"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/monitor"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/network"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/securitycenter"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/storage"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/azure/synapse"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/cloudstack/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/digitalocean/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/digitalocean/spaces"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/github/actions"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/github/branch_protections"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/github/repositories"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/bigquery"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/dns"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/gke"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/iam"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/kms"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/sql"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/google/storage"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/computing"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/dns"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/nas"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/network"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/rdb"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/nifcloud/sslcertificate"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/openstack/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/openstack/networking"
	_ "github.com/aquasecurity/trivy-checks/checks/cloud/oracle/compute"
	_ "github.com/aquasecurity/trivy-checks/checks/kubernetes/network"
)

func init() {
	for _, r := range trules.GetRules() {
		Register(r)
	}
}
