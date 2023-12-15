/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"strconv"
	"log"
	_ "github.com/mattn/go-sqlite3"
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
				AddButton("Save", func() {
					antitank := make_antitank(form_slice[0])
					str, _ := json.Marshal(antitank)
					app.Stop()
					fmt.Println(string(str))
					save_antitank(antitank)
				}).
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
					missile := make_missile(form_slice[1])
					str, _ := json.Marshal(missile)
					app.Stop()
					fmt.Println(string(str))
					save_missile(missile)
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
				AddButton("Save", func() {
					indirect := make_indirect(form_slice[2])
					str, _ := json.Marshal(indirect)
					app.Stop()
					fmt.Println(string(str))
					save_indirect(indirect)
				}).
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

func make_missile (form *tview.Form) Missile {
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

func save_missile (newMissile Missile) {
	db, err := sql.Open("sqlite3", "./tow.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO missiles (Id, Name, Pen, RateOfFire, Generation, TopAttack, AmmoLimited, Aspect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(nil, newMissile.Name, newMissile.Pen, newMissile.RateOfFire, newMissile.Generation, newMissile.TopAttack, newMissile.AmmoLimited, newMissile.Aspect)
	defer stmt.Close()

	fmt.Printf("Added missile %v\n", newMissile.Name)
}

func save_indirect (newIndirect Indirect) {
	db, err := sql.Open("sqlite3", "./tow.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO indirects (Id, Name, RangeMax, RangeMin, Ammo) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec(nil, newIndirect.Name, newIndirect.RangeMax, newIndirect.RangeMin, newIndirect.Ammo)
	defer stmt.Close()

	fmt.Printf("Added indirect %v\n", newIndirect.Name)
}

func save_antitank (newAntiTank AntiTankWeapon) {
	db, err := sql.Open("sqlite3", "./tow.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO antitankweapons (Id, Name, Pen, HEAT, HighExplosive, RateOfFire, Range, Close) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(nil, newAntiTank.Name, newAntiTank.Pen, newAntiTank.HEAT, newAntiTank.HighExplosive, newAntiTank.RateOfFire, newAntiTank.Range, newAntiTank.Close)
	defer stmt.Close()

	fmt.Printf("Added antitank %v\n", newAntiTank.Name)
}

func make_antitank (form *tview.Form) AntiTankWeapon {
	var wep_he bool
	var wep_heat bool
	wep_name := form.GetFormItemByLabel("Name").(*tview.TextArea).GetText()
	rof := form.GetFormItemByLabel("Rate of Fire").(*tview.InputField).GetText()
	wep_rof, _ := strconv.Atoi(rof)
	pen := form.GetFormItemByLabel("Penetration").(*tview.InputField).GetText()
	wep_pen, _ := strconv.Atoi(pen)
	local_range := form.GetFormItemByLabel("Range").(*tview.InputField).GetText()
	wep_range, _ := strconv.Atoi(local_range)
	_, munition := form.GetFormItemByLabel("Munition Type").(*tview.DropDown).GetCurrentOption()
	if munition == "HE" {
		wep_he = true
		wep_heat = false
	} else {
		wep_he = false
		wep_heat = true
	}
	wep_close := form.GetFormItemByLabel("Close Range").(*tview.Checkbox).IsChecked()
	built := AntiTankWeapon {
		Name: wep_name,
		Pen: wep_pen,
		HEAT: wep_heat,
		HighExplosive: wep_he,
		RateOfFire: wep_rof,
		Range: wep_range,
		Close: wep_close,
	}
	return built
}

func make_indirect (form *tview.Form) Indirect {
	var wep_ammo string
	wep_name := form.GetFormItemByLabel("Name").(*tview.TextArea).GetText()
	max := form.GetFormItemByLabel("Range Max").(*tview.InputField).GetText()
	range_max, _ := strconv.Atoi(max)
	min := form.GetFormItemByLabel("Range Min").(*tview.InputField).GetText()
	range_min, _ := strconv.Atoi(min)
	wep_h := form.GetFormItemByLabel("High Explosive Munitions").(*tview.Checkbox).IsChecked()
	wep_s := form.GetFormItemByLabel("Smoke Munitions").(*tview.Checkbox).IsChecked()
	wep_c := form.GetFormItemByLabel("Chemical Munitions").(*tview.Checkbox).IsChecked()
	wep_i := form.GetFormItemByLabel("ICM Munitions").(*tview.Checkbox).IsChecked()
	wep_l := form.GetFormItemByLabel("Laser-guided Munitions").(*tview.Checkbox).IsChecked()
	wep_g := form.GetFormItemByLabel("GPS-guided Munitions").(*tview.Checkbox).IsChecked()
	wep_m := form.GetFormItemByLabel("Artillery-delivered Mine Munitions").(*tview.Checkbox).IsChecked()

	if wep_h == true {
		wep_ammo = wep_ammo + "h"
	}
	if wep_s == true {
		wep_ammo = wep_ammo + "s"
	}
	if wep_c == true {
		wep_ammo = wep_ammo +"c"
	}
	if wep_i == true {
		wep_ammo = wep_ammo +"i"
	}
	if wep_l == true {
		wep_ammo = wep_ammo +"l"
	}
	if wep_g == true {
		wep_ammo = wep_ammo + "g"
	}
	if wep_m == true {
		wep_ammo = wep_ammo + "m"
	}
	built := Indirect {
		Name: wep_name,
		RangeMin: range_min,
		RangeMax: range_max,
		Ammo: wep_ammo,
	}
	return built
}
