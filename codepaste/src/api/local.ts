import { TOKEN_KEY, USERNAME } from "./config"

class LocalStorage {
    static set(key: string, value: string) {
        localStorage.setItem(key, value)
    }
    static get(key: string): string | null {
        return localStorage.getItem(key)
    }
    static remove(key: string) {
        localStorage.removeItem(key)
    }

    static getToken(): string {
        let token = this.get(TOKEN_KEY)
        if (token === undefined || token === null) {
            return ""
        }
        return token
    }
    static setToken(token: string) {
        this.set(TOKEN_KEY, token)
    }
    static removeToken() {
        this.remove(TOKEN_KEY)
    }

    static getUsername(): string {
        let username = this.get(USERNAME)
        if (username === undefined || username === null) {
            return ""
        }
        return username
    }
    static setUsername(username: string) {
        this.set(USERNAME, username)
    }
    static removeUsername() {
        this.remove(USERNAME)
    }
}

export default LocalStorage