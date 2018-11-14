package mistertwo

import "fmt"

// RestoreTask uses 'source:target' pairs defined in the configuration spec
// to copy the targets to the sources. It is the reverse of BackupTask.
func RestoreTask(config *Configuration) {
	fmt.Println(withColor(cyan, "bonclay: restore task\n"))

	var errors []string
	for dst, src := range config.Spec {
		srcPath, err := fullPath(src)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
			continue
		}
		dstPath, err := fullPath(dst)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
			continue
		}

		err = copy(config.Restore.Overwrite, srcPath, dstPath)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
		} else {
			printTaskResponse(true, src, dst)
		}
	}

	if len(errors) > 0 {
		printTaskErrors("restore", errors)
	} else {
		fmt.Println(withColor(green, "\nAll files/directories were successfully restored."))
	}
}
