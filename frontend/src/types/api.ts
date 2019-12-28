export interface Auction {
    id: string,
    active: boolean,
    title: string,
    short_description: string,
    description: string,
    start_price: number,
    minimal_bid: number,
    actual_price: number,
    start_date: Date,
    end_date: Date,
    images: Image[],
    bids: Bid[],
    created_at: Date,
}

export interface Image {
    id: number,
    url: string,
    auction_id: string,
    created_at: Date,
}

export interface Bid {
    id: number,
    bid: number,
    uid: string,
    auction_id: string,
    created_at: Date,
}