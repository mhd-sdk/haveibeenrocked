# This is a brainstorming document, where i wrote down all the ideas i had for the project

## Not implemented
- Hosted Docker images on GitHub with automated build pipelines on pull requests and/or releases
- integration tests using Testcontainers and gherkin
- dependabot for "secure by design"
- Open api Documentation
- Animated wizz with rock music when password matches the database list (msn nostalgia)


## Implemented
- Implement k-anonymity
- automated download of passwords file with mage
- repositories pattern
- Formatting and linting with Prettier/ESLint
- Unit tests
- Use tailwind as it is pretty hyped (I usually use styled component but some changes would be fun) 
- Use gitmojis (or conventionnal commits convention)
- Rock music when password matches the database list
- Modular docker images with a dev and a prod compose environment using environment variables
- Nginx reverse proxy
- The password database is pretty big; having an in-memory cache like Redis can speed up frequent password requests (yes i could easily cache in my go program, but lemme flex my skills pls !)
- Ollama for some AI hype, request llm to give a 5 star rating of the password strength.
- Use bun for the frontend, as it is a new and fast package manager (and i wanted to try it out)