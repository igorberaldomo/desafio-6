Exercicio:
    tem como objetivo a implementação de uma função que dá um tempo de leilão de um produto assim como uma go routine em internal/infra/database/auction/create_auction.go para verificar se o leilão ainda está acontecendo
    

    foi criado um main_test.go para rodar os testes para verificar se a criação de um leilão ocorreu corretamente e se caso o leilão tenha passado o tempo permitido não seja permitido a colocação de novos lances
    para executar os testes execute:
        go test cmd/auction/main_test.go

para executar use:
    docker-compose -f 'docker-compose.yml' up -d --build 

    para descobrir quais auctions estão funcionando:
        envie um metodo GET para http://localhost:8080/auction/
    para ver status de uma auction:
        envie um metodo GET para http://localhost:8080/auction/:auctionId
    para iniciar uma auction:
        envie um metdo POST para http://localhost:8080/auction/
        usando { 
            ProductName: string , 
            Category: string, 
            Description: string, 
            Condition:  ProductCondition }
    para ver o vencedor:
        envie um metodo GET para http://localhost:8080/auction/winner/:auctionId
    para enviar uma BID:
        envie um metodo POST para http://localhost:8080/bid/
        usando {
            UserId    string,
	        AuctionId string, 
	        Amount    float64
        }
    para pegar informação sobre todos os BIDs:
        envie um metodo GET para http://localhost:8080/bid/:auctionId
    para pegar informações sbre o usuario
        envie um metodo GET para http://localhost:8080/user/:userId