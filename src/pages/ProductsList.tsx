import { Table } from 'react-bootstrap'

export default function ProductsList() {
    return (
        <Table striped bordered hover>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Typre</th>
                    <th>Colours</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>1</td>
                    <td>Product 1</td>
                    <td>Type A</td>
                    <td>Red, Blue</td>
                </tr>
                <tr>
                    <td>2</td>
                    <td>Product 2</td>
                    <td>Type B</td>
                    <td>Green, Yellow</td>
                </tr>
                <tr>
                    <td>3</td>
                    <td>Product 3</td>
                    <td>Type A</td>
                    <td>Black, White</td>
                </tr>
            </tbody>
        </Table>
    )
}
