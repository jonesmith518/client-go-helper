package kubectl

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestLogs(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := corev1.PodLogOptions{}
	str,e := pod.Logs(&opts)
	if e !=nil {
		log.Print(e)
	}else {
		log.Print(str)
	}

}

func TestDescribe(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	str,e :=pod.Describe()
	if e !=nil {
		log.Print(e)
	}
	log.Print(str)
}

func TestExec(t *testing.T) {

	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	e := pod.Exec([]string{"ls","-al"})
	if e !=nil {
		log.Print(e)
	}
}

func TestCp(t *testing.T) {
	// kubectl cp /tmp/localfile  api-test-775cf487ff-7zhnj:/opt
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	e := pod.Cp("/tmp/localfile","/opt/")
		if e !=nil {
		log.Print(e)
	}
}


func TestGetAll(t *testing.T) {
	pod := Pod{
		ContainerName: "",
		Name:       "api-test-775cf487ff-7zhnj",
		Namespace:     "dev",
	}
	opts := v1.ListOptions{}
	items ,e := pod.GetAll(&opts)
	if e !=nil {
		log.Print(e)
	}
	for _, v := range items {
		json, _ := json.Marshal(v)
		rawYaml ,_ := yaml.JSONToYAML(json)
		log.Println(string(rawYaml))
	}
	log.Println("=======================")
	pod2 := Pod{
		ContainerName: "",
		Name:       "api-test",
		Namespace:     "dev",
	}
	items2 ,e := pod2.GetAll(&opts)
	if e !=nil {
		log.Print(e)
	}
	for _, v := range items2 {
		json, _ := json.Marshal(v)
		rawYaml ,_ := yaml.JSONToYAML(json)
		log.Println(string(rawYaml))
	}
}