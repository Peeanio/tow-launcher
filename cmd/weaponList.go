/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	// "github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// weaponListCmd represents the weaponList command
var weaponListCmd = &cobra.Command{
	Use:   "weaponList",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// app := tview.NewApplication()
		// weapons := tview.NewTreeNode("weapons").
		// 		SetColor(tcell.ColorRed)
		// tree := tview.NewTreeView().
		// 	SetRoot(weapons).
		// 	SetCurrentNode(weapons)
  //
		// // A helper function which adds the files and directories of the given path
		// // to the given target node.
		// add := func(target *tview.TreeNode, path string) {
		// 	files, err := ioutil.ReadDir(path)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	for _, file := range files {
		// 		node := tview.NewTreeNode(file.Name()).
		// 			SetReference(filepath.Join(path, file.Name())).
		// 			SetSelectable(file.IsDir())
		// 		if file.IsDir() {
		// 			node.SetColor(tcell.ColorGreen)
		// 		}
		// 		target.AddChild(node)
		// 	}
		// }
  //
		// // Add the current directory to the root node.
		// add(weapons, "weapons")
  //
		// // If a directory was selected, open it.
		// tree.SetSelectedFunc(func(node *tview.TreeNode) {
		// 	reference := node.GetReference()
		// 	if reference == nil {
		// 		return // Selecting the root node does nothing.
		// 	}
		// 	children := node.GetChildren()
		// 	if len(children) == 0 {
		// 		// Load and show files in this directory.
		// 		path := reference.(string)
		// 		add(node, path)
		// 	} else {
		// 		// Collapse if visible, expand if collapsed.
		// 		node.SetExpanded(!node.IsExpanded())
		// 	}
		// })
  //
		// if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
		// 	panic(err)
		// }
		db, err := sql.Open("sqlite3", "./tow.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		response, err := db.Query("SELECT * from missiles")
		if err != nil {
			log.Fatal(err)
		}
		defer response.Close()
		weapons := make([]Missile, 0)

		for response.Next() {
			ourMissile := Missile{}
			err = response.Scan(&ourMissile.Id, &ourMissile.Name, &ourMissile.Pen, &ourMissile.RateOfFire, &ourMissile.Generation, &ourMissile.TopAttack, &ourMissile.AmmoLimited, &ourMissile.Aspect)
			if err != nil {
				log.Fatal(err)
			}

		weapons = append(weapons, ourMissile)
	}
		fmt.Println(weapons[0].Name)
	},
}

func init() {
	weaponCmd.AddCommand(weaponListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weaponListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weaponListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
