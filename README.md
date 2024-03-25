# Backend Golang Example

This repo contains a example the backend in golang
This example contains the features below: 
* CRUD - Register Vehicle
    - API
    - Database Conection (Postgres)
* Env Files 
* 

# PROJECT ORGANIZATION 

/your-backend-name
    /cmd
        /api           # Entrypoint da aplicação
    /internal
        /vehicle       # Lógica do domínio específico para veículos
        /repository
        /handler
        /model
        /service
        /storage       # Implementações de interface de armazenamento (ex.: database)
    /pkg
        /response      # Utilidades para respostas HTTP
        /httperror     # Manipulação de erros HTTP
    /config            # Configurações da aplicação (ex.: banco de dados)
    /.env              # Arquivo de variáveis de ambiente
    /scripts           # Scripts de build e setup
    /.gitignore        # Arquivos e diretórios ignorados pelo Git
    /Dockerfile        # Para containerização da aplicação
    /README.md         # Documentação da API
    /go.mod            # Dependências do Go
    /go.sum


# Setup Go 
## MAC - Install with brew
* Befare setup go use this reference to setup dev environment - [scripts](https://github.com/davipeterlini/scripts)
```shell script
brew install go
go version
```

# Start with golang (only first)
```shell script
# its not necessary to use, only use in create the new project
go mod init 
```

# Adding Libs 
## API
```shell script
go get . # update all
go get github.com/gorilla/mux # router
```

## Env File
```shell script
go get . # update all
go get github.com/joho/godotenv # env files.
go get -u github.com/joho/godotenv
```

## Database
```shell script
go get github.com/lib/pq #postgres
```

# Update Libs
```shell script
go get . # update all
go get -u github.com/<dependency> # update specific dependency
```

# Dev - Local
## Start Database in manual
```shell script
brew install --cask rancher-desktop #brew install --cask docker
open -a "Rancher Desktop" # Open GUI for start
docker run hello-world # Test
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
docker ps 
```

# Run Project
```shell script
go run cmd/api/main.go
```















# Debug - Dev
```shell script
go install github.com/go-delve/delve/cmd/dlv@latest
```



# Debug in Golang
## Install Delve
```shell script
go install github.com/go-delve/delve/cmd/dlv@latest
```

## Config Debug with VScode
* Open Project in Go
* Set de Breakpoint in Line
* Open Run and Debug in VSCode
* Click in "Create a launch.json --> GO --> Go Launch Package
* Install envFiles - DotENV
```shell script

```
* Open file launch.json in folder .vscode
# Put the code 
```shell script
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Main",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/api/main.go",
            "envFile": "${workspaceFolder}/.env",  // Especifica o caminho para o seu arquivo .env
            "args": []
        }
    ]
}
```






Altere a service.go para se adequar a nova estrutura
Considere trazer o código completo

Altere a handler.go para se adequar a nova estrutura
Considere trazer o código completo

Altere a main.go (cmd/api) para se adequar a nova estrutura e para que seja possível executar uma chamada da API e o dado ser salvo no banco 
Verifique se o banco está ativo se não estiver faça a chamada shell para subir o docker do banco (docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres). Lembrando que isso só deve ser feito nos testes locais, por isso adicione o .env uma variável para indicar essa informação 
Após a verificação do banco verifique se as tabelas da aplicação foram cridadas (de acordo com as entidades) caso não crie por default 
Para a criação das tabelas use um arquivo SQL na raiz do projeto 
Considere trazer o código completo do main.go

Crie um shell script para gerar o binário da aplicação 

Execute o binário gerado em uma estrutura local de kubernets + postgresql tudo com o rancher 


Ao iniciar a a aplicação cria a tabela vehicles por default 
Considere colocar sql para a criação de tabela em um arquivo, onde estarão todos os SQLs de criação de tabela
Verifique se a tabela existe e Caso exista não deve ser criada novamente 



# Config Kubernets Cluster 
```shell script
brew install kubectl
brew install minikube
minikube start
brew install helm
helm repo add rancher-latest https://releases.rancher.com/server-charts/latest
kubectl get pods --namespace cattle-system
minikube service list
```


# Update dependencie
```shell script

go get .
```

# Add dependencie
```shell script
go get <DEPENDENCY>
```

# Run tests
```shell script
go run poc.go
```

# Run build.sh
## Generate Binary
```shell script
./build.sh build linux
```




## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://gitlab.com/davipeterlini/backend-sales-car.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://gitlab.com/davipeterlini/backend-sales-car/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Set auto-merge](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing (SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thanks to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.



