/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	 "fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/rivo/tview"
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
		app := tview.NewApplication()
		weapons := tview.NewTreeNode("weapons")
		tree := tview.NewTreeView().
			SetRoot(weapons).
			SetCurrentNode(weapons)
		missilesNode := tview.NewTreeNode("missiles").
				SetExpanded(false).
				SetSelectable(true)
		weapons.AddChild(missilesNode)

		missiles := get_missiles()
		for i := 0; i < len(missiles); i ++ {
			new_node := tview.NewTreeNode(missiles[i].Name).
					SetReference(missiles[i])
			missilesNode.AddChild(new_node.Collapse())
		}

		tree.SetSelectedFunc(func(node *tview.TreeNode) {
			children := node.GetChildren()
			reference := node.GetReference()
			if children != nil {
				node.Expand()
			} else if reference != nil {
				app.Stop()
				fmt.Println(reference)
			}
		})

		if err := app.SetRoot(tree, true).Run(); err != nil {
			panic(err)
		}
		// fmt.Println(missiles[0].Name)
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

func get_missiles () []Missile {
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
		return weapons
}
