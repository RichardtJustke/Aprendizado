let desconto = parseFloat(prompt("Digite o valor do desconto (em %):"));
let valorProduto = parseInt(prompt("Digite o valor do produto:"));

function calcularDesconto(valorProduto, desconto){
    let porcentagem = (desconto/100);
    let economia = (valorProduto*porcentagem);

    return valorProduto - economia;
    
}

let resultado  = calcularDesconto(valorProduto, desconto);
alert("O valor do desconto é: R$ " + resultado.toFixed(2));
/*
Escreva uma função calcularDesconto(valor, desconto) que retorna o valor final após aplicar o desconto.

Exemplo: calcularDesconto(100, 10) → 90

Objetivo: lógica simples + porcentagem.
*/