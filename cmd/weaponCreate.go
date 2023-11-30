/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	//"log"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)
// weaponCreateCmd represents the weaponCreate command
var weaponCreateCmd = &cobra.Command{
	Use:   "weaponCreate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		page_slice := []string{"Anti Tank Weapon", "Missile", "Indirect Fire"}
		pages := tview.NewPages()

		form_slice := make([]*tview.Form, 3)
		form_slice[0] = tview.NewForm().
				AddTextArea("Name", "", 40, 1, 40, nil).
				AddInputField("Penetration", "", 3, nil, nil).
				AddDropDown("Munition Type", []string{"HEAT", "High Explosive"}, 0, nil).
				AddInputField("Rate of Fire", "", 3, nil, nil).
				AddInputField("Range", "", 3, nil, nil).
				AddCheckbox("Close Range", false, nil).
				AddButton("Save", nil).
				AddButton("Back", func() {pages.SwitchToPage("Weapon Selection")}).
				AddButton("Quit", func() {app.Stop()})
		form_slice[0].SetBorder(true).SetTitle(" Anti Tank Weapon ")

		form_slice[1] = tview.NewForm().
				AddTextArea("Name", "", 40, 1, 40, nil).
				AddDropDown("Type", []string{"AT", "AA"}, 0, nil).
				AddInputField("Penetration", "", 3, nil, nil).
				AddDropDown("Generation", []string{"1", "2", "3"}, 0, nil).
				AddCheckbox("Top Attack", false, nil).
				AddCheckbox("Ammo Limited", false, nil).
				AddButton("Save", nil).
				AddButton("Back", func() {pages.SwitchToPage("Weapon Selection")}).
				AddButton("Quit", func() {app.Stop()})
		form_slice[1].SetBorder(true).SetTitle(" Missile ")

		form_slice[2] = tview.NewForm().
				AddTextArea("Name", "", 40, 1, 40, nil).
				AddInputField("Range Max", "", 3, nil, nil).
				AddInputField("Range Min", "", 3, nil, nil).
				AddCheckbox("High Explosive Munitions", false, nil).
				AddCheckbox("Smoke Munitions", false, nil).
				AddCheckbox("Chemical Munitions", false, nil).
				AddCheckbox("ICM Munitions", false, nil).
				AddCheckbox("Laser-guided Munitions", false, nil).
				AddCheckbox("GPS-guided Munitions", false, nil).
				AddCheckbox("Artillery-delivered Mine Munitions", false, nil).
				AddButton("Save", nil).
				AddButton("Back", func() {pages.SwitchToPage("Weapon Selection")}).
				AddButton("Quit", func() {app.Stop()})
		form_slice[2].SetBorder(true).SetTitle(" Indirect Fire ")


		plus_quit := []string{"Quit"}
		types := append(page_slice, plus_quit...)
		pages.AddPage(fmt.Sprintf("Weapon Selection"),
			tview.NewModal().
			SetText("Weapon Creation:\nSelect Weapon Type").
			AddButtons(types).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Quit" {
					app.Stop()
				} else {
					pages.SwitchToPage(buttonLabel)
				}
			}),
			true, true) //resize bool, visible bool
		for page := 0; page < len(page_slice); page++ {
			pages.AddPage(fmt.Sprintf(page_slice[page]), form_slice[page], true, false)
		}
		// pages.AddPage(fmt.Sprintf("AntiTankWeapon"),
		// 	form_slice[0],
		// 	true, false)
  //
  //
		// pages.AddPage(fmt.Sprintf("Missile"),
		// 	form_slice[1],
		// 	true, false)

		if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

	},
}

func init() {
	weaponCmd.AddCommand(weaponCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weaponCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weaponCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

