package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// 1- retornar a pasta atual
func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter diretório: ", err)
		return ""
	}
	return dir
}

// 2 - listar arquivos e diretórios
func listFilesAndDirectories() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos: ", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

// 3 - Pegar versão do sistema operacional
func getOSVersion() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sw_vers")
	} else {
		fmt.Println("SO não suportado")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comandos: ", err)
		return
	}
	fmt.Println(string(output))
}

func getSystemInfo() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "systeminfo")
	} else {
		fmt.Println("Sistema operacinal não suporta esse comando")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comandos: ", err)
		return
	}
	fmt.Println(string(output))
}

// Agendar Desligamento
func Shutdown() {
	var cmd *exec.Cmd
	cmd = exec.Command("shutdown", "/s", "/t", "3600")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar desligamento")
		return
	}
}

func cancelShudown() {
	var cmd *exec.Cmd
	cmd = exec.Command("shutdown", "/a")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao cancelar desligamento")
		return
	}
}

func main() {
	fmt.Println("Pasta Atual: ", getCurrentDirectory())

	fmt.Println("Arquivos e Pastas: ")
	listFilesAndDirectories()

	fmt.Println("Sistema operacional: ")
	getOSVersion()

	fmt.Println("Todas as Informações do Sistema: ")
	getSystemInfo()

	Shutdown()

	cancelShudown()
}
