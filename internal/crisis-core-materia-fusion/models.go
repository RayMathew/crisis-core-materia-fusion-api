package crisiscoremateriafusion

type Materia struct {
	Name        string `json:"name"`
	Type        string `json:"materia_type"`
	Grade       int    `json:"grade"`
	DisplayType string `json:"display_type"`
	Description string `json:"description"`
}

type BasicCombinationRule struct {
	FirstMateriaType     MateriaType
	SecondMateriaType    MateriaType
	ResultantMateriaType MateriaType
}

// meaning of "inverse rule NOT applicable": exchanging FirstMateriaType and SecondMateriaType gives a different ResultantMateriaType

var FITBasicRules = []BasicCombinationRule{
	{Fire, Fire, Fire},
	{Ice, Ice, Ice},
	{Lightning, Lightning, Lightning},
	{Fire, Ice, Lightning},
	{Ice, Fire, Lightning}, // inverse of above rule
	{Fire, Lightning, Ice},
	{Lightning, Fire, Ice}, // inverse of above rule
	{Lightning, Ice, Fire},
	{Ice, Lightning, Fire}, // inverse of above rule

	// FIT + Restore Rules
	{Fire, Restore, Fire},
	{Ice, Restore, Ice},
	{Lightning, Restore, Lightning},

	// FIT + Defense Rules are not basic

	// FIT + Status Defense Rules
	{Fire, StatusDefense, FireStatus},
	{Ice, StatusDefense, IceStatus},
	{Lightning, StatusDefense, LightningStatus},

	// FIT + Absorb Magic Rules
	{Fire, AbsorbMagic, AbsorbMagic},
	{Ice, AbsorbMagic, AbsorbMagic},
	{Lightning, AbsorbMagic, AbsorbMagic},

	// FIT + Status Magic Rules
	{Fire, StatusMagic, FireStatus},
	{Ice, StatusMagic, IceStatus},
	{Lightning, StatusMagic, LightningStatus},

	// Fire + FITStatus Rules
	{Fire, FireStatus, FireStatus},
	{Fire, IceStatus, LightningStatus},
	{Fire, LightningStatus, IceStatus},

	// Ice + FITStatus Rules
	{Ice, IceStatus, IceStatus},
	{Ice, FireStatus, LightningStatus},
	{Ice, LightningStatus, FireStatus},

	// Lightning + FITStatus Rules
	{Lightning, LightningStatus, LightningStatus},
	{Lightning, FireStatus, IceStatus},
	{Lightning, IceStatus, FireStatus},

	// FIT + Gravity Rules are not basic

	// FIT + Ultimate Rules
	{Fire, Ultimate, Fire},
	{Ice, Ultimate, Ice},
	{Lightning, Ultimate, Lightning},

	// FIT + Quick Attack Rules
	{Fire, QuickAttack, FireBlade},
	{Ice, QuickAttack, IceBlade},
	{Lightning, QuickAttack, LightningBlade},

	// FIT + Quick Attack Status Rules
	{Fire, QuickAttackStatus, FireStatus},
	{Ice, QuickAttackStatus, IceStatus},
	{Lightning, QuickAttackStatus, LightningStatus},

	// FIT + Blade Arts Rules
	{Fire, BladeArts, FireBlade},
	{Ice, BladeArts, IceBlade},
	{Lightning, BladeArts, LightningBlade},

	// FIT + Blade Arts Status Rules
	{Fire, BladeArtsStatus, FireStatus},
	{Ice, BladeArtsStatus, IceStatus},
	{Lightning, BladeArtsStatus, LightningStatus},

	// Fire + FIT Blade Rules
	{Fire, FireBlade, FireBlade},
	{Fire, IceBlade, LightningBlade},
	{Fire, LightningBlade, IceBlade},

	// Ice + FIT Blade Rules
	{Ice, FireBlade, LightningBlade},
	{Ice, IceBlade, IceBlade},
	{Ice, LightningBlade, FireBlade},

	// Lightning + FIT Blade Rules
	{Lightning, FireBlade, IceBlade},
	{Lightning, IceBlade, FireBlade},
	{Lightning, LightningBlade, LightningBlade},

	// FIT + Absorb Blade Rules
	{Fire, AbsorbBlade, AbsorbMagic},
	{Ice, AbsorbBlade, AbsorbMagic},
	{Lightning, AbsorbBlade, AbsorbMagic},

	// FIT + Item Rules are not basic

	// FIT + Punch Rules
	{Fire, Punch, Fire},
	{Ice, Punch, Ice},
	{Lightning, Punch, Lightning},

	// FIT + HP Up Rules
	{Fire, HPUp, Defense},
	{Ice, HPUp, Defense},
	{Lightning, HPUp, Defense},

	// FIT + MP Up Rules
	{Fire, MPUp, Fire},
	{Ice, MPUp, Ice},
	{Lightning, MPUp, Lightning},

	// FIT + AP Up Rules
	{Fire, APUp, Fire},
	{Ice, APUp, Ice},
	{Lightning, APUp, Lightning},

	// FIT + ATK Up Rules
	{Fire, ATKUp, FireBlade},
	{Ice, ATKUp, IceBlade},
	{Lightning, ATKUp, LightningBlade},

	// FIT + VIT Up Rules
	{Fire, VITUp, Defense},
	{Ice, VITUp, Defense},
	{Lightning, VITUp, Defense},

	// FIT + MAG Up Rules
	{Fire, MAGUp, Fire},
	{Ice, MAGUp, Ice},
	{Lightning, MAGUp, Lightning},

	// FIT + SPR Up Rules
	{Fire, SPRUp, Defense},
	{Ice, SPRUp, Defense},
	{Lightning, SPRUp, Defense},

	// FIT + SP Turbo Rules
	{Fire, SPTurbo, Fire},
	{Ice, SPTurbo, Ice},
	{Lightning, SPTurbo, Lightning},

	// FIT + Libra Rules
	{Fire, Libra, Fire},
	{Ice, Libra, Ice},
	{Lightning, Libra, Lightning},
}

var RestoreBasicRules = []BasicCombinationRule{
	// Restore + FIT Rules
	{Restore, Fire, Restore},
	{Restore, Ice, Restore},
	{Restore, Lightning, Restore},

	// Restore + Restore Rules
	{Restore, Restore, Restore},

	// Restore + Defense Rules are not basic

	// Restore + Status Defense Rules
	{Restore, StatusDefense, Restore},

	// Restore + Absorb Magic Rules
	{Restore, AbsorbMagic, Restore},

	// Restore + Status Magic Rules
	{Restore, StatusMagic, Restore},

	// Restore + FITStatus Rules
	{Restore, FireStatus, Restore},
	{Restore, IceStatus, Restore},
	{Restore, LightningStatus, Restore},

	// Restore + Gravity Rules are not basic

	// Restore + Ultimate Rules
	{Restore, Ultimate, Restore},

	// Restore + Quick Attack Rules
	{Restore, QuickAttack, Restore},

	// Restore + Quick Attack Status Rules
	{Restore, QuickAttackStatus, Restore},

	// Restore + Blade Arts Rules
	{Restore, BladeArts, Restore},

	// Restore + Blade Arts Status Rules
	{Restore, BladeArtsStatus, Restore},

	// Restore + FIT Blade Rules
	{Restore, FireBlade, Restore},
	{Restore, IceBlade, Restore},
	{Restore, LightningBlade, Restore},

	// Restore + Absorb Blade Rules
	{Restore, AbsorbBlade, Restore},

	// Restore + Item Rules are not basic

	// Restore + Punch Rules
	{Restore, Punch, Restore},

	// Restore + HP Up Rules
	{Restore, HPUp, Restore},

	// Restore + MP Up Rules
	{Restore, MPUp, Restore},

	// Restore + AP Up Rules
	{Restore, APUp, Restore},

	// Restore + ATK Up Rules
	{Restore, ATKUp, Restore},

	// Restore + VIT Up Rules
	{Restore, VITUp, Restore},

	// Restore + MAG Up Rules
	{Restore, MAGUp, Restore},

	// Restore + SPR Up Rules
	{Restore, SPRUp, Restore},

	// Restore + SP Turbo Rules
	{Restore, SPTurbo, Restore},

	// Restore + Libra Rules
	{Restore, Libra, Restore},
}

var DefenseBasicRules = []BasicCombinationRule{
	// Defense + FIT Rules
	{Defense, Fire, Defense},
	{Defense, Ice, Defense},
	{Defense, Lightning, Defense},

	// Defense+ Restore Rules
	{Defense, Restore, Defense},

	// Defense + Defense Rules
	{Defense, Defense, Defense},

	// Defense + Status Defense Rules
	{Defense, StatusDefense, Defense},

	// Defense + Absorb Magic Rules
	{Defense, AbsorbMagic, Defense},

	// Defense + Status Magic Rules are not basic

	// Defense + FITStatus Rules are not basic

	// Defense + Gravity Rules are not basic

	// Defense + Ultimate Rules
	{Defense, Ultimate, Defense},

	// Defense + Quick Attack Rules
	{Defense, QuickAttack, Defense},

	// Defense + Quick Attack Status Rules are not basic

	// Defense + Blade Arts Rules
	{Defense, BladeArts, Defense},

	// Defense + Blade Arts Status Rules are not basic

	// Defense + FIT Blade Rules
	{Defense, FireBlade, Defense},
	{Defense, IceBlade, Defense},
	{Defense, LightningBlade, Defense},

	// Defense + Absorb Blade Rules
	{Defense, AbsorbBlade, Defense},

	// Defense + Item Rules are not basic

	// Defense + Punch Rules
	{Defense, Punch, Defense},

	// Defense + HP Up Rules
	{Defense, HPUp, Defense},

	// Defense + MP Up Rules
	{Defense, MPUp, Defense},

	// Defense + AP Up Rules
	{Defense, APUp, Defense},

	// Defense + ATK Up Rules
	{Defense, ATKUp, Defense},

	// Defense + VIT Up Rules
	{Defense, VITUp, Defense},

	// Defense + MAG Up Rules
	{Defense, MAGUp, Defense},

	// Defense + SPR Up Rules
	{Defense, SPRUp, Defense},

	// Defense + SP Turbo Rules
	{Defense, SPTurbo, Defense},

	// Defense + Libra Rules
	{Defense, Libra, Defense},
}

var StatusDefenseBasicRules = []BasicCombinationRule{
	// StatusDefense + FIT Rules
	{StatusDefense, Fire, StatusDefense},
	{StatusDefense, Ice, StatusDefense},
	{StatusDefense, Lightning, StatusDefense},

	// StatusDefense + Restore Rules
	{StatusDefense, Restore, StatusDefense},

	// StatusDefense + Defense Rules
	{StatusDefense, Defense, Defense},

	// StatusDefense + Status Defense Rules
	{StatusDefense, StatusDefense, StatusDefense},

	// StatusDefense + Absorb Magic Rules
	{StatusDefense, AbsorbMagic, StatusDefense},

	// StatusDefense + Status Magic Rules
	{StatusDefense, StatusMagic, StatusDefense},

	// StatusDefense + FITStatus Rules
	{StatusDefense, FireStatus, StatusDefense},
	{StatusDefense, IceStatus, StatusDefense},
	{StatusDefense, LightningStatus, StatusDefense},

	// StatusDefense + Gravity Rules
	{StatusDefense, Gravity, StatusDefense},

	// StatusDefense + Ultimate Rules
	{StatusDefense, Ultimate, StatusDefense},

	// StatusDefense + Quick Attack Rules
	{StatusDefense, QuickAttack, Defense},

	// StatusDefense + Quick Attack Status Rules
	{StatusDefense, QuickAttackStatus, StatusDefense},

	// StatusDefense + Blade Arts Rules
	{StatusDefense, BladeArts, Defense},

	// StatusDefense + Blade Arts Status Rules
	{StatusDefense, BladeArtsStatus, StatusDefense},

	// StatusDefense + FIT Blade Rules
	{StatusDefense, FireBlade, Defense},
	{StatusDefense, IceBlade, Defense},
	{StatusDefense, LightningBlade, Defense},

	// StatusDefense + Absorb Blade Rules
	{StatusDefense, AbsorbBlade, StatusDefense},

	// StatusDefense + Item Rules
	{StatusDefense, Item, StatusDefense},

	// StatusDefense + Punch Rules
	{StatusDefense, Punch, StatusDefense},

	// StatusDefense + HP Up Rules
	{StatusDefense, HPUp, StatusDefense},

	// StatusDefense + MP Up Rules
	{StatusDefense, MPUp, StatusDefense},

	// StatusDefense + AP Up Rules
	{StatusDefense, APUp, StatusDefense},

	// StatusDefense + ATK Up Rules
	{StatusDefense, ATKUp, StatusDefense},

	// StatusDefense + VIT Up Rules
	{StatusDefense, VITUp, StatusDefense},

	// StatusDefense + MAG Up Rules
	{StatusDefense, MAGUp, StatusDefense},

	// StatusDefense + SPR Up Rules
	{StatusDefense, SPRUp, StatusDefense},

	// StatusDefense + SP Turbo Rules
	{StatusDefense, SPTurbo, StatusDefense},

	// StatusDefense + Libra Rules
	{StatusDefense, Libra, StatusDefense},
}

var AbsorbMagicBasicRules = []BasicCombinationRule{
	// AbsorbMagic + FIT Rules
	{AbsorbMagic, Fire, AbsorbMagic},
	{AbsorbMagic, Ice, AbsorbMagic},
	{AbsorbMagic, Lightning, AbsorbMagic},

	// AbsorbMagic + Restore Rules
	{AbsorbMagic, Restore, AbsorbMagic},

	// AbsorbMagic + Defense Rules
	{AbsorbMagic, Defense, AbsorbMagic},

	// AbsorbMagic + Status Defense Rules
	{AbsorbMagic, StatusDefense, AbsorbMagic},

	// AbsorbMagic + Absorb Magic Rules
	{AbsorbMagic, AbsorbMagic, AbsorbMagic},

	// AbsorbMagic + Status Magic Rules
	{AbsorbMagic, StatusMagic, AbsorbMagic},

	// AbsorbMagic + FITStatus Rules
	{AbsorbMagic, FireStatus, AbsorbMagic},
	{AbsorbMagic, IceStatus, AbsorbMagic},
	{AbsorbMagic, LightningStatus, AbsorbMagic},

	// AbsorbMagic + Gravity Rules are not basic

	// AbsorbMagic + Ultimate Rules
	{AbsorbMagic, Ultimate, AbsorbMagic},

	// AbsorbMagic + Quick Attack Rules
	{AbsorbMagic, QuickAttack, AbsorbMagic},

	// AbsorbMagic + Quick Attack Status Rules
	{AbsorbMagic, QuickAttackStatus, AbsorbMagic},

	// AbsorbMagic + Blade Arts Rules
	{AbsorbMagic, BladeArts, AbsorbMagic},

	// AbsorbMagic + Blade Arts Status Rules
	{AbsorbMagic, BladeArtsStatus, AbsorbMagic},

	// AbsorbMagic + FIT Blade Rules
	{AbsorbMagic, FireBlade, AbsorbMagic},
	{AbsorbMagic, IceBlade, AbsorbMagic},
	{AbsorbMagic, LightningBlade, AbsorbMagic},

	// AbsorbMagic + Absorb Blade Rules
	{AbsorbMagic, AbsorbBlade, AbsorbMagic},

	// AbsorbMagic + Item Rules are not basic

	// AbsorbMagic + Punch Rules
	{AbsorbMagic, Punch, AbsorbMagic},

	// AbsorbMagic + HP Up Rules
	{AbsorbMagic, HPUp, AbsorbMagic},

	// AbsorbMagic + MP Up Rules
	{AbsorbMagic, MPUp, AbsorbMagic},

	// AbsorbMagic + AP Up Rules
	{AbsorbMagic, APUp, AbsorbMagic},

	// AbsorbMagic + ATK Up Rules are not basic

	// AbsorbMagic + VIT Up Rules are not basic

	// AbsorbMagic + MAG Up Rules
	{AbsorbMagic, MAGUp, AbsorbMagic},

	// AbsorbMagic + SPR Up Rules
	{AbsorbMagic, SPRUp, AbsorbMagic},

	// AbsorbMagic + SP Turbo Rules
	{AbsorbMagic, SPTurbo, AbsorbMagic},

	// AbsorbMagic + Libra Rules
	{AbsorbMagic, Libra, AbsorbMagic},
}

var StatusMagicBasicRules = []BasicCombinationRule{
	// StatusMagic + FIT Rules
	{StatusMagic, Fire, FireStatus},
	{StatusMagic, Ice, IceStatus},
	{StatusMagic, Lightning, LightningStatus},

	// StatusMagic + Restore Rules
	{StatusMagic, Restore, Restore},

	// StatusMagic + Defense Rules are not basic

	// StatusMagic + Status Defense Rules
	{StatusMagic, StatusDefense, StatusMagic},

	// StatusMagic + Absorb Magic Rules
	{StatusMagic, AbsorbMagic, AbsorbMagic},

	// StatusMagic + Status Magic Rules
	{StatusMagic, StatusMagic, StatusMagic},

	// StatusMagic + FITStatus Rules
	{StatusMagic, FireStatus, FireStatus},
	{StatusMagic, IceStatus, IceStatus},
	{StatusMagic, LightningStatus, IceStatus},

	// StatusMagic + Gravity Rules
	{StatusMagic, Gravity, StatusMagic},

	// StatusMagic + Ultimate Rules
	{StatusMagic, Ultimate, StatusMagic},

	// StatusMagic + Quick Attack Rules
	{StatusMagic, QuickAttack, BladeArtsStatus},

	// StatusMagic + Quick Attack Status Rules
	{StatusMagic, QuickAttackStatus, StatusMagic},

	// StatusMagic + Blade Arts Rules
	{StatusMagic, BladeArts, BladeArtsStatus},

	// StatusMagic + Blade Arts Status Rules
	{StatusMagic, BladeArtsStatus, StatusMagic},

	// StatusMagic + FIT Blade Rules
	{StatusMagic, FireBlade, BladeArtsStatus},
	{StatusMagic, IceBlade, BladeArtsStatus},
	{StatusMagic, LightningBlade, BladeArtsStatus},

	// StatusMagic + Absorb Blade Rules
	{StatusMagic, AbsorbBlade, AbsorbMagic},

	// StatusMagic + Item Rules are not basic

	// StatusMagic + Punch Rules
	{StatusMagic, Punch, StatusMagic},

	// StatusMagic + HP Up Rules
	{StatusMagic, HPUp, StatusDefense},

	// StatusMagic + MP Up Rules
	{StatusMagic, MPUp, StatusMagic},

	// StatusMagic + AP Up Rules
	{StatusMagic, APUp, StatusMagic},

	// StatusMagic + ATK Up Rules
	{StatusMagic, ATKUp, BladeArtsStatus},

	// StatusMagic + VIT Up Rules
	{StatusMagic, VITUp, StatusDefense},

	// StatusMagic + MAG Up Rules
	{StatusMagic, MAGUp, StatusMagic},

	// StatusMagic + SPR Up Rules
	{StatusMagic, SPRUp, StatusDefense},

	// StatusMagic + SP Turbo Rules
	{StatusMagic, SPTurbo, StatusMagic},

	// StatusMagic + Libra Rules
	{StatusMagic, Libra, StatusMagic},
}

var FITStatusBasicRules = []BasicCombinationRule{
	//FireStatus + FIT Rules
	{FireStatus, Fire, FireStatus},
	{FireStatus, Ice, LightningStatus},
	{FireStatus, Lightning, IceStatus},

	//IceStatus + FIT Rules
	{IceStatus, Fire, LightningStatus},
	{IceStatus, Ice, IceStatus},
	{IceStatus, Lightning, FireStatus},

	//LightningStatus + FIT Rules
	{LightningStatus, Fire, IceStatus},
	{LightningStatus, Ice, FireStatus},
	{LightningStatus, Lightning, LightningStatus},

	// FITStatus + Restore Rules
	{FireStatus, Restore, Fire},
	{IceStatus, Restore, Ice},
	{LightningStatus, Restore, Lightning},

	// FITStatus + Defense Rules are not basic

	// FITStatus + Status Defense Rules
	{FireStatus, StatusDefense, Fire},
	{IceStatus, StatusDefense, Ice},
	{LightningStatus, StatusDefense, Lightning},

	// FITStatus + Absorb Magic Rules
	{FireStatus, AbsorbMagic, AbsorbMagic},
	{IceStatus, AbsorbMagic, AbsorbMagic},
	{LightningStatus, AbsorbMagic, AbsorbMagic},

	// FITStatus + Status Magic Rules
	{FireStatus, StatusMagic, FireStatus},
	{IceStatus, StatusMagic, IceStatus},
	{LightningStatus, StatusMagic, LightningStatus},

	// FireStatus + FITStatus Rules
	{FireStatus, FireStatus, FireStatus},
	{FireStatus, IceStatus, LightningStatus},
	{FireStatus, LightningStatus, IceStatus},

	// IceStatus + FITStatus Rules
	{IceStatus, IceStatus, IceStatus},
	{IceStatus, FireStatus, LightningStatus},
	{IceStatus, LightningStatus, FireStatus},

	// LightningStatus + FITStatus Rules
	{LightningStatus, LightningStatus, LightningStatus},
	{LightningStatus, FireStatus, IceStatus},
	{LightningStatus, IceStatus, FireStatus},

	// FITStatus + Gravity Rules are not basic

	// FITStatus + Ultimate Rules
	{FireStatus, Ultimate, FireStatus},
	{IceStatus, Ultimate, IceStatus},
	{LightningStatus, Ultimate, LightningStatus},

	// FITStatus + Quick Attack Rules
	{FireStatus, QuickAttack, FireBlade},
	{IceStatus, QuickAttack, IceBlade},
	{LightningStatus, QuickAttack, LightningBlade},

	// FITStatus + Quick Attack Status Rules
	{FireStatus, QuickAttackStatus, FireStatus},
	{IceStatus, QuickAttackStatus, IceStatus},
	{LightningStatus, QuickAttackStatus, LightningStatus},

	// FITStatus + Blade Arts Rules
	{FireStatus, BladeArts, FireBlade},
	{IceStatus, BladeArts, IceBlade},
	{LightningStatus, BladeArts, LightningBlade},

	// FITStatus + Blade Arts Status Rules
	{FireStatus, BladeArtsStatus, FireStatus},
	{IceStatus, BladeArtsStatus, IceStatus},
	{LightningStatus, BladeArtsStatus, LightningStatus},

	// FireStatus + FIT Blade Rules
	{FireStatus, FireBlade, FireStatus},
	{FireStatus, IceBlade, LightningStatus},
	{FireStatus, LightningBlade, IceStatus},

	// IceStatus + FIT Blade Rules
	{IceStatus, FireBlade, LightningStatus},
	{IceStatus, IceBlade, IceStatus},
	{IceStatus, LightningBlade, FireStatus},

	// LightningStatus + FIT Blade Rules
	{LightningStatus, FireBlade, IceStatus},
	{LightningStatus, IceBlade, FireStatus},
	{LightningStatus, LightningBlade, LightningStatus},

	// FITStatus + Absorb Blade Rules
	{FireStatus, AbsorbBlade, AbsorbMagic},
	{IceStatus, AbsorbBlade, AbsorbMagic},
	{LightningStatus, AbsorbBlade, AbsorbMagic},

	// FITStatus + Item Rules are not basic

	// FITStatus + Punch Rules
	{FireStatus, Punch, FireStatus},
	{IceStatus, Punch, IceStatus},
	{LightningStatus, Punch, LightningStatus},

	// FITStatus + HP Up Rules
	{FireStatus, HPUp, StatusDefense},
	{IceStatus, HPUp, StatusDefense},
	{LightningStatus, HPUp, StatusDefense},

	// FITStatus + MP Up Rules
	{FireStatus, MPUp, FireStatus},
	{IceStatus, MPUp, IceStatus},
	{LightningStatus, MPUp, LightningStatus},

	// FITStatus + AP Up Rules
	{FireStatus, APUp, FireStatus},
	{IceStatus, APUp, IceStatus},
	{LightningStatus, APUp, LightningStatus},

	// FITStatus + ATK Up Rules
	{FireStatus, ATKUp, FireBlade},
	{IceStatus, ATKUp, IceBlade},
	{LightningStatus, ATKUp, LightningBlade},

	// FITStatus + VIT Up Rules
	{FireStatus, VITUp, StatusDefense},
	{IceStatus, VITUp, StatusDefense},
	{LightningStatus, VITUp, StatusDefense},

	// FITStatus + MAG Up Rules
	{FireStatus, MAGUp, FireStatus},
	{IceStatus, MAGUp, IceStatus},
	{LightningStatus, MAGUp, LightningStatus},

	// FITStatus + SPR Up Rules
	{FireStatus, SPRUp, StatusDefense},
	{IceStatus, SPRUp, StatusDefense},
	{LightningStatus, SPRUp, StatusDefense},

	// FITStatus + SP Turbo Rules
	{FireStatus, SPTurbo, FireStatus},
	{IceStatus, SPTurbo, IceStatus},
	{LightningStatus, SPTurbo, LightningStatus},

	// FITStatus + Libra Rules
	{FireStatus, Libra, FireStatus},
	{IceStatus, Libra, IceStatus},
	{LightningStatus, Libra, LightningStatus},
}

var GravityBasicRules = []BasicCombinationRule{
	// Gravity + FIT Rules
	{Gravity, Fire, Gravity},
	{Gravity, Ice, Gravity},
	{Gravity, Lightning, Gravity},

	// Gravity + Restore Rules
	{Gravity, Restore, Gravity},

	// Gravity + Defense Rules
	{Gravity, Defense, Gravity},

	// Gravity + Status Defense Rules
	{Gravity, StatusDefense, Gravity},

	// Gravity + Absorb Magic Rules are not basic

	// Gravity + Status Magic Rules are not basic

	// Gravity + FITStatus Rules
	{Gravity, FireStatus, Gravity},
	{Gravity, IceStatus, Gravity},
	{Gravity, LightningStatus, Gravity},

	// Gravity + Gravity Rules
	{Gravity, Gravity, Gravity},

	// Gravity + Ultimate Rules
	{Gravity, Ultimate, Gravity},

	// Gravity + Quick Attack Rules are not basic

	// Gravity + Quick Attack Status Rules
	{Gravity, QuickAttackStatus, Gravity},

	// Gravity + Blade Arts Rules are not basic

	// Gravity + Blade Arts Status Rules
	{Gravity, BladeArtsStatus, Gravity},

	// Gravity + FIT Blade Rules are not basic

	// Gravity + Absorb Blade Rules are not basic

	// Gravity + Item Rules are not basic

	// Gravity + Punch Rules
	{Gravity, Punch, Gravity},

	// Gravity + HP Up Rules are not basic

	// Gravity + MP Up Rules
	{Gravity, MPUp, Gravity},

	// Gravity + AP Up Rules
	{Gravity, APUp, QuickAttack},

	// Gravity + ATK Up Rules are not basic

	// Gravity + VIT Up Rules are not basic

	// Gravity + MAG Up Rules
	{Gravity, MAGUp, Gravity},

	// Gravity + SPR Up Rules are not basic

	// Gravity + SP Turbo Rules
	{Gravity, SPTurbo, Gravity},

	// Gravity + Libra Rules
	{Gravity, Libra, Gravity},
}

var UltimateBasicRules = []BasicCombinationRule{
	// Ultimate + FIT Rules
	{Ultimate, Fire, Fire},
	{Ultimate, Ice, Ice},
	{Ultimate, Lightning, Lightning},

	// Ultimate + Restore Rules
	{Ultimate, Restore, Restore},

	// Ultimate + Defense Rules
	{Ultimate, Defense, Defense},

	// Ultimate + Status Defense Rules
	{Ultimate, StatusDefense, StatusDefense},

	// Ultimate + Absorb Magic Rules
	{Ultimate, AbsorbMagic, AbsorbMagic},

	// Ultimate + Status Magic Rules
	{Ultimate, StatusMagic, StatusMagic},

	// Ultimate + FITStatus Rules
	{Ultimate, FireStatus, FireStatus},
	{Ultimate, IceStatus, IceStatus},
	{Ultimate, LightningStatus, IceStatus},

	// Ultimate + Gravity Rules
	{Ultimate, Gravity, Gravity},

	// Ultimate + Ultimate Rules
	{Ultimate, Ultimate, Ultimate},

	// Ultimate + Quick Attack Rules
	{Ultimate, QuickAttack, QuickAttack},

	// Ultimate + Quick Attack Status Rules
	{Ultimate, QuickAttackStatus, QuickAttackStatus},

	// Ultimate + Blade Arts Rules
	{Ultimate, BladeArts, BladeArts},

	// Ultimate + Blade Arts Status Rules
	{Ultimate, BladeArtsStatus, BladeArtsStatus},

	// Ultimate + FIT Blade Rules
	{Ultimate, FireBlade, FireBlade},
	{Ultimate, IceBlade, IceBlade},
	{Ultimate, LightningBlade, LightningBlade},

	// Ultimate + Absorb Blade Rules are not basic

	// Ultimate + Item Rules
	{Ultimate, Item, Item},

	// Ultimate + Punch Rules
	{Ultimate, Punch, Ultimate},

	// Ultimate + HP Up Rules
	{Ultimate, HPUp, HPUp},

	// Ultimate + MP Up Rules
	{Ultimate, MPUp, MPUp},

	// Ultimate + AP Up Rules
	{Ultimate, APUp, APUp},

	// Ultimate + ATK Up Rules
	{Ultimate, ATKUp, ATKUp},

	// Ultimate + VIT Up Rules
	{Ultimate, VITUp, VITUp},

	// Ultimate + MAG Up Rules
	{Ultimate, MAGUp, MAGUp},

	// Ultimate + SPR Up Rules
	{Ultimate, SPRUp, SPRUp},

	// Ultimate + SP Turbo Rules
	{Ultimate, SPTurbo, SPTurbo},

	// Ultimate + Libra Rules
	{Ultimate, Libra, Libra},
}

var QuickAttackBasicRules = []BasicCombinationRule{
	// QuickAttack + FIT Rules
	{QuickAttack, Fire, FireBlade},
	{QuickAttack, Ice, IceBlade},
	{QuickAttack, Lightning, LightningBlade},

	// QuickAttack + Restore Rules
	{QuickAttack, Restore, QuickAttack},

	// QuickAttack + Defense Rules are not basic

	// QuickAttack + Status Defense Rules
	{QuickAttack, StatusDefense, QuickAttack},

	// QuickAttack + Absorb Magic Rules are not basic

	// QuickAttack + Status Magic Rules
	{QuickAttack, StatusMagic, BladeArtsStatus},

	// QuickAttack + FITStatus Rules
	{QuickAttack, FireStatus, FireBlade},
	{QuickAttack, IceStatus, IceBlade},
	{QuickAttack, LightningStatus, LightningBlade},

	// QuickAttack + Gravity Rules are not basic

	// QuickAttack + Ultimate Rules
	{QuickAttack, Ultimate, QuickAttack},

	// QuickAttack + Quick Attack Rules
	{QuickAttack, QuickAttack, QuickAttack},

	// QuickAttack + Quick Attack Status Rules
	{QuickAttack, QuickAttackStatus, QuickAttackStatus},

	// QuickAttack + Blade Arts Rules
	{QuickAttack, BladeArts, QuickAttack},

	// QuickAttack + Blade Arts Status Rules
	{QuickAttack, BladeArtsStatus, BladeArtsStatus},

	// QuickAttack + FIT Blade Rules are not basic

	// QuickAttack + Absorb Blade Rules are not basic

	// QuickAttack + Item Rules are not basic

	// QuickAttack + Punch Rules
	{QuickAttack, Punch, QuickAttack},

	// QuickAttack + HP Up Rules
	{QuickAttack, HPUp, QuickAttack},

	// QuickAttack + MP Up Rules
	{QuickAttack, MPUp, QuickAttack},

	// QuickAttack + AP Up Rules
	{QuickAttack, APUp, QuickAttack},

	// QuickAttack + ATK Up Rules
	{QuickAttack, ATKUp, QuickAttack},

	// QuickAttack + VIT Up Rules
	{QuickAttack, VITUp, QuickAttack},

	// QuickAttack + MAG Up Rules
	{QuickAttack, MAGUp, QuickAttack},

	// QuickAttack + SPR Up Rules
	{QuickAttack, SPRUp, QuickAttack},

	// QuickAttack + SP Turbo Rules
	{QuickAttack, SPTurbo, QuickAttack},

	// QuickAttack + Libra Rules
	{QuickAttack, Libra, QuickAttack},
}

var QuickAttackStatusBasicRules = []BasicCombinationRule{
	// QuickAttackStatus + FIT Rules
	{QuickAttackStatus, Fire, QuickAttackStatus},
	{QuickAttackStatus, Ice, QuickAttackStatus},
	{QuickAttackStatus, Lightning, QuickAttackStatus},

	// QuickAttackStatus + Restore Rules
	{QuickAttackStatus, Restore, QuickAttackStatus},

	// QuickAttackStatus + Defense Rules are not basic

	// QuickAttackStatus + Status Defense Rules
	{QuickAttackStatus, StatusDefense, QuickAttackStatus},

	// QuickAttackStatus + Absorb Magic Rules are not basic

	// QuickAttackStatus + Status Magic Rules
	{QuickAttackStatus, StatusMagic, QuickAttackStatus},

	// QuickAttackStatus + FITStatus Rules
	{QuickAttackStatus, FireStatus, QuickAttackStatus},
	{QuickAttackStatus, IceStatus, QuickAttackStatus},
	{QuickAttackStatus, LightningStatus, QuickAttackStatus},

	// QuickAttackStatus + Gravity Rules are not basic

	// QuickAttackStatus + Ultimate Rules
	{QuickAttackStatus, Ultimate, QuickAttackStatus},

	// QuickAttackStatus + Quick Attack Rules
	{QuickAttackStatus, QuickAttack, QuickAttackStatus},

	// QuickAttackStatus + Quick Attack Status Rules
	{QuickAttackStatus, QuickAttackStatus, QuickAttackStatus},

	// QuickAttackStatus + Blade Arts Rules
	{QuickAttackStatus, BladeArts, QuickAttackStatus},

	// QuickAttackStatus + Blade Arts Status Rules
	{QuickAttackStatus, BladeArtsStatus, QuickAttackStatus},

	// QuickAttack + FIT Blade Rules
	{QuickAttackStatus, FireBlade, QuickAttackStatus},
	{QuickAttackStatus, IceBlade, QuickAttackStatus},
	{QuickAttackStatus, LightningBlade, QuickAttackStatus},

	// QuickAttackStatus + Absorb Blade Rules are not basic

	// QuickAttackStatus + Item Rules are not basic

	// QuickAttackStatus + Punch Rules
	{QuickAttackStatus, Punch, QuickAttackStatus},

	// QuickAttackStatus + HP Up Rules
	{QuickAttackStatus, HPUp, StatusDefense},

	// QuickAttackStatus + MP Up Rules
	{QuickAttackStatus, MPUp, StatusMagic},

	// QuickAttackStatus + AP Up Rules
	{QuickAttackStatus, APUp, QuickAttack},

	// QuickAttackStatus + ATK Up Rules
	{QuickAttackStatus, ATKUp, QuickAttackStatus},

	// QuickAttackStatus + VIT Up Rules
	{QuickAttackStatus, VITUp, StatusDefense},

	// QuickAttackStatus + MAG Up Rules
	{QuickAttackStatus, MAGUp, StatusMagic},

	// QuickAttackStatus + SPR Up Rules
	{QuickAttackStatus, SPRUp, StatusDefense},

	// QuickAttackStatus + SP Turbo Rules
	{QuickAttackStatus, SPTurbo, QuickAttackStatus},

	// QuickAttackStatus + Libra Rules
	{QuickAttackStatus, Libra, QuickAttackStatus},
}

var BladeArtsBasicRules = []BasicCombinationRule{
	// BladeArts + FIT Rules
	{BladeArts, Fire, FireBlade},
	{BladeArts, Ice, IceBlade},
	{BladeArts, Lightning, LightningBlade},

	// BladeArts + Restore Rules
	{BladeArts, Restore, BladeArts},

	// BladeArts + Defense Rules are not basic

	// BladeArts + Status Defense Rules
	{BladeArts, StatusDefense, BladeArts},

	// BladeArts + Absorb Magic Rules are not basic

	// BladeArts + Status Magic Rules
	{BladeArts, StatusMagic, BladeArtsStatus},

	// BladeArts + FITStatus Rules
	{BladeArts, FireStatus, FireBlade},
	{BladeArts, IceStatus, IceBlade},
	{BladeArts, LightningStatus, LightningBlade},

	// BladeArts + Gravity Rules
	{BladeArts, Gravity, BladeArts},

	// BladeArts + Ultimate Rules
	{BladeArts, Ultimate, BladeArts},

	// BladeArts + Quick Attack Rules
	{BladeArts, QuickAttack, BladeArts},

	// BladeArts + Quick Attack Status Rules
	{BladeArts, QuickAttackStatus, QuickAttackStatus},

	// BladeArts + Blade Arts Rules
	{BladeArts, BladeArts, BladeArts},

	// BladeArts + Blade Arts Status Rules
	{BladeArts, BladeArtsStatus, BladeArtsStatus},

	// BladeArts + FIT Blade Rules
	{BladeArts, FireBlade, FireBlade},
	{BladeArts, IceBlade, IceBlade},
	{BladeArts, LightningBlade, LightningBlade},

	// BladeArts + Absorb Blade Rules are not basic

	// BladeArts + Item Rules are not basic

	// BladeArts + Punch Rules
	{BladeArts, Punch, BladeArts},

	// BladeArts + HP Up Rules
	{BladeArts, HPUp, BladeArts},

	// BladeArts + MP Up Rules
	{BladeArts, MPUp, BladeArts},

	// BladeArts + AP Up Rules
	{BladeArts, APUp, QuickAttack},

	// BladeArts + ATK Up Rules
	{BladeArts, ATKUp, BladeArts},

	// BladeArts + VIT Up Rules
	{BladeArts, VITUp, BladeArts},

	// BladeArts + MAG Up Rules
	{BladeArts, MAGUp, BladeArts},

	// BladeArts + SPR Up Rules
	{BladeArts, SPRUp, BladeArts},

	// BladeArts + SP Turbo Rules
	{BladeArts, SPTurbo, BladeArts},

	// BladeArts + Libra Rules
	{BladeArts, Libra, BladeArts},
}

var BladeArtsStatusBasicRules = []BasicCombinationRule{
	// BladeArtsStatus + FIT Rules
	{BladeArtsStatus, Fire, BladeArtsStatus},
	{BladeArtsStatus, Ice, BladeArtsStatus},
	{BladeArtsStatus, Lightning, BladeArtsStatus},

	// BladeArtsStatus + Restore Rules
	{BladeArtsStatus, Restore, BladeArtsStatus},

	// BladeArtsStatus + Defense Rules are not basic

	// BladeArtsStatus + Status Defense Rules
	{BladeArtsStatus, StatusDefense, BladeArtsStatus},

	// BladeArtsStatus + Absorb Magic Rules are not basic

	// BladeArtsStatus + Status Magic Rules
	{BladeArtsStatus, StatusMagic, BladeArtsStatus},

	// BladeArtsStatus + FITStatus Rules
	{BladeArtsStatus, FireStatus, BladeArtsStatus},
	{BladeArtsStatus, IceStatus, BladeArtsStatus},
	{BladeArtsStatus, LightningStatus, BladeArtsStatus},

	// BladeArtsStatus + Gravity Rules
	{BladeArtsStatus, Gravity, BladeArtsStatus},

	// BladeArtsStatus + Ultimate Rules
	{BladeArtsStatus, Ultimate, BladeArtsStatus},

	// BladeArtsStatus + Quick Attack Rules
	{BladeArtsStatus, QuickAttack, BladeArtsStatus},

	// BladeArtsStatus + Quick Attack Status Rules
	{BladeArtsStatus, QuickAttackStatus, BladeArtsStatus},

	// BladeArtsStatus + Blade Arts Rules
	{BladeArtsStatus, BladeArts, BladeArtsStatus},

	// BladeArtsStatus + Blade Arts Status Rules
	{BladeArtsStatus, BladeArtsStatus, BladeArtsStatus},

	// BladeArtsStatus + FIT Blade Rules
	{BladeArtsStatus, FireBlade, BladeArtsStatus},
	{BladeArtsStatus, IceBlade, BladeArtsStatus},
	{BladeArtsStatus, LightningBlade, BladeArtsStatus},

	// BladeArtsStatus + Absorb Blade Rules are not basic

	// BladeArtsStatus + Item Rules are not basic

	// BladeArtsStatus + Punch Rules
	{BladeArtsStatus, Punch, BladeArtsStatus},

	// BladeArtsStatus + HP Up Rules
	{BladeArtsStatus, HPUp, StatusDefense},

	// BladeArtsStatus + MP Up Rules
	{BladeArtsStatus, MPUp, StatusMagic},

	// BladeArtsStatus + AP Up Rules
	{BladeArtsStatus, APUp, QuickAttack},

	// BladeArtsStatus + ATK Up Rules
	{BladeArtsStatus, ATKUp, BladeArtsStatus},

	// BladeArtsStatus + VIT Up Rules
	{BladeArtsStatus, VITUp, StatusDefense},

	// BladeArtsStatus + MAG Up Rules
	{BladeArtsStatus, MAGUp, StatusMagic},

	// BladeArtsStatus + SPR Up Rules
	{BladeArtsStatus, SPRUp, StatusDefense},

	// BladeArtsStatus + SP Turbo Rules
	{BladeArtsStatus, SPTurbo, BladeArtsStatus},

	// BladeArtsStatus + Libra Rules
	{BladeArtsStatus, Libra, BladeArtsStatus},
}

var FITBladeBasicRules = []BasicCombinationRule{
	//FireBlade + FIT Rules
	{FireBlade, Fire, FireBlade},
	{FireBlade, Ice, LightningBlade},
	{FireBlade, Lightning, IceBlade},

	//IceBlade + FIT Rules
	{IceBlade, Fire, LightningBlade},
	{IceBlade, Ice, IceBlade},
	{IceBlade, Lightning, FireBlade},

	//LightningBlade + FIT Rules
	{LightningBlade, Fire, IceBlade},
	{LightningBlade, Ice, FireBlade},
	{LightningBlade, Lightning, LightningBlade},

	// FITBlade + Restore Rules are not basic

	// FITBlade + Defense Rules are not basic

	// FITBlade + Status Defense Rules are not basic

	// FITBlade + Absorb Magic Rules are not basic

	// FITBlade + Status Magic Rules
	{FireBlade, StatusMagic, FireBlade},
	{IceBlade, StatusMagic, IceBlade},
	{LightningBlade, StatusMagic, LightningBlade},

	// FireBlade + FITStatus Rules
	{FireBlade, FireStatus, FireStatus},
	{FireBlade, IceStatus, IceStatus},
	{FireBlade, LightningStatus, LightningStatus},

	// IceBlade + FITStatus Rules
	{IceBlade, FireStatus, FireStatus},
	{IceBlade, IceStatus, IceStatus},
	{IceBlade, LightningStatus, LightningStatus},

	// LightningBlade + FITStatus Rules
	{LightningBlade, FireStatus, FireStatus},
	{LightningBlade, IceStatus, IceStatus},
	{LightningBlade, LightningStatus, LightningStatus},

	// FITBlade + Gravity Rules are not basic

	// FITBlade + Ultimate Rules are not basic

	// FITBlade + Quick Attack Rules are not basic

	// FITBlade + Quick Attack Status Rules are not basic

	// FITBlade + Blade Arts Rules are not basic

	// FITBlade + Blade Arts Status Rules are not basic

	// FireBlade + FIT Blade Rules
	{FireBlade, FireBlade, FireBlade},
	{FireBlade, IceBlade, LightningBlade},
	{FireBlade, LightningBlade, IceBlade},

	// IceBlade + FIT Blade Rules
	{IceBlade, FireBlade, LightningBlade},
	{IceBlade, IceBlade, IceBlade},
	{IceBlade, LightningBlade, FireBlade},

	// LightningBlade + FIT Blade Rules
	{LightningBlade, FireBlade, IceBlade},
	{LightningBlade, IceBlade, FireBlade},
	{LightningBlade, LightningBlade, LightningBlade},

	// FITBlade + Absorb Blade Rules are not basic

	// FITBlade + Item Rules are not basic

	// FITBlade + Punch Rules are not basic

	// FITBlade + HP Up Rules are not basic

	// FITBlade + MP Up Rules are not basic

	// FITBlade + AP Up Rules
	{FireBlade, APUp, QuickAttack},
	{IceBlade, APUp, QuickAttack},
	{LightningBlade, APUp, QuickAttack},

	// FITBlade + ATK Up Rules are not basic

	// FITBlade + VIT Up Rules are not basic

	// FITBlade + MAG Up Rules are not basic

	// FITBlade + SPR Up Rules are not basic

	// FITBlade + SP Turbo Rules are not basic

	// FITBlade + Libra Rules are not basic
}

var AbsorbBladeBasicRules = []BasicCombinationRule{
	// AbsorbBlade + FIT Rules are not basic

	// AbsorbBlade + Restore Rules are not basic

	// AbsorbBlade + Defense Rules are not basic

	// AbsorbBlade + Status Defense Rules are not basic

	// AbsorbBlade + Absorb Magic Rules are not basic

	// AbsorbBlade + Status Magic Rules are not basic

	// AbsorbBlade + FITStatus Rules are not basic

	// AbsorbBlade + Gravity Rules are not basic

	// AbsorbBlade + Ultimate Rules are not basic

	// AbsorbBlade + Quick Attack Rules are not basic

	// AbsorbBlade + Quick Attack Status Rules are not basic

	// AbsorbBlade + Blade Arts Status Rules are not basic

	// AbsorbBlade + FIT Blade Rules are not basic

	// AbsorbBlade + Absorb Blade Rules
	{AbsorbBlade, AbsorbBlade, AbsorbBlade},

	// AbsorbBlade + Item Rules are not basic

	// AbsorbBlade + Punch Rules are not basic

	// AbsorbBlade + HP Up Rules are not basic

	// AbsorbBlade + MP Up Rules are not basic

	// AbsorbBlade + AP Up Rules are not basic

	// AbsorbBlade + ATK Up Rules are not basic

	// AbsorbBlade + VIT Up Rules are not basic

	// AbsorbBlade + MAG Up Rules
	{AbsorbBlade, MAGUp, AbsorbMagic},

	// AbsorbBlade + SPR Up Rules
	{AbsorbBlade, SPRUp, AbsorbMagic},

	// BladeArts + SP Turbo Rules are not basic

	// BladeArts + Libra Rules are not basic
}

var ItemBasicRules = []BasicCombinationRule{
	// Item + FIT Rules
	{Item, Fire, Item},
	{Item, Ice, Item},
	{Item, Lightning, Item},

	// Item + Restore Rules
	{Item, Restore, Item},

	// Item + Defense Rules
	{Item, Defense, Item},

	// Item + Status Defense Rules
	{Item, StatusDefense, Item},

	// Item + Absorb Magic Rules
	{Item, AbsorbMagic, Item},

	// Item + Status Magic Rules
	{Item, StatusMagic, Item},

	// Item + FITStatus Rules
	{Item, FireStatus, Item},
	{Item, IceStatus, Item},
	{Item, LightningStatus, Item},

	// Item + Gravity Rules
	{Item, Gravity, Item},

	// Item + Ultimate Rules
	{Item, Ultimate, Item},

	// Item + Quick Attack Rules
	{Item, QuickAttack, Item},

	// Item + Quick Attack Status Rules
	{Item, QuickAttackStatus, Item},

	// Item + Blade Arts Rules
	{Item, BladeArts, Item},

	// Item + Blade Arts Status Rules
	{Item, BladeArtsStatus, Item},

	// Item + FIT Blade Rules
	{Item, FireBlade, Item},
	{Item, IceBlade, Item},
	{Item, LightningBlade, Item},

	// Item + Absorb Blade Rules
	{Item, AbsorbBlade, Item},

	// Item + Item Rules
	{Item, Item, Item},

	// Item + Punch Rules
	{Item, Punch, Item},

	// Item + HP Up Rules
	{Item, HPUp, Item},

	// Item + MP Up Rules
	{Item, MPUp, Item},

	// Item + AP Up Rules
	{Item, APUp, Item},

	// Item + ATK Up Rules
	{Item, ATKUp, Item},

	// Item + VIT Up Rules
	{Item, VITUp, Item},

	// Item + MAG Up Rules
	{Item, MAGUp, Item},

	// Item + SPR Up Rules
	{Item, SPRUp, Item},

	// Item + SP Turbo Rules
	{Item, SPTurbo, Item},

	// Item + Libra Rules
	{Item, Libra, Item},
}

var PunchBasicRules = []BasicCombinationRule{
	// Punch + FIT Rules
	{Punch, Fire, Fire},
	{Punch, Ice, Ice},
	{Punch, Lightning, Lightning},

	// Punch + Restore Rules
	{Punch, Restore, Restore},

	// Punch + Defense Rules
	{Punch, Defense, Defense},

	// Punch + Status Defense Rules
	{Punch, StatusDefense, StatusDefense},

	// Punch + Absorb Magic Rules
	{Punch, AbsorbMagic, AbsorbMagic},

	// Punch + Status Magic Rules
	{Punch, StatusMagic, StatusMagic},

	// Punch + FITStatus Rules
	{Punch, FireStatus, FireStatus},
	{Punch, IceStatus, IceStatus},
	{Punch, LightningStatus, LightningStatus},

	// Punch + Gravity Rules
	{Punch, Gravity, Gravity},

	// Punch + Ultimate Rules
	{Punch, Ultimate, Punch},

	// Punch + Quick Attack Rules
	{Punch, QuickAttack, QuickAttack},

	// Punch + Quick Attack Status Rules
	{Punch, QuickAttackStatus, QuickAttackStatus},

	// Punch + Blade Arts Rules
	{Punch, BladeArts, BladeArts},

	// Punch + Blade Arts Status Rules
	{Punch, BladeArtsStatus, BladeArtsStatus},

	// Punch + FIT Blade Rules are not basic

	// Punch + Absorb Blade Rules are not basic

	// Punch + Item Rules
	{Punch, Item, Item},

	// Punch + Punch Rules
	{Punch, Punch, Punch},

	// Punch + HP Up Rules
	{Punch, HPUp, HPUp},

	// Punch + MP Up Rules
	{Punch, MPUp, MPUp},

	// Punch + AP Up Rules
	{Punch, APUp, APUp},

	// Punch + ATK Up Rules
	{Punch, ATKUp, ATKUp},

	// Punch + VIT Up Rules
	{Punch, VITUp, VITUp},

	// Punch + MAG Up Rules
	{Punch, MAGUp, MAGUp},

	// Punch + SPR Up Rules
	{Punch, SPRUp, SPRUp},

	// Punch + SP Turbo Rules
	{Punch, SPTurbo, SPTurbo},

	// Punch + Libra Rules
	{Punch, Libra, Punch},
}

var HPUpBasicRules = []BasicCombinationRule{
	// HPUp + FIT Rules
	{HPUp, Fire, HPUp},
	{HPUp, Ice, HPUp},
	{HPUp, Lightning, HPUp},

	// HPUp + Restore Rules
	{HPUp, Restore, HPUp},

	// HPUp + Defense Rules are not basic

	// HPUp + Status Defense Rules
	{HPUp, StatusDefense, HPUp},

	// HPUp + Absorb Magic Rules
	{HPUp, AbsorbMagic, HPUp},

	// HPUp + Status Magic Rules
	{HPUp, StatusMagic, HPUp},

	// HPUp + FITStatus Rules
	{HPUp, FireStatus, HPUp},
	{HPUp, IceStatus, HPUp},
	{HPUp, LightningStatus, HPUp},

	// HPUp + Gravity Rules are not basic

	// HPUp + Ultimate Rules
	{HPUp, Ultimate, HPUp},

	// HPUp + Quick Attack Rules
	{HPUp, QuickAttack, HPUp},

	// HPUp + Quick Attack Status Rules
	{HPUp, QuickAttackStatus, HPUp},

	// HPUp + Blade Arts Rules
	{HPUp, BladeArts, HPUp},

	// HPUp + Blade Arts Status Rules
	{HPUp, BladeArtsStatus, HPUp},

	// HPUp + FIT Blade Rules
	{HPUp, FireBlade, HPUp},
	{HPUp, IceBlade, HPUp},
	{HPUp, LightningBlade, HPUp},

	// HPUp + Absorb Blade Rules
	{HPUp, AbsorbBlade, HPUp},

	// HPUp + Item Rules are not basic

	// HPUp + Punch Rules
	{HPUp, Punch, HPUp},

	// HPUp + HP Up Rules
	{HPUp, HPUp, HPUp},

	// HPUp + MP Up Rules
	{HPUp, MPUp, HPUp},

	// HPUp + AP Up Rules
	{HPUp, APUp, HPUp},

	// HPUp + ATK Up Rules
	{HPUp, ATKUp, HPUp},

	// HPUp + VIT Up Rules
	{HPUp, VITUp, HPUp},

	// HPUp + MAG Up Rules
	{HPUp, MAGUp, HPUp},

	// HPUp + SPR Up Rules
	{HPUp, SPRUp, HPUp},

	// HPUp + SP Turbo Rules
	{HPUp, SPTurbo, HPUp},

	// HPUp + Libra Rules
	{HPUp, Libra, HPUp},
}

var MPUpBasicRules = []BasicCombinationRule{
	// MPUp + FIT Rules
	{MPUp, Fire, MPUp},
	{MPUp, Ice, MPUp},
	{MPUp, Lightning, MPUp},

	// MPUp + Restore Rules
	{MPUp, Restore, MPUp},

	// MPUp + Defense Rules are not basic

	// MPUp + Status Defense Rules
	{MPUp, StatusDefense, MPUp},

	// MPUp + Absorb Magic Rules
	{MPUp, AbsorbMagic, MPUp},

	// MPUp + Status Magic Rules
	{MPUp, StatusMagic, MPUp},

	// MPUp + FITStatus Rules
	{MPUp, FireStatus, MPUp},
	{MPUp, IceStatus, MPUp},
	{MPUp, LightningStatus, MPUp},

	// MPUp + Gravity Rules are not basic

	// MPUp + Ultimate Rules
	{MPUp, Ultimate, MPUp},

	// MPUp + Quick Attack Rules
	{MPUp, QuickAttack, MPUp},

	// MPUp + Quick Attack Status Rules
	{MPUp, QuickAttackStatus, MPUp},

	// MPUp + Blade Arts Rules
	{MPUp, BladeArts, MPUp},

	// MPUp + Blade Arts Status Rules
	{MPUp, BladeArtsStatus, MPUp},

	// MPUp + FIT Blade Rules
	{MPUp, FireBlade, MPUp},
	{MPUp, IceBlade, MPUp},
	{MPUp, LightningBlade, MPUp},

	// MPUp + Absorb Blade Rules
	{MPUp, AbsorbBlade, MPUp},

	// MPUp + Item Rules are not basic

	// MPUp + Punch Rules
	{MPUp, Punch, MPUp},

	// MPUp + HP Up Rules
	{MPUp, HPUp, MPUp},

	// MPUp + MP Up Rules
	{MPUp, MPUp, MPUp},

	// MPUp + AP Up Rules
	{MPUp, APUp, MPUp},

	// MPUp + ATK Up Rules
	{MPUp, ATKUp, MPUp},

	// MPUp + VIT Up Rules
	{MPUp, VITUp, MPUp},

	// MPUp + MAG Up Rules
	{MPUp, MAGUp, MPUp},

	// MPUp + SPR Up Rules
	{MPUp, SPRUp, MPUp},

	// MPUp + SP Turbo Rules
	{MPUp, SPTurbo, MPUp},

	// MPUp + Libra Rules
	{MPUp, Libra, MPUp},
}

var APUpBasicRules = []BasicCombinationRule{
	// APUp + FIT Rules
	{APUp, Fire, APUp},
	{APUp, Ice, APUp},
	{APUp, Lightning, APUp},

	// APUp + Restore Rules
	{APUp, Restore, APUp},

	// APUp + Defense Rules are not basic

	// APUp + Status Defense Rules
	{APUp, StatusDefense, APUp},

	// APUp + Absorb Magic Rules
	{APUp, AbsorbMagic, APUp},

	// APUp + Status Magic Rules
	{APUp, StatusMagic, APUp},

	// APUp + FITStatus Rules
	{APUp, FireStatus, APUp},
	{APUp, IceStatus, APUp},
	{APUp, LightningStatus, APUp},

	// APUp + Gravity Rules are not basic

	// APUp + Ultimate Rules
	{APUp, Ultimate, APUp},

	// APUp + Quick Attack Rules
	{APUp, QuickAttack, APUp},

	// APUp + Quick Attack Status Rules
	{APUp, QuickAttackStatus, APUp},

	// APUp + Blade Arts Rules
	{APUp, BladeArts, APUp},

	// APUp + Blade Arts Status Rules
	{APUp, BladeArtsStatus, APUp},

	// APUp + FIT Blade Rules
	{APUp, FireBlade, APUp},
	{APUp, IceBlade, APUp},
	{APUp, LightningBlade, APUp},

	// APUp + Absorb Blade Rules
	{APUp, AbsorbBlade, APUp},

	// APUp + Item Rules are not basic

	// APUp + Punch Rules
	{APUp, Punch, APUp},

	// APUp + HP Up Rules
	{APUp, HPUp, APUp},

	// APUp + MP Up Rules
	{APUp, MPUp, APUp},

	// APUp + AP Up Rules
	{APUp, APUp, APUp},

	// APUp + ATK Up Rules
	{APUp, ATKUp, APUp},

	// APUp + VIT Up Rules
	{APUp, VITUp, APUp},

	// APUp + MAG Up Rules
	{APUp, MAGUp, APUp},

	// APUp + SPR Up Rules
	{APUp, SPRUp, APUp},

	// APUp + SP Turbo Rules
	{APUp, SPTurbo, APUp},

	// APUp + Libra Rules
	{APUp, Libra, APUp},
}

var ATKUpBasicRules = []BasicCombinationRule{
	// ATKUp + FIT Rules
	{ATKUp, Fire, ATKUp},
	{ATKUp, Ice, ATKUp},
	{ATKUp, Lightning, ATKUp},

	// ATKUp + Restore Rules
	{ATKUp, Restore, ATKUp},

	// ATKUp + Defense Rules are not basic

	// ATKUp + Status Defense Rules
	{ATKUp, StatusDefense, ATKUp},

	// ATKUp + Absorb Magic Rules
	{ATKUp, AbsorbMagic, ATKUp},

	// ATKUp + Status Magic Rules
	{ATKUp, StatusMagic, ATKUp},

	// ATKUp + FITStatus Rules
	{ATKUp, FireStatus, ATKUp},
	{ATKUp, IceStatus, ATKUp},
	{ATKUp, LightningStatus, ATKUp},

	// ATKUp + Gravity Rules are not basic

	// ATKUp + Ultimate Rules
	{ATKUp, Ultimate, ATKUp},

	// ATKUp + Quick Attack Rules
	{ATKUp, QuickAttack, ATKUp},

	// ATKUp + Quick Attack Status Rules
	{ATKUp, QuickAttackStatus, ATKUp},

	// ATKUp + Blade Arts Rules
	{ATKUp, BladeArts, ATKUp},

	// ATKUp + Blade Arts Status Rules
	{ATKUp, BladeArtsStatus, ATKUp},

	// ATKUp + FIT Blade Rules
	{ATKUp, FireBlade, ATKUp},
	{ATKUp, IceBlade, ATKUp},
	{ATKUp, LightningBlade, ATKUp},

	// ATKUp + Absorb Blade Rules
	{ATKUp, AbsorbBlade, ATKUp},

	// ATKUp + Item Rules are not basic

	// ATKUp + Punch Rules
	{ATKUp, Punch, ATKUp},

	// ATKUp + HP Up Rules
	{ATKUp, HPUp, ATKUp},

	// ATKUp + MP Up Rules
	{ATKUp, MPUp, ATKUp},

	// ATKUp + AP Up Rules
	{ATKUp, APUp, ATKUp},

	// ATKUp + ATK Up Rules
	{ATKUp, ATKUp, ATKUp},

	// ATKUp + VIT Up Rules
	{ATKUp, VITUp, ATKUp},

	// ATKUp + MAG Up Rules
	{ATKUp, MAGUp, ATKUp},

	// ATKUp + SPR Up Rules
	{ATKUp, SPRUp, ATKUp},

	// ATKUp + SP Turbo Rules
	{ATKUp, SPTurbo, ATKUp},

	// ATKUp + Libra Rules
	{ATKUp, Libra, ATKUp},
}

var VITUpBasicRules = []BasicCombinationRule{
	// VITUp + FIT Rules
	{VITUp, Fire, VITUp},
	{VITUp, Ice, VITUp},
	{VITUp, Lightning, VITUp},

	// VITUp + Restore Rules
	{VITUp, Restore, VITUp},

	// VITUp + Defense Rules are not basic

	// VITUp + Status Defense Rules
	{VITUp, StatusDefense, VITUp},

	// VITUp + Absorb Magic Rules
	{VITUp, AbsorbMagic, VITUp},

	// VITUp + Status Magic Rules
	{VITUp, StatusMagic, VITUp},

	// VITUp + FITStatus Rules
	{VITUp, FireStatus, VITUp},
	{VITUp, IceStatus, VITUp},
	{VITUp, LightningStatus, VITUp},

	// VITUp + Gravity Rules are not basic

	// VITUp + Ultimate Rules
	{VITUp, Ultimate, VITUp},

	// VITUp + Quick Attack Rules
	{VITUp, QuickAttack, VITUp},

	// VITUp + Quick Attack Status Rules
	{VITUp, QuickAttackStatus, VITUp},

	// VITUp + Blade Arts Rules
	{VITUp, BladeArts, VITUp},

	// VITUp + Blade Arts Status Rules
	{VITUp, BladeArtsStatus, VITUp},

	// VITUp + FIT Blade Rules
	{VITUp, FireBlade, VITUp},
	{VITUp, IceBlade, VITUp},
	{VITUp, LightningBlade, VITUp},

	// VITUp + Absorb Blade Rules
	{VITUp, AbsorbBlade, VITUp},

	// VITUp + Item Rules are not basic

	// VITUp + Punch Rules
	{VITUp, Punch, VITUp},

	// VITUp + HP Up Rules
	{VITUp, HPUp, VITUp},

	// VITUp + MP Up Rules
	{VITUp, MPUp, VITUp},

	// VITUp + AP Up Rules
	{VITUp, APUp, VITUp},

	// VITUp + ATK Up Rules
	{VITUp, ATKUp, VITUp},

	// VITUp + VIT Up Rules
	{VITUp, VITUp, VITUp},

	// VITUp + MAG Up Rules
	{VITUp, MAGUp, VITUp},

	// VITUp + SPR Up Rules
	{VITUp, SPRUp, VITUp},

	// VITUp + SP Turbo Rules
	{VITUp, SPTurbo, VITUp},

	// VITUp + Libra Rules
	{VITUp, Libra, VITUp},
}

var MAGUpBasicRules = []BasicCombinationRule{
	// MAGUp + FIT Rules
	{MAGUp, Fire, MAGUp},
	{MAGUp, Ice, MAGUp},
	{MAGUp, Lightning, MAGUp},

	// MAGUp + Restore Rules
	{MAGUp, Restore, MAGUp},

	// MAGUp + Defense Rules are not basic

	// MAGUp + Status Defense Rules
	{MAGUp, StatusDefense, MAGUp},

	// MAGUp + Absorb Magic Rules
	{MAGUp, AbsorbMagic, MAGUp},

	// MAGUp + Status Magic Rules
	{MAGUp, StatusMagic, MAGUp},

	// MAGUp + FITStatus Rules
	{MAGUp, FireStatus, MAGUp},
	{MAGUp, IceStatus, MAGUp},
	{MAGUp, LightningStatus, MAGUp},

	// MAGUp + Gravity Rules are not basic

	// MAGUp + Ultimate Rules
	{MAGUp, Ultimate, MAGUp},

	// MAGUp + Quick Attack Rules
	{MAGUp, QuickAttack, MAGUp},

	// MAGUp + Quick Attack Status Rules
	{MAGUp, QuickAttackStatus, MAGUp},

	// MAGUp + Blade Arts Rules
	{MAGUp, BladeArts, MAGUp},

	// MAGUp + Blade Arts Status Rules
	{MAGUp, BladeArtsStatus, MAGUp},

	// MAGUp + FIT Blade Rules
	{MAGUp, FireBlade, MAGUp},
	{MAGUp, IceBlade, MAGUp},
	{MAGUp, LightningBlade, MAGUp},

	// MAGUp + Absorb Blade Rules
	{MAGUp, AbsorbBlade, MAGUp},

	// MAGUp + Item Rules are not basic

	// MAGUp + Punch Rules
	{MAGUp, Punch, MAGUp},

	// MAGUp + HP Up Rules
	{MAGUp, HPUp, MAGUp},

	// MAGUp + MP Up Rules
	{MAGUp, MPUp, MAGUp},

	// MAGUp + AP Up Rules
	{MAGUp, APUp, MAGUp},

	// MAGUp + ATK Up Rules
	{MAGUp, ATKUp, MAGUp},

	// MAGUp + VIT Up Rules
	{MAGUp, VITUp, MAGUp},

	// MAGUp + MAG Up Rules
	{MAGUp, MAGUp, MAGUp},

	// MAGUp + SPR Up Rules
	{MAGUp, SPRUp, MAGUp},

	// MAGUp + SP Turbo Rules
	{MAGUp, SPTurbo, MAGUp},

	// MAGUp + Libra Rules
	{MAGUp, Libra, MAGUp},
}

var SPRUpBasicRules = []BasicCombinationRule{
	// SPRUp + FIT Rules
	{SPRUp, Fire, SPRUp},
	{SPRUp, Ice, SPRUp},
	{SPRUp, Lightning, SPRUp},

	// SPRUp + Restore Rules
	{SPRUp, Restore, SPRUp},

	// SPRUp + Defense Rules are not basic

	// SPRUp + Status Defense Rules
	{SPRUp, StatusDefense, SPRUp},

	// SPRUp + Absorb Magic Rules
	{SPRUp, AbsorbMagic, SPRUp},

	// SPRUp + Status Magic Rules
	{SPRUp, StatusMagic, SPRUp},

	// SPRUp + FITStatus Rules
	{SPRUp, FireStatus, SPRUp},
	{SPRUp, IceStatus, SPRUp},
	{SPRUp, LightningStatus, SPRUp},

	// SPRUp + Gravity Rules are not basic

	// SPRUp + Ultimate Rules
	{SPRUp, Ultimate, SPRUp},

	// SPRUp + Quick Attack Rules
	{SPRUp, QuickAttack, SPRUp},

	// SPRUp + Quick Attack Status Rules
	{SPRUp, QuickAttackStatus, SPRUp},

	// SPRUp + Blade Arts Rules
	{SPRUp, BladeArts, SPRUp},

	// SPRUp + Blade Arts Status Rules
	{SPRUp, BladeArtsStatus, SPRUp},

	// SPRUp + FIT Blade Rules
	{SPRUp, FireBlade, SPRUp},
	{SPRUp, IceBlade, SPRUp},
	{SPRUp, LightningBlade, SPRUp},

	// SPRUp + Absorb Blade Rules
	{SPRUp, AbsorbBlade, SPRUp},

	// SPRUp + Item Rules are not basic

	// SPRUp + Punch Rules
	{SPRUp, Punch, SPRUp},

	// SPRUp + HP Up Rules
	{SPRUp, HPUp, SPRUp},

	// SPRUp + MP Up Rules
	{SPRUp, MPUp, SPRUp},

	// SPRUp + AP Up Rules
	{SPRUp, APUp, SPRUp},

	// SPRUp + ATK Up Rules
	{SPRUp, ATKUp, SPRUp},

	// SPRUp + VIT Up Rules
	{SPRUp, VITUp, SPRUp},

	// SPRUp + MAG Up Rules
	{SPRUp, MAGUp, SPRUp},

	// SPRUp + SPR Up Rules
	{SPRUp, SPRUp, SPRUp},

	// SPRUp + SP Turbo Rules
	{SPRUp, SPTurbo, SPRUp},

	// SPRUp + Libra Rules
	{SPRUp, Libra, SPRUp},
}

var SPTurboBasicRules = []BasicCombinationRule{
	// SPTurbo + FIT Rules
	{SPTurbo, Fire, SPTurbo},
	{SPTurbo, Ice, SPTurbo},
	{SPTurbo, Lightning, SPTurbo},

	// SPTurbo + Restore Rules
	{SPTurbo, Restore, SPTurbo},

	// SPTurbo + Defense Rules are not basic

	// SPTurbo + Status Defense Rules
	{SPTurbo, StatusDefense, SPTurbo},

	// SPTurbo + Absorb Magic Rules
	{SPTurbo, AbsorbMagic, SPTurbo},

	// SPTurbo + Status Magic Rules
	{SPTurbo, StatusMagic, SPTurbo},

	// SPTurbo + FITStatus Rules
	{SPTurbo, FireStatus, SPTurbo},
	{SPTurbo, IceStatus, SPTurbo},
	{SPTurbo, LightningStatus, SPTurbo},

	// SPTurbo + Gravity Rules are not basic

	// SPTurbo + Ultimate Rules
	{SPTurbo, Ultimate, SPTurbo},

	// SPTurbo + Quick Attack Rules
	{SPTurbo, QuickAttack, SPTurbo},

	// SPTurbo + Quick Attack Status Rules
	{SPTurbo, QuickAttackStatus, SPTurbo},

	// SPTurbo + Blade Arts Rules
	{SPTurbo, BladeArts, SPTurbo},

	// SPTurbo + Blade Arts Status Rules
	{SPTurbo, BladeArtsStatus, SPTurbo},

	// SPTurbo + FIT Blade Rules
	{SPTurbo, FireBlade, SPTurbo},
	{SPTurbo, IceBlade, SPTurbo},
	{SPTurbo, LightningBlade, SPTurbo},

	// SPTurbo + Absorb Blade Rules
	{SPTurbo, AbsorbBlade, SPTurbo},

	// SPTurbo + Item Rules are not basic

	// SPTurbo + Punch Rules
	{SPTurbo, Punch, SPTurbo},

	// SPTurbo + HP Up Rules
	{SPTurbo, HPUp, SPTurbo},

	// SPTurbo + MP Up Rules
	{SPTurbo, MPUp, SPTurbo},

	// SPTurbo + AP Up Rules
	{SPTurbo, APUp, SPTurbo},

	// SPTurbo + ATK Up Rules
	{SPTurbo, ATKUp, SPTurbo},

	// SPTurbo + VIT Up Rules
	{SPTurbo, VITUp, SPTurbo},

	// SPTurbo + MAG Up Rules
	{SPTurbo, MAGUp, SPTurbo},

	// SPTurbo + SPR Up Rules
	{SPTurbo, SPRUp, SPTurbo},

	// SPTurbo + SP Turbo Rules
	{SPTurbo, SPTurbo, SPTurbo},

	// SPTurbo + Libra Rules
	{SPTurbo, Libra, SPTurbo},
}

var LibraBasicRules = []BasicCombinationRule{
	// Libra + FIT Rules
	{Libra, Fire, Fire},
	{Libra, Ice, Ice},
	{Libra, Lightning, Lightning},

	// Libra + Restore Rules
	{Libra, Restore, Restore},

	// Libra + Defense Rules
	{Libra, Defense, Defense},

	// Libra + Status Defense Rules
	{Libra, StatusDefense, StatusDefense},

	// Libra + Absorb Magic Rules
	{Libra, AbsorbMagic, AbsorbMagic},

	// Libra + Status Magic Rules
	{Libra, StatusMagic, StatusMagic},

	// Libra + FITStatus Rules
	{Libra, FireStatus, FireStatus},
	{Libra, IceStatus, IceStatus},
	{Libra, LightningStatus, IceStatus},

	// Libra + Gravity Rules
	{Libra, Gravity, Gravity},

	// Libra + Ultimate Rules
	{Libra, Ultimate, Ultimate},

	// Libra + Quick Attack Rules
	{Libra, QuickAttack, QuickAttack},

	// Libra + Quick Attack Status Rules
	{Libra, QuickAttackStatus, QuickAttackStatus},

	// Libra + Blade Arts Rules
	{Libra, BladeArts, BladeArts},

	// Libra + Blade Arts Status Rules
	{Libra, BladeArtsStatus, BladeArtsStatus},

	// Libra + FIT Blade Rules are not basic

	// Libra + Absorb Blade Rules
	{Libra, AbsorbBlade, AbsorbBlade},

	// Libra + Item Rules
	{Libra, Item, Item},

	// Libra + Punch Rules
	{Libra, Punch, Punch},

	// Libra + HP Up Rules
	{Libra, HPUp, HPUp},

	// Libra + MP Up Rules
	{Libra, MPUp, MPUp},

	// Libra + AP Up Rules
	{Libra, APUp, APUp},

	// Libra + ATK Up Rules
	{Libra, ATKUp, ATKUp},

	// Libra + VIT Up Rules
	{Libra, VITUp, VITUp},

	// Libra + MAG Up Rules
	{Libra, MAGUp, MAGUp},

	// Libra + SPR Up Rules
	{Libra, SPRUp, SPRUp},

	// Libra + SP Turbo Rules
	{Libra, SPTurbo, SPTurbo},

	// Libra + Libra Rules
	{Libra, Libra, Libra},
}

// rules for Dash, Dualcast, Fullcure, DMW are pending

type MateriaFusionService interface {
	// GetAllMateria() ([]Materia, error)
	// GetAllRules() ([]Rule, error)
}

var BasicRuleMap = map[MateriaType][]BasicCombinationRule{
	Fire:              FITBasicRules,
	Ice:               FITBasicRules,
	Lightning:         FITBasicRules,
	Restore:           RestoreBasicRules,
	Defense:           DefenseBasicRules,
	StatusDefense:     StatusDefenseBasicRules,
	AbsorbMagic:       AbsorbMagicBasicRules,
	StatusMagic:       StatusMagicBasicRules,
	FireStatus:        FITStatusBasicRules,
	IceStatus:         FITStatusBasicRules,
	LightningStatus:   FITStatusBasicRules,
	Gravity:           GravityBasicRules,
	Ultimate:          UltimateBasicRules,
	QuickAttack:       QuickAttackBasicRules,
	QuickAttackStatus: QuickAttackStatusBasicRules,
	BladeArts:         BladeArtsBasicRules,
	BladeArtsStatus:   BladeArtsStatusBasicRules,
	FireBlade:         FITBladeBasicRules,
	IceBlade:          FITBladeBasicRules,
	LightningBlade:    FITBladeBasicRules,
	AbsorbBlade:       AbsorbBladeBasicRules,
	Item:              ItemBasicRules,
	Punch:             PunchBasicRules,
	HPUp:              HPUpBasicRules,
	MPUp:              MPUpBasicRules,
	APUp:              APUpBasicRules,
	ATKUp:             ATKUpBasicRules,
	VITUp:             VITUpBasicRules,
	MAGUp:             MAGUpBasicRules,
	SPRUp:             SPRUpBasicRules,
	SPTurbo:           SPTurboBasicRules,
	Libra:             LibraBasicRules,
}
