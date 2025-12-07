export const menuTree = [
  {
    id: "monster",
    label: "怪物数值配置",
    children: [
      { id: "monster_hp", label: "血量" },
      { id: "monster_def", label: "防具" },
      { id: "monster_head", label: "头盔" },
      { id: "monster_weapon", label: "武器" },
      { id: "monster_bullet", label: "子弹" },
      { id: "monster_backpack", label: "背包" },
      { id: "monster_melee_range", label: "近战攻击距离" },
      { id: "monster_melee_speed", label: "近战攻击速度" }
    ]
  },

  {
    id: "weapon",
    label: "武器数值配置",
    children: [
      { id: "wp_name", label: "枪械名称" },
      { id: "wp_range", label: "射程" },
      { id: "wp_bullet", label: "子弹" },
      { id: "wp_capacity", label: "弹夹容量" },
      { id: "wp_reload", label: "换弹时间" },
      { id: "wp_durability", label: "耐久度" },
      { id: "wp_weight", label: "重量" }
    ]
  },

  {
    id: "armor",
    label: "防具/头盔数值配置",
    children: [
      { id: "armor_durability", label: "耐久度" },
      { id: "armor_weight", label: "重量" },
      { id: "armor_repair", label: "维护/修理耗度" }
    ]
  },

  {
    id: "bullet",
    label: "子弹数值配置",
    children: [
      { id: "bullet_level", label: "子弹等级" },
      { id: "bullet_damage", label: "子弹伤害" },
      { id: "bullet_weight", label: "重量" }
    ]
  }
]
