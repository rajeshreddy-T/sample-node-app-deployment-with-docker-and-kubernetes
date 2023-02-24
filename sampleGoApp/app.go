package main

import (
	"context"
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Load the Kubernetes configuration from file
	// home := os.Getenv("HOME")
	kubeconfig := os.Getenv("KUBECONFIG")

	// Use the current context in kubeconfig

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Define the Deployment object
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "example-app",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "example-app",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "example-container",
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	// Trigger the deployment
	result, err := clientset.AppsV1().Deployments("default").Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func int32Ptr(i int32) *int32 { return &i }

// func main() {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		// handle error
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		// handle error
// 	}

// 	deployment := &appsv1.Deployment{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "my-deployment",
// 		},
// 		Spec: appsv1.DeploymentSpec{
// 			Replicas: int32Ptr(3),
// 			Selector: &metav1.LabelSelector{
// 				MatchLabels: map[string]string{
// 					"app": "my-app",
// 				},
// 			},
// 			Template: corev1.PodTemplateSpec{
// 				ObjectMeta: metav1.ObjectMeta{
// 					Labels: map[string]string{
// 						"app": "my-app",
// 					},
// 				},
// 				Spec: corev1.PodSpec{
// 					Containers: []corev1.Container{
// 						{
// 							Name:  "my-container",
// 							Image: "my-image",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	result, err := clientset.AppsV1().Deployments("my-namespace").Create(context.Background(), deployment, metav1.CreateOptions{})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// type RpcServer struct{}

// func int32Ptr(i int32) *int32 { return &i }

// func (s *RpcServer) CreateDeployment(args *DeploymentRequest, reply *DeploymentReply) error {
// 	deployment := &appsv1.Deployment{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: args.Name,
// 		},
// 		Spec: appsv1.DeploymentSpec{
// 			Replicas: int32Ptr(args.Replicas),
// 			Selector: &metav1.LabelSelector{
// 				MatchLabels: map[string]string{
// 					"app": args.AppLabel,
// 				},
// 			},
// 			Template: corev1.PodTemplateSpec{
// 				ObjectMeta: metav1.ObjectMeta{
// 					Labels: map[string]string{
// 						"app": args.AppLabel,
// 					},
// 				},
// 				Spec: corev1.PodSpec{
// 					Containers: []corev1.Container{
// 						{
// 							Name:  args.ContainerName,
// 							Image: args.ImageName,
// 							Ports: []corev1.ContainerPort{
// 								{
// 									ContainerPort: args.Port,
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	result, err := clientset.AppsV1().Deployments(args.Namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
// 	if err != nil {
// 		return err
// 	}
// 	reply.Message = fmt.Sprintf("Deployment created: %s", result.ObjectMeta.Name)
// 	return nil
// }

// func (s *RpcServer) DeleteDeployment(args *DeploymentRequest, reply *DeploymentReply) error {
// 	err := clientset.AppsV1().Deployments(args.Namespace).Delete(context.Background(), args.Name, metav1.DeleteOptions{})
// 	if err != nil {
// 		return err
// 	}
// 	reply.Message = fmt.Sprintf("Deployment deleted: %s", args.Name)
// 	return nil
// }

// func (s *RpcServer) GetDeployment(args *DeploymentRequest, reply *DeploymentReply) error {
// 	result, err := clientset.AppsV1().Deployments(args.Namespace).Get(context.Background(), args.Name, metav1.GetOptions{})
// 	if err != nil {
// 		return err
// 	}
// 	reply.Message = fmt.Sprintf("Deployment found: %s", result.ObjectMeta.Name)
// 	return nil
// }

// func (s *RpcServer) ListDeployments(args *DeploymentRequest, reply *DeploymentReply) error {
// 	result, err := clientset.AppsV1().Deployments(args.Namespace).List(context.Background(), metav1.ListOptions{})
// 	if err != nil {
// 		return err
// 	}
// 	reply.Message = fmt.Sprintf("Deployments found: %d", len(result.Items))
// 	return nil
// }

// func main() {
// 	rpc.Register(new(RpcServer))
// 	rpc.HandleHTTP()
// 	l, e := net.Listen("tcp", ":1234")
// 	if e != nil {
// 		log.Fatal("listen error:", e)
// 	}
// 	http.Serve(l, nil)
// }

// func main() {
// 	client, err := rpc.DialHTTP("tcp", "localhost:1234")
// 	if err != nil {
// 		log.Fatal("dialing:", err)
// 	}

// 	args := &DeploymentRequest{
// 		Name:          "test-deployment",
// 		Namespace:     "default",
// 		Replicas:      1,
// 		AppLabel:      "test-app",
// 		ContainerName: "test-container",
// 		ImageName:     "nginx",
// 		Port:          80,
// 	}
// 	var reply DeploymentReply
// 	err = client.Call("RpcServer.CreateDeployment", args, &reply)
// 	if err != nil {
// 		log.Fatal("arith error:", err)
// 	}
// 	fmt.Println(reply.Message)
// }
