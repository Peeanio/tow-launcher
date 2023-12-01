/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"strconv"
	//"log"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var border bool
// weaponCreateCmd represents the weaponCreate command
var weaponCreateCmd = &cobra.Command{
	Use:   "create",
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
		if border != true{
			form_slice[0].SetBorder(true).SetTitle(" Anti Tank Weapon ")
		}

		form_slice[1] = tview.NewForm().
				AddTextArea("Name", "", 40, 1, 40, nil).
				AddInputField("Rate of Fire", "", 3, nil, nil).
				AddInputField("Penetration", "", 3, nil, nil).
				AddDropDown("Generation", []string{"1", "2", "3"}, 0, nil).
				AddCheckbox("Top Attack", false, nil).
				AddCheckbox("Ammo Limited", false, nil).
				AddDropDown("Aspect", []string{"N/A", "Rear", "All"}, 0, nil).
				AddButton("Save", func() {
					missile := make_missile(form_slice[1], app)
					str, _ := json.Marshal(missile)
					fmt.Println(string(str))
				}).
				AddButton("Back", func() {pages.SwitchToPage("Weapon Selection")}).
				AddButton("Quit", func() {app.Stop()})
		if border != true{
			form_slice[1].SetBorder(true).SetTitle(" Missile ")
		}

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
		if border != true{
			form_slice[2].SetBorder(true).SetTitle(" Indirect Fire ")
		}


		plus_quit := []string{"Quit"}
		types := append(page_slice, plus_quit...)

		selection_page := tview.NewModal().
				SetText("Weapon Creation:\nSelect Weapon Type").
				AddButtons(types).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Quit" {
						app.Stop()
					} else {
						pages.SwitchToPage(buttonLabel)
					}
				})
		if border != true{
			selection_page.SetBorder(true)
		}
		pages.AddPage(fmt.Sprintf("Weapon Selection"),
			selection_page,
			true, true) //resize bool, visible bool
		for page := 0; page < len(page_slice); page++ {
			pages.AddPage(fmt.Sprintf(page_slice[page]), form_slice[page], true, false)
		}

		if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

	},
}

func init() {
	weaponCmd.AddCommand(weaponCreateCmd)

	weaponCreateCmd.Flags().BoolVarP(&border, "border", "b", false, "deactivate border on output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weaponCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weaponCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func make_missile (form *tview.Form, app *tview.Application) Missile {
	wep_name := form.GetFormItemByLabel("Name").(*tview.TextArea).GetText()
	rof := form.GetFormItemByLabel("Rate of Fire").(*tview.InputField).GetText()
	wep_rof, _ := strconv.Atoi(rof)
	pen := form.GetFormItemByLabel("Penetration").(*tview.InputField).GetText()
	wep_pen, _ := strconv.Atoi(pen)
	_, gen := form.GetFormItemByLabel("Generation").(*tview.DropDown).GetCurrentOption()
	wep_gen, _ := strconv.Atoi(gen)
	wep_top := form.GetFormItemByLabel("Top Attack").(*tview.Checkbox).IsChecked()
	wep_limit := form.GetFormItemByLabel("Ammo Limited").(*tview.Checkbox).IsChecked()
	_, wep_aspect := form.GetFormItemByLabel("Aspect").(*tview.DropDown).GetCurrentOption()
	app.Stop()
	built := Missile {
		Name: wep_name,
		Pen: int(wep_pen),
		RateOfFire: int(wep_rof),
		Generation: int(wep_gen),
		TopAttack: wep_top,
		AmmoLimited: wep_limit,
		Aspect: wep_aspect,
	}
	return built
}
