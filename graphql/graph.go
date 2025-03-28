package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	/* accountClient *account.Client
	catalogClient *catalog.Client
	OrderClient   *order.Client
	*/

}

func (s *Server) Account() AccountResolver {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Mutation() MutationResolver {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Query() QueryResolver {
	//TODO implement me
	panic("implement me")
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	/* 	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}
	catalogClient, err := account.NewClient(catalogUrl)
	if err != nil {
		accountClient.close()
		return nil, err
	}
	orderClient, err := account.NewClient(orderUrl)
	if err != nil {
		accountClient.close()
		catalogClient.close()
		return nil, err
	}

	*/
	return &Server{
		/* 	accountClient,
		catalogClient,
		orderClient,

		*/
	}, nil

}

/*
	func (s *Server) Mutation() MutationResolver {
		return &mutationResolver{
			server: s
		}
	}

	func (s *Server) Query() QueryResolver {
		return &queryResolver{
			server: s
		}
	}

	func (s *Server) Account() QueryResolver {
		return &accountResolver{
			server: s
		}
	}
*/
func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(
		Config{
			Resolvers: s,
		})
}
