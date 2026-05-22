package main

import (
	"fmt"
	"lab09/program"
)

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func main() {
	server := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.0.255", "169.0.0.1"},
	}

	resultYaml, err := program.ToYAML(server)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resultJson, err := program.ToJSON(server)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Olexandr Osiychuk")
	fmt.Println("IPZs-21")
	fmt.Println("Task 1(Variant 9)\n\n")

	fmt.Println("---------------------ToYAML---------------------")
	fmt.Println(resultYaml)
	fmt.Println("-----------------ToJSON(lab06)------------------")
	fmt.Println(resultJson)
}
