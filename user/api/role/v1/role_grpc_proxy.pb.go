// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: user/api/role/v1/role.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ RoleServiceServer = (*roleServiceClientProxy)(nil)

// roleServiceClientProxy is the proxy to turn RoleService client to server interface.
type roleServiceClientProxy struct {
	cc RoleServiceClient
}

func NewRoleServiceClientProxy(cc RoleServiceClient) RoleServiceServer {
	return &roleServiceClientProxy{cc}
}

func (c *roleServiceClientProxy) ListRoles(ctx context.Context, in *ListRolesRequest) (*ListRolesResponse, error) {
	return c.cc.ListRoles(ctx, in)
}
func (c *roleServiceClientProxy) GetRole(ctx context.Context, in *GetRoleRequest) (*Role, error) {
	return c.cc.GetRole(ctx, in)
}
func (c *roleServiceClientProxy) CreateRole(ctx context.Context, in *CreateRoleRequest) (*Role, error) {
	return c.cc.CreateRole(ctx, in)
}
func (c *roleServiceClientProxy) UpdateRole(ctx context.Context, in *UpdateRoleRequest) (*Role, error) {
	return c.cc.UpdateRole(ctx, in)
}
func (c *roleServiceClientProxy) DeleteRole(ctx context.Context, in *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	return c.cc.DeleteRole(ctx, in)
}
func (c *roleServiceClientProxy) GetRolePermission(ctx context.Context, in *GetRolePermissionRequest) (*GetRolePermissionResponse, error) {
	return c.cc.GetRolePermission(ctx, in)
}
func (c *roleServiceClientProxy) UpdateRolePermission(ctx context.Context, in *UpdateRolePermissionRequest) (*UpdateRolePermissionResponse, error) {
	return c.cc.UpdateRolePermission(ctx, in)
}
