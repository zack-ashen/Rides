import { useState } from 'react';
import { Routes, Route, Navigate } from 'react-router';

import Landing from './pages/Landing/Landing';

import './App.css';

const App = () => {
  const [token, setToken] = useState<string>();

  const AuthenticatedContent = () => (
    <Routes>
      <Route path="/" element={<Landing setToken={setToken}/>}/>
    </Routes>
  )

  const UnauthenticatedContent = () => (
    <Routes>
      <Route path="*" element={<Landing setToken={setToken}/>} />
    </Routes>
  )


  return token ? <AuthenticatedContent /> : <UnauthenticatedContent />
}

export default App;
