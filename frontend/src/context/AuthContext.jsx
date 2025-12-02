import { createContext, useContext, useState, useEffect } from 'react'
import axios from 'axios'

const AuthContext = createContext()

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)

  // 设置axios默认配置
  axios.defaults.baseURL = '/api'
  axios.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem('token')
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      return config
    },
    (error) => {
      return Promise.reject(error)
    }
  )

  // 响应拦截器 - 处理token刷新
  axios.interceptors.response.use(
    (response) => {
      // 检查响应头中是否有新token
      const newToken = response.headers['x-new-token']
      if (newToken) {
        localStorage.setItem('token', newToken)
      }
      return response
    },
    async (error) => {
      if (error.response?.status === 401) {
        // token过期，清除本地存储并跳转到登录页
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        setUser(null)
        window.location.href = '/login'
      }
      return Promise.reject(error)
    }
  )

  useEffect(() => {
    // 检查本地存储中的token和用户信息
    const token = localStorage.getItem('token')
    const savedUser = localStorage.getItem('user')

    if (token && savedUser) {
      try {
        setUser(JSON.parse(savedUser))
      } catch (e) {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
      }
    }
    setLoading(false)
  }, [])

  const login = async (username, password) => {
    try {
      const response = await axios.post('/auth/login', {
        username,
        password,
      })

      if (response.data.code === 200) {
        const { token, user } = response.data.data
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))
        setUser(user)
        return { success: true }
      } else {
        return { success: false, message: response.data.message }
      }
    } catch (error) {
      return {
        success: false,
        message: error.response?.data?.message || '登录失败，请稍后重试',
      }
    }
  }

  const logout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    setUser(null)
  }

  const value = {
    user,
    login,
    logout,
    loading,
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

