package e2e

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/agentbakere2e/cluster"
	"github.com/Azure/agentbakere2e/scenario"
)

func validateNodeHealth(ctx context.Context, kube *cluster.Kubeclient, vmssName string) (string, error) {
	nodeName, err := waitUntilNodeReady(ctx, kube, vmssName)
	if err != nil {
		return "", fmt.Errorf("error waiting for vmss %s ready: %w", vmssName, err)
	}

	nginxPodName, err := ensureTestNginxPod(ctx, kube, nodeName)
	if err != nil {
		return "", fmt.Errorf("error waiting for vmss %s pod ready: %w", vmssName, err)
	}

	err = waitUntilPodDeleted(ctx, kube, nginxPodName)
	if err != nil {
		return "", fmt.Errorf("error waiting vmss %s pod deleted: %w", vmssName, err)
	}

	return nodeName, nil
}

func validateWasm(ctx context.Context, kube *cluster.Kubeclient, nodeName, privateKey string) error {
	spinPodName, err := ensureWasmPods(ctx, kube, nodeName)
	if err != nil {
		return fmt.Errorf("failed to valiate wasm, unable to ensure wasm pods on node %q: %w", nodeName, err)
	}

	spinPodIP, err := getPodIP(ctx, kube, defaultNamespace, spinPodName)
	if err != nil {
		return fmt.Errorf("on node %s unable to get IP of wasm spin pod %q: %w", nodeName, spinPodName, err)
	}

	debugPodName, err := getDebugPodName(kube)
	if err != nil {
		return fmt.Errorf("on node %s unable to get debug pod name to validate wasm: %w", nodeName, err)
	}

	execResult, err := pollExecOnPod(ctx, kube, defaultNamespace, debugPodName, getWasmCurlCommand(fmt.Sprintf("http://%s/hello", spinPodIP)))
	if err != nil {
		return fmt.Errorf("on node %sunable to execute wasm validation command: %w", nodeName, err)
	}

	if execResult.exitCode != "0" {
		// retry getting the pod IP + curling the hello endpoint if the original curl reports connection refused or a timeout
		// since the wasm spin pod usually restarts at least once after initial creation, giving it a new IP
		if execResult.exitCode == "7" || execResult.exitCode == "28" {
			spinPodIP, err = getPodIP(ctx, kube, defaultNamespace, spinPodName)
			if err != nil {
				return fmt.Errorf(" on node %s unable to get IP of wasm spin pod %q: %w", nodeName, spinPodName, err)
			}

			execResult, err = pollExecOnPod(ctx, kube, defaultNamespace, debugPodName, getWasmCurlCommand(fmt.Sprintf("http://%s/hello", spinPodIP)))
			if err != nil {
				return fmt.Errorf("unable to execute on node %s wasm validation command on wasm pod %q at %s: %w", nodeName, spinPodName, spinPodIP, err)
			}

			if execResult.exitCode != "0" {
				execResult.dumpAll()
				return fmt.Errorf("curl  on node %swasm endpoint on pod %q at %s terminated with exit code %s", nodeName, spinPodName, spinPodIP, execResult.exitCode)
			}
		} else {
			execResult.dumpAll()
			return fmt.Errorf("curl  on node %swasm endpoint on pod %q at %s terminated with exit code %s", nodeName, spinPodName, spinPodIP, execResult.exitCode)
		}
	}

	if err := waitUntilPodDeleted(ctx, kube, spinPodName); err != nil {
		return fmt.Errorf("error waiting for wasm pod deletion on %s: %w", nodeName, err)
	}

	return nil
}

func runLiveVMValidators(ctx context.Context, vmssName, privateIP, sshPrivateKey string, opts *scenarioRunOpts) error {
	podName, err := getDebugPodName(opts.clusterConfig.Kube)
	if err != nil {
		return fmt.Errorf("While running live validator for node %s, unable to get debug pod name: %w", vmssName, err)
	}

	validators := commonLiveVMValidators()
	if opts.scenario.LiveVMValidators != nil {
		validators = append(validators, opts.scenario.LiveVMValidators...)
	}

	for _, validator := range validators {
		desc := validator.Description
		command := validator.Command
		isShellBuiltIn := validator.IsShellBuiltIn
		log.Printf("running live VM validator on %s: %q", vmssName, desc)

		execResult, err := pollExecOnVM(ctx, opts.clusterConfig.Kube, privateIP, podName, sshPrivateKey, command, isShellBuiltIn)
		if err != nil {
			return fmt.Errorf("unable to execute validator on node %s command %q: %w", vmssName, command, err)
		}

		if validator.Asserter != nil {
			err := validator.Asserter(execResult.exitCode, execResult.stdout.String(), execResult.stderr.String())
			if err != nil {
				execResult.dumpAll()
				return fmt.Errorf("failed validator on node %s assertion: %w", vmssName, err)
			}
		}
	}

	return nil
}

func commonLiveVMValidators() []*scenario.LiveVMValidator {
	return []*scenario.LiveVMValidator{
		{
			Description: "assert /etc/default/kubelet should not contain dynamic config dir flag",
			Command:     "cat /etc/default/kubelet",
			Asserter: func(code, stdout, stderr string) error {
				if code != "0" {
					return fmt.Errorf("validator command terminated with exit code %q but expected code 0", code)
				}
				if strings.Contains(stdout, "--dynamic-config-dir") {
					return fmt.Errorf("/etc/default/kubelet should not contain kubelet flag '--dynamic-config-dir', but does")
				}
				return nil
			},
		},
		scenario.SysctlConfigValidator(
			map[string]string{
				"net.ipv4.tcp_retries2":             "8",
				"net.core.message_burst":            "80",
				"net.core.message_cost":             "40",
				"net.core.somaxconn":                "16384",
				"net.ipv4.tcp_max_syn_backlog":      "16384",
				"net.ipv4.neigh.default.gc_thresh1": "4096",
				"net.ipv4.neigh.default.gc_thresh2": "8192",
				"net.ipv4.neigh.default.gc_thresh3": "16384",
			},
		),
		scenario.DirectoryValidator(
			"/var/log/azure/aks",
			[]string{
				"cluster-provision.log",
				"cluster-provision-cse-output.log",
				"cloud-init-files.paved",
				"vhd-install.complete",
				"cloud-config.txt",
			},
		),
	}
}
