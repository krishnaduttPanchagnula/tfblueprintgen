package readme

import "os"

func CreateReadmeFile(filePath string) error {
	content := `# Tfblueprintgen
	This is a cli utility which generates the Modular file structure with the code for your Terraform code to speed up the development`
	return os.WriteFile(filePath, []byte(content), os.ModePerm)
}
