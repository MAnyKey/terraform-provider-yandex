// Code generated by protoc-gen-goext. DO NOT EDIT.

package k8s

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetNodeGroupRequest) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *ListNodeGroupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListNodeGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListNodeGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListNodeGroupsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListNodeGroupsResponse) SetNodeGroups(v []*NodeGroup) {
	m.NodeGroups = v
}

func (m *ListNodeGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListNodeGroupNodesRequest) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *ListNodeGroupNodesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListNodeGroupNodesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListNodeGroupNodesResponse) SetNodes(v []*Node) {
	m.Nodes = v
}

func (m *ListNodeGroupNodesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeleteNodeGroupRequest) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *DeleteNodeGroupMetadata) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *UpdateNodeGroupRequest) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *UpdateNodeGroupRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateNodeGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateNodeGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateNodeGroupRequest) SetNodeTemplate(v *NodeTemplate) {
	m.NodeTemplate = v
}

func (m *UpdateNodeGroupRequest) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *UpdateNodeGroupRequest) SetAllocationPolicy(v *NodeGroupAllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *UpdateNodeGroupRequest) SetDeployPolicy(v *DeployPolicy) {
	m.DeployPolicy = v
}

func (m *UpdateNodeGroupRequest) SetVersion(v *UpdateVersionSpec) {
	m.Version = v
}

func (m *UpdateNodeGroupRequest) SetMaintenancePolicy(v *NodeGroupMaintenancePolicy) {
	m.MaintenancePolicy = v
}

func (m *UpdateNodeGroupRequest) SetAllowedUnsafeSysctls(v []string) {
	m.AllowedUnsafeSysctls = v
}

func (m *UpdateNodeGroupRequest) SetNodeTaints(v []*Taint) {
	m.NodeTaints = v
}

func (m *UpdateNodeGroupRequest) SetNodeLabels(v map[string]string) {
	m.NodeLabels = v
}

func (m *UpdateNodeGroupMetadata) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *CreateNodeGroupRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateNodeGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateNodeGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateNodeGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateNodeGroupRequest) SetNodeTemplate(v *NodeTemplate) {
	m.NodeTemplate = v
}

func (m *CreateNodeGroupRequest) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *CreateNodeGroupRequest) SetAllocationPolicy(v *NodeGroupAllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *CreateNodeGroupRequest) SetDeployPolicy(v *DeployPolicy) {
	m.DeployPolicy = v
}

func (m *CreateNodeGroupRequest) SetVersion(v string) {
	m.Version = v
}

func (m *CreateNodeGroupRequest) SetMaintenancePolicy(v *NodeGroupMaintenancePolicy) {
	m.MaintenancePolicy = v
}

func (m *CreateNodeGroupRequest) SetAllowedUnsafeSysctls(v []string) {
	m.AllowedUnsafeSysctls = v
}

func (m *CreateNodeGroupRequest) SetNodeTaints(v []*Taint) {
	m.NodeTaints = v
}

func (m *CreateNodeGroupRequest) SetNodeLabels(v map[string]string) {
	m.NodeLabels = v
}

func (m *CreateNodeGroupMetadata) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *AutoUpgradeNodeGroupMetadata) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *ListNodeGroupOperationsRequest) SetNodeGroupId(v string) {
	m.NodeGroupId = v
}

func (m *ListNodeGroupOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListNodeGroupOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListNodeGroupOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListNodeGroupOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListNodeGroupOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
