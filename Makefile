prepare:
	echo "making gql definitions"
	make gql
	echo "generating dataloaders"
	make dataloaders

gql:
	go run github.com/99designs/gqlgen generate 2> /dev/null || true # ignoring first run

dataloaders:
	cd graph/model && go run github.com/vektah/dataloaden CityLoader int '*github.com/NunChatSpace/gqlgen-test/graph/model/model.City'
	cd graph/model && go run github.com/vektah/dataloaden StateLoader int '*github.com/NunChatSpace/gqlgen-test/graph/model/model.State'
	cd graph/model && go run github.com/vektah/dataloaden StatesByCountryLoader int '[]*github.com/NunChatSpace/gqlgen-test/graph/model/model.State'
	cd graph/model && go run github.com/vektah/dataloaden CitiesByStateLoader int '[]*github.com/NunChatSpace/gqlgen-test/graph/model/model.City'
	cd graph/model && go run github.com/vektah/dataloaden CountryLoader int '*github.com/NunChatSpace/gqlgen-test/graph/model/model.Country'