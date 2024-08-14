import axios, {type AxiosResponse} from 'axios'


export function getHeaders() {
    return {"Authorization": "Bearer " + import.meta.env.VITE_API_SECRET_KEY}
}

export async function get(module: string, url: string, params: any) {
    const response: AxiosResponse<any> = await axios.get(
        module + "/" + url,
        {headers: getHeaders(), params: params}
    )
    return response.data
}

export async function post(module: string, url: string, params: any) {
    const response: AxiosResponse<any> = await axios.post(
        module + "/" + url, params,
        {headers: getHeaders()}
    )
    return response.data
}
