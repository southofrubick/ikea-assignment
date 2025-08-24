/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react'

import { Button, Form } from 'react-bootstrap'
import type { ColourType, ProductType } from '../api'

interface Product {
    name: string
    type: number
    colours: number[]
}

interface CreateProductFormProps {
    productTypes?: ProductType[]
    colours?: ColourType[]
}

export default function CreateProductForm({
    productTypes,
    colours,
}: CreateProductFormProps) {
    const [product, setProduct] = useState({name: '', type: 0, colours: []} as Product)

    const updateProduct = (event: Event) => {
        const { target } = event

        if (!(target instanceof HTMLInputElement || target instanceof HTMLSelectElement)) {
            return
        }

        const field = target.id
        let value: string | number = target.value

        if (field === 'type') {
            value = parseInt(value)
        }

        setProduct({
            ...product,
            [field]: value,
        })
    }

    const toggleColour = (event: Event) => {
        const { target } = event

        if (!(target instanceof HTMLSelectElement)) {
            return
        }

        setProduct({
            ...product,
            colours: [].slice
                .call(target.selectedOptions)
                .map(option => parseInt(option.value))
        })
    }

    return (
        <Form onSubmit={() => console.log(product)}>
            <Form.Group className="mb-3" controlId="name">
                <Form.Label>Product Name</Form.Label>
                <Form.Control onChange={(e: any) => updateProduct(e)} type="text" placeholder="Enter product name" />
                <Form.Text className="text-muted">
                    Please enter the name of the product.
                </Form.Text>
            </Form.Group>
            <Form.Group className="mb-3" controlId="type">
                <Form.Label>Product Name</Form.Label>
                <Form.Control defaultValue={1} onChange={(e: any) => updateProduct(e)} as="select" className="text-muted">
                    {productTypes.map((productType) => (
                        <option value={productType.id} key={productType.id}>{productType.name}</option>
                    ))}
                </Form.Control>
            </Form.Group>
            <Form.Group className="mb-3" controlId="colours">
                <Form.Label>Product Name</Form.Label>
                <Form.Control onChange={(e: any) => toggleColour(e)} as="select" multiple className="text-muted">
                    {colours.map((colour) => (
                        <option value={colour.id} key={colour.id}>{colour.name}</option>
                    ))}
                </Form.Control>
            </Form.Group>
            <Button variant="primary" type="submit">
                Submit
            </Button>
        </Form>
    )
}
