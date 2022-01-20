import { Routes, Route } from 'react-router';

import Landing from './pages/Landing/Landing';
import Dashboard from './pages/Dashboard/Dashboard';

import './App.css';

const App = () => {
  const AuthenticatedContent = () => (
    <Routes>
      <Route path="/" element={<Dashboard />}/>
    </Routes>
  )

  const UnauthenticatedContent = () => (
    <Routes>
      <Route path="*" element={<Landing />} />
    </Routes>
  )

  const token = localStorage.getItem('token')
  return token ? <AuthenticatedContent /> : <UnauthenticatedContent />
}

export default App;
