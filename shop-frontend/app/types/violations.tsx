import type { Part } from "./part"

export interface Violation {
    product_id: string
	category_id: string
	principal: string
	resource_id: string
	effect: string
	message: string
	allowed: Part[]
}