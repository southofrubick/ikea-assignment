import { useEffect, useState } from 'react'
import { getColours, getProductTypes, type ColourType } from '../api'
import { CreateProductForm } from '../components'

export default function CreateProduct() {
    const [productTypes, setProductTypes] = useState([] as ProductType[])
    const [colours, setColours] = useState([] as ColourType[])

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
                <h3>CreateProduct</h3>
                <CreateProductForm productTypes={productTypes} colours={colours} />
            </div>
        </>
    )
}
