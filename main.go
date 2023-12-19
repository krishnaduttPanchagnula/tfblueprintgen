package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krishnaduttPanchagnula/Tfblueprintgen/lambda"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/readme"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/s3"
	"github.com/krishnaduttPanchagnula/Tfblueprintgen/vpc"
)

var FOLDER_NAME = "terraform-aws"

var ProdFiles = []string{
	filepath.Join(FOLDER_NAME, "Environments", "Production"),
	filepath.Join(FOLDER_NAME, "Environments", "Production", "main.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "Production", "variables.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "Production", "outputs.tf"),
}

var StagingFiles = []string{
	filepath.Join(FOLDER_NAME, "Environments", "Staging"),
	filepath.Join(FOLDER_NAME, "Environments", "Staging", "main.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "Staging", "variables.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "Staging", "outputs.tf"),
}
var UATFiles = []string{
	filepath.Join(FOLDER_NAME, "Environments", "UAT"),
	filepath.Join(FOLDER_NAME, "Environments", "UAT", "main.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "UAT", "variables.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "UAT", "outputs.tf"),
}
var DevFiles = []string{
	filepath.Join(FOLDER_NAME, "Environments", "dev"),
	filepath.Join(FOLDER_NAME, "Environments", "dev", "main.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "dev", "variables.tf"),
	filepath.Join(FOLDER_NAME, "Environments", "dev", "outputs.tf"),
}

var ReadmeFile = []string{
	filepath.Join("terraform-aws", "README.md"),
}

func main() {

	// Define the directory structure

	lambdafilenames := lambda.CreateLambdaFilePathNames(FOLDER_NAME)
	s3filenames := s3.CreateS3FilePathNames(FOLDER_NAME)
	vpcfilenames := vpc.CreateVpcFilePathNames(FOLDER_NAME)

	fileNames := concatMultipleSlices([][]string{lambdafilenames, s3filenames, vpcfilenames, ProdFiles, StagingFiles, UATFiles, DevFiles})
	directoryStructure := append([]string{FOLDER_NAME}, fileNames...)
	// Create directories and files
	for i, path := range directoryStructure {
		switch {
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "lambda", "variables.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := lambda.CreateLambdaVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "lambda", "main.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := lambda.CreateLambdamoduleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "vpc", "variables.tf"):
			// Create variables.tf file with dynamic content for vpc module
			err := vpc.CreateVPCVariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "vpc", "main.tf"):
			// Create variables.tf file with dynamic content for vpc module
			err := vpc.CreateVPCModuleFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "s3", "variables.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := s3.CreateS3VariablesFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".tf" && path == filepath.Join(FOLDER_NAME, "modules", "s3", "main.tf"):
			// Create variables.tf file with dynamic content for lambda module
			err := s3.CreateS3MainFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == ".md" && path == filepath.Join(FOLDER_NAME, "README.md"):
			// Create variables.tf file with dynamic content for lambda module
			err := readme.CreateReadmeFile(path)
			if err != nil {
				fmt.Printf("Error creating %s: %v\n", path, err)
				os.Exit(1)
			}
		case filepath.Ext(path) == "":
			// Create directory
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory %s: %v\n", path, err)
				os.Exit(1)
			}
		default:
			// Create file
			_, err := os.Create(path)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", path, err)
				os.Exit(1)
			}
		}

		// Print status message for each file/directory created
		fmt.Printf("[%d/%d] Created: %s\n", i+1, len(directoryStructure), path)
	}

	fmt.Printf("File structure for %s created successfully.", FOLDER_NAME)
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

// package main

// import (
// 	"fmt"
// 	"os"

// 	tea "github.com/charmbracelet/bubbletea"
// )

// type model struct {
// 	choices  []string         // items on the to-do list
// 	cursor   int              // which to-do list item our cursor is pointing at
// 	selected map[int]struct{} // which to-do items are selected
// }

// func initialModel() model {
// 	return model{
// 		// Our to-do list is a grocery list
// 		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

// 		// A map which indicates which choices are selected. We're using
// 		// the  map like a mathematical set. The keys refer to the indexes
// 		// of the `choices` slice, above.
// 		selected: make(map[int]struct{}),
// 	}
// }
// func (m model) Init() tea.Cmd {
// 	// Just return `nil`, which means "no I/O right now, please."
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {

// 	// Is it a key press?
// 	case tea.KeyMsg:

// 		// Cool, what was the actual key pressed?
// 		switch msg.String() {

// 		// These keys should exit the program.
// 		case "ctrl+c", "q":
// 			return m, tea.Quit

// 		// The "up" and "k" keys move the cursor up
// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}

// 		// The "down" and "j" keys move the cursor down
// 		case "down", "j":
// 			if m.cursor < len(m.choices)-1 {
// 				m.cursor++
// 			}

// 		// The "enter" key and the spacebar (a literal space) toggle
// 		// the selected state for the item that the cursor is pointing at.
// 		case "enter", " ":
// 			_, ok := m.selected[m.cursor]
// 			if ok {
// 				delete(m.selected, m.cursor)
// 			} else {
// 				m.selected[m.cursor] = struct{}{}
// 			}
// 		}
// 	}

// 	// Return the updated model to the Bubble Tea runtime for processing.
// 	// Note that we're not returning a command.
// 	return m, nil
// }

// func (m model) View() string {
// 	// The header
// 	s := "What should we buy at the market?\n\n"

// 	// Iterate over our choices
// 	for i, choice := range m.choices {

// 		// Is the cursor pointing at this choice?
// 		cursor := " " // no cursor
// 		if m.cursor == i {
// 			cursor = ">" // cursor!
// 		}

// 		// Is this choice selected?
// 		checked := " " // not selected
// 		if _, ok := m.selected[i]; ok {
// 			checked = "x" // selected!
// 		}

// 		// Render the row
// 		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
// 	}

// 	// The footer
// 	s += "\nPress q to quit.\n"

// 	// Send the UI for rendering
// 	return s
// }

// func main() {
// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}
// }
