import { Table } from 'react-bootstrap'
import { getProductById, getProducts, type Product } from '../api'
import { useEffect, useState } from 'react'
import type { ProductList } from '../api/products/getProducts'
import { Title } from '../components'

async function populateProductList(products: ProductList[]) {
    const productList = [] as Product[]

    for (let i = 0; i < products.length; i++) {
        productList.push(await getProductById(products[i].id))
    }

    return productList
}

const BORDER_WIDTH = '0px 12px'
const PADDING = '6px 18px'
const WHITE = '#ffffff'
const BG_COLOUR = '#f2f2f2'
const STYLE = { background: WHITE, borderWidth: BORDER_WIDTH, borderColor: BG_COLOUR, padding: PADDING, textAlign: 'center' }
const HEADER_STYLE = { color: WHITE, fontSize: 'large', textAlign: 'left', minWidth: '200px' }

export default function ProductsList() {
    const [products, setProducts] = useState([] as Product[])

    useEffect(() => {
        getProducts().then((list) => {
            populateProductList(list).then(setProducts)
        })
    }, [])

    if (!products?.length && products.length === 0) {
        return <div>Loading...</div>
    }

    return (
        <>
            <div style={{ maxHeight: '65vh', overflowX: 'scroll' }}>
                <Title>Products List</Title>
                <Table striped borderless hover style={STYLE}>
                    <thead>
                        <tr>
                            <th style={{ ...STYLE, ...HEADER_STYLE, background:"#99a7b8", minWidth: '100px' }}>ID</th>
                            <th style={{ ...STYLE, ...HEADER_STYLE, background:"#8aaafb" }}>Name</th>
                            <th style={{ ...STYLE, ...HEADER_STYLE, background:"#5fb4ef" }}>Type</th>
                            <th style={{ ...STYLE, ...HEADER_STYLE, background:"#35c0e1" }}>Colours</th>
                            <th style={{ ...STYLE, ...HEADER_STYLE, background:"#0dc9d4" }}>Created At</th>
                        </tr>
                    </thead>
                    <tbody>
                        {products.map((product) => (
                            <tr key={product.id}>
                                <td style={STYLE}>{product.id}</td>
                                <td style={{ ...STYLE, textAlign: 'left' }}>{product.name}</td>
                                <td style={STYLE}>{product.product_type}</td>
                                <td style={STYLE}>{product.colours.join(', ')}</td>
                                <td style={STYLE}>{new Date(product.created_at).toLocaleString()}</td>
                            </tr>
                        ))}
                    </tbody>
                </Table>
            </div>
        </>
    )
}
