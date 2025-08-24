/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react'

import { Button, Form } from 'react-bootstrap'

interface Product {
    name: string
    type: number
    colours: number[]
}

export default function CreateProductForm() {
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
                    <option value="1">Type A</option>
                    <option value="2">Type B</option>
                    <option value="3">Type C</option>
                </Form.Control>
            </Form.Group>
            <Form.Group className="mb-3" controlId="colours">
                <Form.Label>Product Name</Form.Label>
                <Form.Control onChange={(e: any) => toggleColour(e)} as="select" multiple className="text-muted">
                    <option value="1">Red</option>
                    <option value="2">Blue</option>
                    <option value="3">Black</option>
                </Form.Control>
            </Form.Group>
            <Button variant="primary" type="submit">
                Submit
            </Button>
        </Form>
    )
}
