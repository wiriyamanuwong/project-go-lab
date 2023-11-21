package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brocaar/chirpstack-api/go/as/external/api"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	//"github.com/wiriyamanuwong/project-go-lab/"
)

const (
	serverAddress = "192.168.3.169:8080" // แก้ไขตามที่ ChirpStack ของคุณใช้
	//serverAddress = "https://chirpstack.natoncloud.com"go version
)

func main() {
	// ทำการเชื่อมต่อ gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	//conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{...})))
	if err != nil {
		log.Fatalf("ไม่สามารถเชื่อมต่อกับ server: %v", err)
	}
	defer conn.Close()

	// สร้าง ChirpStack API client
	client := api.NewInternalServiceClient(conn)

	// เรียกใช้งานฟังก์ชัน login
	username := "admin"
	password := "admin"
	response, err := login(client, username, password)
	if err != nil {
		log.Fatalf("เกิดข้อผิดพลาดในระหว่างการ login: %v", err)
	}

	// แสดงผลลัพธ์
	fmt.Printf("Token: %s\n", response.GetJwt())
	fmt.Printf("Token: %s\n", response.GetJwt())
}

// ฟังก์ชัน login
func login(client api.InternalServiceClient, username, password string) (*api.LoginResponse, error) {
	// สร้างข้อมูลการ login
	loginRequest := &api.LoginRequest{
		Username: username,
		Password: password,
	}

	// เรียกใช้งาน gRPC ฟังก์ชัน login
	response, err := client.Login(context.Background(), loginRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}
