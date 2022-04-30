import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { BASE_URL, TIME_OUT } from './config'
import LocalStorage from './local'

interface IReqCallback {
    callback: (token: string) => void
}

class AxiosRequest {

    public instance: AxiosInstance
    // 防止多次刷新
    private isRefreshing: boolean = false
    // 未处理的请求
    private peedingRequest: Array<IReqCallback> = []

    constructor(config: AxiosRequestConfig) {
        this.instance = axios.create(config)
        // 请求拦截器
        this.instance.interceptors.request.use(
            config => {
                const token = LocalStorage.getToken()
                if (token !== "" && config.headers) {
                    config.headers['Authorization'] = "Bearer " + token
                }
                return config
            },
            err => {
                return err
            }
        )
        // 响应拦截器
        this.instance.interceptors.response.use(
            response => {
                return response
            },
            error => {
                if (!error.response) {
                    return Promise.reject(error)
                }
                if (error.response.status === 401 && !error.config.url.includes('/auth/refresh_token')) {
                    const { config } = error.response
                    // 防止多次刷新
                    if (!this.isRefreshing) {
                        this.isRefreshing = true
                        // 尝试在maxRefresh时间内刷新token
                        return this.refreshToken().then(resp => {
                            // 成功
                            if (resp.data.code === 1000 && resp.data.data) {
                                const token = resp.data.data.token
                                console.log('刷新token: ', token);
                                LocalStorage.setToken(token)

                                // 使用新token重新请求
                                if (config.headers) {
                                    config.headers['Authorization'] = `Bearer ${token}`
                                }

                                // 使用新的token重新请求
                                this.peedingRequest.forEach((req) => {
                                    req.callback(token)
                                })
                                this.peedingRequest = []
                                return this.instance(config)
                            }
                        }).catch(err => {
                            // 失败，只能重新登录
                            console.log("token过期，必须重新登录！");
                            LocalStorage.removeToken()
                            LocalStorage.removeUsername()
                            setTimeout(() => {
                                window.location.replace('/login')
                            }, 2500);
                            return Promise.reject(err)
                        }).finally(() => {
                            this.isRefreshing = false
                        })
                    } else {
                        return new Promise(resolve => {
                            this.peedingRequest.push({
                                callback: (token: string) => {
                                    if (config.headers)
                                        config.headers['Authorization'] = `Bearer ${token}`
                                    resolve(this.instance(config))
                                }
                            })
                        })
                    }
                } else if (error.response && error.response.status === 403) {
                }
                return Promise.reject(error)
            }
        )
    }

    refreshToken() {
        interface IToken {
            token: string
            expire: Date
        }
        interface IResponseData {
            code: number,
            msg: string,
            data?: IToken
        }

        return this.get<IResponseData>("/auth/refresh_token", {
            headers: { "Authorization": `Bearer ${LocalStorage.getToken()}` }
        })
    }

    async request<T>(config: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.instance.request<any, any>(config)
    }

    get<T>(url: string, config?: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.request<T>({ url: url, ...config, method: 'GET' })
    }

    post<T>(url: string, config?: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.request<T>({ url: url, ...config, method: 'POST' })
    }

    delete<T>(url: string, config?: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.request<T>({ url: url, ...config, method: 'DELETE' })
    }

    put<T>(url: string, config?: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.request<T>({ url: url, ...config, method: "PUT" })
    }

    patch<T>(url: string, config?: AxiosRequestConfig<any>): Promise<AxiosResponse<T>> {
        return this.request<T>({ url: url, ...config, method: 'PATCH' })
    }
}

const request = new AxiosRequest({
    baseURL: BASE_URL,
    timeout: TIME_OUT,
    headers: {
        "Content-Type": "application/json",
    },
})

export default request

