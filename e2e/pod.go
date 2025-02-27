package e2e

import (
	"context"
	"fmt"

	"github.com/Azure/agentbakere2e/cluster"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"
)

// Returns the name of a pod that's a member of the 'debug' daemonset, running on an aks-nodepool node.
func getDebugPodName(kube *cluster.Kubeclient) (string, error) {
	podList := corev1.PodList{}
	if err := kube.Dynamic.List(context.Background(), &podList, client.MatchingLabels{"app": "debug"}); err != nil {
		return "", fmt.Errorf("failed to list debug pod: %w", err)
	}

	if len(podList.Items) < 1 {
		return "", fmt.Errorf("failed to find debug pod, list by selector returned no results")
	}

	podName := podList.Items[0].Name
	return podName, nil
}

func getPodIP(ctx context.Context, kube *cluster.Kubeclient, namespaceName, podName string) (string, error) {
	pod, err := kube.Typed.CoreV1().Pods(namespaceName).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("unable to get pod %s/%s: %w", namespaceName, podName, err)
	}
	return pod.Status.PodIP, nil
}

func ensureTestNginxPod(ctx context.Context, kube *cluster.Kubeclient, nodeName string) (string, error) {
	nginxPodName := fmt.Sprintf("%s-nginx", nodeName)
	nginxPodManifest := getNginxPodTemplate(nodeName)
	if err := ensurePod(ctx, kube, nginxPodName, nginxPodManifest); err != nil {
		return "", fmt.Errorf("failed to ensure test nginx pod %q: %w", nginxPodName, err)
	}
	return nginxPodName, nil
}

func ensureWasmPods(ctx context.Context, kube *cluster.Kubeclient, nodeName string) (string, error) {
	spinPodName := fmt.Sprintf("%s-wasm-spin", nodeName)
	spinPodManifest := getWasmSpinPodTemplate(nodeName)
	if err := ensurePod(ctx, kube, spinPodName, spinPodManifest); err != nil {
		return "", fmt.Errorf("failed to ensure wasm spin pod %q: %w", spinPodName, err)
	}
	return spinPodName, nil
}

func applyPodManifest(ctx context.Context, kube *cluster.Kubeclient, manifest string) error {
	var podObj corev1.Pod
	if err := yaml.Unmarshal([]byte(manifest), &podObj); err != nil {
		return fmt.Errorf("failed to unmarshal Pod manifest: %w", err)
	}

	desired := podObj.DeepCopy()
	_, err := controllerutil.CreateOrUpdate(ctx, kube.Dynamic, &podObj, func() error {
		podObj = *desired
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to apply Pod manifest: %w", err)
	}

	return nil
}

func ensurePod(ctx context.Context, kube *cluster.Kubeclient, podName, manifest string) error {
	if err := applyPodManifest(ctx, kube, manifest); err != nil {
		return fmt.Errorf("failed to ensure pod: %w", err)
	}
	if err := waitUntilPodRunning(ctx, kube, podName); err != nil {
		return fmt.Errorf("failed to wait for pod to be in running state: %w", err)
	}
	return nil
}
