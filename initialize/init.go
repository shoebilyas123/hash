package initialize

import (
	"fmt"
	"os"

	"github.com/shoebilyas123/hash/common"
)

func Init(argvs []string) {
	pwd, err := os.Getwd()

	if err != nil {
		panic("Error: Cannot get the present working directory.")
	}
	target_dir := pwd

	// Check if the argvs is empty or .
	if len(argvs) > 0 && argvs[0] != "." {
		target_dir = pwd + "/" + argvs[0]
	}

	// If a name is provided create that named directory

	if len(argvs) > 0 && argvs[0] != "." {
		if !common.CheckDirExistence(target_dir) {
			common.HandleCreateDir(target_dir, 0755)
		} else {
			// TODO: Will ask if you want to init that directory?
			// Check for directory name clashes
			fmt.Printf("Navigate into the directory and hash init your project")
		}
	}

	// If the target_dir already has a .hash folder throw appropriate error
	if common.CheckDirExistence(target_dir + "/.hash") {
		panic("Error: Cannot overwrite an already initialized <Repository>")
	}

	// Create a .hash directory followed by a child objects folder
	target_dir += "/.hash"
	common.HandleCreateDir(target_dir, 0755)

	common.HandleCreateDir(target_dir+"/objects", 0755)
	common.HandleCreateFile(target_dir+"/", "HEAD")

	common.HandleCreateDir(target_dir+"/refs", 0755)
	common.HandleCreateDir(target_dir+"/refs/heads", 0755)

	head_ref := "ref: refs/heads/main"

	err = os.WriteFile(target_dir+"/HEAD", []byte(head_ref), 0644)

	if err != nil {
		fmt.Println(err)
	}
}
