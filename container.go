package di2

import (
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Genialngash/graphql-go-test/graph"
)

var (
	container *dig.Container
	once      sync.Once
)

// GetContainer :
func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

// GraphQLHandlerInjector :
func GraphQLHandlerInjector(config *models.Configuration) *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return srv
}

// BuildContainer :
func buildContainer() *dig.Container {
	container := dig.New()
	handlerContainerErrors(
		container.Provide(config.NewConfiguration),
		container.Provide(repositories.NewBootstrapRepository),
		container.Provide(controllers.NewBootstrapController),
		container.Provide(controllers.NewHealthController),
		container.Provide(GraphQLHandlerInjector))

	return container
}
func handlerContainerErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			logger.ConditionalFatal("container", "buildContainer ", err)
		}
	}
}
