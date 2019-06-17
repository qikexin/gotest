/*
Author: lipengwei
Date: 2019/6/4
Description: 
*/
package main

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"fmt"
	"encoding/json"
)

var accesskey string = "LTAIOa0PcnUY75fj"
var secretkey string = "fCHNW0uodHuRvWfXZJKIDLkGp9vCcT"
var reginId string = "cn-shenzhen"

//调用阿里云sdk步骤：1、创建client；2、创建request；3、设置request请求参数；4、将请求转换为acs请求并获取response；5、解析response
type Vminformation struct {
	Code int `json: "code"`
	Msg string `json: "msg"`
	Count int `json: "count"`
	Data []vmdata `json: "data"`
}
type vmdata struct {
	Id string `json: "id"`
	Name string `json: "name"`
	Status string `json: "status"`
	Cpu int `json: "cpu"`
	Mem int `json: "mem"`
	Pip string `json: "pip"`
	Eip string `json: "eip"`
	Os string `json: "os"`
	Time string `json: "time"`
}

type AliInstances struct {
	PageNumber int
	TotalCount int
	PageSize   int
	RequestId  string
	Instances instances
}
type instances struct {
	Instance []vminfo
}
type inneripAddr struct {
	IpAddress []string
}
type vminfo struct {
	InnerIpAddress inneripAddr
	ImageId string
	InstanceTypeFamily string
	VlanId string
	InstanceId string
	EipAddress eipAddr
	InternetMaxBandwidthIn int
	CreditSpecification string
	ZoneId string
	InternetChargeType string
	SpotStrategy string
	StoppedMode string
	SerialNumber string
	IoOptimized bool
	Memory int
	Cpu int
	VpcAttributes vpcAttrbutes
	InternetMaxBandwidthOut int
	DeviceAvailable bool
	SecurityGroupIds securityGroupIds
	SaleCycle string
	SpotPriceLimit float64
	AutoReleaseTime string
	StartTime string
	InstanceName string
	Description string
	ResourceGroupId string
	OSType string
	OSName string
	InstanceNetworkType string
	PublicIpAddress inneripAddr
	HostName string
	InstanceType string
	CreationTime string
	Status string
	ClusterId string
	Recyclable bool
	RegionId string
	GPUSpec string
	DeletionProtection bool
	DedicatedHostAttribute dedicatedHostAttribute
	OperationLocks operationLocks
	GPUAmount int
	InstanceChargeType string
	ExpiredTime string
	DeploymentSetId string
}


type eipAddr struct {
	IpAddress string
	AllocationId string
	InternetChargeType string
}
type vpcAttrbutes struct {
	NatIpAddress string
	PrivateIpAddress inneripAddr
	VSwitchId string
	VpcId string
}
type securityGroupIds struct {
	SecurityGroupId []string
}


type dedicatedHostAttribute struct {
	DedicatedHostId string
	DedicatedHostName string
}

type operationLocks struct {
	LockReason []lockreason
}

type lockreason struct {
	LockMsg string
	LockReason string
}


type Rdsinformation struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int `json:"count"`
	Data []rdsdata `json:"data"`
}

type rdsdata struct {
	Id string `json:"id"`
	Engine string `json:"engine"`
	Engversion string `json:"engversion"`
	Status string `json:"status"`
	CreateTime string `json:"ctime"`
}
type RdsArch struct {
	TotalRecordCount int
	PageNumber int
	RequestId string
	PageRecordCount int
	Items rdsinfo
}

type rdsinfo struct {
	DBInstance []rdsinstance
}

type readonlyinstanceid struct {
	DBInstanceId string
}

type readonlyinstance struct {
	ReadOnlyDBInstanceId []readonlyinstanceid
}

type rdsinstance struct {
	LockMode string
	DBInstanceNetType string
	MasterInstanceId string
	DBInstanceClass string
	ResourceGroupId string
	DBInstanceId string
	VpcCloudInstanceId string
	ZoneId string
	ReadOnlyDBInstanceIds readonlyinstance
	ConnectionMode string
	InstanceNetworkType string
	VSwitchId string
	VpcId string
	Engine string
	MutriORsignle bool
	InsId int
	ExpireTime string
	CreateTime string
	DBInstanceType string
	RegionId string
	EngineVersion string
	LockReason string
	DBInstanceStatus string
	PayType string
}

func main()  {
	client, err := ecs.NewClientWithAccessKey(reginId,accesskey,secretkey)
	if err != nil {
		panic(err)
	}
	request := ecs.CreateDescribeInstancesRequest()
	request.Domain = "ecs.aliyuncs.com"
	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	var aliInstances  AliInstances
	//fmt.Println(response.GetHttpContentString())
	json.Unmarshal(response.GetHttpContentBytes(),&aliInstances)
	fmt.Printf("requestId: %s\n",aliInstances.RequestId)
	fmt.Printf("pageNumber: %d\n",aliInstances.PageNumber)
	fmt.Printf("pageSize: %d\n",aliInstances.PageSize)
	fmt.Printf("totle: %d\n",aliInstances.TotalCount)
	fmt.Printf("instances: %+v\n",aliInstances.Instances)
}
