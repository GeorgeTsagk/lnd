// Code generated by falafel 0.9.1. DO NOT EDIT.
// source: router.proto

package routerrpc

import (
	"context"

	gateway "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func RegisterRouterJSONCallbacks(registry map[string]func(ctx context.Context,
	conn *grpc.ClientConn, reqJSON string, callback func(string, error))) {

	marshaler := &gateway.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
	}

	registry["routerrpc.Router.SendPaymentV2"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SendPaymentRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.SendPaymentV2(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.TrackPaymentV2"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &TrackPaymentRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.TrackPaymentV2(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.TrackPayments"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &TrackPaymentsRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.TrackPayments(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.EstimateRouteFee"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &RouteFeeRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.EstimateRouteFee(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.SendToRoute"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SendToRouteRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.SendToRoute(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.SendToRouteV2"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SendToRouteRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.SendToRouteV2(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.ResetMissionControl"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &ResetMissionControlRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.ResetMissionControl(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.QueryMissionControl"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &QueryMissionControlRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.QueryMissionControl(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.XImportMissionControl"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &XImportMissionControlRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.XImportMissionControl(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.GetMissionControlConfig"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &GetMissionControlConfigRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.GetMissionControlConfig(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.SetMissionControlConfig"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SetMissionControlConfigRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.SetMissionControlConfig(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.QueryProbability"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &QueryProbabilityRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.QueryProbability(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.BuildRoute"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &BuildRouteRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.BuildRoute(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.SubscribeHtlcEvents"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SubscribeHtlcEventsRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.SubscribeHtlcEvents(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.SendPayment"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &SendPaymentRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.SendPayment(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.TrackPayment"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &TrackPaymentRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		stream, err := client.TrackPayment(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		go func() {
			for {
				select {
				case <-stream.Context().Done():
					callback("", stream.Context().Err())
					return
				default:
				}

				resp, err := stream.Recv()
				if err != nil {
					callback("", err)
					return
				}

				respBytes, err := marshaler.Marshal(resp)
				if err != nil {
					callback("", err)
					return
				}
				callback(string(respBytes), nil)
			}
		}()
	}

	registry["routerrpc.Router.UpdateChanStatus"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &UpdateChanStatusRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.UpdateChanStatus(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.XAddLocalChanAliases"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &AddAliasesRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.XAddLocalChanAliases(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}

	registry["routerrpc.Router.XDeleteLocalChanAliases"] = func(ctx context.Context,
		conn *grpc.ClientConn, reqJSON string, callback func(string, error)) {

		req := &DeleteAliasesRequest{}
		err := marshaler.Unmarshal([]byte(reqJSON), req)
		if err != nil {
			callback("", err)
			return
		}

		client := NewRouterClient(conn)
		resp, err := client.XDeleteLocalChanAliases(ctx, req)
		if err != nil {
			callback("", err)
			return
		}

		respBytes, err := marshaler.Marshal(resp)
		if err != nil {
			callback("", err)
			return
		}
		callback(string(respBytes), nil)
	}
}
