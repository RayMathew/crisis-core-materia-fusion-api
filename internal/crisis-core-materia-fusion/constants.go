package crisiscoremateriafusion

type MateriaType string

const (
	Fire              MateriaType = "Fire"
	Ice               MateriaType = "Ice"
	Lightning         MateriaType = "Lightning"
	Restore           MateriaType = "Restore"
	FullCure          MateriaType = "Full Cure"
	Defense           MateriaType = "Defense"
	StatusDefense     MateriaType = "Status Defense"
	AbsorbMagic       MateriaType = "Absorb Magic"
	StatusMagic       MateriaType = "Status Magic"
	FireStatus        MateriaType = "Fire & Status"
	IceStatus         MateriaType = "Ice & Status"
	LightningStatus   MateriaType = "Lightning & Status"
	Gravity           MateriaType = "Gravity"
	Ultimate          MateriaType = "Ultimate"
	QuickAttack       MateriaType = "Quick Attack"
	QuickAttackStatus MateriaType = "Quick Attack & Status"
	BladeArts         MateriaType = "Blade Arts"
	BladeArtsStatus   MateriaType = "Blade Arts & Status"
	FireBlade         MateriaType = "Fire Blade"
	IceBlade          MateriaType = "Ice Blade"
	LightningBlade    MateriaType = "Lightning Blade"
	AbsorbBlade       MateriaType = "Absorb Blade"
	Item              MateriaType = "Item"
	Punch             MateriaType = "Punch"
	HPUp              MateriaType = "HP Up"
	MPUp              MateriaType = "MP Up"
	APUp              MateriaType = "AP Up"
	ATKUp             MateriaType = "ATK Up"
	VITUp             MateriaType = "VIT Up"
	MAGUp             MateriaType = "MAG Up"
	SPRUp             MateriaType = "SPR Up"
	SPTurbo           MateriaType = "SP Turbo"
	Libra             MateriaType = "Libra"
	Dash              MateriaType = "Dash"
	Dualcast          MateriaType = "Dualcast"
	DMW               MateriaType = "DMW"
)

type CacheKey string

const (
	AllMateriaCacheKey CacheKey = "allMateria"
)
