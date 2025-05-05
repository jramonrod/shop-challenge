import type { Violation } from "./violations"

export interface VerifyResults {
    is_valid: boolean
    violations: Violation[]
}
