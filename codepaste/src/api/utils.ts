
export interface IResponseData<T> {
    code: number,
    msg: string,
    data?: T
}

export interface IResponseDataPagination<T> extends IResponseData<T> {
    page: number,
    page_size: number,
    total: number
}

export interface IUserInfo {
    id: number,
    username: string,
    nickname: string,
    created_at: Date
}

export interface IUserTokenInfo {
    token: string,
    expire: Date
}

export interface IPasteInfo {
    id: number,
    key: string,
    title: string,
    lang: string,
    password: string,
    private: boolean,
    content: string,
    created_at?: Date,
    updated_at?: Date,
}

export interface IResultPaste {
    mkey: string,
    title: string,
    password: string,
    lang: string,
    private: boolean,
}

const copyToClipboard = (text: string) => {
    try {
        (async () => {
            await navigator.clipboard.writeText(text)
        })()
    } catch (error) {
        console.log(error);
    }
}

const supportExpireTime = [
    ["0", "阅后即焚"],
    ["180", "三分钟"],
    ["600", "十分钟"],
    ["360", "一小时"],
    ["86400", "一天"],
    ["604800", "一周"],
    ["31536000", "一年"],
    ["-1", "永久"]
]

const supportLangList = [
    ["text", ["text", "纯文本"]],
    ["c_cpp", ["cpp", "C/C++"]],
    ["java", ["java", "Java"]],
    ["golang", ["go", "Golang"]],
    ["python", ["python", "Python"]],
    ["rust", ["rust", "Rust"]],
    ["ruby", ["ruby", "Ruby"]],
    ["php", ["php", "PHP"]],
    ["sql", ["sql", "SQL"]],
    ["javascript", ["javascript", "JavaScript"]],
    ["markdown", ["markdown", "Markdown"]],
    ["json", ["json", "JSON"]],
    ["html", ["html", "HTML"]]
]

const supportLangMap = new Map(supportLangList.map((val) => [val[0], val[1]]))

export { copyToClipboard, supportLangList, supportLangMap, supportExpireTime }