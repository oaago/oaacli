syntax = "proto3";

package rpc.{{.Package}}.{{.Method}};

// 多语言特定包名，用于源代码引用
option go_package = "rpc/{{.Package}}/{{.Method}}";
option java_multiple_files = true;
option java_package = "rpc/{{.Package}}/{{.Method}}";
option objc_class_prefix = "rpc/{{.Package}}/{{.Method}}";
import "github.com/mwitkow/go-proto-validators@v0.3.2/validator.proto";

// 描述该服务的信息
service {{.UpPackage}}{{.UpMethod}} {
  // 描述该方法的功能
  rpc Rpc{{.UpPackage}}{{.UpMethod}}Service ({{.UpPackage}}{{.UpMethod}}Request) returns ({{.UpPackage}}{{.UpMethod}}Reply);
}

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
}
message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z]{2,5}$"}];
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}
// Hello请求参数
message {{.UpPackage}}{{.UpMethod}}Request {
  // 用户名字
  string name = 1;
  InnerMessage innerMessage = 2;
}
// Hello返回结果
message {{.UpPackage}}{{.UpMethod}}Reply {
  // 结果信息
  string message = 1;
  OuterMessage outerMessage = 2;
}