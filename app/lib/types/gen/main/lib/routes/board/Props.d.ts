export type Props = {
    Error: string
    Expired: boolean
    LoggedIn: boolean
    Articles: Article[]
    HasMore: boolean
    Page: number
}

export type ArticlesArticle = {
    AccountID: string
    CreatedAt: number
    Content: string
    Title: string
    ID: string
}
