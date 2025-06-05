# projeto api em go arquitetura mvp

## estrutura de pasta

```sh
    SRC  -> arquivos importantes
    |
    |--- Controllers // Entrada e validação de dados
    |
    |--- model // Regra de negócios e objetos principais
    |
    |--- View // Gerenciamento de dados ( publicos ou privados) 
    Converters
    |--- tester // testes de integrações da aplicação
    |--- configurations // Pacote para conexões e configurações
    |
    etc..
    main.go
    .env
    .gitignore // ignora arquivos sensiveis como o .env e não leva para o repositorio remoto
    .readme.md
```