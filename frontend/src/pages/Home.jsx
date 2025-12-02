import { useAuth } from '../context/AuthContext'
import { useNavigate } from 'react-router-dom'
import './Home.css'

const Home = () => {
  const { user, logout } = useAuth()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  const getRoleName = (role) => {
    const roleMap = {
      1: '超级管理员',
      2: '管理员',
      3: '普通用户',
    }
    return roleMap[role] || '未知'
  }

  return (
    <div className="home-container">
      <header className="home-header">
        <h1>管理系统</h1>
        <div className="user-info">
          <span>欢迎，{user?.username}</span>
          <span className="role-badge">{getRoleName(user?.role)}</span>
          <button onClick={handleLogout} className="logout-button">
            退出登录
          </button>
        </div>
      </header>
      <main className="home-content">
        <div className="welcome-card">
          <h2>欢迎使用管理系统</h2>
          <p>当前登录用户：{user?.username}</p>
          <p>用户角色：{getRoleName(user?.role)}</p>
          <p>用户ID：{user?.id}</p>
          <p>创建时间：{new Date(user?.created_at).toLocaleString('zh-CN')}</p>
        </div>
      </main>
    </div>
  )
}

export default Home

