# Legal Blog Project

## Next Steps (Backend)
- [x] Rewrite types/models.go
- [x] Create a single DB instance for gorm.db
- [x] Complete Post and Category CRUD services
  - [x] Get all posts
  - [x] Get post by ID
  - [x] Create post
  - [x] Update post
  - [x] Delete post
  - [x] Get all categories
  - [ ] Filter posts by categories
  - [x] Create category
  - [x] Delete category
  - [ ] Sort posts by date (service)
- [x] Use the same DB instance inside services

## Next Steps (Frontend: Homework Integration)
To integrate the client side ("Homework" section) as part of the application, follow the steps below:

### 1. **Frontend Framework Setup**
- [ ] Choose a frontend framework for the project (React, Vue.js, or Svelte recommended). Alternatively, pure HTML/CSS/JS can be used.
  - [ ] If using a framework, set up the environment:
    - [ ] Initialize the project (`npm create-react-app`, `vue create`, or `npm init svelte`).
    - [ ] Set up file structure (components, assets, styles).
  - [ ] If using pure HTML/CSS, organize templates and static assets for easy integration with Go templates.

### 2. **Design & Layout Implementation**
  - **Timeline Page (Home)**:
    - [ ] Create a `Timeline` component/page that contains:
      - [ ] Header with "Dicta" (bold font) and "legal poetry by Harbani Ahuja" (small font).
      - [ ] A scrolling timeline that displays:
        - [ ] Case titles (Name of the case)
        - [ ] Date of the case
        - [ ] Category of the case
      - [ ] Add hover effects for cases:
        - [ ] Show the image of the case when hovered.
        - [ ] Change the color of the case based on its category.
      - [ ] Implement smooth left-right scrolling navigation (either via arrows or scroll gesture).
    - [ ] Use Neurial Grotesk font (Light & Bold) across this page.
    - [ ] Add a subtle page load animation (CSS or JS).

  - **Poem Page (Dynamic Route: `/poem/{name}`)**:
    - [ ] Create a `Poem` component/page that contains:
      - [ ] Header: Display the title and citation of the poem.
      - [ ] Highlight a random extract of the poem to push a narrative.
      - [ ] Display a short summary/context of the case (from the backend).
      - [ ] Display a banner image with navigators to move to the next/previous case.
      - [ ] Embed audio of the artist reading the poem.
      - [ ] Display the poem's text.
      - [ ] Add a link to the artist’s work.
      - [ ] Display a random opinion from the case (without context) at the bottom.
      - [ ] Footer: Links to sidebar pages.

  - **Sidebar Navigation (Global Component)**:
    - [ ] Create a sidebar component with links:
      - [ ] List of poems by categories (basic list style, no fancy CSS).
      - [ ] An "About" page for self-promotion.
      - [ ] A "Donate/Support" page with:
        - [ ] Tax-deductible certificate for donations.
        - [ ] Payment gateway integration.
        - [ ] Display messages from donors.
        - [ ] List top donors.
      - [ ] A "Shop" link that redirects to an external e-commerce page.

### 3. **CSS and Styling**
- [ ] Add global styles:
  - [ ] Use Neurial Grotesk font for all relevant sections (Timeline, Poem, Sidebar).
  - [ ] Standardize spacing, padding, and color scheme across all pages.
- [ ] Implement hover effects and animations as described:
  - [ ] Smooth hover effects for images and color changes on the Timeline.
  - [ ] Subtle entrance animations for key elements (Timeline, Poem Page).
- [ ] Ensure responsive design for mobile and tablet views.

### 4. **Backend API Integration**
- [ ] Connect frontend components to the backend using REST API.
  - **Timeline Page**:
    - [ ] Fetch posts (cases) and their metadata (title, date, category) via API.
    - [ ] Implement category filtering to fetch posts by specific categories.
    - [ ] Ensure API supports sorting posts by date for timeline ordering.
  - **Poem Page**:
    - [ ] Fetch poem details (title, citation, text, summary, opinion) by post ID from the API.
    - [ ] Fetch audio (SoundCloud API integration or locally stored audio).
    - [ ] Fetch banner images associated with the poem/case.

### 5. **Integrating SoundCloud API**
- [ ] Research and implement the SoundCloud API to play audio within the Poem page.
  - [ ] Fetch audio files via SoundCloud API and embed them into the page.
  - [ ] Ensure proper audio player UX (e.g., play/pause, volume controls).
  - [ ] If needed, fallback to HTML5 `<audio>` tags for local audio files.

### 6. **Templating and Rendering**
- [ ] If using Go templating, integrate client-side components with Go templates.
  - [ ] Use `templ` or any Go HTML templating engine to render HTML pages (if not using a JS framework).
  - [ ] Connect Go handlers to serve the static frontend files or pre-rendered templates.
  - [ ] Route backend data into the templates for server-side rendering (SSR) if necessary.

## Next Steps (Backend Enhancements)
- [ ] Setup routes for the new client-side pages (Timeline, Poem, Sidebar).
  - [ ] Route: `/` → Timeline
  - [ ] Route: `/poem/{name}` → Poem Page
  - [ ] Route: `/about`, `/donate`, `/shop` → Static or dynamic pages
- [ ] Write error handling for API calls (e.g., post not found, server errors).
- [ ] Implement caching strategies to optimize API responses for the frontend.
- [ ] Add service for sorting posts by date to enable timeline functionality.

## Next Steps (Deployment & Infrastructure)
- [ ] Set up Dockerfile for containerized deployment.
- [ ] Configure CI/CD pipeline to automate testing and deployment.
- [ ] Test the entire client-server workflow on a local environment.
  - [ ] Ensure the frontend is properly linked to backend services.
  - [ ] Fix any CORS issues that may arise when connecting frontend to backend.

## Refactoring:
- [ ] Interface:
  - [x] AdminService
  - [x] UserService
  - [ ] Investigate further opportunities for refactoring code, especially where frontend-backend communication can be optimized.

## Testing:
- [ ] Write unit tests for backend services.
- [ ] Write integration tests to verify frontend-backend functionality (e.g., post fetching, filtering).
- [ ] Add end-to-end tests to simulate real user interactions with the timeline and poem pages.

## Final Steps (Shipping):
- [ ] Perform full testing on staging environment to ensure all frontend and backend services work as intended.
- [ ] Fix any last-minute bugs or design inconsistencies.
- [ ] Deploy to production.
  - [ ] Set up monitoring and error logging for both backend and frontend.
  - [ ] Enable performance optimizations like asset minification, lazy loading of images, and caching.
  - [ ] Track user interaction with analytics tools (optional).

---

