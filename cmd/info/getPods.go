/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

var (
	kubeconfig *string
)

//func formatMemoryUsage(usageStr string) string {
//	return usageStr
// You may want to implement logic to convert the usageStr to a human-readable format
//}

func getPods() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("PodName - %s, PodKind -  %s, Namespace - %s \n", pod.Name, pod.Kind, pod.Namespace)
	}

}

// getPodsCmd represents the getPods command
var getPodsCmd = &cobra.Command{
	Use:   "getPods",
	Short: "get pods in the cluster",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPods called")
		getPods()
	},
}

func init() {
	InfoCmd.AddCommand(getPodsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPodsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPodsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
