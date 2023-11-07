package modifyhostsfile

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/spf13/cobra"

	"github.com/bishopfox/sliver/client/console"
	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
)

func ModifyHostsFileCmd(cmd *cobra.Command, con *console.SliverConsoleClient, args []string) (err error) {
	session, beacon := con.ActiveTarget.GetInteractive()
	if session == nil && beacon == nil {
		return
	}

	if len(args) != 1 {
		con.PrintErrorf("Please specify a domain and an ip.\n")
		return
	}

	param1 := args[0]
	param2, _ := cmd.Flags().GetString("domain")
	param3, _ := cmd.Flags().GetString("ipaddress")

	out, err := con.Rpc.ModifyHostsFile(context.Background(), &sliverpb.ModifyHostsFileReq{
		Request:    con.ActiveTarget.Request(cmd),
		Param1:     param1,
		DomainName: param2,
		Ipaddress:  param3,
	})
	if err != nil {
		con.PrintErrorf("%s\n", err)
		return
	}

	if out.Response != nil && out.Response.Async {
		con.AddBeaconCallback(out.Response.TaskID, func(task *clientpb.BeaconTask) {
			err = proto.Unmarshal(task.Response, out)
			if err != nil {
				con.PrintErrorf("Failed to decode response %s\n", err)
				return
			}

			PrintHostsFile(out, con)
		})
		con.PrintAsyncResponse(out.Response)
	} else {
		PrintHostsFile(out, con)
	}

	return
}

func PrintHostsFile(hw *sliverpb.ModifyHostsFile, con *console.SliverConsoleClient) {
	if hw.Response != nil && hw.Response.Err != "" {
		con.PrintErrorf("%s\n", hw.Response.Err)
		return
	}
	con.PrintInfof("Hosts file contents: %s\n", hw.Output)
}
