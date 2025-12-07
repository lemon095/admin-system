import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Layout from './layout/Layout'
import Home from './pages/Home'
import ArmorValues from './pages/System/Armor/ArmorValues'
import WeaponValues from './pages/System/Weapon/WeaponValues'
import MonsterValues from './pages/System/Monster/MonsterValues'
import BulletValues from './pages/System/Bullet/BulletValues'
import Login from './pages/Login'
import PrivateRoute from './components/PrivateRoute'

export default function AppRoutes() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={
          <PrivateRoute>
            <Layout>
              <Home />
            </Layout>
          </PrivateRoute>
        } />
        <Route path="/system/monster" element={
          <PrivateRoute>
            <Layout>
              <MonsterValues />
            </Layout>
          </PrivateRoute>
        } />
        <Route path="/system/armor" element={
          <PrivateRoute>
            <Layout>
              <ArmorValues />
            </Layout>
          </PrivateRoute>
        } />
        <Route path="/system/bullet" element={
          <PrivateRoute>
            <Layout>
              <BulletValues />
            </Layout>
          </PrivateRoute>
        } />
        <Route path="/system/weapon" element={
          <PrivateRoute>
            <Layout>
              <WeaponValues />
            </Layout>
          </PrivateRoute>
        } />
      </Routes>
    </Router>
  )
}
