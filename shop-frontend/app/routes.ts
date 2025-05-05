import {
    type RouteConfig,
    index,
    route,
  } from "@react-router/dev/routes";
  
export default [
  index("./components/home/homePage.tsx"),
  route("/products/:slug", "./components/builder/builderPage.tsx"),
  route("/cart/:cart_id", "./components/cart/cartPage.tsx"),
] satisfies RouteConfig;
