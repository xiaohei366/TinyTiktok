// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: FavoriteServer.proto

package FavoriteServer

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 用户id
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                         // 用户名称
	FollowCount    int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`       // 关注总数--可选
	FollowerCount  int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"` // 粉丝总数--可选
	IsFollow       bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                // true-已关注，false-未关注
	TotalFavorited int64  `protobuf:"varint,6,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"`                    // 被赞的总次数
	FavoriteCount  int64  `protobuf:"varint,7,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`                      // 喜欢总数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题:
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述--可选
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{2}
}

func (x *BaseResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// 点赞/取消赞
type DouyinFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                   // 用户id
	VideoId    int64 `protobuf:"varint,3,opt,name=videoId,proto3" json:"videoId,omitempty"`                         // 视频id
	ActionType int32 `protobuf:"varint,4,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (x *DouyinFavoriteActionRequest) Reset() {
	*x = DouyinFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionRequest) ProtoMessage() {}

func (x *DouyinFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinFavoriteActionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type DouyinFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *DouyinFavoriteActionResponse) Reset() {
	*x = DouyinFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionResponse) ProtoMessage() {}

func (x *DouyinFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{4}
}

func (x *DouyinFavoriteActionResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// 获取点赞列表
type DouyinFavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *DouyinFavoriteListRequest) Reset() {
	*x = DouyinFavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListRequest) ProtoMessage() {}

func (x *DouyinFavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{5}
}

func (x *DouyinFavoriteListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DouyinFavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp  *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	VideoList []*Video  `protobuf:"bytes,2,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"` // 视频列表
}

func (x *DouyinFavoriteListResponse) Reset() {
	*x = DouyinFavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListResponse) ProtoMessage() {}

func (x *DouyinFavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{6}
}

func (x *DouyinFavoriteListResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *DouyinFavoriteListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type DouyinUserFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *DouyinUserFavoriteRequest) Reset() {
	*x = DouyinUserFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserFavoriteRequest) ProtoMessage() {}

func (x *DouyinUserFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserFavoriteRequest.ProtoReflect.Descriptor instead.
func (*DouyinUserFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{7}
}

func (x *DouyinUserFavoriteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DouyinUserFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp       *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	TotalFavorited int64     `protobuf:"varint,2,opt,name=TotalFavorited,proto3" json:"TotalFavorited,omitempty"` // 被赞的总次数
	FavoriteCount  int64     `protobuf:"varint,3,opt,name=FavoriteCount,proto3" json:"FavoriteCount,omitempty"`   // 喜欢总数量
}

func (x *DouyinUserFavoriteResponse) Reset() {
	*x = DouyinUserFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinUserFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinUserFavoriteResponse) ProtoMessage() {}

func (x *DouyinUserFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinUserFavoriteResponse.ProtoReflect.Descriptor instead.
func (*DouyinUserFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{8}
}

func (x *DouyinUserFavoriteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *DouyinUserFavoriteResponse) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *DouyinUserFavoriteResponse) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type DouyinVideoFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 视频唯一标识
}

func (x *DouyinVideoFavoriteRequest) Reset() {
	*x = DouyinVideoFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinVideoFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinVideoFavoriteRequest) ProtoMessage() {}

func (x *DouyinVideoFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinVideoFavoriteRequest.ProtoReflect.Descriptor instead.
func (*DouyinVideoFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{9}
}

func (x *DouyinVideoFavoriteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DouyinVideoFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp      *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	FavoriteCount int64     `protobuf:"varint,2,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	IsFavorite    bool      `protobuf:"varint,3,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
}

func (x *DouyinVideoFavoriteResponse) Reset() {
	*x = DouyinVideoFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FavoriteServer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinVideoFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinVideoFavoriteResponse) ProtoMessage() {}

func (x *DouyinVideoFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FavoriteServer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinVideoFavoriteResponse.ProtoReflect.Descriptor instead.
func (*DouyinVideoFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_FavoriteServer_proto_rawDescGZIP(), []int{10}
}

func (x *DouyinVideoFavoriteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *DouyinVideoFavoriteResponse) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *DouyinVideoFavoriteResponse) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

var File_FavoriteServer_proto protoreflect.FileDescriptor

var file_FavoriteServer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xdf, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x68, 0x0a, 0x1b, 0x44, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x55, 0x0a, 0x1c, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x34, 0x0a, 0x19, 0x44, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x89,
	0x01, 0x0a, 0x1a, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x19, 0x44, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x1b, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x32, 0xc7, 0x03, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x2a, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61, 0x6f, 0x68, 0x65,
	0x69, 0x33, 0x36, 0x36, 0x2f, 0x54, 0x69, 0x6e, 0x79, 0x54, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FavoriteServer_proto_rawDescOnce sync.Once
	file_FavoriteServer_proto_rawDescData = file_FavoriteServer_proto_rawDesc
)

func file_FavoriteServer_proto_rawDescGZIP() []byte {
	file_FavoriteServer_proto_rawDescOnce.Do(func() {
		file_FavoriteServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_FavoriteServer_proto_rawDescData)
	})
	return file_FavoriteServer_proto_rawDescData
}

var file_FavoriteServer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_FavoriteServer_proto_goTypes = []interface{}{
	(*User)(nil),                         // 0: FavoriteServer.User
	(*Video)(nil),                        // 1: FavoriteServer.Video
	(*BaseResp)(nil),                     // 2: FavoriteServer.BaseResp
	(*DouyinFavoriteActionRequest)(nil),  // 3: FavoriteServer.DouyinFavoriteActionRequest
	(*DouyinFavoriteActionResponse)(nil), // 4: FavoriteServer.DouyinFavoriteActionResponse
	(*DouyinFavoriteListRequest)(nil),    // 5: FavoriteServer.DouyinFavoriteListRequest
	(*DouyinFavoriteListResponse)(nil),   // 6: FavoriteServer.DouyinFavoriteListResponse
	(*DouyinUserFavoriteRequest)(nil),    // 7: FavoriteServer.DouyinUserFavoriteRequest
	(*DouyinUserFavoriteResponse)(nil),   // 8: FavoriteServer.DouyinUserFavoriteResponse
	(*DouyinVideoFavoriteRequest)(nil),   // 9: FavoriteServer.DouyinVideoFavoriteRequest
	(*DouyinVideoFavoriteResponse)(nil),  // 10: FavoriteServer.DouyinVideoFavoriteResponse
}
var file_FavoriteServer_proto_depIdxs = []int32{
	0,  // 0: FavoriteServer.Video.author:type_name -> FavoriteServer.User
	2,  // 1: FavoriteServer.DouyinFavoriteActionResponse.base_resp:type_name -> FavoriteServer.BaseResp
	2,  // 2: FavoriteServer.DouyinFavoriteListResponse.base_resp:type_name -> FavoriteServer.BaseResp
	1,  // 3: FavoriteServer.DouyinFavoriteListResponse.video_list:type_name -> FavoriteServer.Video
	2,  // 4: FavoriteServer.DouyinUserFavoriteResponse.base_resp:type_name -> FavoriteServer.BaseResp
	2,  // 5: FavoriteServer.DouyinVideoFavoriteResponse.base_resp:type_name -> FavoriteServer.BaseResp
	3,  // 6: FavoriteServer.FavoriteService.FavoriteAction:input_type -> FavoriteServer.DouyinFavoriteActionRequest
	5,  // 7: FavoriteServer.FavoriteService.GetFavoriteList:input_type -> FavoriteServer.DouyinFavoriteListRequest
	7,  // 8: FavoriteServer.FavoriteService.GetFavoriteUser:input_type -> FavoriteServer.DouyinUserFavoriteRequest
	9,  // 9: FavoriteServer.FavoriteService.GetFavoriteVideo:input_type -> FavoriteServer.DouyinVideoFavoriteRequest
	4,  // 10: FavoriteServer.FavoriteService.FavoriteAction:output_type -> FavoriteServer.DouyinFavoriteActionResponse
	6,  // 11: FavoriteServer.FavoriteService.GetFavoriteList:output_type -> FavoriteServer.DouyinFavoriteListResponse
	8,  // 12: FavoriteServer.FavoriteService.GetFavoriteUser:output_type -> FavoriteServer.DouyinUserFavoriteResponse
	10, // 13: FavoriteServer.FavoriteService.GetFavoriteVideo:output_type -> FavoriteServer.DouyinVideoFavoriteResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_FavoriteServer_proto_init() }
func file_FavoriteServer_proto_init() {
	if File_FavoriteServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FavoriteServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserFavoriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinUserFavoriteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinVideoFavoriteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_FavoriteServer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinVideoFavoriteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FavoriteServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FavoriteServer_proto_goTypes,
		DependencyIndexes: file_FavoriteServer_proto_depIdxs,
		MessageInfos:      file_FavoriteServer_proto_msgTypes,
	}.Build()
	File_FavoriteServer_proto = out.File
	file_FavoriteServer_proto_rawDesc = nil
	file_FavoriteServer_proto_goTypes = nil
	file_FavoriteServer_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type FavoriteService interface {
	FavoriteAction(ctx context.Context, req *DouyinFavoriteActionRequest) (res *DouyinFavoriteActionResponse, err error)
	GetFavoriteList(ctx context.Context, req *DouyinFavoriteListRequest) (res *DouyinFavoriteListResponse, err error)
	GetFavoriteUser(ctx context.Context, req *DouyinUserFavoriteRequest) (res *DouyinUserFavoriteResponse, err error)
	GetFavoriteVideo(ctx context.Context, req *DouyinVideoFavoriteRequest) (res *DouyinVideoFavoriteResponse, err error)
}