import React from 'react'
import Sidebar from './Sidebar'
import Header from './Header'
import './layout.css'

export default function Layout({ children }) {
  return (
    <div className="app-layout">
      <Sidebar />
      <div className="main-area">
        <Header />
        <div className="content-area">{children}</div>
      </div>
    </div>
  )
}
