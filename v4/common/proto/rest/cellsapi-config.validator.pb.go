// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cellsapi-config.proto

package rest

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/pydio/cells/v4/common/proto/tree"
	_ "github.com/pydio/cells/v4/common/proto/object"
	_ "github.com/pydio/cells/v4/common/proto/ctl"
	_ "github.com/pydio/cells/v4/common/proto/install"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Configuration) Validate() error {
	return nil
}
func (this *ListDataSourceRequest) Validate() error {
	return nil
}
func (this *DataSourceCollection) Validate() error {
	for _, item := range this.DataSources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DataSources", err)
			}
		}
	}
	return nil
}
func (this *DeleteDataSourceResponse) Validate() error {
	return nil
}
func (this *ListPeersAddressesRequest) Validate() error {
	return nil
}
func (this *ListPeersAddressesResponse) Validate() error {
	return nil
}
func (this *ListPeerFoldersRequest) Validate() error {
	return nil
}
func (this *CreatePeerFolderRequest) Validate() error {
	return nil
}
func (this *CreatePeerFolderResponse) Validate() error {
	if this.Node != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Node); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Node", err)
		}
	}
	return nil
}
func (this *ListStorageBucketsRequest) Validate() error {
	if this.DataSource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DataSource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DataSource", err)
		}
	}
	return nil
}
func (this *Process) Validate() error {
	return nil
}
func (this *ListProcessesRequest) Validate() error {
	return nil
}
func (this *ListProcessesResponse) Validate() error {
	for _, item := range this.Processes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Processes", err)
			}
		}
	}
	return nil
}
func (this *ListVersioningPolicyRequest) Validate() error {
	return nil
}
func (this *VersioningPolicyCollection) Validate() error {
	for _, item := range this.Policies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Policies", err)
			}
		}
	}
	return nil
}
func (this *ListVirtualNodesRequest) Validate() error {
	return nil
}
func (this *ListServiceRequest) Validate() error {
	return nil
}
func (this *ServiceCollection) Validate() error {
	for _, item := range this.Services {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Services", err)
			}
		}
	}
	return nil
}
func (this *ControlServiceRequest) Validate() error {
	return nil
}
func (this *DiscoveryRequest) Validate() error {
	return nil
}
func (this *DiscoveryResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ConfigFormRequest) Validate() error {
	return nil
}
func (this *OpenApiResponse) Validate() error {
	return nil
}
func (this *ActionDescription) Validate() error {
	return nil
}
func (this *SchedulerActionsRequest) Validate() error {
	return nil
}
func (this *SchedulerActionsResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SchedulerActionFormRequest) Validate() error {
	return nil
}
func (this *SchedulerActionFormResponse) Validate() error {
	return nil
}
func (this *ListSitesRequest) Validate() error {
	return nil
}
func (this *ListSitesResponse) Validate() error {
	for _, item := range this.Sites {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sites", err)
			}
		}
	}
	return nil
}
