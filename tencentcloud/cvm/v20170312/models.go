// Copyright 1999-2018 Tencent Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20170312

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActionTimer struct {
	// 定时器
	TimerAction *string `json:"TimerAction" name:"TimerAction"`
	// 执行时间
	ActionTime *string `json:"ActionTime" name:"ActionTime"`
	// 扩展数据
	Externals *Externals `json:"Externals" name:"Externals"`
}

type AllocateHostsRequest struct {
	*tchttp.BaseRequest
	// 用于保证请求幂等性的字符串。
	ClientToken *string `json:"ClientToken" name:"ClientToken"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement" name:"Placement"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid" name:"HostChargePrepaid"`
	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。
	HostChargeType *string `json:"HostChargeType" name:"HostChargeType"`
	// CDH实例机型，默认为：'HS1'。
	HostType *string `json:"HostType" name:"HostType"`
	// 购买CDH实例数量。
	HostCount *uint64 `json:"HostCount" name:"HostCount"`
}

func (r *AllocateHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 新创建云子机的实例id列表。
		HostIdSet []*string `json:"HostIdSet" name:"HostIdSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AllocateHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br>可以通过以下方式获取可用的实例ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。<br><li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的`InstanceId`获取实例ID。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 一个或多个待操作的密钥对ID，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-3glfot13`。<br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的`KeyId`获取密钥对ID。
	KeyIds []*string `json:"KeyIds" name:"KeyIds" list`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *AssociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateInstancesKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
}

type CreateImageRequest struct {
	*tchttp.BaseRequest
	// 需要制作镜像的实例ID
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 镜像名称
	ImageName *string `json:"ImageName" name:"ImageName"`
	// 镜像描述
	ImageDescription *string `json:"ImageDescription" name:"ImageDescription"`
	// 软关机失败时是否执行强制关机以制作镜像
	ForcePoweroff *string `json:"ForcePoweroff" name:"ForcePoweroff"`
	// 创建Windows镜像时是否启用Sysprep
	Sysprep *string `json:"Sysprep" name:"Sysprep"`
	// 实例处于运行中时，是否允许关机执行制作镜像任务。
	Reboot *string `json:"Reboot" name:"Reboot"`
}

func (r *CreateImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairRequest struct {
	*tchttp.BaseRequest
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName" name:"KeyName"`
	// 密钥对创建后所属的项目ID。
	// 可以通过以下方式获取项目ID：
	// <li>通过项目列表查询项目ID。
	// <li>通过调用接口DescribeProject，取返回信息中的`projectId `获取项目ID。
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
}

func (r *CreateKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyPairRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 密钥对信息。
		KeyPair *KeyPair `json:"KeyPair" name:"KeyPair"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘类型。数据盘类型限制详见[CVM实例配置](/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。<br><br>该参数对`ResizeInstanceDisk`接口无效。
	DiskType *string `json:"DiskType" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId" name:"DiskId"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[CVM实例配置](/document/product/213/2177)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize *int64 `json:"DiskSize" name:"DiskSize"`
}

type DeleteImagesRequest struct {
	*tchttp.BaseRequest
	// 准备删除的镜像Id列表
	ImageIds []*string `json:"ImageIds" name:"ImageIds" list`
}

func (r *DeleteImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。<br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。
	KeyIds []*string `json:"KeyIds" name:"KeyIds" list`
}

func (r *DeleteKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li> zone - String - 是否必填：否 - （过滤条件）按照可用区过滤。</li>
	// <li> project-id - Integer - 是否必填：否 - （过滤条件）按照项目ID过滤。可通过调用 DescribeProject 查询已创建的项目列表或登录控制台进行查看；也可以调用 AddProject 创建新的项目。</li>
	// <li> host-id - String - 是否必填：否 - （过滤条件）按照CDH ID过滤。CDH ID形如：host-11112222。</li>
	// <li> host-name - String - 是否必填：否 - （过滤条件）按照CDH实例名称过滤。</li>
	// <li> host-state - String - 是否必填：否 - （过滤条件）按照CDH实例状态进行过滤。（PENDING：创建中|LAUNCH_FAILURE：创建失败|RUNNING：运行中|EXPIRED：已过期）</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合查询条件的cdh实例总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// cdh实例详细信息列表
		HostSet []*HostItem `json:"HostSet" name:"HostSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 账户的镜像配额
		ImageNumQuota *int64 `json:"ImageNumQuota" name:"ImageNumQuota"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	// 需要共享的镜像Id
	ImageId *string `json:"ImageId" name:"ImageId"`
}

func (r *DescribeImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 镜像共享信息
		SharePermissionSet []*SharePermission `json:"SharePermissionSet" name:"SharePermissionSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	// 镜像ID列表 。镜像ID如：`img-gvbnzy6f`。array型参数的格式可以参考[API简介](https://cloud.tencent.com/document/api/213/15688)。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。<br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。
	ImageIds []*string `json:"ImageIds" name:"ImageIds" list`
	// 过滤条件，每次请求的`Filters`的上限为0，`Filters.Values`的上限为5。参数不可以同时指定`ImageIds`和`Filters`。详细的过滤条件如下：
	// <li> image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤</li>
	// <li> image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：详见[镜像类型](https://cloud.tencent.com/document/product/213/9452#image_type)。</li>
	// <li> image-state - String - 是否必填： 否 - （过滤条件）按照镜像状态进行过滤。取值范围：详见[镜像状态](https://cloud.tencent.com/document/product/213/9452#image_state)。</li>
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于Offset详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 数量限制，默认为20，最大值为100。关于Limit详见[API简介](/document/api/213/568#.E8.BE.93.E5.85.A5.E5.8F.82.E6.95.B0.E4.B8.8E.E8.BF.94.E5.9B.9E.E5.8F.82.E6.95.B0.E9.87.8A.E4.B9.89)。
	Limit *uint64 `json:"Limit" name:"Limit"`
	// 实例类型，如 `S1.SMALL1`
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。
		ImageSet []*Image `json:"ImageSet" name:"ImageSet" list`
		// 符合要求的镜像数量。
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImportImageOsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImportImageOsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 支持的导入镜像的操作系统类型
		ImportImageOsListSupported []*string `json:"ImportImageOsListSupported" name:"ImportImageOsListSupported" list`
		// 支持的导入镜像的操作系统版本
		ImportImageOsVersionSupported []*string `json:"ImportImageOsVersionSupported" name:"ImportImageOsVersionSupported" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImportImageOsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImportImageOsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceFamilyConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceFamilyConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 实例机型组配置的列表信息
		InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet" name:"InstanceFamilyConfigSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceFamilyConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceFamilyConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInternetBandwidthConfigsRequest struct {
	*tchttp.BaseRequest
	// 待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 带宽配置信息列表。
		InternetBandwidthConfigSet []*InternetBandwidthConfig `json:"InternetBandwidthConfigSet" name:"InternetBandwidthConfigSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceInternetBandwidthConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsRequest struct {
	*tchttp.BaseRequest
	// 每次请求的`Filters`的上限为1，`Filter.Values`的上限为1。
	// Filters.1.Name目前支持“instance-id”，即根据实例 ID 过滤。实例 ID 形如：ins-1w2x3y4z。
	Filters []*Filter `json:"Filters" name:"Filters" list`
}

func (r *DescribeInstanceOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceOperationLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsRequest struct {
	*tchttp.BaseRequest
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照[可用区](https://cloud.tencent.com/document/api/213/9452#zone)过滤。</li>
	// <li> instance-family - String - 是否必填：否 -（过滤条件）按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为1。
	Filters []*Filter `json:"Filters" name:"Filters" list`
}

func (r *DescribeInstanceTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 实例机型配置列表。
		InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet" name:"InstanceTypeConfigSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。（此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`id.N`一节）。每次请求的实例的上限为100。参数不支持同时指定`InstanceIds`和`Filters`。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 过滤条件。
	// <li> zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。</li>
	// <li> project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可通过调用[DescribeProject](https://cloud.tencent.com/document/api/378/4400)查询已创建的项目列表或登录[控制台](https://console.cloud.tencent.com/cvm/index)进行查看；也可以调用[AddProject](https://cloud.tencent.com/document/api/378/4398)创建新的项目。</li>
	// <li> host-id - String - 是否必填：否 - （过滤条件）按照[CDH](https://cloud.tencent.com/document/product/416) ID过滤。[CDH](https://cloud.tencent.com/document/product/416) ID形如：host-11112222。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：ins-11112222。</li>
	// <li> instance-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。</li>
	// <li> instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示[CDH](https://cloud.tencent.com/document/product/416)付费，即只对[CDH](https://cloud.tencent.com/document/product/416)计费，不对[CDH](https://cloud.tencent.com/document/product/416)上的实例计费。 )  </li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的内网IP过滤。</li>
	// <li> public-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的公网IP过滤，包含实例创建时自动分配的IP和实例创建后手动绑定的弹性IP。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例数量。
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 实例详细信息列表。
		InstanceSet []*Instance `json:"InstanceSet" name:"InstanceSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusRequest struct {
	*tchttp.BaseRequest
	// 按照一个或者多个实例ID查询。实例ID形如：`ins-11112222`。此参数的具体格式可参考API[简介](https://cloud.tencent.com/document/api/213/15688)的`id.N`一节）。每次请求的实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit" name:"Limit"`
}

func (r *DescribeInstancesStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的实例状态数量。
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// [实例状态](https://cloud.tencent.com/document/api/213/15738) 列表。
		InstanceStatusSet []*InstanceStatus `json:"InstanceStatusSet" name:"InstanceStatusSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInternetChargeTypeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInternetChargeTypeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 网络计费类型配置。
		InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet" name:"InternetChargeTypeConfigSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternetChargeTypeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInternetChargeTypeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsRequest struct {
	*tchttp.BaseRequest
	// 密钥对ID，密钥对ID形如：`skey-11112222`（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API [简介](https://cloud.tencent.com/document/api/213/15688)的 `id.N` 一节）。参数不支持同时指定 `KeyIds` 和 `Filters`。密钥对ID可以通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询。
	KeyIds []*string `json:"KeyIds" name:"KeyIds" list`
	// 过滤条件。
	// <li> project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可以通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID，或者调用接口 [DescribeProject](https://cloud.tencent.com/document/api/378/4400)，取返回信息中的projectId获取项目ID。</li>
	// <li> key-name - String - 是否必填：否 -（过滤条件）按照密钥对名称过滤。</li>参数不支持同时指定 `KeyIds` 和 `Filters`。
	Filters []*Filter `json:"Filters" name:"Filters" list`
	// 偏移量，默认为0。关于 `Offset` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于 `Limit` 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit" name:"Limit"`
}

func (r *DescribeKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 符合条件的密钥对数量。
		TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
		// 密钥对详细信息列表。
		KeyPairSet []*KeyPair `json:"KeyPairSet" name:"KeyPairSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 地域数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 地域列表信息
		RegionSet []*RegionInfo `json:"RegionSet" name:"RegionSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 可用区数量
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 可用区列表信息
		ZoneSet []*ZoneInfo `json:"ZoneSet" name:"ZoneSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZonesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID，每次请求批量实例的上限为100。<br><br>可以通过以下方式获取可用的实例ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/index)查询实例ID。<br><li>通过调用接口 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) ，取返回信息中的 `InstanceId` 获取密钥对ID。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 密钥对ID列表，每次请求批量密钥对的上限为100。密钥对ID形如：`skey-11112222`。<br><br>可以通过以下方式获取可用的密钥ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥ID。<br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699) ，取返回信息中的 `KeyId` 获取密钥对ID。
	KeyIds []*string `json:"KeyIds" name:"KeyIds" list`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机。<br><li>FALSE：表示在正常关机失败后不进行强制关机。<br><br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *DisassociateInstancesKeyPairsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateInstancesKeyPairsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateInstancesKeyPairsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateInstancesKeyPairsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateInstancesKeyPairsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService" name:"MonitorService"`
}

type Externals struct {
	// 释放地址
	ReleaseAddress *bool `json:"ReleaseAddress" name:"ReleaseAddress"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name" name:"Name"`
	// 字段的过滤值。
	Values []*string `json:"Values" name:"Values" list`
}

type HostItem struct {
	// cdh实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement" name:"Placement"`
	// cdh实例id
	HostId *string `json:"HostId" name:"HostId"`
	// cdh实例类型
	HostType *string `json:"HostType" name:"HostType"`
	// cdh实例名称
	HostName *string `json:"HostName" name:"HostName"`
	// cdh实例付费模式
	HostChargeType *string `json:"HostChargeType" name:"HostChargeType"`
	// cdh实例自动续费标记
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
	// cdh实例创建时间
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// cdh实例过期时间
	ExpiredTime *string `json:"ExpiredTime" name:"ExpiredTime"`
	// cdh实例上已创建云子机的实例id列表
	InstanceIds *string `json:"InstanceIds" name:"InstanceIds"`
	// cdh实例状态
	HostState *string `json:"HostState" name:"HostState"`
	// cdh实例ip
	HostIp *string `json:"HostIp" name:"HostIp"`
	// cdh实例资源信息
	HostResource *HostResource `json:"HostResource" name:"HostResource"`
}

type HostResource struct {
	// cdh实例总cpu核数
	CpuTotal *uint64 `json:"CpuTotal" name:"CpuTotal"`
	// cdh实例可用cpu核数
	CpuAvailable *uint64 `json:"CpuAvailable" name:"CpuAvailable"`
	// cdh实例总内存大小（单位为:GiB）
	MemTotal *float64 `json:"MemTotal" name:"MemTotal"`
	// cdh实例可用内存大小（单位为:GiB）
	MemAvailable *float64 `json:"MemAvailable" name:"MemAvailable"`
	// cdh实例总磁盘大小（单位为:GiB）
	DiskTotal *uint64 `json:"DiskTotal" name:"DiskTotal"`
	// cdh实例可用磁盘大小（单位为:GiB）
	DiskAvailable *uint64 `json:"DiskAvailable" name:"DiskAvailable"`
}

type Image struct {
	// 镜像ID
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 镜像操作系统
	OsName *string `json:"OsName" name:"OsName"`
	// 镜像类型
	ImageType *string `json:"ImageType" name:"ImageType"`
	// 镜像创建时间
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// 镜像名称
	ImageName *string `json:"ImageName" name:"ImageName"`
	// 镜像描述
	ImageDescription *string `json:"ImageDescription" name:"ImageDescription"`
	// 镜像大小
	ImageSize *int64 `json:"ImageSize" name:"ImageSize"`
	// 镜像架构
	Architecture *string `json:"Architecture" name:"Architecture"`
	// 镜像状态
	ImageState *string `json:"ImageState" name:"ImageState"`
	// 镜像来源平台
	Platform *string `json:"Platform" name:"Platform"`
	// 镜像创建者
	ImageCreator *string `json:"ImageCreator" name:"ImageCreator"`
	// 镜像来源
	ImageSource *string `json:"ImageSource" name:"ImageSource"`
}

type ImportImageRequest struct {
	*tchttp.BaseRequest
	// 导入镜像的操作系统架构，`x86_64` 或 `i386`
	Architecture *string `json:"Architecture" name:"Architecture"`
	// 导入镜像的操作系统类型，通过`DescribeImportImageOs`获取
	OsType *string `json:"OsType" name:"OsType"`
	// 导入镜像的操作系统版本，通过`DescribeImportImageOs`获取
	OsVersion *string `json:"OsVersion" name:"OsVersion"`
	// 导入镜像存放的cos地址
	ImageUrl *string `json:"ImageUrl" name:"ImageUrl"`
	// 镜像名称
	ImageName *string `json:"ImageName" name:"ImageName"`
	// 镜像描述
	ImageDescription *string `json:"ImageDescription" name:"ImageDescription"`
	// 只检查参数，不执行任务
	DryRun *bool `json:"DryRun" name:"DryRun"`
	// 是否强制导入，参考[强制导入镜像](https://cloud.tencent.com/document/product/213/12849)
	Force *bool `json:"Force" name:"Force"`
}

func (r *ImportImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairRequest struct {
	*tchttp.BaseRequest
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName" name:"KeyName"`
	// 密钥对创建后所属的[项目](/document/product/378/10863)ID。<br><br>可以通过以下方式获取项目ID：<br><li>通过[项目列表](https://console.cloud.tencent.com/project)查询项目ID。<br><li>通过调用接口 [DescribeProject](https://cloud.tencent.com/document/api/378/4400)，取返回信息中的 `projectId ` 获取项目ID。
	// 
	// 如果是默认项目，直接填0就可以。
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 密钥对的公钥内容，`OpenSSH RSA` 格式。
	PublicKey *string `json:"PublicKey" name:"PublicKey"`
}

func (r *ImportKeyPairRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImportKeyPairResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 密钥对ID。
		KeyId *string `json:"KeyId" name:"KeyId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyPairResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImportKeyPairResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
	// 试运行。
	DryRun *bool `json:"DryRun" name:"DryRun"`
}

func (r *InquiryPriceRenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示对应配置实例的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceRequest struct {
	*tchttp.BaseRequest
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 指定有效的[镜像](/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk" name:"SystemDisk"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService" name:"EnhancedService"`
}

func (r *InquiryPriceResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示重装成对应配置实例的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持`InternetMaxBandwidthOut`参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 带宽生效的终止时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime" name:"EndTime"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示带宽调整为对应大小之后的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *InquiryPriceResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示调整成对应机型实例的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	// 待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且[数据盘类型](/document/api/213/9452#block_device)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	DataDisks []*DataDisk `json:"DataDisks" name:"DataDisks" list`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *InquiryPriceResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示磁盘扩容成对应配置的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesRequest struct {
	*tchttp.BaseRequest
	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement *Placement `json:"Placement" name:"Placement"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[CVM实例配置](https://cloud.tencent.com/document/product/213/2177)描述。若不指定该参数，则默认机型为S1.SMALL1。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk" name:"SystemDisk"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，当前仅支持购买的时候指定一个数据盘。
	DataDisks []*DataDisk `json:"DataDisks" name:"DataDisks" list`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，那么InstanceCount参数只能为1。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount" name:"InstanceCount"`
	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken" name:"ClientToken"`
	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName" name:"HostName"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云主机实例。
	TagSpecification []*TagSpecification `json:"TagSpecification" name:"TagSpecification" list`
	// 实例的市场相关选项，如竞价实例相关参数
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions" name:"InstanceMarketOptions"`
}

func (r *InquiryPriceRunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 该参数表示对应配置实例的价格。
		Price *Price `json:"Price" name:"Price"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例所在的位置。
	Placement *Placement `json:"Placement" name:"Placement"`
	// 实例`ID`。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 实例机型。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 实例的CPU核数，单位：核。
	CPU *int64 `json:"CPU" name:"CPU"`
	// 实例内存容量，单位：`GB`。
	Memory *int64 `json:"Memory" name:"Memory"`
	// 实例业务状态。取值范围：<br><li>NORMAL：表示正常状态的实例<br><li>EXPIRED：表示过期的实例<br><li>PROTECTIVELY_ISOLATED：表示被安全隔离的实例。
	RestrictState *string `json:"RestrictState" name:"RestrictState"`
	// 实例名称。
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	// 实例计费模式。取值范围：<br><li>`PREPAID`：表示预付费，即包年包月<br><li>`POSTPAID_BY_HOUR`：表示后付费，即按量计费<br><li>`CDHPAID`：`CDH`付费，即只对`CDH`计费，不对`CDH`上的实例计费。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 实例系统盘信息。
	SystemDisk *SystemDisk `json:"SystemDisk" name:"SystemDisk"`
	// 实例数据盘信息。只包含随实例购买的数据盘。
	DataDisks []*DataDisk `json:"DataDisks" name:"DataDisks" list`
	// 实例主网卡的内网`IP`列表。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
	// 实例主网卡的公网`IP`列表。
	PublicIpAddresses []*string `json:"PublicIpAddresses" name:"PublicIpAddresses" list`
	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
	// 实例所属虚拟私有网络信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud" name:"VirtualPrivateCloud"`
	// 生产实例所使用的镜像`ID`。
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 自动续费标识。取值范围：<br><li>`NOTIFY_AND_MANUAL_RENEW`：表示通知即将过期，但不自动续费<br><li>`NOTIFY_AND_AUTO_RENEW`：表示通知即将过期，而且自动续费<br><li>`DISABLE_NOTIFY_AND_MANUAL_RENEW`：表示不通知即将过期，也不自动续费。
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	// 到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	ExpiredTime *string `json:"ExpiredTime" name:"ExpiredTime"`
	// 操作系统名称。
	OsName *string `json:"OsName" name:"OsName"`
	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
	// 实例登录设置。目前只返回实例所关联的密钥。
	LoginSettings *LoginSettings `json:"LoginSettings" name:"LoginSettings"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period" name:"Period"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
}

type InstanceFamilyConfig struct {
	// 机型族名称的中文全称。
	InstanceFamilyName *string `json:"InstanceFamilyName" name:"InstanceFamilyName"`
	// 机型族名称的英文简称。
	InstanceFamily *string `json:"InstanceFamily" name:"InstanceFamily"`
}

type InstanceMarketOptionsRequest struct {
	*tchttp.BaseRequest
	// 市场选项类型，当前只支持取值：spot
	MarketType *string `json:"MarketType" name:"MarketType"`
	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions" name:"SpotOptions"`
}

func (r *InstanceMarketOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceMarketOptionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceStatus struct {
	// 实例`ID`。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// [实例状态](/document/api/213/9452#INSTANCE_STATE)。
	InstanceState *string `json:"InstanceState" name:"InstanceState"`
}

type InstanceTypeConfig struct {
	// 可用区。
	Zone *string `json:"Zone" name:"Zone"`
	// 实例机型。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily" name:"InstanceFamily"`
	// GPU核数，单位：核。
	GPU *int64 `json:"GPU" name:"GPU"`
	// CPU核数，单位：核。
	CPU *int64 `json:"CPU" name:"CPU"`
	// 内存容量，单位：`GB`。
	Memory *int64 `json:"Memory" name:"Memory"`
	// 是否支持云硬盘。取值范围：<br><li>`TRUE`：表示支持云硬盘；<br><li>`FALSE`：表示不支持云硬盘。
	CbsSupport *string `json:"CbsSupport" name:"CbsSupport"`
	// 机型状态。取值范围：<br><li>`AVAILABLE`：表示机型可用；<br><li>`UNAVAILABLE`：表示机型不可用。
	InstanceTypeState *string `json:"InstanceTypeState" name:"InstanceTypeState"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。
	InternetChargeType *string `json:"InternetChargeType" name:"InternetChargeType"`
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽](/document/product/213/509)。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut" name:"InternetMaxBandwidthOut"`
	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned" name:"PublicIpAssigned"`
}

type InternetBandwidthConfig struct {
	// 开始时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 结束时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 实例带宽信息。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
}

type InternetChargeTypeConfig struct {
	// 网络计费模式。
	InternetChargeType *string `json:"InternetChargeType" name:"InternetChargeType"`
	// 网络计费模式描述信息。
	Description *string `json:"Description" name:"Description"`
}

type ItemPrice struct {
	// 后续单价，单位：元。
	UnitPrice *float64 `json:"UnitPrice" name:"UnitPrice"`
	// 后续计价单元，可取值范围： <br><li>HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）：<br><li>GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	ChargeUnit *string `json:"ChargeUnit" name:"ChargeUnit"`
	// 预支费用的原价，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice" name:"OriginalPrice"`
	// 预支费用的折扣价，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice" name:"DiscountPrice"`
}

type KeyPair struct {
	// 密钥对的`ID`，是密钥对的唯一标识。
	KeyId *string `json:"KeyId" name:"KeyId"`
	// 密钥对名称。
	KeyName *string `json:"KeyName" name:"KeyName"`
	// 密钥对所属的项目`ID`。
	ProjectId *string `json:"ProjectId" name:"ProjectId"`
	// 密钥对描述信息。
	Description *string `json:"Description" name:"Description"`
	// 密钥对的纯文本公钥。
	PublicKey *string `json:"PublicKey" name:"PublicKey"`
	// 密钥对的纯文本私钥。腾讯云不会保管私钥，请用户自行妥善保存。
	PrivateKey *string `json:"PrivateKey" name:"PrivateKey"`
	// 密钥关联的实例`ID`列表。
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds" name:"AssociatedInstanceIds" list`
	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds" name:"KeyIds" list`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin" name:"KeepImageLogin"`
}

type ModifyHostsAttributeRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds" name:"HostIds" list`
	// CDH实例显示名称。可任意命名，但不得超过60个字符。
	HostName *string `json:"HostName" name:"HostName"`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
}

func (r *ModifyHostsAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyHostsAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHostsAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyHostsAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeRequest struct {
	*tchttp.BaseRequest
	// 镜像ID，形如`img-gvbnzy6f`。镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。<br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 设置新的镜像名称；必须满足下列限制：<br> <li> 不得超过20个字符。<br> <li> 镜像名称不能与已有镜像重复。
	ImageName *string `json:"ImageName" name:"ImageName"`
	// 设置新的镜像描述；必须满足下列限制：<br> <li> 不得超过60个字符。
	ImageDescription *string `json:"ImageDescription" name:"ImageDescription"`
}

func (r *ModifyImageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionRequest struct {
	*tchttp.BaseRequest
	// 镜像ID，形如`img-gvbnzy6f`。镜像Id可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。<br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。 <br>镜像ID必须指定为状态为`NORMAL`的镜像。镜像状态请参考[镜像数据表](/document/api/213/9452#image_state)。
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 接收分享镜像的账号Id列表，array型参数的格式可以参考[API简介](/document/api/213/568)。帐号ID不同于QQ号，查询用户帐号ID请查看[帐号信息](https://console.cloud.tencent.com/developer)中的帐号ID栏。
	AccountIds []*string `json:"AccountIds" name:"AccountIds" list`
	// 操作，包括 `SHARE`，`CANCEL`。其中`SHARE`代表分享操作，`CANCEL`代表取消分享操作。
	Permission *string `json:"Permission" name:"Permission"`
}

func (r *ModifyImageSharePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyImageSharePermissionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyImageSharePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyImageSharePermissionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 实例名称。可任意命名，但不得超过60个字符。
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
}

func (r *ModifyInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 项目ID。项目可以使用[AddProject](https://cloud.tencent.com/doc/api/403/4398)接口创建。后续使用[DescribeInstances](https://cloud.tencent.com/document/api/213/9388)接口查询实例时，项目ID可用于过滤结果。
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
}

func (r *ModifyInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag" name:"RenewFlag"`
}

func (r *ModifyInstancesRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesRenewFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstancesRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstancesRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstancesRenewFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeRequest struct {
	*tchttp.BaseRequest
	// 密钥对ID，密钥对ID形如：`skey-xxxxxxxx`。<br><br>可以通过以下方式获取可用的密钥 ID：<br><li>通过登录[控制台](https://console.cloud.tencent.com/cvm/sshkey)查询密钥 ID。<br><li>通过调用接口 [DescribeKeyPairs](https://cloud.tencent.com/document/api/213/9403) ，取返回信息中的 `KeyId` 获取密钥对 ID。
	KeyId *string `json:"KeyId" name:"KeyId"`
	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `json:"KeyName" name:"KeyName"`
	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。
	Description *string `json:"Description" name:"Description"`
}

func (r *ModifyKeyPairAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyKeyPairAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKeyPairAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyKeyPairAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Placement struct {
	// 实例所属的[可用区](/document/product/213/9452#zone)ID。该参数也可以通过调用  [DescribeZones](/document/api/213/9455) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone" name:"Zone"`
	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](/document/api/378/4400) 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。当前暂不支持。
	HostIds []*string `json:"HostIds" name:"HostIds" list`
}

type Price struct {
	// 描述了实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice" name:"InstancePrice"`
	// 描述了网络价格。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice" name:"BandwidthPrice"`
}

type RebootInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 是否在正常重启失败后选择强制重启实例。取值范围：<br><li>TRUE：表示在正常重启失败后进行强制重启<br><li>FALSE：表示在正常重启失败后不进行强制重启<br><br>默认取值：FALSE。
	ForceReboot *bool `json:"ForceReboot" name:"ForceReboot"`
}

func (r *RebootInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RebootInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域名称，例如，ap-guangzhou
	Region *string `json:"Region" name:"Region"`
	// 地域描述，例如，华南地区(广州)
	RegionName *string `json:"RegionName" name:"RegionName"`
	// 地域是否可用状态
	RegionState *string `json:"RegionState" name:"RegionState"`
}

type RenewHostsRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `json:"HostIds" name:"HostIds" list`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `json:"HostChargePrepaid" name:"HostChargePrepaid"`
}

func (r *RenewHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewHostsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewHostsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceRequest struct {
	*tchttp.BaseRequest
	// 实例ID。可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9388) API返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/9418) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `json:"SystemDisk" name:"SystemDisk"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings" name:"LoginSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService" name:"EnhancedService"`
}

func (r *ResetInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的 `InstanceId` 获取。 每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持 `InternetMaxBandwidthOut` 参数。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
	// 带宽生效的起始时间。格式：`YYYY-MM-DD`，例如：`2016-10-30`。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 带宽生效的终止时间。格式： `YYYY-MM-DD` ，例如：`2016-10-30` 。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过 [`DescribeInstances`](https://cloud.tencent.com/document/api/213/9388)接口返回值中的`ExpiredTime`获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `json:"EndTime" name:"EndTime"`
}

func (r *ResetInstancesInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesInternetMaxBandwidthResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728) API返回值中的`InstanceId`获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>`Linux`实例密码必须8到16位，至少包括两项`[a-z，A-Z]、[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>`Windows`实例密码必须12到16位，至少包括三项`[a-z]，[A-Z]，[0-9]`和`[( ) ~ ~ ! @ # $ % ^ & * - + = _ | { } [ ] : ; ' < > , . ? /]`中的符号。密码不允许以`/`符号开头。<br><li>如果实例即包含`Linux`实例又包含`Windows`实例，则密码复杂度限制按照`Windows`实例的限制。
	Password *string `json:"Password" name:"Password"`
	// 待重置密码的实例操作系统用户名。不得超过64个字符。
	UserName *string `json:"UserName" name:"UserName"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *ResetInstancesPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为1。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *ResetInstancesTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetInstancesTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetInstancesTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetInstancesTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksRequest struct {
	*tchttp.BaseRequest
	// 待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且[数据盘类型](/document/api/213/9452#block_device)为：`CLOUD_BASIC`、`CLOUD_PREMIUM`、`CLOUD_SSD`。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型`InstanceType`限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	DataDisks []*DataDisk `json:"DataDisks" name:"DataDisks" list`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：<br><li>TRUE：表示在正常关机失败后进行强制关机<br><li>FALSE：表示在正常关机失败后不进行强制关机<br><br>默认取值：FALSE。<br><br>强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *ResizeInstanceDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResizeInstanceDisksResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResizeInstanceDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResizeInstanceDisksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesRequest struct {
	*tchttp.BaseRequest
	// 实例[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>CDHPAID：独享母机付费（基于专用宿主机创建，宿主机部分的资源不收费）<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType" name:"InstanceChargeType"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid" name:"InstanceChargePrepaid"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，专用宿主机（对于独享母机付费模式的子机创建）等属性。
	Placement *Placement `json:"Placement" name:"Placement"`
	// 实例机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID_BY_HOUR的子机创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。若不指定该参数，则默认机型为S1.SMALL1。<br><li>对于付费模式为CDHPAID的子机创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的子机，该参数应该为CDH_1C1G。
	InstanceType *string `json:"InstanceType" name:"InstanceType"`
	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId" name:"ImageId"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk" name:"SystemDisk"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，当前仅支持购买的时候指定一个数据盘。
	DataDisks []*DataDisk `json:"DataDisks" name:"DataDisks" list`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，表示每个实例的主网卡ip，而且InstanceCount参数必须与私有网络ip的个数一致。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud" name:"VirtualPrivateCloud"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible" name:"InternetAccessible"`
	// 购买实例数量。取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见[CVM实例购买限制](https://cloud.tencent.com/document/product/213/2664)。
	InstanceCount *int64 `json:"InstanceCount" name:"InstanceCount"`
	// 实例显示名称。<br><li>不指定实例显示名称则默认显示‘未命名’。</li><li>购买多台实例，如果指定模式串`{R:x}`，表示生成数字`[x, x+n-1]`，其中`n`表示购买实例的数量，例如`server_{R:3}`，购买1台时，实例显示名称为`server_3`；购买2台时，实例显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。</li><li>购买多台实例，如果不指定模式串，则在实例显示名称添加后缀`1、2...n`，其中`n`表示购买实例的数量，例如`server_`，购买2台时，实例显示名称分别为`server_1`，`server_2`。
	InstanceName *string `json:"InstanceName" name:"InstanceName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds" name:"SecurityGroupIds" list`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService" name:"EnhancedService"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。<br>更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `json:"ClientToken" name:"ClientToken"`
	// 云服务器的主机名。<br><li>点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。<br><li>Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。<br><li>其他类型（Linux 等）实例：字符长度为[2, 30]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `json:"HostName" name:"HostName"`
	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `json:"ActionTimer" name:"ActionTimer"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云主机实例。
	TagSpecification []*TagSpecification `json:"TagSpecification" name:"TagSpecification" list`
}

func (r *RunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 [DescribeInstancesStatus](https://cloud.tencent.com/document/api/213/15738) 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。
		InstanceIdSet []*string `json:"InstanceIdSet" name:"InstanceIdSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {
	// 是否开启[云监控](/document/product/248)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启[云安全](/document/product/296)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	Enabled *bool `json:"Enabled" name:"Enabled"`
}

type SharePermission struct {
	// 镜像分享时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 镜像分享的账户ID
	Account *string `json:"Account" name:"Account"`
}

type SpotMarketOptions struct {
	// 竞价出价
	MaxPrice *string `json:"MaxPrice" name:"MaxPrice"`
	// 竞价请求类型
	SpotInstanceType *string `json:"SpotInstanceType" name:"SpotInstanceType"`
}

type StartInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
}

func (r *StartInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
	// 是否在正常关闭失败后选择强制关闭实例。取值范围：<br><li>TRUE：表示在正常关闭失败后进行强制关闭<br><li>FALSE：表示在正常关闭失败后不进行强制关闭<br><br>默认取值：FALSE。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
	// 实例的关闭模式。取值范围：<br><li>SOFT_FIRST：表示在正常关闭失败后进行强制关闭<br><li>HARD：直接强制关闭<br><li>SOFT：仅软关机<br>默认取值：SOFT。
	StopType *string `json:"StopType" name:"StopType"`
	// 关机收费模式<br><li>KEEP_CHARGING：关机继续收费<br><li>STOP_CHARGING：关机停止收费<br>默认取值：KEEP_CHARGING。
	StoppedMode *string `json:"StoppedMode" name:"StoppedMode"`
}

func (r *StopInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncImagesRequest struct {
	*tchttp.BaseRequest
	// 镜像ID列表 ，镜像ID可以通过如下方式获取：<br><li>通过[DescribeImages](https://cloud.tencent.com/document/api/213/15715)接口返回的`ImageId`获取。<br><li>通过[镜像控制台](https://console.cloud.tencent.com/cvm/image)获取。<br>镜像ID必须满足限制：<br><li>镜像ID对应的镜像状态必须为`NORMAL`。<br><li>镜像大小小于50GB。<br>镜像状态请参考[镜像数据表](/document/api/213/9452#image_state)。
	ImageIds []*string `json:"ImageIds" name:"ImageIds" list`
	// 目的同步地域列表；必须满足限制：<br><li>不能为源地域，<br><li>必须是一个合法的Region。<br><li>暂不支持部分地域同步。<br>具体地域参数请参考[Region](https://cloud.tencent.com/document/product/213/6091)。
	DestinationRegions []*string `json:"DestinationRegions" name:"DestinationRegions" list`
}

func (r *SyncImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SyncImagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SyncImagesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见[CVM实例配置](/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。
	DiskType *string `json:"DiskType" name:"DiskType"`
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId" name:"DiskId"`
	// 系统盘大小，单位：GB。默认值为 50
	DiskSize *int64 `json:"DiskSize" name:"DiskSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key" name:"Key"`
	// 标签值
	Value *string `json:"Value" name:"Value"`
}

type TagSpecification struct {
	// 标签绑定的资源类型
	ResourceType *string `json:"ResourceType" name:"ResourceType"`
	// 标签对列表
	Tags []*Tag `json:"Tags" name:"Tags" list`
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest
	// 一个或多个待操作的实例ID。可通过[`DescribeInstances`](https://cloud.tencent.com/document/api/213/15728)接口返回值中的`InstanceId`获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds" name:"InstanceIds" list`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigRequest struct {
	*tchttp.BaseRequest
	// 待操作的实例ID。可通过[`DescribeInstances`](document/api/213/9388)接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
	// 私有网络相关信息配置。通过该参数指定私有网络的ID，子网ID，私有网络ip等信息。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud" name:"VirtualPrivateCloud"`
	// 是否对运行中的实例选择强制关机。默认为TRUE。
	ForceStop *bool `json:"ForceStop" name:"ForceStop"`
}

func (r *UpdateInstanceVpcConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceVpcConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateInstanceVpcConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateInstanceVpcConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateInstanceVpcConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](https://console.cloud.tencent.com/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcEx](/document/api/215/1372) ，从接口返回中的`unVpcId`字段获取。
	VpcId *string `json:"VpcId" name:"VpcId"`
	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnetEx](/document/api/215/1371) ，从接口返回中的`unSubnetId`字段获取。
	SubnetId *string `json:"SubnetId" name:"SubnetId"`
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：<br><li>TRUE：表示用作公网网关<br><li>FALSE：表示不用作公网网关<br><br>默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway" name:"AsVpcGateway"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses" name:"PrivateIpAddresses" list`
}

type ZoneInfo struct {
	// 可用区名称，例如，ap-guangzhou-3
	Zone *string `json:"Zone" name:"Zone"`
	// 可用区描述，例如，广州三区
	ZoneName *string `json:"ZoneName" name:"ZoneName"`
	// 可用区ID
	ZoneId *string `json:"ZoneId" name:"ZoneId"`
	// 可用区状态
	ZoneState *string `json:"ZoneState" name:"ZoneState"`
}