package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	pb "github.com/example/mypackage"
)

// Server is a struct that holds a Kubernetes client
type Server struct {
	KubeClient *kubernetes.Clientset
}

// CreatePod is a gRPC method that creates a Kubernetes pod from the input YAML definition
func (s *Server) CreatePod(ctx context.Context, req *pb.PodRequest) (*pb.PodResponse, error) {
	pod := &corev1.Pod{}
	if err := yaml.Unmarshal([]byte(req.PodYaml), pod); err != nil {
		return &pb.PodResponse{Success: false}, err
	}
	_, err := s.KubeClient.CoreV1().Pods(pod.Namespace).Create(pod)
	if err != nil {
		return &pb.PodResponse{Success: false}, err
	}
	return &pb.PodResponse{Success: true}, nil
}

// HandleCreatePod is an HTTP handler function that calls the CreatePod gRPC method
func (s *Server) HandleCreatePod(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	podReq := &pb.PodRequest{PodYaml: string(body)}
	resp, err := s.CreatePod(context.Background(), podReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating pod: %v", err), http.StatusInternalServerError)
		return
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func main() {
	// create Kubernetes client using the official Go client library
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// create gRPC server
	srv := &Server{KubeClient: kubeClient}
	grpcServer := grpc.NewServer()
	pb.RegisterMyKubernetesServiceServer(grpcServer, srv)

	// create HTTP server with Gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/create-pod", srv.HandleCreatePod).Methods("POST")

	// start servers
	go func() {
		if err := grpcServer.Serve(); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", r))
}
