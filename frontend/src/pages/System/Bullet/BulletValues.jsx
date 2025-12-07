import React, { useEffect, useState } from 'react'
import api from '../../../api/http'
import { Link } from 'react-router-dom'

export default function BulletValues(){
  const [list, setList] = useState([])
  useEffect(()=>{ api.get('/api/system/bullet').then(r=>{ if(r.ok) setList(r.data || []) }) },[])
  return (
    <div className="content">
      <h2>子弹数值</h2>
      {list.map(b=>(
        <div key={b.id} className="card">
          <h3>{b.name}</h3>
          <ul>
            <li>等级: {b.level} <Link to={`/system/config-editor?category=bullet&id=${b.id}&field=level`}>编辑</Link></li>
            <li>伤害: {b.damage} <Link to={`/system/config-editor?category=bullet&id=${b.id}&field=damage`}>编辑</Link></li>
            <li>重量: {b.weight} <Link to={`/system/config-editor?category=bullet&id=${b.id}&field=weight`}>编辑</Link></li>
          </ul>
        </div>
      ))}
    </div>
  )
}
