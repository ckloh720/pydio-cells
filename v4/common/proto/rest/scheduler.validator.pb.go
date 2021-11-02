// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scheduler.proto

package rest

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pydio/cells/v4/common/proto/jobs"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UserJobRequest) Validate() error {
	return nil
}
func (this *UserJobResponse) Validate() error {
	return nil
}
func (this *UserJobsCollection) Validate() error {
	for _, item := range this.Jobs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Jobs", err)
			}
		}
	}
	return nil
}
