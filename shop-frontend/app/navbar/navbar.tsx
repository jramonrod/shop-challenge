import { NavLink } from "react-router";

export function Navbar() {
  return (
    <nav className="bg-primary px-4 py-2 flex justify-between items-center">
      <div className="text-xl font-bold">
        <NavLink to="/">Factorial Sports Shop</NavLink>
      </div>
      <ul className="flex space-x-6">
        <li>
          <NavLink to="/" className="hover:text-blue-600">Home</NavLink>
        </li>
      </ul>
    </nav>
  )
}
