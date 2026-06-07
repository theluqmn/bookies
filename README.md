# Bookies - a bookworm's heaven

> [!NOTE]
> For now, Bookies have pivoted towards hosting user-written essays instead of
reviews of books. This project is discontinued and I'm working on something similar but more refined :)

- Write essays for other users to read and rate!
- Earn points!

Check out our [to-do list](https://github.com/theluqmn/bookies/blob/main/TODO.md)

## Guide

1. Download the [latest release](https://github.com/theluqmn/bookies/releases) - only for Windows.
2. Run the main.exe file.
3. Open [localhost:6969/](http://localhost:6969) to view the landing page, click login/signup to create an account
4. You will receive a cookie to prove its you when you log in
5. You may check the database (its a .sqlite file) to check the data
6. That's really it.

## Technical

API routes:

> [!NOTE]
> The API is designed to return HTML for the HTMX frontend to simplify the
system, hence it does not have a traditional REST API that is usable for the public.

- Frontend
  - `/` - landing page
  - `/signup` - signup page
  - `/login` - login page
  - `/essays` - all essays
- Backend
  - `GET /` - home
  - `POST /api/signup` - signup form
  - `POST /api/login` - login form
  - `POST /api/essays` - essay creation form
  - `GET /api/essays` - fetch all essays
  - `GET /api/essays/user` - fetch all essays from a specific user

Database schema:

- `users(id, name, password)`
- `essays(id, language, title, author, content, meta)`

Tech specifications:

- languages: Go, HTML
- libraries: [HTMX](https://htmx.org), [echo](https://echo.labstack.com),
[Tailwind CSS](https://tailwindcss.com/)
- server: [HackClub Nest](https://hackclub.app)

## Notes

Developed by [Luqman](https://theluqmn.github.io). Licensed under the MIT License.
