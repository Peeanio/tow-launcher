package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	Scout := Helicopter {
		Name: "Scout",
		Period: Period{
			Start: 63,
			End: 94,
		},
		BasePV: 126,
		Move: Move{
			Amount: 18,
			Notes: "n",
		},
		Armor: false,
		Pen: 0,
		ROF: 0,
		RNG: 0,
		AI: 0,
		Capacity: 0,
		Equip: "",
		Pods: 2,
		Options: []string{"MG", "rocket", "20mm cannon", "SS-11"},
	}

	Warrior := VehicleGunUnit{
		Name: "Warrior",
		Period: Period{
			Start: 8,
			End: 15,
		},
		Points: 165,
		Move: Move{
			Amount: 8,
			Notes: "t",
		},
		Armor: VehicleArmor{
			Front: 7,
			HClass: "d",
			Flank: 7,
		},
		Gun: []GunWeapon{{
			Pen: 6,
			AmmoType: "kinetic",
			RateOfFire: 4,
			Range: 12,
			AntiInfantry: 0,
		},},
		Msl: nil,
		Indirect: nil,
		Equip: []string{"t", "n"},
		Capacity: 1,
		Notes: []string{"30mm R (UK)"},
	}

	Infantry := Infantry{
		Name: "Infantry",
		Period: Period{
			Start: 9,
			End: 15,
		},
		Points: 40,
		Move: Move{
			Amount: 4,
			Notes: "",
		},
		AntiInfantry: InfantryGun{
			AntiInfantry: -1,
			RateOfFireStill: 6,
			RateOfFireMoving: 2,
			Range: 4,
		},
		LAW: AntiTankWeapon {
			Pen: 2,
			HEAT: false,
			HighExpolsive: true,
			RateOfFire: 1,
			Range: 0,
			Close: true,
		},
		MAW: Missile {
			Name: "NLAW"
			Pen: 14,
			RateOfFire: 1,
			Generation: 3,
			TopAttack: true,
			AmmoLimited: true,
			Aspect: nil,
		},
		SAM: nil,
		Equip: []string{"t"},
		Capacity: 1,
		Notes: []string{"Improvised/MBT LAW"},
	}

	Abbot := VehicleGunUnit{
		Name: "Abbot",
		Period: Period{
			Start: 65,
			End: 95,
		},
		Points: 76,
		Move: Move{
			Amount: 7,
			Notes: "t",
		},
		Armor: VehicleArmor{
			Front: 1,
			HClass: "",
			Flank: 0,
		},
		Gun: []GunWeapon{{
			Pen: 11,
			AmmoType: "heat",
			RateOfFire: 3,
			Range: 10,
			AntiInfantry: -1,
		},},
		Msl: nil,
		Indirect: Indirect{
			RangeMin: 0,
			RangeMax: 172,
			Ammo: []string{"h", "s"},
		},
		Equip: []string{"n", "1"},
		Capacity: nil,
		Notes: []string{"105mm/L37 R (UK)"},
	}

	F22 := Aircraft {
		Name: "F-22 Raptor",
		AvailDate: 05,
		LoadRating: 2,
		GPBombs: false,
		ClusterBombs: true,
		Rockets: false,
		Missiles: true,
		MGStrafe: 0,
		CannonStrafe: 1,
		A10Strafe: 0,
		Armored: false,
	}

	fmt.Println(Scout)
	fmt.Println(F22)
	fmt.Println(Infantry)
	fmt.Println(Warrior)
	fmt.Println(Abbot)
	str_abb, _ := json.Marshal(Abbot)
	fmt.Println(string(str_abb))
}

