/**
 * @author:       wangxuebing
 * @fileName:     alisms.go
 * @date:         2023/2/1 14:23
 * @description:  阿里云短信API接口
 */

package sms_api

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type (
	// AliSmsResp 阿里云发送短信|批量发送短信-响应参数
	AliSmsResp struct {
		Message   string `json:"Message"`   //状态码的描述
		RequestId string `json:"RequestId"` //请求ID
		BizId     string `json:"BizId"`     //发送回执ID
		Code      string `json:"Code"`      //请求状态码[OK:成功;其他错误码:失败]
	}
	// AliSmsQuerySendStatisticsResp 查询短信发送统计信息-响应参数
	AliSmsQuerySendStatisticsResp struct {
		Message   string                            `json:"Message"`   //状态码的描述
		RequestId string                            `json:"RequestId"` //请求ID
		Data      AliSmsQuerySendStatisticsPageResp `json:"data"`      //返回数据
		Code      string                            `json:"Code"`      //请求状态码[OK:成功;其他错误码:失败]
	}
	// AliSmsQuerySendStatisticsPageResp 查询短信发送统计信息-响应参数数据分页
	AliSmsQuerySendStatisticsPageResp struct {
		TotalSize  int                                  `json:"TotalSize"`  //返回数据的总条数
		TargetList []*AliSmsQuerySendStatisticsDataResp `json:"TargetList"` //返回数据列表
	}
	// AliSmsQuerySendStatisticsDataResp 查询短信发送统计信息-响应参数数据
	AliSmsQuerySendStatisticsDataResp struct {
		TotalCount            int    `json:"TotalCount"`            //发送成功的短信条数
		RespondedSuccessCount int    `json:"RespondedSuccessCount"` //接收到回执成功的短信条数
		RespondedFailCount    int    `json:"RespondedFailCount"`    //接收到回执失败的短信条数
		NoRespondedCount      int    `json:"NoRespondedCount"`      //未收到回执的短信条数
		SendDate              string `json:"SendDate"`              //短信发送日期[格式为yyyyMMdd]
	}
)

//阿里云短信平台相关API接口
//

func createAliSmsClient(accessKeyId, accessKeySecret string) (*dysmsapi.Client, error) {
	config := openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("ecs-cn-hangzhou.aliyuncs.com")
	result, err := dysmsapi.NewClient(&config)
	return result, err
}

// SendAliSms 发送短信
// Params:
//
//	accessKeyId:
//	accessKeySecret:
//	signName:短信签名
//	tmpCode:短信模板Code
//	tmpParam:短信模板变量对应的实际值
//	phoneNumbers:电话号码
//
// Returns:
func SendAliSms(accessKeyId, accessKeySecret, signName, tmpCode, tmpParam, phoneNumbers string) (*dysmsapi.SendSmsResponse, error) {
	var resp *dysmsapi.SendSmsResponse
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	smsRequest := dysmsapi.SendSmsRequest{
		PhoneNumbers:  &phoneNumbers,
		SignName:      &signName,
		TemplateCode:  &tmpCode,
		TemplateParam: &tmpParam,
	}
	runtime := &util.RuntimeOptions{}
	runtime.ReadTimeout = tea.Int(10000)
	resp, err = client.SendSmsWithOptions(&smsRequest, runtime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendBatchAliSms 批量发送短信
// Params:
//
//	phoneNumberJson: 接收短信的手机号码(示例: ["1590000****","1350000****"])
//	signNameJson: 短信签名名称(短信签名的个数必须与手机号码的个数相同、内容一一对应)
//	tmpCode: 短信模板CODE
//	tmpParamJson: 短信模板变量对应的实际值(模板变量值的个数必须与手机号码、签名的个数相同、内容一一对应)
//
// Returns:
func SendBatchAliSms(accessKeyId, accessKeySecret string, phoneNumberJson, signNameJson, tmpCode, tmpParamJson string) (*dysmsapi.SendBatchSmsResponse, error) {
	var resp *dysmsapi.SendBatchSmsResponse
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	smsRequest := dysmsapi.SendBatchSmsRequest{
		PhoneNumberJson:   &phoneNumberJson,
		SignNameJson:      &signNameJson,
		TemplateCode:      &tmpCode,
		TemplateParamJson: &tmpParamJson,
	}
	runtime := &util.RuntimeOptions{}
	runtime.ReadTimeout = tea.Int(10000)
	resp, err = client.SendBatchSmsWithOptions(&smsRequest, runtime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryAliSendStatistics 查询短信发送统计信息
// Params:
//
//	isGlobe:   短信发送范围[1: 国内短信发送记录; 2: 国际/港澳台短信发送记录]
//	startDate: 开始日期[格式:yyyyMMdd]
//	endDate:   结束日期[格式:yyyyMMdd]
//	page:      当前页码[默认取值为1]
//	pageSize:  每页显示的条数[取值范围：1~50]
//	tmpType:   模板类型[0: 验证码; 1:通知短信; 2:推广短信(仅支持企业客户); 3:国际/港澳台消息(仅支持企业客户); 7:数字短信;]
//	signName:  签名名称
//
// Returns:
func QueryAliSendStatistics(accessKeyId, accessKeySecret string, isGlobe int32, startDate, endDate string, page, pageSize int32, tmpType *int32, signName *string) (*dysmsapi.QuerySendStatisticsResponse, error) {
	var resp *dysmsapi.QuerySendStatisticsResponse
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	querySendStatisticsReq := dysmsapi.QuerySendStatisticsRequest{
		StartDate:    &startDate,
		EndDate:      &endDate,
		IsGlobe:      &isGlobe,
		PageIndex:    &page,
		PageSize:     &pageSize,
		SignName:     signName,
		TemplateType: tmpType,
	}
	runtime := &util.RuntimeOptions{}
	runtime.ReadTimeout = tea.Int(10000)
	resp, err = client.QuerySendStatisticsWithOptions(&querySendStatisticsReq, runtime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryAliSendDetails 查询短信发送详情
// Params:
//
//	phoneNumber: 接收短信的手机号码
//	bizId:       发送回执ID，即发送流水号
//	sendDate:    短信发送日期，支持查询最近30天的记录
//	pageSize:    分页查看发送记录，指定每页显示的短信记录数量[取值范围:1~50]
//	currentPage: 分页查看发送记录，指定每页显示的短信记录数量
//
// Returns:
func QueryAliSendDetails(accessKeyId, accessKeySecret string, phoneNumber, bizId, sendDate string, pageSize, currentPage int64) (*dysmsapi.QuerySendDetailsResponse, error) {
	var resp *dysmsapi.QuerySendDetailsResponse
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	querySendDetailsReq := dysmsapi.QuerySendDetailsRequest{
		BizId:       &bizId,
		CurrentPage: &currentPage,
		PageSize:    &pageSize,
		PhoneNumber: &phoneNumber,
		SendDate:    &sendDate,
	}
	runtime := &util.RuntimeOptions{}
	runtime.ReadTimeout = tea.Int(10000)
	resp, err = client.QuerySendDetailsWithOptions(&querySendDetailsReq, runtime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
