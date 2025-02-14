// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ela_test

import (
	"net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/open-ness/common/proxy/progutil"
	"github.com/open-ness/edgenode/pkg/ela"
	pb "github.com/open-ness/edgenode/pkg/ela/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("gRPC InterfaceService", func() {
	expectedNetworkInterfaces := []*pb.NetworkInterface{
		{
			Id:                "0000:00:00.0",
			Driver:            pb.NetworkInterface_USERSPACE,
			Type:              pb.NetworkInterface_UPSTREAM,
			MacAddress:        "AA:BB:CC:DD:EE:FF",
			FallbackInterface: "0000:00:00.1",
		},
		{
			Id:                "0000:00:00.1",
			Driver:            pb.NetworkInterface_KERNEL,
			Type:              pb.NetworkInterface_NONE,
			MacAddress:        "00:11:22:33:44:55",
			FallbackInterface: "0000:00:01.2",
		},
	}

	BeforeEach(func() {
		ela.GetInterfaces = func() (*pb.NetworkInterfaces, error) {
			return &pb.NetworkInterfaces{
					NetworkInterfaces: expectedNetworkInterfaces},
				nil
		}
	})

	Describe("GetAll method", func() {
		Specify("will respond", func() {

			lis, err := net.Listen("tcp", ela.Config.ControllerEndpoint)
			prefaceLis := progutil.NewPrefaceListener(lis)
			defer prefaceLis.Close()
			go prefaceLis.Accept() // we only expect 1 connection

			// Then connecting to it from this thread
			conn, err := grpc.Dial("",
				grpc.WithTransportCredentials(transportCreds), grpc.WithDialer(prefaceLis.DialEla))
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()

			client := pb.NewInterfaceServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(),
				3*time.Second)
			defer cancel()

			ifs, err := client.GetAll(ctx, &empty.Empty{},
				grpc.WaitForReady(true))
			Expect(err).ShouldNot(HaveOccurred())

			Expect(ifs).ToNot(BeNil())
			Expect(ifs.NetworkInterfaces).To(HaveLen(2))

			ni := ifs.NetworkInterfaces[0]
			Expect(ni).ToNot(BeNil())
			Expect(ni.Id).To(Equal(expectedNetworkInterfaces[0].Id))
			Expect(ni.Driver).To(Equal(expectedNetworkInterfaces[0].Driver))
			Expect(ni.Type).To(Equal(expectedNetworkInterfaces[0].Type))
			Expect(ni.MacAddress).
				To(Equal(expectedNetworkInterfaces[0].MacAddress))
			Expect(ni.FallbackInterface).
				To(Equal(expectedNetworkInterfaces[0].FallbackInterface))
		})
	})

	Describe("Get method", func() {
		get := func(id string) (*pb.NetworkInterface, error) {

			lis, err := net.Listen("tcp", ela.Config.ControllerEndpoint)
			prefaceLis := progutil.NewPrefaceListener(lis)
			defer prefaceLis.Close()
			go prefaceLis.Accept() // we only expect 1 connection

			// Then connecting to it from this thread
			conn, err := grpc.Dial("",
				grpc.WithTransportCredentials(transportCreds), grpc.WithDialer(prefaceLis.DialEla))
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()

			client := pb.NewInterfaceServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(),
				3*time.Second)
			defer cancel()

			return client.Get(ctx, &pb.InterfaceID{Id: id},
				grpc.WaitForReady(true))
		}

		Context("called with Id corresponding to existing interface", func() {
			Specify("will return that interface", func() {
				ni, err := get("0000:00:00.1")

				Expect(err).ShouldNot(HaveOccurred())

				Expect(ni).ToNot(BeNil())
				Expect(ni.Id).To(Equal(expectedNetworkInterfaces[1].Id))
				Expect(ni.Driver).To(Equal(expectedNetworkInterfaces[1].Driver))
				Expect(ni.Type).To(Equal(expectedNetworkInterfaces[1].Type))
				Expect(ni.MacAddress).
					To(Equal(expectedNetworkInterfaces[1].MacAddress))
				Expect(ni.FallbackInterface).
					To(Equal(expectedNetworkInterfaces[1].FallbackInterface))
			})
		})

		Context("called with Id pointing to not existing interface", func() {
			Specify("will return error", func() {
				ni, err := get("0000:02:02.0")

				Expect(ni).To(BeNil())
				Expect(err).To(HaveOccurred())

				st, ok := status.FromError(err)
				Expect(ok).To(BeTrue())
				Expect(st.Message()).To(Equal("interface not found"))
				Expect(st.Code()).To(Equal(codes.NotFound))
			})
		})

		Context("called with empty Id", func() {
			Specify("will return error", func() {
				ni, err := get("")

				Expect(ni).To(BeNil())
				Expect(err).To(HaveOccurred())

				st, ok := status.FromError(err)
				Expect(ok).To(BeTrue())
				Expect(st.Message()).To(Equal("empty id"))
				Expect(st.Code()).To(Equal(codes.InvalidArgument))
			})
		})

		Context("called with Id", func() {
			Context("failed to obtain interfaces", func() {
				Specify("will return error", func() {
					ela.GetInterfaces = func() (*pb.NetworkInterfaces, error) {
						return nil, errors.New("failure")
					}

					ni, err := get("0000:00:00.0")

					Expect(ni).To(BeNil())
					Expect(err).To(HaveOccurred())

					st, ok := status.FromError(err)
					Expect(ok).To(BeTrue())
					Expect(st.Message()).To(Equal("failed to obtain " +
						"network interfaces: failure"))
					Expect(st.Code()).To(Equal(codes.Unknown))
				})
			})
		})
	})

	Describe("BulkUpdate method", func() {
		bulkUpdate := func(networkInterfaces *pb.NetworkInterfaces) error {

			lis, err := net.Listen("tcp", ela.Config.ControllerEndpoint)
			prefaceLis := progutil.NewPrefaceListener(lis)
			defer prefaceLis.Close()
			go prefaceLis.Accept() // we only expect 1 connection

			// Then connecting to it from this thread
			conn, err := grpc.Dial("",
				grpc.WithTransportCredentials(transportCreds), grpc.WithDialer(prefaceLis.DialEla))
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()

			client := pb.NewInterfaceServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(),
				3*time.Second)
			defer cancel()

			_, err = client.BulkUpdate(ctx, networkInterfaces,
				grpc.WaitForReady(true))
			return err
		}
		var nis *pb.NetworkInterfaces

		BeforeEach(func() {
			nis = &pb.NetworkInterfaces{
				NetworkInterfaces: []*pb.NetworkInterface{
					{
						Id:                "0000:00:00.0",
						Driver:            pb.NetworkInterface_USERSPACE,
						Type:              pb.NetworkInterface_UPSTREAM,
						MacAddress:        "AA:BB:CC:DD:EE:FF",
						FallbackInterface: "0000:00:00.1",
					}}}
		})

		Context("stores request data and runs NTS configuration", func() {
			When("NTS configuration fails", func() {
				It("returns an error, but stores the data", func() {
					ela.NTSConfigurationHandler = func(context.Context) error {
						return errors.New("dummy error")
					}

					err := bulkUpdate(nis)

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(
						And(ContainSubstring("dummy error"),
							ContainSubstring("failed to configure NTS")))

					Expect(ela.InterfaceConfigurationData.
						NetworkInterfaces.NetworkInterfaces).
						To(HaveLen(len(nis.NetworkInterfaces)))
				})
			})

			When("NTS configuration succeeds", func() {
				It("returns no error and stores the data", func() {
					ela.NTSConfigurationHandler = func(context.Context) error {
						return nil
					}

					err := bulkUpdate(nis)

					Expect(err).ToNot(HaveOccurred())

					Expect(ela.InterfaceConfigurationData.
						NetworkInterfaces.NetworkInterfaces).
						To(HaveLen(len(nis.NetworkInterfaces)))
				})
			})
		})

		Describe("verifies request data", func() {
			Context("which contains invalid NetworkInterface", func() {
				When("id is empty", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].Id = ""
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring("Id is nil"))
					})
				})

				When("kernel driver is requested", func() {
					It("is successful", func() {
						ela.NTSConfigurationHandler = func(
							context.Context) error {
							return nil
						}
						nis.NetworkInterfaces[0].Driver =
							pb.NetworkInterface_KERNEL
						err := bulkUpdate(nis)

						Expect(err).ToNot(HaveOccurred())
					})
				})

				When("none type (direction) is requested", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].Type =
							pb.NetworkInterface_NONE
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring(
							"type 'NONE' is not supported"))
					})
				})

				When("mac address is invalid", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].MacAddress =
							"11.22.33.44.55.66.77.88"
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring(
							"invalid MAC address"))
					})
				})

				When("vlan is requested", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].Vlan = 1
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring(
							"Vlan is not supported"))
					})
				})

				When("zones are requested", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].Zones = []string{"dummy"}
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring(
							"Zones are not supported"))
					})
				})

				When("fallback interface is empty", func() {
					It("returns an error", func() {
						nis.NetworkInterfaces[0].FallbackInterface = ""
						err := bulkUpdate(nis)

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(ContainSubstring(
							"FallbackInterface is empty"))
					})
				})
			})
		})
	})
})
