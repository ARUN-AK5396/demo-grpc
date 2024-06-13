package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ARUN-AK5396/demo-grpc/invoicer"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

// Middleware function to handle CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Credentials") // Include Access-Control-Allow-Credentials in allowed headers

		// Handle preflight request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	data := requestData["data"]

	fmt.Printf("Received data: %s\n", data)

	fmt.Fprintf(w, "Received data: %s", data)
}

func HomeFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Arun")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dataset := map[string]interface{}{
		"id":   1,
		"name": "Arun",
		"age":  30,
		"skills": []string{
			"Go",
			"Python",
			"JavaScript",
		},
	}

	jsonData, err := json.Marshal(dataset)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test PDF content"),
		Docx: []byte("test DOCX content"),
	}, nil
}

func main() {
	// Create the main listener.
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a cmux.
	m := cmux.New(lis)

	// Match gRPC requests.
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))

	// Match HTTP requests.
	httpL := m.Match(cmux.HTTP1Fast())

	// Create the gRPC server.
	grpcServer := grpc.NewServer()
	invoicer.RegisterInvoicerServer(grpcServer, &myInvoicerServer{})

	// Create the HTTP server.
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", HomeFunction)
	httpMux.HandleFunc("/data", handlePost)
	httpMux.HandleFunc("/get-data", handleGet)

	// Wrap the HTTP mux with the CORS middleware
	httpServer := &http.Server{
		Handler: enableCors(httpMux),
	}

	// Serve gRPC server.
	go func() {
		log.Println("Serving gRPC on :8080")
		if err := grpcServer.Serve(grpcL); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Serve HTTP server.
	go func() {
		//log.Println("Serving HTTP on :8080")
		if err := httpServer.Serve(httpL); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	// Start cmux serving.
	if err := m.Serve(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
