# tfblueprintgen
This is a cli utility developed using charmbracelet CLI assets,  which generates the Modular file structure with the code for your Terraform code to speed up the development.

In this you can 
- select the environments that would be present in you organization and
- select the resources for which you want the modules to be generated.

Once selected, the tfblueprintgen will generate all the required files and modules with necessary file and code in it. See the tool in action below..!


<video controls src="assets/Tfblueprintgen_charmCLI-ezgif.com-video-to-mp4-converter.mp4" title="tfblueprintgen usage"></video>

## Installation

- From source:
Run the following setps to build your own binary

```shell
git clone https://github.com/krishnaduttPanchagnula/tfblueprintgen.git
cd tfblueprintgen
go build -o tfblueprintgen main.go
```
- Downloading Binary in Linux

```bash
wget https://github.com/krishnaduttPanchagnula/tfblueprintgen/releases/download/0.2/tfblueprintgen

chmod 777 tfblueprintgen

```