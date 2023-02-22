// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespacehandler

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ory/herodot"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDeps interface {
		x.LoggerProvider
		x.WriterProvider
		config.Provider
	}
	handler struct {
		handlerDeps
	}
)

const (
	RouteBase = "/namespaces"
)

func New(d handlerDeps) *handler {
	return &handler{d}
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterNamespacesServiceServer(s, h)
}

func (h *handler) RegisterReadGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return rts.RegisterNamespacesServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (h *handler) RegisterReadGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return rts.RegisterReadServiceHandler(ctx, mux, conn)
}

func (h *handler) ListNamespaces(ctx context.Context, _ *rts.ListNamespacesRequest) (*rts.ListNamespacesResponse, error) {
	m, err := h.Config(ctx).NamespaceManager()
	if err != nil {
		h.Logger().WithError(err).Errorf("could not get namespace manager")
		return nil, herodot.ErrInternalServerError
	}
	namespaces, err := m.Namespaces(ctx)
	if err != nil {
		h.Logger().WithError(err).Errorf("could not get namespaces")
		return nil, herodot.ErrInternalServerError
	}
	apiNamespaces := make([]*rts.Namespace, len(namespaces))
	for i, n := range namespaces {
		apiNamespaces[i] = &rts.Namespace{Name: n.Name}
	}
	return &rts.ListNamespacesResponse{Namespaces: apiNamespaces}, nil
}
