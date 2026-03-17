#!/usr/bin/env node
import gradient from "gradient-string"

console.log(gradient(['red', 'orange']).multiline(`

████████╗███████╗███████╗████████╗███████╗     ██████╗██╗     ██╗
╚══██╔══╝██╔════╝██╔════╝╚══██╔══╝██╔════╝    ██╔════╝██║     ██║
   ██║   █████╗  ███████╗   ██║   █████╗      ██║     ██║     ██║
   ██║   ██╔══╝  ╚════██║   ██║   ██╔══╝      ██║     ██║     ██║
   ██║   ███████╗███████║   ██║   ███████╗    ╚██████╗███████╗██║
   ╚═╝   ╚══════╝╚══════╝   ╚═╝   ╚══════╝     ╚═════╝╚══════╝╚═╝
`)
)

//npm install chalk figlet gradient-string
//https://www.npmjs.com/package/gradient-string(configurações do gradient-string)
//baixar para deixar bonito ksksksks
//ir no site ("https://patorjk.com/software/taag/")
//usar a font ANSI Shadow
//colocar texto
const command = process.argv[2]
if (!command) {
  console.log(`
┌─────────────────────────────┐
│   Bem Vindo ao TESTE CLI    │
└─────────────────────────────┘
Aqui é um ambiente de testes onde vou usar para depois fazer meu portifolio abaixo temos
alguns comandos que você poderá dar no terminal junto com o comando base (ex1)
Command:
testepage1
testepage2
`)
}
