import axios, { AxiosError } from 'axios'

export interface PostNewProductDTO {
    name: string
    product_type_id: number
    colour_ids: number[]
}

export default async function postNewProduct(
    params: PostNewProductDTO,
    abortController?: AbortController,
) {
    const ids: number[] = []

    try {
        const { name, product_type_id, colour_ids } = params

        colour_ids.forEach(async colour_id => {
            const param = JSON.stringify({ name, product_type_id: product_type_id.toString(), colour_id: colour_id.toString() })
            const { data } = await axios.post(
                'http://localhost:8080/api/products',
                param,
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    signal: abortController?.signal,
                },
            )

            ids.push(data.id)
        })

        return ids
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}

