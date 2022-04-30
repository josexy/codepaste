
let BASE_URL = ''

const TOKEN_KEY = "jwt_token"
const USERNAME = "_username"

const TIME_OUT = 5000

const PRODUCTION_URL = "/api/v1"
const DEPLOYMENT_URL = "http://127.0.0.1:10086/api/v1"

if (process.env.NODE_ENV === 'development') {
    BASE_URL = DEPLOYMENT_URL
} else if (process.env.NODE_ENV === 'production') {
    BASE_URL = PRODUCTION_URL
} else {
    BASE_URL = DEPLOYMENT_URL
}

export { BASE_URL, TIME_OUT, TOKEN_KEY, USERNAME }
