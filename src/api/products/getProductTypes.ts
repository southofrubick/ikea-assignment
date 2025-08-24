import axios, { AxiosError } from 'axios'

export interface ProductType {
    id: number
    name: string
    createdAt: string
    updatedAt: string
}

export default async function getProductTypes(
    abortController?: AbortController,
) {
    try {
        const { data } = await axios.get(
            'http://localhost:8080/api/products/types',
            {
                signal: abortController?.signal,
            },
        )

        return data as ProductType[]
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}
