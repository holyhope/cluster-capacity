package plugin

import (
	"context"
	"fmt"

	"github.com/holyhope/cluster-capacity/pkg/logger"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func TotalMemory(ctx context.Context, configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to read kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list nodes: %w", err)
	}

	var total int64 = 0

	for _, node := range nodes.Items {
		total += node.Status.Capacity.Memory().MilliValue()
	}

	logger.FromContext(ctx).Info("%s", resource.NewMilliQuantity(total, resource.BinarySI))

	return nil
}

func TotalCPU(ctx context.Context, configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to read kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list nodes: %w", err)
	}

	var total int64 = 0

	for _, node := range nodes.Items {
		total += node.Status.Capacity.Cpu().MilliValue()
	}

	logger.FromContext(ctx).Info("%s", resource.NewMilliQuantity(total, resource.BinarySI))

	return nil
}
