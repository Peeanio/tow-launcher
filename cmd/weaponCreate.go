/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"
	"log"
	. "github.com/rthornton128/goncurses"
	"github.com/spf13/cobra"
)
const (
	HEIGHT = 10
	WIDTH  = 30
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
		var active int // track selected
		src, err := Init()               // setup starts
		if err != nil {
			log.Fatal("init:", err)
		}
		defer End()
		Raw(true)
		Echo(false)
		Cursor(0)
		src.Clear()
		src.Keypad(true)                 // setup ends

		top_slice := []string{"LAW", "MAW", "Missle - AT", "Missile - AA", "Indirect", "AntiInfantry"}
		top_menu, top_items := create_menu(top_slice)
		maw_slice := []string{"Pen", "HEAT", "HighExpolsive", "RateOfFire", "Range", "Close", "Back", "Save"}
		maw_menu, maw_items := create_menu(maw_slice)
		// maw_items := make([]*MenuItem, len(maw_slice))
		var active_menu *Menu = top_menu
		var active_slice []string = top_slice

		active_menu.Post()

		src.MovePrint(20, 0, "'q' to exit")
		src.Refresh()

		for {
			defer active_menu.Free()
			Update()
			ch := src.GetChar()

			switch Key(ch) {
			case KEY_ESC, 3:
				return
			case KEY_DOWN:
				if active == len(active_slice)-1 {
					active = 0
				} else {
					active += 1
				}
				active_menu.Driver(REQ_DOWN)
			case KEY_UP:
				if active == 0 {
					active = len(active_slice) - 1
				} else {
					active -= 1
				}
				active_menu.Driver(REQ_UP)
			case KEY_RETURN:
				Flash()
				src.Clear()
				if active_menu == top_menu {
					switch active_slice[active] {
					case "LAW":
						src.Print("LAW")
					case "MAW":
						src.Print("MAW")
						active_menu = maw_menu
					}
				} else if active_menu == maw_menu {
					src.Print("Inside maw") //create fillable forms
				}
				active_menu.Post()
				//src.Print(top_slice[active])
				src.Refresh()
				// switch active_menu.Current() {
		 	// 	case "MAW":
		 	// 		active_menu = maw_menu
				// }
			// default:
			// 	src.Print(ch)
			}

		}
		log.Print(maw_items)
		log.Print(top_items)

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

func create_menu(slice_items []string) (*Menu, []*MenuItem) {
	items := make([]*MenuItem, len(slice_items))
	for i, val := range slice_items {
		items[i], _ = NewItem(val, "")
		//defer items[i].Free()
	}
	menu, err := NewMenu(items)
	if err != nil {
		log.Fatal(err)
	}
	//defer menu.Free()
	return menu, items
}
