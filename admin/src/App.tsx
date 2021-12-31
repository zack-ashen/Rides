import { useState } from 'react';
import { Routes, Route, Navigate } from 'react-router';

import Landing from './pages/Landing/Landing';

import './App.css';

const App = () => {
  const [token, setToken] = useState();


  const AuthenticatedContent = () => (
    <Routes>
      <Route path="/">
        <Landing />
      </Route>
    </Routes>
  )

  const UnauthenticatedContent = () => (
    <Routes>
      <Route path="/">
        <Landing />
      </Route>
      <Route path="*">
        <Navigate to="/"/>
      </Route>

    </Routes>
  )


  return <Landing />
}

export default App;
