import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import AppRoutes from './AppRoutes'

import { Container, Nav, Navbar, NavDropdown } from 'react-bootstrap'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
        <Navbar expand="lg" className="bg-body-tertiary">
          <Container>
            <Navbar.Brand>Ikea Assignment</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="me-auto">
                <Nav.Link href="#products">All Products</Nav.Link>
                <Nav.Link href="#products/create">Create new Product</Nav.Link>
              </Nav>
            </Navbar.Collapse>
          </Container>
        </Navbar>
        <main>
            <AppRoutes />
        </main>
    </>
  )
}

export default App
