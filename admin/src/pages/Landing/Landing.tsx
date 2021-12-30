import { useState } from 'react';


import './Landing.css';

const Landing = () => {
  const [orgId, setOrgId] = useState<string>();
  const [password, setPassword] = useState<string>();


  const login = () => {
    

  }

  return (
    <div className="Landing">
      <h1 className="header">🏎️</h1>
      <div className="loginContainer">
        <h2 className="loginHeader">Admin Login</h2>
        <div className='loginForm'>
          <p className="organizationLabel">Organization ID</p>
          <input 
            type="text" 
            className="loginInput" 
            value={orgId}
            onChange={(e) => setOrgId(e.target.value)}/>

          <p className="passwordLabel">Password</p>
          <input
            type="password" 
            className="loginInput" 
            value={password} 
            onChange={(e) => setPassword(e.target.value)}/>
          <button className="signInButton" onClick={login}>Sign In</button>
        </div>
      </div>
    </div>
  )
}

export default Landing;