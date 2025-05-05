import type { Part } from "./part"

export interface Cart {
    uuid:    string           
	status:   string           
    products: { [slug: string]: Part[]}
}
