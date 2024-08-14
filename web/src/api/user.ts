import {post} from "@/api/base";


const module = "/user"


export function APILogin(email: string, password: string) {
    return post(module, "login", {"email": email, "password": password})
}
