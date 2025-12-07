export const formSchema = {
  // 怪物
  monster_hp: [{ label: "血量", prop: "hp", type: "number" }],
  monster_def: [{ label: "防御", prop: "def", type: "number" }],
  monster_head: [{ label: "头盔防御", prop: "head_def", type: "number" }],
  monster_weapon: [{ label: "武器伤害", prop: "weapon_dmg", type: "number" }],
  monster_bullet: [{ label: "子弹类型", prop: "bullet", type: "text" }],
  monster_backpack: [{ label: "背包容量", prop: "capacity", type: "number" }],
  monster_melee_range: [{ label: "近战距离", prop: "range", type: "number" }],
  monster_melee_speed: [{ label: "攻击速度", prop: "speed", type: "number" }],

  // 武器
  wp_name: [{ label: "枪械名称", prop: "name", type: "text" }],
  wp_range: [{ label: "射程", prop: "range", type: "number" }],
  wp_bullet: [{ label: "子弹类型", prop: "bullet", type: "text" }],
  wp_capacity: [{ label: "弹夹容量", prop: "capacity", type: "number" }],
  wp_reload: [{ label: "换弹时间 秒", prop: "reload_time", type: "number" }],
  wp_durability: [{ label: "耐久度", prop: "durability", type: "number" }],
  wp_weight: [{ label: "重量", prop: "weight", type: "number" }],

  // 防具
  armor_durability: [{ label: "耐久度", prop: "durability", type: "number" }],
  armor_weight: [{ label: "重量", prop: "weight", type: "number" }],
  armor_repair: [{ label: "修理耗材", prop: "repair_cost", type: "number" }],

  // 子弹
  bullet_level: [{ label: "子弹等级", prop: "level", type: "number" }],
  bullet_damage: [{ label: "伤害", prop: "damage", type: "number" }],
  bullet_weight: [{ label: "重量", prop: "weight", type: "number" }]
}
