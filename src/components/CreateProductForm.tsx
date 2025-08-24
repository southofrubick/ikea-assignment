import React from 'react'

import { Form } from 'react-bootstrap'

export default function CreateProductForm() {
    return (
        <Form>
            <Form.Group className="mb-3" controlId="formProductName">
                <Form.Label>Product Name</Form.Label>
                <Form.Control type="text" placeholder="Enter product name" />
                <Form.Text className="text-muted">
                    Please enter the name of the product.
                </Form.Text>
            </Form.Group>
            <Form.Group className="mb-3" controlId="formProductType">
                <Form.Label>Product Name</Form.Label>
                <Form.Control type="text" placeholder="Enter product type" />
                <Form.Select className="text-muted">
                    Please enter the name of the product.
                </Form.Select>
            </Form.Group>
            <Form.Group className="mb-3" controlId="formProductColour">
                <Form.Label>Product Name</Form.Label>
                <Form.Control multiple type="text" placeholder="Enter product name" />
                <Form.Select className="text-muted">
                    Please enter the name of the product.
                </Form.Select>
            </Form.Group>
        </Form>
    )
}
