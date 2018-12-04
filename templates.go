package main

var TemplateClassRogue = Player{
	Name:       "Kat",
	Class:      PlayerClassRogue,
	CurrentHp:  8,
	MaxHp:      8,
	ArmorClass: 14,
	Abilities:  []Ability{TemplateAbilityBackstab, TemplateAbilityStrike, TemplateAbilitySnipeShot},
}

var TemplateClassCleric = Player{
	Name:       "Thorgrim",
	Class:      PlayerClassCleric,
	CurrentHp:  8,
	MaxHp:      8,
	ArmorClass: 16,
	Abilities:  []Ability{TemplateAbilityLanceOfFaith, TemplateAbilityDivineFlare},
}

var TemplateClassFighter = Player{
	Name:       "Arjhan",
	Class:      PlayerClassFighter,
	CurrentHp:  10,
	MaxHp:      10,
	ArmorClass: 17,
	Abilities:  []Ability{TemplateAbilityCleave, TemplateAbilityStrike, TemplateAbilityTideOfIron},
}

var TemplateAbilityBackstab = Ability{
	Name:      "Backstab",
	Attack:    10,
	Damage:    2,
	CD:        3,
	CurrentCD: 0,
}

var TemplateAbilityStrike = Ability{
	Name:      "Strike",
	Attack:    7,
	Damage:    1,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilitySnipeShot = Ability{
	Name:      "Snipe shot",
	Attack:    10,
	Damage:    3,
	CD:        5,
	CurrentCD: 0,
}

var TemplateAbilityLanceOfFaith = Ability{
	Name:      "Lance of faith",
	Attack:    6,
	Damage:    1,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilityDivineFlare = Ability{
	Name:      "Divine flare",
	Attack:    10,
	Damage:    2,
	CD:        1,
	CurrentCD: 0,
}

var TemplateAbilityCleave = Ability{
	Name:      "Cleave",
	Attack:    6,
	Damage:    1,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilityTideOfIron = Ability{
	Name:      "Tide of iron",
	Attack:    8,
	Damage:    2,
	CD:        2,
	CurrentCD: 0,
}
