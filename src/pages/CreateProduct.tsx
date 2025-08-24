import { useEffect, useState } from 'react'
import { getColours, getProductTypes, type Colour, type ProductType } from '../api'
import { CreateProductForm, Title } from '../components'

export default function CreateProduct() {
    const [productTypes, setProductTypes] = useState([] as ProductType[])
    const [colours, setColours] = useState([] as Colour[])

    useEffect(() => {
        getProductTypes().then(setProductTypes)
        getColours().then(setColours)
    }, [])

    if (productTypes.length === 0 || colours.length === 0) {
        return <div>Loading...</div>
    }

    return (
        <>
            <div>
                <Title>Create Product</Title>
                <CreateProductForm productTypes={productTypes} colours={colours} />
            </div>
        </>
    )
}
