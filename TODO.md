# Legal Blog Project


## Next steps

- [x] Rewrite types/models.go
- [x] Create a single DB instance to gorm.db
- [ ] Write post and category crud  services
        - [x] Get all posts
        - [x] Get post by id
        - [x] Create post
        - [x] Update post
        - [x] Delete post
        - [x] Get all categories
        - [] Filter posts by categories
        - [x] Create category
        - [x] Delete category
- [x] Use the same DB instance inside services
- [ ] Write Handlers for:
    - [x] Admin
            - [x] Login Handlers
            - [x] GET List of posts handler
            - [x] GET Preview post handler
            - [x] GET List of Categories handler
            - [x] GET Filter Posts by Category handler 
            - [x] PUT Editor Save Changes to post handler
            - [x] POST Add New post from the Editor handler
            - [x] POST New Category from the Category modal handler
            - [x] DELETE post handler
            - [x] DELETE category handler
    - [ ] User
            - [ ] GET List of posts handler
            - [ ] GET post by id handler
- [ ] Add Service to sort posts by date of cases
- [ ] Setup routes for handlers
- [ ] Write a run function inside main file for error logging
- [ ] Create a workable template (using templ) for admin dashboard, preview post, and editor

## Refactoring:

- [ ] Interface:
        - [x] AdminService
        - [x] UserService


## Additional steps

- [ ] Setting up dockerfile
- [ ] Setting up CI/CD pipeline
- [ ] Figure out how to work with images
        - Check ChatGPT response for 19/9/24: Soundclound API for legal blogs
- [ ] Figure out how to work with soundcloud api
- [ ] Get markdown to work with editor and the application in general
