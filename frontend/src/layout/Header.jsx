import React from 'react'
import { useAuth } from '../context/AuthContext'

export default function Header() {
  const { user, logout } = useAuth()
  return (
    <header className="app-header">
      <div className="left">
        <button className="collapse-btn">☰</button>
      </div>
      <div className="right">
        <span>欢迎，{user?.username || '访客'}</span>
        <button onClick={logout} className="logout-btn">退出</button>
      </div>
    </header>
  )
}
