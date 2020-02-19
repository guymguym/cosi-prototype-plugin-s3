module github.com/yard-turkey/cosi-prototype-plugin-s3

go 1.13

require (
	github.com/aws/aws-sdk-go v1.28.5
	github.com/minio/minio-go/v6 v6.0.44 // indirect
	github.com/yard-turkey/cosi-prototype-interface v0.0.0
	google.golang.org/grpc v1.26.0
	k8s.io/klog v0.4.0
	sigs.k8s.io/controller-runtime v0.4.0
)

replace github.com/yard-turkey/cosi-prototype-interface v0.0.0 => /Users/jcope/Workspace/cosi-prototype/cosi-prototype-interface
