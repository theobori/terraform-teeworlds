package command

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
)

const (
	// Terraform command name
	CommandName = "terraform"
)

// Terraform controller
type Terraform struct {
	// Terraform directory used with the flag -chdir
	dir string
}

// Create a new Terraform controller
func NewTerraform(dir string) *Terraform {
	return &Terraform{
		dir: dir,
	}
}

// Set the Terraform directory
func (t *Terraform) SetDir(dir string) *Terraform {
	t.dir = dir

	return t
}

// Execute the Terraform command with arguments in the `dir` directory
func (t *Terraform) Exec(args ...string) (string, error) {
	chdir := "-chdir=" + t.dir

	args = append([]string{chdir}, args...)

	cmd := exec.Command(CommandName, args...)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// Get the Terraform resources
func (t *Terraform) Resources() ([]string, error) {
	res, err := t.Exec("state", "list")
	if err != nil {
		return nil, err
	}

	resources := strings.Split(res, "\n")
	
	var ret []string

	// Remove empty element in the Slice
	for _, name := range resources {
		if name != "" {
			ret = append(ret, name)
		}
	}

	return ret, nil
}

// Destroy a Terraform resource
func (t *Terraform) Destroy(name string) error {
	_, err := t.Exec("destroy", "-auto-approve", "-target="+name)
	if err != nil {
		return err
	}

	return nil
}

// Destroy a random resource
func (t *Terraform) DestroyRandom() (string, error) {
	r, err := t.Resources()
	if err != nil {
		return "", err
	}

	size := len(r)
	if size == 0 {
		return "", fmt.Errorf("no resources availables")
	}

	index := rand.Intn(len(r))
	name := r[index]

	return name, t.Destroy(name)
}
