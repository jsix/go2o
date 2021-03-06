// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foundation_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// * 替换敏感词请求
type ReplaceSensitiveRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text"`
	Replacement          string   `protobuf:"bytes,2,opt,name=Replacement,proto3" json:"Replacement"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceSensitiveRequest) Reset()         { *m = ReplaceSensitiveRequest{} }
func (m *ReplaceSensitiveRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceSensitiveRequest) ProtoMessage()    {}
func (*ReplaceSensitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_6c2a41fd5fca7596, []int{0}
}
func (m *ReplaceSensitiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceSensitiveRequest.Unmarshal(m, b)
}
func (m *ReplaceSensitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceSensitiveRequest.Marshal(b, m, deterministic)
}
func (dst *ReplaceSensitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceSensitiveRequest.Merge(dst, src)
}
func (m *ReplaceSensitiveRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceSensitiveRequest.Size(m)
}
func (m *ReplaceSensitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceSensitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceSensitiveRequest proto.InternalMessageInfo

func (m *ReplaceSensitiveRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ReplaceSensitiveRequest) GetReplacement() string {
	if m != nil {
		return m.Replacement
	}
	return ""
}

func init() {
	proto.RegisterType((*ReplaceSensitiveRequest)(nil), "ReplaceSensitiveRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FoundationServiceClient is the client API for FoundationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoundationServiceClient interface {
	// * 检测是否包含敏感词
	CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error)
	// 获取应用信息,name
	GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 获取地区名称
	GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error)
	// 获取下级区域,code
	GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取支付平台
	GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error)
}

type foundationServiceClient struct {
	cc *grpc.ClientConn
}

func NewFoundationServiceClient(cc *grpc.ClientConn) FoundationServiceClient {
	return &foundationServiceClient{cc}
}

func (c *foundationServiceClient) CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/FoundationService/CheckSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ReplaceSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error) {
	out := new(SSmsApi)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveBoardHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ResourceUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/RegisterApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error) {
	out := new(SSsoApp)
	err := c.cc.Invoke(ctx, "/FoundationService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAllSsoApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error) {
	out := new(SuperLoginResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/SuperValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/FlushSuperPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSyncLoginUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error) {
	out := new(AreaListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetChildAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error) {
	out := new(SMobileAppConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error) {
	out := new(SWxApiConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error) {
	out := new(PaymentPlatformResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetPayPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error) {
	out := new(SGlobMchSaleConf)
	err := c.cc.Invoke(ctx, "/FoundationService/GetGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoundationServiceServer is the server API for FoundationService service.
type FoundationServiceServer interface {
	// * 检测是否包含敏感词
	CheckSensitive(context.Context, *String) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(context.Context, *ReplaceSensitiveRequest) (*String, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(context.Context, *String) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(context.Context, *SmsApiSaveRequest) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(context.Context, *BoardHookSaveRequest) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(context.Context, *String) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(context.Context, *SSsoApp) (*String, error)
	// 获取应用信息,name
	GetApp(context.Context, *String) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(context.Context, *Empty) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(context.Context, *UserPwd) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(context.Context, *UserPwd) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(context.Context, *String) (*String, error)
	// 获取地区名称
	GetAreaNames(context.Context, *GetAreaNamesRequest) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(context.Context, *AreaStringRequest) (*String, error)
	// 获取下级区域,code
	GetChildAreas(context.Context, *Int32) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(context.Context, *Empty) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(context.Context, *SMobileAppConfig) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(context.Context, *Empty) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(context.Context, *SWxApiConfig) (*Result, error)
	// 获取支付平台
	GetPayPlatform(context.Context, *Empty) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(context.Context, *Empty) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(context.Context, *SGlobMchSaleConf) (*Result, error)
}

func RegisterFoundationServiceServer(s *grpc.Server, srv FoundationServiceServer) {
	s.RegisterService(&_FoundationService_serviceDesc, srv)
}

func _FoundationService_CheckSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/CheckSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ReplaceSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceSensitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ReplaceSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, req.(*ReplaceSensitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsApiSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, req.(*SmsApiSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveBoardHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardHookSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveBoardHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, req.(*BoardHookSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ResourceUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ResourceUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_RegisterApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSsoApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).RegisterApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/RegisterApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).RegisterApp(ctx, req.(*SSsoApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetApp(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAllSsoApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAllSsoApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SuperValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SuperValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SuperValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SuperValidate(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_FlushSuperPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/FlushSuperPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSyncLoginUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSyncLoginUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, req.(*GetAreaNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaString(ctx, req.(*AreaStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetChildAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetChildAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMobileAppConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, req.(*SMobileAppConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SWxApiConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, req.(*SWxApiConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetPayPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetPayPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SGlobMchSaleConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, req.(*SGlobMchSaleConf))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoundationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FoundationService",
	HandlerType: (*FoundationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSensitive",
			Handler:    _FoundationService_CheckSensitive_Handler,
		},
		{
			MethodName: "ReplaceSensitive",
			Handler:    _FoundationService_ReplaceSensitive_Handler,
		},
		{
			MethodName: "GetSmsApi",
			Handler:    _FoundationService_GetSmsApi_Handler,
		},
		{
			MethodName: "SaveSmsApi",
			Handler:    _FoundationService_SaveSmsApi_Handler,
		},
		{
			MethodName: "SaveBoardHook",
			Handler:    _FoundationService_SaveBoardHook_Handler,
		},
		{
			MethodName: "ResourceUrl",
			Handler:    _FoundationService_ResourceUrl_Handler,
		},
		{
			MethodName: "RegisterApp",
			Handler:    _FoundationService_RegisterApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _FoundationService_GetApp_Handler,
		},
		{
			MethodName: "GetAllSsoApp",
			Handler:    _FoundationService_GetAllSsoApp_Handler,
		},
		{
			MethodName: "SuperValidate",
			Handler:    _FoundationService_SuperValidate_Handler,
		},
		{
			MethodName: "FlushSuperPwd",
			Handler:    _FoundationService_FlushSuperPwd_Handler,
		},
		{
			MethodName: "GetSyncLoginUrl",
			Handler:    _FoundationService_GetSyncLoginUrl_Handler,
		},
		{
			MethodName: "GetAreaNames",
			Handler:    _FoundationService_GetAreaNames_Handler,
		},
		{
			MethodName: "GetAreaString",
			Handler:    _FoundationService_GetAreaString_Handler,
		},
		{
			MethodName: "GetChildAreas",
			Handler:    _FoundationService_GetChildAreas_Handler,
		},
		{
			MethodName: "GetMoAppConf",
			Handler:    _FoundationService_GetMoAppConf_Handler,
		},
		{
			MethodName: "SaveMoAppConf",
			Handler:    _FoundationService_SaveMoAppConf_Handler,
		},
		{
			MethodName: "GetWxApiConfig",
			Handler:    _FoundationService_GetWxApiConfig_Handler,
		},
		{
			MethodName: "SaveWxApiConfig",
			Handler:    _FoundationService_SaveWxApiConfig_Handler,
		},
		{
			MethodName: "GetPayPlatform",
			Handler:    _FoundationService_GetPayPlatform_Handler,
		},
		{
			MethodName: "GetGlobMchSaleConf_",
			Handler:    _FoundationService_GetGlobMchSaleConf__Handler,
		},
		{
			MethodName: "SaveGlobMchSaleConf_",
			Handler:    _FoundationService_SaveGlobMchSaleConf__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foundation_service.proto",
}

func init() {
	proto.RegisterFile("foundation_service.proto", fileDescriptor_foundation_service_6c2a41fd5fca7596)
}

var fileDescriptor_foundation_service_6c2a41fd5fca7596 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdf, 0x6f, 0xd3, 0x3a,
	0x14, 0xce, 0xae, 0x2e, 0xdd, 0x66, 0xd6, 0x8d, 0x79, 0x43, 0x54, 0x15, 0xa0, 0x11, 0x81, 0x04,
	0x1a, 0x18, 0xe8, 0x90, 0x78, 0xd8, 0x53, 0x37, 0xb1, 0x82, 0xb4, 0x41, 0xd5, 0x30, 0x90, 0x78,
	0x99, 0xdc, 0xe4, 0x34, 0xb5, 0xe6, 0xc4, 0xc6, 0x76, 0xba, 0xf5, 0x0f, 0xe7, 0x1d, 0xd9, 0x49,
	0x53, 0xaf, 0xa5, 0x4f, 0xb5, 0xcf, 0xf7, 0xe3, 0x7c, 0xe7, 0xc8, 0x0d, 0x6a, 0x8d, 0x44, 0x91,
	0x27, 0xd4, 0x30, 0x91, 0x5f, 0x69, 0x50, 0x13, 0x16, 0x03, 0x91, 0x4a, 0x18, 0xd1, 0xde, 0x4a,
	0xb9, 0x18, 0x52, 0x5e, 0xdd, 0x1e, 0x67, 0xa0, 0x35, 0x4d, 0xe1, 0xad, 0xc7, 0x4f, 0x8c, 0x28,
	0xd1, 0xf0, 0x1b, 0x7a, 0x34, 0x00, 0xc9, 0x69, 0x0c, 0x11, 0xe4, 0x9a, 0x19, 0x36, 0x81, 0x01,
	0xfc, 0x2e, 0x40, 0x1b, 0x8c, 0xd1, 0xff, 0xdf, 0xe1, 0xd6, 0xb4, 0xd6, 0x0e, 0xd6, 0x5e, 0x6e,
	0x0e, 0xdc, 0x19, 0x1f, 0xa0, 0xfb, 0x15, 0x3d, 0x83, 0xdc, 0xb4, 0xfe, 0x73, 0x90, 0x5f, 0xea,
	0xfc, 0x59, 0x47, 0xbb, 0x67, 0x75, 0xa7, 0xa8, 0x0c, 0x86, 0x43, 0xb4, 0x7d, 0x3a, 0x86, 0xf8,
	0xba, 0x6e, 0x82, 0xd7, 0x49, 0x64, 0x14, 0xcb, 0xd3, 0xf6, 0x3d, 0x72, 0x22, 0x04, 0x0f, 0x03,
	0xfc, 0x11, 0x3d, 0x58, 0x8c, 0x82, 0x5b, 0x64, 0x45, 0xba, 0xf6, 0x4c, 0x1f, 0x06, 0xf8, 0x00,
	0x6d, 0xf6, 0xc0, 0x44, 0x99, 0xee, 0x4a, 0x36, 0xf7, 0xdd, 0x20, 0x51, 0x59, 0x0a, 0x03, 0x7c,
	0x88, 0x50, 0x44, 0x27, 0x50, 0x51, 0x30, 0x29, 0x0f, 0xb6, 0x34, 0xb7, 0x1b, 0x80, 0x2e, 0xb8,
	0x09, 0x03, 0xfc, 0x1e, 0x35, 0x2d, 0x72, 0x22, 0xa8, 0x4a, 0x3e, 0x0b, 0x71, 0x8d, 0x1f, 0x92,
	0xfa, 0xbc, 0x42, 0xf2, 0xcc, 0xae, 0x45, 0x8b, 0x42, 0xc5, 0x70, 0xa9, 0xf8, 0x3c, 0x83, 0x17,
	0x32, 0xb4, 0x94, 0x94, 0x69, 0x03, 0xaa, 0x2b, 0x25, 0xb6, 0xe9, 0xb4, 0xe8, 0x4a, 0xe9, 0x73,
	0x9e, 0xa0, 0x46, 0x0f, 0x8c, 0x85, 0xef, 0x4c, 0xe1, 0x78, 0x61, 0x80, 0xdf, 0xa0, 0x2d, 0x0b,
	0x73, 0x5e, 0x56, 0x70, 0x83, 0x7c, 0xca, 0xa4, 0x99, 0xb6, 0xf7, 0x2a, 0xf2, 0x39, 0xd3, 0x66,
	0x00, 0x5a, 0x8a, 0x5c, 0x43, 0x18, 0xe0, 0x77, 0xa8, 0x19, 0x15, 0x12, 0xd4, 0x0f, 0xca, 0x59,
	0x42, 0x0d, 0xe0, 0x0d, 0x72, 0xa9, 0x41, 0xf5, 0x6f, 0x12, 0xab, 0xb0, 0xc8, 0xb9, 0x48, 0x59,
	0xee, 0x29, 0x9e, 0xa3, 0xe6, 0x19, 0x2f, 0xf4, 0xd8, 0x81, 0xfd, 0x9b, 0xc4, 0x53, 0x78, 0xc3,
	0xbe, 0x40, 0x3b, 0x76, 0xdd, 0xd3, 0x3c, 0x76, 0xfa, 0x55, 0x03, 0x1f, 0x97, 0x69, 0x15, 0xd0,
	0xaf, 0x34, 0x03, 0x8d, 0xf7, 0x89, 0x7f, 0x9d, 0x2d, 0x71, 0x45, 0x76, 0x82, 0x9a, 0x15, 0xbb,
	0x84, 0x31, 0x26, 0xf3, 0xcb, 0x3f, 0x9e, 0xc0, 0x6b, 0xc7, 0x3f, 0x1d, 0x33, 0x9e, 0x58, 0x9e,
	0xc6, 0x0d, 0xf2, 0x25, 0x37, 0x47, 0x9d, 0xf6, 0xae, 0xd3, 0x2d, 0xb8, 0x1f, 0xba, 0x68, 0x17,
	0x76, 0x89, 0xa7, 0x22, 0x1f, 0xd5, 0x8b, 0xdc, 0x25, 0xd1, 0x85, 0x18, 0x32, 0x0e, 0x15, 0xc0,
	0x52, 0xb7, 0x75, 0xf7, 0x1c, 0xe6, 0xec, 0x65, 0x96, 0xbf, 0x9d, 0x57, 0x68, 0xbb, 0x07, 0xe6,
	0xe7, 0x6d, 0x57, 0xb2, 0x12, 0xac, 0xdd, 0x9b, 0x24, 0xf2, 0xca, 0x2e, 0xc6, 0x8e, 0x75, 0xf6,
	0xb9, 0x77, 0x39, 0xbe, 0xef, 0x07, 0xe7, 0xdb, 0xa7, 0xd3, 0x3e, 0xa7, 0x66, 0x24, 0x54, 0x56,
	0xfb, 0xb6, 0x48, 0x9f, 0x4e, 0xed, 0x7f, 0x6f, 0x86, 0x78, 0x93, 0x76, 0xd0, 0x5e, 0x0f, 0x4c,
	0x8f, 0x8b, 0xe1, 0x45, 0x3c, 0x8e, 0x28, 0x07, 0x6b, 0x7c, 0xe5, 0x0f, 0xbc, 0x80, 0xb9, 0x4e,
	0xfb, 0x36, 0xd6, 0x92, 0x68, 0x99, 0xec, 0xe5, 0x3b, 0x79, 0x8a, 0xf6, 0x62, 0x91, 0x91, 0x94,
	0x99, 0x71, 0x31, 0x24, 0xa9, 0xe8, 0x08, 0xa2, 0x64, 0xfc, 0x6b, 0x9d, 0x1c, 0xbb, 0x0f, 0xcd,
	0xb0, 0xe1, 0x7e, 0x8e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xff, 0x1f, 0x45, 0xb7, 0x04,
	0x00, 0x00,
}
