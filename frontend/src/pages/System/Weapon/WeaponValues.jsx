import React, { useEffect, useState } from 'react'
import api from '../../../api/http'
import { Link } from 'react-router-dom'

export default function WeaponValues(){
  const [list, setList] = useState([])
  useEffect(()=>{ api.get('/api/system/weapon').then(r=>{ if(r.ok) setList(r.data || []) }) },[])
  return (
    <div className="content">
      <h2>武器数值</h2>
      {list.map(w=>(
        <div key={w.id} className="card">
          <h3>{w.name}</h3>
          <ul>
            <li>伤害: {w.damage} <Link to={`/system/config-editor?category=weapon&id=${w.id}&field=damage`}>编辑</Link></li>
            <li>射速: {w.rate} <Link to={`/system/config-editor?category=weapon&id=${w.id}&field=rate`}>编辑</Link></li>
            <li>换弹时间: {w.reload} <Link to={`/system/config-editor?category=weapon&id=${w.id}&field=reload`}>编辑</Link></li>
            <li>耐久度: {w.durability} <Link to={`/system/config-editor?category=weapon&id=${w.id}&field=durability`}>编辑</Link></li>
            <li>重量: {w.weight} <Link to={`/system/config-editor?category=weapon&id=${w.id}&field=weight`}>编辑</Link></li>
          </ul>
        </div>
      ))}
    </div>
  )
}
