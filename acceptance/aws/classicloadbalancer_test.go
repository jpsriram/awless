package awsat

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/elb"
)

func TestClassicLoadbalancer(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		Template("create classicloadbalancer name=my-new-loadbalancer listeners=[HTTPS:443:HTTP:80,TCP:123456:UDP:8080] zones=[us-west-2b,us-west-2c] subnets=[sub-1,sub-2] securitygroups=[sg-1,sg-2]").Mock(&elbMock{
			CreateLoadBalancerFunc: func(input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
				return &elb.CreateLoadBalancerOutput{
					DNSName: String(""),
				}, nil
			}}).
			ExpectInput("CreateLoadBalancer", &elb.CreateLoadBalancerInput{
				AvailabilityZones: []*string{String("us-west-2b"), String("us-west-2c")},
				LoadBalancerName:  String("my-new-loadbalancer"),
				Subnets:           []*string{String("sub-1"), String("sub-2")},
				Listeners: []*elb.Listener{
					{Protocol: String("HTTPS"), LoadBalancerPort: Int64(443), InstanceProtocol: String("HTTP"), InstancePort: Int64(80)},
					{Protocol: String("TCP"), LoadBalancerPort: Int64(123456), InstanceProtocol: String("UDP"), InstancePort: Int64(8080)},
				},
				SecurityGroups: []*string{String("sg-1"), String("sg-2")},
			}).ExpectCalls("CreateLoadBalancer").Run(t)
	})

	t.Run("delete", func(t *testing.T) {
		Template("delete classicloadbalancer name=my-classic-loadb").Mock(&elbMock{
			DeleteLoadBalancerFunc: func(input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
				return nil, nil
			}}).
			ExpectInput("DeleteLoadBalancer", &elb.DeleteLoadBalancerInput{
				LoadBalancerName: String("my-classic-loadb"),
			}).ExpectCalls("DeleteLoadBalancer").Run(t)
	})
}
