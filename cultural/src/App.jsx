import { useState } from 'react'
import indexBottom from '/index-bottom.png'
import logo from './assets/logo.png'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <main className="phone">
        <div className="screen">
          <img src={logo} alt="Logo Cultural" className="logo-img"></img>
          <div className="box">
            <h2>Bem-vindo!</h2>
              <form action="/user/select-type" method="post">
                <button type="submit" name="userType" value="common" className="btn">Sou Usuário</button>
                <button type="submit" name="userType" value="organizer" className="btn">Sou Organizador</button>
              </form>
          </div>
          <img src={indexBottom} alt="Grupo de pessoas" className="footer-img"></img>
        </div>
      </main>
    </>
  )
}

export default App
