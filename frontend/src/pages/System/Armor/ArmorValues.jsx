import React, { useEffect, useState } from 'react'
import api from '../../../api/http'
import { Link } from 'react-router-dom'

export default function ArmorValues(){
  const [list, setList] = useState([])
  useEffect(()=>{ api.get('/api/system/armor').then(r=>{ if(r.ok) setList(r.data || []) }) },[])
  return (
    <div className="content">
      <h2>防具数值</h2>
      {list.map(a=>(
        <div key={a.id} className="card">
          <h3>{a.name}</h3>
          <ul>
            <li>耐久度: {a.durability} <Link to={`/system/config-editor?category=armor&id=${a.id}&field=durability`}>编辑</Link></li>
            <li>重量: {a.weight} <Link to={`/system/config-editor?category=armor&id=${a.id}&field=weight`}>编辑</Link></li>
            <li>维修概率: {a.repair_rate} <Link to={`/system/config-editor?category=armor&id=${a.id}&field=repair_rate`}>编辑</Link></li>
          </ul>
        </div>
      ))}
    </div>
  )
}
