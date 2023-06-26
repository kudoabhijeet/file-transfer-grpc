package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	pb "github.com/kudoabhijeet/file-transfer-service"
	"google.golang.org/grpc"
)

func uploadFile(client pb.FileTransferServiceClient, filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	res, err := client.UploadFile(context.Background(), &pb.FileUploadRequest{
		Filename: filename,
		Content:  content,
	})
	if err != nil {
		log.Fatalf("upload failed: %v", err)
	}
	log.Printf("Response: %s", res.Message)
}

func downloadFile(client pb.FileTransferServiceClient, filename string) {
	stream, err := client.DownloadFile(context.Background(), &pb.FileDownloadRequest{
		Filename: filename,
	})
	if err != nil {
		log.Fatalf("download failed: %v", err)
	}

	var content []byte
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive file content: %v", err)
		}
		content = append(content, res.Content...)
		if err == io.EOF {
			break
		}
	}

	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
	log.Println("File downloaded successfully")
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFileTransferServiceClient(conn)

	uploadFile(client, "file.txt")
	downloadFile(client, "file.txt")
}
