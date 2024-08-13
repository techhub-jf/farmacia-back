# farmacia-back

## Rodando local

Configure seu arquivo `.env`, pode usar como exemplo o arquivo `.env_template`, para iniciar as variáveis de ambiente.

Para iniciar a API em seu local, use o comando `make start`.


## Observações

Se estiver usando VS Code, já está disponível um arquivo `.launch.json` que possui as configurações para depurar o projeto.

Para padronizar o projeto rode o comando `make lint`.


## Lembre de colocar as rotas do swagger como public para testar
## Criar documentação Swagger, se for uma pasta nova precisa incluir ela nesse script
swag init --dir .\,..\app\gateway\api,..\app\domain\entity, ..\app\domain\usecase
