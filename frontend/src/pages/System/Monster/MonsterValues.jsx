import React, { useEffect, useState } from 'react'
import api from '../../../api/http'
import { Link } from 'react-router-dom'

export default function MonsterValues(){
  const [list, setList] = useState([])

  useEffect(()=>{
    api.get('/api/system/monster').then(r=>{
      console.log("后端返回数据 =", r)   // ← 加这个
      if(r.ok) setList(r.data.data.list || [])
    })
  },[])

  const handleAdd = async () => {
    const name = prompt("请输入怪物名称");
    if (!name) return;

    const res = await api.post("/api/system/monster", {
      name,
      hp: 100,
      def: 10,
      helmet: 0,
      weapon: 0
    });

    if (res.ok) {
      alert("新增成功！");
      setList([...list, res.data]);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm("确认删除？")) return;

    const res = await api.delete(`/api/system/monster/${id}`);
    if (res.ok) {
      setList(list.filter(m => m.id !== id));
    }
  };

  return (
    <div className="content">
      <h2>
        怪物数值
        <button onClick={handleAdd}>新增怪物</button>
      </h2>

      {list.map(m => (
        <div key={m.id} className="card">
          <h3>
            {m.name}
            <button onClick={() => handleDelete(m.id)} style={{ marginLeft: 10 }}>
              删除
            </button>
          </h3>

          <ul>
            <li>
              血量: {m.hp}
              <Link to={`/system/config-editor?category=monster&id=${m.id}&field=hp`}>编辑</Link>
            </li>

            <li>
              防具: {m.def}
              <Link to={`/system/config-editor?category=monster&id=${m.id}&field=def`}>编辑</Link>
            </li>

            <li>
              头盔: {m.helmet}
              <Link to={`/system/config-editor?category=monster&id=${m.id}&field=helmet`}>编辑</Link>
            </li>

            <li>
              武器: {m.weapon}
              <Link to={`/system/config-editor?category=monster&id=${m.id}&field=weapon`}>编辑</Link>
            </li>
          </ul>
        </div>
      ))}
    </div>
  )
}
