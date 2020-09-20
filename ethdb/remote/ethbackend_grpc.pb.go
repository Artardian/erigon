// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ETHBACKENDClient is the client API for ETHBACKEND service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ETHBACKENDClient interface {
	Add(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*AddReply, error)
	Etherbase(ctx context.Context, in *EtherbaseRequest, opts ...grpc.CallOption) (*EtherbaseReply, error)
	NetVersion(ctx context.Context, in *NetVersionRequest, opts ...grpc.CallOption) (*NetVersionReply, error)
}

type eTHBACKENDClient struct {
	cc grpc.ClientConnInterface
}

func NewETHBACKENDClient(cc grpc.ClientConnInterface) ETHBACKENDClient {
	return &eTHBACKENDClient{cc}
}

var eTHBACKENDAddStreamDesc = &grpc.StreamDesc{
	StreamName: "Add",
}

func (c *eTHBACKENDClient) Add(ctx context.Context, in *TxRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var eTHBACKENDEtherbaseStreamDesc = &grpc.StreamDesc{
	StreamName: "Etherbase",
}

func (c *eTHBACKENDClient) Etherbase(ctx context.Context, in *EtherbaseRequest, opts ...grpc.CallOption) (*EtherbaseReply, error) {
	out := new(EtherbaseReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/Etherbase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var eTHBACKENDNetVersionStreamDesc = &grpc.StreamDesc{
	StreamName: "NetVersion",
}

func (c *eTHBACKENDClient) NetVersion(ctx context.Context, in *NetVersionRequest, opts ...grpc.CallOption) (*NetVersionReply, error) {
	out := new(NetVersionReply)
	err := c.cc.Invoke(ctx, "/remote.ETHBACKEND/NetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ETHBACKENDService is the service API for ETHBACKEND service.
// Fields should be assigned to their respective handler implementations only before
// RegisterETHBACKENDService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type ETHBACKENDService struct {
	Add        func(context.Context, *TxRequest) (*AddReply, error)
	Etherbase  func(context.Context, *EtherbaseRequest) (*EtherbaseReply, error)
	NetVersion func(context.Context, *NetVersionRequest) (*NetVersionReply, error)
}

func (s *ETHBACKENDService) add(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.Add == nil {
		return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
	}
	in := new(TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/remote.ETHBACKEND/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Add(ctx, req.(*TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ETHBACKENDService) etherbase(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.Etherbase == nil {
		return nil, status.Errorf(codes.Unimplemented, "method Etherbase not implemented")
	}
	in := new(EtherbaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Etherbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/remote.ETHBACKEND/Etherbase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Etherbase(ctx, req.(*EtherbaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *ETHBACKENDService) netVersion(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.NetVersion == nil {
		return nil, status.Errorf(codes.Unimplemented, "method NetVersion not implemented")
	}
	in := new(NetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.NetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/remote.ETHBACKEND/NetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NetVersion(ctx, req.(*NetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterETHBACKENDService registers a service implementation with a gRPC server.
func RegisterETHBACKENDService(s grpc.ServiceRegistrar, srv *ETHBACKENDService) {
	sd := grpc.ServiceDesc{
		ServiceName: "remote.ETHBACKEND",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Add",
				Handler:    srv.add,
			},
			{
				MethodName: "Etherbase",
				Handler:    srv.etherbase,
			},
			{
				MethodName: "NetVersion",
				Handler:    srv.netVersion,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "remote/ethbackend.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewETHBACKENDService creates a new ETHBACKENDService containing the
// implemented methods of the ETHBACKEND service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewETHBACKENDService(s interface{}) *ETHBACKENDService {
	ns := &ETHBACKENDService{}
	if h, ok := s.(interface {
		Add(context.Context, *TxRequest) (*AddReply, error)
	}); ok {
		ns.Add = h.Add
	}
	if h, ok := s.(interface {
		Etherbase(context.Context, *EtherbaseRequest) (*EtherbaseReply, error)
	}); ok {
		ns.Etherbase = h.Etherbase
	}
	if h, ok := s.(interface {
		NetVersion(context.Context, *NetVersionRequest) (*NetVersionReply, error)
	}); ok {
		ns.NetVersion = h.NetVersion
	}
	return ns
}

// UnstableETHBACKENDService is the service API for ETHBACKEND service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableETHBACKENDService interface {
	Add(context.Context, *TxRequest) (*AddReply, error)
	Etherbase(context.Context, *EtherbaseRequest) (*EtherbaseReply, error)
	NetVersion(context.Context, *NetVersionRequest) (*NetVersionReply, error)
}