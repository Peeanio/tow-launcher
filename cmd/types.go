package cmd

//nested or lower values
type VehicleArmor struct {
	Front int
	HClass string
	Flank int
}

type GunWeapon struct {
	Pen int
	AmmoType string // heat or he explict, kinetic is inferenced from lack of tag
	Range int
	RateOfFire int
	AntiInfantry int
}

type Missile struct {
	Id int
	Name string
	Pen int
	RateOfFire int
	Generation int
	TopAttack bool
	AmmoLimited bool
	Aspect any //boolean true for all aspect, false for rear. AA missiles only, nil else
}

type Indirect struct {
	Name string
	RangeMin int
	RangeMax int
	Ammo []string
}

type InfantryGun struct {
	AntiInfantry int
	RateOfFireStill int
	RateOfFireMoving int
	Range int
}

type AntiTankWeapon struct {
	Name string
	Pen int
	HEAT bool
	HighExpolsive bool
	RateOfFire int
	Range int
	Close bool //true for close only, false for ranged value
}

type Move struct {
	Amount int
	Notes string
}

type HelicopterMissile struct {
	PeriodStart int
	PeriodEnd int
	RateOfFire int
	Generation int
	TopAttack bool
	RangeMin int
	RangeMax int
	Aspect bool //true for all aspect, false for rear. AA missiles only
	Notes string
}

type Period struct {
	Start int
	End int
}

type VehicleGunUnit struct {
	Name string
	Period Period
	Points int
	Move Move
	Armor VehicleArmor
	Gun []GunWeapon
	Msl any // could be nil or Missile
	Indirect any // clould be nil or Indirect
	Equip []string
	Capacity any // could be nil or float32, fractions expressed as decimal
	Notes []string
}

type Infantry struct {
	Name string
	Period Period
	Points int
	Move Move
	AntiInfantry InfantryGun
	LAW AntiTankWeapon
	MAW any //shouldbe AntiTankWeapon or Missile
	SAM any // could be nil or Missile
	Equip []string
	Capacity float32
	Notes []string
}

type Aircraft struct {
	Name string
	AvailDate int
	LoadRating int
	GPBombs bool
	ClusterBombs bool
	Rockets bool
	Missiles bool
	MGStrafe int
	CannonStrafe int
	A10Strafe int
	Armored bool
}

type Helicopter struct {
	Name string
	Period Period
	BasePV int
	Move Move
	Armor bool //false for soft, true for 0
	Pen int
	ROF int
	RNG int
	AI int
	Capacity float32
	Equip string
	Pods int
	Options []string
}
