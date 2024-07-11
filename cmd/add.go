package cmd

import (
	"os"
	"strings"

	"github.com/ZeinMansor/go-symfony-cli-graphql/util"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add [num1] [num2]",
	Short: "Gets the sum of 2  commands",
	Long:  `Gets the sum of 2  commands`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Arguments ", strings.Join(args, ", "))
		// println(2 + 2)
		writeToFile(1, 2)
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
}

func writeToFile(x int, y int) {
	// check if file exists

	d1 := []byte("hello\n go \t")
	err := os.WriteFile("./files_output_tst/file1.txt", d1, 0644)
	if err != nil {
		println("Error ")
		util.Check(err)
		// os.Exit(1)
	}
}

type YamlWritable interface {
	WriteToYaml() (string, error)
	// GenerateBoilerplate() (string, error)
	// GenerateSigniture() (string, error)
	// CheckIfExists() (string, error)
	// Overwrite() (string, error)
}

type Query struct {
}

func (q Query) WriteToYaml() (string, error) {
	return "", nil
}

type Mutation struct {
}

func (m Mutation) WriteToYaml() (string, error) {
	return "", nil
}

type ObjType struct {
}

func (ot ObjType) WriteToYaml() (string, error) {
	return "", nil
}

func write(d YamlWritable) string {
	return ""
}

func main() {
	var q Query = Query{}
	var m Mutation = Mutation{}
	var ot ObjType = ObjType{}

	write(q)
	write(m)
	write(ot)

}
