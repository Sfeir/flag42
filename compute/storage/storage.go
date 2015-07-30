package storage

import (
	"os/exec"
)

func ComputeToStorage(dir, store string){
	cmd := exec.Command("gsutil", "cp", "-R", dir, store)
	cmd.Run()
	cmd = exec.Command("gsutil","-m","acl","set","-R","-a","public-read", store)
	cmd.Run()
)

func StorageToCompute(dir, store string){
	cmd := exec.Command("gsutil", "cp", "-R", store, dir)
	cmd.Run()
)
