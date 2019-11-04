package config

import (
	"github.com/spf13/viper"
)

// AwsConf ..
type AwsConf struct {
	S3Key    string
	S3Secret string
	S3Region string
	S3Bucket string
}

var aws AwsConf

// Aws exportable function
func Aws() *AwsConf {
	return &aws
}

func LoadAws() {
	aws = AwsConf{
		S3Key:    viper.GetString("aws.s3_key"),
		S3Secret: viper.GetString("aws.s3_secret"),
		S3Region: viper.GetString("aws.s3_region"),
		S3Bucket: viper.GetString("aws.s3_bucket"),
	}
}
