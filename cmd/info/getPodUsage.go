/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func getPodUsage() {
	fmt.Println("getPodUsage called")

	// Get the config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/phanimullapudi/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// Get the metrics client
	metricsClientset, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "default" // leave empty to get data from all namespaces
	podMetricsList, err := metricsClientset.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, v := range podMetricsList.Items {
		fmt.Printf("%s\n", v.GetName())
		fmt.Printf("%s\n", v.GetNamespace())
		fmt.Printf("%vm\n", v.Containers[0].Usage.Cpu().MilliValue())
		fmt.Printf("%vMi\n", v.Containers[0].Usage.Memory().Value()/(1024*1024))
	}

}

// getPodUsageCmd represents the getPodUsage command
var getPodUsageCmd = &cobra.Command{
	Use:   "getPodUsage",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getPodUsage()
	},
}

func init() {
	InfoCmd.AddCommand(getPodUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPodUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPodUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
