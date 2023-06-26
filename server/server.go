package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/kudoabhijeet/file-transfer-service"
	"google.golang.org/grpc"
)

type fileTransferServer struct {
	pb.UnimplementedFileTransferServiceServer
}

func (s *fileTransferServer) UploadFile(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	file, err := os.Create(req.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(req.Content)
	if err != nil {
		return nil, err
	}

	return &pb.FileUploadResponse{
		Message: "File uploaded successfully",
	}, nil
}

func (s *fileTransferServer) DownloadFile(req *pb.FileDownloadRequest, stream pb.FileTransferService_DownloadFileServer) error {
	file, err := os.Open(req.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = stream.Send(&pb.FileDownloadResponse{
			Content: buffer[:n],
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterFileTransferServiceServer(server, &fileTransferServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server started, listening on port 50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
