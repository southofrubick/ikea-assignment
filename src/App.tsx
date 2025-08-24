import AppRoutes from './AppRoutes'
import { Navbar } from './components'
import { ContentWrapper } from './App.styled'

function App() {
    return (
        <div style={{ fontFamily: 'Josefin Sans, sans-serif' }}>
            <Navbar />
            <main>
                <ContentWrapper>
                    <AppRoutes />
                </ContentWrapper>
            </main>
        </div>
    )
}

export default App
