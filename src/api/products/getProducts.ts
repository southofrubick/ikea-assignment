import axios, { AxiosError } from 'axios'

export interface ProductList {
    id: number
    name: string
    createdAt: string
    updatedAt: string
}

export default async function getProducts(
    abortController?: AbortController,
) {
    try {
        const { data } = await axios.get(
            'http://localhost:8080/api/products',
            {
                signal: abortController?.signal,
            },
        )

        return data as ProductList[]
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}
