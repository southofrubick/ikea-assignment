import axios, { AxiosError } from 'axios'

export interface ColourType {
    id: number
    name: string
    createdAt: string
    updatedAt: string
}

export default async function getColours(
    abortController?: AbortController,
) {
    try {
        const { data } = await axios.get(
            'http://localhost:8080/api/products/colours',
            {
                signal: abortController?.signal,
            },
        )

        return data as ColourType[]
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}
