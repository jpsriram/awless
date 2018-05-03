/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/wallix/awless/cloud"
	"github.com/wallix/awless/logger"
	"github.com/wallix/awless/template/params"
)

type CreateClassicLoadbalancer struct {
	_                 string `action:"create" entity:"classicloadbalancer" awsAPI:"elb" awsCall:"CreateLoadBalancer" awsInput:"elb.CreateLoadBalancerInput" awsOutput:"elb.CreateLoadBalancerOutput"`
	logger            *logger.Logger
	graph             cloud.GraphAPI
	api               elbiface.ELBAPI
	Name              *string   `awsName:"LoadBalancerName" awsType:"awsstr" templateName:"name"`
	AvailabilityZones []*string `awsName:"AvailabilityZones" awsType:"awsstringslice" templateName:"zones"`
	Listeners         []*string `awsName:"Listeners" awsType:"awsclassicloadblisteners" templateName:"listeners"`
	Subnets           []*string `awsName:"Subnets" awsType:"awsstringslice" templateName:"subnets"`
	Securitygroups    []*string `awsName:"SecurityGroups" awsType:"awsstringslice" templateName:"securitygroups"`
	Scheme            *string   `awsName:"Scheme" awsType:"awsstr" templateName:"scheme"`
}

func (cmd *CreateClassicLoadbalancer) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(
		params.Key("name"), params.Key("listeners"),
		params.AtLeastOneOf(params.Key("subnets"), params.Key("zones")),
		params.Opt("scheme", "securitygroups"),
	))
}

func (cmd *CreateClassicLoadbalancer) ExtractResult(i interface{}) string {
	return awssdk.StringValue(cmd.Name)
}

type DeleteClassicLoadbalancer struct {
	_      string `action:"delete" entity:"classicloadbalancer" awsAPI:"elb" awsCall:"DeleteLoadBalancer" awsInput:"elb.DeleteLoadBalancerInput" awsOutput:"elb.DeleteLoadBalancerOutput"`
	logger *logger.Logger
	graph  cloud.GraphAPI
	api    elbiface.ELBAPI
	Name   *string `awsName:"LoadBalancerName" awsType:"awsstr" templateName:"name"`
}

func (cmd *DeleteClassicLoadbalancer) ParamsSpec() params.Spec {
	return params.NewSpec(params.AllOf(params.Key("name")))
}
