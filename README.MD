# Customizable Shop Coding Challenge
# Introduction
Dear readers my name is Jose Rodriguez and I built this shop. It is a flexible, data-driven platform for selling configurable sports products. 
The goal is to enable selling bicycles today; skis, surfboards, etc. tomorrow.

# Trade-offs
## Software Design DDD & Pragmatism
* I chose not to follow by the book DDD efforts because I think it would have been overkill here. For example, I chose to have a single /domain/ directory with the file names acting as the subdomain contexts
* >[!NOTE] I think it's worthwhile to mention that when necessary I did implement some DTOs and mapper classes
  * Similarly the /db/ and /api/ folders could've been infrastructure and interfaces folders respectively to provide a layer of abstraction and interchangeable parts
  * Domain objects as well should not be exposed to the http layer but we add in the json tags and try to keep the code lightweight and non-complex
    * This also extends for writing mapper classes for talking between the infrastructure and the domain. For example, db <-> domain dedicated adapters were not created
* For some of the migrations I have written "unnecessary" at the top. These exist to simplify my life doing local development and to avoid developing an admin panel for creating a state where the app is useable
## Product
* I have decided to not include a typical user creation flow this was with a couple of motivations
  * I want faster tear up and tear down for me to iterate on the challenge, which I can do by simulating a database state useful for it
  * User credentials and session/key management doesn't seem to be the focus of the challenge
* One with a keen eye can already tell that an admin panel for the business owner is missing from the shop. I decided to keep the focus on the core shopping experience: product builder, customizer, cart, and UI.
    * Topics such as authentication and authorization would creep up once we introduce user and admin logins for example.

## Technology Choices
### Backend
* Go + GORM + Postgres
* No OpenAPI codegen
  * I opted for handwritten handlers/models and controllers to stay in my comfort zone and reduce dependencies. In a larger team, I’d adopt an API-first approach with OpenAPI and code generation
### Frontend
> [!NOTE]
> As I mentioned in previous rounds of the interview process, I only have some small years of experience writing front-end code. Apologies if you spot any patterns that might upset you this is a whole world I am still delving into ;). If for nothing else, the front-end part of this challenge can be seen as my ability to use my knowledge + AI to make something somewhat functional.
- Vite
- React Router
-  TailwindCSS (DaisyUI)

## Other Things Left Out of this Challenge
To keep the scope of this challenge focused there are a couple more things left out of the challenge that would be a requirement for a production system:
* Structured Logging
* Monitoring
* Testing
* CI/CD
* Infra

# Referencing Specific Challenge Requirements
Please find answers to the challenge specification in the ./challenge-requirements folder. The following is the location to the requirements: 
| Requirement                  | Location            |
|------------------------------|---------------------|
| Data Model chosen            | `data-model`        |
| Main user actions            | `main-user-actions` |
| Product page                 | `main-user-actions` |
| Add-to-Cart action           | `cart`              |
| Administrative workflows     | `admin-workflows`   |
| New product creation         | `new-product-parts` |
| Pricing rules & calculations | `setting-prices`    |

# How to run locally
## Docker Compose
```bash
cd ./shop-backend
docker-compose build 
docker-compose up -d
```
## Frontend
```bash
cd ./shop-frontend
yarn install
yarn vite
  ```


# About Jose
## Professionally
I'm a software engineer with a strong product mindset and a preference for simple, maintainable solutions. I enjoy taking ownership of problems end-to-end while at the same time making the product and the areas I go through better. I also have a strong proven track record of creating abstractions and tools that can help more than just my immediate team.

## Personally
Outside of code, I greatly enjoy a variety of physical activities like hiking and skiing. I like to game as a hobby whenever my mind craves some puzzling or number crunching activities. I also collect board games and consider myself a very social person.

Aside from that I also code during my freetime, I love developing indie games and contributing to the indie game dev community.
