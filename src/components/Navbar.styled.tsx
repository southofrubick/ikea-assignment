import styled from 'styled-components'

export const NavbarWrapper = styled.div`
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    z-index: 2;
    margin-bottom: 20px;
    border-width: 0 0 1px;
    box-sizing: border-box;
    user-select: none;
`

export const NavbarContainer = styled.div`
    display: flex;
    align-items: center;
    justify-content: space-between;
    max-height: 60px;
    padding: 12px 50px;
`
