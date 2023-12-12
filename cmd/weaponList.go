/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"log"
	// "strings"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// weaponListCmd represents the weaponList command
var weaponListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		weapons := tview.NewTreeNode("Weapons")
		tree := tview.NewTreeView().
			SetRoot(weapons).
			SetCurrentNode(weapons)
		missilesNode := tview.NewTreeNode("Missiles").
				SetExpanded(false).
				SetSelectable(true)
		indirectsNode := tview.NewTreeNode("Indirects").
				SetExpanded(false).
				SetSelectable(true)
		weapons.AddChild(missilesNode).AddChild(indirectsNode)

		missiles := get_missiles()
		indirects := get_indirects()
		for i := 0; i < len(missiles); i ++ {
			new_node := tview.NewTreeNode(missiles[i].Name).
					SetReference(missiles[i])
			missilesNode.AddChild(new_node.Collapse())
		}
		for i := 0; i < len(indirects); i ++ {
			new_node := tview.NewTreeNode(indirects[i].Name).
					SetReference(indirects[i])
			indirectsNode.AddChild(new_node.Collapse())
		}
		selectedView := tview.NewTextView()

		flex := tview.NewFlex().
			AddItem(tree, 0, 1, true).
			AddItem(selectedView, 0, 1, false)
		tree.SetSelectedFunc(func(node *tview.TreeNode) {
			children := node.GetChildren()
			reference := node.GetReference()
			expanded := node.IsExpanded()
			if children != nil && expanded != true {
				node.Expand()
			} else if children != nil && expanded == true {
				node.Collapse()
				selectedView.Clear()
			} else if reference != nil {
				// app.Stop()
				// fmt.Println(reference)
				text, _ := json.Marshal(reference)
				selectedView.SetText(string(text))

			}
		})

		if err := app.SetRoot(flex, true).Run(); err != nil {
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

func get_indirects () []Indirect {
	db, err := sql.Open("sqlite3", "./tow.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	response, err := db.Query("SELECT * from indirects")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Close()
	weapons := make([]Indirect, 0)

	for response.Next() {
		ourIndirect := Indirect{}
		// local_ammo := &ourIndirect.Ammo
		err = response.Scan(&ourIndirect.Id, &ourIndirect.Name, &ourIndirect.RangeMax, &ourIndirect.RangeMin, &ourIndirect.Ammo)//strings.Split(local_ammo, ""))
		if err != nil {
			log.Fatal(err)
		}

		weapons = append(weapons, ourIndirect)
	}
		return weapons
}
