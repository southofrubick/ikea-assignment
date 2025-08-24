import React from 'react'

import AppRoutes from './AppRoutes'
import { Navbar } from './components'
import { ContentWrapper } from './App.styled'

function App() {
    return (
        <>
            <Navbar />
            <main>
                <ContentWrapper>
                    <AppRoutes />
                </ContentWrapper>
            </main>
        </>
    )
}

export default App
