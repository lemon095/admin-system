import React, { useEffect, useState } from 'react'
import { useSearchParams } from 'react-router-dom'
import api from '../../api/http'

export default function ConfigEditor(){
  const [search] = useSearchParams()
  const category = search.get('category')
  const id = search.get('id')
  const field = search.get('field')
  const [item, setItem] = useState({})
  const [value, setValue] = useState('')

  useEffect(()=>{
    if(!category || !id) return
    api.get(`/api/${category}/${id}`).then(r=>{ if(r.ok){ setItem(r.data||{}); setValue((r.data||{})[field] ?? '') } })
  },[category,id,field])

  const save = async ()=>{
    if(!category||!id) returnå
    const payload = {}; payload[field]= (isNaN(Number(value))?value:Number(value))
    await api.post(`/api/${category}/${id}`, { data: payload })
    alert('已保存')
  }

  if(!category) return <div className="content card">未指定分类</div>
  if(!id) return <div className="content card">未指定实体ID</div>
  if(!field) return <div className="content card">未指定字段</div>

  return (
    <div className="content">
      <h3>编辑 {category} #{id} 的 {field}</h3>
      <div className="card">
        <div style={{marginBottom:8}}>
          <label style={{display:'block',marginBottom:6}}>{field}</label>
          <input value={value} onChange={e=>setValue(e.target.value)} style={{padding:8,width:300}} />
        </div>
        <button onClick={save}>保存</button>
      </div>
    </div>
  )
}
