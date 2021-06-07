// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/cloud/retail/v2/product.proto

package retail

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The type of this product.
type Product_Type int32

const (
	// Default value. Default to
	// [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY] if unset.
	Product_TYPE_UNSPECIFIED Product_Type = 0
	// The primary type.
	//
	// As the primary unit for predicting, indexing and search serving, a
	// [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product] is grouped with multiple
	// [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]
	// [Product][google.cloud.retail.v2.Product]s.
	Product_PRIMARY Product_Type = 1
	// The variant type.
	//
	// [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]
	// [Product][google.cloud.retail.v2.Product]s usually share some common
	// attributes on the same
	// [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product]s, but they have variant
	// attributes like different colors, sizes and prices, etc.
	Product_VARIANT Product_Type = 2
	// The collection type. Collection products are bundled
	// [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product]s or
	// [Type.VARIANT][google.cloud.retail.v2.Product.Type.VARIANT]
	// [Product][google.cloud.retail.v2.Product]s that are sold together, such
	// as a jewelry set with necklaces, earrings and rings, etc.
	Product_COLLECTION Product_Type = 3
)

// Enum value maps for Product_Type.
var (
	Product_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PRIMARY",
		2: "VARIANT",
		3: "COLLECTION",
	}
	Product_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"PRIMARY":          1,
		"VARIANT":          2,
		"COLLECTION":       3,
	}
)

func (x Product_Type) Enum() *Product_Type {
	p := new(Product_Type)
	*p = x
	return p
}

func (x Product_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Product_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2_product_proto_enumTypes[0].Descriptor()
}

func (Product_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2_product_proto_enumTypes[0]
}

func (x Product_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product_Type.Descriptor instead.
func (Product_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_proto_rawDescGZIP(), []int{0, 0}
}

// Product availability. If this field is unspecified, the product is
// assumed to be in stock.
type Product_Availability int32

const (
	// Default product availability. Default to
	// [Availability.IN_STOCK][google.cloud.retail.v2.Product.Availability.IN_STOCK]
	// if unset.
	Product_AVAILABILITY_UNSPECIFIED Product_Availability = 0
	// Product in stock.
	Product_IN_STOCK Product_Availability = 1
	// Product out of stock.
	Product_OUT_OF_STOCK Product_Availability = 2
	// Product that is in pre-order state.
	Product_PREORDER Product_Availability = 3
	// Product that is back-ordered (i.e. temporarily out of stock).
	Product_BACKORDER Product_Availability = 4
)

// Enum value maps for Product_Availability.
var (
	Product_Availability_name = map[int32]string{
		0: "AVAILABILITY_UNSPECIFIED",
		1: "IN_STOCK",
		2: "OUT_OF_STOCK",
		3: "PREORDER",
		4: "BACKORDER",
	}
	Product_Availability_value = map[string]int32{
		"AVAILABILITY_UNSPECIFIED": 0,
		"IN_STOCK":                 1,
		"OUT_OF_STOCK":             2,
		"PREORDER":                 3,
		"BACKORDER":                4,
	}
)

func (x Product_Availability) Enum() *Product_Availability {
	p := new(Product_Availability)
	*p = x
	return p
}

func (x Product_Availability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Product_Availability) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2_product_proto_enumTypes[1].Descriptor()
}

func (Product_Availability) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2_product_proto_enumTypes[1]
}

func (x Product_Availability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product_Availability.Descriptor instead.
func (Product_Availability) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_proto_rawDescGZIP(), []int{0, 1}
}

// Product captures all metadata information of items to be recommended or
// searched.
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Full resource name of the product, such as
	// `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/product_id`.
	//
	// The branch ID must be "default_branch".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Immutable. [Product][google.cloud.retail.v2.Product] identifier, which is
	// the final component of [name][google.cloud.retail.v2.Product.name]. For
	// example, this field is "id_1", if
	// [name][google.cloud.retail.v2.Product.name] is
	// `projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/id_1`.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	// Google Merchant Center property
	// [id](https://support.google.com/merchants/answer/6324405). Schema.org
	// Property [Product.sku](https://schema.org/sku).
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. The type of the product. This field is output-only.
	Type Product_Type `protobuf:"varint,3,opt,name=type,proto3,enum=google.cloud.retail.v2.Product_Type" json:"type,omitempty"`
	// Variant group identifier. Must be an
	// [id][google.cloud.retail.v2.Product.id], with the same parent branch with
	// this product. Otherwise, an error is thrown.
	//
	// For [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product]s, this field can only be empty or
	// set to the same value as [id][google.cloud.retail.v2.Product.id].
	//
	// For VARIANT [Product][google.cloud.retail.v2.Product]s, this field cannot
	// be empty. A maximum of 2,000 products are allowed to share the same
	// [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product]. Otherwise, an INVALID_ARGUMENT
	// error is returned.
	//
	// Google Merchant Center Property
	// [item_group_id](https://support.google.com/merchants/answer/6324507).
	// Schema.org Property
	// [Product.inProductGroupWithID](https://schema.org/inProductGroupWithID).
	//
	// This field must be enabled before it can be used. [Learn
	// more](/recommendations-ai/docs/catalog#item-group-id).
	PrimaryProductId string `protobuf:"bytes,4,opt,name=primary_product_id,json=primaryProductId,proto3" json:"primary_product_id,omitempty"`
	// Product categories. This field is repeated for supporting one product
	// belonging to several parallel categories. Strongly recommended using the
	// full path for better search / recommendation quality.
	//
	//
	// To represent full path of category, use '>' sign to separate different
	// hierarchies. If '>' is part of the category name, please replace it with
	// other character(s).
	//
	// For example, if a shoes product belongs to both
	// ["Shoes & Accessories" -> "Shoes"] and
	// ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be
	// represented as:
	//
	//      "categories": [
	//        "Shoes & Accessories > Shoes",
	//        "Sports & Fitness > Athletic Clothing > Shoes"
	//      ]
	//
	// Must be set for [Type.PRIMARY][google.cloud.retail.v2.Product.Type.PRIMARY]
	// [Product][google.cloud.retail.v2.Product] otherwise an INVALID_ARGUMENT
	// error is returned.
	//
	// At most 250 values are allowed per
	// [Product][google.cloud.retail.v2.Product]. Empty values are not allowed.
	// Each value must be a UTF-8 encoded string with a length limit of 5,000
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	// Google Merchant Center property
	// [google_product_category][mc_google_product_category]. Schema.org property
	// [Product.category] (https://schema.org/category).
	//
	// [mc_google_product_category]:
	// https://support.google.com/merchants/answer/6324436
	Categories []string `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	// Required. Product title.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	// Google Merchant Center property
	// [title](https://support.google.com/merchants/answer/6324415). Schema.org
	// property [Product.name](https://schema.org/name).
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	// Product description.
	//
	// This field must be a UTF-8 encoded string with a length limit of 5,000
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	// Google Merchant Center property
	// [description](https://support.google.com/merchants/answer/6324468).
	// schema.org property [Product.description](https://schema.org/description).
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	// Highly encouraged. Extra product attributes to be included. For example,
	// for products, this could include the store name, vendor, style, color, etc.
	// These are very strong signals for recommendation model, thus we highly
	// recommend providing the attributes here.
	//
	// Features that can take on one of a limited number of possible values. Two
	// types of features can be set are:
	//
	// Textual features. some examples would be the brand/maker of a product, or
	// country of a customer. Numerical features. Some examples would be the
	// height/weight of a product, or age of a customer.
	//
	// For example: `{ "vendor": {"text": ["vendor123", "vendor456"]},
	// "lengths_cm": {"numbers":[2.3, 15.4]}, "heights_cm": {"numbers":[8.1, 6.4]}
	// }`.
	//
	// A maximum of 150 attributes are allowed. Otherwise, an INVALID_ARGUMENT
	// error is returned.
	//
	// The key must be a UTF-8 encoded string with a length limit of 5,000
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	Attributes map[string]*CustomAttribute `protobuf:"bytes,12,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Custom tags associated with the product.
	//
	// At most 250 values are allowed per
	// [Product][google.cloud.retail.v2.Product]. This value must be a UTF-8
	// encoded string with a length limit of 1,000 characters. Otherwise, an
	// INVALID_ARGUMENT error is returned.
	//
	// This tag can be used for filtering recommendation results by passing the
	// tag as part of the
	// [PredictRequest.filter][google.cloud.retail.v2.PredictRequest.filter].
	//
	// Google Merchant Center property
	// [custom_label_0–4](https://support.google.com/merchants/answer/6324473).
	Tags []string `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	// Product price and cost information.
	//
	// Google Merchant Center property
	// [price](https://support.google.com/merchants/answer/6324371).
	PriceInfo *PriceInfo `protobuf:"bytes,14,opt,name=price_info,json=priceInfo,proto3" json:"price_info,omitempty"`
	// The timestamp when this [Product][google.cloud.retail.v2.Product] becomes
	// available recommendation and search.
	AvailableTime *timestamppb.Timestamp `protobuf:"bytes,18,opt,name=available_time,json=availableTime,proto3" json:"available_time,omitempty"`
	// The online availability of the [Product][google.cloud.retail.v2.Product].
	// Default to
	// [Availability.IN_STOCK][google.cloud.retail.v2.Product.Availability.IN_STOCK].
	//
	// Google Merchant Center Property
	// [availability](https://support.google.com/merchants/answer/6324448).
	// Schema.org Property [Offer.availability](https://schema.org/availability).
	Availability Product_Availability `protobuf:"varint,19,opt,name=availability,proto3,enum=google.cloud.retail.v2.Product_Availability" json:"availability,omitempty"`
	// The available quantity of the item.
	AvailableQuantity *wrapperspb.Int32Value `protobuf:"bytes,20,opt,name=available_quantity,json=availableQuantity,proto3" json:"available_quantity,omitempty"`
	// Canonical URL directly linking to the product detail page.
	//
	// This field must be a UTF-8 encoded string with a length limit of 5,000
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	//
	// Google Merchant Center property
	// [link](https://support.google.com/merchants/answer/6324416). Schema.org
	// property [Offer.url](https://schema.org/url).
	Uri string `protobuf:"bytes,22,opt,name=uri,proto3" json:"uri,omitempty"`
	// Product images for the product.
	//
	// A maximum of 300 images are allowed.
	//
	// Google Merchant Center property
	// [image_link](https://support.google.com/merchants/answer/6324350).
	// Schema.org property [Product.image](https://schema.org/image).
	Images []*Image `protobuf:"bytes,23,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetType() Product_Type {
	if x != nil {
		return x.Type
	}
	return Product_TYPE_UNSPECIFIED
}

func (x *Product) GetPrimaryProductId() string {
	if x != nil {
		return x.PrimaryProductId
	}
	return ""
}

func (x *Product) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetAttributes() map[string]*CustomAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Product) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Product) GetPriceInfo() *PriceInfo {
	if x != nil {
		return x.PriceInfo
	}
	return nil
}

func (x *Product) GetAvailableTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AvailableTime
	}
	return nil
}

func (x *Product) GetAvailability() Product_Availability {
	if x != nil {
		return x.Availability
	}
	return Product_AVAILABILITY_UNSPECIFIED
}

func (x *Product) GetAvailableQuantity() *wrapperspb.Int32Value {
	if x != nil {
		return x.AvailableQuantity
	}
	return nil
}

func (x *Product) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Product) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

var File_google_cloud_retail_v2_product_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2_product_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf4, 0x08, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x0e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x50,
	0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x4a, 0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x35,
	0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x66, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x03, 0x22, 0x69, 0x0a, 0x0c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x4f, 0x43,
	0x4b, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x41, 0x43, 0x4b, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x04,
	0x3a, 0x84, 0x01, 0xea, 0x41, 0x80, 0x01, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x7d, 0x42, 0xb0, 0x02, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x16,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0xea, 0x41, 0x6c, 0x0a, 0x1c,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x4c, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2_product_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2_product_proto_rawDescData = file_google_cloud_retail_v2_product_proto_rawDesc
)

func file_google_cloud_retail_v2_product_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2_product_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2_product_proto_rawDescData)
	})
	return file_google_cloud_retail_v2_product_proto_rawDescData
}

var file_google_cloud_retail_v2_product_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_retail_v2_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_retail_v2_product_proto_goTypes = []interface{}{
	(Product_Type)(0),             // 0: google.cloud.retail.v2.Product.Type
	(Product_Availability)(0),     // 1: google.cloud.retail.v2.Product.Availability
	(*Product)(nil),               // 2: google.cloud.retail.v2.Product
	nil,                           // 3: google.cloud.retail.v2.Product.AttributesEntry
	(*PriceInfo)(nil),             // 4: google.cloud.retail.v2.PriceInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*wrapperspb.Int32Value)(nil), // 6: google.protobuf.Int32Value
	(*Image)(nil),                 // 7: google.cloud.retail.v2.Image
	(*CustomAttribute)(nil),       // 8: google.cloud.retail.v2.CustomAttribute
}
var file_google_cloud_retail_v2_product_proto_depIdxs = []int32{
	0, // 0: google.cloud.retail.v2.Product.type:type_name -> google.cloud.retail.v2.Product.Type
	3, // 1: google.cloud.retail.v2.Product.attributes:type_name -> google.cloud.retail.v2.Product.AttributesEntry
	4, // 2: google.cloud.retail.v2.Product.price_info:type_name -> google.cloud.retail.v2.PriceInfo
	5, // 3: google.cloud.retail.v2.Product.available_time:type_name -> google.protobuf.Timestamp
	1, // 4: google.cloud.retail.v2.Product.availability:type_name -> google.cloud.retail.v2.Product.Availability
	6, // 5: google.cloud.retail.v2.Product.available_quantity:type_name -> google.protobuf.Int32Value
	7, // 6: google.cloud.retail.v2.Product.images:type_name -> google.cloud.retail.v2.Image
	8, // 7: google.cloud.retail.v2.Product.AttributesEntry.value:type_name -> google.cloud.retail.v2.CustomAttribute
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2_product_proto_init() }
func file_google_cloud_retail_v2_product_proto_init() {
	if File_google_cloud_retail_v2_product_proto != nil {
		return
	}
	file_google_cloud_retail_v2_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_google_cloud_retail_v2_product_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2_product_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2_product_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2_product_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2_product_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2_product_proto = out.File
	file_google_cloud_retail_v2_product_proto_rawDesc = nil
	file_google_cloud_retail_v2_product_proto_goTypes = nil
	file_google_cloud_retail_v2_product_proto_depIdxs = nil
}
