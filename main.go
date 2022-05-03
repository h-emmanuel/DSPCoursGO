package main

import (
	//"fmt"
	//"github.com/mrlaowilly/devgo/cours"
	"fmt"
	"github.com/h-emmanuel/dspcoursGO/internal/fssync"
	"github.com/spf13/cobra"
)

func main() {
	//fmt.Println("Hello World", cours.LeNomDeLaVariablePublic, cours.LeNomDeLaVariablePublicAvecValue)
	//fmt.Println(cours.Hello("Willy"))
	//cours.Master()
	//u := cours.MakeUser("Willy")
	//u.Hello()
	var root = &cobra.Command{
		Use:   "app [source] [target]",
		Short: "Permet de copier des fichiers",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			source := args[0]
			target := args[1]

			if fssync.FolderExist(source) {
				fssync.CreatFolderIfNotExist(target)
				for _, file := range fssync.Scan(source, source) {
					fssync.Copy(source+file, target+file)
				}
			} else {
				fmt.Println("Source folder not exist")
			}

			fmt.Println(source, target)
		},
		DisableFlagsInUseLine: true,
	}

	root.Execute()
}
