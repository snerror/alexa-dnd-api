package main

// Class templates
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

// Ability templates
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

var TemplateAbilitySlice = Ability{
	Name:      "Slice",
	Attack:    9,
	Damage:    2,
	CD:        2,
	CurrentCD: 0,
}

var TemplateAbilityBite = Ability{
	Name:      "Bite",
	Attack:    6,
	Damage:    2,
	CD:        1,
	CurrentCD: 0,
}

var TemplateAbilityDeathShriek = Ability{
	Name:      "Death shriek",
	Attack:    6,
	Damage:    3,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilityClaw = Ability{
	Name:      "Claw",
	Attack:    5,
	Damage:    1,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilityStoneClaw = Ability{
	Name:      "Stone claw",
	Attack:    8,
	Damage:    2,
	CD:        0,
	CurrentCD: 0,
}

var TemplateAbilityFlameStrike = Ability{
	Name:      "Flame strike",
	Attack:    7,
	Damage:    2,
	CD:        0,
	CurrentCD: 0,
}

// Enemy templates
var TemplateEnemySkeleton = Enemy{
	Name:       EnemySkeleton,
	CurrentHp:  1,
	MaxHp:      1,
	ArmorClass: 16,
	Experience: 2,
	Abilities:  []Ability{TemplateAbilitySlice, TemplateAbilityStrike},
}

var TemplateEnemySpider = Enemy{
	Name:       EnemySpider,
	CurrentHp:  1,
	MaxHp:      1,
	ArmorClass: 15,
	Experience: 2,
	Abilities:  []Ability{TemplateAbilityStrike, TemplateAbilityBite},
}

var TemplateEnemyWraith = Enemy{
	Name:       EnemyWraith,
	CurrentHp:  2,
	MaxHp:      2,
	ArmorClass: 15,
	Experience: 3,
	Abilities:  []Ability{TemplateAbilityDeathShriek},
}

var TemplateEnemyZombie = Enemy{
	Name:       EnemyZombie,
	CurrentHp:  1,
	MaxHp:      1,
	ArmorClass: 11,
	Experience: 1,
	Abilities:  []Ability{TemplateAbilityClaw},
}

var TemplateEnemyGargoyle = Enemy{
	Name:       EnemyGargoyle,
	CurrentHp:  2,
	MaxHp:      2,
	ArmorClass: 16,
	Experience: 3,
	Abilities:  []Ability{TemplateAbilityStoneClaw},
}

var TemplateEnemyBlazingSkeleton = Enemy{
	Name:       EnemyBlazingSkeleton,
	CurrentHp:  2,
	MaxHp:      2,
	ArmorClass: 13,
	Experience: 2,
	Abilities:  []Ability{TemplateAbilityFlameStrike},
}
