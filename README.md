## Ideas
- Hosted Docker images on GitHub with automated build pipelines on pull requests
- Animated wizz when password matches the database list (msn nostalgia)
- Rock music when password matches the database list
- Formatting and linting with Prettier/ESLint
- Unit tests + integration tests using Testcontainers
- Use gitmojis (or conventionnal commits convention)
- Try out v0 for initial frontend design
- use tailwind as it is pretty hyped (usually use styled component but some changes would be fun)

## Notes
- I chose to make a multi-stage image for the frontend, first layer for building Vite, second for lightweight production-ready
- Over-engineered, but Nginx reverse proxy allows caching, load balancing, static serving, SSL, and many more. This choice is for scalability.
- Also overengineered, The password database is pretty big; having an in-memory cache like Redis can speed up frequent password requests (yes i could easily cache in my go program, but lemme flex a bit pls ! i really want this job)
- For now, passwords will be loaded at API startup, skipped if the database already contains it. Also, I use Postgres because the job is asking for it :)
- Ollama for some AI hype, mainly used for rating passwords on a 5 stars scale.
- I won't use branching strategies as i'm alone on the project and it will never be deployed in prod, but i'm pretty familliar with the trunk-based development

![Description de l'image](docs/haveibeenrocked-architecture.png)


