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
    let id: number

    try {
        const { name, product_type_id, colour_ids } = params

        for await (let colour_id of colour_ids) {
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

            if (!id) {
                id = data.id
            }
        }

        return id
    } catch (error) {
        const axiosError = error as AxiosError

        console.log(
            `Failed to get product types: ${axiosError.message}`,
            axiosError,
        )

        throw axiosError
    }
}

