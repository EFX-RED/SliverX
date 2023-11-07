package rpc

import (
	"context"

	"github.com/bishopfox/sliver/protobuf/commonpb"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
)

// Modify the victims host file - modifyhostsfile command
func (rpc *Server) ModifyHostsFile(ctx context.Context, req *sliverpb.ModifyHostsFileReq) (*sliverpb.ModifyHostsFile, error) {
	resp := &sliverpb.ModifyHostsFile{Response: &commonpb.Response{}}
	err := rpc.GenericHandler(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
