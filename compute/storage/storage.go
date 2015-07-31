package storage

import (
	"os/exec"
)

func ComputeToStorage(dir, store string) {
	print("dir :")
	print(dir)
	print("\nstore :")
	print(store)
	cmd := exec.Command("gsutil", "cp", "-R", dir, store)
	print("\n\nCaca\n")
	err := cmd.Start()
	print(err.Error())
	cmd = exec.Command("gsutil", "-m", "acl", "set", "-R", "-a", "public-read", store)
	cmd.Run()
}

func StorageToCompute(dir, store string) {
	cmd := exec.Command("gsutil", "cp", "-R", store, dir)
	cmd.Run()
}
