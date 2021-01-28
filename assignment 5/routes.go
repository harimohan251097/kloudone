package main

func initializeRoutes() {

	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{

		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
