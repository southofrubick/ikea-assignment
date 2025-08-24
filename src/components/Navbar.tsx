import { Container, Nav, Navbar as Bar } from 'react-bootstrap'
import { NavbarWrapper } from './Navbar.styled'

export default function Navbar() {
    return (
        <NavbarWrapper>
            <Bar expand="lg" className="bg-body-tertiary">
                <Container>
                    <Bar.Brand>Ikea Assignment</Bar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="#products">All Products</Nav.Link>
                        <Nav.Link href="#products/create">Create new Product</Nav.Link>
                    </Nav>
                </Container>
            </Bar>
        </NavbarWrapper>
    )
}
