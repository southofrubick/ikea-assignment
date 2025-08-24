import { Container, Nav, Navbar as Bar } from 'react-bootstrap'
import { NavbarWrapper } from './Navbar.styled'

export default function Navbar() {
    return (
        <NavbarWrapper>
            <Bar expand="lg" className="bg-body-tertiary" style={{ paddingBottom: '0px', paddingTop: '24px' }}>
                <Container style={{ width: '70vw', margin: '0 10vw' }}>
                    <Bar.Brand style={{ fontFamily: 'Lato', fontStyle: 'italic', fontSize: 'x-large' }}>Ikea Assignment</Bar.Brand>
                    <Nav className="me-auto" variant="tabs" defaultActiveKey="#products">
                        <Nav.Item>
                            <Nav.Link href="#products">All Products</Nav.Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Nav.Link href="#products/create">Create new Product</Nav.Link>
                        </Nav.Item>
                    </Nav>
                </Container>
            </Bar>
        </NavbarWrapper>
    )
}
