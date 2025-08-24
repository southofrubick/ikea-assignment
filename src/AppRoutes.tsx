import React, { type JSX, type ReactNode } from 'react'

import { Route, Routes, Outlet } from 'react-router-dom'
import { CreateProduct, ProductsList } from './pages'

type RouteType = {
    path: string
    element: ReactNode
    permissions?: string[]
}

const PUBLIC_ROUTES = [
    {
        path: 'products',
        element: <ProductsList />,
    },
    {
        path: 'products/create',
        element: <CreateProduct />,
    },
]

const buildRoutes = (routes: RouteType[]) => {
    return routes
        .map((route: RouteType) => (
            <Route path={route.path} element={route.element} key={route.path} />
        ))
}

export default function AppRoutes(): JSX.Element {
    return (
        <Routes>
            <Route path="/" element={<Outlet />}>
                {buildRoutes(PUBLIC_ROUTES)}
            </Route>
        </Routes>
    )
}
