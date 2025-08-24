import axios, { AxiosError } from 'axios'

export interface Product {
    id: number
    name: string
    type: number
    colours: string[]
    product_type: string
    created_at: string
    updated_at: string
}

export default async function getColours(
    id: number,
    abortController?: AbortController,
) {
    try {
        const { data } = await axios.get(
            `http://localhost:8080/api/products/${id}`,
            {
                signal: abortController?.signal,
            },
        )

        return data as Product
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}
