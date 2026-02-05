/* Crie uma função chamada verificarAcesso.

Dentro dela:

Crie uma variável let acesso = false

Use um if para:

Se idade ≥ 18, mudar acesso para true

Mostre o valor de acesso dentro da função

Depois:

Tente mostrar acesso fora da função

Observe o erro e entenda por quê

👉 Objetivo: sentir o escopo na pele (essa aula só entra na cabeça assim) */
let idade = 18
function verificarAcesso(idade){
    let acesso = false
    if(idade >= 18){
        acesso= true
        console.log('Acesso liberado')
    }else{
        console.log('Acesso não liberado')
    }
     return acesso
}
console.log(verificarAcesso(idade))
