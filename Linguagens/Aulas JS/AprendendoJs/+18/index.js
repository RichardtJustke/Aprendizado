let nome = prompt("Digite seu nome completo:");
let idade = parseInt(prompt("Digite sua idade:"));

function idadeMaiorQue18(idade){
    if (idade >= 18){
    alert("Voce tem mais de 18 anos" );
    }else {
    alert("Voce não tem 18 anos ");
    }
}
    
idadeMaiorQue18(idade);
/*
👉 Nível: básico
Função recebe a idade e deve retornar "Maior de idade" ou "Menor de idade".

Objetivo: função + condicional.
*/